package user

type LoginWithUsernamePasswordReqBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Captcha  struct {
		Id   string `json:"id"`
		Code string `json:"code"`
	} `json:"captcha"`
}
