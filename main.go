package main

import (
	"html/template"
	"log"
	"net/http"
	"strings"
	"time"
)

// MemberCard is a struct to store user information
type MemberCard struct {
	Signature string
	Email     string
	TimeOfReg string
	Number    int
}

// check calls log.Fatal on any non-nil error.
func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//Response1 is a struct used in rendering view.html.
type Response1 struct {
	Member    []MemberCard
	Error     bool
	TextError string
}

var NewResp = Response1{}

// viewHandler reads struct and displays it
func viewHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("view.html")
	log.Printf("request for view is: %v\n", *request)
	log.Printf("response for view is %v\n", writer)
	check(err)
	err = html.Execute(writer, NewResp)
	check(err)
}

// createHandler takes a POST request with a signature to add and  appends it to struct.
func createHandler(writer http.ResponseWriter, request *http.Request) {
	log.Printf("request for insert is: %v\n", *request)
	log.Printf("response for insert is %v\n", writer)
	signature := request.FormValue("signature")
	email := strings.ToLower(request.FormValue("email"))

	t := time.Now()
	NewResp.Error = false
	for _, v := range NewResp.Member {
		if email == v.Email {
			NewResp.Error = true
			NewResp.TextError = "User already exist!"
		}
	}
	if len(email) == 0 {
		NewResp.Error = true
		NewResp.TextError = "E-mail is empty!"
	}
	newMember := MemberCard{Signature: signature, Email: email, TimeOfReg: t.Format("02.01.2006"), Number: len(NewResp.Member) + 1}
	if NewResp.Error == false {
		NewResp.Member = append(NewResp.Member, newMember)
	}
	http.Redirect(writer, request, "/", http.StatusFound)
}
func main() {
	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/create", createHandler)
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	log.Fatal(err)
}
