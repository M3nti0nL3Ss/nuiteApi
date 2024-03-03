package controller

import (
	"encoding/json"
	. "nuiteApi/controller/hotelbeds"
	"strings"

	"github.com/gin-gonic/gin"
	. "nuiteApi/models"
	"strconv"
)

func GetHotels(c *gin.Context) {
	params := c.Request.URL.Query()
	stay := &Stay{CheckIn: params["checkin"][0], CheckOut: params["checkout"][0]}
	var occupancies []Occupancy
	_ = json.Unmarshal([]byte(params["occupancies"][0]), &occupancies)

	hotelIdsStr := strings.Split(params["hotelIds"][0], ",")
	hotelIds := make([]int, len(hotelIdsStr))
	for i := range hotelIds {
		hotelIds[i], _ = strconv.Atoi(hotelIdsStr[i])
	}
	hotels := Hotels{Hotel: hotelIds}

	req, _ := json.Marshal(&Request{Stay: *stay, Occupancies: occupancies, Hotels: hotels})
	c.JSON(GetJson(req))
}
