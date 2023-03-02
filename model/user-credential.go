package model

type UserCredential struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AuthHeader struct {
	AuthorizationHeader string `header:"authorization"`
}
