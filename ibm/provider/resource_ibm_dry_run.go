// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package provider

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var dryRunSchema map[string]*schema.Schema

func ResourceIBMDryRun(test *schema.Resource) *schema.Resource {
	dryRunSchema = test.Schema
	return &schema.Resource{
		Create:   resourceIBMDryRunCreate,
		Read:     resourceIBMDryRunRead,
		Update:   resourceIBMDryRunUpdate,
		Delete:   resourceIBMDryRunDelete,
		Importer: &schema.ResourceImporter{},

		Schema: test.Schema,
	}
}

func resourceIBMDryRunCreate(d *schema.ResourceData, meta interface{}) error {

	return resourceIBMDryRunRead(d, meta)
}

func resourceIBMDryRunRead(d *schema.ResourceData, meta interface{}) error {
	d.SetId("testdryrun")
	for k, v := range dryRunSchema {
		if v.Computed {
			value := flattenSchemaElements(v)
			d.Set(k, value)
		}

	}
	return nil
}

func resourceIBMDryRunUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceIBMDryRunRead(d, meta)
}

func resourceIBMDryRunDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func flattenSchemaElements(v *schema.Schema) interface{} {
	switch v.Type {
	case schema.TypeBool:
		result := true
		return result
	case schema.TypeString:
		return "dryrun"
	case schema.TypeInt:
		return 1
	case schema.TypeSet:
		result := flattenTypeSet(v)
		return result
	case schema.TypeList:
		log.Println("**********[ENTER]*********", v)
		result := flattenTypeSet(v)
		return result
	}
	return "dryrun"
}

func flattenTypeSet(v *schema.Schema) interface{} {
	if v.Elem != nil {
		if e, ok := v.Elem.(*schema.Schema); ok {
			if e.Type == schema.TypeString {
				return []string{"dryrun"}
			}
			if e.Type == schema.TypeInt {
				return []int{0}
			}
		}

		if es, ok := v.Elem.(*schema.Resource); ok {
			log.Printf("***********[1]************ %+v", es)
			elemSchemas := make([]map[string]interface{}, 0)
			elemSchema := make(map[string]interface{})
			for k, value := range es.Schema {
				log.Printf("***********[2222]************%s %+v", k, value)
				flattenedValue := flattenSchemaElements(value)
				log.Printf("***********[flattenedValue]************%s %+v", k, flattenedValue)
				if k == "group_id" {
					log.Println("***********[GROUPPPs]************", flattenedValue)
				}
				if k == "rules" {
					log.Println("***********[ROLESS]************", flattenedValue)
				}
				elemSchema[k] = flattenedValue
				elemSchemas = append(elemSchemas, elemSchema)
				log.Println("***********[elemROLESS]************", elemSchemas)
				return elemSchemas

			}
			return elemSchemas
		}
	}

	return nil
}
