## Unreleased

NOTES:

* Updated the return payload for clouds.
* Updated the return payload for networks.

FEATURES:

* **Network Pools** Client methods for `network_pools` API endpoints
* **Network Pool Servers** Client methods for `network_pool_servers` API endpoints
* **Network Routes** Client methods for `network_routes` API endpoints
* **Network Static Routes** Client methods for `network_static_routes` API endpoints
* **Network Subnets** Client methods for `network_subnets` API endpoints
* **Storage Servers** Client methods for `storage_servers` API endpoints

## 0.2.6 (December 22, 2022)

NOTES:

* Updated the return payload for clouds.

FEATURES:

* **Activity** Client methods for `activity` API endpoints
* **Load Balancer Monitors** Client methods for `load_balancer_monitors` API endpoints
* **Load Balancer Pools** Client methods for `load_balancer_pools` API endpoints
* **Load Balancer Types** Client methods for `load_balancer_types` API endpoints
* **Load Balancer Virtual Servers** Client methods for `load_balancer_virtual_servers` API endpoints
* **Load Balancers** Client methods for `load_balancers` API endpoints
* **Network Proxies** Client methods for `network_proxies` API endpoints* 
* **Plugins** Client methods for `plugins` API endpoints
* **Security Packages** Client methods for `security_packages` API endpoints

## 0.2.5 (December 12, 2022)

NOTES:

* Update workflow payloads to return the task sets data.

## 0.2.4 (December 2, 2022)

NOTES:

* Update execute schedule data payload.
* Fix invalid policy data payload.
* Update task data payload.

## 0.2.3 (November 14, 2022)

NOTES:

* Update cluster resource naming policy data type for auto resolve naming conflict setting.

## 0.2.2 (November 10, 2022)

NOTES:

* Fix user role creation response payload struct
* Fix delayed delete policy response payload struct

## 0.2.1 (October 27, 2022)

NOTES:

* Update policies code
* Update data structs

## 0.2.0 (October 7, 2022)

NOTES:

* Added additional tests and examples for existing resources
* Updated the response payloads for existing resources

FEATURES:

* **Alerts** Client methods for `alerts` API endpoints
* **Check Apps** Client methods for `check_apps` API endpoints
* **Deployments** Client methods for `deployments` API endpoints

## 0.1.10 (September 15, 2022)

* **Archives** Client methods for `archive_buckets` API endpoints

## 0.1.9 (September 13, 2022)

FEATURES:

* **Boot Scripts** Client methods for `boot_scripts` API endpoints
* **Preseed Scripts** Client methods for `preseed_scripts` API endpoints

## 0.1.8 (September 12, 2022)

FEATURES:

* **Users** Client methods for `users` API endpoints
* **User Groups** Client methods for `user_groups` API endpoints
* **Whitelabel Settings** Client methods for `whitelabel_settings` API endpoints

## 0.1.7 (September 8, 2022)

FEATURES:

* **Appliance Settings** Client methods for `appliance_settings` API endpoints
* **Backup Settings** Client methods for `backup_settings` API endpoints
* **Cluster Types** Client methods for `cluster_types` API endpoints
* **License** Client methods for `license` API endpoints
* **Provisioning Settings** Client methods for `provisioning_settings` API endpoints
* **VDI Allocations** Client methods for `vdi_allocations` API endpoints
* **VDI Apps** Client methods for `vdi_aps` API endpoints

## 0.1.6 (September 8, 2022)

Retracted Release

## 0.1.5 (August 26, 2022)

FEATURES:

* **Budgets** Client methods for `budgets` API endpoints
* **Credentials** Client methods for `credentials` API endpoints
* **Identity Sources** Client methods for `user_sources` API endpoints
* **Scale Thresholds** Client methods for `scale_thresholds` API endpoints
* **Software Licenses** Client methods for `software_licenses` API endpoints
* **Storage Buckets** Client methods for `storage_buckets` API endpoints
* **VDI Gateways** Client methods for `vdi_gateways` API endpoints
* **VDI Pools** Client methods for `vdi_pools` API endpoints

## 0.1.4 (August 3, 2022)

FEATURES:

* **Check Groups** Client methods for `check_groups` API endpoints
* **Checks** Client methods for `checks` API endpoints
* **Contacts** Client methods for `contacts` API endpoints
* **Incidents** Client methods for `incidents` API endpoints
* **Jobs** Client methods for `jobs` API endpoints
* **Monitoring Apps** Client methods for `monitoring_apps` API endpoints
* **Node Types** Client methods for `node_types` API endpoints
* **Policies** Client methods for `policies` API endpoints
* **Prices** Client methods for `prices` API endpoints
* **Price Sets** Client methods for `price_sets` API endpoints
* **Provision Types** Client methods for `provision_types` API endpoints
* **Roles** Client methods for `roles` API endpoints
* **Service Plans** Client methods for `service_plans` API endpoints
* **Wikis** Client methods for `wikis` API endpoints

## 0.3 (May 17, 2021)

FEATURES:

* **Instance Layouts** Client methods for `instance_layouts` API endpoints
* **Instance Types** Client methods for `instance_types` API endpoints
* **Plans** Client methods for `plans` API endpoints
* **Resource Pools** Client methods for `resource_pools` API endpoints

## 0.2 (May 7, 2021)

FEATURES:

* **Accounts** Client methods for `accounts` (Tenants) API endpoints
* **Apps** Client methods for `apps` API endpoints
* **Blueprints** Client methods for `blueprints` API endpoints
* **Catalog Items** Client methods for `catalog_items` API endpoints
* **Clusters** Client methods for `clusters` API endpoints
* **Environments** Client methods for `environments` API endpoints
* **Option Lists** Client methods for `option_lists` API endpoints
* **Option Types** Client methods for `option_types` API endpoints
* **Task Sets** Client methods for `task_sets` API endpoints
* **Tasks** Client methods for `tasks` API endpoints

## 0.1 (November 27, 2019)

NOTES:

* This is a **BETA** version of the Morpheus Go SDK.

FEATURES:

* **Client** Base client for API communication with remote morpheus appliance.
* **Groups:** Client methods for `groups` API endpoints.
* **Clouds:** Client methods for `clouds` API endpoints.
* **Instances:** Client methods for `instances` API endpoints.
* **Network:** Client methods for `networks` API endpoints.
* **NetworkDomains:** Client methods for `network-domains` API endpoints.
* **Setup:** Client methods for `setup` API endpoints.
* **Whoami:** Client methods for `whoami` API endpoints.
