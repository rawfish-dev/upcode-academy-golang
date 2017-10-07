package common

/* An import is missing here */

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
