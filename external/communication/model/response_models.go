package model

type (
	// StatusResponse is an object to return the current status to a requester
	StatusResponse struct {
		Message string `json:"message"`
	}

	// SimpleResponse
	SimpleResponse struct {
		Result  string `json:"result,omitempty"`
		Message string `json:"message,omitempty"`
		Took    int64  `json:"took,omitempty"`
	}
)
