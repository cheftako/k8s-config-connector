// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccDataprocMetastoreFederation_dataprocMetastoreFederationBasicExample(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"random_suffix": randString(t, 10),
	}

	vcrTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProvidersOiCS,
		CheckDestroy: testAccCheckDataprocMetastoreFederationDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataprocMetastoreFederation_dataprocMetastoreFederationBasicExample(context),
			},
			{
				ResourceName:            "google_dataproc_metastore_federation.default",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"federation_id", "location"},
			},
		},
	})
}

func testAccDataprocMetastoreFederation_dataprocMetastoreFederationBasicExample(context map[string]interface{}) string {
	return Nprintf(`
resource "google_dataproc_metastore_federation" "default" {
  provider      = google-beta
  location      = "us-central1"
  federation_id = "tf-test-fed-1%{random_suffix}"
  version       = "3.1.2"

  backend_metastores {
      rank           = "1"
      name           = google_dataproc_metastore_service.default.id
      metastore_type = "DATAPROC_METASTORE" 
    }
}

resource "google_dataproc_metastore_service" "default" {
  provider   = google-beta
  service_id = "tf-test-fed-1%{random_suffix}"
  location   = "us-central1"
  tier       = "DEVELOPER"


  hive_metastore_config {
    version           = "3.1.2"
    endpoint_protocol = "GRPC"
  }
}
`, context)
}

func testAccCheckDataprocMetastoreFederationDestroyProducer(t *testing.T) func(s *terraform.State) error {
	return func(s *terraform.State) error {
		for name, rs := range s.RootModule().Resources {
			if rs.Type != "google_dataproc_metastore_federation" {
				continue
			}
			if strings.HasPrefix(name, "data.") {
				continue
			}

			config := googleProviderConfig(t)

			url, err := replaceVarsForTest(config, rs, "{{DataprocMetastoreBasePath}}projects/{{project}}/locations/{{location}}/federations/{{federation_id}}")
			if err != nil {
				return err
			}

			billingProject := ""

			if config.BillingProject != "" {
				billingProject = config.BillingProject
			}

			_, err = sendRequest(config, "GET", billingProject, url, config.userAgent, nil)
			if err == nil {
				return fmt.Errorf("DataprocMetastoreFederation still exists at %s", url)
			}
		}

		return nil
	}
}
