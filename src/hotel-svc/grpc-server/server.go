package grpcserver

import (
	"context"
	"log"

	"github.com/maxturyev/booking-system-project/hotel-svc/db"
	pb "github.com/maxturyev/booking-system-project/src/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

// HotelServer fetches info from the database and sends a grpc response
type HotelServer struct {
	pb.UnimplementedHotelServiceServer
	DB *gorm.DB
}

// GetHotels lists hotels from the database
func (s *HotelServer) GetHotels(req *pb.GetHotelsRequest, stream pb.HotelService_GetHotelsServer) error {
	log.Println("GetHotels запрос получен")
	hotels, _ := db.SelectHotels(s.DB)
	for _, hotel := range hotels {
		response := &pb.GetHotelsResponse{
			Hotel: &pb.Hotel{
				HotelID:       int32(hotel.HotelID),
				Name:          hotel.Name,
				Rating:        int32(hotel.Rating),
				Country:       hotel.Country,
				Description:   hotel.Description,
				RoomAvailable: int32(hotel.RoomsAvailable),
				RoomPrice:     float32(hotel.RoomPrice),
				Address:       hotel.Address,
			},
		}
		if err := stream.Send(response); err != nil {
			log.Println("Ошибка при отправке данных клиенту")
			return status.Errorf(codes.Internal, "Ошибка отправки данных")
		}
	}
	return nil
}

func (s *HotelServer) GetHotelPriceByID(ctx context.Context, req *pb.GetHotelPriceByIDRequest) (*pb.GetHotelPriceByIDResponse, error) {
	log.Println("Get ID")
	hotelID := req.GetId()
	hotel, _ := db.SelectHotelByID(s.DB, int(hotelID))
	response := &pb.GetHotelPriceByIDResponse{RoomPrice: hotel.RoomPrice}
	return response, nil
}
