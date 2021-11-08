package bigfix

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceSingleAction() *schema.Resource {
	return &schema.Resource{
		Create: resourceSingleActionCreate,
		Read:   resourceSingleActionRead,
		Delete: resourceSingleActionDelete,

		Schema: map[string]*schema.Schema{
			"title": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"source_release_date": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"target": {
				Type:     schema.TypeSet,
				Required: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"computer_name": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"computer_id": {
							Type:     schema.TypeString,
							Optional: true,
							ForceNew: true,
						},
						"custom_relevance": {
							Type:     schema.TypeBool,
							Optional: true,
							ForceNew: true,
						},
					},
				},
			},

			"source_fixlet": {
				Type:     schema.TypeSet,
				Optional: true,
				ForceNew: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},
					},
				},
			},
			"relevance": {
				Type:     schema.TypeSet,
				Required: true,
				ForceNew: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

// resourceSingleActionSetToState : Function to set some resource properties
func resourceSingleActionSetToState(d *schema.ResourceData, fixletDetail *BESSingleActionRequest) {
	d.Set("title", fixletDetail.SingleAction.Title)
	relevances := make([]interface{}, 0)
	for _, v := range fixletDetail.SingleAction.Relevances {
		relevances = append(relevances, v.Text)
	}
	d.Set("relevance", schema.NewSet(schema.HashString, relevances))
	d.Set("source_release_date", fixletDetail.SingleAction.SourceReleaseDate)
	d.Set("source_release_date", fixletDetail.SingleAction.Actions)

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
				Default:  "application/x-SingleAction-Windows-Shell",
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
		"id":            fixletDetail.SingleAction.DefaultAction.ID,
		"mime_type":     fixletDetail.SingleAction.DefaultAction.ActionScript.MIMEType,
		"action_script": fixletDetail.SingleAction.DefaultAction.ActionScript.Text,
		"description": schema.NewSet(schema.HashResource(descriptionSchema), []interface{}{map[string]interface{}{
			"pre_link":  fixletDetail.SingleAction.DefaultAction.Description.PreLink,
			"link":      fixletDetail.SingleAction.DefaultAction.Description.Link,
			"post_link": fixletDetail.SingleAction.DefaultAction.Description.PostLink,
		}}),
	}

	actions := make([]interface{}, 0)

	for _, v := range fixletDetail.SingleAction.Actions {
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

func resourceSingleActionCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(Config)

	title := d.Get("title").(string)
	description := d.Get("description").(string)

	defaultAction := d.Get("default_action").(*schema.Set).List()[0].(map[string]interface{})
	relevances := make([]Relevance, 0)
	for _, v := range d.Get("relevance").(*schema.Set).List() {
		relevances = append(relevances, Relevance{
			Text: v.(string),
		})
	}
	defaultActionDescription := defaultAction["description"].(*schema.Set).List()[0].(map[string]interface{})

	actions := make([]SingleActionAction, 0)

	for _, v := range d.Get("action").(*schema.Set).List() {
		v := v.(map[string]interface{})
		actionDescription := v["description"].(*schema.Set).List()[0].(map[string]interface{})
		actions = append(actions, SingleActionAction{
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

	fixletRequest := BESSingleActionRequest{
		Xmlns_xsi:                     "http://www.w3.org/2001/XMLSchema-instance",
		Xsi_noNamespaceSchemaLocation: "BES.xsd",
		SingleAction: SingleAction{
			Title:             title,
			Description:       description,
			Source:            "Internal",
			SourceReleaseDate: d.Get("source_release_date").(string),
			Relevances:        relevances,
			Domain:            "BESC",
			Actions:           actions,
			DefaultAction: SingleActionAction{
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

	buff := ParseSingleActionXMLMarshal(fixletRequest)
	url := CreateAction(config.ServerIP, config.Port)
	response, err := config.BfxConnection(POST, url, bytes.NewReader(buff))

	if err != nil {
		log.Printf("[DEBUG] Error reading response body : %v\n", err)
		return err
	}

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

		actionID := responseStruct.Action.ID

		log.Printf("[DEBUG] SingleAction ID is : %s ", actionID)
		d.Set("title", responseStruct.Action.Name)
		d.SetId(actionID)
		return nil
	}
	return err
}

func resourceSingleActionRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(Config)
	log.Printf("[DEBUG] Read Action %s", d.Id())

	id := d.Id()
	url := GetActionDetailAPI(config.ServerIP, config.Port, id)
	response, err := config.BfxConnection(GET, url, nil)
	if err != nil {
		return err
	}

	// Get action
	var fixletDetailStruct BESSingleActionRequest
	if response != nil {
		defer response.Body.Close()

		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Printf("[DEBUG] Error reading response body : %v\n", err)
			return err
		}
		xml.Unmarshal([]byte(data), &fixletDetailStruct)
		resourceSingleActionSetToState(d, &fixletDetailStruct)
		return nil
	}
	return err
}

func resourceSingleActionDelete(d *schema.ResourceData, meta interface{}) error {

	resourceSingleActionRead(d, meta)
	if d.Id() == "" {
		log.Println("[ERROR] Action not found")
		return fmt.Errorf("Action not found")
	}
	config := meta.(Config)
	deleted, err := DeleteAction(config, d.Id())
	if deleted == true {
		log.Printf("[DEBUG] SingleAction having id [%s] deleted successfully \n", d.Id())
		return nil
	}

	log.Printf("[DEBUG] Error while deleting a MAG from bigfix : %v\n", err)
	return err
}
