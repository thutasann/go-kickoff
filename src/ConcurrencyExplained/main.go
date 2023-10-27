package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

const apiKey = "a306ec26b8cf6b336377e33368e74afa";

func fetchWeather(city string, ch chan<-string, wg *sync.WaitGroup) interface{}{
	var data struct{
		Main struct {
			Temp float64 `json:"temp"`
		}
	}

	defer wg.Done();

	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)
	
	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("Error fetching weather for %s: %s\n", city, err)
	}
	defer resp.Body.Close();

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Printf("Error decoding weather data for %s : %s\n", city, err)
	}

	ch <- fmt.Sprintf("This is the %s", city)

	return data
}

func main(){
	cities := []string{"Toronto", "London", "Paris", "Tokyo", "Yangon"}
	startNow := time.Now();

	ch := make(chan string)
	var wg sync.WaitGroup

	for _, city := range cities {
		wg.Add(1)
		go fetchWeather(city, ch, &wg)
	}

	go func(){
		wg.Wait()
		close(ch)
	}()

	for result := range ch {
		fmt.Println("result =>", result)
	}
	
	fmt.Println("The operation takes : ", time.Since(startNow))
}