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
	log.Println("***********")
	for k, v := range dryRunSchema {
		log.Println(k)
		log.Println(v.Type)
		// d.Set(k, "dryrun")
		if v.Computed {
			value := flattenSchemaElements(v)
			log.Println("******[1]******", value)
			d.Set(k, value)
			log.Println("******[1SET1]******")
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
	log.Println("*******[2]******", v.Type)
	switch v.Type {
	case schema.TypeBool:
		result := true
		log.Println("*******[3]******", result)
		return result
	case schema.TypeString:
		result := "dryrun"
		log.Println("*******[4]******", result)
		return "dryrun"
	case schema.TypeInt:
		return 1
	case schema.TypeSet:
		result := flattenTypeSet(v)
		log.Println("*******[5]******", result)
		return result
	case schema.TypeList:
		result := flattenTypeSet(v)
		log.Println("*******[6]******", result)
		return result
	}
	return "dryrun"
}

func flattenTypeSet(v *schema.Schema) interface{} {
	log.Println("[FLATTEN]", v)
	if v.Elem != nil {
		log.Println("***********[77]**********")
		if e, ok := v.Elem.(*schema.Schema); ok {
			if e.Type == schema.TypeString {
				log.Println("***********[7]**********")
				return []string{"dryrun"}
			}
			if e.Type == schema.TypeInt {
				log.Println("***********[8]**********")
				return []int{0}
			}
		}

		if es, ok := v.Elem.(*schema.Resource); ok {
			log.Printf("***********[111]********** %+v", es.Schema)
			// var recursiveFlattenList func(*schema.Schema) interface{}
			// recursiveFlattenList = func(v *schema.Schema) interface{} {
			log.Println("***********[9]**********")
			elemSchema := make(map[string]interface{})
			for k, value := range es.Schema {
				log.Println("***********[99]**********", k, v)
				if value.Type != schema.TypeList && value.Type != schema.TypeSet && value.Type != schema.TypeMap {
					flattenedValue := flattenSchemaElements(v)
					elemSchema[k] = flattenedValue
					log.Println("***********[90]**********", elemSchema)
					return elemSchema
				} else {
					log.Println("***********[10]**********", elemSchema)
					return nil
					// recursiveFlattenList(value)
				}

			}
			return nil
			// }
		}
	}

	return nil
}
