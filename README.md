# 🌦️ Weather Condition (Go Project)

A simple command-line weather application written in **Go**.  
It fetches real-time weather data (temperature, condition, humidity, wind speed) from the **OpenWeather API**.

---

## Features
- Fetches current weather for ANY city
- Displays:
  - 🌡️ Temperature (°C) and (Feels Like)
  - ☁️ Condition (Clear, Clouds, Rain, etc.)
  - 💧 Humidity
  - 🌬️ Wind Speed (km/h)
- Built with Go’s standard libraries (`net/http` and `encoding/json`)
- Easy to run and extend

---

## 📦 Installation & Usage

```bash
# Clone the repo
git clone https://github.com/username/weather-condition.git
cd weather-condition

# Initialize the module (first time only)
go mod init weather-condition

# Edit main.go and add your API key
apiKey := "YOUR_API_KEY"

# Run the program
go run main.go

# Example
Weather in Montreal
Temperature: 23.5°C (Feels like 22.3°C)
Condition  : Clear
Humidity   : 56%
Wind Speed : 12.3 km/h
