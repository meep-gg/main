package lol

var SeasonStart14_1 = int64(1704895200)

var Queue = map[string]bool{
	"RANKED_SOLO_5x5": true,
	"RANKED_FLEX_SR":  true,
}

var SimpleQueue = map[string]bool{
	"solo": true,
	"flex": true,
}

var QueueValue = map[string]string{
	"RANKED_SOLO_5x5": "solo",
	"RANKED_FLEX_SR":  "flex",
}

var QueueValueReverse = map[string]string{
	"solo": "RANKED_SOLO_5x5",
	"flex": "RANKED_FLEX_SR",
}

var Tier = map[string]bool{
	"IRON":        true,
	"BRONZE":      true,
	"SILVER":      true,
	"GOLD":        true,
	"PLATINUM":    true,
	"EMERALD":     true,
	"DIAMOND":     true,
	"MASTER":      true,
	"GRANDMASTER": true,
	"CHALLENGER":  true,
}

var Apex = map[string]bool{
	"MASTER":      true,
	"GRANDMASTER": true,
	"CHALLENGER":  true,
}

var TierValue = map[string]int32{
	"IRON":     0,
	"BRONZE":   400,
	"SILVER":   800,
	"GOLD":     1200,
	"PLATINUM": 1600,
	"EMERALD":  2000,
	"DIAMOND":  2400,
	"APEX":     2800,
}

var Division = map[string]bool{
	"I":   true,
	"II":  true,
	"III": true,
	"IV":  true,
}

var DivisionValue = map[string]int32{
	"I":   300,
	"II":  200,
	"III": 100,
	"IV":  0,
}

var Regional = map[string]bool{
	"americas": true,
}

// "asia":     true,
// "europe":   true,
// "sea":      true,
// "esports":  true,

var RegionalValue = map[string]string{
	"na1":  "americas",
	"br1":  "americas",
	"eun1": "europe",
	"euw1": "europe",
	"jp1":  "asia",
	"kr":   "asia",
	"la1":  "americas",
	"la2":  "americas",
	"oc1":  "americas",
	"tr1":  "europe",
	"ru":   "europe",
	"ph2":  "sea",
	"sg2":  "sea",
	"th2":  "sea",
	"tw2":  "sea",
	"vn2":  "sea",
}
var Platform = map[string]bool{
	"na1": true,
}

// "br1":  true,
// "eun1": true,
// "euw1": true,
// "jp1":  true,
// "kr":   true,
// "la1":  true,
// "la2":  true,
// "oc1":  true,
// "tr1":  true,
// "ru":   true,
// "ph2":  true,
// "sg2":  true,
// "th2":  true,
// "tw2":  true,
// "vn2":  true,
func IsTier(tier string) bool {
	_, exists := Tier[tier]
	return exists
}

func IsDivision(division string) bool {
	_, exists := Division[division]
	return exists
}

func IsRegional(region string) bool {
	_, exists := Regional[region]
	return exists
}

func IsPlatform(platform string) bool {
	_, exists := Platform[platform]
	return exists
}

func IsQueue(queue string) bool {
	_, exists := Queue[queue]
	return exists
}

func IsSimpleQueue(queue string) bool {
	_, exists := SimpleQueue[queue]
	return exists
}

func IsApex(tier string) bool {
	_, exists := Apex[tier]
	return exists
}
