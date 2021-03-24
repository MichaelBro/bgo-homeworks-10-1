package currencies

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"io/ioutil"
	"os"
)

type ValCurs struct {
	XMLName xml.Name `xml:"ValCurs" json:"-"`
	Date    string   `xml:"Date,attr" ,json:"date"`
	Name    string   `xml:"name,attr" ,json:"name"`
	Valute  []Valute `xml:"Valute" ,json:"valute"`
}

type Valute struct {
	ID       string  `xml:"ID,attr" ,json:"id"`
	NumCode  string  `xml:"NumCode" ,json:"num_code"`
	CharCode string  `xml:"CharCode" ,json:"char_code"`
	Nominal  float64 `xml:"Nominal" ,json:"nominal"`
	Name     string  `xml:"Name" ,json:"name"`
	Value    float64 `xml:"Value" ,json:"value"`
}

func ImportFromXml(reader io.Reader) (ValCurs, error) {
	file, err := ioutil.ReadAll(reader)

	if err != nil {
		return ValCurs{}, err
	}

	var decoded ValCurs
	err = xml.Unmarshal(file, &decoded)

	if err != nil {
		return ValCurs{}, err
	}
	return decoded, nil
}

func ExportToJson(v ValCurs) error {
	jsonString, err := json.Marshal(v)

	if err != nil {
		return err
	}

	err = ioutil.WriteFile("currencies.json", jsonString, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func XmlToJson(reader io.Reader) error {
	valCurs, err := ImportFromXml(reader)
	if err != nil {
		return err
	}
	err = ExportToJson(valCurs)
	if err != nil {
		return err
	}
	return nil
}
