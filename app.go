package app

type Shirt struct {
	Id        string
	Size      string
	PaintId   int
	PaperId   int
	TextileId int
}

type Paint struct {
	Id     int
	Type   string
	Color  string
	Volume float64
}

type Paper struct {
	Id   int
	Type string
	Size string
}

type Textile struct {
	Id   int
	Type string
}

type Printer struct {
	Id         int     `db:"id"`
	Status     string  `db:"status"`
	Efficiency float64 `db:"efficiency"`
	StartedAt  int64   `db:"started_at"`
	ExpiredAt  int64   `db:"expired_at"`
}
