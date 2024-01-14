package lol

import "fmt"

func SerializeRank(tier, division string, lp int32) (int32, error) {
	if IsApex(tier) {
		return TierValue["APEX"] + lp, nil
	}
	if !IsTier(tier) {
		return 0, fmt.Errorf("invalid tier: %v", tier)
	}
	if !IsDivision(division) {
		return 0, fmt.Errorf("invalid division: %v", division)
	}
	return TierValue[tier] + DivisionValue[division] + lp, nil
}

func SerializeQueue(queue string) (string, error) {
	if !IsQueue(queue) {
		return "", fmt.Errorf("invalid queue: %v", queue)
	}
	return QueueValue[queue], nil
}

func DeserializeQueue(queue string) (string, error) {
	if !IsSimpleQueue(queue) {
		return "", fmt.Errorf("invalid queue: %v", queue)
	}
	return QueueValueReverse[queue], nil
}

func PlatformToRegional(platform string) (string, error) {
	if !IsPlatform(platform) {
		return "", fmt.Errorf("invalid platform: %v", platform)
	}
	return RegionalValue[platform], nil
}
