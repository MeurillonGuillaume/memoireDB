package model

type (
	// StatusResponse is an object to return the current status to a requester.
	StatusResponse struct {
		Message string `json:"message"`
	}

	SimpleResponse struct {
		Result  string      `json:"result"`
		Value   interface{} `json:"value,omitempty"`
		Message string      `json:"message,omitempty"`
		Error   string      `json:"error,omitempty"`
		Took    int64       `json:"tookNs"`
	}

	RetrieveResponse struct {
		Key   string      `json:"key"`
		Value interface{} `json:"value"`
		Took  int64       `json:"tookNs"`
	}

	ListKeysResponse struct {
		Prefix string      `json:"prefix,omitempty"`
		Keys   interface{} `json:"keys"`
		Took   int64       `json:"tookNs"`
	}
)
