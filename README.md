# WeatherApp
 
A simple app which use API OpenWeatherMap to get the current temperature in a given city .

1. Get API key in OpenWeatherMap: [link](https://openweathermap.org/appid)

2. Change the value `apiKey` in file `main.go` on your own API key.

3. Make sure you have Go installed on your computer.

4. Launch the app using this command: go run main.go

5. Enter the name of the city you want to get the current temperature for.

6. The app will display the current temperature in celsius for the specified city.

## Dependencies

- Go 1.20+
- "net/http" - package for making HTTP requests.
- "encoding/json" - package for decoding JSON data.
