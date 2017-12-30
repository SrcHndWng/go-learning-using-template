package main

import (
	"os"
	"text/template"
)

type Todo struct {
	Name        string
	Description string
}

func main() {
	t := template.Must(template.New("todos").Parse("You have task named \"{{ .Name}}\" with description: \"{{ .Description}}\"\n"))

	todo := Todo{"Test templates", "Let's test a template to see the magic."}
	err := t.Execute(os.Stdout, todo)
	if err != nil {
		panic(err)
	}

	another := Todo{"Go", "Contribute to any Go project"}
	err = t.Execute(os.Stdout, another)
	if err != nil {
		panic(err)
	}
}
