package pkg

type Criteria struct {
	Checkin     Date
	ChecOut     Date
	Destination string
	NumPaxes    int
}


type Response struct {
	Errors []Error
}
type Error struct {
	Code string
	Description string
}

type DestinationSearcherResponse struct {
	Code      string
	Available bool
}

type SearchResponse struct {
	Response
	OptionID   string
	HotelName  string
	Amount     float32
	Currency   string
	Refundable bool
}

type BookResponse struct {
	Response
	BookingID string
}

type Service interface {
	Search(input Criteria) SearchResponse
	Book(optionID string) BookResponse
}
