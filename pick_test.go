package redis

import (
	"fmt"
	"testing"
)

func TestPick(t *testing.T) {

	id := "lb"
	tc := new(TokenContent)

	Pick(id, tc)

	fmt.Println("result :", tc)
	fmt.Println(tc.User)

}
