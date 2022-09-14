package provider

import (
	"context"
	"fmt"
	"time"

	pc "polycode-provider/client"
	"polycode-provider/client/models/item"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceItem() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceItemCreate,
		ReadContext:   resourceItemRead,
		UpdateContext: resourceItemUpdate,
		DeleteContext: resourceItemDelete,
		Schema: map[string]*schema.Schema{
			"last_update": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Last update of the resource",
			},
			"cost": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The item cost",
				ValidateFunc: func(i interface{}, s string) ([]string, []error) {
					if i.(int) < 0 {
						return nil, []error{fmt.Errorf("cost must be a positive integer")}
					}
					return nil, nil
				},
			},
			"hint": {
				Type:        schema.TypeList,
				Required:    true,
				MaxItems:    1,
				Description: "The hint component",
				Elem:        resourceItemDataHint(),
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceItemDataHint() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"text": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The text of the hint",
			},
		},
	}
}

func resourceItemCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*pc.Client)

	var diags diag.Diagnostics

	it := item.Item{
		Type: "hint",
		Data: item.ItemData{
			Text: d.Get("hint.0.text").(string),
		},
		Cost: int64(d.Get("cost").(int)),
	}

	createdItem, err := c.CreateItem(it)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create item",
			Detail:   fmt.Sprintf("Error when creating item: %s", err.Error()),
		})
		return diags
	}

	d.SetId(createdItem.ID)

	tflog.Info(ctx, fmt.Sprintf("Created Item %s", d.Id()))

	return resourceItemRead(ctx, d, m)
}

func resourceItemRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*pc.Client)

	var diags diag.Diagnostics

	tflog.Debug(ctx, fmt.Sprintf("Reading Item %s", d.Id()))

	item, err := c.GetItem(d.Id())
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to get Item",
			Detail:   fmt.Sprintf("Error when getting Item: %s", err.Error()),
		})
		return diags
	}

	tflog.Info(ctx, fmt.Sprintf("Read Item %+v", d.Get("hint")))

	err = d.Set("cost", item.Cost)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to set cost",
			Detail:   fmt.Sprintf("Error when setting cost: %s", err.Error()),
		})
		return diags
	}
	err = d.Set("hint.0.text", item.Data.Text)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to set hint text",
			Detail:   fmt.Sprintf("Error when setting hint text: %s", err.Error()),
		})
		return diags
	}
	err = d.Set("type", item.Type)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to set type",
			Detail:   fmt.Sprintf("Error when setting type: %s", err.Error()),
		})
		return diags
	}

	return diags
}

func resourceItemUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(*pc.Client)

	resourceItemRead(ctx, d, m)

	it := item.Item{
		ID:   d.Id(),
		Type: "hint",
		Data: item.ItemData{
			Text: d.Get("hint.0.text").(string),
		},
		Cost: int64(d.Get("cost").(int)),
	}

	_, err := c.UpdateItem(it)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to update item",
			Detail:   fmt.Sprintf("Error when updating item: %s", err.Error()),
		})
		return diags
	}

	err = d.Set("last_update", d.Set("last_updated", time.Now().Format(time.RFC850)))
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to set last update",
			Detail:   fmt.Sprintf("Error when setting last update: %s", err.Error()),
		})
		return diags
	}

	tflog.Info(ctx, fmt.Sprintf("Updated Item %s", d.Id()))

	return resourceItemRead(ctx, d, m)
}

func resourceItemDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(*pc.Client)

	err := c.DeleteItem(d.Id())
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to delete Item",
			Detail:   fmt.Sprintf("Error when deleting Item: %s", err.Error()),
		})
		return diags
	}

	tflog.Info(ctx, fmt.Sprintf("Deleted Item %s", d.Id()))

	return diags
}
