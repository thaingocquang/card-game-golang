package util

// Paging ...
type Paging struct {
	Page  int   `query:"page" json:"page"`
	Limit int   `query:"limit" json:"limit"`
	Total int64 `query:"total" json:"total"`
}

// Fulfill ...
func (p *Paging) Fulfill() {
	if p.Page <= 0 {
		p.Page = 1
	}

	if p.Limit <= 0 {
		p.Limit = 10
	}
}
