package body

type Mass struct {
	Name string  `json:"name"`
	Mass float64 `json:"mass"`
	X    float64 `json:"x"`
	Y    float64 `json:"y"`
	Z    float64 `json:"z"`
	Vx   float64 `json:"vx"`
	Vy   float64 `json:"vy"`
	Vz   float64 `json:"vz"`
}

type Config struct {
	Steps    int
	Interval float64
	Masses   []Mass
}
