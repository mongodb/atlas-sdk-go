# Atlas SDK GO documentation

* [Learn Concepts for the Atlas Go SDK](./doc_1_concepts.md#learn-concepts-for-the-atlas-go-sdk)
   * [Introduction](./doc_1_concepts.md#introduction)
   * [Fetching Data from the Back End](./doc_1_concepts.md#fetching-data-from-the-back-end)
   * [Performing Data Modification](./doc_1_concepts.md#performing-data-modification)
   * [Experimental Methods](./doc_1_concepts.md#experimental-methods)
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

* [Migrate from the Go HTTP Client to the Atlas Go SDK](./doc_3_migration.md#migrate-from-the-go-http-client-to-the-atlas-go-sdk)
   * [Background](./doc_3_migration.md#background)
   * [Structural Changes](./doc_3_migration.md#structural-changes)
      * [Client Initialization](./doc_3_migration.md#client-initialization)
      * [Error Handling](./doc_3_migration.md#error-handling)
      * [Format of the API Interface](./doc_3_migration.md#format-of-the-api-interface)
      * [Different Naming Conventions for SDK Methods](./doc_3_migration.md#different-naming-conventions-for-sdk-methods)
      * [Multiple Choices when Creating Request Body Objects](./doc_3_migration.md#multiple-choices-when-creating-request-body-objects)

* [Authenticate using the Atlas Go SDK](./doc_4_authentication.md#authenticate-using-the-atlas-go-sdk)
      * [Use the Atlas Go SDK in Your Code](./doc_4_authentication.md#use-the-atlas-go-sdk-in-your-code)

* [Best Practices](./doc_5_best_practices.md#best-practices)
   * [Using Getters Instead of Direct Field Access](./doc_5_best_practices.md#using-getters-instead-of-direct-field-access)
   * [Check for Empty Strings for String Pointers](./doc_5_best_practices.md#check-for-empty-strings-for-string-pointers)
   * [Use Method for Creating Models](./doc_5_best_practices.md#use-method-for-creating-models)
   * [Use golangci-lint Validators](./doc_5_best_practices.md#use-golangci-lint-validators)
      * [Linting Issues](./doc_5_best_practices.md#linting-issues)

* [SDK Reference](./doc_last_reference.md#sdk-reference)
   * [Reference Documentation for SDK Endpoints](./doc_last_reference.md#reference-documentation-for-sdk-endpoints)
   * [Documentation For Models](./doc_last_reference.md#documentation-for-models)

<!-- Created by https://github.com/ekalinin/github-markdown-toc -->
