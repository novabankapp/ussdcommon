package utils

import "time"

func IsNewSession(sessionId string, sessionDict map[string]time.Time) bool {
	if val, ok := sessionDict[sessionId]; ok {
		now := time.Now()
		if now.Sub(val).Hours() >= 1 {
			sessionDict[sessionId] = now
			return true
		} else {
			return false
		}
	} else {
		sessionDict[sessionId] = time.Now()
		return true
	}

}
