package bigfix

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"bigfix": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func TestProvider_impl(t *testing.T) {
	var _ terraform.ResourceProvider = Provider()
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("BFX_SERVER"); v == "" {
		t.Fatal("BFX_SERVER must be set for acceptance test")
	}

	if v := os.Getenv("BFX_PORT"); v == "" {
		t.Fatal("BFX_PORT must be set for acceptance test")
	}

	if v := os.Getenv("BFX_USERNAME"); v == "" {
		t.Fatal("BFX_USERNAME must be set for acceptance test")
	}

	if v := os.Getenv("BFX_PASSWORD"); v == "" {
		t.Fatal("BFX_PASSWORD must be set for acceptance test")
	}
}
