package repo

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

// MySQL implementation of Repository interface
type MySQLRepo struct {
	DbUser string
	DbPass string
	DbHost string
	DbPort string
	DbName string
	db     *sql.DB
}

// Initialize MySQL repository and verify connection
func (r *MySQLRepo) Init() error {
	db, err := sql.Open("mysql", r.DbUser+":"+r.DbPass+"@tcp("+r.DbHost+":"+r.DbPort+")/"+r.DbName)
	if err != nil {
		return err
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute)
	db.SetMaxOpenConns(3)
	db.SetMaxIdleConns(3)

	// verify connection
	err = db.Ping()
	if err != nil {
		return err
	}

	r.db = db
	return nil
}

func (r *MySQLRepo) Save(item Item) (Item, error) {
	if item.Id == NewId {
		return r.insert(item)
	} else {
		return r.update(item)
	}
}

func (r *MySQLRepo) insert(item Item) (Item, error) {
	q := "INSERT INTO practice (date, duration) VALUES (?, ?)"
	stmt, err := r.db.Prepare(q)
	if err != nil {
		return Item{}, err
	}

	res, err := stmt.Exec(item.Date, item.Duration)
	if err != nil {
		return Item{}, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return Item{}, err
	}

	item.Id = int(id)
	return item, nil
}

func (r *MySQLRepo) update(item Item) (Item, error) {
	q := "UPDATE practice SET date=?, duration=? where id=?"
	stmt, err := r.db.Prepare(q)
	if err != nil {
		return Item{}, err
	}

	_, err = stmt.Exec(item.Date, item.Duration, item.Id)
	if err != nil {
		return Item{}, err
	}

	return item, nil
}

func (r *MySQLRepo) Delete(id int) (bool, error) {
	q := "DELETE FROM practice WHERE id=?"
	stmt, err := r.db.Prepare(q)
	if err != nil {
		return false, err
	}

	res, err := stmt.Exec(id)
	if err != nil {
		return false, err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *MySQLRepo) Get(id int) (Item, error) {
	q := "SELECT id, date, duration FROM practice WHERE id=?"
	stmt, err := r.db.Prepare(q)
	if err != nil {
		return Item{}, err
	}

	row := stmt.QueryRow(id)
	item := Item{}
	err = row.Scan(&item.Id, &item.Date, &item.Duration)
	if err != nil {
		return Item{}, err
	}

	return item, nil
}

func (r *MySQLRepo) GetAll(from time.Time, to time.Time) ([]Item, error) {
	var items []Item

	q := "SELECT id, date, duration FROM practice WHERE date >= ? AND date <= ? ORDER BY id"
	stmt, err := r.db.Prepare(q)
	if err != nil {
		return items, err
	}

	rows, err := stmt.Query(from, to)
	if err != nil {
		return items, err
	}

	count := 0
	for rows.Next() {
		count++
		i := Item{}
		err := rows.Scan(&i.Id, &i.Date, &i.Duration)
		if err != nil {
			log.Printf("Failed to scan item: %s", err.Error())
			continue
		}
		items = append(items, i)
	}

	log.Printf("GetAll: succesfully processed %v out of %v", len(items), count)
	return items, nil
}

func (r *MySQLRepo) Close() {
	if r.db == nil {
		return
	}
	err := r.db.Close()
	if err != nil {
		log.Printf("Failed to close db connection: %s", err.Error())
	}
	log.Print("Database connection is closed")
}
