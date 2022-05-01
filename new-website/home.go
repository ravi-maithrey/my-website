package main

//import packages we are going to use
import (
	"fmt"
	"log"
	"net/http"
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

//this will handle urlsprefixed with "articles"
func viewArticle(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/articles/"):] //extract page title by removing slice of first part containing /articles/
	p, _ := loadPage(title)                 //_ allows us to skip storing the err return value as we wont use it
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

//this will craete a new page or open the existing page both as a html form
func editArticles(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):] //extract page title by removing slice of first part containing /articles/
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title} //we are creating a page if such a page doesn't exist
	}

}

func main() {
	http.HandleFunc("/articles/", viewArticle) //to view the articles
	http.HandleFunc("/edit/", editArticles)    //to edit the articles
	http.HandleFunc("/save/", saveArticles)    //to save the articles
	log.Fatal(http.ListenAndServe(":8080", nil))
}
