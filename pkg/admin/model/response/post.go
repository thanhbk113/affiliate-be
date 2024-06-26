package response

type ParPostRes struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type ProductAllResponse struct {
	List  []ParPostRes `json:"list"`
	Total int64        `json:"total"`
	Limit int64        `json:"limit"`
}

type SubPostRes struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	ParID     string `json:"parId"`
	Image     string `json:"image"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type ParPostAllResponse struct {
	List  []ParPostRes `json:"list"`
	Total int64        `json:"total"`
	Limit int64        `json:"limit"`
}

type SubPostAllResponse struct {
	List  []SubPostRes `json:"list"`
	Total int64        `json:"total"`
	Limit int64        `json:"limit"`
}
