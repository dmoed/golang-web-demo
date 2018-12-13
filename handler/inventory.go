package handler

import (
	"apptastic/dashboard/model"
	"apptastic/dashboard/view"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

func InventoryHandler(db *sql.DB, v *view.View) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var year, week = time.Now().ISOWeek()

		inventory, _ := model.GetInventoryByWeekAndYear(db, week, year)
		latest := model.FilterByLatestPerWeek(inventory)

		data, err := json.Marshal(latest)

		if err != nil {
			panic(err.Error())
		}

		v.RenderJSON(w, data)
	})
}

func InventoryTotalSummaryHandler(db *sql.DB, v *view.View) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var err error
		var period = 5
		//var payload []*model.Inventory
		var currentTime = time.Now()
		var currentYear, currentWeek = currentTime.ISOWeek()
		weeks := make(map[string]map[string]string)

		//params := r.URL.Query()
		//var wk = params.Get("week")
		// if wk != "" {
		// 	currentWeek, err = strconv.Atoi(wk)
		// 	if err != nil {
		// 		panic(err.Error())
		// 	}
		// }

		for i := 0; i < period; i++ {

			var previousWeek = (currentWeek - i)
			var previousYear = currentYear

			if previousWeek <= 0 {
				previousWeek = previousWeek + period
				previousYear = currentYear - 1
			}

			var key = fmt.Sprintf("%v_%v", previousYear, previousWeek)
			weeks[key] = map[string]string{
				"year": strconv.Itoa(previousYear),
				"week": strconv.Itoa(previousWeek),
			}
		}

		start := time.Date(2018, 1, 1, 0, 0, 0, 0, time.UTC)
		end := currentTime

		fmt.Println("start", start)
		fmt.Println("end", end)

		inventory, err := model.GetInventoryByDates(db, start, end)

		if err != nil {
			panic(err.Error())
		}

		latest := model.FilterByLatestPerWeek(inventory)
		summary1 := model.GetSummaryWeekProductCustomer(latest)
		summary2 := model.GetSummaryWeek(summary1)

		for wky, v := range weeks {

			var hasWeek = false

			for _, summ := range summary2 {

				if wky == summ.YearWeek {
					hasWeek = true
				}
			}

			if hasWeek == false {
				summary2 = append(summary2, &model.SummaryWeek{
					YearWeek:     wky,
					Year:         v["year"],
					Week:         v["week"],
					Dates:        []string{"Y-m-d", "Y-m-d"},
					TotalTrays:   0,
					TotalBottles: 0,
					Label:        fmt.Sprintf("WK%v", v["week"]),
				})
			}
		}

		data, err := json.Marshal(map[string]interface{}{
			"weeks":   weeks,
			"payload": summary2,
		})

		if err != nil {
			panic(err.Error())
		}

		v.RenderJSON(w, data)
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
