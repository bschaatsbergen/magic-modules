package google

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// resourceFirestoreTtlPolicy returns a *schema.Resource that allows a customer
// to declare a Firestore TTL Policy resource.
func resourceFirestoreTtlPolicy() *schema.Resource {
	return &schema.Resource{
		SchemaVersion: 1,

		Create: resourceFirestoreTtlPolicyCreate,
		Read:   resourceFirestoreTtlPolicyRead,
		Update: resourceFirestoreTtlPolicyUpdate,
		Delete: resourceFirestoreTtlPolicyDelete,

		Schema: map[string]*schema.Schema{
			"project_id": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validateProjectID(),
				Description:  `The project ID. Changing this forces a new project to be created.`,
			},
			"collection_group": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Collection group that the time-to-live policy should be applied on.`,
			},
			"timestamp_field": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Field that represents the time-to-live policy expiration time for documents in a given collection group.`,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceFirestoreTtlPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	project, err := getProject(d, config)
	if err != nil {
		return err
	}
}

func resourceFirestoreTtlPolicyRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceFirestoreTtlPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceFirestoreTtlPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
