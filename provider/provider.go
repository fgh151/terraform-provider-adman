package provider

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"login": {
				Type:     schema.TypeString,
				Required: true,
			},
			"mdpass": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"adman_dns_zone":        zone(),
			"adman_dns_zone_record": record(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	login := d.Get("login").(string)
	mdpass := d.Get("mdpass").(string)

	var diags diag.Diagnostics

	c := AdmanProvider{}

	if login == "" || mdpass == "" {

		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to create client ",
			Detail:   "Login or mdpass is blank",
		})
	}

	c = AdmanProvider{
		Login:  login,
		Mdpass: mdpass,
	}

	return c, diags
}
