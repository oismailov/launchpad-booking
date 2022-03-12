package api_test

import "github.com/oismailov/launchpad-booking/model"

const bookingTestUUID = "ad7fed89-b1f6-426c-ae7e-1a5562295671"

type bookingDataProvider struct {
	booking        model.Booking
	errorMessage   string
	httpStatusCode int
}
