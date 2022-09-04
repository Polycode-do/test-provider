package provider

import (
	"context"
	"fmt"
	pc "polycode-provider/client"
	"polycode-provider/client/models/content"
	"time"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceContent() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceContentCreate,
		ReadContext:   resourceContentRead,
		UpdateContext: resourceContentUpdate,
		DeleteContext: resourceContentDelete,
		Schema: map[string]*schema.Schema{
			"last_update": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Last update of the resource",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The content name",
			},
			"description": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The content description",
			},
			"type": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The content type",
			},
			"reward": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The content reward",
			},
			"container": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Required:    true,
				Description: "The content component",
				Elem:        resourceContentContainer(1),
			},
		},
	}
}

func resourceContentContainer(i int) *schema.Resource {
	if i > 3 {
		return &schema.Resource{
			Schema: map[string]*schema.Schema{},
		}
	}

	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The id of the component",
			},
			"position": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The position where the component will be rendered",
			},
			"orientation": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The orientation of the container",
			},
			"markdown": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "The markdown component",
				Elem:        resourceContentDataMarkdown(),
			},
			"editor": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "The editor component",
				Elem:        resourceContentDataEditor(),
			},
			"container": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The container component",
				Elem:        resourceContentContainer(i + 1),
			},
		},
	}
}

func resourceContentDataMarkdown() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The id of the component",
			},
			"position": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The position where the component will be rendered",
			},
			"content": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The content of the markdown",
			},
		},
	}
}

func resourceContentDataEditor() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceContentCreate,
		ReadContext:   resourceContentRead,
		UpdateContext: resourceContentUpdate,
		DeleteContext: resourceContentDelete,
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The id of the component",
			},
			"position": {
				Type:        schema.TypeInt,
				Required:    true,
				Description: "The position where the component will be rendered",
			},
			"language_settings": {
				Type:        schema.TypeSet,
				Required:    true,
				Description: "List of languages for the editor",
				Elem:        resourceContentLanguage(),
			},
			"hint": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of hints id for the editor",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"validator": {
				Type:        schema.TypeSet,
				Optional:    true,
				Description: "List of validators for the editor",
				Elem:        resourceContentValidator(),
			},
		},
	}
}

func resourceContentValidator() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"id": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "The id of the validator",
			},
			"inputs": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "List of inputs for the validator",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"outputs": {
				Type:        schema.TypeList,
				Required:    true,
				Description: "List of outputs for the validator",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"is_hidden": {
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     false,
				Description: "Whether the validator is hidden",
			},
		},
	}
}

func resourceContentLanguage() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"default_code": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The default code of the language",
			},
			"language": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The language",
			},
			"version": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The version of the language",
			},
		},
	}
}

func resourceContentCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*pc.Client)

	var diags diag.Diagnostics

	co := content.Content{
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Type:        d.Get("type").(string),
		Reward:      int64(d.Get("reward").(int)),
		RootComponent: content.Component{
			Orientation: d.Get("container.0.orientation").(string),
			Type:        "container",
			Data: content.ComponentData{
				Components: serializeChildComponents(d.Get("container.0").(map[string]interface{}), ctx),
			},
		},
		Data: content.ContentData{},
	}

	res, err := c.CreateContent(co)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create content",
			Detail:   fmt.Sprintf("Error when creating content: %s", err.Error()),
		})
		return diags
	}

	d.SetId(res.Data.ID)

	tflog.Info(ctx, fmt.Sprintf("Created Content %s", d.Id()))

	return resourceContentRead(ctx, d, m)
}

func resourceContentRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*pc.Client)

	var diags diag.Diagnostics

	tflog.Debug(ctx, fmt.Sprintf("Reading Content %s", d.Id()))

	res, err := c.GetContent(d.Id())
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to get Content",
			Detail:   fmt.Sprintf("Error when getting Content: %s", err.Error()),
		})
		return diags
	}

	d.Set("name", res.Data.Name)
	d.Set("description", res.Data.Description)
	d.Set("type", res.Data.Type)
	d.Set("reward", res.Data.Reward)
	d.Set("container", deserializeChildComponents(res.Data.IntoContent().RootComponent, 0, ctx))

	return diags
}

func resourceContentUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(*pc.Client)

	co := content.Content{
		ID:          d.Id(),
		Name:        d.Get("name").(string),
		Description: d.Get("description").(string),
		Type:        d.Get("type").(string),
		Reward:      int64(d.Get("reward").(int)),
		RootComponent: content.Component{
			ID:          d.Get("container.0.id").(string),
			Orientation: d.Get("container.0.orientation").(string),
			Type:        "container",
			Data: content.ComponentData{
				Components: serializeChildComponents(d.Get("container.0").(map[string]interface{}), ctx),
			},
		},
		Data: content.ContentData{},
	}

	_, err := c.UpdateContent(co)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to update content",
			Detail:   fmt.Sprintf("Error when updating content: %s", err.Error()),
		})
		return diags
	}

	d.Set("last_update", d.Set("last_updated", time.Now().Format(time.RFC850)))

	tflog.Info(ctx, fmt.Sprintf("Updated Content %s", d.Id()))

	return resourceContentRead(ctx, d, m)
}

func resourceContentDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	c := m.(*pc.Client)

	err := c.DeleteContent(d.Id())
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to delete Content",
			Detail:   fmt.Sprintf("Error when deleting Content: %s", err.Error()),
		})
		return diags
	}

	tflog.Info(ctx, fmt.Sprintf("Deleted Content %s", d.Id()))

	return diags
}

func serializeChildComponents(rootComponent map[string]interface{}, ctx context.Context) []content.Component {
	length := 0
	for key, val := range rootComponent {
		if key == "markdown" || key == "editor" {
			length += len(val.(*schema.Set).List())
		}
		if key == "container" && val != nil {
			length += len(val.([]interface{}))
		}
	}

	tflog.Debug(ctx, fmt.Sprintf("Serializing root Component %s with %d child components", rootComponent["id"], length))

	childComponents := make([]content.Component, length)

	for key, val := range rootComponent {
		switch key {
		case "markdown":
			for _, v := range val.(*schema.Set).List() {
				markdown := v.(map[string]interface{})
				position := markdown["position"].(int)
				childComponents[position-1] = content.Component{
					ID:   markdown["id"].(string),
					Type: "markdown",
					Data: content.ComponentData{
						Markdown: markdown["content"].(string),
					},
				}
			}
		case "editor":
			for _, v := range val.(*schema.Set).List() {
				editor := v.(map[string]interface{})

				languages := make([]content.Language, 0)
				for _, language := range editor["language_settings"].(*schema.Set).List() {
					languages = append(languages, content.Language{
						DefaultCode: language.(map[string]interface{})["default_code"].(string),
						Language:    language.(map[string]interface{})["language"].(string),
						Version:     language.(map[string]interface{})["version"].(string),
					})
				}

				validators := make([]content.Validator, 0)
				for _, validator := range editor["validator"].(*schema.Set).List() {
					inputs := make([]string, len(validator.(map[string]interface{})["inputs"].([]interface{})))
					for key, val := range validator.(map[string]interface{})["inputs"].([]interface{}) {
						inputs[key] = val.(string)
					}
					outputs := make([]string, len(validator.(map[string]interface{})["outputs"].([]interface{})))
					for key, val := range validator.(map[string]interface{})["outputs"].([]interface{}) {
						outputs[key] = val.(string)
					}

					validators = append(validators, content.Validator{
						ID:       validator.(map[string]interface{})["id"].(string),
						IsHidden: validator.(map[string]interface{})["is_hidden"].(bool),
						Input: content.ValidatorInput{
							Stdin: inputs,
						},
						Output: content.ValidatorOutput{
							Stdout: outputs,
						},
					})
				}

				hints := make([]content.ItemIdentifier, 0)
				for _, item := range editor["hint"].([]interface{}) {
					hints = append(hints, content.ItemIdentifier{
						ID: item.(string),
					})
				}

				position := editor["position"].(int)
				childComponents[position-1] = content.Component{
					ID:   editor["id"].(string),
					Type: "editor",
					Data: content.ComponentData{
						EditorSettings: content.EditorSettings{
							Languages: languages,
						},
						Validators: validators,
						Items:      hints,
					},
				}
			}
		case "container":
			for _, v := range val.([]interface{}) {
				container := v.(map[string]interface{})
				position := container["position"].(int)
				childComponents[position-1] = content.Component{
					ID:   container["id"].(string),
					Type: "container",
					Data: content.ComponentData{
						Components: serializeChildComponents(container, ctx),
					},
					Orientation: container["orientation"].(string),
				}
			}
		}
	}

	return childComponents
}

func deserializeChildComponents(rootComponent content.Component, position int, ctx context.Context) []interface{} {
	container := make([]interface{}, 0)
	container = append(container, map[string]interface{}{
		"id":          rootComponent.ID,
		"orientation": rootComponent.Orientation,
		"position":    position,
	})

	markdown := make([]interface{}, 0)
	editor := make([]interface{}, 0)

	tflog.Debug(ctx, fmt.Sprintf("Deserializing root Component %s", rootComponent.ID))

	for key, childComponent := range rootComponent.Data.Components {
		switch childComponent.Type {
		case "markdown":
			markdown = append(markdown, map[string]interface{}{
				"id":       childComponent.ID,
				"content":  childComponent.Data.Markdown,
				"position": key + 1,
			})
		case "editor":
			languages := make([]interface{}, 0)
			for _, language := range childComponent.Data.EditorSettings.Languages {
				languages = append(languages, map[string]interface{}{
					"default_code": language.DefaultCode,
					"language":     language.Language,
					"version":      language.Version,
				})
			}

			validators := make([]interface{}, 0)
			for _, validator := range childComponent.Data.Validators {
				inputs := make([]interface{}, len(validator.Input.Stdin))
				for key, val := range validator.Input.Stdin {
					inputs[key] = val
				}
				outputs := make([]interface{}, len(validator.Output.Stdout))
				for key, val := range validator.Output.Stdout {
					outputs[key] = val
				}

				validators = append(validators, map[string]interface{}{
					"id":        validator.ID,
					"is_hidden": validator.IsHidden,
					"inputs":    inputs,
					"outputs":   outputs,
				})
			}

			hints := make([]interface{}, 0)
			for _, item := range childComponent.Data.Items {
				hints = append(hints, item.ID)
			}

			editor = append(editor, map[string]interface{}{
				"id":       childComponent.ID,
				"hint":     hints,
				"position": key + 1,
			})

			editor[0].(map[string]interface{})["language_settings"] = schema.NewSet(schema.HashResource(resourceContentLanguage()), languages)
			editor[0].(map[string]interface{})["validator"] = schema.NewSet(schema.HashResource(resourceContentValidator()), validators)
		case "container":
			container[0].(map[string]interface{})["container"] = deserializeChildComponents(childComponent, key+1, ctx)
		}
	}

	container[0].(map[string]interface{})["markdown"] = schema.NewSet(schema.HashResource(resourceContentDataMarkdown()), markdown)
	container[0].(map[string]interface{})["editor"] = schema.NewSet(schema.HashResource(resourceContentDataEditor()), editor)

	return container
}
