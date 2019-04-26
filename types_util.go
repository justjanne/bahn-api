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

type bahnTime struct {
	time.Time
}

const bahnTimeLayout = "06-01-02 15:04:05.999"

func (t *bahnTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if t == nil || t.IsZero() {
		return xml.Attr{}, nil
	} else {
		return xml.Attr{
			Name:  name,
			Value: t.Format(bahnTimeLayout),
		}, nil
	}
}

func (t *bahnTime) UnmarshalXMLAttr(attr xml.Attr) error {
	if attr.Value != "" {
		value, err := time.Parse(bahnTimeLayout, attr.Value)
		if err != nil {
			return err
		}
		t.Time = value
	}
	return nil
}

func (t *bahnTime) Value() *time.Time {
	if t != nil {
		return &t.Time
	} else {
		return nil
	}
}

type shortBahnTime struct {
	time.Time
}

const shortBahnTimeLayout = "0601021504"

func (t *shortBahnTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if t == nil || t.IsZero() {
		return xml.Attr{}, nil
	} else {
		return xml.Attr{
			Name:  name,
			Value: t.Format(shortBahnTimeLayout),
		}, nil
	}
}

func (t *shortBahnTime) UnmarshalXMLAttr(attr xml.Attr) error {
	if attr.Value != "" {
		value, err := time.Parse(shortBahnTimeLayout, attr.Value)
		if err != nil {
			return err
		}
		t.Time = value
	}
	return nil
}

func (t *shortBahnTime) Value() *time.Time {
	if t != nil {
		return &t.Time
	} else {
		return nil
	}
}

type mediumBahnTime struct {
	time.Time
}

const mediumBahnTimeLayout = "2006-01-02T15:04:05"

func (t *mediumBahnTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if t == nil || t.IsZero() {
		return xml.Attr{}, nil
	} else {
		return xml.Attr{
			Name:  name,
			Value: t.Format(mediumBahnTimeLayout),
		}, nil
	}
}

func (t *mediumBahnTime) UnmarshalXMLAttr(attr xml.Attr) error {
	if attr.Value != "" {
		value, err := time.Parse(mediumBahnTimeLayout, attr.Value)
		if err != nil {
			return err
		}
		t.Time = value
	}
	return nil
}

func (t *mediumBahnTime) MarshalJSON() ([]byte, error) {
	var text string
	if t == nil || t.IsZero() {
		text = ""
	} else {
		text = t.Format(mediumBahnTimeLayout)
	}
	return json.Marshal(&text)
}

func (t *mediumBahnTime) UnmarshalJSON(data []byte) error {
	var err error

	var text string
	if err := json.Unmarshal(data, &text); err != nil {
		return err
	}

	if text != "" {
		var value time.Time
		if value, err = time.Parse(mediumBahnTimeLayout, text); err != nil {
			return err
		}
		t.Time = value
	}
	return nil
}

func (t *mediumBahnTime) Value() *time.Time {
	if t != nil {
		return &t.Time
	} else {
		return nil
	}
}

type bahnDate struct {
	time.Time
}

const bahnDateLayout = "2006-01-02"

func (t *bahnDate) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if t == nil || t.IsZero() {
		return xml.Attr{}, nil
	} else {
		return xml.Attr{
			Name:  name,
			Value: t.Format(bahnDateLayout),
		}, nil
	}
}

func (t *bahnDate) UnmarshalXMLAttr(attr xml.Attr) error {
	if attr.Value != "" {
		value, err := time.Parse(bahnDateLayout, attr.Value)
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
		text = t.Format(bahnDateLayout)
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
		if value, err = time.Parse(bahnDateLayout, text); err != nil {
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
