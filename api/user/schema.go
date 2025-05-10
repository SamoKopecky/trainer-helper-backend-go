package user

type userPostRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type userPutRequest struct {
	Id       string `json:"id"`
	Nickname string `json:"nickname"`
}
