package main

import (
	"context"
	"encoding/json"
	"fmt"
	"karangiri121/railway-scrapper/internal/config"
	"karangiri121/railway-scrapper/internal/db"
	"karangiri121/railway-scrapper/internal/models"
	"log"
	"net/http"
	"time"

	"gorm.io/gorm"
)

func main() {
	config := config.MustLoad()
	db, err := db.New(config)
	if err != nil {
		log.Fatal(err.Error())
	}
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	for _, ch := range "cp" {

		url := fmt.Sprintf("https://www.railyatri.in/trains/sort_by_char/%c", ch)

		resp, err := client.Get(url)
		if err != nil {
			fmt.Println("API request failed:", err)
			return
		}
		defer resp.Body.Close()

		var result map[string]interface{}

		err = json.NewDecoder(resp.Body).Decode(&result)
		if err != nil {
			panic(err)
		}

		data := result["trains"]
		var trains []models.Train

		for _, item := range data.([]interface{}) {
			trainMap := item.(map[string]interface{})
			train := models.TrainFromMap(trainMap)
			trains = append(trains, train)
		}
		if len(trains) == 0 {
			continue
		}
		for _, train := range trains {
			err := gorm.G[models.Train](db).Create(context.Background(), &train)
			if err != nil {
				fmt.Printf("Error On %c %s %s", ch, err.Error(), train.ID)
			}
		}
		// temp := db.Create(trains)

		// if temp.Error != nil {
		// 	fmt.Printf("Error On %c %s", ch, temp.Error.Error())
		// }
		time.Sleep(1 * time.Second)
	}

	var finalTrains []models.Train
	db.Find(&finalTrains)

	fmt.Println(len(finalTrains))
}
