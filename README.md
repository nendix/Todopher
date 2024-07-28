# TaskGopher

A simple file-based todo manager written in Go

### Usage:

**tg** [command] [options] <br>

#### Commands:

- **add** [todo] [dd-mm-yy] _- Add a new todo_
- **edit** [id] [new todo] [new dd-mm-yy] _- Edit a todo_
- **mark** [id1 id2 ...] _- Mark todos as completed_
- **unmark** [id1 id2 ...] _- Unmark todos as not completed_
- **list** _- List all todos_
- **search** [key_word] _- List all todos that contain the keyword_
- **sort** [id] [by_status|by_date] _- Sort todos by status or by date_
- **delete** [id1 id2 ...] _- Delete todos_

#### NOTE:

TaskGopher saves the todos in _your_home_dir_/todo/todos.txt <br>
To change the default saving path and file name you have to modify the **filename variable** inside the **todo.go** file with your desired path and filename.
