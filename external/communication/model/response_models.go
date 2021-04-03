package model

// StatusResponse is an object to return the current status to a requester
type StatusResponse struct {
	Message string `json:"message"`
}
