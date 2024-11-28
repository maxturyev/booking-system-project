package sqlwork

import (
	"fmt"
	"time"
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
