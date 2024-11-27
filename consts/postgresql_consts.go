package consts

const (
	Host      = "localhost"
	Port      = "5433"
	User      = "postgres"
	Password  = "password123"
	DBBooking = "database"
	DBHotel   = "dataBaseHotel"
)

const (
	CreateTableBookingCommand = "CREATE TABLE ? (hotel_id SERIAL PRIMARY KEY, date_start DATE, date_end DATE, " +
		"room_category INTEGER DEFAULT 0, room_count INTEGER DEFAULT 0, price INTEGER DEFAULT 0," +
		"guest_amount INTEGER DEFAULT 0, is_cancelled BOOLEAN DEFAULT FALSE)"
	InsertIntoBookingCommand = "INSERT INTO ? (hotel_id, date_start, date_end, room_category, " +
		"room_count, price, guest_amount, is_cancelled) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	DeleteFromBookingCommand = "DELETE FROM ? WHERE hotel_id = ?"
)
