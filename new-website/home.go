package main

//import packages we are going to use
import (
	"fmt"
	"os"
)

//defining a struct which would define the basis of a page
type Page struct {
	Title string
	Body  []byte
}

//create a functoin which saves the file as .html and returns an error type
func (p *Page) save() error {
	filename := p.Title + ".html"
	return os.WriteFile(filename, p.Body, 0600) //0600 = read/write perms for current user
}

//func to load page
func loadPage(title string) (*Page, error) {
	filename := title + ".html"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	p1 := &Page{Title: "TestPage", Body: []byte("<h1>This is a sample Page.</h1>")}
	p1.save()
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}
