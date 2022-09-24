package provider

import (
	"context"
	"fmt"
	"time"

	pc "polycode-provider/client"
	"polycode-provider/client/models/module"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceModule() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceModuleCreate,
		ReadContext:   resourceModuleRead,
		UpdateContext: resourceModuleUpdate,
		DeleteContext: resourceModuleDelete,
		Schema: map[string]*schema.Schema{
			"last_update": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Last update of the resource",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the module",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(string)
					if len(v) < 3 {
						errs = append(errs, fmt.Errorf("%q must be at least 3 characters", key))
					}
					return
				},
			},
			"description": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Description of the module",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(string)
					if len(v) < 3 {
						errs = append(errs, fmt.Errorf("%q must be at least 3 characters", key))
					}
					return
				},
			},
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Type of the module",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(string)
					if v != "challenge" && v != "practice" && v != "certification" && v != "submodule" {
						errs = append(errs, fmt.Errorf("%q must be hint", key))
					}
					return
				},
			},
			"tags": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "Tags of the module",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"reward": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "Reward of the module",
				ValidateFunc: func(val interface{}, key string) (warns []string, errs []error) {
					v := val.(int)
					if v < 0 {
						errs = append(errs, fmt.Errorf("%q must be positive", key))
					}
					return
				},
			},
			"module": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "List of modules id",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"content": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "List of content id",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
		},
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
	}
}

func resourceModuleCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*pc.Client)

	var diags diag.Diagnostics

	modules := make([]module.ModuleIdentifier, 0)
	for _, v := range d.Get("module").([]interface{}) {
		modules = append(modules, module.ModuleIdentifier{
			ID: v.(string),
		})
	}
	contents := make([]module.ContentIdentifier, 0)
	for _, v := range d.Get("content").([]interface{}) {
		contents = append(contents, module.ContentIdentifier{
			ID: v.(string),
		})
	}
	tags := make([]string, 0)
	for _, v := range d.Get("tags").([]interface{}) {
		tags = append(tags, v.(string))
	}

	mo := module.Module{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Type:        d.Get("type").(string),
		Tags:        tags,
		Reward:      int64(d.Get("reward").(int)),
		Modules:     modules,
		Contents:    contents,
	}

	createdModule, err := c.CreateModule(mo)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create module",
			Detail:   fmt.Sprintf("Error when creating module: %s", err.Error()),
		})
		return diags
	}

	d.SetId(createdModule.ID)

	tflog.Info(ctx, fmt.Sprintf("Created Module %s", d.Id()))

	return resourceModuleRead(ctx, d, m)
}

func resourceModuleRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*pc.Client)

	var diags diag.Diagnostics

	tflog.Debug(ctx, fmt.Sprintf("Reading Module %s", d.Id()))

	module, err := c.GetModule(d.Id())
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to get Module",
			Detail:   fmt.Sprintf("Error when getting Module: %s", err.Error()),
		})
		return diags
	}

	modules := make([]string, 0)
	for _, v := range module.Modules {
		modules = append(modules, v.ID)
	}
	contents := make([]string, 0)
	for _, v := range module.Contents {
		contents = append(contents, v.ID)
	}

	err = d.Set("name", module.Name)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to set name",
			Detail:   fmt.Sprintf("Error when setting name: %s", err.Error()),
		})
	}
	err = d.Set("description", module.Description)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to set description",
			Detail:   fmt.Sprintf("Error when setting description: %s", err.Error()),
		})
	}
	err = d.Set("type", module.Type)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to set type",
			Detail:   fmt.Sprintf("Error when setting type: %s", err.Error()),
		})
	}
	err = d.Set("tags", module.Tags)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to set tags",
			Detail:   fmt.Sprintf("Error when setting tags: %s", err.Error()),
		})
	}
	err = d.Set("reward", module.Reward)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to set reward",
			Detail:   fmt.Sprintf("Error when setting reward: %s", err.Error()),
		})
	}
	err = d.Set("module", modules)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to set module",
			Detail:   fmt.Sprintf("Error when setting module: %s", err.Error()),
		})
	}
	err = d.Set("content", contents)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to set content",
			Detail:   fmt.Sprintf("Error when setting content: %s", err.Error()),
		})
	}

	return diags
}

func resourceModuleUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(*pc.Client)

	modules := make([]module.ModuleIdentifier, 0)
	for _, v := range d.Get("module").([]interface{}) {
		modules = append(modules, module.ModuleIdentifier{
			ID: v.(string),
		})
	}
	contents := make([]module.ContentIdentifier, 0)
	for _, v := range d.Get("content").([]interface{}) {
		contents = append(contents, module.ContentIdentifier{
			ID: v.(string),
		})
	}
	tags := make([]string, 0)
	for _, v := range d.Get("tags").([]interface{}) {
		tags = append(tags, v.(string))
	}

	mo := module.Module{
		ID:          d.Id(),
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Type:        d.Get("type").(string),
		Tags:        tags,
		Reward:      int64(d.Get("reward").(int)),
		Modules:     modules,
		Contents:    contents,
	}

	_, err := c.UpdateModule(mo)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to update module",
			Detail:   fmt.Sprintf("Error when updating module: %s", err.Error()),
		})
		return diags
	}

	err = d.Set("last_update", time.Now().Format(time.RFC850))
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to set last update",
			Detail:   fmt.Sprintf("Error when setting last update: %s", err.Error()),
		})
		return diags
	}

	tflog.Info(ctx, fmt.Sprintf("Updated Module %s", d.Id()))

	return resourceModuleRead(ctx, d, m)
}

func resourceModuleDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(*pc.Client)

	err := c.DeleteModule(d.Id())
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to delete Module",
			Detail:   fmt.Sprintf("Error when deleting Module: %s", err.Error()),
		})
		return diags
	}

	tflog.Info(ctx, fmt.Sprintf("Deleted Module %s", d.Id()))

	return diags
}
