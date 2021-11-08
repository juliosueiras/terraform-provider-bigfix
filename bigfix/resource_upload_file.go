package bigfix

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceUploadFile() *schema.Resource {
	return &schema.Resource{
		Create: resourceUploadFileCreate,
		Read:   resourceUploadFileRead,
		Delete: resourceUploadFileDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
			},
			"filename": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"is_private": {
				Type:     schema.TypeBool,
				Default:  false,
				Optional: true,
				ForceNew: true,
			},
			"sha1": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"sha256": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"size": {
				Type:     schema.TypeInt,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"url": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
				Optional: true,
			},
		},
	}
}

// resourceUploadFileSetToState : Function to set some resource properties
func resourceUploadFileSetToState(d *schema.ResourceData, fileDetail *FileCreationResponse) error {
	d.Set("name", fileDetail.FileUpload.Name)
	d.Set("url", fileDetail.FileUpload.URL)
	d.Set("sha1", fileDetail.FileUpload.SHA1)
	d.Set("sha256", fileDetail.FileUpload.SHA256)
	d.Set("size", fileDetail.FileUpload.Size)
	return nil
}

func resourceUploadFileCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(Config)

	filename := d.Get("filename").(string)
	isPrivate := d.Get("is_private").(bool)

	url := CreateUploadFile(config.ServerIP, config.Port, isPrivate)
	response, err := config.newfileUploadRequest(url, map[string]string{}, "file", filename)

	if err != nil {
		log.Printf("[DEBUG] Error reading response body : %v\n", err)
		return err
	}

	var responseStruct FileCreationResponse
	if response != nil {
		defer response.Body.Close()
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Printf("[DEBUG] Error reading response body : %v\n", err)
			return err
		}

		xml.Unmarshal([]byte(data), &responseStruct)
		if responseStruct.FileUpload.Name != "" {
			log.Println("[DEBUG] File created successfully !")
		}

		fileLoc := responseStruct.FileUpload.Name

		log.Printf("[DEBUG] UploadFile ID is : %s ", fileLoc)
		d.SetId(fileLoc)

		if !isPrivate {
			url := GetUploadFileDetailAPI(config.ServerIP, config.Port, fileLoc) + "?private=0"
			config.BfxConnection(POST, url, nil)
		}

		return nil
	}
	return err
}

func resourceUploadFileRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(Config)
	log.Printf("[DEBUG] Read Action %s", d.Id())

	id := d.Id()
	url := GetUploadFileDetailAPI(config.ServerIP, config.Port, id)
	response, err := config.BfxConnection(GET, url, nil)
	if err != nil {
		return err
	}

	// Get action
	var fileResponse FileCreationResponse
	if response != nil {
		defer response.Body.Close()

		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			log.Printf("[DEBUG] Error reading response body : %v\n", err)
			return err
		}
		xml.Unmarshal([]byte(data), &fileResponse)
		return resourceUploadFileSetToState(d, &fileResponse)
	}
	return err
}

func resourceUploadFileDelete(d *schema.ResourceData, meta interface{}) error {

	resourceUploadFileRead(d, meta)
	if d.Id() == "" {
		log.Println("[ERROR] Action not found")
		return fmt.Errorf("Action not found")
	}
	config := meta.(Config)
	deleted, err := DeleteUploadFile(config, d.Id())
	if deleted == true {
		log.Printf("[DEBUG] UploadFile having id [%s] deleted successfully \n", d.Id())
		return nil
	}

	log.Printf("[DEBUG] Error while deleting a MAG from bigfix : %v\n", err)
	return err
}
