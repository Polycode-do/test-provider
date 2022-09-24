package provider

import (
	"context"
	"fmt"

	polycode "polycode-provider/client"

	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"host": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The host of the Polycode API to interact with",
				DefaultFunc: schema.EnvDefaultFunc("POLYCODE_HOST", nil),
			},
			"username": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The Polycode username to connect with",
				DefaultFunc: schema.EnvDefaultFunc("POLYCODE_USERNAME", nil),
			},
			"password": {
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
				Description: "The Polycode password to connect with",
				DefaultFunc: schema.EnvDefaultFunc("POLYCODE_PASSWORD", nil),
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"polycode_content": resourceContent(),
			"polycode_item":    resourceItem(),
			"polycode_module":  resourceModule(),
		},
		DataSourcesMap:       map[string]*schema.Resource{},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	username := d.Get("username").(string)
	password := d.Get("password").(string)

	var host *string

	hVal, ok := d.GetOk("host")
	if ok {
		tempHost := hVal.(string)
		host = &tempHost
	}

	var diags diag.Diagnostics

	if (username != "") && (password != "") {
		c, err := polycode.NewClient(host, &username, &password)
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Error,
				Summary:  "Unable to create Polycode client",
				Detail:   fmt.Sprintf("Unable to authenticate user Polycode client: %s", err.Error()),
			})

			return nil, diags
		}

		tflog.Debug(ctx, fmt.Sprintf("Authenticated client with user %s", username))

		return c, diags
	}

	c, err := polycode.NewClient(host, nil, nil)
	if err != nil {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create Polycode client",
			Detail:   fmt.Sprintf("Unable to create anonymous Polycode client: %s", err.Error()),
		})
		return nil, diags
	}

	tflog.Debug(ctx, "Authenticated anonymous client")

	return c, diags
}
