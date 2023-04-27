package model

type SearchDto struct {
	From          string `json:"from"`
	To            string `json:"to"`
	DepartureDate string `json:"departure_date"`
}

type SearchResultDto struct {
	ID            int    `json:"id"`
	DepartureTime string `json:"departure_time"`
	ArrivalTime   string `json:"arrival_time"`
	AirlineName   string `json:"airline_name"`
	Price         int64  `json:"price"`
	DepartureDate string `json:"departure_date"`
}

type FlightDetailsDto struct {
	ID            int    `json:"id"`
	AirlineName   string `json:"Airline_name"`
	From          string `json:"Departure"`
	FromCode      string `json:"Departure_code"`
	To            string `json:"Arrival"`
	ToCode        string `json:"Arrival_code"`
	DepartureDate string `json:"Departure_date"`
	ReturnDate    string `json:"Return_date"`
}
