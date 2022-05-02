package utils

import (
	"errors"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/sirupsen/logrus"
)

var sshKeyFile = []string{
	"id_rsa",
	"id_ed25519",
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// UniqueStrings returns slice of unique string given a string slice
func UniqueStrings(strSlice []string) []string {
	u := make([]string, 0, len(strSlice))
	m := make(map[string]bool)

	for _, val := range strSlice {
		if _, ok := m[val]; !ok {
			m[val] = true
			u = append(u, val)
		}
	}

	return u
}

func GenerateRand(n int) string {
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func DefaultSSHKeyPath() (string, error) {
	dirname, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	for _, k := range sshKeyFile {
		privSshKey := filepath.Join(dirname, ".ssh", k)
		logrus.Info(privSshKey)

		_, err := os.Stat(privSshKey)
		if err == nil {
			return privSshKey, nil
		}
	}

	return "", errors.New("ssh keys not found")
}
