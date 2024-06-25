package response

type IntroduceResponse struct {
	ID        string `json:"id"`
	Content   string `json:"content,omitempty"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
