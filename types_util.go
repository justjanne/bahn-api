package bahn

import (
	"encoding/json"
	"encoding/xml"
	"strings"
	"time"
)

type bahnStringList []string

func (s *bahnStringList) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if s == nil {
		return xml.Attr{}, nil
	} else {
		return xml.Attr{
			Name:  name,
			Value: strings.Join(*s, "|"),
		}, nil
	}
}

func (s *bahnStringList) UnmarshalXMLAttr(attr xml.Attr) error {
	*s = strings.Split(attr.Value, "|")
	return nil
}

func (s *bahnStringList) Value() []string {
	if s != nil {
		return *s
	} else {
		return make([]string, 0)
	}
}

type timeLong struct {
	time.Time
}

const TimeLayoutLong = "06-01-02 15:04:05.999"

func (t *timeLong) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if t == nil || t.IsZero() {
		return xml.Attr{}, nil
	} else {
		return xml.Attr{
			Name:  name,
			Value: t.Format(TimeLayoutLong),
		}, nil
	}
}

func (t *timeLong) UnmarshalXMLAttr(attr xml.Attr) error {
	if attr.Value != "" {
		value, err := time.Parse(TimeLayoutLong, attr.Value)
		if err != nil {
			return err
		}
		t.Time = value
	}
	return nil
}

func (t *timeLong) Value() *time.Time {
	if t != nil {
		return &t.Time
	} else {
		return nil
	}
}

type timeShort struct {
	time.Time
}

const TimeLayoutShort = "0601021504"

func (t *timeShort) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if t == nil || t.IsZero() {
		return xml.Attr{}, nil
	} else {
		return xml.Attr{
			Name:  name,
			Value: t.Format(TimeLayoutShort),
		}, nil
	}
}

func (t *timeShort) UnmarshalXMLAttr(attr xml.Attr) error {
	if attr.Value != "" {
		value, err := time.Parse(TimeLayoutShort, attr.Value)
		if err != nil {
			return err
		}
		t.Time = value
	}
	return nil
}

func (t *timeShort) Value() *time.Time {
	if t != nil {
		return &t.Time
	} else {
		return nil
	}
}

type timeMediumShort struct {
	time.Time
}

const TimeLayoutMediumShort = "200601021504"

func (t *timeMediumShort) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if t == nil || t.IsZero() {
		return xml.Attr{}, nil
	} else {
		return xml.Attr{
			Name:  name,
			Value: t.Format(TimeLayoutMediumShort),
		}, nil
	}
}

func (t *timeMediumShort) UnmarshalXMLAttr(attr xml.Attr) error {
	if attr.Value != "" {
		value, err := time.Parse(TimeLayoutMediumShort, attr.Value)
		if err != nil {
			return err
		}
		t.Time = value
	}
	return nil
}

func (t *timeMediumShort) Value() *time.Time {
	if t != nil {
		return &t.Time
	} else {
		return nil
	}
}

type timeMedium struct {
	time.Time
}

const TimeLayoutMedium = "2006-01-02T15:04:05"

func (t *timeMedium) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if t == nil || t.IsZero() {
		return xml.Attr{}, nil
	} else {
		return xml.Attr{
			Name:  name,
			Value: t.Format(TimeLayoutMedium),
		}, nil
	}
}

func (t *timeMedium) UnmarshalXMLAttr(attr xml.Attr) error {
	if attr.Value != "" {
		value, err := time.Parse(TimeLayoutMedium, attr.Value)
		if err != nil {
			return err
		}
		t.Time = value
	}
	return nil
}

func (t *timeMedium) MarshalJSON() ([]byte, error) {
	var text string
	if t == nil || t.IsZero() {
		text = ""
	} else {
		text = t.Format(TimeLayoutMedium)
	}
	return json.Marshal(&text)
}

func (t *timeMedium) UnmarshalJSON(data []byte) error {
	var err error

	var text string
	if err := json.Unmarshal(data, &text); err != nil {
		return err
	}

	if text != "" {
		var value time.Time
		if value, err = time.Parse(TimeLayoutMedium, text); err != nil {
			return err
		}
		t.Time = value
	}
	return nil
}

func (t *timeMedium) Value() *time.Time {
	if t != nil {
		return &t.Time
	} else {
		return nil
	}
}

type bahnDate struct {
	time.Time
}

const DateLayoutLong = "2006-01-02"

func (t *bahnDate) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if t == nil || t.IsZero() {
		return xml.Attr{}, nil
	} else {
		return xml.Attr{
			Name:  name,
			Value: t.Format(DateLayoutLong),
		}, nil
	}
}

func (t *bahnDate) UnmarshalXMLAttr(attr xml.Attr) error {
	if attr.Value != "" {
		value, err := time.Parse(DateLayoutLong, attr.Value)
		if err != nil {
			return err
		}
		t.Time = value
	}
	return nil
}

func (t *bahnDate) MarshalJSON() ([]byte, error) {
	var text string
	if t == nil || t.IsZero() {
		text = ""
	} else {
		text = t.Format(DateLayoutLong)
	}
	return json.Marshal(&text)
}

func (t *bahnDate) UnmarshalJSON(data []byte) error {
	var err error

	var text string
	if err := json.Unmarshal(data, &text); err != nil {
		return err
	}

	if text != "" {
		var value time.Time
		if value, err = time.Parse(DateLayoutLong, text); err != nil {
			return err
		}
		t.Time = value
	}
	return nil
}

func (t *bahnDate) Value() *time.Time {
	if t != nil {
		return &t.Time
	} else {
		return nil
	}
}
