package incapsula

import (
	"context"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIncapsulaGcpPscTarget() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIncapsulaGcpPscTargetRead,
		Description: "Provides the GCP PSC (Private Service Connect) target address for a given region.",

		Schema: map[string]*schema.Schema{
			"region": {
				Type:        schema.TypeString,
				Description: "The region for which to retrieve the GCP PSC target.",
				Required:    true,
			},
			"account_id": {
				Type:        schema.TypeString,
				Description: "The account ID (caid) to scope the request to. When omitted, the account of the API credentials is used.",
				Optional:    true,
			},
			"address": {
				Type:        schema.TypeString,
				Description: "The GCP PSC target address for the specified region.",
				Computed:    true,
			},
		},
	}
}

func dataSourceIncapsulaGcpPscTargetRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := m.(*Client)

	region := d.Get("region").(string)
	accountID := d.Get("account_id").(string)

	resp, err := client.GetPscTarget(region, accountID)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("address", resp.Address); err != nil {
		return diag.FromErr(err)
	}
	if accountID != "" {
		d.SetId(accountID + "/" + region)
	} else {
		d.SetId(region)
	}

	return nil
}
