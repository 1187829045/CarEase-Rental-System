package _struct

type SendSMSForm struct {
	Mobile string `form:"mobile" json:"mobile" binding:"required,len=11"`
}

type CarListQuery struct {
	Brand    string   `form:"brand"`
	Model    string   `form:"model"`
	Status   *int8    `form:"status"`
	MinPrice *float64 `form:"minPrice"`
	MaxPrice *float64 `form:"maxPrice"`
}

type CarUpdateForm struct {
	ID           uint     `json:"id" binding:"required"`
	Brand        *string  `json:"brand"`
	Model        *string  `json:"model"`
	Color        *string  `json:"color"`
	LicensePlate *string  `json:"licensePlate"`
	Seats        *int8    `json:"seats"`
	FuelType     *string  `json:"fuelType"`
	Displacement *float64 `json:"displacement"`
	DriveType    *string  `json:"driveType"`
	Status       *int8    `json:"status"`
	DailyRent    *float64 `json:"dailyRent"`
	Mileage      *int64   `json:"mileage"`
	Description  *string  `json:"description"`
}
