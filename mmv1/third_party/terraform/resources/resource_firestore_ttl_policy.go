package google

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"google.golang.org/api/storagetransfer/v1"
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
				Description: `Foobar.`,
			},
			"timestamp_field": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `Foobar.`,
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

	transferJob := &storagetransfer.TransferJob{
		Description:        d.Get("description").(string),
		ProjectId:          project,
		Status:             d.Get("status").(string),
		Schedule:           expandTransferSchedules(d.Get("schedule").([]interface{})),
		TransferSpec:       expandTransferSpecs(d.Get("transfer_spec").([]interface{})),
		NotificationConfig: expandTransferJobNotificationConfig(d.Get("notification_config").([]interface{})),
	}

	var res *storagetransfer.TransferJob

	err = retry(func() error {
		res, err = config.NewStorageTransferClient(userAgent).TransferJobs.Create(transferJob).Do()
		return err
	})

	if err != nil {
		fmt.Printf("Error creating transfer job %v: %v", transferJob, err)
		return err
	}

	if err := d.Set("name", res.Name); err != nil {
		return fmt.Errorf("Error setting name: %s", err)
	}

	name := GetResourceNameFromSelfLink(res.Name)
	d.SetId(fmt.Sprintf("%s/%s", project, name))

	return resourceStorageTransferJobRead(d, meta)
}
