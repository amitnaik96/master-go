package main

import (
	"os"
	"text/template"
)

func templateFn() {
	// fill value in {{.}}
	templ := template.Must(template.New("mytemplate").Parse("Value: {{.}}\n"))

	templ.Execute(os.Stdout, "Hello")
	templ.Execute(os.Stdout, []string{"a", "b"})
}
