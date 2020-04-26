package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// https://stackoverflow.com/a/44359967/5042046
func PrettyJson(data interface{}) (string, error) {
	const (
		empty = ""
		tab   = "\t"
	)

	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent(empty, tab)

	err := encoder.Encode(data)
	if err != nil {
		return empty, err
	}
	return buffer.String(), nil
}

func FormatDateTime(s string) string {
	if s == "" {
		return ""
	}
	t, err := time.Parse("2006/01/02", strings.TrimSpace(s))
	if err != nil {
		fmt.Println(err)
		return ""
	} else {
		return t.Format("02-01-2006")
	}
}

func ConvertStringToNumber(s string) (int, error) {
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	} else {
		return num, nil
	}
}