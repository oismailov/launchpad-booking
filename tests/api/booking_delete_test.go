package api_test

import (
	"encoding/json"
	"fmt"
	"github.com/oismailov/launchpad-booking/application"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDeleteBooking_Success(t *testing.T) {
	defer RestoreTestDatabaseState()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/api/bookings/%s", bookingTestUUID), nil)

	router.ServeHTTP(w, req)

	var response Response

	if err = json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Error(err)
	}

	assert.Equal(t, http.StatusOK, w.Code)

	bookings, err := application.GetAllBookings()

	if err != nil {
		t.Error(err)
	}

	assert.Empty(t, bookings)
	assert.NotEmpty(t, response)

	assert.Equal(t, "booking has been deleted", response.Message)
}

func TestDeleteBooking_InvalidBookingUUID(t *testing.T) {
	defer RestoreTestDatabaseState()
	invalidUUID := "abc"

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", fmt.Sprintf("/api/bookings/%s", invalidUUID), nil)

	router.ServeHTTP(w, req)

	var response Response

	if err = json.Unmarshal(w.Body.Bytes(), &response); err != nil {
		t.Error(err)
	}

	assert.Equal(t, http.StatusBadRequest, w.Code)

	bookings, err := application.GetAllBookings()

	if err != nil {
		t.Error(err)
	}

	assert.NotEmpty(t, bookings)
	assert.NotEmpty(t, response)

	assert.Equal(t, "invalid booking UUID", response.Message)
}
