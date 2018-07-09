package cmd_test

import (
	"bufio"
	"errors"
	"strings"
	"testing"

	"github.com/isaacasensio/memo/cmd"
	"github.com/stretchr/testify/assert"
)

type errorPasswordReader struct {
}

func (pr errorPasswordReader) ReadPassword() (string, error) {
	return "", errors.New("error :)")
}

type stubPasswordReader struct {
	Password string
}

func (pr stubPasswordReader) ReadPassword() (string, error) {
	return pr.Password, nil
}
func TestRunReturnsErrorWhenReadPasswordFails(t *testing.T) {
	var reader errorPasswordReader
	scanner := bufio.NewScanner(strings.NewReader(""))
	result, err := cmd.Run(reader, scanner)
	assert.Equal(t, errors.New("error :)"), err)
	assert.Equal(t, "", result)
}
func TestRunReturnsRequestedRunes(t *testing.T) {
	input := "2\n5\n6\n"
	scanner := bufio.NewScanner(strings.NewReader(input))
	pr := stubPasswordReader{Password: "password"}
	result, err := cmd.Run(pr, scanner)
	assert.NoError(t, err)
	assert.Equal(t, "awo", result)
}

func TestRunOnlyReturnsThreeRunes(t *testing.T) {
	input := "1\n5\n6\n2"
	scanner := bufio.NewScanner(strings.NewReader(input))
	pr := stubPasswordReader{Password: "password"}
	result, err := cmd.Run(pr, scanner)
	assert.NoError(t, err)
	assert.Equal(t, "pwo", result)
}

func TestRunReturnsErrorWhenNoneIntegerValuesAreProvidedForPosition(t *testing.T) {
	input := "2\n5\na\n8"
	scanner := bufio.NewScanner(strings.NewReader(input))
	pr := stubPasswordReader{Password: "password"}
	result, err := cmd.Run(pr, scanner)
	assert.Error(t, err)
	assert.Equal(t, "", result)
}

func TestRunReturnsErrorWhenEmptyPasswordIsProvided(t *testing.T) {
	scanner := bufio.NewScanner(strings.NewReader(""))
	pr := stubPasswordReader{Password: ""}
	result, err := cmd.Run(pr, scanner)
	assert.Error(t, err)
	assert.Equal(t, "", result)
}
