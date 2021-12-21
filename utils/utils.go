package utils

import (
	"bytes"
	"strings"
	"text/template"
)

func Msg(fmt string, args map[string]interface{}) (str string) {
	var msg bytes.Buffer

	tmpl, err := template.New("").Option("missingkey=zero").Parse(fmt)

	if err != nil {
		return fmt
	}

	tmpl.Execute(&msg, args)
	return msg.String()
}

func CleanQueryString(value string) string {
	r := strings.ReplaceAll(value, "\n", "")
	r = strings.ReplaceAll(value, "\t", "")
	r = strings.ReplaceAll(value, "<no value>", "0")

	return r
}
