package pagination

import "project-e-commerces/entities"

type ProductPagination struct {
	Limit        int                `json:"limit"`
	Page         int                `json:"page"`
	TotalRows    int                `json:"total_rows"`
	FirstPage    string             `json:"first_page"`
	LastPage     string             `json:"last_page"`
	PreviousPage string             `json:"previous_page"`
	NextPage     string             `json:"next_page"`
	FromRow      int                `json:"from_row"`
	ToRow        int                `json:"to_row"`
	Rows         []entities.Product `json:"rows"`
}
