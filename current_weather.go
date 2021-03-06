package owm

import (
	"errors"
	"net/url"
)

type CurrentWeatherAPI struct {
	*OwmAPI
}

func NewCurrentWeatherAPI(config *Config) (*CurrentWeatherAPI, error) {
	if !ValidateConfig(config) {
		return nil, errors.New("Invalid Config")
	}

	return &CurrentWeatherAPI{NewOwmAPI(config, currentURL)}, nil
}

type cityNameParams struct {
	country string
	name    string
	state   string
}

func (c *cityNameParams) setQuery(values url.Values) {
	if c.name != "" {
		name := c.name

		if c.state != "" {
			name += "," + c.state
		}

		if c.country != "" {
			name += "," + c.country
		}

		values.Set("q", name)
	}
}

type CityNameOption func(*cityNameParams)

func WithStateOption(state string) CityNameOption {
	return func(c *cityNameParams) {
		c.state = state
	}
}

func WithCountryOption(country string) CityNameOption {
	return func(c *cityNameParams) {
		c.country = country
	}
}

func (c *CurrentWeatherAPI) GetByCityName(name string, opts ...CityNameOption) (*CurrentWeather, error) {
	params := &cityNameParams{name: name}

	for _, opt := range opts {
		opt(params)
	}

	c.Params = params

	weather := &CurrentWeather{}
	err := c.getAndLoad(weather)

	return weather, err
}

type cityIDParams struct {
	id string
}

func (c *cityIDParams) setQuery(values url.Values) {
	if c.id != "" {
		values.Set("id", c.id)
	}
}

func (c *CurrentWeatherAPI) GetByCityID(id string) (*CurrentWeather, error) {
	c.Params = &cityIDParams{id: id}

	weather := &CurrentWeather{}
	err := c.getAndLoad(weather)

	return weather, err
}

type coordParams struct {
	coord *Coord
}

func (c *coordParams) setQuery(values url.Values) {
	if c.coord != nil {
		values.Set("lat", c.coord.Lat.String())
		values.Set("lon", c.coord.Lon.String())
	}
}

func (c *CurrentWeatherAPI) GetByCoord(coord *Coord) (*CurrentWeather, error) {
	if !ValidateCoord(coord) {
		return nil, errors.New("Invalid Coord")
	}

	c.Params = &coordParams{coord: coord}

	weather := &CurrentWeather{}
	err := c.getAndLoad(weather)

	return weather, err
}

type zipCodeParams struct {
	zipCode string
}

func (z *zipCodeParams) setQuery(values url.Values) {
	if z.zipCode != "" {
		values.Set("zip", z.zipCode)
	}
}

func (c *CurrentWeatherAPI) GetByZIPCode(zipCode string) (*CurrentWeather, error) {
	c.Params = &zipCodeParams{zipCode: zipCode}

	weather := &CurrentWeather{}
	err := c.getAndLoad(weather)

	return weather, err
}
