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

	}
	return nil
}

func resourceIBMDryRunUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceIBMDryRunRead(d, meta)
}

func resourceIBMDryRunDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
