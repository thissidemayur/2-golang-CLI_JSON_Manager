package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/thissidemayur/cli-json-manager/internal/types"
)

func readJSONFile(filePath string)([]types.Record , error){

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var records []types.Record
	err = json.Unmarshal(data, &records)
	if err != nil {
		return nil, err
	}
	
	return records, nil
}

func TestAddRecord(t *testing.T){
	// create a temprory json file
	tmpFile,err:=os.CreateTemp("", "test_data_*.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpFile.Name())// clean up after test(even test fails)
	defer tmpFile.Close()

	record := types.Record{
		ID: 0,
		Name: "Mayur",
	}
	// initialize Manager with temp file
	m := Manager{fileName: tmpFile.Name()}

	// add record
	err = m.AddRecord(&record.Name)
	if err != nil {
		t.Fatalf("AddRecord() failed: %v", err)
	}

	// verify record 
	data,err :=readJSONFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("Failed to read JSON file: %v", err)
	}

	if len(data) !=1 {
		t.Fatalf("Expected 1 record, got %d", len(data))
	}

	if data[0].Name != record.Name {
		t.Errorf("Expected Name %s, got %s", record.Name, data[0].Name)
	}

	fmt.Println("✅ TestAddRecord passed")

}

func TestListRecord(t *testing.T) {
	// create  tempraory file
	tmpFile,err:=os.CreateTemp("", "test_data_*.json");
	if err != nil {
		t.Fatalf("Failed to create temp file: %v",err)
	}
	defer tmpFile.Close()
	defer os.Remove(tmpFile.Name())

	// write mock data into the file
	data:=[]types.Record{
		{Name: "bhawaniputra",ID: 0},
		{Name:"taraPutra",ID: 1},
		{Name:"kaliPutra",ID: 2},
	}

	records,_ :=json.MarshalIndent(data,""," ")
	
	os.WriteFile(tmpFile.Name(),records,0644)

	// now matching data with our input
	m := Manager{fileName: tmpFile.Name()}
	data, err = m.ListRecord()
	if err != nil {
		t.Fatalf("ListRecords() failed: %v", err)
	}

	if len(data) != 3 {
		t.Errorf("Expected 3 records, got %d", len(data))
	}

	fmt.Println("✅ listRecord() successfully test.")
}

// Test for DeleteRecord
func TestDeleteRecord(t *testing.T) {
	// create temp file
	tmpFile, err := os.CreateTemp("", "test_data_*.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer tmpFile.Close()
	defer os.Remove(tmpFile.Name())

	// write mock data into the file
	data := []types.Record{
		{Name: "alpha", ID: 0},
		{Name: "beta", ID: 1},
		{Name: "gamma", ID: 2},
	}

	records, _ := json.MarshalIndent(data, "", " ")
	err = os.WriteFile(tmpFile.Name(), records, 0644)
	if err != nil {
		t.Fatalf("Failed to write JSON file: %v", err)
	}

	// delete record with ID 1
	m := Manager{fileName: tmpFile.Name()}
	err = m.DeleteRecord(1)
	if err != nil {
		t.Fatalf("DeleteRecord() failed: %v", err)
	}

	// verify record deletion
	remainingRecords, err := readJSONFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("Failed to read JSON file: %v", err)
	}

	if len(remainingRecords) != 2 {
		t.Errorf("Expected 2 records after deletion, got %d", len(remainingRecords))
	}

	for _,record := range remainingRecords{
		if record.ID == 0  && record.Name != "alpha"{
			t.Errorf("Expected Name 'alpha' for ID 0, got %s", record.Name)
	}
		if record.ID == 2 && record.Name != "gamma"{
			t.Errorf("Expected Name 'gamma' for ID 2, got %s", record.Name)
		}
	}

	fmt.Println("✅ DeleteRecord() successfully test.")
}

// test for update record
func TestUpdateRecord(t *testing.T) {
	// create temp file
	tmpFile, err := os.CreateTemp("", "test_data_*.json")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer tmpFile.Close()
	defer os.Remove(tmpFile.Name())

	// write mock data into the file
	data := []types.Record{
		{Name: "delta", ID: 0},
		{Name: "epsilon", ID: 1},
	}

	records, _ := json.MarshalIndent(data, "", " ")
	err = os.WriteFile(tmpFile.Name(), records, 0644)
	if err != nil {
		t.Fatalf("Failed to write JSON file: %v", err)
	}
	// update record with id 1
	m:= Manager{fileName: tmpFile.Name()}
	err = m.UpdateRecord(1, "zeta")
	if err != nil {
		t.Fatalf("UpdateRecord() failed: %v", err)
	}

	// verify update record
	updatedRecords, err := readJSONFile(tmpFile.Name())
	if err != nil {
		t.Fatalf("Failed to read JSON file: %v", err)
	}

	if len(updatedRecords) != 2 {
		t.Errorf("Expected 2 records after update, got %d", len(updatedRecords))
	}

	for _, record := range updatedRecords {
		if record.ID == 1 && record.Name != "zeta" {
			t.Errorf("Expected Name 'zeta' for ID 1, got %s", record.Name)
		}
	}

	fmt.Println("✅ UpdateRecord() successfully test.")
}