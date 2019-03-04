package database

import (
	"log"
	"time"

	"github.com/go-pg/pg"
	_ "github.com/lib/pq"
)

// New Creates database connection to a postgres database
func NewPostgres(DSN string, timeout int, enableQueryLog bool) (*pg.DB, error) {
	/* u, err := pg.ParseURL(DSN)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	*/

	// TODO: Need to use DSN to make this dynamic
	u := &pg.Options{
		User:     "postgres",
		Password: "postgres",
		Database: "arya_auth",
		Addr:     "127.0.0.1:5432",
	}

	db := pg.Connect(u)

	_, err := db.Exec("Select 1")
	if err != nil {
		log.Println("Error While Registering DB")
		return nil, err
	}

	log.Println("Successfully Registered DB")

	if timeout > 0 {
		db.WithTimeout(time.Second * time.Duration(timeout))
	}

	if enableQueryLog {
		db.OnQueryProcessed(func(event *pg.QueryProcessedEvent) {
			if query, err := event.FormattedQuery(); err == nil {
				log.Printf("%s | %s", time.Since(event.StartTime), query)
			}
		})
	}

	return db, nil
}
