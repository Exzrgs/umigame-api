package tasks

import (
	"time"

	"database/sql"
)

func ExeTasks(db *sql.DB) {
	go func() {
		t := time.NewTicker(time.Minute)
		for {
			select {
			case <-t.C:
				ActivateCheck(db)
			}
		}
		t.Stop()
	}()
}
