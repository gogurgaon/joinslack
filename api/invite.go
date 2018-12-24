//Package api has the web api as well as go sdk for using slack apis
package api

import (
	"errors"
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
	 * We will parse the email and the user's full name
	 * Then we will use the slack api to invite the user
	 * Will return the response as the template page
	 */
	//paring the template
	t := template.New("thanks.html")                   // Create a template.
	te, err := t.ParseFiles("./templates/thanks.html") // Parse template file.
	if err != nil {
		log.Println("Error while loading the thanks template file", err.Error())
	}

	//parsing the user's full name and email
	email := req.PostFormValue("email")
	name := req.PostFormValue("name")
	if len(email) == 0 || len(name) == 0 {
		ErrorResponse(res, errors.New("Couldn't find the email and full name"))
		return
	}

	//inviting the user
	err = Invite(email)
	if err != nil {
		ErrorResponse(res, err)
		return
	}

	//finally giving out the response template
	err = te.Execute(res, struct {
		Workspace string
		Message   string
	}{
		*config.WORKSPACENAME,
		"Thanks for joining " + *config.WORKSPACENAME + ". Check your email for the invite",
	}) // merge.
	if err != nil {
		log.Println("Error while executing the thanks template", err.Error())
	}
}

//ErrorResponse will return the error response to the client
func ErrorResponse(res http.ResponseWriter, er error) {
	/*
	 * Then we will parse the template for error page
	 * Will return the response as the error page
	 */
	//paring the template
	t := template.New("error.html")                   // Create a template.
	te, err := t.ParseFiles("./templates/error.html") // Parse template file.
	if err != nil {
		log.Println("Error while loading the error template file", err.Error())
	}
	err = te.Execute(res, er.Error()) // merge.
	if err != nil {
		log.Println("Error while executing the error template", err.Error())
	}
}
