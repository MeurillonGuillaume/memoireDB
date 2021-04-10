package model

type (
	// InsertModel is a model used to accept incoming Insert value requests
	InsertModel struct {
		Key   string      `json:"key"`
		Value interface{} `json:"value"`
	}

	// RetrieveModel is a model used to accept incoming Retrieve value requests
	RetrieveModel struct {
		Key string `json:"key"`
	}

	// ListKeyModel is a model used to accept incoming List Key requests with optional prefixes
	ListKeyModel struct {
		Prefix string `json:"prefix,omitempty"`
	}
)
