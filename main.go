package main

//import
import (
	"encoding/json"
	"fmt"
	"net/http"
)

type WeatherData struct {
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		Humidity  int     `json:"humidity"`
	} `json:"main"`
	Weather []struct {
		Main string `json:"main"`
	} `json:"weather"`
	Wind struct {
		Speed float64 `json:"speed"` // m/s with units=metric
	} `json:"wind"`
	Name    string `json:"name"`
	Cod     int    `json:"cod"`
	Message string `json:"message"`
}

//API Key to be taken from Openweather
func main() {
apiKey := os.Getenv("OPENWEATHER_API_KEY")
if apiKey == "" {
    fmt.Println("Set OPENWEATHER_API_KEY")
    return
}
  //tested for Montreal
	city := "Montreal"

	resp, err := http.Get(fmt.Sprintf(
		"https://api.openweathermap.org/data/2.5/weather?q=%s&units=metric&appid=%s",
		city, apiKey,
	))

  //handle error(logic)
  if err != nil {
		fmt.Println("Error fetching weather:", err)
		return
	}
	defer resp.Body.Close()

	var data WeatherData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	// Handle API-level errors (if bad key, then city not found)
	if resp.StatusCode != http.StatusOK || (data.Cod != 0 && data.Cod != 200) {
		if data.Message != "" {
			fmt.Printf("API error: %s (code %d)\n", data.Message, data.Cod)
		} else {
			fmt.Printf("API error: HTTP %d\n", resp.StatusCode)
		}
		return
	}

	if len(data.Weather) == 0 {
		fmt.Println("No weather data returned (key may still be activating). Try again in a few minutes.")
		return
	}

  //print
	kmh := data.Wind.Speed * 3.6
	fmt.Printf("Weather in %s\n", data.Name)
	fmt.Printf("Temperature: %.1f°C (Feels like %.1f°C)\n", data.Main.Temp, data.Main.FeelsLike)
	fmt.Printf("Condition  : %s\n", data.Weather[0].Main)
	fmt.Printf("Humidity   : %d%%\n", data.Main.Humidity)
	fmt.Printf("Wind Speed : %.1f km/h\n", kmh)
}
