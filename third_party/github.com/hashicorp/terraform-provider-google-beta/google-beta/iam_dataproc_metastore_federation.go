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

	"github.com/hashicorp/errwrap"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/cloudresourcemanager/v1"
)

var DataprocMetastoreFederationIamSchema = map[string]*schema.Schema{
	"project": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"location": {
		Type:     schema.TypeString,
		Computed: true,
		Optional: true,
		ForceNew: true,
	},
	"federation_id": {
		Type:             schema.TypeString,
		Required:         true,
		ForceNew:         true,
		DiffSuppressFunc: compareSelfLinkOrResourceName,
	},
}

type DataprocMetastoreFederationIamUpdater struct {
	project      string
	location     string
	federationId string
	d            TerraformResourceData
	Config       *Config
}

func DataprocMetastoreFederationIamUpdaterProducer(d TerraformResourceData, config *Config) (ResourceIamUpdater, error) {
	values := make(map[string]string)

	project, _ := getProject(d, config)
	if project != "" {
		if err := d.Set("project", project); err != nil {
			return nil, fmt.Errorf("Error setting project: %s", err)
		}
	}
	values["project"] = project
	location, _ := getLocation(d, config)
	if location != "" {
		if err := d.Set("location", location); err != nil {
			return nil, fmt.Errorf("Error setting location: %s", err)
		}
	}
	values["location"] = location
	if v, ok := d.GetOk("federation_id"); ok {
		values["federation_id"] = v.(string)
	}

	// We may have gotten either a long or short name, so attempt to parse long name if possible
	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/federations/(?P<federation_id>[^/]+)", "(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<federation_id>[^/]+)", "(?P<location>[^/]+)/(?P<federation_id>[^/]+)", "(?P<federation_id>[^/]+)"}, d, config, d.Get("federation_id").(string))
	if err != nil {
		return nil, err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &DataprocMetastoreFederationIamUpdater{
		project:      values["project"],
		location:     values["location"],
		federationId: values["federation_id"],
		d:            d,
		Config:       config,
	}

	if err := d.Set("project", u.project); err != nil {
		return nil, fmt.Errorf("Error setting project: %s", err)
	}
	if err := d.Set("location", u.location); err != nil {
		return nil, fmt.Errorf("Error setting location: %s", err)
	}
	if err := d.Set("federation_id", u.GetResourceId()); err != nil {
		return nil, fmt.Errorf("Error setting federation_id: %s", err)
	}

	return u, nil
}

func DataprocMetastoreFederationIdParseFunc(d *schema.ResourceData, config *Config) error {
	values := make(map[string]string)

	project, _ := getProject(d, config)
	if project != "" {
		values["project"] = project
	}

	location, _ := getLocation(d, config)
	if location != "" {
		values["location"] = location
	}

	m, err := getImportIdQualifiers([]string{"projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/federations/(?P<federation_id>[^/]+)", "(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<federation_id>[^/]+)", "(?P<location>[^/]+)/(?P<federation_id>[^/]+)", "(?P<federation_id>[^/]+)"}, d, config, d.Id())
	if err != nil {
		return err
	}

	for k, v := range m {
		values[k] = v
	}

	u := &DataprocMetastoreFederationIamUpdater{
		project:      values["project"],
		location:     values["location"],
		federationId: values["federation_id"],
		d:            d,
		Config:       config,
	}
	if err := d.Set("federation_id", u.GetResourceId()); err != nil {
		return fmt.Errorf("Error setting federation_id: %s", err)
	}
	d.SetId(u.GetResourceId())
	return nil
}

func (u *DataprocMetastoreFederationIamUpdater) GetResourceIamPolicy() (*cloudresourcemanager.Policy, error) {
	url, err := u.qualifyFederationUrl("getIamPolicy")
	if err != nil {
		return nil, err
	}

	project, err := getProject(u.d, u.Config)
	if err != nil {
		return nil, err
	}
	var obj map[string]interface{}

	userAgent, err := generateUserAgentString(u.d, u.Config.userAgent)
	if err != nil {
		return nil, err
	}

	policy, err := sendRequest(u.Config, "GET", project, url, userAgent, obj)
	if err != nil {
		return nil, errwrap.Wrapf(fmt.Sprintf("Error retrieving IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	out := &cloudresourcemanager.Policy{}
	err = Convert(policy, out)
	if err != nil {
		return nil, errwrap.Wrapf("Cannot convert a policy to a resource manager policy: {{err}}", err)
	}

	return out, nil
}

func (u *DataprocMetastoreFederationIamUpdater) SetResourceIamPolicy(policy *cloudresourcemanager.Policy) error {
	json, err := ConvertToMap(policy)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	obj["policy"] = json

	url, err := u.qualifyFederationUrl("setIamPolicy")
	if err != nil {
		return err
	}
	project, err := getProject(u.d, u.Config)
	if err != nil {
		return err
	}

	userAgent, err := generateUserAgentString(u.d, u.Config.userAgent)
	if err != nil {
		return err
	}

	_, err = sendRequestWithTimeout(u.Config, "POST", project, url, userAgent, obj, u.d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return errwrap.Wrapf(fmt.Sprintf("Error setting IAM policy for %s: {{err}}", u.DescribeResource()), err)
	}

	return nil
}

func (u *DataprocMetastoreFederationIamUpdater) qualifyFederationUrl(methodIdentifier string) (string, error) {
	urlTemplate := fmt.Sprintf("{{DataprocMetastoreBasePath}}%s:%s", fmt.Sprintf("projects/%s/locations/%s/federations/%s", u.project, u.location, u.federationId), methodIdentifier)
	url, err := replaceVars(u.d, u.Config, urlTemplate)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (u *DataprocMetastoreFederationIamUpdater) GetResourceId() string {
	return fmt.Sprintf("projects/%s/locations/%s/federations/%s", u.project, u.location, u.federationId)
}

func (u *DataprocMetastoreFederationIamUpdater) GetMutexKey() string {
	return fmt.Sprintf("iam-dataprocmetastore-federation-%s", u.GetResourceId())
}

func (u *DataprocMetastoreFederationIamUpdater) DescribeResource() string {
	return fmt.Sprintf("dataprocmetastore federation %q", u.GetResourceId())
}
