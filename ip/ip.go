package ip

import (
	"fmt"
	"net"
)

type Location struct {
	Ip          string `json:"ip"`
	Country     string `json:"country"`
	CountryCode string `json:"loc2"`
	State       string `json:"state"`
	StateCode   string `json:"loc4"`
	City        string `json:"city"`
	CityCode    string `json:"loc6"`
	Code        string `json:"loc"`
}

type DbInfo struct {
	reader *reader
}

type Db struct {
	reader *reader
}

func NewDb(name string) (*Db, error) {
	r, e := newReader(name, &DbInfo{})
	if e != nil {
		return nil, e
	}

	return &Db{
		reader: r,
	}, nil
}

func (db *Db) Find(addr string) (*Location, error) {
	ipv := net.ParseIP(addr)
	if ipv == nil {
		return nil, fmt.Errorf("%s", "ip format error.")
	}

	loc, err := db.reader.find1(addr, "CN")
	if err != nil {
		return nil, err
	}

	code := GetRegionCode(loc[0], loc[1], loc[2])
	cCode := ""
	sCode := ""
	ccCode := ""
	if len(code) >= 2 {
		cCode = code[:2]
		sCode = code[:2]
		ccCode = code[:2]
	}
	if len(code) >= 4 {
		sCode = code[:4]
		ccCode = code[:4]
	}
	if len(code) >= 6 {
		ccCode = code[:6]
	}

	return &Location{
		Ip:          ipv.To4().String(),
		Country:     loc[0],
		State:       loc[1],
		City:        loc[2],
		CountryCode: cCode,
		StateCode:   sCode,
		CityCode:    ccCode,
		Code:        code,
	}, nil
}

type IrregularLocation struct {
	Country []string
	State   []string
	City    []string
}

func GetRegionCode(c, s, cc string) string {
	country := GetCountry(c)
	if country == nil {
		return "ZZ"
	}

	state := country.GetState(s)
	if state == nil {
		return country.Code
	}

	city := state.GetCity(cc)
	if city == nil {
		return state.Code
	}
	return city.Code
}
