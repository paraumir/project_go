package models

type Item struct {
	ID     int     `json:"id"`
	Title  string  `json:"title"`
	Flavor string  `json:"flavor"`
	Rating float64 `json:"rating"`
	Image  string  `json:"image"`
}
