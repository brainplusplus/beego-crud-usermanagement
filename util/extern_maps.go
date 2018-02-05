package util

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"bitbucket.org/tixid/TixId_Crawler_Go/logging"
)

var log = logging.MustGetLogger()
var maps = make(map[string](map[string]string))

const (
	M_ERROR_CODES_EN    = "error_codes_en"
	M_ERROR_CODES_ID    = "error_codes_id"
)

type Map struct {
	Name    string            `json:"map_name"`
	Entries map[string]string `json:"entries"`
}

type MapsDocument struct {
	Maps []Map `json:"maps"`
}

// Call this to cause the map data to be loaded from the specified file
// and placed into the private maps map
func InitializeMaps(fileName string) error {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Error("Error reading maps file ", err.Error())
		return err
	}
	log.Info("found map data file map_data.json")
	doc := MapsDocument{}
	if err := json.Unmarshal(bytes, &doc); err != nil {
		return err
	}
	for _, elem := range doc.Maps {
		log.Debug( "populating map ", elem.Name)
		maps[elem.Name] = elem.Entries
	}
	log.Info( "parsed map data file map_data.json")

	return nil
}

// Call this to get a specific map
// NOTE that the caller is then RESPONSIBLE for handling
// cases where the keys of interest are not present
// Preferably use MustGetString() or LookupFunc()
// Panics if the map of interest is not found
func MustGetMap(mapName string) map[string]string {
	m, ok := maps[mapName]
	if !ok {
		log.Error("missing map ", mapName)
		panic("missing map " + mapName)
	}
	return m
}


// Call this to get the value associated with a key in a given map
// Panics if either the map or key are not found
func MustGetString(mapName string, key string) string {
	m := MustGetMap(mapName)

	v, ok := m[key]
	if !ok {
		log.Error("missing key ", key," in map ", mapName)
		panic("missing key " + key + "in map " + mapName)
	}
	return v
}

type MappingFunc func(string) string

// Returns a function that can be used to retrieve keys from
// the specified map
// The returned function will panic if asked to retrieve from a map
// or a key that does not exist
func LookupFunc(mapName string) MappingFunc {

	return func(key string) string {
		m := MustGetMap(mapName)
		v, ok := m[key]
		if !ok {
			log.Error( "missing key ", key," in map ", mapName)
			panic("missing key " + key + " in map " + mapName)
		}
		return v
	}
}

// working the same as LookupFunc, but if the key is not found, it will override with another key
func LookupFuncWithDefaultKey(mapName string, defaultKey string) MappingFunc {

	return func(key string) string {
		m := MustGetMap(mapName)
		v, ok := m[key]
		if !ok {
			return LookupFunc(mapName)(defaultKey)
		}
		return v
	}
}

func LookupFuncIgnoreCase(mapName string) MappingFunc {

	return func(key string) string {
		m := MustGetMap(mapName)
		v, ok := m[strings.TrimSpace(strings.ToLower(key))]
		if !ok {
			log.Error( "missing key ", key," in map ", mapName)
			panic("missing key " + key + " in map " + mapName)
		}
		return v
	}
}

func LookupFuncIgnoreCaseNoPanic(mapName string) MappingFunc {

	return func(key string) string {
		m := MustGetMap(mapName)
		v, ok := m[strings.TrimSpace(strings.ToLower(key))]
		if !ok {
			log.Error( "missing key ", key," in map ", mapName)
			return ""
		}
		return v
	}
}