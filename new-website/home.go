package main

//import packages we are going to use
import (
	"html/template"
	"log"
	"net/http"
	"os"
)

//defining a struct which would define the basis of a page
type Page struct {
	Title string
	Body  []byte
}

//to reduce rewriting same code again, we package the templating part into a func
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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
	renderTemplate(w, "article", p)
}

//this will craete a new page or open the existing page both as a html form
//The .Title and .Body dotted identifiers in the template file refer to p.Title and p.Body
func editArticles(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):] //extract page title by removing slice of first part containing /articles/
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title} //we are creating a page if such a page doesn't exist
	}
	renderTemplate(w, "edit", p) //executes what has been written into the page
}

// this func hadnles the form submission, i.e., submission of articles edited by editArticles
func saveArticles(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/articles/"+title, http.StatusFound)
}

//creating a global variable to load template files into cache to prevent loading again and again
var templates = template.Must(template.ParseFiles("edit.html", "article.html"))

func main() {
	http.HandleFunc("/articles/", viewArticle) //to view the articles
	http.HandleFunc("/edit/", editArticles)    //to edit the articles
	http.HandleFunc("/save/", saveArticles)    //to save the articles
	log.Fatal(http.ListenAndServe(":8080", nil))
}
