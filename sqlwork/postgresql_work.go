package sqlwork

import (
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/maxturyev/booking-system-project/consts"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Понять зачем sslmode
// Для начала работы прописать в константах правильные параметры для локальной машины, после этого создать бд,
// создать таблицу и после этого можно добавлять заказы. Проверять, правильно ли все получилось, можно через GetBookInfo
func CreateDatabase(namedb string) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable", consts.Host, consts.Port,
		consts.User, consts.Password)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	createDatabaseCommand := fmt.Sprintf("CREATE DATABASE %s", namedb)
	db.Exec(createDatabaseCommand)
}

func DropDatabase(namedb string) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable", consts.Host, consts.Port,
		consts.User, consts.Password)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	dropDatabaseCommand := fmt.Sprintf("DROP DATABASE %s", namedb)
	db.Exec(dropDatabaseCommand)
}

func CreateTable(dbname, nameTable, command string) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", consts.Host, consts.Port,
		consts.User, consts.Password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.Exec(command, nameTable)
}

func DropTable(dbname, nameTable string) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", consts.Host, consts.Port,
		consts.User, consts.Password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	dropTableCommand := fmt.Sprintf("DROP TABLE %s", nameTable)
	db.Exec(dropTableCommand)
}

func InsertInto(dbname, nameTable, command string, hotel_id int32, date_start, date_end string,
	room_category, room_count, price, guest_amount int32, is_cancelled bool) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", consts.Host, consts.Port,
		consts.User, consts.Password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.Exec(command, nameTable, hotel_id, date_start, date_end, room_category, room_count, price, guest_amount, is_cancelled)
}

func DeleteFrom(dbname, nameTable, command string, hotel_id int32) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", consts.Host, consts.Port,
		consts.User, consts.Password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.Exec(command, nameTable, hotel_id)
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

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", consts.Host, consts.Port,
		consts.User, consts.Password, consts.DBBooking)
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

	return answer
}
