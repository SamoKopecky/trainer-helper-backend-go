package user_handler

type userPostRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}
