package sqlwork

import (
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = "5433"
	user     = "postgres"
	password = "password123"
	dbname   = "database"
)

type BookInfo struct {
	hotel_id      int32
	date_start    time.Time
	date_end      time.Time
	room_category int32
	room_count    int32
	price         int32
	guest_amount  int32
	is_cancelled  bool
}

func (book BookInfo) GetInfo() string {
	return fmt.Sprintf("hotel_id: %d\ndate_start: %s\ndate_end %s\nroom_category: %d\nroom_count: %d\nprice: %d\n"+
		"guest_amount: %d\nis_cancelled: %t\n", book.hotel_id, book.date_start, book.date_end, book.room_category,
		book.room_count, book.price, book.guest_amount, book.is_cancelled)
}

// Понять зачем sslmode
// Для начала работы прописать в константах правильные параметры для локальной машины, после этого создать бд,
// создать таблицу и после этого можно добавлять заказы. Проверять, правильно ли все получилось, можно через GetBookInfo
func CreateDatabase(namedb string) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable", host, port, user, password)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	createDatabaseCommand := fmt.Sprintf("CREATE DATABASE %s", namedb)
	db.Exec(createDatabaseCommand)
}

func DropDatabase(namedb string) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable", host, port, user, password)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	dropDatabaseCommand := fmt.Sprintf("DROP DATABASE %s", namedb)
	db.Exec(dropDatabaseCommand)
}

func CreateTable(nameTable string) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	createTableCommand := fmt.Sprintf("CREATE TABLE %s (hotel_id SERIAL PRIMARY KEY, date_start DATE, date_end DATE, "+
		"room_category INTEGER DEFAULT 0, room_count INTEGER DEFAULT 0, price INTEGER DEFAULT 0,"+
		"guest_amount INTEGER DEFAULT 0, is_cancelled BOOLEAN DEFAULT FALSE)", nameTable)
	db.Exec(createTableCommand)
}

func DropTable(nameTable string) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	dropTableCommand := fmt.Sprintf("DROP TABLE %s", nameTable)
	db.Exec(dropTableCommand)
}

func InsertInto(nameTable string, hotel_id int32, date_start, date_end string,
	room_category, room_count, price, guest_amount int32, is_cancelled bool) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.Exec("INSERT INTO ? (hotel_id, date_start, date_end, room_category, "+
		"room_count, price, guest_amount, is_cancelled) VALUES (?, ?, ?, ?, ?, ?, ?, ?)", nameTable,
		hotel_id, date_start, date_end, room_category, room_count, price, guest_amount, is_cancelled)
}

func GetBookInfo(namedb string, hotel_id int) BookInfo {
	// dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	// db, err := sql.Open("postgres", dsn)

	// if err != nil {
	// 	panic(err)
	// }
	// // defer db.Close()

	// createDatabaseCommand := fmt.Sprintf("INSERT INTO %s VALUES (%s %s %s)", "booking_data", 0, "2020-10-11", "2020-10-17",
	// 	2, 1, 12000, 2, true)
	// db.Exec(createDatabaseCommand)

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	var results []map[string]interface{}
	exec_string := fmt.Sprintf("SELECT * FROM %s WHERE hotel_id = %d", namedb, hotel_id)
	db.Raw(exec_string).Scan(&results)

	// fmt.Println(reflect.TypeOf(results[0]["hotel_id"]))

	var answer BookInfo
	answer.hotel_id = results[0]["hotel_id"].(int32)
	answer.date_start = results[0]["date_start"].(time.Time)
	answer.date_end = results[0]["date_end"].(time.Time)
	answer.room_category = results[0]["room_category"].(int32)
	answer.room_count = results[0]["room_count"].(int32)
	answer.price = results[0]["price"].(int32)
	answer.guest_amount = results[0]["guest_amount"].(int32)
	answer.is_cancelled = results[0]["is_cancelled"].(bool)

	fmt.Println(answer.GetInfo())

	return answer
}
