package utils

import (
	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"
)

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

var alphaNumeric = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func StrLower(s string) string {
	return strings.ToLower(s)
}

func StrTrim(s string) string {
	return strings.TrimSpace(s)
}

func StrRandom(length int) string {
	result := make([]rune, length)
	for i := range result {
		result[i] = alphaNumeric[rand.Intn(len(alphaNumeric))]
	}
	return string(result)
}

func Panicln(format string, a ...interface{}) {
	log.Panicln(fmt.Errorf(format, a...))
}
