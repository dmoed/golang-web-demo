package handler

import (
	"apptastic/dashboard/model"
	"apptastic/dashboard/monitor"
	"apptastic/dashboard/view"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func InventoryHandler(db *sql.DB, v *view.View, logger *log.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		monitor.PrintMemUsage()

		yr, wk := time.Now().ISOWeek()
		year := getQueryParamInt(r, "year", yr)
		week := getQueryParamInt(r, "week", wk)

		start := model.WeekStart(year, week)
		end := start.AddDate(0, 0, 6)

		inventory, err := model.GetInventoryByDates(db, start, end)

		if err != nil {
			panic(err.Error())
		}

		data, err := json.Marshal(map[string]interface{}{
			"payload": inventory,
		})

		monitor.PrintMemUsage()

		if err != nil {
			panic(err.Error())
		}

		v.RenderJSON(w, data)
	})
}

func InventoryTotalSummaryHandler(db *sql.DB, v *view.View, logger *log.Logger) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		layout := "Mon 2 Jan"
		yr, wk := time.Now().ISOWeek()
		year := getQueryParamInt(r, "year", yr)
		week := getQueryParamInt(r, "week", wk)

		period := 26
		end := model.WeekStart(year, week).AddDate(0, 0, 6)
		start := end.AddDate(0, 0, -(period*7)+1)
		weeks := model.GetWeeks(year, week, period)

		fmt.Println("start", start)
		fmt.Println("end", end)

		inventory, err := model.GetInventoryByDates(db, start, end)

		if err != nil {
			panic(err.Error())
		}

		monitor.PrintMemUsage()

		summary := model.GetSummaryWeek(model.GetSummaryWeekProductCustomer(model.FilterByLatestPerWeek(inventory)))
		monitor.PrintMemUsage()

		for _, v := range weeks {

			hasWeek := false

			ywk := fmt.Sprintf("%v_%v", v.Year, v.Week)

			for _, summ := range summary {

				if ywk == summ.YearWeek {
					hasWeek = true
				}
			}

			if hasWeek == false {
				summary = append(summary, &model.SummaryWeek{
					YearWeek:     ywk,
					Year:         strconv.Itoa(v.Year),
					Week:         strconv.Itoa(v.Week),
					Dates:        []string{v.StartDate.Format(layout), v.EndDate.Format(layout)},
					TotalTrays:   0,
					TotalBottles: 0,
					Label:        fmt.Sprintf("WK %d", v.Week),
				})
			}
		}

		data, err := json.Marshal(map[string]interface{}{
			"weeks":   weeks,
			"start":   start,
			"end":     end,
			"payload": summary,
		})

		if err != nil {
			panic(err.Error())
		}

		v.RenderJSON(w, data)
	})
}

func getQueryParamString(r *http.Request, key string, falllback string) string {

	u := r.URL.Query()

	if val := u.Get(key); val != "" {

		return val
	}

	return falllback
}

func getQueryParamInt(r *http.Request, key string, falllback int) int {

	val := r.URL.Query().Get(key)

	if newval, err := strconv.Atoi(val); err == nil {

		return newval
	}

	return falllback
}
