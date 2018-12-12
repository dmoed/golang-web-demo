package handler

import (
	"apptastic/dashboard/model"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func InventoryHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var year, week = time.Now().ISOWeek()

		inventory, _ := model.GetInventoryByWeekAndYear(db, week, year)

		data, err := json.Marshal(inventory)

		if err != nil {
			panic(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	})
}

func InventoryTotalSummaryHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error
		var period = 12
		var weeks []string
		var payload []*model.Inventory
		var currentYear, currentWeek = time.Now().ISOWeek()

		params := r.URL.Query()
		var wk = params.Get("week")

		if wk != "" {

			currentWeek, err = strconv.Atoi(wk)

			if err != nil {
				panic(err.Error())
			}
		}

		for i := 0; i < period; i++ {

			var previousWeek = (currentWeek - i)
			var previousYear = currentYear

			if previousWeek <= 0 {

				previousWeek = previousWeek + 52
				previousYear = currentYear - 1
			}

			fmt.Println("prevWeek", previousYear, previousWeek)

			weeks = append(weeks, fmt.Sprintf("%v_%v", previousYear, previousWeek))
		}

		data, err := json.Marshal(payload)

		if err != nil {
			panic(err.Error())
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(data)
	})
}

func previousWeek(week int) int {

	var prevWeek int

	if week == 1 {

		prevWeek = 52

	} else {

		prevWeek = week - 1
	}

	return prevWeek
}

func WeekStart(year, week int) time.Time {
	// Start from the middle of the year:
	t := time.Date(year, 7, 1, 0, 0, 0, 0, time.UTC)

	// Roll back to Monday:
	if wd := t.Weekday(); wd == time.Sunday {
		t = t.AddDate(0, 0, -6)
	} else {
		t = t.AddDate(0, 0, -int(wd)+1)
	}

	// Difference in weeks:
	_, w := t.ISOWeek()
	t = t.AddDate(0, 0, (week-w)*7)

	return t
}
