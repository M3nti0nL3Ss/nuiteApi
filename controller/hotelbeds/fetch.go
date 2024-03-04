package hotelbeds

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	. "nuiteApi/models"
	"os"
	"strconv"
)

func formatHotelBedsResponse(data []byte) (interface{}, error) {
	// fmt.Println(data)
	var hotels HotelsData
	err := json.Unmarshal(data, &hotels)
	if err != nil {
		fmt.Println("Error decoding hotel data:", err)
		return nil, err
	}

	var result HotelsResult
	for _, hotel := range hotels.Hotels.Hotels {
		price, _ := strconv.ParseFloat(hotel.MinRate, 64)
		result.Data = append(result.Data, struct {
			HotelId  string  `json:"hotelId"`
			Currency string  `json:"currency"`
			Price    float64 `json:"price"`
		}{
			HotelId:  fmt.Sprintf("%d", hotel.Code),
			Currency: hotel.Currency,
			Price:    price,
		})
	}

	return result, err
}

func GetJson(jsonStr []byte) (int, interface{}) {
	url := os.Getenv("ENDPOINT")

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Set("Api-key", os.Getenv("API_KEY"))
	req.Header.Set("X-Signature", GenerateXSignature())

	//res, err := httputil.DumpRequest(req, true)
	//if err != nil {
	//	log.Fatal(err)
	//}
	// fmt.Println(string(res))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	res, err := formatHotelBedsResponse(body)
	if err != nil {
		return 500, Error{Error: err}
	}
	if resp.StatusCode == 401 {
		return resp.StatusCode, ErrorMessage{Message: "Quota exceeded"}
	}
	return resp.StatusCode, res
}
