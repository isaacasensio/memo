package cmd

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

// PasswordReader returns password read from a reader
type PasswordReader interface {
	ReadPassword() (string, error)
}

// StdInPasswordReader default stdin password reader
type StdInPasswordReader struct {
}

// ReadPassword reads password from stdin
func (pr StdInPasswordReader) ReadPassword() (string, error) {
	pwd, error := terminal.ReadPassword(int(syscall.Stdin))
	return string(pwd), error
}

func readPosition(s *bufio.Scanner) (int, error) {
	s.Scan()
	return strconv.Atoi(s.Text())
}

func readPassword(pr PasswordReader) (string, error) {
	pwd, err := pr.ReadPassword()
	if err != nil {
		return "", err
	}
	if len(pwd) == 0 {
		return "", errors.New("empty password provided")
	}
	return pwd, nil
}

// Run reads memorable information and chars index from stdin and returns runes matching requested information.
func Run(pr PasswordReader, s *bufio.Scanner) (string, error) {
	fmt.Println("Enter memorable information / password: ")
	pwd, err := readPassword(pr)
	if err != nil {
		return "", err
	}

	fmt.Println("")

	memo := []rune(pwd)
	result := make([]rune, 0)

	for i := 0; i < 3; i++ {
		fmt.Print("Character: ")
		p, err := readPosition(s)
		if err != nil {
			return "", err
		}
		if p > len(memo) {
			return "", fmt.Errorf("position: %d > password length: %d", p, len(memo))
		}
		result = append(result, memo[p-1])
	}

	return string(result), nil
}
