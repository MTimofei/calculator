package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Input() (example string, err error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("enter an example: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		text = ""
		return text, err
	}
	text = strings.TrimSpace(text)
	return text, nil
}
