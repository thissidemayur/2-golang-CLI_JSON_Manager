package main

import (
	"flag"
	"fmt"
	"os"

	logger "github.com/thissidemayur/cli-json-manager"
	"github.com/thissidemayur/cli-json-manager/internal/cli/commands"
)

func main() {
	// Detect environment (default: dev)
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	// Initialize logger
	logger.InitLogger(env)

	// Global flags
	fileFlag := flag.String("file", "data.json", "Path to JSON file (default: data.json)")
	flag.Parse()

	if flag.NArg() < 1 {
		printUsage()
		os.Exit(1)
	}

	cmd := flag.Arg(0)

	// Handle subcommands
	switch cmd {
	case "add":
		runAdd(*fileFlag)
	case "list":
		runList(*fileFlag)
	case "delete":
		runDelete(*fileFlag)
	case "update":
		runUpdate(*fileFlag)
	default:
		logger.Error("Unknown command", "command", cmd)
		printUsage()
		os.Exit(1)
	}
}

// ============================ COMMAND HANDLERS ============================

// ADD
func runAdd(file string) {
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	name := addCmd.String("name", "", "Name to add")
	addCmd.Parse(flag.Args()[1:])

	if *name == "" {
		fmt.Println("❌ Missing required flag: --name")
		os.Exit(2)
	}

	mgr, err := commands.NewManager(file)
	if err != nil {
		logger.Error("Error creating manager", "error", err)
		fmt.Println("⚠️  Failed to initialize manager.")
		os.Exit(3)
	}

	if err := mgr.AddRecord(name); err != nil {
		logger.Error("Error adding record", "error", err)
		fmt.Println("❌ Failed to add record:", err)
		os.Exit(4)
	}

	fmt.Printf("✅ Added record: %q successfully!\n", *name)
}

// LIST
func runList(file string) {
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	listCmd.Parse(flag.Args()[1:])

	mgr, err := commands.NewManager(file)
	if err != nil {
		logger.Error("Error creating manager", "error", err)
		fmt.Println("⚠️  Failed to initialize manager.")
		os.Exit(3)
	}

	records, err := mgr.ListRecord()
	if err != nil {
		logger.Error("Error listing records", "error", err)
		fmt.Println("❌ Failed to list records:", err)
		os.Exit(4)
	}

	if len(records) == 0 {
		fmt.Println("ℹ️  No records found.")
		return
	}

	fmt.Println("\n📋 List of Records:")
	for _, record := range records {
		fmt.Printf("  - ID: %-3d Name: %s\n", record.ID, record.Name)
	}
}

// DELETE
func runDelete(file string) {
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	id := deleteCmd.Int("id", 0, "ID to delete")
	deleteCmd.Parse(flag.Args()[1:])

	mgr, err := commands.NewManager(file)
	if err != nil {
		logger.Error("Error creating manager", "error", err)
		fmt.Println("⚠️  Failed to initialize manager.")
		os.Exit(3)
	}

	if err := mgr.DeleteRecord(*id); err != nil {
		logger.Error("Error deleting record", "id", *id, "error", err)
		fmt.Printf("❌ Failed to delete record with ID %d: %v\n", *id, err)
		os.Exit(4)
	}

	fmt.Printf("🗑️  Deleted record with ID %d successfully.\n", *id)
}

// UPDATE
func runUpdate(file string) {
	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	id := updateCmd.Int("id", 0, "ID to update")
	name := updateCmd.String("name", "", "New name")
	updateCmd.Parse(flag.Args()[1:])

	mgr, err := commands.NewManager(file)
	if err != nil {
		logger.Error("Error creating manager", "error", err)
		fmt.Println("⚠️  Failed to initialize manager.")
		os.Exit(3)
	}

	if err := mgr.UpdateRecord(*id, *name); err != nil {
		logger.Error("Error updating record", "id", *id, "error", err)
		fmt.Printf("❌ Failed to update record with ID %d: %v\n", *id, err)
		os.Exit(4)
	}

	fmt.Printf("✏️  Updated record ID %d → %q successfully.\n", *id, *name)
}

// USAGE HELP
func printUsage() {
	fmt.Println(`
Usage:
  cli-json-manager [--file <fileName>] <command> [options]

Commands:
  add     --name <string>         Add a new record
  list                           List all records
  delete  --id <int>             Delete a record by ID
  update  --id <int> --name <string>  Update record name

Examples:
  cli-json-manager add --name "Bhawaniputra"
  cli-json-manager list
  cli-json-manager delete --id 2
  cli-json-manager update --id 3 --name "TaraPutra"
`)
}
