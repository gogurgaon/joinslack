package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
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
	 * We will read the body of the response
	 * Will parse the response into json
	 */
	//hitting the team api
	log.Println("Going to fetch the team details")
	resp, err := http.Get(config.TEAMAPIURL)
	if err != nil {
		//handling the error
		log.Println("Error while getting the team info")
		return nil, err
	}

	//ready the response
	robots, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		//error while reading the responser
		log.Println("Error while reading the response body of team api")
		return nil, err
	}

	//json parsing the result
	result := map[string]interface{}{}
	dec := json.NewDecoder(strings.NewReader(string(robots)))
	err = dec.Decode(&result)
	if err != nil {
		//error while decoding the result
		log.Println("Error while parsing the response to json map[string]interface{}")
		return nil, err
	}

	return result, nil
}
