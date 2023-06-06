// Licensed under the Mozilla Public License v2.0

package dnsservices

import (
	//"context"
	"fmt"
	"log"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	//"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	dnsLinkedZones = "linked_zones"
	//DnsLinkedZoneInstanceID             = "instance_id"
	//DnsLinkedZoneName                   = "name"
	//DnsLinkedZoneDescription            = "description"
	//DnsLinkedZoneLinkedTo               = "linked_to"
	//DnsLinkedZoneState                  = "state"
	//DnsLinkedZoneLabel                  = "label"
	//DnsLinkedZoneApprovalRequiredBefore = "approval_required_before"
	//DnsLinkedZoneCreatedOn              = "created_on"
	//DnsLinkedZoneModifiedOn             = "modified_on"
	//dnsLZOffset                         = "offset"
	//dnLSZLimit                          = "limit"
)

func DataSourceIBMDNSLinkedZones() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMDNSLinkedZonesRead,
		Schema: map[string]*schema.Schema{
			DnsLinkedZoneInstanceID: {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The GUID of the DNS Services instance.",
			},
			dnsLinkedZones: {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of linked zones.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						DnsLinkedZoneName: {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The name of the linked zone.",
						},
						DnsLinkedZoneDescription: {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Descriptive text of the linked zone.",
						},
						DnsLinkedZoneLinkedTo: {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The unique identifier of the zone it is linked to.",
						},
						DnsLinkedZoneState: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The state of the linked zone.",
						},
						DnsLinkedZoneLabel: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The label of the linked zone.",
						},
						DnsLinkedZoneApprovalRequiredBefore: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The time when the approval is required before.",
						},
						DnsLinkedZoneCreatedOn: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The time when the linked zone is created.",
						},
						DnsLinkedZoneModifiedOn: {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The recent time when the linked zone is modified.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMDNSLinkedZonesRead(d *schema.ResourceData, meta interface{}) error {

	sess, err := meta.(conns.ClientSession).PrivateDNSClientSession()
	if err != nil {
		return err
	}
	instanceID := d.Get(DnsLinkedZoneInstanceID).(string)

	//listLinkedZonesOptions := sess.NewListLinkedZonesOptions(instanceID)
	opt := sess.NewListLinkedZonesOptions(instanceID)
	//availableDNSZones, resp, err := sess.ListLinkedZonesWithContext(context, listLinkedZonesOptions)
	result, resp, err := sess.ListLinkedZones(opt)
	if err != nil {
		return fmt.Errorf("[ERROR] Error listing the Linked Zones %s:%s", err, resp)
	}
	log.Printf("[INFO] Got Linked Zones %+v", result)
	DnsLinkedZones := make([]map[string]interface{}, 0)
	for _, instance := range result.LinkedDnszones {
		log.Printf("[INFO] Got Linked Zone %+v", instance)
		DnsLinkedZone := map[string]interface{}{}
		DnsLinkedZone[DnsLinkedZoneInstanceID] = instance.InstanceID
		log.Printf("[INFO] Got Linked Zone Name %+v", *instance.Name)
		DnsLinkedZone[DnsLinkedZoneName] = *instance.Name
		DnsLinkedZone[DnsLinkedZoneDescription] = *instance.Description
		DnsLinkedZone[DnsLinkedZoneLinkedTo] = *instance.LinkedTo
		DnsLinkedZone[DnsLinkedZoneState] = *instance.State
		DnsLinkedZone[DnsLinkedZoneLabel] = *instance.Label
		DnsLinkedZone[DnsLinkedZoneApprovalRequiredBefore] = *instance.ApprovalRequiredBefore
		DnsLinkedZone[DnsLinkedZoneCreatedOn] = *instance.CreatedOn
		DnsLinkedZone[DnsLinkedZoneModifiedOn] = *instance.ModifiedOn
		log.Printf("[INFO] Created DNS Zone %+v", DnsLinkedZone)
		DnsLinkedZones = append(DnsLinkedZones, DnsLinkedZone)
	}
	log.Printf("[INFO] Created DNS Zones %+v", DnsLinkedZones)
	d.SetId(dataSourceIBMDNSLinkedZoneID(d))
	d.Set(dnsLinkedZones, DnsLinkedZones)
	//d.Set(DnsLinkedZoneName)
	return nil
}

func dataSourceIBMDNSLinkedZoneID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
