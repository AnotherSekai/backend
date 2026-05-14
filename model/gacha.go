package model

type Gacha struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	AbName string `json:"assetbundleName"`
}

type GachaResponse struct {
	Data []Gacha `json:"data"`
}
