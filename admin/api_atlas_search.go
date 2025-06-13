// Code based on the AtlasAPI V2 OpenAPI file

package admin

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type AtlasSearchApi interface {

	/*
		CreateAtlasSearchDeployment Create Search Nodes

		Creates Search Nodes for the specified cluster.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Label that identifies the cluster to create Search Nodes for.
		@param apiSearchDeploymentRequest Creates Search Nodes for the specified cluster.
		@return CreateAtlasSearchDeploymentApiRequest
	*/
	CreateAtlasSearchDeployment(ctx context.Context, groupId string, clusterName string, apiSearchDeploymentRequest *ApiSearchDeploymentRequest) CreateAtlasSearchDeploymentApiRequest
	/*
		CreateAtlasSearchDeployment Create Search Nodes


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateAtlasSearchDeploymentApiParams - Parameters for the request
		@return CreateAtlasSearchDeploymentApiRequest
	*/
	CreateAtlasSearchDeploymentWithParams(ctx context.Context, args *CreateAtlasSearchDeploymentApiParams) CreateAtlasSearchDeploymentApiRequest

	// Method available only for mocking purposes
	CreateAtlasSearchDeploymentExecute(r CreateAtlasSearchDeploymentApiRequest) (*ApiSearchDeploymentResponse, *http.Response, error)

	/*
		CreateAtlasSearchIndex Create One Atlas Search Index

		Creates one Atlas Search index on the specified collection. Atlas Search indexes define the fields on which to create the index and the analyzers to use when creating the index. Only clusters running MongoDB v4.2 or later can use Atlas Search. To use this resource, the requesting Service Account or API Key must have the Project Data Access Admin role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Name of the cluster that contains the collection on which to create an Atlas Search index.
		@param searchIndexCreateRequest Creates one Atlas Search index on the specified collection.
		@return CreateAtlasSearchIndexApiRequest
	*/
	CreateAtlasSearchIndex(ctx context.Context, groupId string, clusterName string, searchIndexCreateRequest *SearchIndexCreateRequest) CreateAtlasSearchIndexApiRequest
	/*
		CreateAtlasSearchIndex Create One Atlas Search Index


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateAtlasSearchIndexApiParams - Parameters for the request
		@return CreateAtlasSearchIndexApiRequest
	*/
	CreateAtlasSearchIndexWithParams(ctx context.Context, args *CreateAtlasSearchIndexApiParams) CreateAtlasSearchIndexApiRequest

	// Method available only for mocking purposes
	CreateAtlasSearchIndexExecute(r CreateAtlasSearchIndexApiRequest) (*SearchIndexResponse, *http.Response, error)

	/*
		CreateAtlasSearchIndexDeprecated Create One Atlas Search Index

		Creates one Atlas Search index on the specified collection. Atlas Search indexes define the fields on which to create the index and the analyzers to use when creating the index. Only clusters running MongoDB v4.2 or later can use Atlas Search. To use this resource, the requesting Service Account or API Key must have the Project Data Access Admin role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Name of the cluster that contains the collection on which to create an Atlas Search index.
		@param clusterSearchIndex Creates one Atlas Search index on the specified collection.
		@return CreateAtlasSearchIndexDeprecatedApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for AtlasSearchApi
	*/
	CreateAtlasSearchIndexDeprecated(ctx context.Context, groupId string, clusterName string, clusterSearchIndex *ClusterSearchIndex) CreateAtlasSearchIndexDeprecatedApiRequest
	/*
		CreateAtlasSearchIndexDeprecated Create One Atlas Search Index


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param CreateAtlasSearchIndexDeprecatedApiParams - Parameters for the request
		@return CreateAtlasSearchIndexDeprecatedApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for AtlasSearchApi
	*/
	CreateAtlasSearchIndexDeprecatedWithParams(ctx context.Context, args *CreateAtlasSearchIndexDeprecatedApiParams) CreateAtlasSearchIndexDeprecatedApiRequest

	// Method available only for mocking purposes
	CreateAtlasSearchIndexDeprecatedExecute(r CreateAtlasSearchIndexDeprecatedApiRequest) (*ClusterSearchIndex, *http.Response, error)

	/*
		DeleteAtlasSearchDeployment Delete Search Nodes

		Deletes the Search Nodes for the specified cluster.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Label that identifies the cluster to delete.
		@return DeleteAtlasSearchDeploymentApiRequest
	*/
	DeleteAtlasSearchDeployment(ctx context.Context, groupId string, clusterName string) DeleteAtlasSearchDeploymentApiRequest
	/*
		DeleteAtlasSearchDeployment Delete Search Nodes


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteAtlasSearchDeploymentApiParams - Parameters for the request
		@return DeleteAtlasSearchDeploymentApiRequest
	*/
	DeleteAtlasSearchDeploymentWithParams(ctx context.Context, args *DeleteAtlasSearchDeploymentApiParams) DeleteAtlasSearchDeploymentApiRequest

	// Method available only for mocking purposes
	DeleteAtlasSearchDeploymentExecute(r DeleteAtlasSearchDeploymentApiRequest) (*http.Response, error)

	/*
		DeleteAtlasSearchIndex Remove One Atlas Search Index by ID

		Removes one Atlas Search index that you identified with its unique ID. To use this resource, the requesting Service Account or API Key must have the Project Data Access Admin role. This deletion is eventually consistent.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Name of the cluster that contains the database and collection with one or more Application Search indexes.
		@param indexId Unique 24-hexadecimal digit string that identifies the Atlas Search index. Use the [Get All Atlas Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Atlas Search indexes.
		@return DeleteAtlasSearchIndexApiRequest
	*/
	DeleteAtlasSearchIndex(ctx context.Context, groupId string, clusterName string, indexId string) DeleteAtlasSearchIndexApiRequest
	/*
		DeleteAtlasSearchIndex Remove One Atlas Search Index by ID


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteAtlasSearchIndexApiParams - Parameters for the request
		@return DeleteAtlasSearchIndexApiRequest
	*/
	DeleteAtlasSearchIndexWithParams(ctx context.Context, args *DeleteAtlasSearchIndexApiParams) DeleteAtlasSearchIndexApiRequest

	// Method available only for mocking purposes
	DeleteAtlasSearchIndexExecute(r DeleteAtlasSearchIndexApiRequest) (*http.Response, error)

	/*
		DeleteAtlasSearchIndexByName Remove One Atlas Search Index by Name

		Removes one Atlas Search index that you identified with its database, collection, and name. To use this resource, the requesting Service Account or API Key must have the Project Data Access Admin role. This deletion is eventually consistent.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Name of the cluster that contains the database and collection with one or more Application Search indexes.
		@param collectionName Name of the collection that contains one or more Atlas Search indexes.
		@param databaseName Label that identifies the database that contains the collection with one or more Atlas Search indexes.
		@param indexName Name of the Atlas Search index to delete.
		@return DeleteAtlasSearchIndexByNameApiRequest
	*/
	DeleteAtlasSearchIndexByName(ctx context.Context, groupId string, clusterName string, collectionName string, databaseName string, indexName string) DeleteAtlasSearchIndexByNameApiRequest
	/*
		DeleteAtlasSearchIndexByName Remove One Atlas Search Index by Name


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteAtlasSearchIndexByNameApiParams - Parameters for the request
		@return DeleteAtlasSearchIndexByNameApiRequest
	*/
	DeleteAtlasSearchIndexByNameWithParams(ctx context.Context, args *DeleteAtlasSearchIndexByNameApiParams) DeleteAtlasSearchIndexByNameApiRequest

	// Method available only for mocking purposes
	DeleteAtlasSearchIndexByNameExecute(r DeleteAtlasSearchIndexByNameApiRequest) (*http.Response, error)

	/*
		DeleteAtlasSearchIndexDeprecated Remove One Atlas Search Index

		Removes one Atlas Search index that you identified with its unique ID. To use this resource, the requesting Service Account or API Key must have the Project Data Access Admin role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Name of the cluster that contains the database and collection with one or more Application Search indexes.
		@param indexId Unique 24-hexadecimal digit string that identifies the Atlas Search index. Use the [Get All Atlas Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Atlas Search indexes.
		@return DeleteAtlasSearchIndexDeprecatedApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for AtlasSearchApi
	*/
	DeleteAtlasSearchIndexDeprecated(ctx context.Context, groupId string, clusterName string, indexId string) DeleteAtlasSearchIndexDeprecatedApiRequest
	/*
		DeleteAtlasSearchIndexDeprecated Remove One Atlas Search Index


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param DeleteAtlasSearchIndexDeprecatedApiParams - Parameters for the request
		@return DeleteAtlasSearchIndexDeprecatedApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for AtlasSearchApi
	*/
	DeleteAtlasSearchIndexDeprecatedWithParams(ctx context.Context, args *DeleteAtlasSearchIndexDeprecatedApiParams) DeleteAtlasSearchIndexDeprecatedApiRequest

	// Method available only for mocking purposes
	DeleteAtlasSearchIndexDeprecatedExecute(r DeleteAtlasSearchIndexDeprecatedApiRequest) (*http.Response, error)

	/*
		GetAtlasSearchDeployment Return All Search Nodes

		Returns the Search Nodes for the specified cluster. Deprecated versions: v2-{2024-05-30}, v2-{2023-01-01}

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Label that identifies the cluster to return the Search Nodes for.
		@return GetAtlasSearchDeploymentApiRequest
	*/
	GetAtlasSearchDeployment(ctx context.Context, groupId string, clusterName string) GetAtlasSearchDeploymentApiRequest
	/*
		GetAtlasSearchDeployment Return All Search Nodes


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetAtlasSearchDeploymentApiParams - Parameters for the request
		@return GetAtlasSearchDeploymentApiRequest
	*/
	GetAtlasSearchDeploymentWithParams(ctx context.Context, args *GetAtlasSearchDeploymentApiParams) GetAtlasSearchDeploymentApiRequest

	// Method available only for mocking purposes
	GetAtlasSearchDeploymentExecute(r GetAtlasSearchDeploymentApiRequest) (*ApiSearchDeploymentResponse, *http.Response, error)

	/*
		GetAtlasSearchIndex Return One Atlas Search Index by ID

		Returns one Atlas Search index in the specified project. You identify this index using its unique ID. Atlas Search index contains the indexed fields and the analyzers used to create the index. To use this resource, the requesting Service Account or API Key must have the Project Data Access Read Write role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Name of the cluster that contains the collection with one or more Atlas Search indexes.
		@param indexId Unique 24-hexadecimal digit string that identifies the Application Search [index](https://dochub.mongodb.org/core/index-definitions-fts). Use the [Get All Application Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Application Search indexes.
		@return GetAtlasSearchIndexApiRequest
	*/
	GetAtlasSearchIndex(ctx context.Context, groupId string, clusterName string, indexId string) GetAtlasSearchIndexApiRequest
	/*
		GetAtlasSearchIndex Return One Atlas Search Index by ID


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetAtlasSearchIndexApiParams - Parameters for the request
		@return GetAtlasSearchIndexApiRequest
	*/
	GetAtlasSearchIndexWithParams(ctx context.Context, args *GetAtlasSearchIndexApiParams) GetAtlasSearchIndexApiRequest

	// Method available only for mocking purposes
	GetAtlasSearchIndexExecute(r GetAtlasSearchIndexApiRequest) (*SearchIndexResponse, *http.Response, error)

	/*
		GetAtlasSearchIndexByName Return One Atlas Search Index by Name

		Returns one Atlas Search index in the specified project. You identify this index using its database, collection and name. Atlas Search index contains the indexed fields and the analyzers used to create the index. To use this resource, the requesting Service Account or API Key must have the Project Data Access Read Write role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Name of the cluster that contains the collection with one or more Atlas Search indexes.
		@param collectionName Name of the collection that contains one or more Atlas Search indexes.
		@param databaseName Label that identifies the database that contains the collection with one or more Atlas Search indexes.
		@param indexName Name of the Atlas Search index to return.
		@return GetAtlasSearchIndexByNameApiRequest
	*/
	GetAtlasSearchIndexByName(ctx context.Context, groupId string, clusterName string, collectionName string, databaseName string, indexName string) GetAtlasSearchIndexByNameApiRequest
	/*
		GetAtlasSearchIndexByName Return One Atlas Search Index by Name


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetAtlasSearchIndexByNameApiParams - Parameters for the request
		@return GetAtlasSearchIndexByNameApiRequest
	*/
	GetAtlasSearchIndexByNameWithParams(ctx context.Context, args *GetAtlasSearchIndexByNameApiParams) GetAtlasSearchIndexByNameApiRequest

	// Method available only for mocking purposes
	GetAtlasSearchIndexByNameExecute(r GetAtlasSearchIndexByNameApiRequest) (*SearchIndexResponse, *http.Response, error)

	/*
		GetAtlasSearchIndexDeprecated Return One Atlas Search Index

		Returns one Atlas Search index in the specified project. You identify this index using its unique ID. Atlas Search index contains the indexed fields and the analyzers used to create the index. To use this resource, the requesting Service Account or API Key must have the Project Data Access Read Write role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Name of the cluster that contains the collection with one or more Atlas Search indexes.
		@param indexId Unique 24-hexadecimal digit string that identifies the Application Search [index](https://dochub.mongodb.org/core/index-definitions-fts). Use the [Get All Application Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Application Search indexes.
		@return GetAtlasSearchIndexDeprecatedApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for AtlasSearchApi
	*/
	GetAtlasSearchIndexDeprecated(ctx context.Context, groupId string, clusterName string, indexId string) GetAtlasSearchIndexDeprecatedApiRequest
	/*
		GetAtlasSearchIndexDeprecated Return One Atlas Search Index


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param GetAtlasSearchIndexDeprecatedApiParams - Parameters for the request
		@return GetAtlasSearchIndexDeprecatedApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for AtlasSearchApi
	*/
	GetAtlasSearchIndexDeprecatedWithParams(ctx context.Context, args *GetAtlasSearchIndexDeprecatedApiParams) GetAtlasSearchIndexDeprecatedApiRequest

	// Method available only for mocking purposes
	GetAtlasSearchIndexDeprecatedExecute(r GetAtlasSearchIndexDeprecatedApiRequest) (*ClusterSearchIndex, *http.Response, error)

	/*
		ListAtlasSearchIndexes Return All Atlas Search Indexes for One Collection

		Returns all Atlas Search indexes on the specified collection. Atlas Search indexes contain the indexed fields and the analyzers used to create the indexes. To use this resource, the requesting Service Account or API Key must have the Project Data Access Read Write role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Name of the cluster that contains the collection with one or more Atlas Search indexes.
		@param collectionName Name of the collection that contains one or more Atlas Search indexes.
		@param databaseName Label that identifies the database that contains the collection with one or more Atlas Search indexes.
		@return ListAtlasSearchIndexesApiRequest
	*/
	ListAtlasSearchIndexes(ctx context.Context, groupId string, clusterName string, collectionName string, databaseName string) ListAtlasSearchIndexesApiRequest
	/*
		ListAtlasSearchIndexes Return All Atlas Search Indexes for One Collection


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListAtlasSearchIndexesApiParams - Parameters for the request
		@return ListAtlasSearchIndexesApiRequest
	*/
	ListAtlasSearchIndexesWithParams(ctx context.Context, args *ListAtlasSearchIndexesApiParams) ListAtlasSearchIndexesApiRequest

	// Method available only for mocking purposes
	ListAtlasSearchIndexesExecute(r ListAtlasSearchIndexesApiRequest) ([]SearchIndexResponse, *http.Response, error)

	/*
		ListAtlasSearchIndexesCluster Return All Atlas Search Indexes for One Cluster

		Returns all Atlas Search indexes on the specified cluster. Atlas Search indexes contain the indexed fields and the analyzers used to create the indexes. To use this resource, the requesting Service Account or API Key must have the Project Data Access Read Write role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Name of the cluster that contains the collection with one or more Atlas Search indexes.
		@return ListAtlasSearchIndexesClusterApiRequest
	*/
	ListAtlasSearchIndexesCluster(ctx context.Context, groupId string, clusterName string) ListAtlasSearchIndexesClusterApiRequest
	/*
		ListAtlasSearchIndexesCluster Return All Atlas Search Indexes for One Cluster


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListAtlasSearchIndexesClusterApiParams - Parameters for the request
		@return ListAtlasSearchIndexesClusterApiRequest
	*/
	ListAtlasSearchIndexesClusterWithParams(ctx context.Context, args *ListAtlasSearchIndexesClusterApiParams) ListAtlasSearchIndexesClusterApiRequest

	// Method available only for mocking purposes
	ListAtlasSearchIndexesClusterExecute(r ListAtlasSearchIndexesClusterApiRequest) ([]SearchIndexResponse, *http.Response, error)

	/*
		ListAtlasSearchIndexesDeprecated Return All Atlas Search Indexes for One Collection

		Returns all Atlas Search indexes on the specified collection. Atlas Search indexes contain the indexed fields and the analyzers used to create the indexes. To use this resource, the requesting Service Account or API Key must have the Project Data Access Read Write role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Name of the cluster that contains the collection with one or more Atlas Search indexes.
		@param collectionName Name of the collection that contains one or more Atlas Search indexes.
		@param databaseName Human-readable label that identifies the database that contains the collection with one or more Atlas Search indexes.
		@return ListAtlasSearchIndexesDeprecatedApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for AtlasSearchApi
	*/
	ListAtlasSearchIndexesDeprecated(ctx context.Context, groupId string, clusterName string, collectionName string, databaseName string) ListAtlasSearchIndexesDeprecatedApiRequest
	/*
		ListAtlasSearchIndexesDeprecated Return All Atlas Search Indexes for One Collection


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param ListAtlasSearchIndexesDeprecatedApiParams - Parameters for the request
		@return ListAtlasSearchIndexesDeprecatedApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for AtlasSearchApi
	*/
	ListAtlasSearchIndexesDeprecatedWithParams(ctx context.Context, args *ListAtlasSearchIndexesDeprecatedApiParams) ListAtlasSearchIndexesDeprecatedApiRequest

	// Method available only for mocking purposes
	ListAtlasSearchIndexesDeprecatedExecute(r ListAtlasSearchIndexesDeprecatedApiRequest) ([]ClusterSearchIndex, *http.Response, error)

	/*
		UpdateAtlasSearchDeployment Update Search Nodes

		Updates the Search Nodes for the specified cluster. Deprecated versions: v2-{2023-01-01}

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Label that identifies the cluster to update the Search Nodes for.
		@param apiSearchDeploymentRequest Updates the Search Nodes for the specified cluster.
		@return UpdateAtlasSearchDeploymentApiRequest
	*/
	UpdateAtlasSearchDeployment(ctx context.Context, groupId string, clusterName string, apiSearchDeploymentRequest *ApiSearchDeploymentRequest) UpdateAtlasSearchDeploymentApiRequest
	/*
		UpdateAtlasSearchDeployment Update Search Nodes


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateAtlasSearchDeploymentApiParams - Parameters for the request
		@return UpdateAtlasSearchDeploymentApiRequest
	*/
	UpdateAtlasSearchDeploymentWithParams(ctx context.Context, args *UpdateAtlasSearchDeploymentApiParams) UpdateAtlasSearchDeploymentApiRequest

	// Method available only for mocking purposes
	UpdateAtlasSearchDeploymentExecute(r UpdateAtlasSearchDeploymentApiRequest) (*ApiSearchDeploymentResponse, *http.Response, error)

	/*
		UpdateAtlasSearchIndex Update One Atlas Search Index by ID

		Updates one Atlas Search index that you identified with its unique ID. Atlas Search indexes define the fields on which to create the index and the analyzers to use when creating the index. To use this resource, the requesting Service Account or API Key must have the Project Data Access Admin role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Name of the cluster that contains the collection whose Atlas Search index you want to update.
		@param indexId Unique 24-hexadecimal digit string that identifies the Atlas Search [index](https://dochub.mongodb.org/core/index-definitions-fts). Use the [Get All Atlas Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Atlas Search indexes.
		@param searchIndexUpdateRequest Details to update on the Atlas Search index.
		@return UpdateAtlasSearchIndexApiRequest
	*/
	UpdateAtlasSearchIndex(ctx context.Context, groupId string, clusterName string, indexId string, searchIndexUpdateRequest *SearchIndexUpdateRequest) UpdateAtlasSearchIndexApiRequest
	/*
		UpdateAtlasSearchIndex Update One Atlas Search Index by ID


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateAtlasSearchIndexApiParams - Parameters for the request
		@return UpdateAtlasSearchIndexApiRequest
	*/
	UpdateAtlasSearchIndexWithParams(ctx context.Context, args *UpdateAtlasSearchIndexApiParams) UpdateAtlasSearchIndexApiRequest

	// Method available only for mocking purposes
	UpdateAtlasSearchIndexExecute(r UpdateAtlasSearchIndexApiRequest) (*SearchIndexResponse, *http.Response, error)

	/*
		UpdateAtlasSearchIndexByName Update One Atlas Search Index by Name

		Updates one Atlas Search index that you identified with its database, collection name, and index name. Atlas Search indexes define the fields on which to create the index and the analyzers to use when creating the index. To use this resource, the requesting Service Account or API Key must have the Project Data Access Admin role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Name of the cluster that contains the collection whose Atlas Search index you want to update.
		@param collectionName Name of the collection that contains one or more Atlas Search indexes.
		@param databaseName Label that identifies the database that contains the collection with one or more Atlas Search indexes.
		@param indexName Name of the Atlas Search index to update.
		@param searchIndexUpdateRequest Details to update the Atlas Search index with.
		@return UpdateAtlasSearchIndexByNameApiRequest
	*/
	UpdateAtlasSearchIndexByName(ctx context.Context, groupId string, clusterName string, collectionName string, databaseName string, indexName string, searchIndexUpdateRequest *SearchIndexUpdateRequest) UpdateAtlasSearchIndexByNameApiRequest
	/*
		UpdateAtlasSearchIndexByName Update One Atlas Search Index by Name


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateAtlasSearchIndexByNameApiParams - Parameters for the request
		@return UpdateAtlasSearchIndexByNameApiRequest
	*/
	UpdateAtlasSearchIndexByNameWithParams(ctx context.Context, args *UpdateAtlasSearchIndexByNameApiParams) UpdateAtlasSearchIndexByNameApiRequest

	// Method available only for mocking purposes
	UpdateAtlasSearchIndexByNameExecute(r UpdateAtlasSearchIndexByNameApiRequest) (*SearchIndexResponse, *http.Response, error)

	/*
		UpdateAtlasSearchIndexDeprecated Update One Atlas Search Index

		Updates one Atlas Search index that you identified with its unique ID. Atlas Search indexes define the fields on which to create the index and the analyzers to use when creating the index. To use this resource, the requesting Service Account or API Key must have the Project Data Access Admin role.

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
		@param clusterName Name of the cluster that contains the collection whose Atlas Search index to update.
		@param indexId Unique 24-hexadecimal digit string that identifies the Atlas Search [index](https://dochub.mongodb.org/core/index-definitions-fts). Use the [Get All Atlas Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Atlas Search indexes.
		@param clusterSearchIndex Details to update on the Atlas Search index.
		@return UpdateAtlasSearchIndexDeprecatedApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for AtlasSearchApi
	*/
	UpdateAtlasSearchIndexDeprecated(ctx context.Context, groupId string, clusterName string, indexId string, clusterSearchIndex *ClusterSearchIndex) UpdateAtlasSearchIndexDeprecatedApiRequest
	/*
		UpdateAtlasSearchIndexDeprecated Update One Atlas Search Index


		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@param UpdateAtlasSearchIndexDeprecatedApiParams - Parameters for the request
		@return UpdateAtlasSearchIndexDeprecatedApiRequest

		Deprecated: this method has been deprecated. Please check the latest resource version for AtlasSearchApi
	*/
	UpdateAtlasSearchIndexDeprecatedWithParams(ctx context.Context, args *UpdateAtlasSearchIndexDeprecatedApiParams) UpdateAtlasSearchIndexDeprecatedApiRequest

	// Method available only for mocking purposes
	UpdateAtlasSearchIndexDeprecatedExecute(r UpdateAtlasSearchIndexDeprecatedApiRequest) (*ClusterSearchIndex, *http.Response, error)
}

// AtlasSearchApiService AtlasSearchApi service
type AtlasSearchApiService service

type CreateAtlasSearchDeploymentApiRequest struct {
	ctx                        context.Context
	ApiService                 AtlasSearchApi
	groupId                    string
	clusterName                string
	apiSearchDeploymentRequest *ApiSearchDeploymentRequest
}

type CreateAtlasSearchDeploymentApiParams struct {
	GroupId                    string
	ClusterName                string
	ApiSearchDeploymentRequest *ApiSearchDeploymentRequest
}

func (a *AtlasSearchApiService) CreateAtlasSearchDeploymentWithParams(ctx context.Context, args *CreateAtlasSearchDeploymentApiParams) CreateAtlasSearchDeploymentApiRequest {
	return CreateAtlasSearchDeploymentApiRequest{
		ApiService:                 a,
		ctx:                        ctx,
		groupId:                    args.GroupId,
		clusterName:                args.ClusterName,
		apiSearchDeploymentRequest: args.ApiSearchDeploymentRequest,
	}
}

func (r CreateAtlasSearchDeploymentApiRequest) Execute() (*ApiSearchDeploymentResponse, *http.Response, error) {
	return r.ApiService.CreateAtlasSearchDeploymentExecute(r)
}

/*
CreateAtlasSearchDeployment Create Search Nodes

Creates Search Nodes for the specified cluster.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Label that identifies the cluster to create Search Nodes for.
	@return CreateAtlasSearchDeploymentApiRequest
*/
func (a *AtlasSearchApiService) CreateAtlasSearchDeployment(ctx context.Context, groupId string, clusterName string, apiSearchDeploymentRequest *ApiSearchDeploymentRequest) CreateAtlasSearchDeploymentApiRequest {
	return CreateAtlasSearchDeploymentApiRequest{
		ApiService:                 a,
		ctx:                        ctx,
		groupId:                    groupId,
		clusterName:                clusterName,
		apiSearchDeploymentRequest: apiSearchDeploymentRequest,
	}
}

// CreateAtlasSearchDeploymentExecute executes the request
//
//	@return ApiSearchDeploymentResponse
func (a *AtlasSearchApiService) CreateAtlasSearchDeploymentExecute(r CreateAtlasSearchDeploymentApiRequest) (*ApiSearchDeploymentResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ApiSearchDeploymentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AtlasSearchApiService.CreateAtlasSearchDeployment")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/deployment"
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
	if r.apiSearchDeploymentRequest == nil {
		return localVarReturnValue, nil, reportError("apiSearchDeploymentRequest is required and must be specified")
	}

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
	localVarPostBody = r.apiSearchDeploymentRequest
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

type CreateAtlasSearchIndexApiRequest struct {
	ctx                      context.Context
	ApiService               AtlasSearchApi
	groupId                  string
	clusterName              string
	searchIndexCreateRequest *SearchIndexCreateRequest
}

type CreateAtlasSearchIndexApiParams struct {
	GroupId                  string
	ClusterName              string
	SearchIndexCreateRequest *SearchIndexCreateRequest
}

func (a *AtlasSearchApiService) CreateAtlasSearchIndexWithParams(ctx context.Context, args *CreateAtlasSearchIndexApiParams) CreateAtlasSearchIndexApiRequest {
	return CreateAtlasSearchIndexApiRequest{
		ApiService:               a,
		ctx:                      ctx,
		groupId:                  args.GroupId,
		clusterName:              args.ClusterName,
		searchIndexCreateRequest: args.SearchIndexCreateRequest,
	}
}

func (r CreateAtlasSearchIndexApiRequest) Execute() (*SearchIndexResponse, *http.Response, error) {
	return r.ApiService.CreateAtlasSearchIndexExecute(r)
}

/*
CreateAtlasSearchIndex Create One Atlas Search Index

Creates one Atlas Search index on the specified collection. Atlas Search indexes define the fields on which to create the index and the analyzers to use when creating the index. Only clusters running MongoDB v4.2 or later can use Atlas Search. To use this resource, the requesting Service Account or API Key must have the Project Data Access Admin role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Name of the cluster that contains the collection on which to create an Atlas Search index.
	@return CreateAtlasSearchIndexApiRequest
*/
func (a *AtlasSearchApiService) CreateAtlasSearchIndex(ctx context.Context, groupId string, clusterName string, searchIndexCreateRequest *SearchIndexCreateRequest) CreateAtlasSearchIndexApiRequest {
	return CreateAtlasSearchIndexApiRequest{
		ApiService:               a,
		ctx:                      ctx,
		groupId:                  groupId,
		clusterName:              clusterName,
		searchIndexCreateRequest: searchIndexCreateRequest,
	}
}

// CreateAtlasSearchIndexExecute executes the request
//
//	@return SearchIndexResponse
func (a *AtlasSearchApiService) CreateAtlasSearchIndexExecute(r CreateAtlasSearchIndexApiRequest) (*SearchIndexResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *SearchIndexResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AtlasSearchApiService.CreateAtlasSearchIndex")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/indexes"
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
	if r.searchIndexCreateRequest == nil {
		return localVarReturnValue, nil, reportError("searchIndexCreateRequest is required and must be specified")
	}

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
	localVarPostBody = r.searchIndexCreateRequest
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

type CreateAtlasSearchIndexDeprecatedApiRequest struct {
	ctx                context.Context
	ApiService         AtlasSearchApi
	groupId            string
	clusterName        string
	clusterSearchIndex *ClusterSearchIndex
}

type CreateAtlasSearchIndexDeprecatedApiParams struct {
	GroupId            string
	ClusterName        string
	ClusterSearchIndex *ClusterSearchIndex
}

func (a *AtlasSearchApiService) CreateAtlasSearchIndexDeprecatedWithParams(ctx context.Context, args *CreateAtlasSearchIndexDeprecatedApiParams) CreateAtlasSearchIndexDeprecatedApiRequest {
	return CreateAtlasSearchIndexDeprecatedApiRequest{
		ApiService:         a,
		ctx:                ctx,
		groupId:            args.GroupId,
		clusterName:        args.ClusterName,
		clusterSearchIndex: args.ClusterSearchIndex,
	}
}

func (r CreateAtlasSearchIndexDeprecatedApiRequest) Execute() (*ClusterSearchIndex, *http.Response, error) {
	return r.ApiService.CreateAtlasSearchIndexDeprecatedExecute(r)
}

/*
CreateAtlasSearchIndexDeprecated Create One Atlas Search Index

Creates one Atlas Search index on the specified collection. Atlas Search indexes define the fields on which to create the index and the analyzers to use when creating the index. Only clusters running MongoDB v4.2 or later can use Atlas Search. To use this resource, the requesting Service Account or API Key must have the Project Data Access Admin role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Name of the cluster that contains the collection on which to create an Atlas Search index.
	@return CreateAtlasSearchIndexDeprecatedApiRequest

Deprecated
*/
func (a *AtlasSearchApiService) CreateAtlasSearchIndexDeprecated(ctx context.Context, groupId string, clusterName string, clusterSearchIndex *ClusterSearchIndex) CreateAtlasSearchIndexDeprecatedApiRequest {
	return CreateAtlasSearchIndexDeprecatedApiRequest{
		ApiService:         a,
		ctx:                ctx,
		groupId:            groupId,
		clusterName:        clusterName,
		clusterSearchIndex: clusterSearchIndex,
	}
}

// CreateAtlasSearchIndexDeprecatedExecute executes the request
//
//	@return ClusterSearchIndex
//
// Deprecated
func (a *AtlasSearchApiService) CreateAtlasSearchIndexDeprecatedExecute(r CreateAtlasSearchIndexDeprecatedApiRequest) (*ClusterSearchIndex, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ClusterSearchIndex
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AtlasSearchApiService.CreateAtlasSearchIndexDeprecated")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes"
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
	if r.clusterSearchIndex == nil {
		return localVarReturnValue, nil, reportError("clusterSearchIndex is required and must be specified")
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
	localVarPostBody = r.clusterSearchIndex
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

type DeleteAtlasSearchDeploymentApiRequest struct {
	ctx         context.Context
	ApiService  AtlasSearchApi
	groupId     string
	clusterName string
}

type DeleteAtlasSearchDeploymentApiParams struct {
	GroupId     string
	ClusterName string
}

func (a *AtlasSearchApiService) DeleteAtlasSearchDeploymentWithParams(ctx context.Context, args *DeleteAtlasSearchDeploymentApiParams) DeleteAtlasSearchDeploymentApiRequest {
	return DeleteAtlasSearchDeploymentApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
	}
}

func (r DeleteAtlasSearchDeploymentApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteAtlasSearchDeploymentExecute(r)
}

/*
DeleteAtlasSearchDeployment Delete Search Nodes

Deletes the Search Nodes for the specified cluster.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Label that identifies the cluster to delete.
	@return DeleteAtlasSearchDeploymentApiRequest
*/
func (a *AtlasSearchApiService) DeleteAtlasSearchDeployment(ctx context.Context, groupId string, clusterName string) DeleteAtlasSearchDeploymentApiRequest {
	return DeleteAtlasSearchDeploymentApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
	}
}

// DeleteAtlasSearchDeploymentExecute executes the request
func (a *AtlasSearchApiService) DeleteAtlasSearchDeploymentExecute(r DeleteAtlasSearchDeploymentApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AtlasSearchApiService.DeleteAtlasSearchDeployment")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/deployment"
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

type DeleteAtlasSearchIndexApiRequest struct {
	ctx         context.Context
	ApiService  AtlasSearchApi
	groupId     string
	clusterName string
	indexId     string
}

type DeleteAtlasSearchIndexApiParams struct {
	GroupId     string
	ClusterName string
	IndexId     string
}

func (a *AtlasSearchApiService) DeleteAtlasSearchIndexWithParams(ctx context.Context, args *DeleteAtlasSearchIndexApiParams) DeleteAtlasSearchIndexApiRequest {
	return DeleteAtlasSearchIndexApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
		indexId:     args.IndexId,
	}
}

func (r DeleteAtlasSearchIndexApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteAtlasSearchIndexExecute(r)
}

/*
DeleteAtlasSearchIndex Remove One Atlas Search Index by ID

Removes one Atlas Search index that you identified with its unique ID. To use this resource, the requesting Service Account or API Key must have the Project Data Access Admin role. This deletion is eventually consistent.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Name of the cluster that contains the database and collection with one or more Application Search indexes.
	@param indexId Unique 24-hexadecimal digit string that identifies the Atlas Search index. Use the [Get All Atlas Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Atlas Search indexes.
	@return DeleteAtlasSearchIndexApiRequest
*/
func (a *AtlasSearchApiService) DeleteAtlasSearchIndex(ctx context.Context, groupId string, clusterName string, indexId string) DeleteAtlasSearchIndexApiRequest {
	return DeleteAtlasSearchIndexApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
		indexId:     indexId,
	}
}

// DeleteAtlasSearchIndexExecute executes the request
func (a *AtlasSearchApiService) DeleteAtlasSearchIndexExecute(r DeleteAtlasSearchIndexApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AtlasSearchApiService.DeleteAtlasSearchIndex")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/indexes/{indexId}"
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)
	if r.indexId == "" {
		return nil, reportError("indexId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"indexId"+"}", url.PathEscape(r.indexId), -1)

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

type DeleteAtlasSearchIndexByNameApiRequest struct {
	ctx            context.Context
	ApiService     AtlasSearchApi
	groupId        string
	clusterName    string
	collectionName string
	databaseName   string
	indexName      string
}

type DeleteAtlasSearchIndexByNameApiParams struct {
	GroupId        string
	ClusterName    string
	CollectionName string
	DatabaseName   string
	IndexName      string
}

func (a *AtlasSearchApiService) DeleteAtlasSearchIndexByNameWithParams(ctx context.Context, args *DeleteAtlasSearchIndexByNameApiParams) DeleteAtlasSearchIndexByNameApiRequest {
	return DeleteAtlasSearchIndexByNameApiRequest{
		ApiService:     a,
		ctx:            ctx,
		groupId:        args.GroupId,
		clusterName:    args.ClusterName,
		collectionName: args.CollectionName,
		databaseName:   args.DatabaseName,
		indexName:      args.IndexName,
	}
}

func (r DeleteAtlasSearchIndexByNameApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteAtlasSearchIndexByNameExecute(r)
}

/*
DeleteAtlasSearchIndexByName Remove One Atlas Search Index by Name

Removes one Atlas Search index that you identified with its database, collection, and name. To use this resource, the requesting Service Account or API Key must have the Project Data Access Admin role. This deletion is eventually consistent.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Name of the cluster that contains the database and collection with one or more Application Search indexes.
	@param collectionName Name of the collection that contains one or more Atlas Search indexes.
	@param databaseName Label that identifies the database that contains the collection with one or more Atlas Search indexes.
	@param indexName Name of the Atlas Search index to delete.
	@return DeleteAtlasSearchIndexByNameApiRequest
*/
func (a *AtlasSearchApiService) DeleteAtlasSearchIndexByName(ctx context.Context, groupId string, clusterName string, collectionName string, databaseName string, indexName string) DeleteAtlasSearchIndexByNameApiRequest {
	return DeleteAtlasSearchIndexByNameApiRequest{
		ApiService:     a,
		ctx:            ctx,
		groupId:        groupId,
		clusterName:    clusterName,
		collectionName: collectionName,
		databaseName:   databaseName,
		indexName:      indexName,
	}
}

// DeleteAtlasSearchIndexByNameExecute executes the request
func (a *AtlasSearchApiService) DeleteAtlasSearchIndexByNameExecute(r DeleteAtlasSearchIndexByNameApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AtlasSearchApiService.DeleteAtlasSearchIndexByName")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/indexes/{databaseName}/{collectionName}/{indexName}"
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)
	if r.collectionName == "" {
		return nil, reportError("collectionName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"collectionName"+"}", url.PathEscape(r.collectionName), -1)
	if r.databaseName == "" {
		return nil, reportError("databaseName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"databaseName"+"}", url.PathEscape(r.databaseName), -1)
	if r.indexName == "" {
		return nil, reportError("indexName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"indexName"+"}", url.PathEscape(r.indexName), -1)

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

type DeleteAtlasSearchIndexDeprecatedApiRequest struct {
	ctx         context.Context
	ApiService  AtlasSearchApi
	groupId     string
	clusterName string
	indexId     string
}

type DeleteAtlasSearchIndexDeprecatedApiParams struct {
	GroupId     string
	ClusterName string
	IndexId     string
}

func (a *AtlasSearchApiService) DeleteAtlasSearchIndexDeprecatedWithParams(ctx context.Context, args *DeleteAtlasSearchIndexDeprecatedApiParams) DeleteAtlasSearchIndexDeprecatedApiRequest {
	return DeleteAtlasSearchIndexDeprecatedApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
		indexId:     args.IndexId,
	}
}

func (r DeleteAtlasSearchIndexDeprecatedApiRequest) Execute() (*http.Response, error) {
	return r.ApiService.DeleteAtlasSearchIndexDeprecatedExecute(r)
}

/*
DeleteAtlasSearchIndexDeprecated Remove One Atlas Search Index

Removes one Atlas Search index that you identified with its unique ID. To use this resource, the requesting Service Account or API Key must have the Project Data Access Admin role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Name of the cluster that contains the database and collection with one or more Application Search indexes.
	@param indexId Unique 24-hexadecimal digit string that identifies the Atlas Search index. Use the [Get All Atlas Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Atlas Search indexes.
	@return DeleteAtlasSearchIndexDeprecatedApiRequest

Deprecated
*/
func (a *AtlasSearchApiService) DeleteAtlasSearchIndexDeprecated(ctx context.Context, groupId string, clusterName string, indexId string) DeleteAtlasSearchIndexDeprecatedApiRequest {
	return DeleteAtlasSearchIndexDeprecatedApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
		indexId:     indexId,
	}
}

// DeleteAtlasSearchIndexDeprecatedExecute executes the request
// Deprecated
func (a *AtlasSearchApiService) DeleteAtlasSearchIndexDeprecatedExecute(r DeleteAtlasSearchIndexDeprecatedApiRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod = http.MethodDelete
		localVarPostBody   any
		formFiles          []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AtlasSearchApiService.DeleteAtlasSearchIndexDeprecated")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes/{indexId}"
	if r.groupId == "" {
		return nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)
	if r.indexId == "" {
		return nil, reportError("indexId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"indexId"+"}", url.PathEscape(r.indexId), -1)

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

type GetAtlasSearchDeploymentApiRequest struct {
	ctx         context.Context
	ApiService  AtlasSearchApi
	groupId     string
	clusterName string
}

type GetAtlasSearchDeploymentApiParams struct {
	GroupId     string
	ClusterName string
}

func (a *AtlasSearchApiService) GetAtlasSearchDeploymentWithParams(ctx context.Context, args *GetAtlasSearchDeploymentApiParams) GetAtlasSearchDeploymentApiRequest {
	return GetAtlasSearchDeploymentApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
	}
}

func (r GetAtlasSearchDeploymentApiRequest) Execute() (*ApiSearchDeploymentResponse, *http.Response, error) {
	return r.ApiService.GetAtlasSearchDeploymentExecute(r)
}

/*
GetAtlasSearchDeployment Return All Search Nodes

Returns the Search Nodes for the specified cluster. Deprecated versions: v2-{2024-05-30}, v2-{2023-01-01}

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Label that identifies the cluster to return the Search Nodes for.
	@return GetAtlasSearchDeploymentApiRequest
*/
func (a *AtlasSearchApiService) GetAtlasSearchDeployment(ctx context.Context, groupId string, clusterName string) GetAtlasSearchDeploymentApiRequest {
	return GetAtlasSearchDeploymentApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
	}
}

// GetAtlasSearchDeploymentExecute executes the request
//
//	@return ApiSearchDeploymentResponse
func (a *AtlasSearchApiService) GetAtlasSearchDeploymentExecute(r GetAtlasSearchDeploymentApiRequest) (*ApiSearchDeploymentResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ApiSearchDeploymentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AtlasSearchApiService.GetAtlasSearchDeployment")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/deployment"
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2025-03-12+json"}

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

type GetAtlasSearchIndexApiRequest struct {
	ctx         context.Context
	ApiService  AtlasSearchApi
	groupId     string
	clusterName string
	indexId     string
}

type GetAtlasSearchIndexApiParams struct {
	GroupId     string
	ClusterName string
	IndexId     string
}

func (a *AtlasSearchApiService) GetAtlasSearchIndexWithParams(ctx context.Context, args *GetAtlasSearchIndexApiParams) GetAtlasSearchIndexApiRequest {
	return GetAtlasSearchIndexApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
		indexId:     args.IndexId,
	}
}

func (r GetAtlasSearchIndexApiRequest) Execute() (*SearchIndexResponse, *http.Response, error) {
	return r.ApiService.GetAtlasSearchIndexExecute(r)
}

/*
GetAtlasSearchIndex Return One Atlas Search Index by ID

Returns one Atlas Search index in the specified project. You identify this index using its unique ID. Atlas Search index contains the indexed fields and the analyzers used to create the index. To use this resource, the requesting Service Account or API Key must have the Project Data Access Read Write role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Name of the cluster that contains the collection with one or more Atlas Search indexes.
	@param indexId Unique 24-hexadecimal digit string that identifies the Application Search [index](https://dochub.mongodb.org/core/index-definitions-fts). Use the [Get All Application Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Application Search indexes.
	@return GetAtlasSearchIndexApiRequest
*/
func (a *AtlasSearchApiService) GetAtlasSearchIndex(ctx context.Context, groupId string, clusterName string, indexId string) GetAtlasSearchIndexApiRequest {
	return GetAtlasSearchIndexApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
		indexId:     indexId,
	}
}

// GetAtlasSearchIndexExecute executes the request
//
//	@return SearchIndexResponse
func (a *AtlasSearchApiService) GetAtlasSearchIndexExecute(r GetAtlasSearchIndexApiRequest) (*SearchIndexResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *SearchIndexResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AtlasSearchApiService.GetAtlasSearchIndex")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/indexes/{indexId}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return localVarReturnValue, nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)
	if r.indexId == "" {
		return localVarReturnValue, nil, reportError("indexId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"indexId"+"}", url.PathEscape(r.indexId), -1)

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

type GetAtlasSearchIndexByNameApiRequest struct {
	ctx            context.Context
	ApiService     AtlasSearchApi
	groupId        string
	clusterName    string
	collectionName string
	databaseName   string
	indexName      string
}

type GetAtlasSearchIndexByNameApiParams struct {
	GroupId        string
	ClusterName    string
	CollectionName string
	DatabaseName   string
	IndexName      string
}

func (a *AtlasSearchApiService) GetAtlasSearchIndexByNameWithParams(ctx context.Context, args *GetAtlasSearchIndexByNameApiParams) GetAtlasSearchIndexByNameApiRequest {
	return GetAtlasSearchIndexByNameApiRequest{
		ApiService:     a,
		ctx:            ctx,
		groupId:        args.GroupId,
		clusterName:    args.ClusterName,
		collectionName: args.CollectionName,
		databaseName:   args.DatabaseName,
		indexName:      args.IndexName,
	}
}

func (r GetAtlasSearchIndexByNameApiRequest) Execute() (*SearchIndexResponse, *http.Response, error) {
	return r.ApiService.GetAtlasSearchIndexByNameExecute(r)
}

/*
GetAtlasSearchIndexByName Return One Atlas Search Index by Name

Returns one Atlas Search index in the specified project. You identify this index using its database, collection and name. Atlas Search index contains the indexed fields and the analyzers used to create the index. To use this resource, the requesting Service Account or API Key must have the Project Data Access Read Write role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Name of the cluster that contains the collection with one or more Atlas Search indexes.
	@param collectionName Name of the collection that contains one or more Atlas Search indexes.
	@param databaseName Label that identifies the database that contains the collection with one or more Atlas Search indexes.
	@param indexName Name of the Atlas Search index to return.
	@return GetAtlasSearchIndexByNameApiRequest
*/
func (a *AtlasSearchApiService) GetAtlasSearchIndexByName(ctx context.Context, groupId string, clusterName string, collectionName string, databaseName string, indexName string) GetAtlasSearchIndexByNameApiRequest {
	return GetAtlasSearchIndexByNameApiRequest{
		ApiService:     a,
		ctx:            ctx,
		groupId:        groupId,
		clusterName:    clusterName,
		collectionName: collectionName,
		databaseName:   databaseName,
		indexName:      indexName,
	}
}

// GetAtlasSearchIndexByNameExecute executes the request
//
//	@return SearchIndexResponse
func (a *AtlasSearchApiService) GetAtlasSearchIndexByNameExecute(r GetAtlasSearchIndexByNameApiRequest) (*SearchIndexResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *SearchIndexResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AtlasSearchApiService.GetAtlasSearchIndexByName")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/indexes/{databaseName}/{collectionName}/{indexName}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return localVarReturnValue, nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)
	if r.collectionName == "" {
		return localVarReturnValue, nil, reportError("collectionName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"collectionName"+"}", url.PathEscape(r.collectionName), -1)
	if r.databaseName == "" {
		return localVarReturnValue, nil, reportError("databaseName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"databaseName"+"}", url.PathEscape(r.databaseName), -1)
	if r.indexName == "" {
		return localVarReturnValue, nil, reportError("indexName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"indexName"+"}", url.PathEscape(r.indexName), -1)

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

type GetAtlasSearchIndexDeprecatedApiRequest struct {
	ctx         context.Context
	ApiService  AtlasSearchApi
	groupId     string
	clusterName string
	indexId     string
}

type GetAtlasSearchIndexDeprecatedApiParams struct {
	GroupId     string
	ClusterName string
	IndexId     string
}

func (a *AtlasSearchApiService) GetAtlasSearchIndexDeprecatedWithParams(ctx context.Context, args *GetAtlasSearchIndexDeprecatedApiParams) GetAtlasSearchIndexDeprecatedApiRequest {
	return GetAtlasSearchIndexDeprecatedApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
		indexId:     args.IndexId,
	}
}

func (r GetAtlasSearchIndexDeprecatedApiRequest) Execute() (*ClusterSearchIndex, *http.Response, error) {
	return r.ApiService.GetAtlasSearchIndexDeprecatedExecute(r)
}

/*
GetAtlasSearchIndexDeprecated Return One Atlas Search Index

Returns one Atlas Search index in the specified project. You identify this index using its unique ID. Atlas Search index contains the indexed fields and the analyzers used to create the index. To use this resource, the requesting Service Account or API Key must have the Project Data Access Read Write role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Name of the cluster that contains the collection with one or more Atlas Search indexes.
	@param indexId Unique 24-hexadecimal digit string that identifies the Application Search [index](https://dochub.mongodb.org/core/index-definitions-fts). Use the [Get All Application Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Application Search indexes.
	@return GetAtlasSearchIndexDeprecatedApiRequest

Deprecated
*/
func (a *AtlasSearchApiService) GetAtlasSearchIndexDeprecated(ctx context.Context, groupId string, clusterName string, indexId string) GetAtlasSearchIndexDeprecatedApiRequest {
	return GetAtlasSearchIndexDeprecatedApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
		indexId:     indexId,
	}
}

// GetAtlasSearchIndexDeprecatedExecute executes the request
//
//	@return ClusterSearchIndex
//
// Deprecated
func (a *AtlasSearchApiService) GetAtlasSearchIndexDeprecatedExecute(r GetAtlasSearchIndexDeprecatedApiRequest) (*ClusterSearchIndex, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ClusterSearchIndex
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AtlasSearchApiService.GetAtlasSearchIndexDeprecated")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes/{indexId}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return localVarReturnValue, nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)
	if r.indexId == "" {
		return localVarReturnValue, nil, reportError("indexId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"indexId"+"}", url.PathEscape(r.indexId), -1)

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

type ListAtlasSearchIndexesApiRequest struct {
	ctx            context.Context
	ApiService     AtlasSearchApi
	groupId        string
	clusterName    string
	collectionName string
	databaseName   string
}

type ListAtlasSearchIndexesApiParams struct {
	GroupId        string
	ClusterName    string
	CollectionName string
	DatabaseName   string
}

func (a *AtlasSearchApiService) ListAtlasSearchIndexesWithParams(ctx context.Context, args *ListAtlasSearchIndexesApiParams) ListAtlasSearchIndexesApiRequest {
	return ListAtlasSearchIndexesApiRequest{
		ApiService:     a,
		ctx:            ctx,
		groupId:        args.GroupId,
		clusterName:    args.ClusterName,
		collectionName: args.CollectionName,
		databaseName:   args.DatabaseName,
	}
}

func (r ListAtlasSearchIndexesApiRequest) Execute() ([]SearchIndexResponse, *http.Response, error) {
	return r.ApiService.ListAtlasSearchIndexesExecute(r)
}

/*
ListAtlasSearchIndexes Return All Atlas Search Indexes for One Collection

Returns all Atlas Search indexes on the specified collection. Atlas Search indexes contain the indexed fields and the analyzers used to create the indexes. To use this resource, the requesting Service Account or API Key must have the Project Data Access Read Write role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Name of the cluster that contains the collection with one or more Atlas Search indexes.
	@param collectionName Name of the collection that contains one or more Atlas Search indexes.
	@param databaseName Label that identifies the database that contains the collection with one or more Atlas Search indexes.
	@return ListAtlasSearchIndexesApiRequest
*/
func (a *AtlasSearchApiService) ListAtlasSearchIndexes(ctx context.Context, groupId string, clusterName string, collectionName string, databaseName string) ListAtlasSearchIndexesApiRequest {
	return ListAtlasSearchIndexesApiRequest{
		ApiService:     a,
		ctx:            ctx,
		groupId:        groupId,
		clusterName:    clusterName,
		collectionName: collectionName,
		databaseName:   databaseName,
	}
}

// ListAtlasSearchIndexesExecute executes the request
//
//	@return []SearchIndexResponse
func (a *AtlasSearchApiService) ListAtlasSearchIndexesExecute(r ListAtlasSearchIndexesApiRequest) ([]SearchIndexResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue []SearchIndexResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AtlasSearchApiService.ListAtlasSearchIndexes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/indexes/{databaseName}/{collectionName}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return localVarReturnValue, nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)
	if r.collectionName == "" {
		return localVarReturnValue, nil, reportError("collectionName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"collectionName"+"}", url.PathEscape(r.collectionName), -1)
	if r.databaseName == "" {
		return localVarReturnValue, nil, reportError("databaseName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"databaseName"+"}", url.PathEscape(r.databaseName), -1)

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

type ListAtlasSearchIndexesClusterApiRequest struct {
	ctx         context.Context
	ApiService  AtlasSearchApi
	groupId     string
	clusterName string
}

type ListAtlasSearchIndexesClusterApiParams struct {
	GroupId     string
	ClusterName string
}

func (a *AtlasSearchApiService) ListAtlasSearchIndexesClusterWithParams(ctx context.Context, args *ListAtlasSearchIndexesClusterApiParams) ListAtlasSearchIndexesClusterApiRequest {
	return ListAtlasSearchIndexesClusterApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     args.GroupId,
		clusterName: args.ClusterName,
	}
}

func (r ListAtlasSearchIndexesClusterApiRequest) Execute() ([]SearchIndexResponse, *http.Response, error) {
	return r.ApiService.ListAtlasSearchIndexesClusterExecute(r)
}

/*
ListAtlasSearchIndexesCluster Return All Atlas Search Indexes for One Cluster

Returns all Atlas Search indexes on the specified cluster. Atlas Search indexes contain the indexed fields and the analyzers used to create the indexes. To use this resource, the requesting Service Account or API Key must have the Project Data Access Read Write role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Name of the cluster that contains the collection with one or more Atlas Search indexes.
	@return ListAtlasSearchIndexesClusterApiRequest
*/
func (a *AtlasSearchApiService) ListAtlasSearchIndexesCluster(ctx context.Context, groupId string, clusterName string) ListAtlasSearchIndexesClusterApiRequest {
	return ListAtlasSearchIndexesClusterApiRequest{
		ApiService:  a,
		ctx:         ctx,
		groupId:     groupId,
		clusterName: clusterName,
	}
}

// ListAtlasSearchIndexesClusterExecute executes the request
//
//	@return []SearchIndexResponse
func (a *AtlasSearchApiService) ListAtlasSearchIndexesClusterExecute(r ListAtlasSearchIndexesClusterApiRequest) ([]SearchIndexResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue []SearchIndexResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AtlasSearchApiService.ListAtlasSearchIndexesCluster")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/indexes"
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
	localVarHTTPHeaderAccepts := []string{"application/vnd.atlas.2024-05-30+json"}

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

type ListAtlasSearchIndexesDeprecatedApiRequest struct {
	ctx            context.Context
	ApiService     AtlasSearchApi
	groupId        string
	clusterName    string
	collectionName string
	databaseName   string
}

type ListAtlasSearchIndexesDeprecatedApiParams struct {
	GroupId        string
	ClusterName    string
	CollectionName string
	DatabaseName   string
}

func (a *AtlasSearchApiService) ListAtlasSearchIndexesDeprecatedWithParams(ctx context.Context, args *ListAtlasSearchIndexesDeprecatedApiParams) ListAtlasSearchIndexesDeprecatedApiRequest {
	return ListAtlasSearchIndexesDeprecatedApiRequest{
		ApiService:     a,
		ctx:            ctx,
		groupId:        args.GroupId,
		clusterName:    args.ClusterName,
		collectionName: args.CollectionName,
		databaseName:   args.DatabaseName,
	}
}

func (r ListAtlasSearchIndexesDeprecatedApiRequest) Execute() ([]ClusterSearchIndex, *http.Response, error) {
	return r.ApiService.ListAtlasSearchIndexesDeprecatedExecute(r)
}

/*
ListAtlasSearchIndexesDeprecated Return All Atlas Search Indexes for One Collection

Returns all Atlas Search indexes on the specified collection. Atlas Search indexes contain the indexed fields and the analyzers used to create the indexes. To use this resource, the requesting Service Account or API Key must have the Project Data Access Read Write role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Name of the cluster that contains the collection with one or more Atlas Search indexes.
	@param collectionName Name of the collection that contains one or more Atlas Search indexes.
	@param databaseName Human-readable label that identifies the database that contains the collection with one or more Atlas Search indexes.
	@return ListAtlasSearchIndexesDeprecatedApiRequest

Deprecated
*/
func (a *AtlasSearchApiService) ListAtlasSearchIndexesDeprecated(ctx context.Context, groupId string, clusterName string, collectionName string, databaseName string) ListAtlasSearchIndexesDeprecatedApiRequest {
	return ListAtlasSearchIndexesDeprecatedApiRequest{
		ApiService:     a,
		ctx:            ctx,
		groupId:        groupId,
		clusterName:    clusterName,
		collectionName: collectionName,
		databaseName:   databaseName,
	}
}

// ListAtlasSearchIndexesDeprecatedExecute executes the request
//
//	@return []ClusterSearchIndex
//
// Deprecated
func (a *AtlasSearchApiService) ListAtlasSearchIndexesDeprecatedExecute(r ListAtlasSearchIndexesDeprecatedApiRequest) ([]ClusterSearchIndex, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue []ClusterSearchIndex
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AtlasSearchApiService.ListAtlasSearchIndexesDeprecated")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes/{databaseName}/{collectionName}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return localVarReturnValue, nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)
	if r.collectionName == "" {
		return localVarReturnValue, nil, reportError("collectionName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"collectionName"+"}", url.PathEscape(r.collectionName), -1)
	if r.databaseName == "" {
		return localVarReturnValue, nil, reportError("databaseName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"databaseName"+"}", url.PathEscape(r.databaseName), -1)

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

type UpdateAtlasSearchDeploymentApiRequest struct {
	ctx                        context.Context
	ApiService                 AtlasSearchApi
	groupId                    string
	clusterName                string
	apiSearchDeploymentRequest *ApiSearchDeploymentRequest
}

type UpdateAtlasSearchDeploymentApiParams struct {
	GroupId                    string
	ClusterName                string
	ApiSearchDeploymentRequest *ApiSearchDeploymentRequest
}

func (a *AtlasSearchApiService) UpdateAtlasSearchDeploymentWithParams(ctx context.Context, args *UpdateAtlasSearchDeploymentApiParams) UpdateAtlasSearchDeploymentApiRequest {
	return UpdateAtlasSearchDeploymentApiRequest{
		ApiService:                 a,
		ctx:                        ctx,
		groupId:                    args.GroupId,
		clusterName:                args.ClusterName,
		apiSearchDeploymentRequest: args.ApiSearchDeploymentRequest,
	}
}

func (r UpdateAtlasSearchDeploymentApiRequest) Execute() (*ApiSearchDeploymentResponse, *http.Response, error) {
	return r.ApiService.UpdateAtlasSearchDeploymentExecute(r)
}

/*
UpdateAtlasSearchDeployment Update Search Nodes

Updates the Search Nodes for the specified cluster. Deprecated versions: v2-{2023-01-01}

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Label that identifies the cluster to update the Search Nodes for.
	@return UpdateAtlasSearchDeploymentApiRequest
*/
func (a *AtlasSearchApiService) UpdateAtlasSearchDeployment(ctx context.Context, groupId string, clusterName string, apiSearchDeploymentRequest *ApiSearchDeploymentRequest) UpdateAtlasSearchDeploymentApiRequest {
	return UpdateAtlasSearchDeploymentApiRequest{
		ApiService:                 a,
		ctx:                        ctx,
		groupId:                    groupId,
		clusterName:                clusterName,
		apiSearchDeploymentRequest: apiSearchDeploymentRequest,
	}
}

// UpdateAtlasSearchDeploymentExecute executes the request
//
//	@return ApiSearchDeploymentResponse
func (a *AtlasSearchApiService) UpdateAtlasSearchDeploymentExecute(r UpdateAtlasSearchDeploymentApiRequest) (*ApiSearchDeploymentResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ApiSearchDeploymentResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AtlasSearchApiService.UpdateAtlasSearchDeployment")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/deployment"
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
	if r.apiSearchDeploymentRequest == nil {
		return localVarReturnValue, nil, reportError("apiSearchDeploymentRequest is required and must be specified")
	}

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
	localVarPostBody = r.apiSearchDeploymentRequest
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

type UpdateAtlasSearchIndexApiRequest struct {
	ctx                      context.Context
	ApiService               AtlasSearchApi
	groupId                  string
	clusterName              string
	indexId                  string
	searchIndexUpdateRequest *SearchIndexUpdateRequest
}

type UpdateAtlasSearchIndexApiParams struct {
	GroupId                  string
	ClusterName              string
	IndexId                  string
	SearchIndexUpdateRequest *SearchIndexUpdateRequest
}

func (a *AtlasSearchApiService) UpdateAtlasSearchIndexWithParams(ctx context.Context, args *UpdateAtlasSearchIndexApiParams) UpdateAtlasSearchIndexApiRequest {
	return UpdateAtlasSearchIndexApiRequest{
		ApiService:               a,
		ctx:                      ctx,
		groupId:                  args.GroupId,
		clusterName:              args.ClusterName,
		indexId:                  args.IndexId,
		searchIndexUpdateRequest: args.SearchIndexUpdateRequest,
	}
}

func (r UpdateAtlasSearchIndexApiRequest) Execute() (*SearchIndexResponse, *http.Response, error) {
	return r.ApiService.UpdateAtlasSearchIndexExecute(r)
}

/*
UpdateAtlasSearchIndex Update One Atlas Search Index by ID

Updates one Atlas Search index that you identified with its unique ID. Atlas Search indexes define the fields on which to create the index and the analyzers to use when creating the index. To use this resource, the requesting Service Account or API Key must have the Project Data Access Admin role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Name of the cluster that contains the collection whose Atlas Search index you want to update.
	@param indexId Unique 24-hexadecimal digit string that identifies the Atlas Search [index](https://dochub.mongodb.org/core/index-definitions-fts). Use the [Get All Atlas Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Atlas Search indexes.
	@return UpdateAtlasSearchIndexApiRequest
*/
func (a *AtlasSearchApiService) UpdateAtlasSearchIndex(ctx context.Context, groupId string, clusterName string, indexId string, searchIndexUpdateRequest *SearchIndexUpdateRequest) UpdateAtlasSearchIndexApiRequest {
	return UpdateAtlasSearchIndexApiRequest{
		ApiService:               a,
		ctx:                      ctx,
		groupId:                  groupId,
		clusterName:              clusterName,
		indexId:                  indexId,
		searchIndexUpdateRequest: searchIndexUpdateRequest,
	}
}

// UpdateAtlasSearchIndexExecute executes the request
//
//	@return SearchIndexResponse
func (a *AtlasSearchApiService) UpdateAtlasSearchIndexExecute(r UpdateAtlasSearchIndexApiRequest) (*SearchIndexResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *SearchIndexResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AtlasSearchApiService.UpdateAtlasSearchIndex")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/indexes/{indexId}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return localVarReturnValue, nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)
	if r.indexId == "" {
		return localVarReturnValue, nil, reportError("indexId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"indexId"+"}", url.PathEscape(r.indexId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.searchIndexUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("searchIndexUpdateRequest is required and must be specified")
	}

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
	localVarPostBody = r.searchIndexUpdateRequest
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

type UpdateAtlasSearchIndexByNameApiRequest struct {
	ctx                      context.Context
	ApiService               AtlasSearchApi
	groupId                  string
	clusterName              string
	collectionName           string
	databaseName             string
	indexName                string
	searchIndexUpdateRequest *SearchIndexUpdateRequest
}

type UpdateAtlasSearchIndexByNameApiParams struct {
	GroupId                  string
	ClusterName              string
	CollectionName           string
	DatabaseName             string
	IndexName                string
	SearchIndexUpdateRequest *SearchIndexUpdateRequest
}

func (a *AtlasSearchApiService) UpdateAtlasSearchIndexByNameWithParams(ctx context.Context, args *UpdateAtlasSearchIndexByNameApiParams) UpdateAtlasSearchIndexByNameApiRequest {
	return UpdateAtlasSearchIndexByNameApiRequest{
		ApiService:               a,
		ctx:                      ctx,
		groupId:                  args.GroupId,
		clusterName:              args.ClusterName,
		collectionName:           args.CollectionName,
		databaseName:             args.DatabaseName,
		indexName:                args.IndexName,
		searchIndexUpdateRequest: args.SearchIndexUpdateRequest,
	}
}

func (r UpdateAtlasSearchIndexByNameApiRequest) Execute() (*SearchIndexResponse, *http.Response, error) {
	return r.ApiService.UpdateAtlasSearchIndexByNameExecute(r)
}

/*
UpdateAtlasSearchIndexByName Update One Atlas Search Index by Name

Updates one Atlas Search index that you identified with its database, collection name, and index name. Atlas Search indexes define the fields on which to create the index and the analyzers to use when creating the index. To use this resource, the requesting Service Account or API Key must have the Project Data Access Admin role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Name of the cluster that contains the collection whose Atlas Search index you want to update.
	@param collectionName Name of the collection that contains one or more Atlas Search indexes.
	@param databaseName Label that identifies the database that contains the collection with one or more Atlas Search indexes.
	@param indexName Name of the Atlas Search index to update.
	@return UpdateAtlasSearchIndexByNameApiRequest
*/
func (a *AtlasSearchApiService) UpdateAtlasSearchIndexByName(ctx context.Context, groupId string, clusterName string, collectionName string, databaseName string, indexName string, searchIndexUpdateRequest *SearchIndexUpdateRequest) UpdateAtlasSearchIndexByNameApiRequest {
	return UpdateAtlasSearchIndexByNameApiRequest{
		ApiService:               a,
		ctx:                      ctx,
		groupId:                  groupId,
		clusterName:              clusterName,
		collectionName:           collectionName,
		databaseName:             databaseName,
		indexName:                indexName,
		searchIndexUpdateRequest: searchIndexUpdateRequest,
	}
}

// UpdateAtlasSearchIndexByNameExecute executes the request
//
//	@return SearchIndexResponse
func (a *AtlasSearchApiService) UpdateAtlasSearchIndexByNameExecute(r UpdateAtlasSearchIndexByNameApiRequest) (*SearchIndexResponse, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *SearchIndexResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AtlasSearchApiService.UpdateAtlasSearchIndexByName")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/search/indexes/{databaseName}/{collectionName}/{indexName}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return localVarReturnValue, nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)
	if r.collectionName == "" {
		return localVarReturnValue, nil, reportError("collectionName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"collectionName"+"}", url.PathEscape(r.collectionName), -1)
	if r.databaseName == "" {
		return localVarReturnValue, nil, reportError("databaseName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"databaseName"+"}", url.PathEscape(r.databaseName), -1)
	if r.indexName == "" {
		return localVarReturnValue, nil, reportError("indexName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"indexName"+"}", url.PathEscape(r.indexName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.searchIndexUpdateRequest == nil {
		return localVarReturnValue, nil, reportError("searchIndexUpdateRequest is required and must be specified")
	}

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
	localVarPostBody = r.searchIndexUpdateRequest
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

type UpdateAtlasSearchIndexDeprecatedApiRequest struct {
	ctx                context.Context
	ApiService         AtlasSearchApi
	groupId            string
	clusterName        string
	indexId            string
	clusterSearchIndex *ClusterSearchIndex
}

type UpdateAtlasSearchIndexDeprecatedApiParams struct {
	GroupId            string
	ClusterName        string
	IndexId            string
	ClusterSearchIndex *ClusterSearchIndex
}

func (a *AtlasSearchApiService) UpdateAtlasSearchIndexDeprecatedWithParams(ctx context.Context, args *UpdateAtlasSearchIndexDeprecatedApiParams) UpdateAtlasSearchIndexDeprecatedApiRequest {
	return UpdateAtlasSearchIndexDeprecatedApiRequest{
		ApiService:         a,
		ctx:                ctx,
		groupId:            args.GroupId,
		clusterName:        args.ClusterName,
		indexId:            args.IndexId,
		clusterSearchIndex: args.ClusterSearchIndex,
	}
}

func (r UpdateAtlasSearchIndexDeprecatedApiRequest) Execute() (*ClusterSearchIndex, *http.Response, error) {
	return r.ApiService.UpdateAtlasSearchIndexDeprecatedExecute(r)
}

/*
UpdateAtlasSearchIndexDeprecated Update One Atlas Search Index

Updates one Atlas Search index that you identified with its unique ID. Atlas Search indexes define the fields on which to create the index and the analyzers to use when creating the index. To use this resource, the requesting Service Account or API Key must have the Project Data Access Admin role.

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param groupId Unique 24-hexadecimal digit string that identifies your project. Use the [/groups](#tag/Projects/operation/listProjects) endpoint to retrieve all projects to which the authenticated user has access.  **NOTE**: Groups and projects are synonymous terms. Your group id is the same as your project id. For existing groups, your group/project id remains the same. The resource and corresponding endpoints use the term groups.
	@param clusterName Name of the cluster that contains the collection whose Atlas Search index to update.
	@param indexId Unique 24-hexadecimal digit string that identifies the Atlas Search [index](https://dochub.mongodb.org/core/index-definitions-fts). Use the [Get All Atlas Search Indexes for a Collection API](https://docs.atlas.mongodb.com/reference/api/fts-indexes-get-all/) endpoint to find the IDs of all Atlas Search indexes.
	@return UpdateAtlasSearchIndexDeprecatedApiRequest

Deprecated
*/
func (a *AtlasSearchApiService) UpdateAtlasSearchIndexDeprecated(ctx context.Context, groupId string, clusterName string, indexId string, clusterSearchIndex *ClusterSearchIndex) UpdateAtlasSearchIndexDeprecatedApiRequest {
	return UpdateAtlasSearchIndexDeprecatedApiRequest{
		ApiService:         a,
		ctx:                ctx,
		groupId:            groupId,
		clusterName:        clusterName,
		indexId:            indexId,
		clusterSearchIndex: clusterSearchIndex,
	}
}

// UpdateAtlasSearchIndexDeprecatedExecute executes the request
//
//	@return ClusterSearchIndex
//
// Deprecated
func (a *AtlasSearchApiService) UpdateAtlasSearchIndexDeprecatedExecute(r UpdateAtlasSearchIndexDeprecatedApiRequest) (*ClusterSearchIndex, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPatch
		localVarPostBody    any
		formFiles           []formFile
		localVarReturnValue *ClusterSearchIndex
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AtlasSearchApiService.UpdateAtlasSearchIndexDeprecated")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/atlas/v2/groups/{groupId}/clusters/{clusterName}/fts/indexes/{indexId}"
	if r.groupId == "" {
		return localVarReturnValue, nil, reportError("groupId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"groupId"+"}", url.PathEscape(r.groupId), -1)
	if r.clusterName == "" {
		return localVarReturnValue, nil, reportError("clusterName is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"clusterName"+"}", url.PathEscape(r.clusterName), -1)
	if r.indexId == "" {
		return localVarReturnValue, nil, reportError("indexId is empty and must be specified")
	}
	localVarPath = strings.Replace(localVarPath, "{"+"indexId"+"}", url.PathEscape(r.indexId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.clusterSearchIndex == nil {
		return localVarReturnValue, nil, reportError("clusterSearchIndex is required and must be specified")
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
	localVarPostBody = r.clusterSearchIndex
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
