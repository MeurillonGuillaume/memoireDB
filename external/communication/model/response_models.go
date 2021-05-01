package model

type (
	// StatusResponse is an object to return the current status to a requester.
	StatusResponse struct {
		Message string `json:"message"`
	}

	SimpleResponse struct {
		ID      string      `json:"id"`
		Result  string      `json:"result"`
		Value   interface{} `json:"value,omitempty"`
		Message string      `json:"message,omitempty"`
		Error   string      `json:"error,omitempty"`
		Took    int64       `json:"tookNs"`
	}

	RetrieveResponse struct {
		ID    string      `json:"id"`
		Key   string      `json:"key"`
		Value interface{} `json:"value"`
		Took  int64       `json:"tookNs"`
	}

	ListKeysResponse struct {
		ID     string      `json:"id"`
		Prefix string      `json:"prefix,omitempty"`
		Keys   interface{} `json:"keys"`
		Took   int64       `json:"tookNs"`
	}
)
