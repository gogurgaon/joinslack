//Join Slack application will send automatic invites to anyone who sign up to the workspace
package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/melvinodsa/joinslack/api"
	"github.com/melvinodsa/joinslack/config"
)

func main() {
	/*
	 * We will first load the configs
	 * Signup page as index page will be registered with the handler
	 * User sign up action will be registered with the handler
	 * Will serve the static contents
	 * Start the server
	 */
	//Loading the configs
	config.LoadConfig()

	//Signup page
	http.HandleFunc("/", api.SignupPage)

	//Signup request from user
	http.HandleFunc("/signup", api.Signup)

	//static file contents
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir(*config.STATIC))))

	//starting the server
	log.Println("Serving at localhost:", strconv.Itoa(*config.PORT), "...")
	log.Fatal(http.ListenAndServe(":"+strconv.Itoa(*config.PORT), nil))
}
