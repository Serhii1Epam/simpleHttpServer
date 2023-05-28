package userdata

type UserData struct {
	User     string `json:"user"`
	Password string `json:"password"`
	ErrMsg   string `json:"errMsg"`
}
