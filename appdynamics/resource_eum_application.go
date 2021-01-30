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

func resourceEUMApplication() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceEUMApplicationCreate,
		ReadContext:   resourceEUMApplicationRead,
		UpdateContext: resourceEUMApplicationUpdate,
		DeleteContext: resourceEUMApplicationDelete,
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
			"eum_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceEUMApplicationCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

  provider_data := m.(map[string]string)
  base_url := provider_data["base_url"]
	token := provider_data["token"]

  url := base_url + "/controller/restui/allApplications/createApplication?applicationType=WEB"
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

	resourceEUMApplicationRead(ctx, d, m)

	return diags
}

func resourceEUMApplicationRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	provider_data := m.(map[string]string)
  base_url := provider_data["base_url"]
	token := provider_data["token"]

  url := base_url + "/controller/restui/eumApplications/getAllEumApplicationsData?time-range=last_1_hour.BEFORE_NOW.-1.-1.60"
	bearer := "Bearer " + token

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", bearer)
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	type Entries []struct {
		Name        string `json:"name"`
		AppKey      string `json:"appKey"`
		ID          int    `json:"id"`
	}

	data := Entries{}
	_ = json.Unmarshal([]byte(body), &data)

	for i := 0; i < len(data); i++ {
		if (data[i].Name == d.Get("name").(string)) {
			d.SetId(fmt.Sprint(data[i].ID))
			d.Set("eum_key", data[i].AppKey)
		}
	}

	return diags
}

func resourceEUMApplicationUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	provider_data := m.(map[string]string)
  base_url := provider_data["base_url"]
	token := provider_data["token"]

  url := base_url + "/controller/restui/allApplications/updateApplicationDetails"
	bearer := "Bearer " + token

	current_version, err := strconv.ParseInt(d.Get("version").(string), 10, 64)
	if err != nil {
		return diag.FromErr(err)
	}
	new_version := current_version + 1
	new_version_string := fmt.Sprint(new_version)
	d.Set("version", new_version_string)

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

	d.Set("last_updated", time.Now().Format(time.RFC850))

	return resourceEUMApplicationRead(ctx, d, m)
}

func resourceEUMApplicationDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

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
