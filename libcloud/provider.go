package libcloud

import (
	"github.com/hashicorp/terraform/helper/schema"
)

// Provider function
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"instance": resourceInstance(),
		},
	}
}
