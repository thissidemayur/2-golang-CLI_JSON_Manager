package commands

import (
	logger "github.com/thissidemayur/cli-json-manager"
	"github.com/thissidemayur/cli-json-manager/internal/types"
)

// ========================== CRUD OPERATIONS ==========================

// ADD
func (m *Manager) AddRecord(name *string) error {
	records, err := m.ReadRecord()
	if err != nil {
		logger.Error("Error reading records", "error", err)
		return err
	}

	newRecord := types.Record{
		ID:   len(records) + 1,
		Name: *name,
	}
	records = append(records, newRecord)

	if err := m.SaveRecords(records); err != nil {
		logger.Error("Error saving records", "error", err)
		return err
	}

	logger.Info("Record added successfully", "id", newRecord.ID, "name", newRecord.Name)
	return nil
}

// LIST
func (m *Manager) ListRecord() ([]types.Record, error) {
	records, err := m.ReadRecord()
	if err != nil {
		logger.Error("Error reading records", "error", err)
		return nil, err
	}

	if len(records) == 0 {
		logger.Info("No records found", "file", m.fileName)
		return nil, nil
	}

	logger.Info("Records fetched successfully", "count", len(records))
	return records, nil
}

// DELETE
func (m *Manager) DeleteRecord(id int) error {
	if id <= 0 {
		logger.Warn("Invalid record ID", "id", id)
		return ErrInvalidId
	}

	records, err := m.ReadRecord()
	if err != nil {
		logger.Error("Error reading records", "error", err)
		return err
	}

	newRecords := []types.Record{}
	found := false
	for _, r := range records {
		if r.ID == id {
			found = true
			continue
		}
		newRecords = append(newRecords, r)
	}

	if !found {
		logger.Warn("Record not found", "id", id)
		return ErrNotFound
	}

	// Reassign sequential IDs
	for i := range newRecords {
		newRecords[i].ID = i + 1
	}

	if err := m.SaveRecords(newRecords); err != nil {
		logger.Error("Error saving records after delete", "error", err)
		return err
	}

	logger.Info("Record deleted", "id", id)
	return nil
}

// UPDATE
func (m *Manager) UpdateRecord(id int, newName string) error {
	if id <= 0 {
		logger.Warn("Invalid record ID", "id", id)
		return ErrInvalidId
	}
	if newName == "" {
		logger.Warn("Empty new name provided", "id", id)
		return ErrEmptyName
	}

	records, err := m.ReadRecord()
	if err != nil {
		logger.Error("Error reading records", "error", err)
		return err
	}

	found := false
	for i := range records {
		if records[i].ID == id {
			found = true
			records[i].Name = newName
			break
		}
	}

	if !found {
		logger.Warn("Record not found for update", "id", id)
		return ErrNotFound
	}

	if err := m.SaveRecords(records); err != nil {
		logger.Error("Error saving updated records", "error", err)
		return err
	}

	logger.Info("Record updated", "id", id, "newName", newName)
	return nil
}
