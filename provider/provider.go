package provider

import (
	"context"
	"os"

	"polycode-provider/client"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func New() tfsdk.Provider {
	return &Provider{}
}

type Provider struct {
	configured bool
	client     *client.Client
}

func (provider *Provider) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"host": {
				Type:        types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "The host of the Polycode server.",
			},
			"password": {
				Type:        types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "The password of the admin that will be used to make requests to the API.",
			},
			"username": {
				Type:        types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "The username of the admin that will be used to make requests to the API.",
			},
		},
	}, nil
}

type ProviderData struct {
	Username types.String `tfsdk:"username"`
	Host     types.String `tfsdk:"host"`
	Password types.String `tfsdk:"password"`
}

func (provider *Provider) Configure(ctx context.Context, req tfsdk.ConfigureProviderRequest, res *tfsdk.ConfigureProviderResponse) {
	var config ProviderData
	diags := req.Config.Get(ctx, &config)
	res.Diagnostics.Append(diags...)
	if res.Diagnostics.HasError() {
		return
	}

	var username string
	if config.Username.Unknown {
		res.Diagnostics.AddWarning(
			"Unable to create client",
			"Cannot use unknown value as username",
		)
		return
	}

	if config.Username.Null {
		username = os.Getenv("POLYCODE_USERNAME")
	} else {
		username = config.Username.Value
	}

	if username == "" {
		res.Diagnostics.AddError(
			"Unable to find username",
			"Username cannot be an empty string",
		)
		return
	}

	var password string
	if config.Password.Unknown {
		res.Diagnostics.AddError(
			"Unable to create client",
			"Cannot use unknown value as password",
		)
		return
	}

	if config.Password.Null {
		password = os.Getenv("POLYCODE_PASSWORD")
	} else {
		password = config.Password.Value
	}

	if password == "" {
		res.Diagnostics.AddError(
			"Unable to find password",
			"password cannot be an empty string",
		)
		return
	}

	var host string
	if config.Host.Unknown {
		res.Diagnostics.AddError(
			"Unable to create client",
			"Cannot use unknown value as host",
		)
		return
	}

	if config.Host.Null {
		host = os.Getenv("POLYCODE_HOST")
	} else {
		host = config.Host.Value
	}

	if host == "" {
		res.Diagnostics.AddError(
			"Unable to find host",
			"Host cannot be an empty string",
		)
		return
	}

	c, err := client.NewClient(&host, &username, &password)
	if err != nil {
		res.Diagnostics.AddError(
			"Unable to create client",
			"Unable to create polycode client:\n\n"+err.Error(),
		)
		return
	}

	provider.client = c
	provider.configured = true
}

func (provider *Provider) GetResources(_ context.Context) (map[string]tfsdk.ResourceType, diag.Diagnostics) {
	return map[string]tfsdk.ResourceType{}, nil
}

func (provider *Provider) GetDataSources(_ context.Context) (map[string]tfsdk.DataSourceType, diag.Diagnostics) {
	return map[string]tfsdk.DataSourceType{
		"polycode_user": DataSourceUserType{},
	}, nil
}
