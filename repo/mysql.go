package repo

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

// MySQL implementation of Repository interface
type MySQLRepo struct {
	db *sql.DB
}

func (r *MySQLRepo) Init() {
	//TODO: extract properties
	dbDriver := "mysql"
	dbUser := "gpractice"
	dbPass := "123"
	dbName := "gpractice"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic("failed to initialize db connection")
	}

	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 2)
	db.SetMaxOpenConns(3)
	db.SetMaxIdleConns(3)
	r.db = db
	return
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

func (r *MySQLRepo) GetAll() ([]Item, error) {
	var items []Item

	q := "SELECT id, date, duration FROM practice"
	stmt, err := r.db.Prepare(q)
	if err != nil {
		return items, err
	}

	rows, err := stmt.Query()
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

	log.Printf("Items=%+v, succesfully processed %v out of %v", items, len(items), count)
	return items, nil
}
