package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

func PrintBody(body io.ReadCloser) {
	buf := new(bytes.Buffer)
	_, _ = buf.ReadFrom(body)
	fmt.Println(buf.String())
}

func PrintStruct(data interface{}) {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Print(string(b))
}
