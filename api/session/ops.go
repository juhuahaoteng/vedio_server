package session

import "sync"



var sessionMap *sync.Map

func init()  {
	sessionMap = &sync.Map{}
}

func LoadSessionsFromDB()  {
	
}
func GenerateNewSessionId(un string) string {
	return "a"
}

func IsSessionExpired(sid string) (string,bool) {
	return "a",false
}


