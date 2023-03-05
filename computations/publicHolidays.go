package computations

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
	"time"
)

// holidays is global variable. I use it to store the entries from db, so I don't have to make a query at every call
var holidays = make(map[string]bool)

// GetRomanianHolidays establishes a connection to a PostgreSQL database and retrieves all the dates for Romanian holidays from the "Holidays" table. It stores the holidays in a map.
func GetRomanianHolidays(connection string) (map[string]bool, error) {
	conn, err := pgx.Connect(context.Background(), connection)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer func(conn *pgx.Conn, ctx context.Context) {
		err := conn.Close(ctx)
		if err != nil {

		}
	}(conn, context.Background()) //closing the connection

	rows, err := conn.Query(context.Background(), "SELECT * FROM Holidays")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var holiday time.Time
		err = rows.Scan(&holiday)
		if err != nil {
			return nil, err
		}
		holidays[holiday.Format("2006-01-02")] = true
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return holidays, nil
}

// IsHoliday verifies if a date is a holiday or not
func IsHoliday(date time.Time) bool {
	key := date.Format("2006-01-02")
	_, ok := holidays[key]
	return ok
}
