package route

import (
	"bufio"
	"errors"
	"os"
	"strconv"
	"strings"
)

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

func NewRoute() *Route {
	return &Route{}
}

func (r *Route) LoadPositions() error {
	
  if r.ID == "" {	return errors.New("route id not informed")	}

  file, error := os.Open("destinatios/"+ r.ID +".txt")
  if error != nil { return error }

  defer file.Close()

  scanner := bufio.NewScanner(file)
  for scanner.Scan() {

    data := strings.Split(scanner.Text(), ",")
    
    latitude, error := strconv.ParseFloat(data[0], 64)
    if error != nil { return nil }
    
    longitude, err := strconv.ParseFloat(data[1], 64)
    if err != nil { return nil }

    r.Positions = append(
      r.Positions, 
      Position{
      	Latitude:  latitude,
      	Longitude: longitude,
      })
  }

  return nil
}
