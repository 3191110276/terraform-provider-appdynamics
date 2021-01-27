package appdynamics

import (
	"context"
	"fmt"
	"strings"
	"net/http"
	"io/ioutil"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceApplication() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceApplicationCreate,
		ReadContext:   resourceApplicationRead,
		UpdateContext: resourceApplicationUpdate,
		DeleteContext: resourceApplicationDelete,
		Schema: map[string]*schema.Schema{
			"last_updated": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: false,
			},
		},
	}
}

func resourceApplicationCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	url := d.Get("base_url").(string) + "/controller/restui/allApplications/createApplication?applicationType=APM%0A"
	bearer := "Bearer " + d.Get("token").(string)

	payload := strings.NewReader("{\"name\": \"apitest\", \"description\": \"\"}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//items := d.Get("items").([]interface{})

	resourceApplicationRead(ctx, d, m)

	return diags
}

func resourceApplicationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	return diags
}

func resourceApplicationUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	//orderID := d.Id()

	return resourceApplicationRead(ctx, d, m)
}

func resourceApplicationDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//orderID := d.Id()

	// d.SetId("") is automatically called assuming delete returns no errors, but
	// it is added here for explicitness.
	d.SetId("")

	return diags
}
