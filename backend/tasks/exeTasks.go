package tasks

import (
	"time"

	"github.com/jmoiron/sqlx"
)

func ExeTasks(db *sqlx.DB) {
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
