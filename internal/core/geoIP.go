package core

import (
	"log"
	"net"

	"github.com/oschwald/geoip2-golang"
)

type IPLocater interface {
	IPLocation(ip net.IP) GeoInfo
}

type GeoIP2Locater struct {
	dbFile string
}

func (l GeoIP2Locater) IPLocation(ip net.IP) GeoInfo {
	db, err := geoip2.Open(l.dbFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	record, err := db.City(ip)
	if err != nil {
		log.Fatal(err)
	}
	return GeoInfo{
		Coordinates: [2]float64{record.Location.Longitude, record.Location.Latitude},
		Country:     record.Country.Names["en"],
		City:        record.City.Names["en"],
	}
}

type GeoInfo struct {
	Coordinates [2]float64
	Country     string
	City        string
}
