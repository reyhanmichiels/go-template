package main

import (
	"errors"
	"fmt"
	"html/template"
	"os"
	"strings"
)

type Repository struct {
	Inter string
	Strct string
}

func main() {
	var handlerName string

	fmt.Scanln(&handlerName)

	handlerName = strings.ToLower(handlerName) + "Repository"

	repo := Repository{
		Inter: strings.Title(handlerName),
		Strct: handlerName,
	}

	file, err := CreateFileHandler(handlerName)

	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	tmplHandler, err := template.New("handler.tmpl").ParseFiles("handler.tmpl")
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}

	err = tmplHandler.Execute(file, repo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
}

func CreateFileHandler(name string) (*os.File, error) {
	path := "handler/" + name + ".go"

	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		err := os.MkdirAll("handler", os.ModePerm)
		if err != nil {
			return nil, err
		}

		f, err := os.Create(path)
		if err != nil {
			return nil, err
		}

		return f, nil
	}

	return nil, errors.New("handler already exist")

}
