package route

type Route struct {
  ID        string      `json:"routeId"`
  ClientID  string      `json:"clientID"`
  Positions []Position  `json:"positions"` 
}




