package user_handler

type userPostRequest struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type userDeleteRequest struct {
	Id string `json:"id"`
}

type userPutRequest struct {
	Id       string `json:"id"`
	Nickname string `json:"nickname"`
}
