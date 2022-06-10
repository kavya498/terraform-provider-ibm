// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package provider

import (
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
		result := flattenTypeSet(v)
		return result
	case schema.TypeMap:
		result := make(map[string]interface{})
		result["dryrun_key"] = "dryrun"
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
				return []int{1}
			}
		}

		if es, ok := v.Elem.(*schema.Resource); ok {
			elemSchemas := make([]map[string]interface{}, 0)
			elemSchema := make(map[string]interface{})
			for k, value := range es.Schema {
				flattenedValue := flattenSchemaElements(value)
				elemSchema[k] = flattenedValue
			}
			elemSchemas = append(elemSchemas, elemSchema)
			return elemSchemas
		}
	}

	return nil
}
