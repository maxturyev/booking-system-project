package grpcserver

import (
	pb "github.com/maxturyev/booking-system-project/mocks/grpc"
	"gorm.io/gorm"
)

// HotelServer fetches info from the database and sends a grpc response
type HotelServer struct {
	pb.UnimplementedHotelServiceServer
	DB *gorm.DB
}
