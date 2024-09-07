# TaskGopher

A simple file-based todo manager written in Go

### Usage:

**tg** [command] [options] <br>

#### Commands:
- **help** _- Print the cmd list_
- **init** _- Initialize the enviroment_
- **setlist** [listname] _- Change the todo list_
- **add** [todo] [dd-mm-yy] _- Add a new todo_
- **edit** [id] [new todo] [new dd-mm-yy] _- Edit a todo_
- **mark** [id1 id2 ...] _- Mark todos as completed_
- **unmark** [id1 id2 ...] _- Unmark todos as not completed_
- **list** _- List all todos_
- **search** [key_word] _- List all todos that contain the keyword_
- **sort** [id] [by_status|by_date] _- Sort todos by status or by date_
- **delete** [id1 id2 ...] _- Delete todos_

#### NOTE:

TaskGopher saves the todos in *HOME_DIR/todo* and uses a *.env* file that stores the current todo list<br>
To **setup the enviroment** you need to run: `tg init` and then `tg setlist todos`.<br>
To **delete a list or change his name** you can use your system file manager or terminal commands,
just remember to update the *.env* file with `tg setlist [new_list_name]`.
