package rates

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"golang.org/x/time/rate"
)

type Limit struct {
	Name      string
	Reservoir int
	Interval  time.Duration
}

var MainRates = []Limit{
	{
		Name:      "main-sb",
		Reservoir: 500,
		Interval:  10 * time.Second,
	},
	{
		Name:      "main-lb",
		Reservoir: 30000,
		Interval:  10 * time.Minute,
	},
}

var client *redis.Client
var riotAPIKey string

func InitRedis(redisAddr string) {
	client = redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})
}

func IsAllowed(ctx context.Context, key string, limit int, window time.Duration) (bool, error) {
	luaScript := `
	local current = redis.call("GET", KEYS[1])
	if current and tonumber(current) >= tonumber(ARGV[1]) then
			return 0
	else
			return 1
	end
	`
	result, err := client.Eval(ctx, luaScript, []string{key}, limit, int(window.Seconds())).Result()
	if err != nil {
		return false, err
	}

	return result.(int64) == 1, nil
}

func incrementRate(ctx context.Context, key string, limit int, window time.Duration) (bool, error) {
	luaScript := `
	local current = redis.call("GET", KEYS[1])
	if current and tonumber(current) >= tonumber(ARGV[1]) then
			return 0
	else
			redis.call("INCR", KEYS[1])
			if not current then
					redis.call("EXPIRE", KEYS[1], ARGV[2])
			end
			return 1
	end
	`
	result, err := client.Eval(ctx, luaScript, []string{key}, limit, int(window.Seconds())).Result()
	if err != nil {
		return false, err
	}

	return result.(int64) == 1, nil
}

func CheckRates(ctx context.Context, region string, rates []Limit) error {
	var rateList []Limit = append(MainRates, rates...)
	limiter := rate.NewLimiter(rate.Every(10*time.Second), 1) // Adjust the parameters according to your needs

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if limiter.Allow() {
				allAllowed, err := checkAllRates(ctx, region, rateList)
				if err != nil {
					return err
				}
				if allAllowed {
					return incrementAllRates(ctx, region, rateList)
				}
			} else {
				// Wait for the next permission of the rate limiter
				time.Sleep(limiter.Reserve().Delay())
			}
		}
	}
}

func checkAllRates(ctx context.Context, region string, rates []Limit) (bool, error) {
	for _, rate := range rates {
		allowed, err := IsAllowed(ctx, region+"-"+rate.Name, rate.Reservoir, rate.Interval)
		if err != nil {
			return false, err
		}
		if !allowed {
			return false, nil
		}
	}
	return true, nil
}

func incrementAllRates(ctx context.Context, region string, rates []Limit) error {
	for _, rate := range rates {
		if _, err := incrementRate(ctx, region+"-"+rate.Name, rate.Reservoir, rate.Interval); err != nil {
			return err
		}
	}
	return nil
}

// // Checks all rates in array and increments all rates if all are allowed
// // if any rate is not allowed, it will sleep for the TTL of that rate limit key
// func CheckRates(region string, rates []Limit) error {
// 	var rateList []Limit = append(MainRates, rates...)
// 	maxRetries := 3

// 	for retry := 0; retry < maxRetries; retry++ {
// 		allAllowed := true
// 		// check all rates
// 		for _, rate := range rateList {
// 			allowed, err := IsAllowed(region+"-"+rate.Name, rate.Reservoir, rate.Interval)
// 			if err != nil {
// 				return err
// 			}
// 			if !allowed {
// 				allAllowed = false
// 				// Get TTL for the rate limit key
// 				ttl, err := client.TTL(ctx, region+"-"+rate.Name).Result()
// 				if err != nil {
// 					return err
// 				}
// 				time.Sleep(time.Duration(ttl) * time.Second)
// 				break
// 			}
// 		}

// 		// increment all rates
// 		if allAllowed {
// 			for _, rate := range rateList {
// 				_, err := incrementRate(region+"-"+rate.Name, rate.Reservoir, rate.Interval)
// 				if err != nil {
// 					return err
// 				}
// 			}
// 			return nil
// 		}
// 	}
// 	return fmt.Errorf("rate limit exceeded and max retries reached")
// }
