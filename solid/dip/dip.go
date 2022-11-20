package dip

import "fmt"

// // This is againt DIP because a lower level (db) function is being
// // depended by a high level function
// type MySQL struct{}

// func (db MySQL) QuerySomeData() []string {
// 	return []string{"inf1", "inf2", "inf3"}
// }

type DBConn interface {
	QuerySomeData() []string
}

type MySQL struct{}

func (db MySQL) QuerySomeData() []string {
	return []string{"inf1", "inf2", "inf3"}
}

type MyRepository struct {
	// // This is against DIP as its relying on an implementation instead of abstraction
	// db MySQL

	// This fixes the above to support DIP
	db DBConn
}

func (r MyRepository) GetData() []string {
	return r.db.QuerySomeData()
}

func RunDip() {
	mysqlDB := MySQL{}
	repo := MyRepository{db: mysqlDB}
	fmt.Println(repo.GetData())
}
