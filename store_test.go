package redis

import "testing"

func TestStore(t *testing.T) {
	id := "lb"
	tc := new(TokenContent)
	tc.User = "liub"
	tc.TokenCode = "11111"
	tc.Domain = "www.sumg.press"
	tc.Type = "code"
	Stroe(id, tc, 5)
}
