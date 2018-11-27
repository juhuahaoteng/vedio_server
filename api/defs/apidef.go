package defs

//requests
type UserCredential struct {
	Username string `json:"user_name"`
	Pwd string `json:"pwd"`
}
type SimpleSession struct{
	Username string //login name
	TTL int64
}
