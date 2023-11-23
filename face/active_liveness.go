package face

// ActiveLiveness stores active liveness detection result
// from the given input
type ActiveLiveness struct {
	Status string               `json:"status"`
	Result ActiveLivenessResult `json:"result"`
}

type ActiveLivenessResult struct {
	GestureStatus   string `json:"gesture_status"`
	DetectedGesture string `json:"detected_gesture"`
}
