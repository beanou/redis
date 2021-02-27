package redis

import (
	"fmt"
	"testing"
)

func TestPick(t *testing.T) {

	id := "lb"
	tc := new(TokenContent)

	err := Pick(id, tc)

	fmt.Println("err", err)
	fmt.Println("result :", tc)
	fmt.Println(tc.User)

}
