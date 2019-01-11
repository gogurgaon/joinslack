# joinslack ![](https://travis-ci.com/melvinodsa/joinslack.svg?branch=master)
Automatic slack sign up

## Installation
```sh
go get github.com/gogurgaon/joinslack
cd $GOPATH/src/github.com/gogurgaon/joinslack
dep ensure #if you want to use dep tool.
go install
joinslack
```

## Help
Run the following command to get help on how to use the application.
```sh
joinslack --help
```


## Configuration
The application requires few configurations. It can be configured via a config.json file at the location from which the application is run.
These configurations can also be set from command line.

*   **Token** (mandatory) - Token with which the application has to authenticate with slack. It can be generated via this [link](https://api.slack.com/custom-integrations/legacy-tokens).
*   **Workspace** (mandatory) - Name of your workspace which is used in the sub-domain of slack. Eg. testingdevgroup.slack.com. testingdevgroup is the workspace name
*   **WorkspaceName** (mandatory) - Display name of your workspace
*   **InviteURL**  (mandatory) - Custom invite url
*   Static -   Relative directory where the static asset files are kept
*   Port - Port on which application has to run
*   ConfigFile - Relative location of the config file

An example configuration looks like this :-

```json
{
	"Port": 9090,
	"Workspace": "gogurgaon",
	"WorkspaceName": "Gurgaon Golang Meetup",
	"Token": "2349d3bn90jkkasd9034rj-not-anactualtoken"
}
```
