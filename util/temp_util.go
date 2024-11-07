package util

import (
	"embed"
	"fmt"
	"text/template"
)

var f embed.FS

var templates = template.Must(template.ParseFS(f, "templates/*"))

func GenTemp() {
	fmt.Println(templates)
}
