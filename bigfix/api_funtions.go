package bigfix

import (
	"log"
	"net/url"
)

// ConnectBigFixAPI : Bigfix Authentication.
//
// BigFix API Doc :  https://bigfix.me/restapi/?id=300
func ConnectBigFixAPI(serverIP string, port string) string {
	connect := "https://" + serverIP + ":" + port + "/api/login"
	log.Println("[DEBUG] Getting response from URL : \n\n", connect)
	return connect
}

// CreateMAGAPI : Create Multiple Action Group.
//
// BigFix API Doc : https://bigfix.me/restapi/?id=261
func CreateMAGAPI(serverIP string, port string) string {
	connect := "https://" + serverIP + ":" + port + "/api/actions"
	log.Println("[DEBUG] Getting response from URL : \n\n", connect)
	return connect
}

// GetActionStatusAPI : Get Status of Action.
//
// BigFix API Doc : https://bigfix.me/restapi/?id=272
func GetActionStatusAPI(serverIP string, port string, actionID string) string {
	connect := "https://" + serverIP + ":" + port + "/api/action/" + actionID + "/status"
	log.Println("[DEBUG] Getting response from URL : \n\n", connect)
	return connect
}

// GetActionDetailAPI : Get Action detail of given ID.
//
// BigFix API Doc : https://bigfix.me/restapi/?id=295
func GetActionDetailAPI(serverIP string, port string, id string) string {
	connect := "https://" + serverIP + ":" + port + "/api/action/" + id
	log.Println("[DEBUG] Getting response from URL : \n\n", connect)
	return connect
}

// GetDeleteActionAPI : Get fixlet details.
//
// BigFix API Doc : https://bigfix.me/restapi/?id=271
func GetDeleteActionAPI(serverIP string, port string, id string) string {
	connect := "https://" + serverIP + ":" + port + "/api/action/" + id
	log.Println("[DEBUG] Getting response from URL : \n\n", connect)
	return connect

}

// GetAllSitesAPI : Get list of all sites.
//
// BigFix API Doc : https://bigfix.me/restapi/?id=316
func GetAllSitesAPI(serverIP string, port string) string {
	connect := "https://" + serverIP + ":" + port + "/api/sites"
	log.Println("[DEBUG] Getting response from URL : \n\n", connect)
	return connect
}

// GetComputerDetailAPI : Get computer details.
// Following API's are created with the help of relevance language queries
//
// BigFix API Doc : https://bigfix.me/restapi/?id=306
func GetComputerDetailAPI(serverIP string, port string, computerName string) string {
	query := "relevance=(name+of+it,values+of+property+results+whose+(name+of+property+of+it=\"ID\")+of+it,values+of+property+results+whose+(name+of+property+of+it=\"OS\")+of+it,values+of+property+results+whose+(name+of+property+of+it=\"RAM\")+of+it,values+of+property+results+whose+(name+of+property+of+it=\"DNS+Name\")+of+it,values+of+property+results+whose+(name+of+property+of+it=\"CPU\")+of+it,values+of+property+results+whose+(name+of+property+of+it=\"Relay\")+of+it,values+of+property+results+whose+(name+of+property+of+it=\"Last+Report+Time\")+of+it,(concatenations+of+(values+of+it+as+string))+of+property+results+whose+(name+of+property+of+it=\"IP+Address\")+of+it)+of+bes+computer+whose+(name+of+it+is+\"" + url.QueryEscape(computerName) + "\")"
	connect := "https://" + serverIP + ":" + port + "/api/query?" + query
	log.Println("[DEBUG] Getting response from URL : \n\n", connect)
	return connect
}

// GetComputerCountAPI : Get Computer details.
//
// BigFix API Doc : https://bigfix.me/restapi/?id=306
func GetComputerCountAPI(serverIP string, port string) string {
	query := "relevance=number+of+bes+computers"
	connect := "https://" + serverIP + ":" + port + "/api/query?" + query
	log.Println("[DEBUG] Getting response from URL : \n\n", connect)
	return connect
}

// GetFixletMetaDataAPI : Get fixlet details.
//
// BigFix API Doc : https://bigfix.me/restapi/?id=306
func GetFixletMetaDataAPI(serverIP string, port string, name string) string {
	query := "relevance=(name+of+site+of+it,id+of+it,+name+of+it,content+id+of+default+action+of+it)+of+bes+fixlets+whose+(name+of+it+is+\"" + url.QueryEscape(name) + "\")"
	connect := "https://" + serverIP + ":" + port + "/api/query?" + query
	log.Println("[DEBUG] Getting response from URL : \n\n", connect)
	return connect
}

// GetRelevantFixletsAPI : Get relevant Fixlets for particular computer.
//
// BigFix API Doc : https://bigfix.me/restapi/?id=306
func GetRelevantFixletsAPI(serverIP string, port string, siteName string, computerID string) string {
	query := "relevance=(name+of+site+of+it,id+of+it,+name+of+it)+of+relevant+fixlets+whose+(name+of+site+of+it+is+\"" + url.QueryEscape(siteName) + "\"+and+exists+content+id+of+Default+Action+of+it)+of+bes+computer+whose+(id+of+it+is+" + url.QueryEscape(computerID) + ")"
	connect := "https://" + serverIP + ":" + port + "/api/query?" + query
	log.Println("[DEBUG] Getting response from URL : \n\n", connect)
	return connect
}

// GetCountofRelevantFixletsAPI : Get Count of relevant fixlets for particular computer
//
// BigFix API Doc : https://bigfix.me/restapi/?id=306
func GetCountofRelevantFixletsAPI(serverIP string, port string, siteName string, computerID string) string {
	query := "relevance=number+of+(name+of+site+of+it,id+of+it,+name+of+it)+of+relevant+fixlets+whose+(name+of+site+of+it+is+\"" + url.QueryEscape(siteName) + "\"+and+exists+content+id+of+Default+Action+of+it)+of+bes+computer+whose+(id+of+it+is+" + url.QueryEscape(computerID) + ")"
	connect := "https://" + serverIP + ":" + port + "/api/query?" + query
	log.Println("[DEBUG] Getting response from URL : \n\n", connect)
	return connect
}

// GetComputerNameAPI : Get list of all sites.
//
// BigFix API Doc : https://bigfix.me/restapi/?id=306
func GetComputerNameAPI(serverIP string, port string, id string) string {
	query := "relevance=(name+of+it)+of+bes+computer+whose+(id+of+it+is+" + id + "+)"
	connect := "https://" + serverIP + ":" + port + "/api/query?" + query
	log.Println("[DEBUG] Getting response from URL : \n\n", connect)
	return connect
}
