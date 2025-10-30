package storage

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/thissidemayur/cli-json-manager/internal/types"
)

const fileName = "data.json"


func ReadRecord() []types.Record {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return []types.Record{}
	}

	// get file data	
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println("❌ Error reading file:", err)
		return []types.Record{}
	}

	var records []types.Record
	// unmarshal (convert json to go)
	if err := json.Unmarshal(data, &records); err != nil {
		fmt.Println("❌ Error unmarshalling data:", err)
		return []types.Record{}
	}
	return records

}


func SaveRecords(records []types.Record) {
	// marshal (convert go to json)
	data, err := json.MarshalIndent(records," ","\n")
	if err != nil {
		fmt.Println("❌ Error marshalling data:", err)
		return
	}
	
	// write to json file
	if err := os.WriteFile(fileName, data, 0644); err != nil {
		fmt.Println("❌ Error writing file:", err)
		return
	}
}