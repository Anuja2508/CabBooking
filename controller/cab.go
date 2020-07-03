package controller

import (
	"GoCab/model"

	"github.com/jinzhu/gorm"
)

// Cab :- For initalizing database instance
type Cab struct {
	database *gorm.DB
}

// NewCabController will create new user controller
func NewCabController(database *gorm.DB) Cab {
	return Cab{
		database: database,
	}
}

// FindNearByCab :- For finding near by location
func (controller Cab) FindNearByCab(latititute float64, longitiute float64) ([]model.NearByCabList, error) {
	var nearByCabList []model.NearByCabList

	if err := controller.database.Raw(`select * from (
	      SELECT  *,( 3959 * acos( cos( radians(17.048385) ) * cos( radians(latitute) ) * cos( radians(longititute) - radians(-0.575623) ) + sin( radians(17.048385) ) * sin( radians(latitute) ) ) )* 1.609344
 						AS distance FROM cabs
					 )al  
		  ORDER BY al.distance asc limit 5`).Scan(&nearByCabList).Error; err != nil {
		return nil, err

	}

	return nearByCabList, nil

}

// BookCab :- For storing booking details in Booking table
func (controller Cab) BookCab(br *model.BookingRequest) error {
	var b model.Booking
	b.CabID = br.CabID
	b.UserID = br.UserID
	b.Time = br.Time
	b.DestinationLocation = br.DestinationLocation
	b.PickupLocation = br.PickupLocation
	b.Status = 0
	if err := controller.database.Create(&b).Error; err != nil {
		return err
	}
	return nil

}

// UpdateBookingStatus :- For updating status of booking
func (controller Cab) UpdateBookingStatus(Status int, bookingID uint) error {

	if err := controller.database.Table("bookings").Where("id = ?", bookingID).UpdateColumn("status", Status).Error; err != nil {
		return err
	}
	return nil

}

// HistoryOfBooking :- For geeting of login user booking details
func (controller Cab) HistoryOfBooking(userID uint) ([]model.Booking, error) {

	var b []model.Booking

	if err := controller.database.Preload("User").Preload("Cab").Where("user_id = ?", userID).Find(&b).Error; err != nil {
		return b, err
	}
	return b, nil

}
