package api_test

import (
	"encoding/json"
	"github.com/oismailov/launchpad-booking/model"
	"github.com/oismailov/launchpad-booking/persistance"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllBookings_Success(t *testing.T) {
	defer RestoreTestDatabaseState()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/bookings", nil)

	router.ServeHTTP(w, req)

	var response []model.Booking

	if err = json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Error(err)
	}

	assert.Equal(t, http.StatusOK, w.Code)

	bookings, err := persistance.GetAllBookings()

	if err != nil {
		t.Error(err)
	}

	assert.NotEmpty(t, bookings)
	assert.NotEmpty(t, response)

	assert.Equal(t, 1, len(bookings))
	assert.Equal(t, 1, len(response))

	bookingEntity := bookings[0]
	bookingFromResponse := response[0]

	assert.Equal(t, bookingEntity.UUID, bookingFromResponse.UUID)
	assert.Equal(t, bookingEntity.LaunchDate, bookingFromResponse.LaunchDate)
	assert.Equal(t, bookingEntity.DestinationID, bookingFromResponse.DestinationID)
	assert.Equal(t, bookingEntity.LaunchDate, bookingFromResponse.LaunchDate)
	assert.Equal(t, bookingEntity.Birthday, bookingFromResponse.Birthday)
	assert.Equal(t, bookingEntity.FirstName, bookingFromResponse.FirstName)
	assert.Equal(t, bookingEntity.LastName, bookingFromResponse.LastName)
	assert.Equal(t, bookingEntity.Gender, bookingFromResponse.Gender)
}

func TestGetAllBookings_EmptyResult(t *testing.T) {
	defer RestoreTestDatabaseState()

	err := persistance.DeleteBookingByUUID(bookingTestUUID)
	if err != nil {
		t.Error(err)
	}

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/bookings", nil)
	router.ServeHTTP(w, req)

	var response []model.Booking

	if err = json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Error(err)
	}

	assert.Equal(t, http.StatusOK, w.Code)

	bookings, err := persistance.GetAllBookings()

	if err != nil {
		t.Error(err)
	}

	assert.Empty(t, bookings)
	assert.Empty(t, response)
}
