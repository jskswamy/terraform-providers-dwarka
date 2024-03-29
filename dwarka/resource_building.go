package dwarka

import (
	"context"
	"terraform-provider-dwarka/client/dwarka"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceBuilding() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceBuildingCreate,
		ReadContext:   resourceBuildingRead,
		UpdateContext: resourceBuildingUpdate,
		DeleteContext: resourceBuildingDelete,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Required: true,
			},
			"lat": {
				Type:     schema.TypeFloat,
				Required: true,
			},
			"lan": {
				Type:     schema.TypeFloat,
				Required: true,
			},
		},
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
	}
}

func resourceBuildingCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*dwarka.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	building := dwarka.Building{
		Name:        d.Get("name").(string),
		Lat:         d.Get("lat").(float64),
		Lan:         d.Get("lan").(float64),
		Description: d.Get("description").(string),
	}

	b, err := c.CreateBuilding(building)
	if err != nil {
		return diag.FromErr(err)
	}

	d.SetId(b.Name)

	resourceBuildingRead(ctx, d, m)

	return diags
}

func resourceBuildingRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*dwarka.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	orderID := d.Id()

	building, err := c.GetBuilding(orderID)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("name", building.Name); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("lat", building.Lat); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("lan", building.Lan); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("description", building.Description); err != nil {
		return diag.FromErr(err)
	}

	return diags
}

func resourceBuildingUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*dwarka.Client)

	buildingID := d.Id()

	if d.HasChanges("name", "lat", "lan", "description") {
		building := dwarka.Building{
			Name:        d.Get("name").(string),
			Lat:         d.Get("lat").(float64),
			Lan:         d.Get("lan").(float64),
			Description: d.Get("description").(string),
		}

		_, err := c.UpdateBuilding(buildingID, building)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	return resourceBuildingRead(ctx, d, m)
}

func resourceBuildingDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*dwarka.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	buildingID := d.Id()

	err := c.DeleteBuilding(buildingID)
	if err != nil {
		return diag.FromErr(err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
