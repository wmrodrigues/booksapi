package structs

// HeartbeatResponse encapsulate the basic attributes to sent a response to client
type HeartbeatResponse struct {
	Status string `json:"status"`
	Code   int    `json:"code"`
}
