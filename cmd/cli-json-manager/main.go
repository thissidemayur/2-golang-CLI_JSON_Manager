package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/thissidemayur/cli-json-manager/internal/cli/commands"
)


func main(){
	// CLI Commands
	addCmd := flag.NewFlagSet("add" , flag.ExitOnError)
	name:= addCmd.String("name","","Name to add")

	listCmd := flag.NewFlagSet("list",flag.ExitOnError)

	deleteCmd := flag.NewFlagSet("delete",flag.ExitOnError)
	deleteId := deleteCmd.Int("id",0,"Id to delete")

	updateCmd := flag.NewFlagSet("update",flag.ExitOnError)
	updateId := updateCmd.Int("id",0,"Id to update")
	updateName := updateCmd.String("name","","New Name")
	// ensure subcommand are provided
	if len(os.Args) < 2 {
		fmt.Println("Usage: cli-json-manager <command> [options] <value>")
		fmt.Println("Commands: add, list, delete, update")
		return
	}

	// switch between command
	switch os.Args[1]{
	case "add":
		addCmd.Parse(os.Args[2:]);
		if *name =="" {
			fmt.Println("❌ Name is required. Usage: cli-json-manager add -name <name>")
			return;
		} 
		commands.AddRecord(name)

	case "list":
		listCmd.Parse(os.Args[2:])
		commands.ListRecord()

	case "delete":
		deleteCmd.Parse(os.Args[2:])
			if  *deleteId == 0 {
			fmt.Println("❌ ID is required. Usage: cli-json-manager delete -id <id>")
			return;
		}
		commands.DeleteRecord(*deleteId)

	case "update":
		updateCmd.Parse(os.Args[2:])
			if  *updateId == 0 {
			fmt.Println("❌ ID is required. Usage: cli-json-manager update -id <id>")
			return;
			}

			if *name =="" {
			fmt.Println("❌ Name is required. Usage: cli-json-manager update -name <name>")
			return;
		
		}
		commands.UpdateRecord(*updateId,*updateName)

	default:
		fmt.Println("❌ Unknown command: ",os.Args[1],	" Available commands: add, list, delete, update")
	}
}