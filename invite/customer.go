package invite

// Customer holds customer info
type Customer struct {
	Location
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
}
