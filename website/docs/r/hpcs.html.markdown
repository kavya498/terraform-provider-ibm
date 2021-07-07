---
subcategory: "Hyper Protect Crypto Service (HPCS)"
layout: "ibm"
page_title: "IBM : Hyper Protect Crypto Service instance"
description: |-
  Manages IBM Cloud Hyper Protect Crypto Service Instance.
---

# ibm\_hpcs

Manages HPCS resource. This allows hpcs sub-resources to be added to an existing hpcs instance.

## Example Usage

```terraform
resource ibm_hpcs hpcs {
  location             = "us-south"
  name                 = "test-hpcs"
  plan                 = "standard"
  units                = 2
  signature_threshold  = 1
  revocation_threshold = 1
  admins {
    name  = "admin1"
    key   = "/cloudTKE/1.sigkey"
    token = "sensitive1234"
  }
  admins {
    name  = "admin2"
    key   = "/cloudTKE/2.sigkey"
    token = "sensitive1234"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name used to identify the HPCS instance in the IBM Cloud UI.
* `plan` - (Required, string) The service plan for this HPCS' instance
* `location` - (Required, string) The location for this HPCS' instance
* `units` -(Required, string) Number of Crypto Units.
* `tags` - (Optional, array of strings) Tags associated with the instance.
* `service_endpoints` - (Optional, string) Types of the service endpoints that can be set to a resource instance. Possible values are 'public', 'private', 'public-and-private'.
* `resource_group_id` - (Optional, string) The Id of Resource Group
* `signature_threshold`- (Required, int)Signature Threshold of HSM.
* `revocation_threshold` - (Required, int) Revocation Threshold of HSM.
* `admins` - (Required, List) List of Admins for Crypto Units
  * `name` - (Required, string) Name of Admin
  * `key` - (Required, string) Path to the Signature file.
  * `token` - (Required, string) Password to access Signature file.
* `signature_server_url` - (Optional, string) URL of signing service.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - (String) The unique identifier of this hpcs instance.
* `plan` - (String) The service plan for this HPCS' instance
* `location` - (String) The location for this HPCS' instance
* `service` - (String) The service type (`hs-crypto`) of the instance.
* `status` - (String) Status of the hpcs instance.
* `state` - (String) The current state of the instance. For example, if the instance is deleted, it will return removed.
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
* `resource_aliases_url` - (String) The relative path to the resource aliases for the instance.
* `resource_bindings_url` - (String) The relative path to the resource bindings for the instance.
* `resource_keys_url` - (String) The relative path to the resource keys for the instance.
* `created_at` - (String) The date when the instance was created.
* `created_by` - (String) The subject who created the instance.
* `update_at` - (String) The date when the instance was last updated.
* `update_by` - (String) The subject who updated the instance.
* `deleted_at` - (String) The date when the instance was deleted.
* `deleted_by` - (String) The subject who deleted the instance.
* `scheduled_reclaim_at` - (String) The date when the instance was scheduled for reclamation.
* `scheduled_reclaim_by` - (String) The subject who initiated the instance reclamation.
* `restored_at` - (String) The date when the instance under reclamation was restored.
* `restored_by` - (String) The subject who restored the instance back from reclamation.
