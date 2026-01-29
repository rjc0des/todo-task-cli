package cmd

import (
	"fmt"
	"os"
	"task-cli/internal/model"
	"task-cli/internal/store"
	"text/tabwriter"
)

func ListCommand(args string) error {
	db, err := store.Load()

	if err != nil {
		return fmt.Errorf("Error while loading the data\n")
	}

	fmt.Println("Tasks:")

	if len(db.Tasks) == 0 {
		return fmt.Errorf("No tasks found\n")
	}

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 3, ' ', 0)
	fmt.Fprintln(w, "id\tdescription\tstatus\tcreated_at\tupdated_at")

	if len(args) == 0 {

		for _, t := range db.Tasks {
			fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\n", t.ID, t.Description, t.Status, t.CreatedAt.Format("2006-01-02 15:04"),
				t.UpdatedAt.Format("2006-01-02 15:04"))
		}

	} else {
		if !model.TaskStatus.IsStatusValid(model.TaskStatus(args)) {
			return fmt.Errorf("status invalid\n")
		}

		for _, task := range db.Tasks {
			if task.Status != model.TaskStatus(args) {
				continue
			}

			fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\n", task.ID, task.Description, task.Status, task.CreatedAt.Format("2006-01-02 15:04"),
				task.UpdatedAt.Format("2006-01-02 15:04"))
		}
	}

	w.Flush()
	return nil
}
