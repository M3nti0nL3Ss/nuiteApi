package models

type Response struct {
	Data []hotel `json:"data"`
}

type hotel struct {
	HotelId  int     `json:"hotelId"`
	Currency string  `json:"currency"`
	Price    float32 `json:"price"`
}

type HotelParser struct {
	Code     int    `json:"code"`
	MinRate  string `json:"minRate"`
	Currency string `json:"currency"`
}

func (h HotelParser) Error() string {
	panic("HotelParser error Handler")
}

type HotelsData struct {
	Hotels struct {
		Hotels []HotelParser `json:"hotels"`
	} `json:"hotels"`
}

type HotelsResult struct {
	Data []struct {
		HotelId  string  `json:"hotelId"`
		Currency string  `json:"currency"`
		Price    float64 `json:"price"`
	} `json:"data"`
}
