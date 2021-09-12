package main

import (
	"bytes"
	"fmt"
	// "strings"
)

func main() {
	comma("12345678")
	comma("105803495842345678")
	comma("123")
	comma("12")
	comma("3")
	comma("")
	comma("354839")
}

func comma (s string) string {

	var buf bytes.Buffer

	strlen := len(s)
	i := strlen % 3
	iterations := (strlen - i) / 3

	buf.WriteString(s[:i])

	for j := 0; j < iterations; j++ {
		if buf.Len() > 0 {
			buf.WriteByte(',')
		}
		buf.WriteString(s[i:i+3])
		i += 3
	}

	fmt.Println(buf.String())

	return ""
}