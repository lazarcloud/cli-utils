package cli

import (
	"bufio"
	"os"
)

func ReadInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	choice, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return choice, nil
}
