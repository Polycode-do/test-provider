package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type DataSourceUserType struct{}

func (dataSource DataSourceUserType) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Type:        types.StringType,
				Required:    true,
				Description: "The ID of the user.",
			},
			"username": {
				Type:        types.StringType,
				Computed:    true,
				Optional:    true,
				Description: "The username of the user.",
			},
			"description": {
				Type:        types.StringType,
				Computed:    true,
				Optional:    true,
				Description: "The description of the user.",
			},
			"points": {
				Type:        types.Int64Type,
				Computed:    true,
				Optional:    true,
				Description: "The points of the user.",
			},
			"rank": {
				Type:        types.Int64Type,
				Computed:    true,
				Optional:    true,
				Description: "The rank of the user.",
			},
		},
	}, nil
}

func (dataSource DataSourceUserType) NewDataSource(_ context.Context, provider tfsdk.Provider) (tfsdk.DataSource, diag.Diagnostics) {
	return DataSourceUser{
		provider: *(provider.(*Provider)),
	}, nil
}

type DataSourceUser struct {
	provider Provider
}

func (dataSource DataSourceUser) Read(ctx context.Context, req tfsdk.ReadDataSourceRequest, res *tfsdk.ReadDataSourceResponse) {
	var state User
	diags := req.Config.Get(ctx, &state)
	res.Diagnostics.Append(diags...)
	if res.Diagnostics.HasError() {
		return
	}

	userID := state.ID.Value

	user, err := dataSource.provider.client.GetUser(userID)
	if err != nil {
		res.Diagnostics.AddError(
			"Error reading user",
			"Could not read userID"+userID+": "+err.Error(),
		)
		return
	}

	state.ID = types.String{Value: user.Data.ID}
	state.Username = types.String{Value: user.Data.Username}
	state.Description = types.String{Value: user.Data.Description}
	state.Points = types.Int64{Value: user.Data.Points}
	state.Rank = types.Int64{Value: user.Data.Rank}

	diags = res.State.Set(ctx, &state)
	res.Diagnostics.Append(diags...)
	if res.Diagnostics.HasError() {
		return
	}
}
