package bigfix

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
)

// API methods

// GET :
var GET = "GET"

// POST :
var POST = "POST"

// PUT :
var PUT = "PUT"

// DELETE :
var DELETE = "DELETE"

// Config : BigFix login credentials structure
type Config struct {
	ServerIP string
	Port     string
	Username string
	Password string
}

// BFXConfig : BFXConfig is per-provider, specifies where to connect to bigfix
func BFXConfig(d *schema.ResourceData) (interface{}, error) {

	serverip := d.Get("server").(string)
	// Check if field is not empty
	if serverip == "" {
		return nil, fmt.Errorf("Bigfix server ip not found ")
	}

	username := d.Get("username").(string)
	// Check if field is not empty
	if username == "" {
		return nil, fmt.Errorf("Bigfix  username not found")
	}

	password := d.Get("password").(string)
	// Check if field is not empty
	if password == "" {
		return nil, fmt.Errorf("Bigfix  password not found")
	}

	port := d.Get("port").(string)
	// Check if field is not empty
	if port == "" {
		return nil, fmt.Errorf("Bigfix server port not found")
	}

	// Check connection to bfx server
	bfxURL := ConnectBigFixAPI(serverip, port) //get url for connection
	_, err := http.NewRequest(GET, bfxURL, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	config := Config{
		ServerIP: serverip,
		Username: username,
		Password: password,
		Port:     port,
	}

	return config, nil
}

// BfxConnection : Will Connect to BigFix Server
func (config Config) BfxConnection(method string, url string, buffer io.Reader) (*http.Response, error) {

	//Initialize HTTPS client to skip SSL certificate verification
	flagSSL := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	skipSslVerify := &http.Client{
		Transport: flagSSL,
		// set default timeout to 300 secs
		Timeout: 300 * time.Second,
	}
	log.Println("[DEBUG] Initialize HTTPS client to skip SSL certificate verification.")

	//make request
	request, err := http.NewRequest(method, url, buffer)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
		return nil, err
	}
	log.Println("[DEBUG] make request to the BigFix.")

	//Configure basic authentication and content-type
	request.SetBasicAuth(config.Username, config.Password)
	request.Header.Set("Content-Type", "application/xml")

	// Make request to BigFix Rest API
	response, err1 := skipSslVerify.Do(request)
	if err1 != nil {
		return nil, err1
	}

	//check the status of the request i.e. response

	// Status = 200 //update 2XX to 3XX
	if response.StatusCode >= http.StatusOK && response.StatusCode < 400 {
		log.Println("[DEBUG] HTTP response OK  ")
		return response, nil
	}

	// catch all other status code
	return nil, fmt.Errorf("\tMethod: %s \n \tURL: %s \n \tStatusCode: [%d] \n \tStatus: %s", method, url, response.StatusCode, response.Status)
}
