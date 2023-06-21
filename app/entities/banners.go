package entities

type Banners struct {
	ID        uint
	Title     string
	Thumbnail string
	Status    string
	Type      string
	Redirect  string
	CreatedAt int
	EndAt     int
	Segment   string
}
