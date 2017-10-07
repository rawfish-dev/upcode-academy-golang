package common

import (
	"fmt"
)

type Message struct {
	Value string `json:"value"`
}

func BuildMessagePad(header string) string {
	return fmt.Sprintf(`
	********************
	%s
	********************
	`, header)
}
