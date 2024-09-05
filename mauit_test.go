package main

import (
	"encoding/json"
	"fmt"
	"mauit/models"
	"net/http"
	"strings"
	"testing"
)

var Dieters []models.Dieter
var Entries []models.Entry

func AddDieters() {

    fmt.Printf("Adding Dieters")
    newdieter := "{\"ID\": 1, \"Name\": \"Matt\", \"Calories\": 1600}"

    resp, err := http.Post("http://localhost:9090/dieters/add", "application/json", strings.NewReader(newdieter)) 

    if err != nil {
        return
    }

    defer resp.Body.Close()
    dec := json.NewDecoder(resp.Body)
    var success models.Dieter
    err = dec.Decode(&success)
    Dieters = append(Dieters, success)

}

func AddEntries() {
    fmt.Printf("Adding Entries")
    Entries = []models.Entry {
        {ID: 1, FoodID:1, MealID:1, Calories:100}, 
        {ID: 2, FoodID:1, MealID:1, Calories:100}, 
        {ID: 3, FoodID:2, MealID:1, Calories:150}, 
        {ID: 1, FoodID:3, MealID:2, Calories:300}, 
    }
}

func TestMauit(t *testing.T) {

    AddDieters()
    got := Dieters 
    want := Dieters != nil 

    if !want {
        t.Errorf("got %v, wanted a list of dieters", got)
    }

}

