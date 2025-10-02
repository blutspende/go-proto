package astmmodels

import (
	"time"

	"github.com/blutspende/bloodlab-common/encoding"
	"github.com/blutspende/bloodlab-common/timezone"
	"github.com/blutspende/go-astm/v3/enums/lineseparator"
	"github.com/blutspende/go-astm/v3/enums/notation"
)

// Configuration struct for the whole process
type Configuration struct {
	Encoding                   encoding.Encoding
	LineSeparator              string
	AutoDetectLineSeparator    bool
	TimeZone                   timezone.TimeZone
	EnforceSequenceNumberCheck bool
	Notation                   string
	DefaultDecimalPrecision    int
	RoundLastDecimal           bool
	KeepShortDateTimeZone      bool
	EscapeOutputStrings        bool
	Delimiters                 Delimiters
	TimeLocation               *time.Location
}

var DefaultConfiguration = Configuration{
	Encoding:                   encoding.ISO8859_1,
	LineSeparator:              lineseparator.LF,
	AutoDetectLineSeparator:    true,
	TimeZone:                   timezone.EuropeBerlin,
	EnforceSequenceNumberCheck: true,
	Notation:                   notation.Standard,
	DefaultDecimalPrecision:    3,
	RoundLastDecimal:           true,
	KeepShortDateTimeZone:      true,
	EscapeOutputStrings:        false,
	Delimiters:                 DefaultDelimiters,
	TimeLocation:               nil,
}

// Delimiters used in ASTM parsing
type Delimiters struct {
	Field        string
	Repeat       string
	Component    string
	SubComponent string
	Escape       string
}

// TODO: rename and reorganise astm and hl7 default delimiters

var DefaultDelimiters = Delimiters{
	Field:     `|`,
	Repeat:    `\`,
	Component: `^`,
	Escape:    `&`,
}
var HL7DefaultDelimiters = Delimiters{
	Field:        `|`,
	Repeat:       `~`,
	Component:    `^`,
	SubComponent: `&`,
	Escape:       `\`,
}
