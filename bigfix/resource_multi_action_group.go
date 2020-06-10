package bigfix

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceMultiActionGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceMultiActionGroupCreate,
		Read:   resourceMultiActionGroupRead,
		Delete: resourceMultiActionGroupDelete,

		Schema: map[string]*schema.Schema{
			// required values
			"input_file_name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"target_computer_name": {
				Type:        schema.TypeString,
				Description: "Computer name of target computer",
				Optional:    true,
				Computed:    true,
			},

			"target_computer_id": {
				Type:        schema.TypeString,
				Description: "Computer ID of target computer",
				Sensitive:   true,
				Required:    true,
				ForceNew:    true,
			},
			"site_name": {
				Type:        schema.TypeSet,
				Description: "name of site to get relevant fixlets",
				Required:    true,
				ForceNew:    true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			// computed values
			"title": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"relevance": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"state": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

// resourceFixletSetToState : Function to set some resource properties
func resourceMAGSetToState(d *schema.ResourceData, action *Action, state string) {
	d.Set("title", action.MultipleActionGroup.Title)
	d.Set("relevance", action.MultipleActionGroup.Relevance)
	d.Set("state", state)
}

func resourceMultiActionGroupCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(Config)

	inputFile := d.Get("input_file_name").(string)
	targetComputerID := d.Get("target_computer_id").(string)
	site := []string{}
	siteSet := d.Get("site_name").(*schema.Set)
	for _, v := range siteSet.List() {
		site = append(site, v.(string))
	}

	//get name of computer
	computerName := getNameOfComputer(config, targetComputerID)

	// get list of members of MAG
	var memberListStruct []SourcedMemberAction
	memberListStruct, err := SetSourcedMemberList(config, site, targetComputerID)
	if err != nil {
		return err
	}

	// get values from XML file
	file, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Printf("[DEBUG] Error while reading XML file : %v\n", err)
		return fmt.Errorf("error while reading file :{ %s }", err)
	}
	log.Println("[DEBUG] input file data : ", string(file))

	// set file data into struct
	var magStruct MAGFile
	err1 := xml.Unmarshal(file, &magStruct)
	if err1 != nil {
		log.Printf("[DEBUG] Error while parsing XML file : %v\n", err1)
		return fmt.Errorf("error while parsing XML file:{ %s }", err1)
	}

	// create buffer using file content and list data
	buff := ParseMAGXMLMarshal(magStruct, targetComputerID, site, memberListStruct)
	url := CreateMAGAPI(config.ServerIP, config.Port)
	response, err := config.BfxConnection(POST, url, bytes.NewReader(buff))
	if err != nil {
		return (err)
	}

	// Check whether action is created
	var responseStruct ActionCreationResponse
	if response != nil {
		defer response.Body.Close()
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Printf("[DEBUG] Error reading response body : %v\n", err)
			return err
		}

		xml.Unmarshal([]byte(data), &responseStruct)
		if responseStruct.Action.ID != "" {
			log.Println("[DEBUG] MAG created successfully !")
		}

		magID := responseStruct.Action.ID

		// get status of action
		actionStatus := GetActionStatus(config, magID)

		log.Printf("[DEBUG] Action ID is : %s ", magID)
		d.Set("title", responseStruct.Action.Name)
		d.Set("target_computer_name", computerName)
		d.Set("state", actionStatus)
		d.SetId(magID)
		return nil
	}
	return err
}

func resourceMultiActionGroupRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(Config)
	log.Printf("[DEBUG] Read Action %s", d.Id())

	id := d.Id()

	// get status of action
	actionStatus := GetActionStatus(config, id)

	// Get URL for Action Details using ID
	url := GetActionDetailAPI(config.ServerIP, config.Port, id)
	response, err := config.BfxConnection(GET, url, nil)
	if err != nil {
		return err
	}

	// Get action
	var actionDetailStruct Action
	if response != nil {
		defer response.Body.Close()

		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Printf("[DEBUG] Error reading response body : %v\n", err)
			return err
		}
		xml.Unmarshal([]byte(data), &actionDetailStruct)
		resourceMAGSetToState(d, &actionDetailStruct, actionStatus)
		return nil
	}
	return err
}

func resourceMultiActionGroupDelete(d *schema.ResourceData, meta interface{}) error {

	resourceMultiActionGroupRead(d, meta)
	if d.Id() == "" {
		log.Println("[ERROR] Action not found")
		return fmt.Errorf("Action not found")
	}
	config := meta.(Config)
	deleted, err := DeleteAction(config, d.Id())
	if deleted == true {
		log.Printf("[DEBUG] MAG having id [%s] deleted successfully \n", d.Id())
		return nil
	}

	log.Printf("[DEBUG] Error while deleting a MAG from bigfix : %v\n", err)
	return err
}
