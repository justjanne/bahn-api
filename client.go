package bahn

import (
	"fmt"
	"golang.org/x/net/html/charset"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"net/url"
)

type ApiClient struct {
	IrisBaseUrl          string
	CoachSequenceBaseUrl string
	HafasBaseUrl         string
	HttpClient           *http.Client
}

func (c *ApiClient) Station(evaId int64) ([]Station, error) {
	var err error

	uri := fmt.Sprintf("%s/timetable/station/%d", c.IrisBaseUrl, evaId)

	var stations []Station

	var response *http.Response
	if response, err = c.HttpClient.Get(uri); err != nil {
		return stations, err
	}

	if stations, err = StationsFromReader(response.Body); err != nil {
		return stations, err
	}

	if err = response.Body.Close(); err != nil {
		return stations, err
	}

	return stations, err
}

func (c *ApiClient) Timetable(evaId int64, date time.Time) (Timetable, error) {
	var err error

	BahnFormat := "060102/15"
	uri := fmt.Sprintf("%s/timetable/plan/%d/%s", c.IrisBaseUrl, evaId, date.Format(BahnFormat))

	var timetable Timetable

	var response *http.Response
	if response, err = c.HttpClient.Get(uri); err != nil {
		return timetable, err
	}

	if timetable, err = TimetableFromReader(response.Body); err != nil {
		return timetable, err
	}

	if err = response.Body.Close(); err != nil {
		return timetable, err
	}

	return timetable, err
}

func (c *ApiClient) RealtimeAll(evaId int64, date time.Time) (Timetable, error) {
	var err error

	uri := fmt.Sprintf("%s/timetable/fchg/%d", c.IrisBaseUrl, evaId)

	var timetable Timetable

	var response *http.Response
	if response, err = c.HttpClient.Get(uri); err != nil {
		return timetable, err
	}

	if timetable, err = TimetableFromReader(response.Body); err != nil {
		return timetable, err
	}

	if err = response.Body.Close(); err != nil {
		return timetable, err
	}

	return timetable, err
}

func (c *ApiClient) RealtimeRecent(evaId int64, date time.Time) (Timetable, error) {
	var err error

	uri := fmt.Sprintf("%s/timetable/rchg/%d", c.IrisBaseUrl, evaId)

	var timetable Timetable

	var response *http.Response
	if response, err = c.HttpClient.Get(uri); err != nil {
		return timetable, err
	}

	if timetable, err = TimetableFromReader(response.Body); err != nil {
		return timetable, err
	}

	if err = response.Body.Close(); err != nil {
		return timetable, err
	}

	return timetable, err
}

func (c *ApiClient) WingDefinition(parent string, wing string) (WingDefinition, error) {
	var err error

	uri := fmt.Sprintf("%s/timetable/wingdef/%s/%s", c.IrisBaseUrl, parent, wing)

	var wingDefinition WingDefinition

	var response *http.Response
	if response, err = c.HttpClient.Get(uri); err != nil {
		return wingDefinition, err
	}

	if wingDefinition, err = WingDefinitionFromReader(response.Body); err != nil {
		return wingDefinition, err
	}

	if err = response.Body.Close(); err != nil {
		return wingDefinition, err
	}

	return wingDefinition, err
}

func (c *ApiClient) CoachSequence(line string, date time.Time) (CoachSequence, error) {
	var err error

	uri := fmt.Sprintf("%s/%s/%s", c.CoachSequenceBaseUrl, line, date.Format(TimeLayoutShort))

	var coachSequence CoachSequence

	var response *http.Response
	if response, err = c.HttpClient.Get(uri); err != nil {
		return coachSequence, err
	}

	if coachSequence, err = CoachSequenceFromReader(response.Body); err != nil {
		return coachSequence, err
	}

	if err = response.Body.Close(); err != nil {
		return coachSequence, err
	}

	return coachSequence, err
}

func (c *ApiClient) Suggestions(line string, date time.Time) ([]Suggestion, error) {
	var err error

	uri := fmt.Sprintf("%s/trainsearch.exe/dn", c.HafasBaseUrl)

	var suggestions []Suggestion

	DateFormat := "02.01.2006"
	body := url.Values{
		"maxResults": []string{"50"},
		"trainname":  []string{line},
		"date":       []string{date.Format(DateFormat)},
		"L":          []string{"vs_json.vs_hap"},
	}

	var response *http.Response
	if response, err = c.HttpClient.PostForm(uri, body); err != nil {
		return suggestions, err
	}

	var utf8reader io.Reader
	if utf8reader, err = charset.NewReaderLabel("ISO 8859-1", response.Body); err != nil {
		return suggestions, nil
	}

	var content []byte
	if content, err = ioutil.ReadAll(utf8reader); err != nil {
		return suggestions, err
	}
	strippedContent := string(content)
	strippedContent = strings.TrimPrefix(strippedContent, "TSLs.sls = ")
	strippedContent = strings.TrimSuffix(strippedContent, ";")

	if suggestions, err = SuggestionsFromBytes([]byte(strippedContent)); err != nil {
		return suggestions, err
	}

	if err = response.Body.Close(); err != nil {
		return suggestions, err
	}

	return suggestions, err
}

func (c *ApiClient) HafasMessages(trainlink string) ([]HafasMessage, error) {
	var err error

	uri := fmt.Sprintf("%s/traininfo.exe/dn/%s?rt=1&ajax=1", c.HafasBaseUrl, trainlink)

	var messages []HafasMessage
	request, err := http.NewRequest("GET", uri, nil)

	var response *http.Response
	if response, err = c.HttpClient.Do(request); err != nil {
		return messages, err
	}

	if messages, err = HafasMessagesFromReader(response.Body); err != nil {
		return messages, err
	}

	if err = response.Body.Close(); err != nil {
		return messages, err
	}

	return messages, err
}
