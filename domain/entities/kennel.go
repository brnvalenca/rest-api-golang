package entities

type Kennel struct {
	ID            int    `json:"id"`
	ContactNumber string `json:"contact_number"`
	Name          string `json:"name"`
	Dogs          []Dog  `json:"dogs"`
	Address       Address
}
