package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Promotion struct{
	PromotionID uint `json: "promotion_id"`
	Name string `json:"name"`
	DiscountType string `json:"discount_type"`
	DiscountValue float64 `json:"discount_value"`
	StartDate time.Time `json:"start_date"`
	EndDate time.Time `json:"end_date"` 
}


var promotion []Promotion

func CreatePromotionData(c echo.Context) error {
	var promo Promotion

	// Check Invalid Data Base On Promotion Struct
	if err := c.Bind(&promo); err != nil {
		return echo.NewHTTPError(http.StatusBadGateway, "Invalid Promotion Data") 
	}

	// Check Duplication Data Interatively
	for _, p := range promotions{
		if p.Promotion == promo.PromotionID && ((p.StartDate.Equal(promo.StartDate) || p.StartDate.Before(promo.StartDate)) && (p.EndDate.Equal(promo.EndDate) || p.EndDate.After(promo.EndDate))) {
			return echo.NewHTTPError(http.Status.Conflict, "Duplicate promotion found")
		}
	}

	// Append Data to Database
	promotion = append(promotions, promo)

	// Return Data Already Input
	return c.JSON(http.StatusCreated, promo)
}

func GetAllPromotionData(c echo.Context) error {
	return c.JSON(http.StatusOK, promotions)

}

func GetPromotionByID(c echo.Context) error {
	promotion_id, err := strconv.Atoi(c.Param("promtion_id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Invalid Promotion ID")
	}

	for _, promo := range promotions {
		if int(promo.PromotionID) == promotion_id {
			return c.JSON(http.StatusOK, promo)
		}
	}

	return echo.NewHTTPError(http.StatusNotFound, "Promotion Not Found")
}



func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
}