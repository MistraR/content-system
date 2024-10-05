package utils

import "fmt"

func GetAuthKey(sessionId string) string {
	authKey := fmt.Sprintf("session_auth:" + sessionId)
	return authKey
}

func GetSessionKey(userId string) string {
	sessionKey := fmt.Sprintf("session_id:" + userId)
	return sessionKey
}
