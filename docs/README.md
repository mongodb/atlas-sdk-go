# Atlas SDK GO documentation

* [Learn Concepts for the Atlas Go SDK](./doc_1_concepts.md#learn-concepts-for-the-atlas-go-sdk)
   * [Introduction](./doc_1_concepts.md#introduction)
   * [Fetching Data from the Back End](./doc_1_concepts.md#fetching-data-from-the-back-end)
   * [Performing Data Modification](./doc_1_concepts.md#performing-data-modification)
   * [Read Only and Write Only Fields](./doc_1_concepts.md#read-only-and-write-only-fields)
   * [Example](./doc_1_concepts.md#example)
   * [Release Strategy (Semantic Versioning)](./doc_1_concepts.md#release-strategy-semantic-versioning)
      * [Versioning Rules](./doc_1_concepts.md#versioning-rules)
         * [Major Version (vYYYYMMDDXXX.0.0)](./doc_1_concepts.md#major-version-vyyyymmddxxx00)
         * [Minor Version (vYYYYMMDDXXX.Y.0)](./doc_1_concepts.md#minor-version-vyyyymmddxxxy0)
         * [Patch Version (vYYYYMMDDXXX.Y.Z)](./doc_1_concepts.md#patch-version-vyyyymmddxxxyz)
   * [Example Version: v20230201001.0.0](./doc_1_concepts.md#example-version-v2023020100100)

* [Error Handling](./doc_2_error_handling.md#error-handling)
   * [Fetching Error Object](./doc_2_error_handling.md#fetching-error-object)
   * [Checking for the Existence of a Specific Error Code](./doc_2_error_handling.md#checking-for-the-existence-of-a-specific-error-code)
   * [Checking for the Existence of a Specific Response Status Code](./doc_2_error_handling.md#checking-for-the-existence-of-a-specific-response-status-code)
   * [Mocking Errors](./doc_2_error_handling.md#mocking-errors)

* [Migrate from the Go HTTP Client to the Atlas Go SDK](./doc_3_migration.md#migrate-from-the-go-http-client-to-the-atlas-go-sdk)
   * [Background](./doc_3_migration.md#background)
   * [Structural Changes](./doc_3_migration.md#structural-changes)
      * [Client Initialization](./doc_3_migration.md#client-initialization)
      * [Error Handling](./doc_3_migration.md#error-handling)
      * [Format of the API Interface](./doc_3_migration.md#format-of-the-api-interface)
      * [Different Naming Conventions for SDK Methods](./doc_3_migration.md#different-naming-conventions-for-sdk-methods)
      * [Multiple Choices when Creating Request Body Objects](./doc_3_migration.md#multiple-choices-when-creating-request-body-objects)

* [Authenticating with the Atlas Go SDK](./doc_4_authentication.md#authenticating-with-the-atlas-go-sdk)
   * [Using the Atlas Go SDK in Your Code with Digest Authentication](./doc_4_authentication.md#using-the-atlas-go-sdk-in-your-code-with-digest-authentication)
   * [(Preview) Using the Atlas Go SDK with Service Account Authentication](./doc_4_authentication.md#preview-using-the-atlas-go-sdk-with-service-account-authentication)
   * [OAuth Authentication](./doc_4_authentication.md#oauth-authentication)
      * [Authenticating with OAuth ClientID and ClientSecret](./doc_4_authentication.md#authenticating-with-oauth-clientid-and-clientsecret)
      * [Admin API Authentication using Service Accounts](./doc_4_authentication.md#admin-api-authentication-using-service-accounts)
      * [Revocation](./doc_4_authentication.md#revocation)
      * [OAuth Token Cache](./doc_4_authentication.md#oauth-token-cache)

* [Best Practices](./doc_5_best_practices.md#best-practices)
   * [Using Getters Instead of Direct Field Access](./doc_5_best_practices.md#using-getters-instead-of-direct-field-access)
   * [Check for Empty Strings for String Pointers](./doc_5_best_practices.md#check-for-empty-strings-for-string-pointers)
   * [Working with Date Fields](./doc_5_best_practices.md#working-with-date-fields)
   * [Working with Objects](./doc_5_best_practices.md#working-with-objects)
   * [Working with Pointers](./doc_5_best_practices.md#working-with-pointers)
   * [Working with Arrays](./doc_5_best_practices.md#working-with-arrays)
   * [Working with Binary Responses](./doc_5_best_practices.md#working-with-binary-responses)
   * [Use Method for Creating Models](./doc_5_best_practices.md#use-method-for-creating-models)
   * [Use golangci-lint Validators](./doc_5_best_practices.md#use-golangci-lint-validators)
      * [Linting Issues](./doc_5_best_practices.md#linting-issues)

* [SDK Reference](./doc_last_reference.md#sdk-reference)
   * [Reference Documentation for SDK Endpoints](./doc_last_reference.md#reference-documentation-for-sdk-endpoints)
   * [Documentation For Models](./doc_last_reference.md#documentation-for-models)

<!-- Created by https://github.com/ekalinin/github-markdown-toc -->
