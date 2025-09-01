// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

type ClustersApi interface {

	/*
		AutoScalingConfiguration Return Auto Scaling Configuration for One Sharded Cluster

		Returns the internal configuration of AutoScaling for sharded clusters. This endpoint can be used for diagnostic purposes to ensure that sharded clusters updated from older APIs have gained support for AutoScaling each shard independently.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies this cluster.
		@return AutoScalingConfigurationApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for ClustersApi
	*/
	AutoScalingConfiguration(ctx context.Context, groupId string, clusterName string) AutoScalingConfigurationApiRequest
	/*
		AutoScalingConfiguration Return Auto Scaling Configuration for One Sharded Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param AutoScalingConfigurationApiParams - Parameters for the request
		@return AutoScalingConfigurationApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for ClustersApi
	*/
	AutoScalingConfigurationWithParams(ctx context.Context, args *AutoScalingConfigurationApiParams) AutoScalingConfigurationApiRequest

	// Method available only for mocking purposes
	AutoScalingConfigurationExecute(r AutoScalingConfigurationApiRequest) (*ClusterDescriptionAutoScalingModeConfiguration, *http.Response, error)

	/*
			CreateCluster Create One Cluster in One Project

			Creates one cluster in the specified project. Clusters contain a group of hosts that maintain the same data set. This resource can create clusters with asymmetrically-sized shards. Each project supports up to 25 database deployments. To use this resource, the requesting Service Account or API Key must have the Project Owner role. This feature is not available for serverless clusters.

		Please note that using an instanceSize of M2 or M5 will create a Flex cluster instead. Support for the instanceSize of M2 or M5 will be discontinued in January 2026. We recommend using the createFlexCluster API for such configurations moving forward. Deprecated versions: v2-{2024-08-05}, v2-{2023-02-01}, v2-{2023-01-01}

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
			@param clusterDescription20240805 Cluster to create in this project.
			@return CreateClusterApiRequest
	*/
	CreateCluster(ctx context.Context, groupId string, clusterDescription20240805 *ClusterDescription20240805) CreateClusterApiRequest
	/*
		CreateCluster Create One Cluster in One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateClusterApiParams - Parameters for the request
		@return CreateClusterApiRequest
	*/
	CreateClusterWithParams(ctx context.Context, args *CreateClusterApiParams) CreateClusterApiRequest

	// Method available only for mocking purposes
	CreateClusterExecute(r CreateClusterApiRequest) (*ClusterDescription20240805, *http.Response, error)

	/*
			DeleteCluster Remove One Cluster from One Project

			Removes one cluster from the specified project. The cluster must have termination protection disabled in order to be deleted. To use this resource, the requesting Service Account or API Key must have the Project Owner role. This feature is not available for serverless clusters.

		This endpoint can also be used on Flex clusters that were created using the [createCluster](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/v2/#tag/Clusters/operation/createCluster) endpoint or former M2/M5 clusters that have been migrated to Flex clusters until January 2026. Please use the deleteFlexCluster endpoint for Flex clusters instead. Deprecated versions: v2-{2023-01-01}

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
			@param clusterName Human-readable label that identifies the cluster.
			@return DeleteClusterApiRequest
	*/
	DeleteCluster(ctx context.Context, groupId string, clusterName string) DeleteClusterApiRequest
	/*
		DeleteCluster Remove One Cluster from One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteClusterApiParams - Parameters for the request
		@return DeleteClusterApiRequest
	*/
	DeleteClusterWithParams(ctx context.Context, args *DeleteClusterApiParams) DeleteClusterApiRequest

	// Method available only for mocking purposes
	DeleteClusterExecute(r DeleteClusterApiRequest) (*http.Response, error)

	/*
			GetCluster Return One Cluster from One Project

			Returns the details for one cluster in the specified project. Clusters contain a group of hosts that maintain the same data set. The response includes clusters with asymmetrically-sized shards. To use this resource, the requesting Service Account or API Key must have the Project Read Only role. This feature is not available for serverless clusters.

		This endpoint can also be used on Flex clusters that were created using the [createCluster](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/v2/#tag/Clusters/operation/createCluster) endpoint or former M2/M5 clusters that have been migrated to Flex clusters until January 2026. Please use the getFlexCluster endpoint for Flex clusters instead. Deprecated versions: v2-{2023-02-01}, v2-{2023-01-01}

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
			@param clusterName Human-readable label that identifies this cluster.
			@return GetClusterApiRequest
	*/
	GetCluster(ctx context.Context, groupId string, clusterName string) GetClusterApiRequest
	/*
		GetCluster Return One Cluster from One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetClusterApiParams - Parameters for the request
		@return GetClusterApiRequest
	*/
	GetClusterWithParams(ctx context.Context, args *GetClusterApiParams) GetClusterApiRequest

	// Method available only for mocking purposes
	GetClusterExecute(r GetClusterApiRequest) (*ClusterDescription20240805, *http.Response, error)

	/*
		GetClusterAdvancedConfiguration Return Advanced Configuration Options for One Cluster

		Returns the advanced configuration details for one cluster in the specified project. Clusters contain a group of hosts that maintain the same data set. Advanced configuration details include the read/write concern, index and oplog limits, and other database settings. This feature isn't available for `M0` free clusters, `M2` and `M5` shared-tier clusters, flex clusters, or serverless clusters. To use this resource, the requesting Service Account or API Key must have the Project Read Only role. Deprecated versions: v2-{2023-01-01}

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster.
		@return GetClusterAdvancedConfigurationApiRequest
	*/
	GetClusterAdvancedConfiguration(ctx context.Context, groupId string, clusterName string) GetClusterAdvancedConfigurationApiRequest
	/*
		GetClusterAdvancedConfiguration Return Advanced Configuration Options for One Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetClusterAdvancedConfigurationApiParams - Parameters for the request
		@return GetClusterAdvancedConfigurationApiRequest
	*/
	GetClusterAdvancedConfigurationWithParams(ctx context.Context, args *GetClusterAdvancedConfigurationApiParams) GetClusterAdvancedConfigurationApiRequest

	// Method available only for mocking purposes
	GetClusterAdvancedConfigurationExecute(r GetClusterAdvancedConfigurationApiRequest) (*ClusterDescriptionProcessArgs20240805, *http.Response, error)

	/*
		GetClusterStatus Return Status of All Cluster Operations

		Returns the status of all changes that you made to the specified cluster in the specified project. Use this resource to check the progress MongoDB Cloud has made in processing your changes. The response does not include the deployment of new dedicated clusters. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster.
		@return GetClusterStatusApiRequest
	*/
	GetClusterStatus(ctx context.Context, groupId string, clusterName string) GetClusterStatusApiRequest
	/*
		GetClusterStatus Return Status of All Cluster Operations


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetClusterStatusApiParams - Parameters for the request
		@return GetClusterStatusApiRequest
	*/
	GetClusterStatusWithParams(ctx context.Context, args *GetClusterStatusApiParams) GetClusterStatusApiRequest

	// Method available only for mocking purposes
	GetClusterStatusExecute(r GetClusterStatusApiRequest) (*ClusterStatus, *http.Response, error)

	/*
		GetSampleDatasetLoadStatus Return Status of Sample Dataset Load for One Cluster

		Checks the progress of loading the sample dataset into one cluster. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param sampleDatasetId Unique 24-hexadecimal digit string that identifies the loaded sample dataset.
		@return GetSampleDatasetLoadStatusApiRequest
	*/
	GetSampleDatasetLoadStatus(ctx context.Context, groupId string, sampleDatasetId string) GetSampleDatasetLoadStatusApiRequest
	/*
		GetSampleDatasetLoadStatus Return Status of Sample Dataset Load for One Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetSampleDatasetLoadStatusApiParams - Parameters for the request
		@return GetSampleDatasetLoadStatusApiRequest
	*/
	GetSampleDatasetLoadStatusWithParams(ctx context.Context, args *GetSampleDatasetLoadStatusApiParams) GetSampleDatasetLoadStatusApiRequest

	// Method available only for mocking purposes
	GetSampleDatasetLoadStatusExecute(r GetSampleDatasetLoadStatusApiRequest) (*SampleDatasetStatus, *http.Response, error)

	/*
		GrantMongoDbEmployeeAccess Grant MongoDB Employee Cluster Access for One Cluster

		Grants MongoDB employee cluster access for the given duration and at the specified level for one cluster.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies this cluster.
		@param employeeAccessGrant Grant access level and expiration.
		@return GrantMongoDbEmployeeAccessApiRequest
	*/
	GrantMongoDbEmployeeAccess(ctx context.Context, groupId string, clusterName string, employeeAccessGrant *EmployeeAccessGrant) GrantMongoDbEmployeeAccessApiRequest
	/*
		GrantMongoDbEmployeeAccess Grant MongoDB Employee Cluster Access for One Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GrantMongoDbEmployeeAccessApiParams - Parameters for the request
		@return GrantMongoDbEmployeeAccessApiRequest
	*/
	GrantMongoDbEmployeeAccessWithParams(ctx context.Context, args *GrantMongoDbEmployeeAccessApiParams) GrantMongoDbEmployeeAccessApiRequest

	// Method available only for mocking purposes
	GrantMongoDbEmployeeAccessExecute(r GrantMongoDbEmployeeAccessApiRequest) (*http.Response, error)

	/*
		ListCloudProviderRegions Return All Cloud Provider Regions

		Returns the list of regions available for the specified cloud provider at the specified tier. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@return ListCloudProviderRegionsApiRequest
	*/
	ListCloudProviderRegions(ctx context.Context, groupId string) ListCloudProviderRegionsApiRequest
	/*
		ListCloudProviderRegions Return All Cloud Provider Regions


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListCloudProviderRegionsApiParams - Parameters for the request
		@return ListCloudProviderRegionsApiRequest
	*/
	ListCloudProviderRegionsWithParams(ctx context.Context, args *ListCloudProviderRegionsApiParams) ListCloudProviderRegionsApiRequest

	// Method available only for mocking purposes
	ListCloudProviderRegionsExecute(r ListCloudProviderRegionsApiRequest) (*PaginatedApiAtlasProviderRegions, *http.Response, error)

	/*
			ListClusters Return All Clusters in One Project

			Returns the details for all clusters in the specific project to which you have access. Clusters contain a group of hosts that maintain the same data set. The response includes clusters with asymmetrically-sized shards. To use this resource, the requesting Service Account or API Key must have the Project Read Only role. This feature is not  available for serverless clusters.

		This endpoint can also be used on Flex clusters that were created using the [createCluster](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/v2/#tag/Clusters/operation/createCluster) endpoint or former M2/M5 clusters that have been migrated to Flex clusters until January 2026. Please use the listFlexClusters endpoint for Flex clusters instead. Deprecated versions: v2-{2023-02-01}, v2-{2023-01-01}

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
			@return ListClustersApiRequest
	*/
	ListClusters(ctx context.Context, groupId string) ListClustersApiRequest
	/*
		ListClusters Return All Clusters in One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListClustersApiParams - Parameters for the request
		@return ListClustersApiRequest
	*/
	ListClustersWithParams(ctx context.Context, args *ListClustersApiParams) ListClustersApiRequest

	// Method available only for mocking purposes
	ListClustersExecute(r ListClustersApiRequest) (*PaginatedClusterDescription20240805, *http.Response, error)

	/*
		ListClustersForAllProjects Return All Authorized Clusters in All Projects

		Returns the details for all clusters in all projects to which you have access. Clusters contain a group of hosts that maintain the same data set. The response does not include multi-cloud clusters. To use this resource, the requesting Service Account or API Key can have any cluster-level role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return ListClustersForAllProjectsApiRequest
	*/
	ListClustersForAllProjects(ctx context.Context) ListClustersForAllProjectsApiRequest
	/*
		ListClustersForAllProjects Return All Authorized Clusters in All Projects


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListClustersForAllProjectsApiParams - Parameters for the request
		@return ListClustersForAllProjectsApiRequest
	*/
	ListClustersForAllProjectsWithParams(ctx context.Context, args *ListClustersForAllProjectsApiParams) ListClustersForAllProjectsApiRequest

	// Method available only for mocking purposes
	ListClustersForAllProjectsExecute(r ListClustersForAllProjectsApiRequest) (*PaginatedOrgGroup, *http.Response, error)

	/*
		LoadSampleDataset Load Sample Dataset into One Cluster

		Requests loading the MongoDB sample dataset into the specified cluster. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param name Human-readable label that identifies the cluster into which you load the sample dataset.
		@return LoadSampleDatasetApiRequest
	*/
	LoadSampleDataset(ctx context.Context, groupId string, name string) LoadSampleDatasetApiRequest
	/*
		LoadSampleDataset Load Sample Dataset into One Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param LoadSampleDatasetApiParams - Parameters for the request
		@return LoadSampleDatasetApiRequest
	*/
	LoadSampleDatasetWithParams(ctx context.Context, args *LoadSampleDatasetApiParams) LoadSampleDatasetApiRequest

	// Method available only for mocking purposes
	LoadSampleDatasetExecute(r LoadSampleDatasetApiRequest) (*SampleDatasetStatus, *http.Response, error)

	/*
		PinFeatureCompatibilityVersion Pin Feature Compatibility Version for One Cluster in One Project

		Pins the Feature Compatibility Version (FCV) to the current MongoDB version and sets the pin expiration date. If an FCV pin already exists for the cluster, calling this method will only update the expiration date of the existing pin and will not repin the FCV.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies this cluster.
		@param pinFCV Optional request params for tuning FCV pinning configuration.
		@return PinFeatureCompatibilityVersionApiRequest
	*/
	PinFeatureCompatibilityVersion(ctx context.Context, groupId string, clusterName string, pinFCV *PinFCV) PinFeatureCompatibilityVersionApiRequest
	/*
		PinFeatureCompatibilityVersion Pin Feature Compatibility Version for One Cluster in One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param PinFeatureCompatibilityVersionApiParams - Parameters for the request
		@return PinFeatureCompatibilityVersionApiRequest
	*/
	PinFeatureCompatibilityVersionWithParams(ctx context.Context, args *PinFeatureCompatibilityVersionApiParams) PinFeatureCompatibilityVersionApiRequest

	// Method available only for mocking purposes
	PinFeatureCompatibilityVersionExecute(r PinFeatureCompatibilityVersionApiRequest) (*http.Response, error)

	/*
		RevokeMongoDbEmployeeAccess Revoke MongoDB Employee Cluster Access for One Cluster

		Revokes a previously granted MongoDB employee cluster access.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies this cluster.
		@return RevokeMongoDbEmployeeAccessApiRequest
	*/
	RevokeMongoDbEmployeeAccess(ctx context.Context, groupId string, clusterName string) RevokeMongoDbEmployeeAccessApiRequest
	/*
		RevokeMongoDbEmployeeAccess Revoke MongoDB Employee Cluster Access for One Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param RevokeMongoDbEmployeeAccessApiParams - Parameters for the request
		@return RevokeMongoDbEmployeeAccessApiRequest
	*/
	RevokeMongoDbEmployeeAccessWithParams(ctx context.Context, args *RevokeMongoDbEmployeeAccessApiParams) RevokeMongoDbEmployeeAccessApiRequest

	// Method available only for mocking purposes
	RevokeMongoDbEmployeeAccessExecute(r RevokeMongoDbEmployeeAccessApiRequest) (*http.Response, error)

	/*
		TestFailover Test Failover for One Cluster

		Starts a failover test for the specified cluster in the specified project. Clusters contain a group of hosts that maintain the same data set. A failover test checks how MongoDB Cloud handles the failure of the cluster's primary node. During the test, MongoDB Cloud shuts down the primary node and elects a new primary. To use this resource, the requesting Service Account or API Key must have the Project Cluster Manager role. Deprecated versions: v2-{2023-01-01}

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster.
		@return TestFailoverApiRequest
	*/
	TestFailover(ctx context.Context, groupId string, clusterName string) TestFailoverApiRequest
	/*
		TestFailover Test Failover for One Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param TestFailoverApiParams - Parameters for the request
		@return TestFailoverApiRequest
	*/
	TestFailoverWithParams(ctx context.Context, args *TestFailoverApiParams) TestFailoverApiRequest

	// Method available only for mocking purposes
	TestFailoverExecute(r TestFailoverApiRequest) (*http.Response, error)

	/*
		UnpinFeatureCompatibilityVersion Unpin Feature Compatibility Version for One Cluster in One Project

		Unpins the current fixed Feature Compatibility Version (FCV). This feature is not available for clusters on rapid release.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies this cluster.
		@return UnpinFeatureCompatibilityVersionApiRequest
	*/
	UnpinFeatureCompatibilityVersion(ctx context.Context, groupId string, clusterName string) UnpinFeatureCompatibilityVersionApiRequest
	/*
		UnpinFeatureCompatibilityVersion Unpin Feature Compatibility Version for One Cluster in One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UnpinFeatureCompatibilityVersionApiParams - Parameters for the request
		@return UnpinFeatureCompatibilityVersionApiRequest
	*/
	UnpinFeatureCompatibilityVersionWithParams(ctx context.Context, args *UnpinFeatureCompatibilityVersionApiParams) UnpinFeatureCompatibilityVersionApiRequest

	// Method available only for mocking purposes
	UnpinFeatureCompatibilityVersionExecute(r UnpinFeatureCompatibilityVersionApiRequest) (*http.Response, error)

	/*
		UpdateCluster Update One Cluster in One Project

		Updates the details for one cluster in the specified project. Clusters contain a group of hosts that maintain the same data set. This resource can update clusters with asymmetrically-sized shards. To update a cluster's termination protection, the requesting Service Account or API Key must have the Project Owner role. For all other updates, the requesting Service Account or API Key must have the Project Cluster Manager role. You can't modify a paused cluster (`paused : true`). You must call this endpoint to set `paused : false`. After this endpoint responds with `paused : false`, you can call it again with the changes you want to make to the cluster. This feature is not available for serverless clusters. Deprecated versions: v2-{2024-08-05}, v2-{2023-02-01}, v2-{2023-01-01}

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster.
		@param clusterDescription20240805 Cluster to update in the specified project.
		@return UpdateClusterApiRequest
	*/
	UpdateCluster(ctx context.Context, groupId string, clusterName string, clusterDescription20240805 *ClusterDescription20240805) UpdateClusterApiRequest
	/*
		UpdateCluster Update One Cluster in One Project


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateClusterApiParams - Parameters for the request
		@return UpdateClusterApiRequest
	*/
	UpdateClusterWithParams(ctx context.Context, args *UpdateClusterApiParams) UpdateClusterApiRequest

	// Method available only for mocking purposes
	UpdateClusterExecute(r UpdateClusterApiRequest) (*ClusterDescription20240805, *http.Response, error)

	/*
		UpdateClusterAdvancedConfiguration Update Advanced Configuration Options for One Cluster

		Updates the advanced configuration details for one cluster in the specified project. Clusters contain a group of hosts that maintain the same data set. Advanced configuration details include the read/write concern, index and oplog limits, and other database settings. To use this resource, the requesting Service Account or API Key must have the Project Cluster Manager role. This feature isn't available for `M0` free clusters, `M2` and `M5` shared-tier clusters, flex clusters, or serverless clusters. Deprecated versions: v2-{2023-01-01}

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Human-readable label that identifies the cluster.
		@param clusterDescriptionProcessArgs20240805 Advanced configuration details to add for one cluster in the specified project.
		@return UpdateClusterAdvancedConfigurationApiRequest
	*/
	UpdateClusterAdvancedConfiguration(ctx context.Context, groupId string, clusterName string, clusterDescriptionProcessArgs20240805 *ClusterDescriptionProcessArgs20240805) UpdateClusterAdvancedConfigurationApiRequest
	/*
		UpdateClusterAdvancedConfiguration Update Advanced Configuration Options for One Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateClusterAdvancedConfigurationApiParams - Parameters for the request
		@return UpdateClusterAdvancedConfigurationApiRequest
	*/
	UpdateClusterAdvancedConfigurationWithParams(ctx context.Context, args *UpdateClusterAdvancedConfigurationApiParams) UpdateClusterAdvancedConfigurationApiRequest

	// Method available only for mocking purposes
	UpdateClusterAdvancedConfigurationExecute(r UpdateClusterAdvancedConfigurationApiRequest) (*ClusterDescriptionProcessArgs20240805, *http.Response, error)

	/*
			UpgradeSharedCluster Upgrade One Shared-Tier Cluster

			Upgrades a shared-tier cluster to a Flex or Dedicated (M10+) cluster in the specified project. To use this resource, the requesting Service Account or API Key must have the Project Cluster Manager role. Each project supports up to 25 clusters.

		This endpoint can also be used to upgrade Flex clusters that were created using the [createCluster](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/v2/#tag/Clusters/operation/createCluster) API or former M2/M5 clusters that have been migrated to Flex clusters, using instanceSizeName to “M2” or “M5” until January 2026. This functionality will be available until January 2026, after which it will only be available for M0 clusters. Please use the upgradeFlexCluster endpoint instead.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
			@param legacyAtlasTenantClusterUpgradeRequest Details of the shared-tier cluster upgrade in the specified project.
			@return UpgradeSharedClusterApiRequest
	*/
	UpgradeSharedCluster(ctx context.Context, groupId string, legacyAtlasTenantClusterUpgradeRequest *LegacyAtlasTenantClusterUpgradeRequest) UpgradeSharedClusterApiRequest
	/*
		UpgradeSharedCluster Upgrade One Shared-Tier Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpgradeSharedClusterApiParams - Parameters for the request
		@return UpgradeSharedClusterApiRequest
	*/
	UpgradeSharedClusterWithParams(ctx context.Context, args *UpgradeSharedClusterApiParams) UpgradeSharedClusterApiRequest

	// Method available only for mocking purposes
	UpgradeSharedClusterExecute(r UpgradeSharedClusterApiRequest) (*LegacyAtlasCluster, *http.Response, error)

	/*
			UpgradeSharedClusterToServerless Upgrade One Shared-Tier Cluster to One Serverless Instance

			This endpoint has been deprecated as of February 2025 as we no longer support the creation of new serverless instances. Please use the upgradeFlexCluster endpoint to upgrade Flex clusters.

		 Upgrades a shared-tier cluster to a serverless instance in the specified project. To use this resource, the requesting Service Account or API Key must have the Project Cluster Manager role.

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
			@param serverlessInstanceDescription Details of the shared-tier cluster upgrade in the specified project.
			@return UpgradeSharedClusterToServerlessApiRequest

			Deprecated: this method has been deprecated. Please check the latest resource version for ClustersApi
	*/
	UpgradeSharedClusterToServerless(ctx context.Context, groupId string, serverlessInstanceDescription *ServerlessInstanceDescription) UpgradeSharedClusterToServerlessApiRequest
	/*
		UpgradeSharedClusterToServerless Upgrade One Shared-Tier Cluster to One Serverless Instance


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpgradeSharedClusterToServerlessApiParams - Parameters for the request
		@return UpgradeSharedClusterToServerlessApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for ClustersApi
	*/
	UpgradeSharedClusterToServerlessWithParams(ctx context.Context, args *UpgradeSharedClusterToServerlessApiParams) UpgradeSharedClusterToServerlessApiRequest

	// Method available only for mocking purposes
	UpgradeSharedClusterToServerlessExecute(r UpgradeSharedClusterToServerlessApiRequest) (*ServerlessInstanceDescription, *http.Response, error)
}

// ClustersApiService ClustersApi service
type ClustersApiService service

type AutoScalingConfigurationApiRequest struct {
	ctx         context.Context
	ApiService  ClustersApi
	groupId     string
	clusterName string
}

type AutoScalingConfigurationApiParams struct {
	GroupId     string
	ClusterName string
}

func (a *ClustersApiService) AutoScalingConfigurationWithParams(ctx context.Context, args *AutoScalingConfigurationApiParams) AutoScalingConfigurationApiRequest {
	return AutoScalingConfigurationApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
	}
}

func (r AutoScalingConfigurationApiRequest) Execute() (*ClusterDescriptionAutoScalingModeConfiguration, *http.Response, error) {
	return r.ApiService.AutoScalingConfigurationExecute(r)
}

/*
AutoScalingConfiguration Return Auto Scaling Configuration for One Sharded Cluster

Returns the internal configuration of AutoScaling for sharded clusters. This endpoint can be used for diagnostic purposes to ensure that sharded clusters updated from older APIs have gained support for AutoScaling each shard independently.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies this cluster.
	@return AutoScalingConfigurationApiRequest

Deprecated
*/
func (a *ClustersApiService) AutoScalingConfiguration(ctx context.Context, groupId string, clusterName string) AutoScalingConfigurationApiRequest {
	return AutoScalingConfigurationApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
	}
}

// AutoScalingConfigurationExecute executes the request
//
//	@return ClusterDescriptionAutoScalingModeConfiguration
//
// Deprecated
func (a *ClustersApiService) AutoScalingConfigurationExecute(r AutoScalingConfigurationApiRequest) (*ClusterDescriptionAutoScalingModeConfiguration, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ClusterDescriptionAutoScalingModeConfiguration
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClustersApiService.AutoScalingConfiguration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/autoScalingConfiguration"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return localVarReturnValue, nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-08-05+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		defer localVarHTTPResponse.Body.Close()
		buf, readErr := io.ReadAll(localVarHTTPResponse.Body)
		if readErr != nil {
			err = readErr
		}
		newErr := &GenericOpenAPIError{
			body:  buf,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type CreateClusterApiRequest struct {
	ctx                        context.Context
	ApiService                 ClustersApi
	groupId                    string
	clusterDescription20240805 *ClusterDescription20240805
}

type CreateClusterApiParams struct {
	GroupId                    string
	ClusterDescription20240805 *ClusterDescription20240805
}

func (a *ClustersApiService) CreateClusterWithParams(ctx context.Context, args *CreateClusterApiParams) CreateClusterApiRequest {
	return CreateClusterApiRequest{
		ApiService:                 a,
		ctx:                        ctx,
		groupId:                    args.GroupId,
		clusterDescription20240805: args.ClusterDescription20240805,
	}
}

func (r CreateClusterApiRequest) Execute() (*ClusterDescription20240805, *http.Response, error) {
	return r.ApiService.CreateClusterExecute(r)
}

/*
CreateCluster Create One Cluster in One Project

Creates one cluster in the specified project. Clusters contain a group of hosts that maintain the same data set. This resource can create clusters with asymmetrically-sized shards. Each project supports up to 25 database deployments. To use this resource, the requesting Service Account or API Key must have the Project Owner role. This feature is not available for serverless clusters.

Please note that using an instanceSize of M2 or M5 will create a Flex cluster instead. Support for the instanceSize of M2 or M5 will be discontinued in January 2026. We recommend using the createFlexCluster API for such configurations moving forward. Deprecated versions: v2-{2024-08-05}, v2-{2023-02-01}, v2-{2023-01-01}

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return CreateClusterApiRequest
*/
func (a *ClustersApiService) CreateCluster(ctx context.Context, groupId string, clusterDescription20240805 *ClusterDescription20240805) CreateClusterApiRequest {
	return CreateClusterApiRequest{
		ApiService:                 a,
		ctx:                        ctx,
		groupId:                    groupId,
		clusterDescription20240805: clusterDescription20240805,
	}
}

// CreateClusterExecute executes the request
//
//	@return ClusterDescription20240805
func (a *ClustersApiService) CreateClusterExecute(r CreateClusterApiRequest) (*ClusterDescription20240805, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ClusterDescription20240805
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClustersApiService.CreateCluster")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.clusterDescription20240805 == nil {
		return localVarReturnValue, nil, reportError("clusterDescription20240805 is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2024-10-23+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-10-23+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.clusterDescription20240805
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		defer localVarHTTPResponse.Body.Close()
		buf, readErr := io.ReadAll(localVarHTTPResponse.Body)
		if readErr != nil {
			err = readErr
		}
		newErr := &GenericOpenAPIError{
			body:  buf,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type DeleteClusterApiRequest struct {
	ctx           context.Context
	ApiService    ClustersApi
	groupId       string
	clusterName   string
	retainBackups *bool
}

type DeleteClusterApiParams struct {
	GroupId       string
	ClusterName   string
	RetainBackups *bool
}

func (a *ClustersApiService) DeleteClusterWithParams(ctx context.Context, args *DeleteClusterApiParams) DeleteClusterApiRequest {
	return DeleteClusterApiRequest{
		ApiService:    a,
		ctx:           ctx,
		groupId:       args.GroupId,
		clusterName:   args.ClusterName,
		retainBackups: args.RetainBackups,
	}
}

// Flag that indicates whether to retain backup snapshots for the deleted dedicated cluster.
func (r DeleteClusterApiRequest) RetainBackups(retainBackups bool) DeleteClusterApiRequest {
	r.retainBackups = &retainBackups
	return r
}

func (r DeleteClusterApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteClusterExecute(r)
}

/*
DeleteCluster Remove One Cluster from One Project

Removes one cluster from the specified project. The cluster must have termination protection disabled in order to be deleted. To use this resource, the requesting Service Account or API Key must have the Project Owner role. This feature is not available for serverless clusters.

This endpoint can also be used on Flex clusters that were created using the [createCluster](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/v2/#tag/Clusters/operation/createCluster) endpoint or former M2/M5 clusters that have been migrated to Flex clusters until January 2026. Please use the deleteFlexCluster endpoint for Flex clusters instead. Deprecated versions: v2-{2023-01-01}

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster.
	@return DeleteClusterApiRequest
*/
func (a *ClustersApiService) DeleteCluster(ctx context.Context, groupId string, clusterName string) DeleteClusterApiRequest {
	return DeleteClusterApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
	}
}

// DeleteClusterExecute executes the request
func (a *ClustersApiService) DeleteClusterExecute(r DeleteClusterApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClustersApiService.DeleteCluster")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}"
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.retainBackups != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "retainBackups", r.retainBackups, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-02-01+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type GetClusterApiRequest struct {
	ctx         context.Context
	ApiService  ClustersApi
	groupId     string
	clusterName string
}

type GetClusterApiParams struct {
	GroupId     string
	ClusterName string
}

func (a *ClustersApiService) GetClusterWithParams(ctx context.Context, args *GetClusterApiParams) GetClusterApiRequest {
	return GetClusterApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
	}
}

func (r GetClusterApiRequest) Execute() (*ClusterDescription20240805, *http.Response, error) {
	return r.ApiService.GetClusterExecute(r)
}

/*
GetCluster Return One Cluster from One Project

Returns the details for one cluster in the specified project. Clusters contain a group of hosts that maintain the same data set. The response includes clusters with asymmetrically-sized shards. To use this resource, the requesting Service Account or API Key must have the Project Read Only role. This feature is not available for serverless clusters.

This endpoint can also be used on Flex clusters that were created using the [createCluster](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/v2/#tag/Clusters/operation/createCluster) endpoint or former M2/M5 clusters that have been migrated to Flex clusters until January 2026. Please use the getFlexCluster endpoint for Flex clusters instead. Deprecated versions: v2-{2023-02-01}, v2-{2023-01-01}

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies this cluster.
	@return GetClusterApiRequest
*/
func (a *ClustersApiService) GetCluster(ctx context.Context, groupId string, clusterName string) GetClusterApiRequest {
	return GetClusterApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
	}
}

// GetClusterExecute executes the request
//
//	@return ClusterDescription20240805
func (a *ClustersApiService) GetClusterExecute(r GetClusterApiRequest) (*ClusterDescription20240805, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ClusterDescription20240805
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClustersApiService.GetCluster")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return localVarReturnValue, nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-08-05+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		defer localVarHTTPResponse.Body.Close()
		buf, readErr := io.ReadAll(localVarHTTPResponse.Body)
		if readErr != nil {
			err = readErr
		}
		newErr := &GenericOpenAPIError{
			body:  buf,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type GetClusterAdvancedConfigurationApiRequest struct {
	ctx         context.Context
	ApiService  ClustersApi
	groupId     string
	clusterName string
}

type GetClusterAdvancedConfigurationApiParams struct {
	GroupId     string
	ClusterName string
}

func (a *ClustersApiService) GetClusterAdvancedConfigurationWithParams(ctx context.Context, args *GetClusterAdvancedConfigurationApiParams) GetClusterAdvancedConfigurationApiRequest {
	return GetClusterAdvancedConfigurationApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
	}
}

func (r GetClusterAdvancedConfigurationApiRequest) Execute() (*ClusterDescriptionProcessArgs20240805, *http.Response, error) {
	return r.ApiService.GetClusterAdvancedConfigurationExecute(r)
}

/*
GetClusterAdvancedConfiguration Return Advanced Configuration Options for One Cluster

Returns the advanced configuration details for one cluster in the specified project. Clusters contain a group of hosts that maintain the same data set. Advanced configuration details include the read/write concern, index and oplog limits, and other database settings. This feature isn't available for `M0` free clusters, `M2` and `M5` shared-tier clusters, flex clusters, or serverless clusters. To use this resource, the requesting Service Account or API Key must have the Project Read Only role. Deprecated versions: v2-{2023-01-01}

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster.
	@return GetClusterAdvancedConfigurationApiRequest
*/
func (a *ClustersApiService) GetClusterAdvancedConfiguration(ctx context.Context, groupId string, clusterName string) GetClusterAdvancedConfigurationApiRequest {
	return GetClusterAdvancedConfigurationApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
	}
}

// GetClusterAdvancedConfigurationExecute executes the request
//
//	@return ClusterDescriptionProcessArgs20240805
func (a *ClustersApiService) GetClusterAdvancedConfigurationExecute(r GetClusterAdvancedConfigurationApiRequest) (*ClusterDescriptionProcessArgs20240805, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ClusterDescriptionProcessArgs20240805
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClustersApiService.GetClusterAdvancedConfiguration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/processArgs"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return localVarReturnValue, nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-08-05+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		defer localVarHTTPResponse.Body.Close()
		buf, readErr := io.ReadAll(localVarHTTPResponse.Body)
		if readErr != nil {
			err = readErr
		}
		newErr := &GenericOpenAPIError{
			body:  buf,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type GetClusterStatusApiRequest struct {
	ctx         context.Context
	ApiService  ClustersApi
	groupId     string
	clusterName string
}

type GetClusterStatusApiParams struct {
	GroupId     string
	ClusterName string
}

func (a *ClustersApiService) GetClusterStatusWithParams(ctx context.Context, args *GetClusterStatusApiParams) GetClusterStatusApiRequest {
	return GetClusterStatusApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
	}
}

func (r GetClusterStatusApiRequest) Execute() (*ClusterStatus, *http.Response, error) {
	return r.ApiService.GetClusterStatusExecute(r)
}

/*
GetClusterStatus Return Status of All Cluster Operations

Returns the status of all changes that you made to the specified cluster in the specified project. Use this resource to check the progress MongoDB Cloud has made in processing your changes. The response does not include the deployment of new dedicated clusters. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster.
	@return GetClusterStatusApiRequest
*/
func (a *ClustersApiService) GetClusterStatus(ctx context.Context, groupId string, clusterName string) GetClusterStatusApiRequest {
	return GetClusterStatusApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
	}
}

// GetClusterStatusExecute executes the request
//
//	@return ClusterStatus
func (a *ClustersApiService) GetClusterStatusExecute(r GetClusterStatusApiRequest) (*ClusterStatus, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ClusterStatus
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClustersApiService.GetClusterStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/status"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return localVarReturnValue, nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		defer localVarHTTPResponse.Body.Close()
		buf, readErr := io.ReadAll(localVarHTTPResponse.Body)
		if readErr != nil {
			err = readErr
		}
		newErr := &GenericOpenAPIError{
			body:  buf,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type GetSampleDatasetLoadStatusApiRequest struct {
	ctx             context.Context
	ApiService      ClustersApi
	groupId         string
	sampleDatasetId string
}

type GetSampleDatasetLoadStatusApiParams struct {
	GroupId         string
	SampleDatasetId string
}

func (a *ClustersApiService) GetSampleDatasetLoadStatusWithParams(ctx context.Context, args *GetSampleDatasetLoadStatusApiParams) GetSampleDatasetLoadStatusApiRequest {
	return GetSampleDatasetLoadStatusApiRequest{
		ApiService:      a,
		ctx:             ctx,
		groupId:         args.GroupId,
		sampleDatasetId: args.SampleDatasetId,
	}
}

func (r GetSampleDatasetLoadStatusApiRequest) Execute() (*SampleDatasetStatus, *http.Response, error) {
	return r.ApiService.GetSampleDatasetLoadStatusExecute(r)
}

/*
GetSampleDatasetLoadStatus Return Status of Sample Dataset Load for One Cluster

Checks the progress of loading the sample dataset into one cluster. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param sampleDatasetId Unique 24-hexadecimal digit string that identifies the loaded sample dataset.
	@return GetSampleDatasetLoadStatusApiRequest
*/
func (a *ClustersApiService) GetSampleDatasetLoadStatus(ctx context.Context, groupId string, sampleDatasetId string) GetSampleDatasetLoadStatusApiRequest {
	return GetSampleDatasetLoadStatusApiRequest{
		ApiService:      a,
		ctx:             ctx,
		groupId:         groupId,
		sampleDatasetId: sampleDatasetId,
	}
}

// GetSampleDatasetLoadStatusExecute executes the request
//
//	@return SampleDatasetStatus
func (a *ClustersApiService) GetSampleDatasetLoadStatusExecute(r GetSampleDatasetLoadStatusApiRequest) (*SampleDatasetStatus, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *SampleDatasetStatus
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClustersApiService.GetSampleDatasetLoadStatus")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/sampleDatasetLoad/{sampleDatasetId}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.sampleDatasetId == "" {
		return localVarReturnValue, nil, reportError("sampleDatasetId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"sampleDatasetId"+"}", url.PathEscape(r.sampleDatasetId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		defer localVarHTTPResponse.Body.Close()
		buf, readErr := io.ReadAll(localVarHTTPResponse.Body)
		if readErr != nil {
			err = readErr
		}
		newErr := &GenericOpenAPIError{
			body:  buf,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type GrantMongoDbEmployeeAccessApiRequest struct {
	ctx                 context.Context
	ApiService          ClustersApi
	groupId             string
	clusterName         string
	employeeAccessGrant *EmployeeAccessGrant
}

type GrantMongoDbEmployeeAccessApiParams struct {
	GroupId             string
	ClusterName         string
	EmployeeAccessGrant *EmployeeAccessGrant
}

func (a *ClustersApiService) GrantMongoDbEmployeeAccessWithParams(ctx context.Context, args *GrantMongoDbEmployeeAccessApiParams) GrantMongoDbEmployeeAccessApiRequest {
	return GrantMongoDbEmployeeAccessApiRequest{
		ApiService:          a,
		ctx:                 ctx,
		groupId:             args.GroupId,
		clusterName:         args.ClusterName,
		employeeAccessGrant: args.EmployeeAccessGrant,
	}
}

func (r GrantMongoDbEmployeeAccessApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.GrantMongoDbEmployeeAccessExecute(r)
}

/*
GrantMongoDbEmployeeAccess Grant MongoDB Employee Cluster Access for One Cluster

Grants MongoDB employee cluster access for the given duration and at the specified level for one cluster.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies this cluster.
	@return GrantMongoDbEmployeeAccessApiRequest
*/
func (a *ClustersApiService) GrantMongoDbEmployeeAccess(ctx context.Context, groupId string, clusterName string, employeeAccessGrant *EmployeeAccessGrant) GrantMongoDbEmployeeAccessApiRequest {
	return GrantMongoDbEmployeeAccessApiRequest{
		ApiService:          a,
		ctx:                 ctx,
		groupId:             groupId,
		clusterName:         clusterName,
		employeeAccessGrant: employeeAccessGrant,
	}
}

// GrantMongoDbEmployeeAccessExecute executes the request
func (a *ClustersApiService) GrantMongoDbEmployeeAccessExecute(r GrantMongoDbEmployeeAccessApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClustersApiService.GrantMongoDbEmployeeAccess")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}:grantMongoDBEmployeeAccess"
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.employeeAccessGrant == nil {
		return nil, reportError("employeeAccessGrant is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2024-08-05+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-08-05+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.employeeAccessGrant
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ListCloudProviderRegionsApiRequest struct {
	ctx          context.Context
	ApiService   ClustersApi
	groupId      string
	includeCount *bool
	itemsPerPage *int
	pageNum      *int
	providers    *[]string
	tier         *string
}

type ListCloudProviderRegionsApiParams struct {
	GroupId      string
	IncludeCount *bool
	ItemsPerPage *int
	PageNum      *int
	Providers    *[]string
	Tier         *string
}

func (a *ClustersApiService) ListCloudProviderRegionsWithParams(ctx context.Context, args *ListCloudProviderRegionsApiParams) ListCloudProviderRegionsApiRequest {
	return ListCloudProviderRegionsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		groupId:      args.GroupId,
		includeCount: args.IncludeCount,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
		providers:    args.Providers,
		tier:         args.Tier,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListCloudProviderRegionsApiRequest) IncludeCount(includeCount bool) ListCloudProviderRegionsApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListCloudProviderRegionsApiRequest) ItemsPerPage(itemsPerPage int) ListCloudProviderRegionsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListCloudProviderRegionsApiRequest) PageNum(pageNum int) ListCloudProviderRegionsApiRequest {
	r.pageNum = &pageNum
	return r
}

// Cloud providers whose regions to retrieve. When you specify multiple providers, the response can return only tiers and regions that support multi-cloud clusters.
func (r ListCloudProviderRegionsApiRequest) Providers(providers []string) ListCloudProviderRegionsApiRequest {
	r.providers = &providers
	return r
}

// Cluster tier for which to retrieve the regions.
func (r ListCloudProviderRegionsApiRequest) Tier(tier string) ListCloudProviderRegionsApiRequest {
	r.tier = &tier
	return r
}

func (r ListCloudProviderRegionsApiRequest) Execute() (*PaginatedApiAtlasProviderRegions, *http.Response, error) {
	return r.ApiService.ListCloudProviderRegionsExecute(r)
}

/*
ListCloudProviderRegions Return All Cloud Provider Regions

Returns the list of regions available for the specified cloud provider at the specified tier. To use this resource, the requesting Service Account or API Key must have the Project Read Only role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ListCloudProviderRegionsApiRequest
*/
func (a *ClustersApiService) ListCloudProviderRegions(ctx context.Context, groupId string) ListCloudProviderRegionsApiRequest {
	return ListCloudProviderRegionsApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// ListCloudProviderRegionsExecute executes the request
//
//	@return PaginatedApiAtlasProviderRegions
func (a *ClustersApiService) ListCloudProviderRegionsExecute(r ListCloudProviderRegionsApiRequest) (*PaginatedApiAtlasProviderRegions, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedApiAtlasProviderRegions
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClustersApiService.ListCloudProviderRegions")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/provider/regions"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.includeCount != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeCount", r.includeCount, "")
	} else {
		var defaultValue bool = true
		r.includeCount = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeCount", r.includeCount, "")
	}
	if r.itemsPerPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	} else {
		var defaultValue int = 100
		r.itemsPerPage = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	}
	if r.pageNum != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNum", r.pageNum, "")
	} else {
		var defaultValue int = 1
		r.pageNum = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNum", r.pageNum, "")
	}
	if r.providers != nil {
		t := *r.providers
		// Workaround for unused import
		_ = reflect.Append
		parameterAddToHeaderOrQuery(localVarQueryParams, "providers", t, "multi")

	}
	if r.tier != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "tier", r.tier, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		defer localVarHTTPResponse.Body.Close()
		buf, readErr := io.ReadAll(localVarHTTPResponse.Body)
		if readErr != nil {
			err = readErr
		}
		newErr := &GenericOpenAPIError{
			body:  buf,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ListClustersApiRequest struct {
	ctx                               context.Context
	ApiService                        ClustersApi
	groupId                           string
	includeCount                      *bool
	itemsPerPage                      *int
	pageNum                           *int
	includeDeletedWithRetainedBackups *bool
}

type ListClustersApiParams struct {
	GroupId                           string
	IncludeCount                      *bool
	ItemsPerPage                      *int
	PageNum                           *int
	IncludeDeletedWithRetainedBackups *bool
}

func (a *ClustersApiService) ListClustersWithParams(ctx context.Context, args *ListClustersApiParams) ListClustersApiRequest {
	return ListClustersApiRequest{
		ApiService:                        a,
		ctx:                               ctx,
		groupId:                           args.GroupId,
		includeCount:                      args.IncludeCount,
		itemsPerPage:                      args.ItemsPerPage,
		pageNum:                           args.PageNum,
		includeDeletedWithRetainedBackups: args.IncludeDeletedWithRetainedBackups,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListClustersApiRequest) IncludeCount(includeCount bool) ListClustersApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListClustersApiRequest) ItemsPerPage(itemsPerPage int) ListClustersApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListClustersApiRequest) PageNum(pageNum int) ListClustersApiRequest {
	r.pageNum = &pageNum
	return r
}

// Flag that indicates whether to return Clusters with retain backups.
func (r ListClustersApiRequest) IncludeDeletedWithRetainedBackups(includeDeletedWithRetainedBackups bool) ListClustersApiRequest {
	r.includeDeletedWithRetainedBackups = &includeDeletedWithRetainedBackups
	return r
}

func (r ListClustersApiRequest) Execute() (*PaginatedClusterDescription20240805, *http.Response, error) {
	return r.ApiService.ListClustersExecute(r)
}

/*
ListClusters Return All Clusters in One Project

Returns the details for all clusters in the specific project to which you have access. Clusters contain a group of hosts that maintain the same data set. The response includes clusters with asymmetrically-sized shards. To use this resource, the requesting Service Account or API Key must have the Project Read Only role. This feature is not  available for serverless clusters.

This endpoint can also be used on Flex clusters that were created using the [createCluster](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/v2/#tag/Clusters/operation/createCluster) endpoint or former M2/M5 clusters that have been migrated to Flex clusters until January 2026. Please use the listFlexClusters endpoint for Flex clusters instead. Deprecated versions: v2-{2023-02-01}, v2-{2023-01-01}

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return ListClustersApiRequest
*/
func (a *ClustersApiService) ListClusters(ctx context.Context, groupId string) ListClustersApiRequest {
	return ListClustersApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
	}
}

// ListClustersExecute executes the request
//
//	@return PaginatedClusterDescription20240805
func (a *ClustersApiService) ListClustersExecute(r ListClustersApiRequest) (*PaginatedClusterDescription20240805, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedClusterDescription20240805
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClustersApiService.ListClusters")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.includeCount != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeCount", r.includeCount, "")
	} else {
		var defaultValue bool = true
		r.includeCount = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeCount", r.includeCount, "")
	}
	if r.itemsPerPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	} else {
		var defaultValue int = 100
		r.itemsPerPage = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	}
	if r.pageNum != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNum", r.pageNum, "")
	} else {
		var defaultValue int = 1
		r.pageNum = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNum", r.pageNum, "")
	}
	if r.includeDeletedWithRetainedBackups != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeDeletedWithRetainedBackups", r.includeDeletedWithRetainedBackups, "")
	} else {
		var defaultValue bool = false
		r.includeDeletedWithRetainedBackups = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeDeletedWithRetainedBackups", r.includeDeletedWithRetainedBackups, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-08-05+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		defer localVarHTTPResponse.Body.Close()
		buf, readErr := io.ReadAll(localVarHTTPResponse.Body)
		if readErr != nil {
			err = readErr
		}
		newErr := &GenericOpenAPIError{
			body:  buf,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ListClustersForAllProjectsApiRequest struct {
	ctx          context.Context
	ApiService   ClustersApi
	includeCount *bool
	itemsPerPage *int
	pageNum      *int
}

type ListClustersForAllProjectsApiParams struct {
	IncludeCount *bool
	ItemsPerPage *int
	PageNum      *int
}

func (a *ClustersApiService) ListClustersForAllProjectsWithParams(ctx context.Context, args *ListClustersForAllProjectsApiParams) ListClustersForAllProjectsApiRequest {
	return ListClustersForAllProjectsApiRequest{
		ApiService:   a,
		ctx:          ctx,
		includeCount: args.IncludeCount,
		itemsPerPage: args.ItemsPerPage,
		pageNum:      args.PageNum,
	}
}

// Flag that indicates whether the response returns the total number of items (**totalCount**) in the response.
func (r ListClustersForAllProjectsApiRequest) IncludeCount(includeCount bool) ListClustersForAllProjectsApiRequest {
	r.includeCount = &includeCount
	return r
}

// Number of items that the response returns per page.
func (r ListClustersForAllProjectsApiRequest) ItemsPerPage(itemsPerPage int) ListClustersForAllProjectsApiRequest {
	r.itemsPerPage = &itemsPerPage
	return r
}

// Number of the page that displays the current set of the total objects that the response returns.
func (r ListClustersForAllProjectsApiRequest) PageNum(pageNum int) ListClustersForAllProjectsApiRequest {
	r.pageNum = &pageNum
	return r
}

func (r ListClustersForAllProjectsApiRequest) Execute() (*PaginatedOrgGroup, *http.Response, error) {
	return r.ApiService.ListClustersForAllProjectsExecute(r)
}

/*
ListClustersForAllProjects Return All Authorized Clusters in All Projects

Returns the details for all clusters in all projects to which you have access. Clusters contain a group of hosts that maintain the same data set. The response does not include multi-cloud clusters. To use this resource, the requesting Service Account or API Key can have any cluster-level role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return ListClustersForAllProjectsApiRequest
*/
func (a *ClustersApiService) ListClustersForAllProjects(ctx context.Context) ListClustersForAllProjectsApiRequest {
	return ListClustersForAllProjectsApiRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// ListClustersForAllProjectsExecute executes the request
//
//	@return PaginatedOrgGroup
func (a *ClustersApiService) ListClustersForAllProjectsExecute(r ListClustersForAllProjectsApiRequest) (*PaginatedOrgGroup, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *PaginatedOrgGroup
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClustersApiService.ListClustersForAllProjects")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/clusters"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.includeCount != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeCount", r.includeCount, "")
	} else {
		var defaultValue bool = true
		r.includeCount = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "includeCount", r.includeCount, "")
	}
	if r.itemsPerPage != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	} else {
		var defaultValue int = 100
		r.itemsPerPage = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "itemsPerPage", r.itemsPerPage, "")
	}
	if r.pageNum != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNum", r.pageNum, "")
	} else {
		var defaultValue int = 1
		r.pageNum = &defaultValue
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageNum", r.pageNum, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		defer localVarHTTPResponse.Body.Close()
		buf, readErr := io.ReadAll(localVarHTTPResponse.Body)
		if readErr != nil {
			err = readErr
		}
		newErr := &GenericOpenAPIError{
			body:  buf,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type LoadSampleDatasetApiRequest struct {
	ctx        context.Context
	ApiService ClustersApi
	groupId    string
	name       string
}

type LoadSampleDatasetApiParams struct {
	GroupId string
	Name    string
}

func (a *ClustersApiService) LoadSampleDatasetWithParams(ctx context.Context, args *LoadSampleDatasetApiParams) LoadSampleDatasetApiRequest {
	return LoadSampleDatasetApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    args.GroupId,
		name:       args.Name,
	}
}

func (r LoadSampleDatasetApiRequest) Execute() (*SampleDatasetStatus, *http.Response, error) {
	return r.ApiService.LoadSampleDatasetExecute(r)
}

/*
LoadSampleDataset Load Sample Dataset into One Cluster

Requests loading the MongoDB sample dataset into the specified cluster. To use this resource, the requesting Service Account or API Key must have the Project Owner role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param name Human-readable label that identifies the cluster into which you load the sample dataset.
	@return LoadSampleDatasetApiRequest
*/
func (a *ClustersApiService) LoadSampleDataset(ctx context.Context, groupId string, name string) LoadSampleDatasetApiRequest {
	return LoadSampleDatasetApiRequest{
		ApiService: a,
		ctx:        ctx,
		groupId:    groupId,
		name:       name,
	}
}

// LoadSampleDatasetExecute executes the request
//
//	@return SampleDatasetStatus
func (a *ClustersApiService) LoadSampleDatasetExecute(r LoadSampleDatasetApiRequest) (*SampleDatasetStatus, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *SampleDatasetStatus
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClustersApiService.LoadSampleDataset")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/sampleDatasetLoad/{name}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.name == "" {
		return localVarReturnValue, nil, reportError("name is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(r.name), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		defer localVarHTTPResponse.Body.Close()
		buf, readErr := io.ReadAll(localVarHTTPResponse.Body)
		if readErr != nil {
			err = readErr
		}
		newErr := &GenericOpenAPIError{
			body:  buf,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type PinFeatureCompatibilityVersionApiRequest struct {
	ctx         context.Context
	ApiService  ClustersApi
	groupId     string
	clusterName string
	pinFCV      *PinFCV
}

type PinFeatureCompatibilityVersionApiParams struct {
	GroupId     string
	ClusterName string
	PinFCV      *PinFCV
}

func (a *ClustersApiService) PinFeatureCompatibilityVersionWithParams(ctx context.Context, args *PinFeatureCompatibilityVersionApiParams) PinFeatureCompatibilityVersionApiRequest {
	return PinFeatureCompatibilityVersionApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
		pinFCV:      args.PinFCV,
	}
}

func (r PinFeatureCompatibilityVersionApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.PinFeatureCompatibilityVersionExecute(r)
}

/*
PinFeatureCompatibilityVersion Pin Feature Compatibility Version for One Cluster in One Project

Pins the Feature Compatibility Version (FCV) to the current MongoDB version and sets the pin expiration date. If an FCV pin already exists for the cluster, calling this method will only update the expiration date of the existing pin and will not repin the FCV.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies this cluster.
	@return PinFeatureCompatibilityVersionApiRequest
*/
func (a *ClustersApiService) PinFeatureCompatibilityVersion(ctx context.Context, groupId string, clusterName string, pinFCV *PinFCV) PinFeatureCompatibilityVersionApiRequest {
	return PinFeatureCompatibilityVersionApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
		pinFCV:      pinFCV,
	}
}

// PinFeatureCompatibilityVersionExecute executes the request
func (a *ClustersApiService) PinFeatureCompatibilityVersionExecute(r PinFeatureCompatibilityVersionApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClustersApiService.PinFeatureCompatibilityVersion")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}:pinFeatureCompatibilityVersion"
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2024-05-30+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-05-30+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.pinFCV
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type RevokeMongoDbEmployeeAccessApiRequest struct {
	ctx         context.Context
	ApiService  ClustersApi
	groupId     string
	clusterName string
}

type RevokeMongoDbEmployeeAccessApiParams struct {
	GroupId     string
	ClusterName string
}

func (a *ClustersApiService) RevokeMongoDbEmployeeAccessWithParams(ctx context.Context, args *RevokeMongoDbEmployeeAccessApiParams) RevokeMongoDbEmployeeAccessApiRequest {
	return RevokeMongoDbEmployeeAccessApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
	}
}

func (r RevokeMongoDbEmployeeAccessApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.RevokeMongoDbEmployeeAccessExecute(r)
}

/*
RevokeMongoDbEmployeeAccess Revoke MongoDB Employee Cluster Access for One Cluster

Revokes a previously granted MongoDB employee cluster access.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies this cluster.
	@return RevokeMongoDbEmployeeAccessApiRequest
*/
func (a *ClustersApiService) RevokeMongoDbEmployeeAccess(ctx context.Context, groupId string, clusterName string) RevokeMongoDbEmployeeAccessApiRequest {
	return RevokeMongoDbEmployeeAccessApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
	}
}

// RevokeMongoDbEmployeeAccessExecute executes the request
func (a *ClustersApiService) RevokeMongoDbEmployeeAccessExecute(r RevokeMongoDbEmployeeAccessApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClustersApiService.RevokeMongoDbEmployeeAccess")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}:revokeMongoDBEmployeeAccess"
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-08-05+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type TestFailoverApiRequest struct {
	ctx         context.Context
	ApiService  ClustersApi
	groupId     string
	clusterName string
}

type TestFailoverApiParams struct {
	GroupId     string
	ClusterName string
}

func (a *ClustersApiService) TestFailoverWithParams(ctx context.Context, args *TestFailoverApiParams) TestFailoverApiRequest {
	return TestFailoverApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
	}
}

func (r TestFailoverApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.TestFailoverExecute(r)
}

/*
TestFailover Test Failover for One Cluster

Starts a failover test for the specified cluster in the specified project. Clusters contain a group of hosts that maintain the same data set. A failover test checks how MongoDB Cloud handles the failure of the cluster's primary node. During the test, MongoDB Cloud shuts down the primary node and elects a new primary. To use this resource, the requesting Service Account or API Key must have the Project Cluster Manager role. Deprecated versions: v2-{2023-01-01}

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster.
	@return TestFailoverApiRequest
*/
func (a *ClustersApiService) TestFailover(ctx context.Context, groupId string, clusterName string) TestFailoverApiRequest {
	return TestFailoverApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
	}
}

// TestFailoverExecute executes the request
func (a *ClustersApiService) TestFailoverExecute(r TestFailoverApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClustersApiService.TestFailover")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/restartPrimaries"
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-02-01+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type UnpinFeatureCompatibilityVersionApiRequest struct {
	ctx         context.Context
	ApiService  ClustersApi
	groupId     string
	clusterName string
}

type UnpinFeatureCompatibilityVersionApiParams struct {
	GroupId     string
	ClusterName string
}

func (a *ClustersApiService) UnpinFeatureCompatibilityVersionWithParams(ctx context.Context, args *UnpinFeatureCompatibilityVersionApiParams) UnpinFeatureCompatibilityVersionApiRequest {
	return UnpinFeatureCompatibilityVersionApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
	}
}

func (r UnpinFeatureCompatibilityVersionApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.UnpinFeatureCompatibilityVersionExecute(r)
}

/*
UnpinFeatureCompatibilityVersion Unpin Feature Compatibility Version for One Cluster in One Project

Unpins the current fixed Feature Compatibility Version (FCV). This feature is not available for clusters on rapid release.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies this cluster.
	@return UnpinFeatureCompatibilityVersionApiRequest
*/
func (a *ClustersApiService) UnpinFeatureCompatibilityVersion(ctx context.Context, groupId string, clusterName string) UnpinFeatureCompatibilityVersionApiRequest {
	return UnpinFeatureCompatibilityVersionApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
	}
}

// UnpinFeatureCompatibilityVersionExecute executes the request
func (a *ClustersApiService) UnpinFeatureCompatibilityVersionExecute(r UnpinFeatureCompatibilityVersionApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodPost
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClustersApiService.UnpinFeatureCompatibilityVersion")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}:unpinFeatureCompatibilityVersion"
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-05-30+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type UpdateClusterApiRequest struct {
	ctx                        context.Context
	ApiService                 ClustersApi
	groupId                    string
	clusterName                string
	clusterDescription20240805 *ClusterDescription20240805
}

type UpdateClusterApiParams struct {
	GroupId                    string
	ClusterName                string
	ClusterDescription20240805 *ClusterDescription20240805
}

func (a *ClustersApiService) UpdateClusterWithParams(ctx context.Context, args *UpdateClusterApiParams) UpdateClusterApiRequest {
	return UpdateClusterApiRequest{
		ApiService:                 a,
		ctx:                        ctx,
		groupId:                    args.GroupId,
		clusterName:                args.ClusterName,
		clusterDescription20240805: args.ClusterDescription20240805,
	}
}

func (r UpdateClusterApiRequest) Execute() (*ClusterDescription20240805, *http.Response, error) {
	return r.ApiService.UpdateClusterExecute(r)
}

/*
UpdateCluster Update One Cluster in One Project

Updates the details for one cluster in the specified project. Clusters contain a group of hosts that maintain the same data set. This resource can update clusters with asymmetrically-sized shards. To update a cluster's termination protection, the requesting Service Account or API Key must have the Project Owner role. For all other updates, the requesting Service Account or API Key must have the Project Cluster Manager role. You can't modify a paused cluster (`paused : true`). You must call this endpoint to set `paused : false`. After this endpoint responds with `paused : false`, you can call it again with the changes you want to make to the cluster. This feature is not available for serverless clusters. Deprecated versions: v2-{2024-08-05}, v2-{2023-02-01}, v2-{2023-01-01}

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster.
	@return UpdateClusterApiRequest
*/
func (a *ClustersApiService) UpdateCluster(ctx context.Context, groupId string, clusterName string, clusterDescription20240805 *ClusterDescription20240805) UpdateClusterApiRequest {
	return UpdateClusterApiRequest{
		ApiService:                 a,
		ctx:                        ctx,
		groupId:                    groupId,
		clusterName:                clusterName,
		clusterDescription20240805: clusterDescription20240805,
	}
}

// UpdateClusterExecute executes the request
//
//	@return ClusterDescription20240805
func (a *ClustersApiService) UpdateClusterExecute(r UpdateClusterApiRequest) (*ClusterDescription20240805, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ClusterDescription20240805
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClustersApiService.UpdateCluster")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return localVarReturnValue, nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.clusterDescription20240805 == nil {
		return localVarReturnValue, nil, reportError("clusterDescription20240805 is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2024-10-23+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-10-23+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.clusterDescription20240805
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		defer localVarHTTPResponse.Body.Close()
		buf, readErr := io.ReadAll(localVarHTTPResponse.Body)
		if readErr != nil {
			err = readErr
		}
		newErr := &GenericOpenAPIError{
			body:  buf,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type UpdateClusterAdvancedConfigurationApiRequest struct {
	ctx                                   context.Context
	ApiService                            ClustersApi
	groupId                               string
	clusterName                           string
	clusterDescriptionProcessArgs20240805 *ClusterDescriptionProcessArgs20240805
}

type UpdateClusterAdvancedConfigurationApiParams struct {
	GroupId                               string
	ClusterName                           string
	ClusterDescriptionProcessArgs20240805 *ClusterDescriptionProcessArgs20240805
}

func (a *ClustersApiService) UpdateClusterAdvancedConfigurationWithParams(ctx context.Context, args *UpdateClusterAdvancedConfigurationApiParams) UpdateClusterAdvancedConfigurationApiRequest {
	return UpdateClusterAdvancedConfigurationApiRequest{
		ApiService:                            a,
		ctx:                                   ctx,
		groupId:                               args.GroupId,
		clusterName:                           args.ClusterName,
		clusterDescriptionProcessArgs20240805: args.ClusterDescriptionProcessArgs20240805,
	}
}

func (r UpdateClusterAdvancedConfigurationApiRequest) Execute() (*ClusterDescriptionProcessArgs20240805, *http.Response, error) {
	return r.ApiService.UpdateClusterAdvancedConfigurationExecute(r)
}

/*
UpdateClusterAdvancedConfiguration Update Advanced Configuration Options for One Cluster

Updates the advanced configuration details for one cluster in the specified project. Clusters contain a group of hosts that maintain the same data set. Advanced configuration details include the read/write concern, index and oplog limits, and other database settings. To use this resource, the requesting Service Account or API Key must have the Project Cluster Manager role. This feature isn't available for `M0` free clusters, `M2` and `M5` shared-tier clusters, flex clusters, or serverless clusters. Deprecated versions: v2-{2023-01-01}

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Human-readable label that identifies the cluster.
	@return UpdateClusterAdvancedConfigurationApiRequest
*/
func (a *ClustersApiService) UpdateClusterAdvancedConfiguration(ctx context.Context, groupId string, clusterName string, clusterDescriptionProcessArgs20240805 *ClusterDescriptionProcessArgs20240805) UpdateClusterAdvancedConfigurationApiRequest {
	return UpdateClusterAdvancedConfigurationApiRequest{
		ApiService:                            a,
		ctx:                                   ctx,
		groupId:                               groupId,
		clusterName:                           clusterName,
		clusterDescriptionProcessArgs20240805: clusterDescriptionProcessArgs20240805,
	}
}

// UpdateClusterAdvancedConfigurationExecute executes the request
//
//	@return ClusterDescriptionProcessArgs20240805
func (a *ClustersApiService) UpdateClusterAdvancedConfigurationExecute(r UpdateClusterAdvancedConfigurationApiRequest) (*ClusterDescriptionProcessArgs20240805, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ClusterDescriptionProcessArgs20240805
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClustersApiService.UpdateClusterAdvancedConfiguration")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/processArgs"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return localVarReturnValue, nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.clusterDescriptionProcessArgs20240805 == nil {
		return localVarReturnValue, nil, reportError("clusterDescriptionProcessArgs20240805 is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2024-08-05+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-08-05+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.clusterDescriptionProcessArgs20240805
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		defer localVarHTTPResponse.Body.Close()
		buf, readErr := io.ReadAll(localVarHTTPResponse.Body)
		if readErr != nil {
			err = readErr
		}
		newErr := &GenericOpenAPIError{
			body:  buf,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type UpgradeSharedClusterApiRequest struct {
	ctx                                    context.Context
	ApiService                             ClustersApi
	groupId                                string
	legacyAtlasTenantClusterUpgradeRequest *LegacyAtlasTenantClusterUpgradeRequest
}

type UpgradeSharedClusterApiParams struct {
	GroupId                                string
	LegacyAtlasTenantClusterUpgradeRequest *LegacyAtlasTenantClusterUpgradeRequest
}

func (a *ClustersApiService) UpgradeSharedClusterWithParams(ctx context.Context, args *UpgradeSharedClusterApiParams) UpgradeSharedClusterApiRequest {
	return UpgradeSharedClusterApiRequest{
		ApiService:                             a,
		ctx:                                    ctx,
		groupId:                                args.GroupId,
		legacyAtlasTenantClusterUpgradeRequest: args.LegacyAtlasTenantClusterUpgradeRequest,
	}
}

func (r UpgradeSharedClusterApiRequest) Execute() (*LegacyAtlasCluster, *http.Response, error) {
	return r.ApiService.UpgradeSharedClusterExecute(r)
}

/*
UpgradeSharedCluster Upgrade One Shared-Tier Cluster

Upgrades a shared-tier cluster to a Flex or Dedicated (M10+) cluster in the specified project. To use this resource, the requesting Service Account or API Key must have the Project Cluster Manager role. Each project supports up to 25 clusters.

This endpoint can also be used to upgrade Flex clusters that were created using the [createCluster](https://www.mongodb.com/docs/atlas/reference/api-resources-spec/v2/#tag/Clusters/operation/createCluster) API or former M2/M5 clusters that have been migrated to Flex clusters, using instanceSizeName to “M2” or “M5” until January 2026. This functionality will be available until January 2026, after which it will only be available for M0 clusters. Please use the upgradeFlexCluster endpoint instead.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return UpgradeSharedClusterApiRequest
*/
func (a *ClustersApiService) UpgradeSharedCluster(ctx context.Context, groupId string, legacyAtlasTenantClusterUpgradeRequest *LegacyAtlasTenantClusterUpgradeRequest) UpgradeSharedClusterApiRequest {
	return UpgradeSharedClusterApiRequest{
		ApiService:                             a,
		ctx:                                    ctx,
		groupId:                                groupId,
		legacyAtlasTenantClusterUpgradeRequest: legacyAtlasTenantClusterUpgradeRequest,
	}
}

// UpgradeSharedClusterExecute executes the request
//
//	@return LegacyAtlasCluster
func (a *ClustersApiService) UpgradeSharedClusterExecute(r UpgradeSharedClusterApiRequest) (*LegacyAtlasCluster, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *LegacyAtlasCluster
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClustersApiService.UpgradeSharedCluster")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/tenantUpgrade"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.legacyAtlasTenantClusterUpgradeRequest == nil {
		return localVarReturnValue, nil, reportError("legacyAtlasTenantClusterUpgradeRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2023-01-01+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.legacyAtlasTenantClusterUpgradeRequest
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		defer localVarHTTPResponse.Body.Close()
		buf, readErr := io.ReadAll(localVarHTTPResponse.Body)
		if readErr != nil {
			err = readErr
		}
		newErr := &GenericOpenAPIError{
			body:  buf,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type UpgradeSharedClusterToServerlessApiRequest struct {
	ctx                           context.Context
	ApiService                    ClustersApi
	groupId                       string
	serverlessInstanceDescription *ServerlessInstanceDescription
}

type UpgradeSharedClusterToServerlessApiParams struct {
	GroupId                       string
	ServerlessInstanceDescription *ServerlessInstanceDescription
}

func (a *ClustersApiService) UpgradeSharedClusterToServerlessWithParams(ctx context.Context, args *UpgradeSharedClusterToServerlessApiParams) UpgradeSharedClusterToServerlessApiRequest {
	return UpgradeSharedClusterToServerlessApiRequest{
		ApiService:                    a,
		ctx:                           ctx,
		groupId:                       args.GroupId,
		serverlessInstanceDescription: args.ServerlessInstanceDescription,
	}
}

func (r UpgradeSharedClusterToServerlessApiRequest) Execute() (*ServerlessInstanceDescription, *http.Response, error) {
	return r.ApiService.UpgradeSharedClusterToServerlessExecute(r)
}

/*
UpgradeSharedClusterToServerless Upgrade One Shared-Tier Cluster to One Serverless Instance

This endpoint has been deprecated as of February 2025 as we no longer support the creation of new serverless instances. Please use the upgradeFlexCluster endpoint to upgrade Flex clusters.

	Upgrades a shared-tier cluster to a serverless instance in the specified project. To use this resource, the requesting Service Account or API Key must have the Project Cluster Manager role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@return UpgradeSharedClusterToServerlessApiRequest

Deprecated
*/
func (a *ClustersApiService) UpgradeSharedClusterToServerless(ctx context.Context, groupId string, serverlessInstanceDescription *ServerlessInstanceDescription) UpgradeSharedClusterToServerlessApiRequest {
	return UpgradeSharedClusterToServerlessApiRequest{
		ApiService:                    a,
		ctx:                           ctx,
		groupId:                       groupId,
		serverlessInstanceDescription: serverlessInstanceDescription,
	}
}

// UpgradeSharedClusterToServerlessExecute executes the request
//
//	@return ServerlessInstanceDescription
//
// Deprecated
func (a *ClustersApiService) UpgradeSharedClusterToServerlessExecute(r UpgradeSharedClusterToServerlessApiRequest) (*ServerlessInstanceDescription, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ServerlessInstanceDescription
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClustersApiService.UpgradeSharedClusterToServerless")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/tenantUpgradeToServerless"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.serverlessInstanceDescription == nil {
		return localVarReturnValue, nil, reportError("serverlessInstanceDescription is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/vnd.atlas.2023-01-01+json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header (only first one)
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2023-01-01+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.serverlessInstanceDescription
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := a.client.makeApiError(localVarHTTPResponse, localVarHTTPMethod, localVarPath)
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarHTTPResponse.Body, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		defer localVarHTTPResponse.Body.Close()
		buf, readErr := io.ReadAll(localVarHTTPResponse.Body)
		if readErr != nil {
			err = readErr
		}
		newErr := &GenericOpenAPIError{
			body:  buf,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
