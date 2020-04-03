package privatedns

import (
	"context"
	"fmt"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/privatedns/mgmt/2018-09-01/privatedns"
	"github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2018-05-01/resources"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/clients"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/tags"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/internal/timeouts"
	"github.com/terraform-providers/terraform-provider-azurerm/azurerm/utils"
)

func dataSourceArmPrivateDnsZone() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceArmPrivateDnsZoneRead,

		Timeouts: &schema.ResourceTimeout{
			Read: schema.DefaultTimeout(5 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"resource_group_name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},

			"number_of_record_sets": {
				Type:     schema.TypeInt,
				Computed: true,
			},

			"max_number_of_record_sets": {
				Type:     schema.TypeInt,
				Computed: true,
			},

			"max_number_of_virtual_network_links": {
				Type:     schema.TypeInt,
				Computed: true,
			},

			"max_number_of_virtual_network_links_with_registration": {
				Type:     schema.TypeInt,
				Computed: true,
			},

			"name_servers": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
			},

			"tags": tags.SchemaDataSource(),
		},
	}
}

func dataSourceArmPrivateDnsZoneRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*clients.Client).PrivateDns.PrivateZonesClient
	ctx, cancel := timeouts.ForRead(meta.(*clients.Client).StopContext, d)
	defer cancel()

	name := d.Get("name").(string)
	resourceGroup := d.Get("resource_group_name").(string)

	var (
		resp privatedns.PrivateZone
		err  error
	)
	if resourceGroup != "" {
		resp, err = client.Get(ctx, resourceGroup, name)
		if err != nil {
			if utils.ResponseWasNotFound(resp.Response) {
				return fmt.Errorf("Error: Private DNS Zone %q (Resource Group %q) was not found", name, resourceGroup)
			}
			return fmt.Errorf("error reading Private DNS Zone %q (Resource Group %q): %+v", name, resourceGroup, err)
		}
	} else {
		rgClient := meta.(*clients.Client).Resource.GroupsClient

		resp, resourceGroup, err = findPrivateZone(client, rgClient, ctx, name)
		if err != nil {
			return err
		}

		if resourceGroup == "" {
			return fmt.Errorf("Error: Private DNS Zone %q was not found", name)
		}
	}

	d.Set("name", name)
	d.Set("resource_group_name", resourceGroup)

	if props := resp.PrivateZoneProperties; props != nil {
		d.Set("number_of_record_sets", props.NumberOfRecordSets)
		d.Set("max_number_of_record_sets", props.MaxNumberOfRecordSets)
		d.Set("max_number_of_virtual_network_links", props.MaxNumberOfVirtualNetworkLinks)
		d.Set("max_number_of_virtual_network_links_with_registration", props.MaxNumberOfVirtualNetworkLinksWithRegistration)
	}

	return tags.FlattenAndSet(d, resp.Tags)
}

func findPrivateZone(client *privatedns.PrivateZonesClient, rgClient *resources.GroupsClient, ctx context.Context, name string) (privatedns.PrivateZone, string, error) {
	groups, err := rgClient.List(ctx, "", nil)
	if err != nil {
		return privatedns.PrivateZone{}, "", fmt.Errorf("Error listing Resource Groups: %+v", err)
	}

	for _, g := range groups.Values() {
		resourceGroup := *g.Name

		privateZones, err := client.ListByResourceGroup(ctx, resourceGroup, nil)
		if err != nil {
			//return privatedns.PrivateZone{}, "", fmt.Errorf("Error listing Private DNS Zones (Resource Group: %s): %+v", resourceGroup, err)
			continue
		}

		for _, z := range privateZones.Values() {
			if *z.Name == name {
				return z, resourceGroup, nil
			}
		}
	}

	return privatedns.PrivateZone{}, "", nil
}
