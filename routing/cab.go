package routing

import (
	"GoCab/controller"
	"GoCab/model"
	"errors"
	"fmt"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

// Cab router
type Cab struct {
	cabController controller.Cab
}

// NewCabRouter will create new cab router
func NewCabRouter(
	cabController controller.Cab,

) Cab {
	return Cab{

		cabController: cabController,
	}

}

// Register will create new routes
func (router Cab) Register(group *echo.Group) {
	group.POST("/near-by-cab", router.nearBycab)
	group.POST("/book-cab", router.bookCab)
	group.POST("/booking-status", router.updatebookingStatus)
	group.GET("/booking-history", router.historyOfBooking)

}

func (router Cab) nearBycab(context echo.Context) error {
	type Request struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}
	req := new(Request)
	// bind request to context
	if err := context.Bind(req); err != nil {
		log.Error(err)
		return errors.New("Invalid Request")
	}
	// create user in database
	cabList, err := router.cabController.FindNearByCab(req.Latitude, req.Longitude)
	if err != nil {
		return err
	}
	// return response
	return context.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "success",
		"cabList": cabList,
	})

}

func (router Cab) bookCab(context echo.Context) error {

	req := new(model.BookingRequest)
	// bind request to context
	if err := context.Bind(req); err != nil {
		log.Error(err)
		return errors.New("Invalid Request")
	}
	user := context.Get("User").(*model.User)
	fmt.Println(user)
	req.UserID = user.ID
	// create user in database
	err := router.cabController.BookCab(req)
	if err != nil {
		return err
	}
	// return response
	return context.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "success",
		"Message": "cab is on the way to pick you",
	})

}

func (router Cab) updatebookingStatus(context echo.Context) error {

	type Request struct {
		Status    int  `json:"status"`
		BookingID uint `json:"booking_id"`
	}

	req := new(Request)
	// bind request to context
	if err := context.Bind(req); err != nil {
		log.Error(err)
		return errors.New("Invalid Request")
	}

	// create user in database
	err := router.cabController.UpdateBookingStatus(req.Status, req.BookingID)
	if err != nil {
		return err
	}
	// return response
	return context.JSON(http.StatusOK, map[string]interface{}{
		"Status":  "success",
		"Message": "Your Booking status is updated",
	})

}

func (router Cab) historyOfBooking(context echo.Context) error {
	user := context.Get("User").(*model.User)

	// create user in database
	bookingList, err := router.cabController.HistoryOfBooking(user.ID)
	if err != nil {
		return err
	}
	// return response
	return context.JSON(http.StatusOK, map[string]interface{}{
		"Status":      "success",
		"BookingList": bookingList,
	})

}
