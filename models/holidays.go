package models

type Holiday struct {
	ChName string `json:"ch_name"`
	EnName string `json:"en_name"`
	Date   string `json:"date"`
	Count  int    `json:"count"`
}

type Holidays []Holiday
