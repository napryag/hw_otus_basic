package types

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"encoding/xml"

	"go.mongodb.org/mongo-driver/v2/bson"
	"gopkg.in/yaml.v3"
)

type Book struct {
	ID     int     `json:"id" xml:"id" yaml:"id" bson:"id"`
	Title  string  `json:"title" xml:"title" yaml:"title" bson:"title"`
	Author string  `json:"author" xml:"author" yaml:"author" bson:"author"`
	Year   int     `json:"year" xml:"year" yaml:"year" bson:"year"`
	Size   int     `json:"size" xml:"size" yaml:"size" bson:"size"`
	Rate   float32 `json:"rate" xml:"rate" yaml:"rate" bson:"rate"`
	Sample []byte
}

func EncodeGOB(b *Book) (bytes.Buffer, error) {
	var buff bytes.Buffer
	enc := gob.NewEncoder(&buff)
	err := enc.Encode(b)
	return buff, err
}

func DecodeGOB(a bytes.Buffer, b *Book) error {
	dec := gob.NewDecoder(&a)
	err := dec.Decode(&b)
	return err
}

func MarshalJSON(a *Book) ([]byte, error) {
	b, err := json.Marshal(a)
	return b, err
}

func UnmarshalJSON(b []byte, c *Book) error {
	err := json.Unmarshal(b, &c)
	return err
}

func MarshalXML(a *Book) ([]byte, error) {
	b, err := xml.Marshal(a)
	return b, err
}

func UnmarshalXML(b []byte, c *Book) error {
	err := xml.Unmarshal(b, &c)
	return err
}

func MarshalYAML(a *Book) ([]byte, error) {
	b, err := yaml.Marshal(a)
	return b, err
}

func UnmarshalYAML(b []byte, c *Book) error {
	err := yaml.Unmarshal(b, &c)
	return err
}

func MarshalBSON(a *Book) ([]byte, error) {
	b, err := bson.Marshal(a)
	return b, err
}

func UnmarshalBSON(b []byte, c *Book) error {
	err := bson.Unmarshal(b, &c)
	return err
}
