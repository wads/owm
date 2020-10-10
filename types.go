package openweathermap

import (
	"encoding/json"
	"fmt"
	"strconv"
)

type degree float64

func (d degree) String() string {
	return strconv.FormatFloat(float64(d), 'f', -1, 64)
}

type Coord struct {
	Lat degree
	Lon degree
}

func ValidateCoord(coord *Coord) bool {
	if coord.Lat < -90 && coord.Lat > 90 {
		return false
	}

	if coord.Lon < -180 && coord.Lon > 180 {
		return false
	}

	return true
}

type BoundingBox struct {
	LatTop    degree
	LatBottom degree
	LonLeft   degree
	LonRight  degree
	Zoom      int
}

func (b *BoundingBox) String() string {
	return fmt.Sprintf("%f,%f,%f,%f,%d", b.LonLeft, b.LatBottom, b.LonRight, b.LatTop, b.Zoom)
}

func ValidateBoundingBox(bbox *BoundingBox) bool {
	if bbox.LatTop < -90 && bbox.LatTop > 90 {
		return false
	}

	if bbox.LatBottom < -90 && bbox.LatBottom > 90 {
		return false
	}

	if bbox.LonLeft < -180 && bbox.LonLeft > 180 {
		return false
	}

	if bbox.LonRight < -180 && bbox.LonRight > 180 {
		return false
	}

	return true
}

type Weather struct {
	ID          int
	Main        string
	Description string
	Icon        string
}

type Main struct {
	Temp      float32
	Pressure  int
	Humidity  int
	TempMin   float32 `json:"temp_min"`
	TempMax   float32 `json:"temp_max"`
	SeaLevel  float32 `json:"sea_level"`
	GrndLevel float32 `json:"grnd_level"`
}

type Wind struct {
	Speed float32
	Deg   int
}

type Clouds struct {
	All   int
	Today int
}

type Rain struct {
	OneH   float32 `json:"1h"`
	ThreeH float32 `json:"3h"`
}

type Snow struct {
	OneH   float32 `json:"1h"`
	ThreeH float32 `json:"3h"`
}

type Sys struct {
	Type    int
	ID      int
	Message float32
	Country string
	Sunrise int
	Sunset  int
}

type CurrentAndForecastWeather struct {
	Lat            float32    `json:"lat"`
	Lon            float32    `json:"lon"`
	Timezone       string     `json:"timezone"`
	TimezoneOffset int        `json:"timezone_offset"`
	Current        Current    `json:"current"`
	Minutely       []Minutely `json:"minutely"`
	Hourly         []Hourly   `json:"hourly"`
	Daily          []Daily    `json:"daily"`
}

type Current struct {
	Dt         int       `json:"dt"`
	Sunrise    int       `json:"sunrise"`
	Sunset     int       `json:"sunset"`
	Temp       float32   `json:"temp"`
	FeelsLike  float32   `json:"feels_like"`
	Pressure   int       `json:"pressure"`
	Humidity   int       `json:"humidity"`
	DewPoint   float32   `json:"dew_point"`
	Clouds     int       `json:"clouds"`
	Uvi        float32   `json:"uvi"`
	Visibility int       `json:"visibility"`
	WindSpeed  float32   `json:"wind_speed"`
	WindGust   float32   `json:"wind_gust"`
	WindDeg    int       `json:"wind_deg"`
	Rain       Rain      `json:"rain"`
	Snow       Snow      `json:"snow"`
	Weather    []Weather `json:"weather"`
}

type Minutely struct {
	Dt            int     `json:"dt"`
	Precipitation float32 `json:"precipitation"`
}

type Hourly struct {
	Dt         int       `json:"dt"`
	Temp       float32   `json:"temp"`
	FeelsLike  float32   `json:"feels_like"`
	Pressure   int       `json:"pressure"`
	Humidity   int       `json:"humidity"`
	DewPoint   float32   `json:"dew_point"`
	Clouds     int       `json:"clouds"`
	Visibility int       `json:"visibility"`
	WindSpeed  float32   `json:"wind_speed"`
	WindGust   float32   `json:"wind_gust"`
	WindDeg    int       `json:"wind_deg"`
	Rain       Rain      `json:"rain"`
	Snow       Snow      `json:"snow"`
	Weather    []Weather `json:"weather"`
}

type Daily struct {
	Dt         int       `json:"dt"`
	Temp       Temp      `json:"temp"`
	FeelsLike  FeelsLike `json:"feels_like"`
	Pressure   int       `json:"pressure"`
	Humidity   int       `json:"humidity"`
	DewPoint   float32   `json:"dew_point"`
	Clouds     int       `json:"clouds"`
	Visibility int       `json:"visibility"`
	WindSpeed  float32   `json:"wind_speed"`
	WindGust   float32   `json:"wind_gust"`
	WindDeg    int       `json:"wind_deg"`
	Rain       float32   `json:"rain"`
	Snow       Snow      `json:"snow"`
	Weather    []Weather `json:"weather"`
}

type TimeOfDay struct {
	Morn  float32 `json:"morn"`
	Day   float32 `json:"day"`
	Eve   float32 `json:"eve"`
	Night float32 `json:"night"`
}

type Temp struct {
	TimeOfDay
	Min float32 `json:"min"`
	Max float32 `json:"max"`
}

type FeelsLike struct {
	TimeOfDay
}

type CurrentWeather struct {
	Coord      Coord
	Weather    []Weather
	Base       string
	Main       Main
	Visibility int
	Wind       Wind
	Clouds     Clouds
	Rain       Rain
	Snow       Snow
	Dt         int
	Sys        Sys
	Timezone   int
	ID         int
	Name       string
	Cod        int
}

type CurrentCitiesWeather struct {
	Cod      int
	Calctime float32
	Cnt      int
	List     []CurrentWeather
}

func (c *CurrentCitiesWeather) UnmarshalJSON(data []byte) error {
	type Alias CurrentCitiesWeather
	a := &struct {
		Cod json.RawMessage
		*Alias
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &a); err != nil {
		return err
	}

	var cod string
	if err := json.Unmarshal(a.Cod, &cod); err == nil {
		i, err := strconv.Atoi(cod)
		if err != nil {
			return err
		}
		c.Cod = i
	}

	return nil
}
