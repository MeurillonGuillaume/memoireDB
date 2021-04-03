package model

type (
	// InsertModel is a model used to accept incoming Insert value requests
	InsertModel struct {
		Key   string      `json:"key"`
		Value interface{} `json:"value"`
	}
)
