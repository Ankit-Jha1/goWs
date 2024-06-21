package handlers

import (
	"log"
	"net/http"

	"github.com/CloudyKit/jet/v6"
)

// created something with all the html
var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./html"),
	// restart nhi krna padega => standard library me h ye feature?
	jet.InDevelopmentMode(),
)

// we want to render template
func Home(w http.ResponseWriter, r *http.Request) {
	err := renderPage(w, "home.jet", nil)
	if err != nil {
		log.Println("render page err ==> ", err)
	}
}

// code to render template
// any fxn can use this
// saara html milaake 1 views variable bana, ab koi page render yha krenge
// writer lia h, means render matlab user ko bhej denge uske screen pr
func renderPage(w http.ResponseWriter, tmpl string, data jet.VarMap) error {
	// view template nikal rhe views me se
	view, err := views.GetTemplate(tmpl)
	if err != nil {
		log.Println("get template err ==> ", err)
		return err
	}
	// nil is context = not in this course
	// view me w, data for view de dia
	err = view.Execute(w, data, nil)
	if err != nil {
		log.Println("view execute err ==>", err)
		return err
	}
	return nil
}
