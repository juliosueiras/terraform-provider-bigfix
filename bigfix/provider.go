package bigfix

import (
	"log"
	"os"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

// Provider : Provider returns a terraform.ResourceProvider.
// Contains registry of Data sources and Resources
func Provider() terraform.ResourceProvider {
	// actual provider
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"port": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Bigfix server port number",
				DefaultFunc: schema.EnvDefaultFunc("BFX_PORT", nil),
			},
			"username": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Username of BES Console Operator Profile",
				DefaultFunc: schema.EnvDefaultFunc("BFX_USERNAME", nil),
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Password of BES Console Operator Profile",
				DefaultFunc: schema.EnvDefaultFunc("BFX_PASSWORD", nil),
			},
			"server": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Bigfix server address",
				DefaultFunc: schema.EnvDefaultFunc("BFX_SERVER", nil),
			},
		},

		//Supported Data Sources by this provider
		DataSourcesMap: map[string]*schema.Resource{
			"bigfix_computer": dataSourceComputer(),
			"bigfix_fixlet":   dataSourceFixlet(),
		},

		//Supported Resources by this provider
		ResourcesMap: map[string]*schema.Resource{
			"bigfix_multiple_action_group": resourceMultiActionGroup(),
			"bigfix_fixlet":                resourceFixlet(),
			"bigfix_task":                  resourceTask(),
			"bigfix_upload_file":           resourceUploadFile(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config, err := BFXConfig(d)
	log.Println("[DEBUG] mConfig on provider is ")
	if err != nil {
		log.Println("[ERROR] Failed to establish connection with BigFix server.")
		os.Exit(1)
		return nil, err

	}
	log.Println("[DEBUG] connecting to BigFix Server...")
	return config, nil
}
