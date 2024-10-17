# task-tracker

### https://roadmap.sh/projects/task-tracker

```
# Adding a new task
go run . add "Buy groceries"
# Output: Task added successfully (ID: 1)

# Updating and deleting tasks
go run . update 1 "Buy groceries and cook dinner"
go run . delete 1

# Marking a task as in progress or done
go run . mark-in-progress 1
go run . mark-done 1

# Listing all tasks
go run . list

# Listing tasks by status
go run . list done
go run . list todo
go run . list in-progress
```