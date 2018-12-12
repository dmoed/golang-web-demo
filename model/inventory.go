package model

import (
	"database/sql"
	"log"
	"time"
)

type Inventory struct {
	ID         string `json:"id"`
	CustomerID string `json:"customer_id"`
	ProductID  string `json:"product_id"`
	LocationID string `json:"location_id"`
	Quantity   string `json:"quantity"`
	Week       string `json:"week"`
	Year       string `json:"year"`
	Date       string `json:"date"`
}

const layout = "2006-01-02 15:04:05"

//GetInventoryByWeek action
func GetInventoryByWeekAndYear(db *sql.DB, week int, year int) ([]*Inventory, error) {

	results, err := db.Query("select id, customer_id, product_id, location_id, quantity, week, year, date from inventory where week = ? and year = ?", week, year)

	if err != nil {
		return nil, err
	}

	data := make([]*Inventory, 0)

	for results.Next() {

		inv := new(Inventory)

		err = results.Scan(&inv.ID, &inv.CustomerID, &inv.ProductID, &inv.LocationID, &inv.Quantity, &inv.Week, &inv.Year, &inv.Date)

		if err != nil {
			return nil, err
		}

		data = append(data, inv)
	}

	return data, err
}

func FilterByLatestPerWeek(list []*Inventory) []*Inventory {

	type Group struct {
		Index []int `json:"index"`
		Date  int64 `json:"date"`
	}

	GroupMap := make(map[string]*Group, 0)

	for k, v := range list {

		key := v.Week + "_" + v.ProductID + "_" + v.CustomerID + "_" + v.LocationID
		t, err := time.Parse(layout, v.Date)
		if err != nil {
			log.Fatal(err)
		}

		if element, ok := GroupMap[key]; ok == false {

			GroupMap[key] = &Group{[]int{k}, t.Unix()}

		} else {

			if element.Date > t.Unix() {
				element.Index = append(element.Index, k)
			}
		}
	}

	var payload []*Inventory

	for _, gm := range GroupMap {

		for _, ind := range gm.Index {

			payload = append(payload, list[ind])
		}
	}

	return payload
}

func GetInventoryProductSummary() {

}
