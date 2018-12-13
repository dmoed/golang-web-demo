package model

import (
	"database/sql"
	"fmt"
	"time"
)

type Inventory struct {
	ID         string `json:"id"`
	CustomerID string `json:"customer_id"`
	ProductID  string `json:"product_id"`
	LocationID string `json:"location_id"`
	Quantity   int    `json:"quantity"`
	Week       string `json:"week"`
	Year       string `json:"year"`
	Date       string `json:"date"`
	PrTray     int    `json:"pr_tray"`
	Location   string `json:"location"`
	UUID       string `json:"uuid"`
}

type SummaryWeekProductCustomer struct {
	ID           string    `json:"id"`
	CustomerID   string    `json:"customer_id"`
	ProductID    string    `json:"product_id"`
	TotalTrays   float32   `json:"total_trays"`
	TotalBottles int       `json:"total_bottles"`
	Week         string    `json:"week"`
	Year         string    `json:"year"`
	YearWeek     string    `json:"year_week"`
	Date         time.Time `json:"date"`
}

type SummaryWeek struct {
	YearWeek     string   `json:"year_week"`
	Year         string   `json:"year"`
	Week         string   `json:"week"`
	Dates        []string `json:"dates"`
	Label        string   `json:"label"`
	TotalTrays   int      `json:"total_trays"`
	TotalBottles int      `json:"total_bottles"`
}

const layout = "2006-01-02 15:04:05"

//GetInventoryByWeekAndYear queries and returns inventory from year and week
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

//GetInventoryByDates queries and returns inventory from start to end date
func GetInventoryByDates(db *sql.DB, start time.Time, end time.Time) ([]*Inventory, error) {

	var startDate = start.Format(layout)
	var endDate = end.Format(layout)

	var sql = `select i.id, i.customer_id, i.product_id, i.location_id, 
				i.quantity, i.uuid, i.week, i.year, i.date, p.pr_tray,
				l.uuid as location_uuid
				from inventory as i
				LEFT JOIN inventory_location AS l ON l.id = i.location_id 
				LEFT JOIN product as p ON p.id = i.product_id
				where i.active is true 
				and i.date between ? and ?
				order by date desc`

	results, err := db.Query(sql, startDate, endDate)

	if err != nil {
		return nil, err
	}

	data := make([]*Inventory, 0)

	for results.Next() {

		inv := new(Inventory)

		err = results.Scan(&inv.ID, &inv.CustomerID, &inv.ProductID, &inv.LocationID, &inv.Quantity, &inv.UUID, &inv.Week, &inv.Year, &inv.Date, &inv.PrTray, &inv.Location)

		if err != nil {
			return nil, err
		}

		data = append(data, inv)
	}

	return data, err
}

//FilterByLatestPerWeek returns latest inventory per week
func FilterByLatestPerWeek(list []*Inventory) []*Inventory {

	type Group struct {
		Index []int `json:"index"`
		Date  int64 `json:"date"`
	}

	GroupMap := make(map[string]*Group, 0)

mainLoop:
	for k, v := range list {

		key := v.Week + "_" + v.ProductID + "_" + v.CustomerID + "_" + v.LocationID
		t, err := time.Parse(layout, v.Date)
		if err != nil {
			panic(err.Error())
		}

		if element, ok := GroupMap[key]; ok == false {

			GroupMap[key] = &Group{
				Index: []int{k},
				Date:  t.Unix(),
			}

		} else {

			if element.Date < t.Unix() {

				//replace all
				element.Index = []int{k}
				element.Date = t.Unix()

			} else if element.Date > t.Unix() {

				//do nothing

			} else {

				//check for duplicates
				for _, ind := range element.Index {

					if list[ind].UUID == v.UUID {
						continue mainLoop
					}
				}

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

//
func GetSummaryWeekProductCustomer(inventory []*Inventory) []*SummaryWeekProductCustomer {

	list := make(map[string]*SummaryWeekProductCustomer, 0)

	for _, v := range inventory {

		var key = fmt.Sprintf("%v_%v_%v_%v", v.Year, v.Week, v.CustomerID, v.ProductID)
		var qtyTrays float32
		var qtyBottles int

		switch v.Location {
		case "Magazijn":
			qtyTrays = float32(v.Quantity)
			qtyBottles = v.Quantity * v.PrTray
		default:
			qtyTrays = float32(v.Quantity) / float32(v.PrTray)
			qtyBottles = v.Quantity
		}

		if val, ok := list[key]; ok == false {

			list[key] = &SummaryWeekProductCustomer{
				ID:           v.ProductID,
				CustomerID:   v.CustomerID,
				ProductID:    v.ProductID,
				YearWeek:     key,
				Year:         v.Year,
				Week:         v.Week,
				Date:         time.Now(),
				TotalTrays:   qtyTrays,
				TotalBottles: qtyBottles,
			}

		} else {

			val.TotalTrays += qtyTrays
			val.TotalBottles += qtyBottles
		}
	}

	var payload []*SummaryWeekProductCustomer

	for _, s := range list {

		payload = append(payload, s)
	}

	return payload
}

//
func GetSummaryWeek(s []*SummaryWeekProductCustomer) []*SummaryWeek {

	SummaryWeekMap := make(map[string]*SummaryWeek, 0)

	for _, v := range s {

		var key = fmt.Sprintf("%v_%v", v.Year, v.Week)
		var totalTrays = int(v.TotalTrays + 0.5)
		var totalBottles = v.TotalBottles

		if val, ok := SummaryWeekMap[key]; ok == false {

			SummaryWeekMap[key] = &SummaryWeek{
				YearWeek:     key,
				Year:         v.Year,
				Week:         v.Week,
				Dates:        []string{"Y-m-d", "Y-m-d"},
				TotalTrays:   totalTrays,
				TotalBottles: totalBottles,
				Label:        fmt.Sprintf("WK%v", v.Week),
			}

		} else {

			val.TotalTrays += totalTrays
			val.TotalBottles += totalBottles
		}
	}

	var payload []*SummaryWeek

	for _, s := range SummaryWeekMap {

		payload = append(payload, s)
	}

	return payload
}

func GetInventoryProductSummary() {

	// SummaryMap := make(map[string]*Summary, 0)

	// for _, v := range inventory {

	// 	var key = fmt.Sprintf("%v_%v", v.Year, v.Week)
	// 	var qty int

	// 	switch v.Location {
	// 	case "Magazijn":
	// 		qty = v.Quantity
	// 	default:
	// 		qty = v.Quantity / v.PrTray
	// 	}

	// 	if val, ok := SummaryMap[key]; ok == false {

	// 		SummaryMap[key] = &Summary{
	// 			YearWeek: key,
	// 			Year:     v.Year,
	// 			Week:     v.Week,
	// 			Date:     time.Now(),
	// 			Total:    qty,
	// 		}

	// 	} else {

	// 		val.Total = val.Total + qty
	// 	}
	// }
}
