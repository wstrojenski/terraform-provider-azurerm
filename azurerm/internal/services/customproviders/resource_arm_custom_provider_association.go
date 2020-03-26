package customproviders

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/services/preview/customproviders/mgmt/2018-09-01-preview/customproviders"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/tf"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/clients"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/features"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/timeouts"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/helpers/azure"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/customproviders/parse"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/services/customproviders/validate"
	azSchema "github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tf/schema"
)

func resourceArmCustomProviderAssociation() *schema.Resource {
	return &schema.Resource{
		Create: resourceArmCustomProviderAssociationCreateUpdate,
		Read:   resourceArmCustomProviderAssociationRead,
		Update: resourceArmCustomProviderAssociationCreateUpdate,
		Delete: resourceArmCustomProviderAssociationDelete,
		Importer: azSchema.ValidateResourceIDPriorToImport(func(id string) error {
			_, err := parse.AssociationID(id)
			return err
		}),

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(30 * time.Minute),
			Read:   schema.DefaultTimeout(5 * time.Minute),
			Update: schema.DefaultTimeout(30 * time.Minute),
			Delete: schema.DefaultTimeout(30 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.CustomProviderName,
			},

			"scope": {
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: azure.ValidateResourceID,
			},

			"target_resource_id": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: azure.ValidateResourceID,
			},
		},
	}
}

func resourceArmCustomProviderAssociationCreateUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).CustomProviders.AssociationClient
	ctx, cancel := timeouts.ForCreateUpdate(meta.(*clients.Client).StopContext, d)
	defer cancel()

	name := d.Get("name").(string)
	scope := d.Get("scope").(string)

	if features.ShouldResourcesBeImported() && d.IsNewResource() {
		existing, err := client.Get(ctx, scope, name)
		if err != nil {
			if !utils.ResponseWasNotFound(existing.Response) {
				return fmt.Errorf("checking for presence of existing Custom Provider Association %q (Scope %q): %s", name, scope, err)
			}
		}

		if existing.ID != nil && *existing.ID != "" {
			return tf.ImportAsExistsError("azurerm_custom_provider_association", *existing.ID)
		}
	}

	association := customproviders.Association{
		AssociationProperties: &customproviders.AssociationProperties{
			TargetResourceID: utils.String(d.Get("target_resource_id").(string)),
		},
	}

	scope = strings.TrimPrefix(scope, "/")

	future, err := client.CreateOrUpdate(ctx, scope, name, association)
	if err != nil {
		return fmt.Errorf("creating/updating Custom Provider Association %q (Scope %q): %+v", name, scope, err)
	}

	if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
		return fmt.Errorf("waiting for completion of Custom Provider Association %q (Scope %q): %+v", name, scope, err)
	}

	resp, err := client.Get(ctx, scope, name)
	if err != nil {
		return fmt.Errorf("retrieving Custom Provider Association %q (Scope %q): %+v", name, scope, err)
	}

	if resp.ID == nil || *resp.ID == "" {
		return fmt.Errorf("cannot read Custom Provider Association %q (Scope %q) ID", name, scope)
	}

	d.SetId(*resp.ID)

	return resourceArmCustomProviderAssociationRead(d, meta)
}

func resourceArmCustomProviderAssociationRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).CustomProviders.AssociationClient
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.AssociationID(d.Id())
	if err != nil {
		return err
	}

	resp, err := client.Get(ctx, id.ResourceId, id.Name)
	if err != nil {
		if utils.ResponseWasNotFound(resp.Response) {
			d.SetId("")
			return nil
		}

		return fmt.Errorf("retrieving Custom Provider Association %q (Scope %q): %+v", id.Name, id.ResourceId, err)
	}

	d.Set("name", resp.Name)
	d.Set("scope", id.ResourceId)

	return nil
}

func resourceArmCustomProviderAssociationDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).CustomProviders.AssociationClient
	ctx, cancel := timeouts.ForDelete(meta.(*clients.Client).StopContext, d)
	defer cancel()

	id, err := parse.AssociationID(d.Id())
	if err != nil {
		return err
	}

	future, err := client.Delete(ctx, id.ResourceId, id.Name)
	if err != nil {
		return fmt.Errorf("deleting Custom Provider Association %q (Scope %q): %+v", id.Name, id.ResourceId, err)
	}

	if err = future.WaitForCompletionRef(ctx, client.Client); err != nil {
		return fmt.Errorf("waiting for deletion of Custom Provider Association %q (Scope %q): %+v", id.Name, id.ResourceId, err)
	}

	return nil
}
