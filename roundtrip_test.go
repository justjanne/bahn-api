package bahn

import (
	"bufio"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"testing"
)

func decodeXml(target interface{}, filename string) error {
	var err error
	var xmlFile *os.File

	if xmlFile, err = os.Open(filename); err != nil {
		return err
	}

	if err = xml.NewDecoder(xmlFile).Decode(&target); err != nil {
		return err
	}

	if err = xmlFile.Close(); err != nil {
		return err
	}

	return nil
}

func decodeJson(target interface{}, filename string) error {
	var err error
	var jsonFile *os.File

	if jsonFile, err = os.Open(filename); err != nil {
		return err
	}

	if err = json.NewDecoder(jsonFile).Decode(&target); err != nil {
		return err
	}

	if err = jsonFile.Close(); err != nil {
		return err
	}

	return nil
}

func encodeYaml(target interface{}, filename string) error {
	var err error
	var yamlFile *os.File

	if yamlFile, err = os.Create(filename); err != nil {
		return err
	}

	var data []byte
	if data, err = yaml.Marshal(&target); err != nil {
		return err
	}

	writer := bufio.NewWriter(yamlFile)
	if _, err = writer.Write(data); err != nil {
		return err
	}

	if err = writer.Flush(); err != nil {
		return err
	}

	if err = yamlFile.Close(); err != nil {
		return err
	}

	return nil
}

func encodeJson(target interface{}, filename string) error {
	var err error
	var jsonFile *os.File

	if jsonFile, err = os.Create(filename); err != nil {
		return err
	}

	if err = json.NewEncoder(jsonFile).Encode(&target); err != nil {
		return err
	}

	if err = jsonFile.Close(); err != nil {
		return err
	}

	return nil
}

func encodeXml(target interface{}, filename string) error {
	var err error
	var xmlFile *os.File

	if xmlFile, err = os.Create(filename); err != nil {
		return err
	}

	if err = xml.NewEncoder(xmlFile).Encode(&target); err != nil {
		return err
	}

	if err = xmlFile.Close(); err != nil {
		return err
	}

	return nil
}

func xmlInput(raw interface{}, filename string) {
	var err error

	if err = decodeXml(raw, filename+".xml"); err != nil {
		panic(err.Error())
	}
}

func jsonInput(raw interface{}, filename string) {
	var err error

	if err = decodeJson(raw, filename+".json"); err != nil {
		panic(err.Error())
	}
}

func xmlOutput(raw interface{}, data interface{}, filename string) {
	var err error

	if err = encodeXml(raw, filename+".roundtrip.xml"); err != nil {
		panic(err.Error())
	}

	if err = encodeYaml(data, filename+".yaml"); err != nil {
		panic(err.Error())
	}

	if err = encodeJson(data, filename+".json"); err != nil {
		panic(err.Error())
	}
}

func jsonOutput(raw interface{}, data interface{}, filename string) {
	var err error

	if err = encodeJson(raw, filename+".roundtrip.json"); err != nil {
		panic(err.Error())
	}

	if err = encodeYaml(data, filename+".yaml"); err != nil {
		panic(err.Error())
	}

	if err = encodeJson(data, filename+".json"); err != nil {
		panic(err.Error())
	}
}

const InputFolder = "data"
const OutputFolder = "out"

var stationData []byte
var timetableData []byte
var realtimeData []byte
var wingDefinitionData []byte
var coachSequenceData []byte

func TestMain(m *testing.M) {
	var err error

	if stationData, err = ioutil.ReadFile(fmt.Sprintf("%s/%s/%d.xml", InputFolder, "iris_station", 0)); err != nil {
		panic(err)
	}
	if timetableData, err = ioutil.ReadFile(fmt.Sprintf("%s/%s/%d.xml", InputFolder, "iris_timetable", 0)); err != nil {
		panic(err)
	}
	if realtimeData, err = ioutil.ReadFile(fmt.Sprintf("%s/%s/%d.xml", InputFolder, "iris_realtime", 0)); err != nil {
		panic(err)
	}
	if wingDefinitionData, err = ioutil.ReadFile(fmt.Sprintf("%s/%s/%d.xml", InputFolder, "iris_wingdef", 0)); err != nil {
		panic(err)
	}
	if coachSequenceData, err = ioutil.ReadFile(fmt.Sprintf("%s/%s/%d.json", InputFolder, "apps_wagenreihung", 0)); err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}

func BenchmarkStation(b *testing.B) {
	if _, err := StationsFromBytes(stationData); err != nil {
		b.Error(err.Error())
	}
}

func BenchmarkTimetable(b *testing.B) {
	if _, err := TimetableFromBytes(timetableData); err != nil {
		b.Error(err.Error())
	}
}

func BenchmarkRealtime(b *testing.B) {
	if _, err := TimetableFromBytes(realtimeData); err != nil {
		b.Error(err.Error())
	}
}

func BenchmarkWingDefinition(b *testing.B) {
	if _, err := WingDefinitionFromBytes(wingDefinitionData); err != nil {
		b.Error(err.Error())
	}
}

func BenchmarkCoachSequence(b *testing.B) {
	if _, err := CoachSequenceFromBytes(coachSequenceData); err != nil {
		b.Error(err.Error())
	}
}

func BenchmarkRoundtrip(b *testing.B) {
	for i := 0; i < 3; i++ {
		var raw rawStations
		folderName := "iris_station"
		input := fmt.Sprintf("%s/%s/%d", InputFolder, folderName, i)
		output := fmt.Sprintf("%s/%s/%d", OutputFolder, folderName, i)
		xmlInput(&raw, input)
		data := parseStations(raw)
		xmlOutput(&raw, &data, output)
	}

	for i := 0; i < 5; i++ {
		var raw rawTimetable
		folderName := "iris_timetable"
		input := fmt.Sprintf("%s/%s/%d", InputFolder, folderName, i)
		output := fmt.Sprintf("%s/%s/%d", OutputFolder, folderName, i)
		xmlInput(&raw, input)
		data := parseTimetable(raw)
		xmlOutput(&raw, &data, output)
	}

	for i := 0; i < 5; i++ {
		var raw rawTimetable
		folderName := "iris_realtime"
		input := fmt.Sprintf("%s/%s/%d", InputFolder, folderName, i)
		output := fmt.Sprintf("%s/%s/%d", OutputFolder, folderName, i)
		xmlInput(&raw, input)
		data := parseTimetable(raw)
		xmlOutput(&raw, &data, output)
	}

	for i := 0; i < 3; i++ {
		var raw rawWingDefinition
		folderName := "iris_wingdef"
		input := fmt.Sprintf("%s/%s/%d", InputFolder, folderName, i)
		output := fmt.Sprintf("%s/%s/%d", OutputFolder, folderName, i)
		xmlInput(&raw, input)
		data := parseWingDefinition(raw)
		xmlOutput(&raw, &data, output)
	}

	for i := 0; i < 227; i++ {
		var raw rawCoachSequence
		folderName := "apps_wagenreihung"
		input := fmt.Sprintf("%s/%s/%d", InputFolder, folderName, i)
		output := fmt.Sprintf("%s/%s/%d", OutputFolder, folderName, i)
		jsonInput(&raw, input)
		data := parseCoachSequence(raw)
		jsonOutput(&raw, &data, output)
	}
}
