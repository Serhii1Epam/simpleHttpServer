package hasher

import (
	"fmt"
	"testing"
)

/* Sha256 from "TestPass" == eddef9e8e578c2a560c3187c4152c8b6f3f90c1dcf8c88b386ac1a9a96079c2c
 */
func TestCheckPasswordHash(t *testing.T) {
	want := "eddef9e8e578c2a560c3187c4152c8b6f3f90c1dcf8c88b386ac1a9a96079c2c"
	if !CheckPasswordHash("TestPass", want) {
		t.Errorf("CheckPasswordHash() => want %q", want)
	}
}

/* Sha256 from "TestPass1" == 4ee33bac59675856c9d8f9ddfaf21368a08f8afe7827516c6d031b8859064229
 */
func TestHashPassword(t *testing.T) {
	want := "4ee33bac59675856c9d8f9ddfaf21368a08f8afe7827516c6d031b8859064229"
	if got, err := HashPassword("TestPass1"); got != want {
		fmt.Printf("%x", got)
		t.Errorf("HashPassword() = %q, want %q, err %s", got, want, err)
	}
}
