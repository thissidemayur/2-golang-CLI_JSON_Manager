package commands

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/thissidemayur/cli-json-manager/internal/types"
)

// Manager struct to handle records
type Manager struct {
	fileName string
}

// NewManager creates a new Manager instance(constructor)
func NewManager(fileName string) *Manager {
	return &Manager{fileName: fileName}
}

// add a new record
func (m *Manager) AddRecord(name *string) {
	records, err := m.ReadRecord()
	if err != nil {
		fmt.Println("❌ Error reading records:", err)
		return
	}

	newRecord := types.Record{
		ID:   len(records),
		Name: *name,
	}

	records = append(records, newRecord)

	// marshal (convert go to json)
	m.SaveRecords(records)

	fmt.Println("✅ Record added successfully!")
}

// read all records
func (m* Manager) ListRecord(){
	 records, err := m.ReadRecord()
	 if err != nil {
		 fmt.Println("❌ Error reading records:", err)
		 return
	 }
	 if len(records) == 0 {
		 fmt.Println("ℹ️ No records found in ", m.fileName)
		 return
	 }

	 fmt.Println("📋 List of Records:")
	 for _, record := range records {
		 fmt.Printf("  - ID: %d, Name: %s\n", record.ID, record.Name)
	 }	
}

// delete  a record
func (m *Manager) DeleteRecord(id int){	
records, err := m.ReadRecord()
	if err != nil {
		fmt.Println("❌ Error reading records:", err)
		return
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
	m.SaveRecords(newRecord)

	fmt.Printf("🗑️  Deleted record with ID %d\n", id)
}

// update a record
func (m *Manager) UpdateRecord(id int, newName string) {
	records, err := m.ReadRecord()
	if err != nil {
		fmt.Println("❌ Error reading records:", err)
		return
	}
	if len(records) == 0 {
		fmt.Println(" 📂No records found.", m.fileName)
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
	m.SaveRecords(records)
	fmt.Printf("✏️  Updated record ID %d → New Name: %s\n", id, newName)

}

// ensure file exists
func (m *Manager) EnsureFile() error {
	_, err := os.Stat(m.fileName)
	if os.IsNotExist(err) {
		file, err := os.Create(m.fileName)
		if err != nil {
			return err
		}
		defer file.Close()
		file.Write([]byte("[]")) // initialize with empty JSON array	
	}
	return nil
}

// read records from file
func (m *Manager) ReadRecord() ([]types.Record , error) {
	if err := m.EnsureFile(); err != nil {
		return nil , err
	}

	// get file data	
	data, err := os.ReadFile(m.fileName)
	if err != nil {
		fmt.Println("❌ Error reading file:", err)
		return nil, err
	}

	var people []types.Record
	// unmarshal (convert json to go)
	if err := json.Unmarshal(data, &people); err != nil {
		fmt.Println("❌ Error unmarshalling data:", err)
		return nil, err
	}
	return people, nil

}


func (m *Manager) SaveRecords(records []types.Record) error {
	// marshal (convert go to json)
	data, err := json.MarshalIndent(records, "", " ")
	if err != nil {
		fmt.Println("❌ Error marshalling data:", err)
		return err
	}
	
	// write to json file
	if err := os.WriteFile(m.fileName, data, 0644); err != nil {
		fmt.Println("❌ Error writing file:", err)
		return err
	}
	return nil
}

