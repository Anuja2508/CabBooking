package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

// User :- Table structure for storing register users.
type User struct {
	gorm.Model
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	MobileNumber string `json:"phone_number"`
	Email        string `json:"email"`
}

// Cab :- Table structure for storing cab details
type Cab struct {
	gorm.Model
	NumberPlate     string  `json:"number_plate"`
	DriverFirstName string  `json:"driver_first_name"`
	DriverLastName  string  `json:"driver_last_name"`
	Latitute        float64 `json:"latitute"`
	Longititute     float64 `json:"longititue"`
	Available       bool    `json:"available"`
}

// Booking :- Table structure for storing cab bookings , status 0 means pending from driver side , 1 Accepted by the driver , 2 ride completed
type Booking struct {
	gorm.Model
	UserID              uint
	User                User `gorm:"foreignkey:ID;association_foreignkey:UserID"`
	CabID               uint
	Cab                 Cab `gorm:"foreignkey:ID;association_foreignkey:CabID"`
	Time                *time.Time
	PickupLocation      string
	DestinationLocation string
	Status              int
}

// UserClaims are custom claims extending default ones.
type UserClaims struct {
	PhoneNumber string `json:"phone_number"`
	jwt.StandardClaims
}

// NearByCabList :- Response struct for getting near by list
type NearByCabList struct {
	ID              uint    `json:"cab_id"`
	NumberPlate     string  `json:"number_plate"`
	DriverFirstName string  `json:"driver_first_name"`
	DriverLastName  string  `json:"driver_last_name"`
	Distance        float64 `json:"distance"`
}

// BookingRequest :- Request body for booking cab
type BookingRequest struct {
	UserID              uint
	CabID               uint       `json:"cab_id"`
	Time                *time.Time `json:"time"`
	PickupLocation      string     `json:"pickup_location"`
	DestinationLocation string     `json:"destination_location"`
	Status              int        `json:"status"`
}
