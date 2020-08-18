package provider

import (
	"github.com/ani0989/terraform-provider-example/api/client"
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[String]*schema.Schema{
			"username": {
				Type:     schema.TypString,
				Required: false,
			},
			"password": {
				Type:     schema.TypString,
				Required: false,
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"example_item": resourceItem(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	return client.NewBasicAuthClient(username, password), nil
}
