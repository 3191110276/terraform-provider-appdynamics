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

func resourceDBCollector() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceDBCollectorCreate,
		ReadContext:   resourceDBCollectorRead,
		UpdateContext: resourceDBCollectorUpdate,
		DeleteContext: resourceDBCollectorDelete,
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
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"agent_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
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
			"debugc": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceDBCollectorCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

  provider_data := m.(map[string]string)
  base_url := provider_data["base_url"]

  url := base_url + "/controller/rest/databases/collectors/create"

  req_string := "{\"type\":\"TYPE\",\"name\":\"NAME\",\"hostname\":\"HOST\",\"port\":\"PORT\",\"username\":\"USER\",\"password\":\"PASSWORD\",\"enabled\":true,\"agentName\":\"AGENT\"}"
  req_string = strings.Replace(req_string, "TYPE", d.Get("type").(string), 1)
	req_string = strings.Replace(req_string, "NAME", d.Get("name").(string), 1)
	req_string = strings.Replace(req_string, "HOST", d.Get("hostname").(string), 1)
	req_string = strings.Replace(req_string, "PORT", d.Get("port").(string), 1)
	req_string = strings.Replace(req_string, "USER", d.Get("username").(string), 1)
	req_string = strings.Replace(req_string, "PASSWORD", d.Get("password").(string), 1)
	req_string = strings.Replace(req_string, "AGENT", d.Get("agent_name").(string), 1)

  payload := strings.NewReader(req_string)

	req, _ := http.NewRequest("POST", url, payload)

	//req.SetBasicAuth(provider_data["username"], provider_data["password"])

  req.Header.Add("Authorization", "Basic bWltYXVyZXJAY2VlcjptaW1hdXJlcg==")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("cache-control", "no-cache")

	res, err := http.DefaultClient.Do(req)
  if err != nil {}

	defer res.Body.Close()
	//body, _ := ioutil.ReadAll(res.Body)

  if res.StatusCode >= 200 && res.StatusCode < 300 {
		d.Set("debuga", "Status 200")
	} else {
		d.Set("debuga", "Status not 200")
	}

	resourceDBCollectorRead(ctx, d, m)

	return diags
}

func resourceDBCollectorRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	provider_data := m.(map[string]string)
  base_url := provider_data["base_url"]

  url := base_url + "/controller/rest/databases/collectors"

	req, _ := http.NewRequest("GET", url, nil)

	req.SetBasicAuth(provider_data["username"], provider_data["password"])

	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	type Entries []struct {
		PerformanceState interface{} `json:"performanceState"`
		CollectorStatus  string      `json:"collectorStatus"`
		EventSummary     interface{} `json:"eventSummary"`
		ConfigID         int         `json:"configId"`
		NodeID           int         `json:"nodeId"`
		Config           struct {
			ID                       int           `json:"id"`
			Version                  int           `json:"version"`
			Name                     string        `json:"name"`
			NameUnique               bool          `json:"nameUnique"`
			BuiltIn                  bool          `json:"builtIn"`
			CreatedBy                interface{}   `json:"createdBy"`
			CreatedOn                int64         `json:"createdOn"`
			ModifiedBy               interface{}   `json:"modifiedBy"`
			ModifiedOn               int64         `json:"modifiedOn"`
			Type                     string        `json:"type"`
			Hostname                 string        `json:"hostname"`
			UseWindowsAuth           bool          `json:"useWindowsAuth"`
			Username                 string        `json:"username"`
			Password                 string        `json:"password"`
			Port                     int           `json:"port"`
			LoggingEnabled           bool          `json:"loggingEnabled"`
			Enabled                  bool          `json:"enabled"`
			ExcludedSchemas          interface{}   `json:"excludedSchemas"`
			JdbcConnectionProperties []interface{} `json:"jdbcConnectionProperties"`
			DatabaseName             string        `json:"databaseName"`
			FailoverPartner          interface{}   `json:"failoverPartner"`
			ConnectAsSysdba          bool          `json:"connectAsSysdba"`
			UseServiceName           bool          `json:"useServiceName"`
			Sid                      string        `json:"sid"`
			CustomConnectionString   interface{}   `json:"customConnectionString"`
			EnterpriseDB             bool          `json:"enterpriseDB"`
			UseSSL                   bool          `json:"useSSL"`
			EnableOSMonitor          bool          `json:"enableOSMonitor"`
			HostOS                   string        `json:"hostOS"`
			UseLocalWMI              bool          `json:"useLocalWMI"`
			HostDomain               interface{}   `json:"hostDomain"`
			HostUsername             string        `json:"hostUsername"`
			HostPassword             interface{}   `json:"hostPassword"`
			DbInstanceIdentifier     interface{}   `json:"dbInstanceIdentifier"`
			Region                   interface{}   `json:"region"`
			CertificateAuth          bool          `json:"certificateAuth"`
			RemoveLiterals           bool          `json:"removeLiterals"`
			SSHPort                  int           `json:"sshPort"`
			AgentName                string        `json:"agentName"`
			DbCyberArkEnabled        bool          `json:"dbCyberArkEnabled"`
			DbCyberArkApplication    interface{}   `json:"dbCyberArkApplication"`
			DbCyberArkSafe           interface{}   `json:"dbCyberArkSafe"`
			DbCyberArkFolder         interface{}   `json:"dbCyberArkFolder"`
			DbCyberArkObject         interface{}   `json:"dbCyberArkObject"`
			HwCyberArkEnabled        bool          `json:"hwCyberArkEnabled"`
			HwCyberArkApplication    interface{}   `json:"hwCyberArkApplication"`
			HwCyberArkSafe           interface{}   `json:"hwCyberArkSafe"`
			HwCyberArkFolder         interface{}   `json:"hwCyberArkFolder"`
			HwCyberArkObject         interface{}   `json:"hwCyberArkObject"`
			OrapkiSslEnabled         bool          `json:"orapkiSslEnabled"`
			OrasslClientAuthEnabled  bool          `json:"orasslClientAuthEnabled"`
			OrasslTruststoreLoc      interface{}   `json:"orasslTruststoreLoc"`
			OrasslTruststoreType     interface{}   `json:"orasslTruststoreType"`
			OrasslTruststorePassword interface{}   `json:"orasslTruststorePassword"`
			OrasslKeystoreLoc        interface{}   `json:"orasslKeystoreLoc"`
			OrasslKeystoreType       interface{}   `json:"orasslKeystoreType"`
			OrasslKeystorePassword   interface{}   `json:"orasslKeystorePassword"`
			LdapEnabled              bool          `json:"ldapEnabled"`
			CustomMetrics            interface{}   `json:"customMetrics"`
			SubConfigs               []interface{} `json:"subConfigs"`
			JmxPort                  int           `json:"jmxPort"`
			BackendIds               []int         `json:"backendIds"`
			ExtraProperties          []interface{} `json:"extraProperties"`
		} `json:"config"`
		LicensesUsed int `json:"licensesUsed"`
	}

	data := Entries{}
	_ = json.Unmarshal([]byte(body), &data)

	for i := 0; i < len(data); i++ {
	if (data[i].Config.Name == d.Get("name").(string)) {
	  d.Set("debuga", "found_entry")
	}
	}

	d.SetId("1")

	return diags
}

func resourceDBCollectorUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	provider_data := m.(map[string]string)
  base_url := provider_data["base_url"]

  url := base_url + "/controller/restui/allApplications/updateApplicationDetails"

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

	req.SetBasicAuth(provider_data["username"], provider_data["password"])

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	//body, _ := ioutil.ReadAll(res.Body)

	d.Set("last_updated", time.Now().Format(time.RFC850))

	return resourceDBCollectorRead(ctx, d, m)
}

func resourceDBCollectorDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	provider_data := m.(map[string]string)
  base_url := provider_data["base_url"]

  url := base_url + "/controller/rest/databases/collectors/REPLACEID"
	url = strings.Replace(url, "REPLACEID", d.Id(), 1)

	req, _ := http.NewRequest("DELETE", url, nil)

	req.SetBasicAuth(provider_data["username"], provider_data["password"])

	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

	d.SetId("")

	return diags
}
