package bigfix

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
)

func TestAccResourceMultipleActionGroup_basic(t *testing.T) {
	// 	Change this title as per mentioned in xml file and also
	//	add complete id here before running test
	title := "Multiple Action Group for Testing [ 11509943 ]"
	resourceName := "bigfix_multiple_action_group.test"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccResourceMultipleActionGroupDestroy(resourceName),
		Steps: []resource.TestStep{
			{
				Config: testAccCheckResourceMultipleActionGroupConfig(),
				Check: resource.ComposeTestCheckFunc(
					testAccResourceMultipleActionGroupExists(resourceName),
					resource.TestCheckResourceAttr(
						resourceName, "title", title),
				),
			},
		},
	})
}

func testAccResourceMultipleActionGroupExists(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Not Found: " + name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Multiple Action Group not added to Actions")
		}

		actionID := rs.Primary.Attributes["id"]

		config := testAccProvider.Meta().(Config)
		url := GetActionDetailAPI(config.ServerIP, config.Port, actionID)

		response, err := config.BfxConnection(GET, url, nil)
		if err != nil {
			return err
		}

		defer response.Body.Close()
		var actionStruct Action
		data, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		}

		err = xml.Unmarshal(data, &actionStruct)
		if err != nil {
			return err
		}

		if actionStruct.MultipleActionGroup.Title == "" {
			return fmt.Errorf("No multiple action group found")
		}
		return nil
	}
}

func testAccResourceMultipleActionGroupDestroy(name string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[name]

		if !ok {
			return fmt.Errorf("Not Found: %s", name)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("Multiple Action Group not added to Actions")
		}

		actionID := rs.Primary.Attributes["id"]

		config := testAccProvider.Meta().(Config)
		url := GetActionDetailAPI(config.ServerIP, config.Port, actionID)
		response, err := config.BfxConnection(GET, url, nil)
		if err != nil {
			return nil
		}

		if response != nil {
			return fmt.Errorf("Bad : Check  %q  still exists", rs.Primary.ID)
		}

		return nil
	}
}

func testAccCheckResourceMultipleActionGroupConfig() string {
	return fmt.Sprintf(`
	
	provider "bigfix" {
		server = "%s"
		port = %s
		username = "%s"
		password = "%s"
	}

	data  "bigfix_computer" "test"{
		name = "%s"
	}

   	resource "bigfix_multiple_action_group" "test" {
		input_file_name = "%s"
		target_computer_id = "${data.bigfix_computer.test.id}"
		site_name = ["BES Support","BES Inventory and License"]
  	}
	`, os.Getenv("BFX_SERVER"), os.Getenv("BFX_PORT"), os.Getenv("BFX_USERNAME"), os.Getenv("BFX_PASSWORD"), os.Getenv("BFX_COMPUTER"),
		os.Getenv("BFX_INPUT_FILE_NAME"))
}
