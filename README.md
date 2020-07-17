<!-- markdownlint-disable MD041-->
[![Go Report Card](https://goreportcard.com/badge/github.com/EricNeid/go-openweather?style=flat-square)](https://goreportcard.com/report/github.com/EricNeid/go-openweather)
![Go](https://github.com/EricNeid/go-openweather/workflows/Go/badge.svg)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/EricNeid/go-openweather)
[![Release](https://img.shields.io/github/release/EricNeid/go-openweather.svg?style=flat-square)](https://github.com/EricNeid/go-openweather/releases/latest)
[![Gitpod Ready-to-Code](https://img.shields.io/badge/Gitpod-Ready--to--Code-blue?logo=gitpod)](https://gitpod.io/#https://github.com/EricNeid/go-openweather)

# About

This Repo contains golang library to query OpenWetherMaps (<http://openweathermap.org/>) for weather information.

* current weather: <http://openweathermap.org/current>
* 5 days forecast: <http://openweathermap.org/forecast5>

## Install

```bash
go get github.com/EricNeid/openweather
```

## Documentation

Is available on ``godoc``:

<https://godoc.org/github.com/EricNeid/go-openweather>

## Examples

Consuming the library:

```go
import "github.com/EricNeid/openweather"

// create a query
q := openweather.NewQueryForCity(readAPIKey(), "Berlin,de")

// obtain data
resp, err := q.Weather()

// enjoy
fmt.Println(resp.Name) // Berlin
fmt.Println(resp.Weather[0].Description) // misc
fmt.Println(resp.Main.Temp) // 1
```

See the test files for more example.

A simple client for testing is also included:

```bash
go build cmd/openweatherclient
openweatherclient -key <OpenWeather API Key> -city Berlin,de
```
