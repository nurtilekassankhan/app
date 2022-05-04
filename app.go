package app

import "time"

type AvailableResources struct {
	PaintVolume   float64
	PaperVolume   float64
	TextileVolume float64
}

type RequiredResources struct {
	PaintVolume   float64
	PaperVolume   float64
	TextileVolume float64
	SpentTime     time.Duration
}

type Order struct {
	Id      int
	Status  string
	ShirtId int `db:"shirt_id" json:"shirt_id"`
	Amount  float64
}

type Shirt struct {
	Id            string
	Size          string
	PaintId       int     `db:"paint_id" json:"paint_id"`
	PaperId       int     `db:"paper_id" json:"paper_id"`
	TextileId     int     `db:"textile_id" json:"textile_id"`
	PaintVolume   float64 `db:"paint_volume" json:"paint_volume"`
	PaperVolume   float64 `db:"paper_volume" json:"paper_volume"`
	TextileVolume float64 `db:"textile_volume" json:"textile_volume"`
}

type Paint struct {
	Id     int
	Type   string
	Color  string
	Volume float64
}

type Paper struct {
	Id     int
	Type   string
	Volume float64
}

type Textile struct {
	Id     int
	Type   string
	Volume float64
}

type Printer struct {
	Id         int     `db:"id"`
	Status     string  `db:"status"`
	Efficiency float64 `db:"efficiency"` //производительность футболка/час
	StartedAt  int64   `db:"started_at"`
	ExpiredAt  int64   `db:"expired_at"`
}
