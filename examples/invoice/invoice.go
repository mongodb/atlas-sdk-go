package main

import (
	"fmt"
	"log"
	"os"

	"context"

	"go.mongodb.org/atlas-sdk/v20250312015/admin"
	"go.mongodb.org/atlas-sdk/v20250312015/examples"
)

func main() {
	ctx := context.Background()
	url := os.Getenv("MONGODB_ATLAS_BASE_URL")
	if url == "" {
		url = "https://cloud.mongodb.com"
	}
	apiKey := os.Getenv("MONGODB_ATLAS_PUBLIC_KEY")
	apiSecret := os.Getenv("MONGODB_ATLAS_PRIVATE_KEY")
	orgID := os.Getenv("MONGODB_ATLAS_ORG_ID")
	invoiceID := os.Getenv("MONGODB_ATLAS_INVOICE_ID")
	if apiKey == "" || apiSecret == "" || orgID == "" || invoiceID == "" {
		log.Fatal("Environment variables MONGODB_ATLAS_PUBLIC_KEY, MONGODB_ATLAS_PRIVATE_KEY, MONGODB_ATLAS_ORG_ID and MONGODB_ATLAS_INVOICE_ID are required")
	}

	sdk, err := admin.NewClient(
		admin.UseDigestAuth(apiKey, apiSecret),
		admin.UseBaseURL(url))
	examples.HandleErr(err, nil)

	invoice, response, err := sdk.InvoicesApi.GetInvoice(ctx, orgID, invoiceID).Execute()
	examples.HandleErr(err, response)

	fmt.Printf("Content-Type: %v\n", response.Header.Get("Content-Type"))
	fmt.Printf("Invoice object: %v\n", invoice)
	fmt.Printf("Org: %s, Created: %v, AmountBilledCents: %d\n", invoice.GetOrgId(), invoice.GetCreated(), invoice.GetAmountBilledCents())
}
