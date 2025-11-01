package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/thissidemayur/cli-json-manager/internal/cli/commands"
)

func main() {
	// global flag share across all commands
	fileFlag := flag.String("file", "data.json", "JSON file to store data (default: data.json)")

	// parse only top-level (global) flags so user can do: cli-json-manager --file=user.json add...
	flag.Parse()

	// ensure subcommand are provided like add, list, delete, update
	if flag.NArg() < 1 {
		fmt.Println("❌ Missing command.")
		fmt.Println("Usage: cli-json-manager [--file <fileName] <command> [options]")
		fmt.Println("Commands: add, list, delete, update")
		os.Exit(1)
	}

	// first positional argument- subcommand (add, list, delete, update)
	cmd := flag.Arg(0)

	switch cmd {
	// ---------------- ADD ----------------
	case "add":
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)
		name := addCmd.String("name", "", "Name to add")
		addCmd.Parse(flag.Args()[1:]) // parse arguments *after* add
		if *name == "" {
			fmt.Println("❌ missing required flag: --name")
			os.Exit(2)
		}
		mgr, err := commands.NewManager(*fileFlag)
		if err != nil {
			fmt.Println("❌ Error creating manager:", err)
			os.Exit(3)

		}
		if err := mgr.AddRecord(name); err != nil {
			fmt.Println(err)
			os.Exit(4)
		}

	//  ---------------- LIST ----------------
	case "list":
		listCmd := flag.NewFlagSet("list", flag.ExitOnError)
		listCmd.Parse(flag.Args()[1:])
		mgr, err := commands.NewManager(*fileFlag)
		if err != nil {
			fmt.Println("❌ Error creating manager:", err)
			os.Exit(3)
		}
		if err := mgr.ListRecord(); err != nil {
			fmt.Println(err)
			os.Exit(3)
		}

	//  ---------------- DELETE ----------------
	case "delete":
		deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
		deleteId := deleteCmd.Int("id", 0, "Id to delete")
		deleteCmd.Parse(flag.Args()[1:])
		mgr, err := commands.NewManager(*fileFlag)
		if err != nil {
			fmt.Println("❌ Error creating manager:", err)
			os.Exit(3)
		}
		if err := mgr.DeleteRecord(*deleteId); err != nil {
			fmt.Println(err)
			os.Exit(4)
		}

	//  ---------------- UPDATE ----------------
	case "update":
		updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
		updateId := updateCmd.Int("id", 0, "Id to update")
		updateName := updateCmd.String("name", "", "New Name")
		updateCmd.Parse(flag.Args()[1:])

		mgr, err := commands.NewManager(*fileFlag)
		if err != nil {
			fmt.Println("❌ Error creating manager:", err)
			os.Exit(3)
		}
		if err := mgr.UpdateRecord(*updateId, *updateName); err != nil {
			fmt.Println(err)
			os.Exit(4)
		}

	//  ---------------- DEFAULT/UNKNOWN COMMAND ----------------
	default:
		fmt.Println("❌ Unknown command: ", cmd)
		fmt.Println("Available commands: add, list, delete, update")
		os.Exit(1)

	}
}

/*
Exit codes:
0 -> ✅success
1 -> ❌ general error (missing command, unknown command)
2 -> ❌ Bad usage, incorrect command syntax [command-specific error (missing flags, invalid input)]
3 -> ❌ Runtime error [manager creation error (file issues, etc.)]
4 -> ❌ Operation-specific error (add, delete, update failures)
*/
