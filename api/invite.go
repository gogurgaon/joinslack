//Package api has the web api as well as go sdk for using slack apis
package api

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gogurgaon/joinslack/config"
)

/*
 * This file contains the api for automatically inviting users to workspace
 */

//SignupPage displays the sign up page for the workspace
func SignupPage(res http.ResponseWriter, req *http.Request) {
	/*
	 * Then we will parse the template for signup page
	 * Will return the response as the template page
	 */
	//paring the template
	t := template.New("signup.html")                   // Create a template.
	te, err := t.ParseFiles("./templates/signup.html") // Parse template file.
	if err != nil {
		log.Println("Error while loading the signup template file", err.Error())
	}
	err = te.Execute(res, *config.WORKSPACENAME) // merge.
	if err != nil {
		log.Println("Error while executing the signup template", err.Error())
	}
}

//Signup is the user request to signup for a workspace
func Signup(res http.ResponseWriter, req *http.Request) {
	/*
	 * Then we will parse the template for thank you page
	 * Will return the response as the template page
	 */
	//paring the template
	t := template.New("thanks.html")                   // Create a template.
	te, err := t.ParseFiles("./templates/thanks.html") // Parse template file.
	if err != nil {
		log.Println("Error while loading the thanks template file", err.Error())
	}
	err = te.Execute(res, *config.WORKSPACENAME) // merge.
	if err != nil {
		log.Println("Error while executing the thanks template", err.Error())
	}
}
