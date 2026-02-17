package car

type CarUpdateForm struct {
	ID           uint     `json:"id" binding:"required"`
	Brand        *string  `json:"brand"`
	Model        *string  `json:"model"`
	Color        *string  `json:"color"`
	LicensePlate *string  `json:"licensePlate"`
	Displacement *float64 `json:"displacement"`
	DriveType    *string  `json:"driveType"`
	Status       *int8    `json:"status"`
	DailyRent    *float64 `json:"dailyRent"`
	Mileage      *int64   `json:"mileage"`
	Description  *string  `json:"description"`
}
