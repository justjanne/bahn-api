package main

import (
	"encoding/json"
	"encoding/xml"
	"strings"
	"time"
)

type BahnStringList []string

func (s *BahnStringList) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if s == nil {
		return xml.Attr{}, nil
	} else {
		return xml.Attr{
			Name:  name,
			Value: strings.Join(*s, "|"),
		}, nil
	}
}

func (s *BahnStringList) UnmarshalXMLAttr(attr xml.Attr) error {
	*s = strings.Split(attr.Value, "|")
	return nil
}

func (s *BahnStringList) Value() []string {
	if s != nil {
		return *s
	} else {
		return make([]string, 0)
	}
}

type BahnTime struct {
	time.Time
}

const BahnTimeLayout = "06-01-02 15:04:05.999"

func (t *BahnTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if t == nil || t.IsZero() {
		return xml.Attr{}, nil
	} else {
		return xml.Attr{
			Name:  name,
			Value: t.Format(BahnTimeLayout),
		}, nil
	}
}

func (t *BahnTime) UnmarshalXMLAttr(attr xml.Attr) error {
	if attr.Value != "" {
		value, err := time.Parse(BahnTimeLayout, attr.Value)
		if err != nil {
			return err
		}
		t.Time = value
	}
	return nil
}

func (t *BahnTime) Value() *time.Time {
	if t != nil {
		return &t.Time
	} else {
		return nil
	}
}

type ShortBahnTime struct {
	time.Time
}

const ShortBahnTimeLayout = "0601021504"

func (t *ShortBahnTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if t == nil || t.IsZero() {
		return xml.Attr{}, nil
	} else {
		return xml.Attr{
			Name:  name,
			Value: t.Format(ShortBahnTimeLayout),
		}, nil
	}
}

func (t *ShortBahnTime) UnmarshalXMLAttr(attr xml.Attr) error {
	if attr.Value != "" {
		value, err := time.Parse(ShortBahnTimeLayout, attr.Value)
		if err != nil {
			return err
		}
		t.Time = value
	}
	return nil
}

func (t *ShortBahnTime) Value() *time.Time {
	if t != nil {
		return &t.Time
	} else {
		return nil
	}
}

type MediumBahnTime struct {
	time.Time
}

const MediumBahnTimeLayout = "2006-01-02T15:04:05"

func (t *MediumBahnTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if t == nil || t.IsZero() {
		return xml.Attr{}, nil
	} else {
		return xml.Attr{
			Name:  name,
			Value: t.Format(MediumBahnTimeLayout),
		}, nil
	}
}

func (t *MediumBahnTime) UnmarshalXMLAttr(attr xml.Attr) error {
	if attr.Value != "" {
		value, err := time.Parse(MediumBahnTimeLayout, attr.Value)
		if err != nil {
			return err
		}
		t.Time = value
	}
	return nil
}

func (t *MediumBahnTime) MarshalJSON() ([]byte, error) {
	var text string
	if t == nil || t.IsZero() {
		text = ""
	} else {
		text = t.Format(MediumBahnTimeLayout)
	}
	return json.Marshal(&text)
}

func (t *MediumBahnTime) UnmarshalJSON(data []byte) error {
	var err error

	var text string
	if err := json.Unmarshal(data, &text); err != nil {
		return err
	}

	if text != "" {
		var value time.Time
		if value, err = time.Parse(MediumBahnTimeLayout, text); err != nil {
			return err
		}
		t.Time = value
	}
	return nil
}

func (t *MediumBahnTime) Value() *time.Time {
	if t != nil {
		return &t.Time
	} else {
		return nil
	}
}

type BahnDate struct {
	time.Time
}

const BahnDateLayout = "2006-01-02"

func (t *BahnDate) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if t == nil || t.IsZero() {
		return xml.Attr{}, nil
	} else {
		return xml.Attr{
			Name:  name,
			Value: t.Format(BahnDateLayout),
		}, nil
	}
}

func (t *BahnDate) UnmarshalXMLAttr(attr xml.Attr) error {
	if attr.Value != "" {
		value, err := time.Parse(BahnDateLayout, attr.Value)
		if err != nil {
			return err
		}
		t.Time = value
	}
	return nil
}

func (t *BahnDate) MarshalJSON() ([]byte, error) {
	var text string
	if t == nil || t.IsZero() {
		text = ""
	} else {
		text = t.Format(BahnDateLayout)
	}
	return json.Marshal(&text)
}

func (t *BahnDate) UnmarshalJSON(data []byte) error {
	var err error

	var text string
	if err := json.Unmarshal(data, &text); err != nil {
		return err
	}

	if text != "" {
		var value time.Time
		if value, err = time.Parse(BahnDateLayout, text); err != nil {
			return err
		}
		t.Time = value
	}
	return nil
}

func (t *BahnDate) Value() *time.Time {
	if t != nil {
		return &t.Time
	} else {
		return nil
	}
}
