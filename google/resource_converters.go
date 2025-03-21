// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------
package google

// ResourceConverter returns a map of terraform resource types (i.e. `google_project`)
// to a slice of ResourceConverters.
//
// Modelling of relationships:
// terraform resources to CAI assets as []ResourceConverter:
// 1:1 = [ResourceConverter{Convert: convertAbc}]                  (len=1)
// 1:N = [ResourceConverter{Convert: convertAbc}, ...]             (len=N)
// N:1 = [ResourceConverter{Convert: convertAbc, merge: mergeAbc}] (len=1)
func ResourceConverters() map[string][]ResourceConverter {
	return map[string][]ResourceConverter{
		"google_compute_firewall":                          {resourceConverterComputeFirewall()},
		"google_compute_disk":                              {resourceConverterComputeDisk()},
		"google_compute_forwarding_rule":                   {resourceConverterComputeForwardingRule()},
		"google_compute_global_forwarding_rule":            {resourceConverterComputeGlobalForwardingRule()},
		"google_compute_instance":                          {resourceConverterComputeInstance()},
		"google_compute_network":                           {resourceConverterComputeNetwork()},
		"google_compute_subnetwork":                        {resourceConverterComputeSubnetwork()},
		"google_storage_bucket":                            {resourceConverterStorageBucket()},
		"google_sql_database_instance":                     {resourceConverterSQLDatabaseInstance()},
		"google_container_cluster":                         {resourceConverterContainerCluster()},
		"google_container_node_pool":                       {resourceConverterContainerNodePool()},
		"google_bigquery_dataset":                          {resourceConverterBigQueryDataset()},
		"google_spanner_instance":                          {resourceConverterSpannerInstance()},
		"google_project_service":                           {resourceConverterServiceUsage()},
		"google_pubsub_subscription":                       {resourceConverterPubsubSubscription()},
		"google_pubsub_topic":                              {resourceConverterPubsubTopic()},
		"google_kms_crypto_key":                            {resourceConverterKMSCryptoKey()},
		"google_kms_key_ring":                              {resourceConverterKMSKeyRing()},
		"google_filestore_instance":                        {resourceConverterFilestoreInstance()},
		"google_access_context_manager_service_perimeter":  {resourceConverterAccessContextManagerServicePerimeter()},
		"google_bigquery_table_iam_policy":                 {resourceConverterBigQueryTableIamPolicy()},
		"google_bigquery_table_iam_binding":                {resourceConverterBigQueryTableIamBinding()},
		"google_bigquery_table_iam_member":                 {resourceConverterBigQueryTableIamMember()},
		"google_binary_authorization_attestor_iam_policy":  {resourceConverterBinaryAuthorizationAttestorIamPolicy()},
		"google_binary_authorization_attestor_iam_binding": {resourceConverterBinaryAuthorizationAttestorIamBinding()},
		"google_binary_authorization_attestor_iam_member":  {resourceConverterBinaryAuthorizationAttestorIamMember()},
		"google_cloudfunctions_function_iam_policy":        {resourceConverterCloudFunctionsCloudFunctionIamPolicy()},
		"google_cloudfunctions_function_iam_binding":       {resourceConverterCloudFunctionsCloudFunctionIamBinding()},
		"google_cloudfunctions_function_iam_member":        {resourceConverterCloudFunctionsCloudFunctionIamMember()},
		"google_compute_disk_iam_policy":                   {resourceConverterComputeDiskIamPolicy()},
		"google_compute_disk_iam_binding":                  {resourceConverterComputeDiskIamBinding()},
		"google_compute_disk_iam_member":                   {resourceConverterComputeDiskIamMember()},
		"google_compute_image_iam_policy":                  {resourceConverterComputeImageIamPolicy()},
		"google_compute_image_iam_binding":                 {resourceConverterComputeImageIamBinding()},
		"google_compute_image_iam_member":                  {resourceConverterComputeImageIamMember()},
		"google_compute_instance_iam_policy":               {resourceConverterComputeInstanceIamPolicy()},
		"google_compute_instance_iam_binding":              {resourceConverterComputeInstanceIamBinding()},
		"google_compute_instance_iam_member":               {resourceConverterComputeInstanceIamMember()},
		"google_compute_region_disk_iam_policy":            {resourceConverterComputeRegionDiskIamPolicy()},
		"google_compute_region_disk_iam_binding":           {resourceConverterComputeRegionDiskIamBinding()},
		"google_compute_region_disk_iam_member":            {resourceConverterComputeRegionDiskIamMember()},
		"google_compute_subnetwork_iam_policy":             {resourceConverterComputeSubnetworkIamPolicy()},
		"google_compute_subnetwork_iam_binding":            {resourceConverterComputeSubnetworkIamBinding()},
		"google_compute_subnetwork_iam_member":             {resourceConverterComputeSubnetworkIamMember()},
		"google_data_catalog_entry_group_iam_policy":       {resourceConverterDataCatalogEntryGroupIamPolicy()},
		"google_data_catalog_entry_group_iam_binding":      {resourceConverterDataCatalogEntryGroupIamBinding()},
		"google_data_catalog_entry_group_iam_member":       {resourceConverterDataCatalogEntryGroupIamMember()},
		"google_data_catalog_tag_template_iam_policy":      {resourceConverterDataCatalogTagTemplateIamPolicy()},
		"google_data_catalog_tag_template_iam_binding":     {resourceConverterDataCatalogTagTemplateIamBinding()},
		"google_data_catalog_tag_template_iam_member":      {resourceConverterDataCatalogTagTemplateIamMember()},
		"google_healthcare_consent_store_iam_policy":       {resourceConverterHealthcareConsentStoreIamPolicy()},
		"google_healthcare_consent_store_iam_binding":      {resourceConverterHealthcareConsentStoreIamBinding()},
		"google_healthcare_consent_store_iam_member":       {resourceConverterHealthcareConsentStoreIamMember()},
		"google_iap_web_iam_policy":                        {resourceConverterIapWebIamPolicy()},
		"google_iap_web_iam_binding":                       {resourceConverterIapWebIamBinding()},
		"google_iap_web_iam_member":                        {resourceConverterIapWebIamMember()},
		"google_iap_tunnel_instance_iam_policy":            {resourceConverterIapTunnelInstanceIamPolicy()},
		"google_iap_tunnel_instance_iam_binding":           {resourceConverterIapTunnelInstanceIamBinding()},
		"google_iap_tunnel_instance_iam_member":            {resourceConverterIapTunnelInstanceIamMember()},
		"google_iap_tunnel_iam_policy":                     {resourceConverterIapTunnelIamPolicy()},
		"google_iap_tunnel_iam_binding":                    {resourceConverterIapTunnelIamBinding()},
		"google_iap_tunnel_iam_member":                     {resourceConverterIapTunnelIamMember()},
		"google_notebooks_instance_iam_policy":             {resourceConverterNotebooksInstanceIamPolicy()},
		"google_notebooks_instance_iam_binding":            {resourceConverterNotebooksInstanceIamBinding()},
		"google_notebooks_instance_iam_member":             {resourceConverterNotebooksInstanceIamMember()},
		"google_privateca_ca_pool_iam_policy":              {resourceConverterPrivatecaCaPoolIamPolicy()},
		"google_privateca_ca_pool_iam_binding":             {resourceConverterPrivatecaCaPoolIamBinding()},
		"google_privateca_ca_pool_iam_member":              {resourceConverterPrivatecaCaPoolIamMember()},
		"google_pubsub_topic_iam_policy":                   {resourceConverterPubsubTopicIamPolicy()},
		"google_pubsub_topic_iam_binding":                  {resourceConverterPubsubTopicIamBinding()},
		"google_pubsub_topic_iam_member":                   {resourceConverterPubsubTopicIamMember()},
		"google_secret_manager_secret_iam_policy":          {resourceConverterSecretManagerSecretIamPolicy()},
		"google_secret_manager_secret_iam_binding":         {resourceConverterSecretManagerSecretIamBinding()},
		"google_secret_manager_secret_iam_member":          {resourceConverterSecretManagerSecretIamMember()},
		"google_endpoints_service_iam_policy":              {resourceConverterServiceManagementServiceIamPolicy()},
		"google_endpoints_service_iam_binding":             {resourceConverterServiceManagementServiceIamBinding()},
		"google_endpoints_service_iam_member":              {resourceConverterServiceManagementServiceIamMember()},
		"google_storage_bucket_iam_policy":                 {resourceConverterStorageBucketIamPolicy()},
		"google_storage_bucket_iam_binding":                {resourceConverterStorageBucketIamBinding()},
		"google_storage_bucket_iam_member":                 {resourceConverterStorageBucketIamMember()},
		"google_project": {
			resourceConverterProject(),
			resourceConverterProjectBillingInfo(),
		},
		"google_bigtable_instance": {
			resourceConverterBigtableInstance(),
			resourceConverterBigtableCluster(),
		},
		"google_organization_iam_policy":     {resourceConverterOrganizationIamPolicy()},
		"google_organization_iam_binding":    {resourceConverterOrganizationIamBinding()},
		"google_organization_iam_member":     {resourceConverterOrganizationIamMember()},
		"google_project_organization_policy": {resourceConverterProjectOrgPolicy()},
		"google_folder_iam_policy":           {resourceConverterFolderIamPolicy()},
		"google_folder_iam_binding":          {resourceConverterFolderIamBinding()},
		"google_folder_iam_member":           {resourceConverterFolderIamMember()},
		"google_project_iam_policy":          {resourceConverterProjectIamPolicy()},
		"google_project_iam_binding":         {resourceConverterProjectIamBinding()},
		"google_project_iam_member":          {resourceConverterProjectIamMember()},
	}
}
