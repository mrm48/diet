package main

import (
	"encoding/json"
	"mauit/models"
	"net/http"
	"strings"
	"testing"
)

var Dieters     []models.Dieter
var Entries     []models.Entry
var newDieters  []models.Dieter
var newDieter   models.Dieter


func AddDieters() {

    newDieter.Name = "Matt"
    newDieter.ID = 1
    newDieter.Calories = 1600
    newDieters = append(newDieters, newDieter)
    newDieter.Name = "Jack"
    newDieter.ID = 2
    newDieter.Calories = 1600
    newDieters = append(newDieters, newDieter)

    for _, v := range newDieters {
        addDieter, err := json.Marshal(v)

        if err != nil {
            return
        }
        resp, err := http.Post("http://localhost:9090/dieters/add", "application/json", strings.NewReader(string(addDieter)))

        if err != nil {
            return
        }

        defer resp.Body.Close()
        dec := json.NewDecoder(resp.Body)
        var success models.Dieter
        err = dec.Decode(&success)
        Dieters = append(Dieters, success)
    }

}

func QueryDieters() {

    resp, err := http.Get("http://localhost:9090/dieters/all") 

    if err != nil {
        return
    }

    defer resp.Body.Close()

    dec := json.NewDecoder(resp.Body)
    var success []models.Dieter
    err = dec.Decode(&success)
    Dieters = success

}

func AddEntries() {
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
        t.Errorf("got %v, wanted a new dieter", got)
    } 

}

func TestQuery(t *testing.T) {
    QueryDieters()

    got := Dieters
    want := newDieters 

    if got == nil {
        t.Errorf("got %v, wanted a list of dieters", got)
    } else {
        for k, v := range got { 
            if v.Name != want[k].Name || v.ID != want[k].ID || v.Calories != want[k].Calories {
                t.Errorf("got %v, wanted \"%v\"; got %v, wanted %v; got %v, wanted %v", v.Name, want[k].Name, v.ID, want[k].ID, v.Calories, want[k].Calories)
            }
        }
    } 
}

