package main

import (
	"encoding/json"
	"fmt"

	"go.mongodb.org/atlas-sdk/v20231115008/admin"
)

func main() {
	var num int = 0
	myStruct := admin.DiskBackupSnapshotSchedule{
		ReferenceHourOfDay:&num,
	}

	// Print the struct before encoding to JSON
	fmt.Println("Before encoding:", myStruct)

	// Encode struct to JSON
	encoded, err := json.Marshal(myStruct)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	// Print the encoded JSON
	fmt.Println("Encoded JSON:", string(encoded))
}
