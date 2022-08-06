package test

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"testing"
	"time"

	"github.com/fiure-cms/fiure-default-theme/data"
)

var distFolder = "dist"
var templatesFolder = "templates"
var layoutsFolder = "layouts"
var partialsFolder = "partials"

func Test404Template(t *testing.T) {

	layout := "error"
	filename := "404"

	// Get MyData
	md := getMyData()

	// Get MyFile
	f, err := createFile(filename)
	if err != nil {
		log.Fatalf("file create: %s", err)

		t.Errorf("Error: %v", err)
	}

	defer f.Close()

	// Render Template
	tmpl := template.Must(template.New(filename).Funcs(getMyFuncs()).ParseFiles(getMyFileList(layout, fmt.Sprintf("../%s/%s.tmpl", templatesFolder, filename))...))
	err = tmpl.ExecuteTemplate(f, layout, md)
	if err != nil {
		log.Fatalf("template execution: %s", err)

		t.Errorf("Error: %v", err)
	}
}

func TestIndexTemplate(t *testing.T) {

	layout := "base"
	filename := "index"

	// Get MyData
	md := getMyData()

	// Get MyFile
	f, err := createFile(filename)
	if err != nil {
		log.Fatalf("file create: %s", err)

		t.Errorf("Error: %v", err)
	}

	defer f.Close()

	// Render Template
	tmpl := template.Must(template.New(filename).Funcs(getMyFuncs()).ParseFiles(getMyFileList(layout, fmt.Sprintf("../%s/%s.tmpl", templatesFolder, filename))...))
	err = tmpl.ExecuteTemplate(f, layout, md)
	if err != nil {
		log.Fatalf("template execution: %s", err)

		t.Errorf("Error: %v", err)
	}
}

func TestPageTemplate(t *testing.T) {

	layout := "base"
	filename := "page"

	// Get MyData
	md := getMyData()

	// Get MyFile
	f, err := createFile(filename)
	if err != nil {
		log.Fatalf("file create: %s", err)

		t.Errorf("Error: %v", err)
	}

	defer f.Close()

	// Render Template
	tmpl := template.Must(template.New(filename).Funcs(getMyFuncs()).ParseFiles(getMyFileList(layout, fmt.Sprintf("../%s/%s.tmpl", templatesFolder, filename))...))
	err = tmpl.ExecuteTemplate(f, layout, md)
	if err != nil {
		log.Fatalf("template execution: %s", err)

		t.Errorf("Error: %v", err)
	}
}

func TestSearchTemplate(t *testing.T) {

	layout := "base"
	filename := "search"

	// Get MyData
	md := getMyData()

	// Get MyFile
	f, err := createFile(filename)
	if err != nil {
		log.Fatalf("file create: %s", err)

		t.Errorf("Error: %v", err)
	}

	defer f.Close()

	// Render Template
	tmpl := template.Must(template.New(filename).Funcs(getMyFuncs()).ParseFiles(getMyFileList(layout, fmt.Sprintf("../%s/%s.tmpl", templatesFolder, filename))...))
	err = tmpl.ExecuteTemplate(f, layout, md)
	if err != nil {
		log.Fatalf("template execution: %s", err)

		t.Errorf("Error: %v", err)
	}
}

func TestTermTemplate(t *testing.T) {

	layout := "base"
	filename := "term"

	// Get MyData
	md := getMyData()

	// Get MyFile
	f, err := createFile(filename)
	if err != nil {
		log.Fatalf("file create: %s", err)

		t.Errorf("Error: %v", err)
	}

	defer f.Close()

	// Render Template
	tmpl := template.Must(template.New(filename).Funcs(getMyFuncs()).ParseFiles(getMyFileList(layout, fmt.Sprintf("../%s/%s.tmpl", templatesFolder, filename))...))
	err = tmpl.ExecuteTemplate(f, layout, md)
	if err != nil {
		log.Fatalf("template execution: %s", err)

		t.Errorf("Error: %v", err)
	}
}

func TestAuthorTemplate(t *testing.T) {

	layout := "base"
	filename := "author"

	// Get MyData
	md := getMyData()

	// Get MyFile
	f, err := createFile(filename)
	if err != nil {
		log.Fatalf("file create: %s", err)

		t.Errorf("Error: %v", err)
	}

	defer f.Close()

	// Render Template
	tmpl := template.Must(template.New(filename).Funcs(getMyFuncs()).ParseFiles(getMyFileList(layout, fmt.Sprintf("../%s/%s.tmpl", templatesFolder, filename))...))
	err = tmpl.ExecuteTemplate(f, layout, md)
	if err != nil {
		log.Fatalf("template execution: %s", err)

		t.Errorf("Error: %v", err)
	}
}

func TestSingleTemplate(t *testing.T) {

	layout := "base"
	filename := "single"

	// Get MyData
	md := getMyData()

	// Get MyFile
	f, err := createFile(filename)
	if err != nil {
		log.Fatalf("file create: %s", err)

		t.Errorf("Error: %v", err)
	}

	defer f.Close()

	// Render Template
	tmpl := template.Must(template.New(filename).Funcs(getMyFuncs()).ParseFiles(getMyFileList(layout, fmt.Sprintf("../%s/%s.tmpl", templatesFolder, filename))...))
	err = tmpl.ExecuteTemplate(f, layout, md)
	if err != nil {
		log.Fatalf("template execution: %s", err)

		t.Errorf("Error: %v", err)
	}
}

func getMyData() data.MyData {
	md := data.MyData{}
	md.GetDataFromFile("../data/data.json")

	return md
}

func getMyFuncs() template.FuncMap {
	return template.FuncMap{
		"convertUpdateTime": convertUpdateTime,
		"add":               add,
		"remove":            remove,
	}
}

func getMyFileList(layout string, files ...string) []string {

	requiredFiles := []string{
		fmt.Sprintf("../%s/%s/%s.tmpl", templatesFolder, layoutsFolder, layout),
		fmt.Sprintf("../%s/%s/%s.tmpl", templatesFolder, partialsFolder, "header"),
		fmt.Sprintf("../%s/%s/%s.tmpl", templatesFolder, partialsFolder, "pagination"),
		fmt.Sprintf("../%s/%s/%s.tmpl", templatesFolder, partialsFolder, "sidebar"),
		fmt.Sprintf("../%s/%s/%s.tmpl", templatesFolder, partialsFolder, "footer"),
	}

	return append(requiredFiles, files...)
}

func createFile(filename string) (*os.File, error) {
	return os.Create(fmt.Sprintf("../%s/%s.html", distFolder, filename))
}

func convertUpdateTime(updated int64) time.Time {
	return time.Unix(updated, 0)
}

func add(a int, b int) int {
	return a + b
}

func remove(a int, b int) int {
	return a - b
}
