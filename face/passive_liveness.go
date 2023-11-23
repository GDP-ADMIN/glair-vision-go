package face

// PassiveLiveness stores passive liveness detection result
// from the given input
type PassiveLiveness struct {
	Status string                `json:"status"`
	Result PassiveLivenessResult `json:"result"`
}

type PassiveLivenessResult struct {
	SpoofPercentage float64 `json:"spoof_percentage"`
	Status          string  `json:"status"`
}
