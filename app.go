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
	ShirtId int
	Amount  float64
}

type Shirt struct {
	Id            string
	Size          string
	PaintId       int
	PaperId       int
	TextileId     int
	PaintVolume   float64
	PaperVolume   float64
	TextileVolume float64
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
