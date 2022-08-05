package test

import (
	"html/template"
	"log"
	"os"
	"testing"
	"time"

	"github.com/fiure-cms/fiure-default-theme/data"
)

func Test404Template(t *testing.T) {

	// Get MyData
	md := getMyData()

	// Get MyFile
	f, err := os.Create("../dist/404.html")
	if err != nil {
		log.Fatalf("file create: %s", err)

		t.Errorf("Error: %v", err)
	}

	defer f.Close()

	// Render Template
	tmpl := template.Must(template.New("404.html").Funcs(getMyFuncs()).ParseFiles(getMyFileList("../templates/layouts/error.tmpl", "../templates/404.tmpl")...))
	err = tmpl.ExecuteTemplate(f, "error", md)
	if err != nil {
		log.Fatalf("template execution: %s", err)

		t.Errorf("Error: %v", err)
	}
}

func TestIndexTemplate(t *testing.T) {

	// Get MyData
	md := getMyData()

	// Get MyFile
	f, err := os.Create("../dist/index.html")
	if err != nil {
		log.Fatalf("file create: %s", err)

		t.Errorf("Error: %v", err)
	}

	defer f.Close()

	// Render Template
	tmpl := template.Must(template.New("index.html").Funcs(getMyFuncs()).ParseFiles(getMyFileList("", "../templates/index.tmpl")...))
	err = tmpl.ExecuteTemplate(f, "base", md)
	if err != nil {
		log.Fatalf("template execution: %s", err)

		t.Errorf("Error: %v", err)
	}
}

func TestLoginTemplate(t *testing.T) {

	// Get MyData
	md := getMyData()

	// Render Template
	tmpl := template.Must(template.New("login.html").Funcs(getMyFuncs()).ParseFiles(getMyFileList("", "../templates/login.tmpl")...))
	err := tmpl.ExecuteTemplate(os.Stdout, "base", md)
	if err != nil {
		log.Fatalf("template execution: %s", err)

		t.Errorf("Error: %v", err)
	}
}

func TestPageTemplate(t *testing.T) {

	// Get MyData
	md := getMyData()

	// Render Template
	tmpl := template.Must(template.New("page.html").Funcs(getMyFuncs()).ParseFiles(getMyFileList("", "../templates/page.tmpl")...))
	err := tmpl.ExecuteTemplate(os.Stdout, "base", md)
	if err != nil {
		log.Fatalf("template execution: %s", err)

		t.Errorf("Error: %v", err)
	}
}

func TestSearchTemplate(t *testing.T) {

	// Get MyData
	md := getMyData()

	// Render Template
	tmpl := template.Must(template.New("search.html").Funcs(getMyFuncs()).ParseFiles(getMyFileList("", "../templates/search.tmpl")...))
	err := tmpl.ExecuteTemplate(os.Stdout, "base", md)
	if err != nil {
		log.Fatalf("template execution: %s", err)

		t.Errorf("Error: %v", err)
	}
}

func TestSingleTemplate(t *testing.T) {

	// Get MyData
	md := getMyData()

	// Render Template
	tmpl := template.Must(template.New("single.html").Funcs(getMyFuncs()).ParseFiles(getMyFileList("", "../templates/single.tmpl")...))
	err := tmpl.ExecuteTemplate(os.Stdout, "base", md)
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
	if layout == "" {
		layout = "../templates/layouts/base.tmpl"
	}

	requiredFiles := []string{
		layout,
		"../templates/partials/header.tmpl",
		"../templates/partials/pagination.tmpl",
		"../templates/partials/sidebar.tmpl",
		"../templates/partials/footer.tmpl",
	}

	return append(requiredFiles, files...)
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
