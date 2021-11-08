package bigfix

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceTask() *schema.Resource {
	return &schema.Resource{
		Create: resourceTaskCreate,
		Update: resourceTaskUpdate,
		Read:   resourceTaskRead,
		Delete: resourceTaskDelete,

		Schema: map[string]*schema.Schema{
			"title": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"source_release_date": {
				Type:     schema.TypeString,
				Required: true,
			},
			"default_action": {
				Type:     schema.TypeSet,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "Action1",
						},
						"mime_type": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "application/x-Fixlet-Windows-Shell",
						},
						"action_script": {
							Type:     schema.TypeString,
							Required: true,
						},
						"description": {
							Type:     schema.TypeSet,
							MaxItems: 1,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"pre_link": {
										Type:     schema.TypeString,
										Required: true,
									},
									"link": {
										Type:     schema.TypeString,
										Required: true,
									},
									"post_link": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},
					},
				},
			},

			"action": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Required: true,
						},
						"mime_type": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "application/x-Fixlet-Windows-Shell",
						},
						"action_script": {
							Type:     schema.TypeString,
							Required: true,
						},
						"description": {
							Type:     schema.TypeSet,
							MaxItems: 1,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"pre_link": {
										Type:     schema.TypeString,
										Required: true,
									},
									"link": {
										Type:     schema.TypeString,
										Required: true,
									},
									"post_link": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},
					},
				},
			},

			"site_name": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Default:  "",
			},

			"site_type": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"relevance": {
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

// resourceTaskSetToState : Function to set some resource properties
func resourceTaskSetToState(d *schema.ResourceData, fixletDetail *BESTaskRequest) {
	d.Set("title", fixletDetail.Task.Title)
	relevances := make([]interface{}, 0)
	for _, v := range fixletDetail.Task.Relevances {
		relevances = append(relevances, v.Text)
	}
	d.Set("relevance", schema.NewSet(schema.HashString, relevances))
	d.Set("source_release_date", fixletDetail.Task.SourceReleaseDate)
	d.Set("source_release_date", fixletDetail.Task.Actions)

	descriptionSchema := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"pre_link": {
				Type:     schema.TypeString,
				Required: true,
			},
			"link": {
				Type:     schema.TypeString,
				Required: true,
			},
			"post_link": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}

	defaultActionSchema := &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Action1",
			},
			"mime_type": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "application/x-Fixlet-Windows-Shell",
			},
			"action_script": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeSet,
				MaxItems: 1,
				Required: true,
				Elem:     descriptionSchema,
			},
		},
	}

	defaultAction := map[string]interface{}{
		"id":            fixletDetail.Task.DefaultAction.ID,
		"mime_type":     fixletDetail.Task.DefaultAction.ActionScript.MIMEType,
		"action_script": fixletDetail.Task.DefaultAction.ActionScript.Text,
		"description": schema.NewSet(schema.HashResource(descriptionSchema), []interface{}{map[string]interface{}{
			"pre_link":  fixletDetail.Task.DefaultAction.Description.PreLink,
			"link":      fixletDetail.Task.DefaultAction.Description.Link,
			"post_link": fixletDetail.Task.DefaultAction.Description.PostLink,
		}}),
	}

	actions := make([]interface{}, 0)

	for _, v := range fixletDetail.Task.Actions {
		actions = append(actions, map[string]interface{}{
			"id":            v.ID,
			"mime_type":     v.ActionScript.MIMEType,
			"action_script": v.ActionScript.Text,
			"description": schema.NewSet(schema.HashResource(descriptionSchema), []interface{}{map[string]interface{}{
				"pre_link":  v.Description.PreLink,
				"link":      v.Description.Link,
				"post_link": v.Description.PostLink,
			}}),
		})
	}

	d.Set("default_action", schema.NewSet(schema.HashResource(defaultActionSchema), []interface{}{defaultAction}))
	d.Set("action", schema.NewSet(schema.HashResource(defaultActionSchema), actions))
}

func resourceTaskCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(Config)

	title := d.Get("title").(string)
	description := d.Get("description").(string)
	siteName := d.Get("site_name").(string)
	siteType := d.Get("site_type").(string)

	defaultAction := d.Get("default_action").(*schema.Set).List()[0].(map[string]interface{})
	relevances := make([]Relevance, 0)
	for _, v := range d.Get("relevance").(*schema.Set).List() {
		relevances = append(relevances, Relevance{
			Text: v.(string),
		})
	}
	defaultActionDescription := defaultAction["description"].(*schema.Set).List()[0].(map[string]interface{})

	actions := make([]FixletAction, 0)

	for _, v := range d.Get("action").(*schema.Set).List() {
		v := v.(map[string]interface{})
		actionDescription := v["description"].(*schema.Set).List()[0].(map[string]interface{})
		actions = append(actions, FixletAction{
			ID: v["id"].(string),
			ActionScript: ActionScript{
				MIMEType: v["mime_type"].(string),
				Text:     v["action_script"].(string),
			},
			Description: ActionDescription{
				PostLink: actionDescription["post_link"].(string),
				PreLink:  actionDescription["pre_link"].(string),
				Link:     actionDescription["link"].(string),
			},
		})
	}

	fixletRequest := BESTaskRequest{
		Xmlns_xsi:                     "http://www.w3.org/2001/XMLSchema-instance",
		Xsi_noNamespaceSchemaLocation: "BES.xsd",
		Task: Task{
			Title:             title,
			Description:       description,
			Source:            "Internal",
			SourceReleaseDate: d.Get("source_release_date").(string),
			Relevances:        relevances,
			Domain:            "BESC",
			Actions:           actions,
			DefaultAction: FixletAction{
				ID: defaultAction["id"].(string),
				ActionScript: ActionScript{
					MIMEType: defaultAction["mime_type"].(string),
					Text:     defaultAction["action_script"].(string),
				},
				Description: ActionDescription{
					PostLink: defaultActionDescription["post_link"].(string),
					PreLink:  defaultActionDescription["pre_link"].(string),
					Link:     defaultActionDescription["link"].(string),
				},
			},
		},
	}

	buff := ParseTaskXMLMarshal(fixletRequest)
	url := CreateTask(config.ServerIP, config.Port, siteName, siteType)
	response, err := config.BfxConnection(POST, url, bytes.NewReader(buff))

	if err != nil {
		log.Printf("[DEBUG] Error reading response body : %v\n", err)
		return err
	}

	var responseStruct TaskCreationResponse
	if response != nil {
		defer response.Body.Close()
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Printf("[DEBUG] Error reading response body : %v\n", err)
			return err
		}

		xml.Unmarshal([]byte(data), &responseStruct)
		if responseStruct.Task.ID != "" {
			log.Println("[DEBUG] MAG created successfully !")
		}

		fixletID := responseStruct.Task.ID

		log.Printf("[DEBUG] Task ID is : %s ", fixletID)
		d.Set("title", responseStruct.Task.Name)
		d.SetId(fixletID)
		return nil
	}
	return err
}

func resourceTaskUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(Config)

	title := d.Get("title").(string)
	description := d.Get("description").(string)
	siteName := d.Get("site_name").(string)
	siteType := d.Get("site_type").(string)

	defaultAction := d.Get("default_action").(*schema.Set).List()[0].(map[string]interface{})
	relevances := make([]Relevance, 0)
	for _, v := range d.Get("relevance").(*schema.Set).List() {
		relevances = append(relevances, Relevance{
			Text: v.(string),
		})
	}
	defaultActionDescription := defaultAction["description"].(*schema.Set).List()[0].(map[string]interface{})

	actions := make([]FixletAction, 0)

	for _, v := range d.Get("action").(*schema.Set).List() {
		v := v.(map[string]interface{})
		actionDescription := v["description"].(*schema.Set).List()[0].(map[string]interface{})
		actions = append(actions, FixletAction{
			ID: v["id"].(string),
			ActionScript: ActionScript{
				MIMEType: v["mime_type"].(string),
				Text:     v["action_script"].(string),
			},
			Description: ActionDescription{
				PostLink: actionDescription["post_link"].(string),
				PreLink:  actionDescription["pre_link"].(string),
				Link:     actionDescription["link"].(string),
			},
		})
	}

	fixletRequest := BESTaskRequest{
		Xmlns_xsi:                     "http://www.w3.org/2001/XMLSchema-instance",
		Xsi_noNamespaceSchemaLocation: "BES.xsd",
		Task: Task{
			Title:             title,
			Description:       description,
			Source:            "Internal",
			SourceReleaseDate: d.Get("source_release_date").(string),
			Relevances:        relevances,
			Domain:            "BESC",
			Actions:           actions,
			DefaultAction: FixletAction{
				ID: defaultAction["id"].(string),
				ActionScript: ActionScript{
					MIMEType: defaultAction["mime_type"].(string),
					Text:     defaultAction["action_script"].(string),
				},
				Description: ActionDescription{
					PostLink: defaultActionDescription["post_link"].(string),
					PreLink:  defaultActionDescription["pre_link"].(string),
					Link:     defaultActionDescription["link"].(string),
				},
			},
		},
	}

	buff := ParseTaskXMLMarshal(fixletRequest)
	url := UpdateTask(config.ServerIP, config.Port, d.Id(), siteName, siteType)
	response, err := config.BfxConnection(PUT, url, bytes.NewReader(buff))

	if err != nil {
		log.Printf("[DEBUG] Error reading response body : %v\n", err)
		return err
	}

	if response != nil {
		defer response.Body.Close()
		_, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Printf("[DEBUG] Error reading response body : %v\n", err)
			return err
		}
		return nil
	}
	return err
}

func resourceTaskRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(Config)
	log.Printf("[DEBUG] Read Action %s", d.Id())

	id := d.Id()
	siteName := d.Get("site_name").(string)
	siteType := d.Get("site_type").(string)
	url := GetTaskDetailAPI(config.ServerIP, config.Port, id, siteName, siteType)
	response, err := config.BfxConnection(GET, url, nil)
	if err != nil {
		return err
	}

	// Get action
	var fixletDetailStruct BESTaskRequest
	if response != nil {
		defer response.Body.Close()

		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Printf("[DEBUG] Error reading response body : %v\n", err)
			return err
		}
		xml.Unmarshal([]byte(data), &fixletDetailStruct)
		resourceTaskSetToState(d, &fixletDetailStruct)
		return nil
	}
	return err
}

func resourceTaskDelete(d *schema.ResourceData, meta interface{}) error {

	resourceTaskRead(d, meta)
	if d.Id() == "" {
		log.Println("[ERROR] Action not found")
		return fmt.Errorf("Action not found")
	}
	config := meta.(Config)
	siteName := d.Get("site_name").(string)
	siteType := d.Get("site_type").(string)
	deleted, err := DeleteTask(config, d.Id(), siteName, siteType)
	if deleted == true {
		log.Printf("[DEBUG] Task having id [%s] deleted successfully \n", d.Id())
		return nil
	}

	log.Printf("[DEBUG] Error while deleting a MAG from bigfix : %v\n", err)
	return err
}
