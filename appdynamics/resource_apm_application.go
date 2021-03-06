package appdynamics

import (
	"context"
	"fmt"
	"strings"
	"time"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceAPMApplication() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceAPMApplicationCreate,
		ReadContext:   resourceAPMApplicationRead,
		UpdateContext: resourceAPMApplicationUpdate,
		DeleteContext: resourceAPMApplicationDelete,
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
		},
	}
}

func resourceAPMApplicationCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

  provider_data := m.(map[string]string)
  base_url := provider_data["base_url"]
	token := provider_data["token"]

	url := base_url + "/controller/restui/allApplications/createApplication?applicationType=APM"
	bearer := "Bearer " + token

	req_string := "{\"name\": \"APPNAME\", \"description\": \"DESCRIPTION\"}"
  req_string = strings.Replace(req_string, "APPNAME", d.Get("name").(string), 1)
	req_string = strings.Replace(req_string, "DESCRIPTION", d.Get("description").(string), 1)

  payload := strings.NewReader(req_string)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	//body, _ := ioutil.ReadAll(res.Body)

	d.Set("version", "1")

	resourceAPMApplicationRead(ctx, d, m)

	return diags
}

func resourceAPMApplicationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

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
		if (d.Id() == "") {
			if (data[i].Name == d.Get("name").(string)) {
				d.SetId(fmt.Sprint(data[i].ID))
				d.Set("name", data[i].Name)
				d.Set("description", data[i].Description)
			}
		} else {
			if (fmt.Sprint(data[i].ID) == d.Id()) {
				d.Set("name", data[i].Name)
				d.Set("description", data[i].Description)
			}
		}
	}

	return diags
}

func resourceAPMApplicationUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	provider_data := m.(map[string]string)
  base_url := provider_data["base_url"]
	token := provider_data["token"]

  url := base_url + "/controller/restui/allApplications/updateApplicationDetails"
	bearer := "Bearer " + token

  req_string := "{\n\t\"id\":APPID,\n\t\"version\":APPVERSION,\n\t\"name\":\"APPNAME\",\"description\":\"DESCRIPTION\"\n\t\n}"
  req_string = strings.Replace(req_string, "APPID", d.Id(), 1)
	req_string = strings.Replace(req_string, "APPVERSION", d.Get("version").(string), 1)
	req_string = strings.Replace(req_string, "APPNAME", d.Get("name").(string), 1)
	req_string = strings.Replace(req_string, "DESCRIPTION", d.Get("description").(string), 1)

	payload := strings.NewReader(req_string)

	req, _ := http.NewRequest("POST", url, payload)

  req.Header.Add("Authorization", bearer)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	//body, _ := ioutil.ReadAll(res.Body)

	current_version, err := strconv.ParseInt(d.Get("version").(string), 10, 64)
	if err != nil {
		return diag.FromErr(err)
	}
	new_version := current_version + 1
	new_version_string := fmt.Sprint(new_version)
	d.Set("version", new_version_string)

	d.Set("last_updated", time.Now().Format(time.RFC850))

	return resourceAPMApplicationRead(ctx, d, m)
}

func resourceAPMApplicationDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

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
