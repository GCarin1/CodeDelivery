package route

type route struct {
	ID        string
	ClientID  string
	Positions []Positions
}
type Positions struct {
	lat  float64
	long float64
}
