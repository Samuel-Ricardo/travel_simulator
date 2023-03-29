package route

import (
	"bufio"
	"encoding/json"
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

/*
* {"clientId":"1","routeId":"1"}
*/
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

  file, error := os.Open("destinations/"+ r.ID +".txt")
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

func (r *Route) ExportJsonPositions() ([]string, error) {

  var route PartialRoutePosition
  var result []string
  total := len(r.Positions)

  for key, value := range r.Positions {
    
    route.ID = r.ID
    route.ClientID = r.ClientID
    route.Position = []float64{value.Latitude, value.Longitude}
    route.Finished = false

    if total-1 == key { route.Finished = true } 

    jsonRoute, err := json.Marshal(route)
    if err !=  nil { return nil, err }

    result = append(result, string(jsonRoute))
  }
  
  return result, nil
}
