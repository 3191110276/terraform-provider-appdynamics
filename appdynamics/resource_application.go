package appdynamics

import (
	"context"
	"fmt"
	"strings"
	"time"
	"net/http"
	"io/ioutil"
	"encoding/json"

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
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"debuga": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"debugb": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceApplicationCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

  provider_data := m.(map[string]string)
  base_url := provider_data["base_url"]
	token := provider_data["token"]

	url := base_url + "/controller/restui/allApplications/createApplication?applicationType=APM"
	bearer := "Bearer " + token

	//type req_template struct {
  //  name        string
  //  description string
  //}

	//req_go := &req_template{
	//		name:   "tftesting1234",
	//		description: "abcd",
	//}

	//req_json, _ := json.Marshal(req_go)
	//d.Set("debuga", string(req_json))
	//payload := strings.NewReader(string(req_json))

	req_string := "{\"name\": \"APPNAME\", \"description\": \"DESCRIPTION\"}"
  //d.Set("debuga", strings.Replace(req_string, "APPNAME", d.Get("name").(string), 1))
	//d.Set("debugb", strings.Replace(req_string, "DESCRIPTION", d.Get("description").(string), 1))

	payload := strings.NewReader(req_string)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	//body, _ := ioutil.ReadAll(res.Body)

	d.Set("version", "1")

	resourceApplicationRead(ctx, d, m)

	return diags
}

func resourceApplicationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	provider_data := m.(map[string]string)
  base_url := provider_data["base_url"]
	token := provider_data["token"]

  url := base_url + "/controller/rest/applications?output=json"
	bearer := "Bearer " + token

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", bearer)
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	type Entries []struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		ID          int    `json:"id"`
		AccountGUID string `json:"accountGuid"`
	}

	data := Entries{}
	_ = json.Unmarshal([]byte(body), &data)

	for i := 0; i < len(data); i++ {
		if (data[i].Name == d.Get("name").(string)) {
			d.SetId(fmt.Sprint(data[i].ID))
		}
	}

	return diags
}

func resourceApplicationUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	//appID := d.Id()

	provider_data := m.(map[string]string)
  base_url := provider_data["base_url"]
	token := provider_data["token"]

  url := base_url + "/controller/restui/allApplications/updateApplicationDetails"
	bearer := "Bearer " + token

  payload := strings.NewReader("{\n\t\"id\":7558,\n\t\"version\":6,\n\t\"name\":\"apitest7\",\n\t\"description\":\"\",\n\t\"active\":true,\n\t\"running\":false,\n\t\"eumAppName\":null\n}")

	req, _ := http.NewRequest("POST", url, payload)

  req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

	d.Set("last_updated", time.Now().Format(time.RFC850))

	return resourceApplicationRead(ctx, d, m)
}

func resourceApplicationDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	provider_data := m.(map[string]string)
  base_url := provider_data["base_url"]
	token := provider_data["token"]

  url := base_url + "/controller/restui/allApplications/deleteApplication"
	bearer := "Bearer " + token

  payload := strings.NewReader(d.Id())

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

	d.SetId("")

	return diags
}
