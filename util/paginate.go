package util

type Paginate struct {
	Page    int32 `json:"page"`
	PerPage int32 `json:"per_page"`
}

func (p *Paginate) GetPerPageOffset() (int32, int32) {
	return p.PerPage, (p.Page - 1) * p.PerPage
}
