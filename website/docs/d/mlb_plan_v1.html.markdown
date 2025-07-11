---
layout: "ecl"
page_title: "Enterprise Cloud: ecl_mlb_plan_v1"
sidebar_current: "docs-ecl-datasource-mlb-plan-v1"
description: |-
  Use this data source to get information of a plan within Enterprise Cloud Managed Load Balancer.
---

# ecl\_mlb\_plan\_v1

Use this data source to get information of a plan within Enterprise Cloud Managed Load Balancer.

## Example Usage

```hcl
data "ecl_mlb_plan_v1" "ha_50m_4if" {
  name = "50M_HA_4IF"
}
```

## Argument Reference

The following arguments are supported:

* `id` - (Optional) ID of the resource
* `name` - (Optional) Name of the resource
    * This field accepts single-byte characters only
* `description` - (Optional) Description of the resource
    * This field accepts single-byte characters only
* `bandwidth` - (Optional) Bandwidth of the plan
    * Must be one of these values:
        * `"50M"`
        * `"200M"`
        * `"1000M"`
        * `"3000M"`
* `redundancy` - (Optional) Redundancy of the plan
    * Must be one of these values:
        * `"HA"`
        * `"SINGLE"`
* `max_number_of_interfaces` - (Optional) Maximum number of interfaces for the plan
* `max_number_of_health_monitors` - (Optional) Maximum number of health monitors for the plan
* `max_number_of_listeners` - (Optional) Maximum number of listeners for the plan
* `max_number_of_policies` - (Optional) Maximum number of policies for the plan
* `max_number_of_routes` - (Optional) Maximum number of routes for the plan
* `max_number_of_target_groups` - (Optional) Maximum number of target groups for the plan
* `max_number_of_members` - (Optional) Maximum number of members for the target group of the plan
* `max_number_of_rules` - (Optional) Maximum number of rules for the policy of the plan
* `max_number_of_conditions` - (Optional) Maximum number of conditions in the rule of the plan
* `enabled` - (Optional) Whether a new load balancer can be created with this plan

## Attributes Reference

`id` is set to the ID of the found plan.<br>
In addition, the following attributes are exported:

* `name` - Name of the plan
* `description` - Description of the plan
* `bandwidth` - Bandwidth of the load balancer
* `redundancy` - Redundancy of the load balancer
* `max_number_of_interfaces` - Maximum number of interfaces for the load balancer
* `max_number_of_health_monitors` - Maximum number of health monitors for the load balancer
* `max_number_of_listeners` - Maximum number of listeners for the load balancer
* `max_number_of_policies` - Maximum number of policies for the load balancer
* `max_number_of_routes` - Maximum number of routes for the load balancer
* `max_number_of_target_groups` - Maximum number of target groups for the load balancer
* `max_number_of_members` - Maximum number of members for a target group
* `max_number_of_rules` - Maximum number of rules for a policy
* `max_number_of_conditions` - Maximum number of conditions in a rule
* `max_number_of_server_name_indications` - Maximum number of Server Name Indications (SNIs) in a policy
* `enabled` - Whether a new load balancer can be created with this plan
