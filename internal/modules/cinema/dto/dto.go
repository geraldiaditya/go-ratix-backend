package dto

import "github.com/geraldiaditya/ratix-backend/internal/modules/cinema/domain"

type CityResponse struct {
	Cities []string `json:"cities"`
}

type BrandResponse struct {
	Brands []string `json:"brands"`
}

type CinemaResponse struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	Brand      string  `json:"brand"`
	City       string  `json:"city"`
	Address    string  `json:"address"`
	Rating     float64 `json:"rating"`
	Lat        float64 `json:"lat"`
	Lon        float64 `json:"lon"`
	PictureURL string  `json:"picture_url"`
	Distance   float64 `json:"distance_km,omitempty"`
}

type SeatLayoutResponse struct {
	Layout SeatLayout `json:"layout"`
	Legend SeatLegend `json:"legend"`
}

type SeatLayout struct {
	Rows  int    `json:"rows"`
	Cols  int    `json:"cols"`
	Seats []Seat `json:"seats"`
}

type Seat struct {
	Row    string  `json:"row"`
	Number int     `json:"number"`
	Status string  `json:"status"` // available, occupied, selected
	Type   string  `json:"type"`   // standard, premium
	Price  float64 `json:"price"`
}

type SeatLegend struct {
	Available string `json:"available"`
	Occupied  string `json:"occupied"`
	Selected  string `json:"selected"`
}

func ToCinemaResponse(c domain.Cinema) CinemaResponse {
	return CinemaResponse{
		ID:         c.ID,
		Name:       c.Name,
		Brand:      c.Brand,
		City:       c.City,
		Address:    c.Address,
		Rating:     c.Rating,
		Lat:        c.Lat,
		Lon:        c.Lon,
		PictureURL: c.PictureURL,
		Distance:   c.Distance,
	}
}
