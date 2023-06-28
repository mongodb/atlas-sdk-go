package test

import (
	"fmt"

	"go.mongodb.org/atlas-sdk/v20230201001/admin"
)

func ExampleNewClient() {
	apiKey := "test"
	apiSecret := "test"
	host := "https://cloud.mongodb.com"

	sdk, err := admin.NewClient(
		admin.UseDigestAuth(apiKey, apiSecret),
		admin.UseBaseURL(host),
		admin.UseDebug(false))

	if err != nil {
		fmt.Println(err)
	}

	serverURL := sdk.GetConfig().Servers[0].URL
	fmt.Println(serverURL)
	// Output:
	// https://cloud.mongodb.com
}
