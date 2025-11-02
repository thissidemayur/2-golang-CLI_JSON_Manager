package commands

import (
	"fmt"

	"github.com/thissidemayur/cli-json-manager/internal/types"
)

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
func (m *Manager) ListRecord() ([]types.Record, error) {
	records, err := m.ReadRecord()
	if err != nil {
		return nil, err
	}
	if len(records) == 0 {
		fmt.Printf("ℹ️ No records found in %q\n	", m.fileName)
		return nil, nil
	}

	fmt.Println("📋 List of Records:")
	for _, record := range records {
		fmt.Printf("  - ID: %d, Name: %s\n", record.ID, record.Name)
	}
	return records, nil
}
func (m *Manager) DeleteRecord(id int) error {
	if id <= 0 {
		return ErrInvalidId
	}

	records, err := m.ReadRecord()
	if err != nil {
		return err
	}

	found := false
	newRecord := []types.Record{}
	for _, record := range records {
		if record.ID == id {
			found = true
			continue
		} else {
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
	if err := m.SaveRecords(records); err != nil {
		return err
	}
	fmt.Printf("✏️  Updated record ID %d → New Name: %s\n", id, newName)
	return nil
}
