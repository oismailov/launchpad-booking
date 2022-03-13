package api_test

import (
	"encoding/json"
	"github.com/oismailov/launchpad-booking/application"
	"github.com/oismailov/launchpad-booking/model"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateBooking_Success(t *testing.T) {
	defer RestoreTestDatabaseState()

	for k, booking := range CreateBookingDataProviderWithSuccessCases() {
		w := httptest.NewRecorder()

		b, err := json.Marshal(booking)

		if err != nil {
			t.Error(err)
		}

		req, _ := http.NewRequest("POST", "/api/bookings", strings.NewReader(string(b)))

		router.ServeHTTP(w, req)

		var response model.Booking

		if err = json.Unmarshal(w.Body.Bytes(), &response); err != nil {
			t.Error(err)
		}

		assert.Equal(t, http.StatusCreated, w.Code)

		bookings, err := application.GetAllBookings()

		if err != nil {
			t.Error(err)
		}

		assert.NotEmpty(t, bookings)
		assert.NotEmpty(t, response)

		assert.Equal(t, k+2, len(bookings))
	}
}

func CreateBookingDataProviderWithSuccessCases() []model.Booking {
	return []model.Booking{
		{
			UUID:          "b6ff1be9-26c5-470f-9abd-67b484fb86d6",
			FirstName:     "Sarah",
			LastName:      "Connor",
			Gender:        "female",
			Birthday:      "1984-01-01",
			LaunchpadID:   "5fa1f220-12fc-439b-8127-6598432eee0e",
			DestinationID: "dfa32bd1-5ddd-4052-b427-df344d3bbd65",
			LaunchDate:    "2029-01-01",
		}, {
			UUID:          "7e78b00a-39ba-4d5f-9914-b7509d92961b",
			FirstName:     "John",
			LastName:      "Smith",
			Gender:        "male",
			Birthday:      "2020-01-01",
			LaunchpadID:   "5fa1f220-12fc-439b-8127-6598432eee0e",
			DestinationID: "9c71e341-ec77-43d1-bd73-e1e0ab6be7f0",
			LaunchDate:    "2022-01-01",
		},
	}
}

func TestCreateBooking_WithInvalidRequestBody(t *testing.T) {
	defer RestoreTestDatabaseState()

	for _, payload := range CreateBookingDataProviderWithInvalidRequestBody() {
		w := httptest.NewRecorder()

		b, err := json.Marshal(payload.booking)

		if err != nil {
			t.Error(err)
		}

		req, _ := http.NewRequest("POST", "/api/bookings", strings.NewReader(string(b)))

		router.ServeHTTP(w, req)

		var response Response

		if err = json.Unmarshal(w.Body.Bytes(), &response); err != nil {
			t.Error(err)
		}

		assert.Equal(t, payload.httpStatusCode, w.Code)

		bookings, err := application.GetAllBookings()

		if err != nil {
			t.Error(err)
		}

		assert.NotEmpty(t, bookings)
		assert.NotEmpty(t, response)

		assert.Equal(t, 1, len(bookings))
		assert.Equal(t, payload.errorMessage, response.Message)
	}
}

func CreateBookingDataProviderWithInvalidRequestBody() []bookingDataProvider {
	return []bookingDataProvider{
		{
			model.Booking{
				UUID:          "b6ff1be9-26c5-470f-9abd-67b484fb86d6",
				FirstName:     "Sarah",
				LastName:      "Connor",
				Gender:        "female",
				Birthday:      "1984-01-01",
				LaunchpadID:   "e3ddfea9-bba6-4560-9cf7-3ff88abdfcc2",
				DestinationID: "dfa32bd1-5ddd-4052-b427-df344d3bbd65",
				LaunchDate:    "2029-01-01",
			},
			"unable to create a booking", //invalid launchpad id
			http.StatusBadRequest,
		},
		{
			model.Booking{
				UUID:          "7e78b00a-39ba-4d5f-9914-b7509d92961b",
				FirstName:     "John",
				LastName:      "Smith",
				Gender:        "male",
				Birthday:      "2020-01-01",
				LaunchpadID:   "5fa1f220-12fc-439b-8127-6598432eee0e",
				DestinationID: "9c71e341-ec77-43d1-bd73-e1e0ab6be7f0",
				LaunchDate:    "2025-01-01",
			},
			"unable to create a booking", //duplicate with local booking
			http.StatusBadRequest,
		},
		{
			model.Booking{
				UUID:          "b86d488d-e4d4-48ef-ac09-202298e77f92",
				FirstName:     "Mike",
				LastName:      "Doe",
				Gender:        "male",
				Birthday:      "2020-01-01",
				LaunchpadID:   "5fa1f220-12fc-439b-8127-6598432eee0e",
				DestinationID: "ac46ecaf-39c6-4109-af98-814247262cdb",
				LaunchDate:    "2023-01-01",
			},
			"unable to create a booking", //duplicate with SpaceX launches
			http.StatusBadRequest,
		},
		{
			model.Booking{
				UUID:          "b86d488d-e4d4-48ef-ac09-202298e77f92",
				FirstName:     "Mike",
				LastName:      "Doe",
				Gender:        "male",
				Birthday:      "2020-01-0",
				LaunchpadID:   "5fa1f220-12fc-439b-8127-6598432eee0e",
				DestinationID: "ec1a1df6-f581-4133-a567-658942a3094a",
				LaunchDate:    "2025-01-01",
			},
			"invalid birthday",
			http.StatusBadRequest,
		},
		{
			model.Booking{
				UUID:          "b86d488d-e4d4-48ef-ac09-202298e77f92",
				FirstName:     "Mike",
				LastName:      "Doe",
				Gender:        "male",
				Birthday:      "2020-01-01",
				LaunchpadID:   "5fa1f220-12fc-439b-8127-6598432eee0e",
				DestinationID: "ec1a1df6-f581-4133-a567-658942a3094a",
				LaunchDate:    "2025-01-0",
			},
			"invalid launch_date",
			http.StatusBadRequest,
		},
	}
}
