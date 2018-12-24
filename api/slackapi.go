package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gogurgaon/joinslack/config"
)

/*
 * This file contains the utility functions to communicate with the slack apis
 */

//GetTeamInfo returns the info of the team.
//It will return nil as info if any error occurs and error.
//Else it will return info and error as nil.
func GetTeamInfo() (map[string]interface{}, error) {
	/*
	 * We will hit the slack api for getting the team info
	 * Will process the response from the api
	 */
	//hitting the team api
	log.Println("Going to fetch the team details")
	resp, err := http.Get(config.TEAMAPIURL)
	if err != nil {
		//handling the error
		log.Println("Error while getting the team info")
		return nil, err
	}

	//processing the response
	return processResponse(resp)
}

//Invite send join invites the user with the given email.
//It uses the invite url to send out invites. It will use post request to
//communicate with the slack server
func Invite(email string) error {
	/*
	 * We will hit the slack api for sending invites
	 * We will process the response from the slack api server
	 */
	//hitting the slack invite api
	log.Println("Going to send out invite to the user", "<"+email+">")
	log.Println("Invite URL:", *config.INVITEURL)
	resp, err := http.PostForm(*config.INVITEURL, url.Values{
		"email": {email}, "token": {*config.TOKEN},
	})
	if err != nil {
		//error while hitting the slack invite api
		log.Println("Error while sending invite")
		return err
	}

	//processing the response
	_, err = processResponse(resp)
	return err
}

//processResponse processes the response body passed on to it.
//If any error in the api response we will return the same
func processResponse(resp *http.Response) (map[string]interface{}, error) {
	/*
	 * We will first parse the response body to json
	 * Then will check whether ok key is present in the response
	 * If not then we will check the kind of error and return it
	 */
	//json parsing the response
	//reading the response
	response, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		//error while reading the response
		log.Println("Error while reading the response")
		return nil, err
	}
	//json decoding
	result := map[string]interface{}{}
	dec := json.NewDecoder(strings.NewReader(string(response)))
	err = dec.Decode(&result)
	if err != nil {
		//error while decoding the result
		log.Println("Error while parsing the response to json map[string]interface{}")
		return nil, err
	}

	log.Println("Got the response from the server", result)
	//checking ok is present in the response
	if fine, ok := result["ok"]; ok && fine.(bool) {
		//ok is present. will return the response
		return result, nil
	}

	//there is some error in the response
	log.Println("There is an error response from the slack api", result["error"])
	return result, checkError(result["error"].(string))
}

//checkError will check for the type of error based on the err argument passed to it
//It matches the err argument with the slack api response error codes
func checkError(err string) error {
	/*
	 * We will use switch case to identify the error
	 */
	switch err {
	case "invalid_auth":
		return errors.New("The authentication token for the application has expired. PLease contact the admin")
	case "already_in_team":
		return errors.New("You have already joined the team")
	case "invalid_email":
		return errors.New("The email provided by you is invalid. Please check the email")
	case "already_invited":
		return errors.New("You are already invited to the team. Please check your inbox")
	}
	return nil
}
