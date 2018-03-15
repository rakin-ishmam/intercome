package invite

// Customer holds customer info
type Customer struct {
	Lat    float64 `json:"latitude"`
	Lng    float64 `json:"longitude"`
	UserID int     `json:"user_id"`
	Name   string  `json:"name"`
}
