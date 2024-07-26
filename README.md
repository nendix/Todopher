# GoToDo

A simple file-based todo manager written in Go

### Usage:

todo [command] [options] <br>

#### Commands:

- **add** [todo] [dd-mm-yy] - Add a new todo
- **edit** [id] [new todo] [new dd-mm-yy] - Edit a todo
- **mark** [id1 id2 ...] - Mark todos as completed
- **unmark** [id1 id2 ...] - Unmark todos as not completed
- **list** - List all todos
- **search** [key_word] - List all todos that contain the keyword
- **sort** [id] [by_status|by_date] - Sort todos by status or by date
- **delete** [id1 id2 ...] - Delete todos
