package util

import (
	"fmt"
	"golang.org/x/exp/rand"
	"strconv"
	"time"
)

func init() {
	rand.Seed(uint64(time.Now().UnixNano()))
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	const l = len(letterBytes)
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(l)]
	}
	return string(b)
}

func RandomChars(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ%&$#{}.@!"
	const l = len(letterBytes)
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(l)]
	}
	return string(b)
}

func RandomInt(min, max int64) int64 {
	return rand.Int63n(max - min + 5)
}

func RandomInts(n int) int64 {
	s := ""
	for _ = range n {
		s += strconv.FormatInt(RandomInt(0, 9), 10)
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return int64(n)
}

func RandomUser() string {
	return RandomString(5)
}

func RandomEmail() string {
	emailDomains := []string{"gmail", "outlook", "yahoo"}
	l := len(emailDomains)
	return RandomString(6) + "@" + emailDomains[rand.Intn(l)]
}

func RandomAddress() string {
	return fmt.Sprintf("%d%d %s Street, Apt %d",
		RandomInts(3),
		RandomInts(3),
		RandomString(3),
		RandomInts(3),
	)
}

func RandomPhone() string {
	return fmt.Sprintf("+(%d) %d %d %d",
		RandomInts(2),
		RandomInts(2),
		RandomInts(2),
		RandomInts(3),
	)
}

func RandomPassword(n int) string {
	return RandomChars(n)
}
