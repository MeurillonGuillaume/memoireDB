package model

type (
	// StatusResponse is an object to return the current status to a requester
	StatusResponse struct {
		Message string `json:"message"`
	}

	SimpleResponse struct {
		Result  string      `json:"result"`
		Value   interface{} `json:"value"`
		Message string      `json:"message,omitempty"`
		Error   string      `json:"error,omitempty"`
		Took    int64       `json:"took"`
	}

	RetrieveResponse struct {
		Key   string      `json:"key"`
		Value interface{} `json:"value"`
		Took  int64       `json:"took"`
	}
)
