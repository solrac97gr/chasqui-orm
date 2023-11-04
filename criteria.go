package chasqui

type Criteria struct {
	Pagination struct {
		Limit  int `json:"limit,omitempty"`
		Offset int `json:"offset,omitempty"`
	} `json:"pagination,omitempty"`
	Query struct {
		Filters []struct {
			Conditions []struct {
				Field    string      `json:"field,omitempty"`
				Operator string      `json:"operator,omitempty"`
				Value    interface{} `json:"value,omitempty"`
			} `json:"conditions,omitempty"`
			Logic string `json:"logic,omitempty"`
		} `json:"filters,omitempty"`
		Orders []struct {
			Field string `json:"field,omitempty"`
			Order string `json:"order,omitempty"`
		} `json:"orders,omitempty"`
		Logic string `json:"logic,omitempty"`
	} `json:"query,omitempty"`
}
