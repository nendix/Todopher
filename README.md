# TaskGopher

A simple file-based todo manager written in Go

### Usage:

**tg** [command] [options] <br>

#### Commands:

- **add** [todo] [dd-mm-yy] - Add a new todo
- **edit** [id] [new todo] [new dd-mm-yy] - Edit a todo
- **mark** [id1 id2 ...] - Mark todos as completed
- **unmark** [id1 id2 ...] - Unmark todos as not completed
- **list** - List all todos
- **search** [key_word] - List all todos that contain the keyword
- **sort** [id] [by_status|by_date] - Sort todos by status or by date
- **delete** [id1 id2 ...] - Delete todos

#### NOTE:

TaskGopher saves the todos in _your_home_dir_/todo/todos.txt <br>
To change the default saving path and file name you have to modify the **filename variable** inside the **todo.go** file with your desired path and filename.
