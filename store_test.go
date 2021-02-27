package redis

import (
	"fmt"
	"testing"
)

type TokenContent struct {
	User      string
	TokenCode string
	Domain    string
	Type      string
}

func TestStore(t *testing.T) {
	id := "lb"
	tc := new(TokenContent)
	tc.User = "liub"
	tc.TokenCode = "11111"
	tc.Domain = "www.sumg.press"
	tc.Type = "code"
	err := Stroe(id, tc, 7200)
	fmt.Println("error:", err)
}
