package bigfix

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccDataSourceFixlet_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDataFixletConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.bigfix_fixlet.test", "name", os.Getenv("BFX_FIXLET")),
				),
			},
		},
	})
}

func testAccDataSourceFixlet(src, n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		service := s.RootModule().Resources[src]
		serviceResource := service.Primary.Attributes

		search := s.RootModule().Resources[n]
		searchResource := search.Primary.Attributes

		testArrtributes := []string{
			"id",
			"site_name",
			"default_action",
			"name",
		}

		for _, attribute := range testArrtributes {
			if searchResource[attribute] != serviceResource[attribute] {
				return fmt.Errorf("Expected Computer parameter `%s` to be `%s` but got `%s`", attribute, serviceResource[attribute], searchResource[attribute])
			}
		}
		return nil
	}
}

func testAccCheckDataFixletConfig() string {
	return fmt.Sprintf(`
	provider "bigfix" {
		server = "%s"
		port = %s
		username = "%s"
		password = "%s"
	  }
	  
	  data  "bigfix_fixlet" "test"{
		name = "%s"
   }
	`, os.Getenv("BFX_SERVER"),
		os.Getenv("BFX_PORT"),
		os.Getenv("BFX_USERNAME"),
		os.Getenv("BFX_PASSWORD"),
		os.Getenv("BFX_FIXLET"))
}
