package main

import (
	"fmt"
	"os"

	"context"

	"go.mongodb.org/atlas-sdk/admin"
	"go.mongodb.org/atlas-sdk/examples"
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
	// Values provided as part of env variables
	// See: https://www.mongodb.com/docs/atlas/app-services/authentication/api-key/
	apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
	apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")
	url := os.Getenv("MONGODB_ATLAS_URL")

	sdk, err := admin.NewClient(
		admin.UseDigestAuth(apiKey, apiSecret),
		admin.UseBaseURL(url),
		admin.UseDebug(true))
	examples.HandleErr(err, nil)

	current := new(admin.CloudDatabaseUser)
	update(current)

	params := &admin.UpdateDatabaseUserApiParams{
		GroupId:      current.GroupId,
		DatabaseName: current.DatabaseName,
		Username:     current.Username,
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
	current.Scopes = make([]admin.UserScope, 0)
	current.Roles = make([]admin.DatabaseUserRole, 0)
	current.Roles = append(current.Roles, *admin.NewDatabaseUserRole(AdminDB, "readWrite"))
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
