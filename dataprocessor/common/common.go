package common

type CountryWrapper struct {
	Countries []Country `json:"countries"`
}

type StateWrapper struct {
	States []State `json:"states"`
}

type CityWrapper struct {
	Cities []City `json:"cities"`
}

type State struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CountryID string `json:"country_id"`
}

type Country struct {
	ID          string `json:"id"`
	CountryCode string `json:"sortname"`
	Name        string `json:"name"`
}

type City struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	StateID string `json:"state_id"`
}

type RowData struct {
	ID          string
	CountryName string
	StateName   string
	CityName    string
}
