---
subcategory: "Hyper Protect Crypto Service (HPCS)"
layout: "ibm"
page_title: "IBM : Hyper Protect Crypto Service instance"
description: |-
  Get information on an IBM Cloud Hyper Protect Crypto Service Instance.
---

# ibm\_hpcs

Imports a read only copy of an existing HPCS resource.

## Example Usage

```terraform
data "ibm_hpcs" "hpcs_instance" {
  name    = "test"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name used to identify the HPCS instance in the IBM Cloud UI.
* `resource_group_id` -(Optional, string) The Id of Resource Group

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - (String) The unique identifier of this hpcs instance.
* `plan` - (String) The service plan for this HPCS' instance
* `location` - (String) The location for this HPCS' instance
* `service` - (String) The service type (`hs-crypto`) of the instance.
* `status` - (String) Status of the hpcs instance.
* `guid` - (String) Unique identifier of resource instance.
* `crn` - (String) CRN of HPCS Instance
* `extensions` - (List) The extended metadata as a map associated with the resource instance.
* `hsm_info` - (List) HSM config of HPCS Instance Crypto Units.
  * `hsm_id` - (String) HSM ID
  * `hsm_location` - (String) HSM Location
  * `hsm_type` - (String) HSM Type.
  * `signature_threshold`- (Int) Signature Threshold for Crypto Units
  * `revocation_threshold` - ((Int)) Revocation Threshold for Crypto Units
  * `admins` - (List) List of Admins for Crypto Units
    * `name` - (String) Name of Admin
    * `ski` - (String) Admin SKI
  * `new_mk_status` - (String) Status of New Marster Key Registry
  * `new_mkvp` - (String) New Marster Key Registry Pattern
  * `current_mk_status` - (String) Status of Current Marster Key Registry
  * `current_mkvp` - (String) Current Marster Key Registry Pattern.