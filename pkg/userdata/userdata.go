package userdata

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type UserData struct {
	User     string `json:"user"`
	Password string `json:"password"`
}

type IParse interface {
	ParseBody() *UserData
}

type JsonBytes []byte
type UndefinedBytes []byte

func New() *UserData {
	return &UserData{}
}

func (u *UserData) CreateUser(w http.ResponseWriter) error {
	fmt.Fprintf(w, "CreateUser = User : %s Password : %s\n", u.User, u.Password)
	return nil
}

func (u *UserData) LoginUser(w http.ResponseWriter) error {
	fmt.Fprintf(w, "LoginUser = User : %s Password : %s\n", u.User, u.Password)
	return nil
}

func (in *JsonBytes) ParseBody() *UserData {
	inUser := New()
	json.Unmarshal(*in, &inUser)
	return inUser

}

func (in *UndefinedBytes) ParseBody() *UserData {
	var data []byte = *in
	inUser := New()
	str := strings.Fields(string(data))
	inUser.User = str[0]
	inUser.Password = str[1]
	return inUser

}

func Parse(i IParse) *UserData {
	return i.ParseBody()
}
