package types

import (
	"sync"
	"time"
)

//Row struct
type Row struct {
	BOARDID                 string `xml:"BOARDID,attr" json:"BOARDID"`
	TRADEDATE               string `xml:"TRADEDATE,attr" json:"TRADEDATE"`
	SHORTNAME               string `xml:"SHORTNAME,attr" json:"SHORTNAME"`
	SECID                   string `xml:"SECID,attr" json:"SECID "`
	NUMTRADES               string `xml:"NUMTRADES,attr" json:"NUMTRADES"`
	VALUE                   string `xml:"VALUE,attr" json:"VALUE"`
	OPEN                    string `xml:"OPEN,attr" json:"OPEN"`
	LOW                     string `xml:"LOW,attr" json:"LOW"`
	HIGH                    string `xml:"HIGH,attr" json:"HIGH"`
	LEGALCLOSEPRICE         string `xml:"LEGALCLOSEPRICE,attr" json:"LEGALCLOSEPRICE"`
	WAPRICE                 string `xml:"WAPRICE,attr" json:"WAPRICE"`
	CLOSE                   string `xml:"CLOSE,attr" json:"CLOSE"`
	VOLUME                  string `xml:"VOLUME,attr" json:"VOLUME"`
	MARKETPRICE2            string `xml:"MARKETPRICE2,attr" json:"MARKETPRICE2"`
	MARKETPRICE3            string `xml:"MARKETPRICE3,attr" json:"MARKETPRICE3"`
	ADMITTEDQUOTE           string `xml:"ADMITTEDQUOTE,attr" json:"ADMITTEDQUOTE"`
	MP2VALTRD               string `xml:"MP2VALTRD,attr" json:"MP2VALTRD"`
	MARKETPRICE3TRADESVALUE string `xml:"MARKETPRICE3TRADESVALUE,attr" json:"MARKETPRICE3TRADESVALUE"`
	ADMITTEDVALUE           string `xml:"ADMITTEDVALUE,attr" json:"ADMITTEDVALUE"`
	WAVAL                   string `xml:"WAVAL,attr" json:"WAVAL"`
	TRADINGSESSION          string `xml:"TRADINGSESSION,attr" json:"TRADINGSESSION"`
}

//Rows struct
type Rows struct {
	Rows []Row `xml:"row"`
}

//Column struct
type Column struct {
	Name       string `xml:"name,attr"`
	ColumnType string `xml:"type,attr"`
}

//Columns struct
type Columns struct {
	Columns []Column `xml:"column"`
}

//Metadata struct
type Metadata struct {
	Columns Columns `xml:"columns"`
}

//Data struct
type Data struct {
	ID       string   `xml:"id,attr"`
	Metadata Metadata `xml:"metadata"`
	Rows     Rows     `xml:"rows"`
}

//Document struct
type Document struct {
	Data []Data `xml:"data"`
}

//Cache struct
type Cache struct {
	LastUpdate time.Time
	Data       Document
	RWMutex    sync.RWMutex
}
