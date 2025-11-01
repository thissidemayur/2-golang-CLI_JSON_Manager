package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/thissidemayur/cli-json-manager/internal/types"
)

// sentimental error (resuable constant)
var (
	ErrNotFound = errors.New("record not found")
	ErrInvalidId = errors.New("invalid ID")
	ErrFileNotFound = errors.New("file not found")
	ErrEmptyName = errors.New("name cannot be empty")
)
// Manager struct to handle records
type Manager struct {
	fileName string
}

// NewManager creates a new Manager instance(constructor)
func NewManager(fileName string) (*Manager ,error ){
	m:= &Manager{fileName: fileName}
	if err := m.EnsureFile(); err != nil {
		return nil, fmt.Errorf("ensure file failed: %w", err)
	}
	return m, nil
}

//  ========================== CRUD on File  ==========================

// add a new record
func (m *Manager) AddRecord(name *string) error {
	records, err := m.ReadRecord()
	if err != nil {
		return fmt.Errorf("❌ Error reading records: %w", err)
	}

	newRecord := types.Record{
		ID:   len(records),
		Name: *name,
	}

	records = append(records, newRecord)

	// marshal (convert go to json)
	if err := m.SaveRecords(records); err != nil {
		return err
	}

	fmt.Printf("✅ Added Record: \"Id: %d, Name: %q\" successfully!\n", newRecord.ID, newRecord.Name)
	return nil
}

// read all records
func (m *Manager) ListRecord() error {
	records, err := m.ReadRecord()
	 if err != nil {
		 return err
	 }
	 if len(records) == 0 {
		 fmt.Printf("ℹ️ No records found in %q\n	", m.fileName)
		 return nil
	 }

	 fmt.Println("📋 List of Records:")
	 for _, record := range records {
		 fmt.Printf("  - ID: %d, Name: %s\n", record.ID, record.Name)
	 }
	 return nil
}

// delete  a record
func (m *Manager) DeleteRecord(id int) error {
	if id <= 0 {
		return ErrInvalidId
	}

	records, err := m.ReadRecord()
	if err != nil {
return err
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
		 return ErrNotFound
	}

	// reassign IDs (so they remain sequential)
	for i := range newRecord {
		newRecord[i].ID = i + 1
	}

	// marshal (convert go to json)
	m.SaveRecords(newRecord)

	fmt.Printf("🗑️  Deleted record with ID %d\n", id)
	return nil
}

// update a record
func (m *Manager) UpdateRecord(id int, newName string) error {
	if id == 0 {
		return ErrInvalidId
	}
	if newName == "" {
		return ErrEmptyName
	}

	records, err := m.ReadRecord()
	if err != nil {
		return err
	}
	
	found := false
	for i, r := range records {
		if r.ID == id {
			found = true
			records[i].Name = newName
			break
		}
	}

	if !found {
		return ErrNotFound
	}
	
	// marshal (convert go to json)
	if err:=m.SaveRecords(records); err != nil {
		return err
	}
	fmt.Printf("✏️  Updated record ID %d → New Name: %s\n", id, newName)
	return nil
}

// ========================== HELPER Functions ==========================
 
// ensure file exists
func (m *Manager) EnsureFile() error {
	_, err := os.Stat(m.fileName)
	if os.IsNotExist(err) {
		file, err := os.Create(m.fileName)
		if err != nil {
			return err
		}
		defer file.Close()
		file.Write([]byte("[]")) // initialize  empty JSON array	
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
		return nil, fmt.Errorf("❌ Error reading file: %w", err)
	}

	var people []types.Record
	// unmarshal (convert json to go)
	if err := json.Unmarshal(data, &people); err != nil {
		return nil, fmt.Errorf("❌ Invalid JSON format: %w", err)
	}
	return people, nil

}

// save records to file [ convert go to json and write to file]
func (m *Manager) SaveRecords(records []types.Record) error {
	// marshal (convert go to json)
	data, err := json.MarshalIndent(records, "", " ")
	if err != nil {
		fmt.Println("❌ Error marshalling data:", err)
		return err
	}
	
	// write to json file
	if err := os.WriteFile(m.fileName, data, 0644); err != nil {
		return fmt.Errorf("❌ Error writing file: %w", err)
	}
	return nil
}

