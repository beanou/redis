package redis

// import "encoding/json"

// 此结构在使用项目中定义
type TokenContent struct {
	User      string
	TokenCode string
	Domain    string
	Type      string
}

// func (tc *TokenContent) MarshalBinary() ([]byte, error) {
// return json.Marshal(tc)
// }

type Results struct {
	ErrCode int64
	ErrMsg  string
	// Data    *TokenContent
}
