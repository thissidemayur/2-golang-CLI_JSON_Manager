package commands

import (
	"fmt"

	"github.com/thissidemayur/cli-json-manager/internal/storage"
	"github.com/thissidemayur/cli-json-manager/internal/types"
)

// add a new record
func AddRecord(name *string) {
	records := storage.ReadRecord()
	newRecord := types.Record{
		ID:   len(records),
		Name: *name,
	}

	records = append(records, newRecord)

	// marshal (convert go to json)
	storage.SaveRecords(records)

	fmt.Println("✅ Record added successfully!")
}

// read all records
func ListRecord(){
	 records:= storage.ReadRecord()
	 if len(records) == 0{
		fmt.Println("ℹ️ No records found.")
		return
	 }

	 fmt.Println("📋 List of Records:")
	 for _, record := range records {
		 fmt.Printf("  - ID: %d, Name: %s\n", record.ID, record.Name)
	 }	
}

// delete  a record
func DeleteRecord(id int){	
records := storage.ReadRecord()
	if len(records) == 0 {
		fmt.Println("ℹ️ No records found to delete.")
		return;
	}

	found := false
	newRecord := [] types.Record{}
	for _,record := range records {
		if record.ID == id {
			found = true
			continue
		} else{
			newRecord = append(newRecord, record)
		}
	}

	if !found {
		fmt.Printf("⚠️  No record found with ID %d\n", id)
		return
	}

	// reassign IDs (so they remain sequential)
	for i := range newRecord {
		newRecord[i].ID = i + 1
	}

	// marshal (convert go to json)
	storage.SaveRecords(newRecord)

	fmt.Printf("🗑️  Deleted record with ID %d\n", id)
}

// update a record
func UpdateRecord(id int,newName string){
	records := storage.ReadRecord()
	if len(records) == 0 {
		fmt.Println("ℹ️  No records found to update.")
		return
	}
	found := false
	for i, r := range records {
		if r.ID == id {
			found = true
			records[i].Name = newName
			return
		}
	}

	if !found {
		fmt.Printf("⚠️  No record found with ID %d\n", id)
		return
	}
	
	// marshal (convert go to json)
	storage.SaveRecords(records)
	fmt.Printf("✏️  Updated record ID %d → New Name: %s\n", id, newName)

}

/*
	os.Args[0] -> programe name/path
	os.Args[1] -> first argument- the subcommand(add,list)
	os.Args[2] -> the rest of the flag or parameter (eg -name Mayur)
*/