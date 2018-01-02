package main

import (
	"html/template"
	"os"
)

type entry struct {
	Name string
	Done bool
}

type ToDo struct {
	User string
	List []entry
}

func main() {
	// prepare data
	entry1 := entry{"dinner", false}
	entry2 := entry{"breakfast", true}
	entry3 := entry{"lunch", false}
	var entries []entry
	entries = append(entries, entry1)
	entries = append(entries, entry2)
	entries = append(entries, entry3)
	todo := ToDo{"user1", entries}

	t := template.Must(template.ParseGlob("*.html")) // cannot use "template.New()"
	err := t.Execute(os.Stdout, todo)
	if err != nil {
		panic(err)
	}
}
