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
	Caches               []CacheBackend
}

const cacheTimestamp = "2006-01-02T15:04"

func (c *ApiClient) loadWithCache(key string, result interface{}, load func(interface{}) error) error {
	for _, cache := range c.Caches {
		if err := cache.Get(key, &result); err == nil {
			for _, targetCache := range c.Caches {
				if targetCache == cache {
					break
				}
				_ = targetCache.Set(key, result)
			}
			return nil
		}
	}
	var err error
	if err = load(&result); err == nil {
		for _, cache := range c.Caches {
			_ = cache.Set(key, result)
		}
	}
	return err
}

func (c *ApiClient) Station(evaId int64) ([]Station, error) {
	key := fmt.Sprintf("realtime_recent %d", evaId)
	var result []Station
	err := c.loadWithCache(key, &result, func(target interface{}) error {
		var err error
		target, err = c.loadStation(evaId)
		return err
	})
	return result, err
}

func (c *ApiClient) loadStation(evaId int64) ([]Station, error) {
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
	key := fmt.Sprintf("timetable %d %s", evaId, date.Format(cacheTimestamp))

	var result Timetable
	err := c.loadWithCache(key, &result, func(target interface{}) error {
		var err error
		target, err = c.loadTimetable(evaId, date)
		return err
	})
	return result, err
}

func (c *ApiClient) loadTimetable(evaId int64, date time.Time) (Timetable, error) {
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
	key := fmt.Sprintf("realtime_all %d %s", evaId, date.Format(cacheTimestamp))

	var result Timetable
	err := c.loadWithCache(key, &result, func(target interface{}) error {
		var err error
		target, err = c.loadRealtimeAll(evaId, date)
		return err
	})
	return result, err
}

func (c *ApiClient) loadRealtimeAll(evaId int64, date time.Time) (Timetable, error) {
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
	key := fmt.Sprintf("realtime_recent %d %s", evaId, date.Format(cacheTimestamp))

	var result Timetable
	err := c.loadWithCache(key, &result, func(target interface{}) error {
		var err error
		target, err = c.loadRealtimeRecent(evaId, date)
		return err
	})
	return result, err
}

func (c *ApiClient) loadRealtimeRecent(evaId int64, date time.Time) (Timetable, error) {
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
	key := fmt.Sprintf("wing_definition %s %s", parent, wing)

	var result WingDefinition
	err := c.loadWithCache(key, &result, func(target interface{}) error {
		var err error
		target, err = c.loadWingDefinition(parent, wing)
		return err
	})
	return result, err
}

func (c *ApiClient) loadWingDefinition(parent string, wing string) (WingDefinition, error) {
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
	key := fmt.Sprintf("coach_sequence %s %s", line, date.Format(cacheTimestamp))

	var result CoachSequence
	err := c.loadWithCache(key, &result, func(target interface{}) error {
		var err error
		target, err = c.loadCoachSequence(line, date)
		return err
	})
	return result, err
}

func (c *ApiClient) loadCoachSequence(line string, date time.Time) (CoachSequence, error) {
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
	key := fmt.Sprintf("suggestions %s %s", line, date.Format(cacheTimestamp))

	var result []Suggestion
	err := c.loadWithCache(key, &result, func(target interface{}) error {
		var err error
		target, err = c.loadSuggestions(line, date)
		return err
	})
	return result, err
}

func (c *ApiClient) loadSuggestions(line string, date time.Time) ([]Suggestion, error) {
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
	key := fmt.Sprintf("hafas_messages %s", trainlink)

	var result []HafasMessage
	err := c.loadWithCache(key, &result, func(target interface{}) error {
		var err error
		target, err = c.loadHafasMessages(trainlink)
		return err
	})
	return result, err
}

func (c *ApiClient) loadHafasMessages(trainlink string) ([]HafasMessage, error) {
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
