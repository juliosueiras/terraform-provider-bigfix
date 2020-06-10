package bigfix

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/hashicorp/terraform/helper/schema"
)

func dataSourceComputer() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceComputerRead,

		Schema: map[string]*schema.Schema{
			// required value
			"name": {
				Type:        schema.TypeString,
				Description: "Name of computer",
				Required:    true,
				ForceNew:    true,
			},
			// computed values
			"id": {
				Type:        schema.TypeString,
				Description: "ID of computer",
				Computed:    true,
				Optional:    true,
			},
			"ip_address": {
				Type:        schema.TypeString,
				Description: "IP Address of computer",
				Computed:    true,
				Optional:    true,
			},
			"os": {
				Type:        schema.TypeString,
				Description: "Operating System of computer",
				Computed:    true,
				Optional:    true,
			},
			"cpu": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"last_report_time": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"dns_name": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"ram": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"subnet_address": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"computer_type": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"relay": {
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
		},
	}
}

func dataSourceComputerRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(Config)

	computerName := d.Get("name").(string)

	// Get details of computer with computer name
	log.Printf("[INFO] Reading Bigfix Computers")
	url := GetComputerDetailAPI(config.ServerIP, config.Port, computerName)
	Response, err := config.BfxConnection(GET, url, nil)
	if err != nil {
		return err
	}

	var computerdetailstruct ComputerDetails
	if Response != nil {
		defer Response.Body.Close()

		data, err := ioutil.ReadAll(Response.Body)
		if err != nil {
			fmt.Println(err)
		}

		//log.Println("[DEBUG] Response is :")
		//println(string(data))
		xml.Unmarshal([]byte(data), &computerdetailstruct)
		if computerdetailstruct.Query.Error != "" {
			return fmt.Errorf("No computer found with name `%s`.\n\t Error of Relevance Query is : { %s } ", computerName, computerdetailstruct.Query.Error)
		}

		d.Set("id", computerdetailstruct.Query.Result.Tuple.Answer[1].Text)
		d.Set("name", computerdetailstruct.Query.Result.Tuple.Answer[0].Text)
		d.Set("ip_address", computerdetailstruct.Query.Result.Tuple.Answer[8].Text)
		d.Set("os", computerdetailstruct.Query.Result.Tuple.Answer[2].Text)
		d.Set("cpu", computerdetailstruct.Query.Result.Tuple.Answer[5].Text)
		d.Set("last_report_time", computerdetailstruct.Query.Result.Tuple.Answer[7].Text)
		d.Set("dns_name", computerdetailstruct.Query.Result.Tuple.Answer[4].Text)
		d.Set("relay", computerdetailstruct.Query.Result.Tuple.Answer[6].Text)
		d.Set("ram", computerdetailstruct.Query.Result.Tuple.Answer[3].Text)

		d.SetId(computerdetailstruct.Query.Result.Tuple.Answer[1].Text)

	}
	return nil
}
