package face

// FaceMatching stores face matching result from the given input
type FaceMatching struct {
	Status string             `json:"status"`
	Result FaceMatchingResult `json:"result"`
}

// FaceMatchingResult stores the face matching metadata from
// GLAIR Vision Face Matching API
type FaceMatchingResult struct {
	MatchPercentage float64 `json:"match_percentage"`
	MatchStatus     string  `json:"match_status"`
	MatchStatusCode string  `json:"match_status_code"`
}
