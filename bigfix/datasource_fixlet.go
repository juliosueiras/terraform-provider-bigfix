package bigfix

import (
	"encoding/xml"
	"fmt"
	"github.com/hashicorp/terraform/helper/schema"
	"io/ioutil"
	"log"
)

func dataSourceFixlet() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceFixletRead,

		Schema: map[string]*schema.Schema{
			//required value
			"name": {
				Type:        schema.TypeString,
				Description: "title of fixlet",
				Required:    true,
				ForceNew:    true,
			},
			// optional value
			"id": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"site_name": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"default_action": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
		},
	}
}

func dataSourceFixletRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(Config)

	name := d.Get("name").(string)

	// Check whether Fixlet exists with provided name
	log.Println("[DEBUG] Reading List of Fixlets...")
	url := GetFixletMetaDataAPI(config.ServerIP, config.Port, name)
	response, err := config.BfxConnection(GET, url, nil)
	if err != nil {
		return err
	}

	var fixletMetaDataStruct FixletMetaData
	if response != nil {
		defer response.Body.Close()

		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		}
		log.Println("[DEBUG] response is :")
		log.Println(string(data))
		xml.Unmarshal([]byte(data), &fixletMetaDataStruct)
		if fixletMetaDataStruct.Query.Result.Tuple.Text == "" {
			return fmt.Errorf("No Fixlet found with name `%s`", name)
		}

		d.Set("name", fixletMetaDataStruct.Query.Result.Tuple.Answer[2].Text)
		d.Set("id", fixletMetaDataStruct.Query.Result.Tuple.Answer[1].Text)
		d.Set("site_name", fixletMetaDataStruct.Query.Result.Tuple.Answer[0].Text)
		d.Set("default_action", fixletMetaDataStruct.Query.Result.Tuple.Answer[3].Text)

		d.SetId(fixletMetaDataStruct.Query.Result.Tuple.Answer[1].Text)

	}
	return nil
}
