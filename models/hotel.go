package models

// Hotel defines the structure for a hotel API
type Hotel struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Hotelier       int    `json:"hotelier_id"`
	Rating         int    `json:"rating"`
	Country        string `json:"country"`
	Address        string `json:"address"`
	RoomsAvailable int    `json:"rooms_available"`
	CreatedOn      string `json:"-"`
	UpdatedOn      string `json:"-"`
	DeletedOn      string `json:"-"`
}
