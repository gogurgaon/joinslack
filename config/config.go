//Package config has the implementation for loading the config required by the application
//from environment variables, command line and config file.
//Config file is given priority over environment variables and command line args.
package config

import (
	"encoding/json"
	"flag"
	"html/template"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
 * This file has the implementation for loading the config from config file and from command line flags
 */

//PORT for the application. One can provide the port through cmd argument flags
var PORT = flag.Int("port", 80, "port for the chat application")

//STATIC is the static assets directory for the application.
//By default it is the one specified by the assets directory. If developer want to have
//to have something different, it can be done.
var STATIC = flag.String("static", "assets", "static assets directory for the application")

//CONFIGFILEPATH is the path to the config file for thr application
var CONFIGFILEPATH = flag.String("config-file", "./config.json", "config file location")

//WORKSPACE is the name of the workspace
var WORKSPACE = flag.String("workspace", "testingdevgroup", "name of the workspace where you want the users to jon")

//WORKSPACENAME is the name of the workspace to be used as display name
var WORKSPACENAME = flag.String("workspace-name", "Working Test Group", "display name of the workspace")

//WORKSPACELOGO is the url pointing towards the logo of the workspace team
var WORKSPACELOGO = flag.String("workspace-logo-url", "https://avatars3.githubusercontent.com/u/45892404?s=200&v=4", "url pointing to the default logo for your team")

//INVITEURLTEMPLATE is the url template to be hit for sending out url to join the slack workspace.
//It is a text template.
var INVITEURLTEMPLATE = template.Must(template.New("invite-url").Parse("https://{{.}}.slack.com/api/users.admin.invite"))

//INVITEURL is the url to be hit for sending out url to join the slack workspace.
var INVITEURL = flag.String("invite-url", "https://testingdevgroup.slack.com/api/users.admin.invite", "the entire invite url for the workspace")

//TOKEN is the token to be used for the application.
var TOKEN = flag.String("token", "", "token for authentication of the admin user")

//USERAPIURLTEMPLATE is the url template to be hit for getting the list of users in thew workspace.
//It is a text template.
var USERAPIURLTEMPLATE = template.Must(template.New("user-list-url").Parse("https://{{.Workspace}}.slack.com/api/users.list?token={{.Token}}&presence=1"))

//USERAPIURL is the url to be hit for getting the list of users in thew workspace.
var USERAPIURL = ""

//TEAMAPIURLTEMPLATE is the url template to be hit for getting the team info.
//It is a text template.
var TEAMAPIURLTEMPLATE = template.Must(template.New("team-url").Parse("https://{{.}}.slack.com/api/team.info?token={{.Token}}"))

//TEAMAPIURL is the url to be hit for getting the team info.
var TEAMAPIURL = ""

//Config has the basic config structure required by the application.
//config file however will have priority over Command line args
type Config struct {
	Port          int    //Port of the application
	Static        string //Static files location
	Workspace     string //Workspace is the name of the workspace
	WorkspaceName string //WorkspaceName is the display name of the workspace
	InviteURL     string //InviteUrl is the url to be used for inviting ther users
	Token         string //Token for the application
	WorkspaceLogo string //WorkspaceLogo is the url pointing towards the workspace logo
}

//LoadConfig will load the configuratoions required by the application
func LoadConfig() {
	/*
	 * We will parse the flags
	 * Then we will load the config file
	 */

	//parsing the flags
	flag.Parse()

	//loading the config file
	loadConfigFile()
}

func loadConfigFile() {
	/*
	 * Opening the config file
	 * Then we will decode the config file
	 * If the command line flags are not empty update them with config
	 * Environment variables get more priority over other methods with deployment in mind
	 */
	//Loading the config file
	file, err := os.Open(*CONFIGFILEPATH)
	if err != nil {
		log.Println("No config file found", err.Error())
		return
	}
	dec := json.NewDecoder(file)

	//decode the config
	config := Config{}
	err = dec.Decode(&config)
	if err != nil {
		log.Fatal("Error while reading the config file", err.Error())
	}

	//updating the command line args based on the config file
	//port
	if config.Port != 0 {
		*PORT = config.Port
	}
	if len(os.Getenv("PORT")) != 0 {
		p, err := strconv.Atoi(os.Getenv("PORT"))
		if err != nil {
			log.Fatal("Couldn't convert the port to an integer", err.Error())
		}
		*PORT = p
	}

	//static
	if len(config.Static) != 0 {
		*STATIC = config.Static
	}
	if len(os.Getenv("STATIC")) != 0 {
		*STATIC = os.Getenv("STATIC")
	}

	//workspace
	if len(config.Workspace) != 0 {
		*WORKSPACE = config.Workspace
	}
	if len(os.Getenv("WORKSPACE")) != 0 {
		*WORKSPACE = os.Getenv("WORKSPACE")
	}

	//workspace name
	if len(config.WorkspaceName) != 0 {
		*WORKSPACENAME = config.WorkspaceName
	}
	if len(os.Getenv("WORKSPACENAME")) != 0 {
		*WORKSPACENAME = os.Getenv("WORKSPACENAME")
	}

	//Checking whether the workspace subdomain name is empty or not
	if len(*WORKSPACE) == 0 {
		log.Fatal("Couldn't find the workspace id. Please provide it through config file or command line options")
	}

	//Workspace logo
	if len(config.WorkspaceLogo) != 0 {
		*WORKSPACELOGO = config.WorkspaceLogo
	}
	if len(os.Getenv("WORKSPACELOGO")) != 0 {
		*WORKSPACELOGO = os.Getenv("WORKSPACELOGO")
	}

	//invite url
	if len(config.InviteURL) != 0 {
		*INVITEURL = config.InviteURL
	}
	if len(os.Getenv("INVITEURL")) != 0 {
		*INVITEURL = os.Getenv("INVITEURL")
	}

	//if workspace is not empty we will use the template to set invite url
	b := strings.Builder{}
	INVITEURLTEMPLATE.Execute(&b, *WORKSPACE)
	temp := b.String()
	if len(temp) != 0 {
		*INVITEURL = temp
	}

	//token
	if len(config.Token) != 0 {
		*TOKEN = config.Token
	}
	if len(os.Getenv("TOKEN")) != 0 {
		*TOKEN = os.Getenv("TOKEN")
	}

	//verifying whether there exist a token
	if len(*TOKEN) == 0 {
		log.Fatal("Couldn't find the token. Please provide it through config file or command line options")
	}

	//setting the user api url
	b = strings.Builder{}
	USERAPIURLTEMPLATE.Execute(&b, struct {
		Workspace string
		Token     string
	}{*WORKSPACE, *TOKEN})
	temp = b.String()
	if len(temp) != 0 {
		USERAPIURL = temp
	}

	//setting the team api url
	b = strings.Builder{}
	TEAMAPIURLTEMPLATE.Execute(&b, struct {
		Workspace string
		Token     string
	}{*WORKSPACE, *TOKEN})
	temp = b.String()
	if len(temp) != 0 {
		TEAMAPIURL = temp
	}
}
