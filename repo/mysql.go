package repo

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/vk23/gpractice/model"
	"log"
	"time"
)

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

func (r *MySQLRepo) Save(item model.Item) model.Item {
	result := model.Item{}
	if item.Id == NewId {
		result = r.insert(item)
	} else {
		result = r.update(item)
	}
	log.Printf("Save result=%v\n", result)
	return result
}

func (r *MySQLRepo) insert(item model.Item) model.Item {
	q := "INSERT INTO practice (date, duration) VALUES (?,?)"
	stmt, err := r.db.Prepare(q)
	//TODO: fix panics
	if err != nil {
		panic(err.Error())
	}
	res, err := stmt.Exec(item.Date, item.Duration)
	if err != nil {
		panic(err.Error())
	}
	id, err := res.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	item.Id = uint64(id)
	return item
}

func (r *MySQLRepo) update(item model.Item) model.Item {
	q := "UPDATE practice SET date=?, duration=? where id=?"
	stmt, err := r.db.Prepare(q)
	//TODO: fix panics
	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(item.Date, item.Duration, item.Id)
	if err != nil {
		panic(err.Error())
	}
	return item
}

func (r *MySQLRepo) Delete(id uint64) bool {
	q := "DELETE FROM practice WHERE id=?"
	stmt, err := r.db.Prepare(q)
	//TODO: fix panics
	if err != nil {
		panic(err.Error())
	}
	res, err := stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}
	count, err := res.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	return count > 0
}

func (r *MySQLRepo) Get(id uint64) model.Item {
	q := "SELECT id, date, duration FROM practice WHERE id=?"
	stmt, err := r.db.Prepare(q)
	if err != nil {
		panic(err.Error())
	}
	row := stmt.QueryRow(id)
	item := model.Item{}
	err = row.Scan(&item.Id, &item.Date, &item.Duration)
	if err != nil {
		panic(err.Error())
	}
	log.Printf("Get result=%v%n", item)

	return item
}

func (r *MySQLRepo) GetAll() []model.Item {
	q := "SELECT id, date, duration FROM practice"
	stmt, err := r.db.Prepare(q)
	if err != nil {
		panic(err.Error())
	}
	rows, err := stmt.Query()
	if err != nil {
		panic(err.Error())
	}

	var items []model.Item
	for rows.Next() {
		i := model.Item{}
		err := rows.Scan(&i.Id, &i.Date, &i.Duration)
		if err != nil {
			panic(err.Error())
		}
		items = append(items, i)
	}

	log.Printf("Items=%+v, count=%v", items, len(items))
	return items
}
