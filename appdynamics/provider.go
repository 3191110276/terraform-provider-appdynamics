package appdynamics

import (
	"context"

	//"github.com/hashicorp-demoapp/hashicups-client-go"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider -
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"base_url": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
		  },
			"token": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
			},
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Sensitive:   true,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"appdynamics_apm_application": resourceAPMApplication(),
			"appdynamics_eum_application": resourceEUMApplication(),
			"appdynamics_db_collector": resourceDBCollector(),
		},
		ConfigureContextFunc: providerConfigure,
	}
}

func providerConfigure(ctx context.Context, d *schema.ResourceData) (interface{}, diag.Diagnostics) {
	base_url := d.Get("base_url").(string)
	token := d.Get("token").(string)
	username := d.Get("username").(string)
	password := d.Get("password").(string)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	if (base_url != "") && (token != "") && (username != "") && (password != "") {
		provider_data := map[string]string{
			"base_url": base_url,
			"token":token,
			"username":username,
			"password":password,
		}

		return provider_data, diags
	}

	return nil, diags
}
