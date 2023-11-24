package glair

// Session stores session data from any GLAIR Vision
// session requests
type Session struct {
	Status     string `json:"status"`
	SuccessURL string `json:"success_url"`
	CancelURL  string `json:"cancel_url"`
	URL        string `json:"url"`
}

// SessionsInput stores sessions request input
type SessionsInput struct {
	// SuccessURL represents redirection URL
	// when the session is concluded successfully
	//
	// SuccessURL is required
	SuccessURL string
	// CancelURL represents redirection URL
	// when user presses the back button during
	// the session or the session encountered an
	// error when processing the request
	//
	// CancelURL is optional
	CancelURL *string
}
