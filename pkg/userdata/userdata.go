package userdata

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/Serhii1Epam/simpleHttpServer/pkg/appdb"
	"github.com/Serhii1Epam/simpleHttpServer/pkg/hasher"
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

func NewUser() *UserData {
	return &UserData{}
}

func (u *UserData) CreateUser(w http.ResponseWriter, db *appdb.Database) error {
	fmt.Fprintf(w, "CreateUser = User : %s Password : %s\n", u.User, u.Password)
	hsr := hasher.NewHasher(u.Password)
	calculated_hash, _ := hsr.HashPassword()
	fmt.Fprintf(w, "hsr[%v]\n calc[%s]\n", hsr, calculated_hash)
	db.Insert(*hsr, u.User)
	fmt.Fprintf(w, "DB: [%v]\n", db)
	return nil
}

func (u *UserData) LoginUser(w http.ResponseWriter, db *appdb.Database) error {
	fmt.Fprintf(w, "LoginUser = User : %s Password : %s\n", u.User, u.Password)
	hsr := hasher.NewHasher(u.Password)
	if !hsr.CheckPasswordHash(string(db.Select(u.User))) {
		fmt.Fprintf(w, "Invalid User:[%v], \n", u)
	}

	return nil
}

func (in *JsonBytes) ParseBody() *UserData {
	inUser := NewUser()
	json.Unmarshal(*in, &inUser)
	return inUser

}

func (in *UndefinedBytes) ParseBody() *UserData {
	var data []byte = *in
	inUser := NewUser()
	str := strings.Fields(string(data))
	inUser.User = str[0]
	inUser.Password = str[1]
	return inUser

}

func Parse(i IParse) *UserData {
	return i.ParseBody()
}
