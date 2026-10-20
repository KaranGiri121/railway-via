package models

type Train struct {
	ID              string `json:"id"`
	Name            string `json:"name"`
	Origin          string `json:"origin"`
	Destination     string `json:"destination"`
	OriginCode      string `json:"origin_code"`
	DestinationCode string `json:"destination_code"`
}

func TrainFromMap(data map[string]interface{}) Train {
	return Train{
		ID:              data["train_number"].(string),
		Name:            data["train_name"].(string),
		OriginCode:      data["train_src"].(string),
		DestinationCode: data["train_dstn"].(string),
		Origin:          data["src_name"].(string),
		Destination:     data["dstn_name"].(string),
	}
}

type Station struct {
	ID          string `json:"id"`
	TrainId     string `json:"train_id"`
	StationCode string `json:"station_code"`
	StationName string `json:"station_name"`
	Time        string `json:"time"`
}
