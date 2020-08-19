package provider

import (
	"fmt"
	"regexp"

	"github.com/ani0989/terraform-provider-example/api/client"

	"github.com/hashicorp/terraform/helper/schema"
)

func resourceItem() *schema.Resource {
	return &schema.Resource{
		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				Description:  "The name of the resource",
				ForceNew:     true,
				ValidateFunc: validateName,
			},
		},
		Read:   resourceCreateItem,
		Create: resourceCreateItem,
		Update: resourceCreateItem,
		Delete: resourceCreateItem,
	}
}

func validateName(v interface{}, k string) (ws []string, es []error) {
	var errs []error
	var warns []string
	value, ok := v.(string)
	if !ok {
		errs = append(errs, fmt.Errorf("Expected name to be string"))
		return warns, errs
	}
	whiteSpace := regexp.MustCompile(`\s+`)
	if whiteSpace.Match([]byte(value)) {
		errs = append(errs, fmt.Errorf("No whitespaces allowed"))
		return warns, errs
	}
	return warns, errs
}

func resourceCreateItem(d *schema.ResourceData, m interface{}) error {
	apiClient := m.(*client.Client)
	fmt.Println("Schema : ")
	fmt.Println(d)

	todo := client.Todo{ID: 1}

	err := apiClient.PostReq(&todo)
	if err != nil {
		return err
	}
	fmt.Println("Get Req")
	fmt.Println(apiClient.GetReq())
	return nil
}
