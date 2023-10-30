package mysql

import (
	"database/sql"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"goWithHtmx/model"
	"os"
)

type MysqlRepo struct {
	db *sql.DB
}

type mysqlItem struct {
	ID        uint64 `json:"id"`
	Title     string `json:"title"`
	Available bool   `json:"available"`
}

func (m mysqlItem) ToItem() model.Item {
	i := model.Item{}
	i.SetID(m.ID)
	i.SetTile(m.Title)
	i.SetAvailable(m.Available)
	return i
}

func New(connectionString string) (*MysqlRepo, error) {
	dsn := flag.String("dsn", os.Getenv("DSN"), connectionString)
	flag.Parse()

	db, err := sql.Open("mysql", *dsn)
	if err != nil {
		panic(err.Error())
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)

	return &MysqlRepo{
		db: db,
	}, nil
}

func (m MysqlRepo) GetAll() ([]model.Item, error) {
	var items []model.Item

	query, err := m.db.Query(`select * from usssers;`)
	if err != nil {
		return items, err
	}

	defer query.Close()

	for query.Next() {
		var i mysqlItem
		err = query.Scan(&i.ID, &i.Title, &i.Available)
		if err != nil {
			return nil, err
		}
		item := i.ToItem()
		items = append(items, item)
	}
	return items, err
}

func (m MysqlRepo) Get(id uint64) (model.Item, error) {
	//TODO implement me
	panic("implement me")
}

func (m MysqlRepo) Add(customer model.Item) error {
	//TODO implement me
	panic("implement me")
}

func (m MysqlRepo) Delete(id uint64) error {
	//TODO implement me
	panic("implement me")
}

func (m MysqlRepo) UpdateStatus(id uint64) error {
	//TODO implement me
	panic("implement me")
}
