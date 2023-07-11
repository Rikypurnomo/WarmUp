package authdto

type LoginResponse struct {
	Password string `json:"token"`
}

type LoginTokenResponse struct {
	NewUuid string
	Expired int64
	Token string

}
