module github.com/3191110276/terraform-provider-appdynamics

go 1.13

require (
	//github.com/hashicorp-demoapp/hashicups-client-go v0.0.0-20200508203820-4c67e90efb8e
	//github.com/hashicorp-demoapp/hashicups-client-go v0.0.0-20200508203820-4c67e90efb8e
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.0.0-rc.2
)

replace (
    github.com/3191110276/terraform-provider-appdynamics/appdynamics => "./appdynamics"
)
