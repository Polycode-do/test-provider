package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type User struct {
	ID          types.String `tfsdk:"id"`
	Username    types.String `tfsdk:"username"`
	Description types.String `tfsdk:"description"`
	Points      types.Int64  `tfsdk:"points"`
	Rank        types.Int64  `tfsdk:"rank"`
}
