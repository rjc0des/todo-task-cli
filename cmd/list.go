package cmd

import (
	"fmt"
	"os"
	"task-cli/internal/store"
	"text/tabwriter"
)

func ListCommand() {
	db, err := store.Load()

	if err != nil {
		fmt.Println("Error while loading the data")
		return
	}

	fmt.Println("Tasks:")

	w := tabwriter.NewWriter(os.Stdout, 1, 1, 3, ' ', 0)

	fmt.Fprintln(w, "id\tdescription\tstatus\tcreated_at\tupdated_at")

	for _, t := range db.Tasks {
		fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%s\n", t.ID, t.Description, t.Status, t.CreatedAt.Format("2006-01-02 15:04"),
			t.UpdatedAt.Format("2006-01-02 15:04"))
	}

	w.Flush()
}
