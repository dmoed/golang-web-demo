package handler

import (
	"apptastic/dashboard/auth"
	"database/sql"
	"encoding/json"
	"html/template"
	"net/http"
)

func DashboardHandler(db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//get user here
		user, err := auth.GetUser(r, db)

		if err != nil {
			panic(err.Error())
		}

		react := &react{
			InitialState: reactState{
				User: UserProps{
					ID:            user.ID,
					Username:      user.Username,
					Email:         user.Email,
					Displayname:   user.Username,
					ProfileImage:  "",
					Roles:         []string{"ROLE_USER"},
					Authenticated: true,
				},
			},
			Props: reactProps{
				Config: reactPropsConfig{
					BaseUrl:    "localhost",
					AppName:    "GOLANG",
					AppVersion: "0.0.1",
					AppLogo:    "",
				},
			},
		}

		encoded, _ := json.Marshal(react)
		myJSON := template.JS(encoded)

		t, _ := template.ParseFiles("templates/react.html")
		t.Execute(w, map[string]interface{}{
			"json": myJSON,
		})
	})
}

type react struct {
	InitialState reactState `json:"__initial_state__"`
	Props        reactProps `json:"__props__"`
}

type reactState struct {
	User UserProps `json:"user"`
}

type reactProps struct {
	Config reactPropsConfig `json:"config"`
}

type reactPropsConfig struct {
	BaseUrl    string `json:"base_url"`
	Host       string `json:"host"`
	Scheme     string `json:"scheme"`
	Routes     string `json:"routes"`
	AppName    string `json:"app_name"`
	AppLogo    string `json:"app_logo"`
	AppVersion string `json:"app_version"`
}

type UserProps struct {
	ID            int      `json:"id"`
	Username      string   `json:"username"`
	Displayname   string   `json:"displayname"`
	Email         string   `json:"email"`
	ProfileImage  string   `json:"profile_image"`
	Roles         []string `json:"roles"`
	Authenticated bool     `json:"authenticated"`
}
