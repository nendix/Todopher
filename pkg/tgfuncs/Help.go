package tg

import "fmt"

func PrintHelp() {
	fmt.Println("Usage: tg [command] [options]")
	fmt.Println("Commands:")
	fmt.Println("  list                                     - List all todos")
	fmt.Println("  add [todo] [dd-mm-yy]                    - Add a new todo")
	fmt.Println("  edit [id] [new todo] [new dd-mm-yy]      - Edit a todo")
	fmt.Println("  mark [id1 id2 ...]                       - Mark todos as completed")
	fmt.Println("  unmark [id1 id2 ...]                     - Unmark todos as not completed")
	fmt.Println("  search [key_word]                        - List all todos that contain the keyword")
	fmt.Println("  sort [id] [by_status|by_date]            - Sort todos by status or by date")
	fmt.Println("  delete [id1 id2 ...]                     - Delete todos")
}
