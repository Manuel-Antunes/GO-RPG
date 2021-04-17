package main

import (
	"html/template"
	"os"
)

func main() {
	tmpString := "Hello, {{.}}"
	tmpl := template.Must(template.New("tela").Parse(tmpString))
	err := tmpl.Execute(os.Stdout, "sla")
	if err != nil {
		panic("deu ruim")
	}
}
