package main

import (
	"fmt"
	"log"
	"os"

	"context"

	"github.com/mongodb/atlas-sdk-go/admin"
	"github.com/mongodb/atlas-sdk-go/examples"
)

const (
	AdminDB             = "admin"
	ExternalAuthDB      = "$external"
	defaultResourceType = "CLUSTER"
)

/*
* MongoDB Atlas Go SDK Example for updating database users.
 */
func main() {
	ctx := context.Background()
	// Service Accounts are the recommended way to authenticate programmatically with the MongoDB Atlas API
	// See: https://www.mongodb.com/docs/cloud-manager/tutorial/manage-programmatic-api-keys/
	clientID := os.Getenv("MONGODB_ATLAS_CLIENT_ID")
	clientSecret := os.Getenv("MONGODB_ATLAS_CLIENT_SECRET")
	url := os.Getenv("MONGODB_ATLAS_BASE_URL")

	if clientID == "" || clientSecret == "" {
		log.Fatal("MONGODB_ATLAS_CLIENT_ID and MONGODB_ATLAS_CLIENT_SECRET must be set")
	}

	sdk, err := admin.NewClient(
		admin.UseBaseURL(url),
		admin.UseOAuthAuth(ctx, clientID, clientSecret),
		admin.UseDebug(true))
	examples.HandleErr(err, nil)

	current := new(admin.CloudDatabaseUser)
	update(current)

	_, response, err := sdk.ProjectsApi.GetGroup(ctx, current.GroupId).Execute()
	if err != nil {
		fmt.Println("Project missconfigured. Did you set the correct values in update() function?")
		examples.HandleErr(err, response)
	}

	params := &admin.UpdateDatabaseUserApiParams{
		GroupId:           current.GroupId,
		DatabaseName:      current.DatabaseName,
		Username:          current.Username,
		CloudDatabaseUser: current,
	}
	dbUser, response, err := sdk.DatabaseUsersApi.UpdateDatabaseUserWithParams(ctx, params).Execute()

	examples.HandleErr(err, response)
	fmt.Println(dbUser)
}

// Here we should provide input parameters for the dbUser update we wish to apply
func update(current *admin.CloudDatabaseUser) *admin.CloudDatabaseUser {
	// Note that this values require manual edits
	current.GroupId = "groupId"
	current.Username = "user"
	current.Password = admin.PtrString("password")
	current.Scopes = &[]admin.UserScope{}
	current.Roles = []admin.DatabaseUserRole{*admin.NewDatabaseUserRole(AdminDB, "readWrite")}
	current.DatabaseName = getAuthDB(current)
	return current
}

// GetAuthDB determines the authentication database based on the type of user.
// LDAP, X509 and AWSIAM should all use $external.
// SCRAM-SHA should use admin.
func getAuthDB(user *admin.CloudDatabaseUser) string {
	// base documentation https://registry.terraform.io/providers/mongodb/mongodbatlas/latest/docs/resources/database_user
	_, isX509 := adminX509Type[admin.GetOrDefault(user.X509Type, "")]
	_, isIAM := awsIAMType[admin.GetOrDefault(user.AwsIAMType, "")]

	// just USER is external
	isLDAP := user.LdapAuthType != nil && *user.LdapAuthType == "USER"

	if isX509 || isIAM || isLDAP {
		return ExternalAuthDB
	}
	return AdminDB
}

var adminX509Type = map[string]struct{}{
	"MANAGED":  {},
	"CUSTOMER": {},
}

var awsIAMType = map[string]struct{}{
	"USER": {},
	"ROLE": {},
}
