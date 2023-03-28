package route

type Route struct {
	ID        string     `json:"routeId"`
	ClientID  string     `json:"clientID"`
	Positions []Position `json:"positions"`
}

type Position struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type PartialRoutePosition struct {
	ID       string    `json:"routeId"`
	ClientID string    `json:"clientId"`
	Position []float64 `json:"position"`
	Finished bool      `json:"finished"`
}
