package examples

import (
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/atlas-sdk/v20250312003/admin"
)

func HandleErr(err error, resp *http.Response) {
	if err == nil {
		return
	}

	if resp != nil {
		fmt.Println(resp.Body)
	}
	// Printing generic message
	fmt.Println(err.Error())
	apiErr, _ := admin.AsError(err)
	log.Fatalf("Error when performing SDK request: %v", apiErr.GetDetail())
}
