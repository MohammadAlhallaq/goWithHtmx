package model

type Item struct {
	id        uint64
	title     string
	available bool
}

func (p Item) GetID() uint64 {
	return p.id
}

func (p Item) GetTitle() string {
	return p.title
}

func (p Item) GetPrice() bool {
	return p.available
}

func (p Item) SetID(id uint64) {
	p.id = id
}

func (p Item) SetTile(title string) {
	p.title = title
}

func (p Item) SetAvailable(available bool) {
	p.available = available
}

//func createItem(item string) error {
//	stm := `insert into todos(todo, available) values ($1, $2)`
//	_, err := db.Exec(stm, item, false)
//
//	if err != nil {
//		return err
//	}
//	return err
//}

//func GetAllItems() ([]Item, error) {
//	fmt.Println(db)
//	var items []Item
//	stm := `select * from users;`
//
//	rows, err := db.Query(stm)
//	if err != nil {
//		return nil, err
//	}
//	defer db.Close()
//
//	for rows.Next() {
//		var title string
//		var available bool
//		var id uint64
//
//		err = rows.Scan(&id, &title, &available)
//		if err != nil {
//			return nil, err
//		}
//		item := Item{Id: id, Title: title, Available: available}
//		items = append(items, item)
//	}
//	return items, err
//}

//func getItem(id uint64) (Item, error) {
//	item := Item{Id: id}
//	stm := `select * from items where id=$i;`
//	row, err := db.Query(stm, id)
//
//	if err != nil {
//		return item, err
//	}
//
//	for row.Next() {
//		var title string
//		var available bool
//		var id uint64
//
//		err = row.Scan(&id, &title, &available)
//		if err != nil {
//			return item, err
//		}
//
//		item.Title = title
//		item.Available = available
//	}
//	return item, err
//}

//func markAvailable(id uint64) error {
//	item, err := getItem(id)
//	if err != nil {
//		return err
//	}
//	stm := `update items where id=$1 set available=$2;`
//	_, err = db.Exec(stm, id, !item.Available)
//
//	return err
//}
//
//func deleteItem(id uint64) error {
//
//	stm := `delete items where id=$1`
//	_, err := db.Exec(stm, id)
//
//	return err
//}
