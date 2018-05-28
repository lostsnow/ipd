package ip

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"sort"
	"strings"
)

type Location struct {
	Ip          string
	Country     string
	CountryCode string
	State       string
	StateCode   string
	City        string
	CityCode    string
	Code        string
}

type Db struct {
	file   *os.File
	index  []byte
	data   []byte
	offset int
}

func NewDb(name string) (*Db, error) {
	db := &Db{}

	if err := db.load(name); err != nil {
		return nil, err
	}

	return db, nil
}

func (db *Db) load(fn string) error {
	var err error
	db.file, err = os.Open(fn)
	if err != nil {
		return err
	}
	b4 := make([]byte, 4)
	_, err = db.file.Read(b4)
	if err != nil {
		return err
	}
	off := int(binary.BigEndian.Uint32(b4))
	db.offset = off

	_, err = db.file.Seek(262148, 0)
	if err != nil {
		return err
	}
	l := off - 262148 - 262144
	db.index = make([]byte, l)
	_, err = db.file.Read(db.index)
	if err != nil {
		return err
	}

	db.data, err = ioutil.ReadAll(db.file)
	if err != nil {
		return err
	}
	return nil
}

func (db *Db) Find(s string) (*Location, error) {
	ipv := net.ParseIP(s)
	if ipv == nil {
		return nil, fmt.Errorf("%s", "ip format error.")
	}

	low := 0
	mid := 0
	high := int(len(db.index)/9) - 1

	val := binary.BigEndian.Uint32(ipv.To4())

	for low <= high {
		mid = int((low + high) / 2)
		pos := mid * 9

		var start uint32
		if mid > 0 {
			pos1 := (mid - 1) * 9
			start = binary.BigEndian.Uint32(db.index[pos1:pos1+4]) + 1
		} else {
			start = 0
		}

		end := binary.BigEndian.Uint32(db.index[pos : pos+4])

		if val < start {
			high = mid - 1
		} else if val > end {
			low = mid + 1
		} else {
			off := int(binary.LittleEndian.Uint32([]byte{
				db.index[pos+4],
				db.index[pos+5],
				db.index[pos+6],
				0,
			}))
			l := int(db.index[pos+7])*256 + int(db.index[pos+8])

			loc := strings.Split(string(db.data[off:off+l]), "\t")
			if len(loc) < 4 {
				return nil, fmt.Errorf("%s", "ip location error.")
			}
			code := GetRegionCode(loc[0], loc[1], loc[2])

			cCode := ""
			sCode := ""
			ccCode := ""
			if len(code) >= 2 {
				cCode = code[:2]
				sCode = code[:2]
			}
			if len(code) >= 4 {
				sCode = code[:4]
			}
			if len(code) >= 6 {
				ccCode = code[:6]
			}
			return &Location{
				Ip:          intToIP(val).String(),
				Country:     loc[0],
				State:       loc[1],
				City:        loc[2],
				CountryCode: cCode,
				StateCode:   sCode,
				CityCode:    ccCode,
				Code:        code,
			}, nil
		}
	}

	return nil, fmt.Errorf("%s", "not found")
}

type IrregularLocation struct {
	Country []string
	State   []string
	City    []string
}

func (db *Db) Check() (*IrregularLocation, error) {
	i := &IrregularLocation{}

	iCountry := map[string]bool{}
	iState := map[string]bool{}
	iCity := map[string]bool{}

	for pos := 0; pos < db.offset-262148-262144; pos += 9 {

		// end := binary.BigEndian.Uint32(db.index[pos : pos+4])
		off := binary.LittleEndian.Uint32(append(db.index[pos+4:pos+7], byte(0)))
		length := int(db.index[pos+7])*256 + int(db.index[pos+8])

		_, err := db.file.Seek(int64(off)+int64(length), 0)
		if err != nil {
			return nil, err
		}

		loc := strings.Split(string(db.data[off:off+uint32(length)]), "\t")

		country := GetCountry(loc[0])
		if _, ok := irregularCountries[loc[0]]; country == nil && !ok {
			if _, ok = iCountry[loc[0]]; !ok {
				i.Country = append(i.Country, loc[0])
				iCountry[loc[0]] = true
			}
		}
		if country == nil {
			continue
		}

		if country.Code != "CN" || loc[1] == "中国" {
			continue
		}

		state := country.GetState(loc[1])
		if _, ok := iState[loc[0]+loc[1]]; state == nil && !ok {
			i.State = append(i.State, loc[0]+loc[1])
			iState[loc[0]+loc[1]] = true
		}
		if state == nil || loc[2] == "" {
			continue
		}

		noCity := map[string]bool{
			"北京": true,
			"天津": true,
			"上海": true,
			"重庆": true,
			"台湾": true,
		}
		if _, ok := noCity[loc[1]]; ok {
			continue
		}

		city := state.GetCity(loc[2])
		if _, ok := iCity[loc[0]+loc[1]+loc[2]]; city == nil && !ok {
			i.City = append(i.City, loc[0]+loc[1]+loc[2])
			iCity[loc[0]+loc[1]+loc[2]] = true
		}
	}

	sort.Strings(i.Country)
	sort.Strings(i.State)
	sort.Strings(i.City)

	return i, nil
}

func intToIP(nn uint32) net.IP {
	ip := make(net.IP, 4)
	binary.BigEndian.PutUint32(ip, nn)
	return ip
}

func GetRegionCode(c, s, cc string) string {
	country := GetCountry(c)
	if country == nil {
		return ""
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
