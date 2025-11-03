package commands

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/thissidemayur/cli-json-manager/internal/types"
)

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
func (m *Manager) ReadRecord() ([]types.Record, error) {
	if err := m.EnsureFile(); err != nil {
		return nil, err
	}

	// get file data
	data, err := os.ReadFile(m.fileName)
	if err != nil {
		return nil, fmt.Errorf("❌ Error reading file: %w", err)
	}

	// handle empty file
	if len(data) == 0 {
		return []types.Record{}, nil
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
