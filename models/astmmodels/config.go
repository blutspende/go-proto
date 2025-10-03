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
	Protocol                   Protocol
	ProtocolVersion            string
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

var DefaultConfigurationASTM = Configuration{
	Protocol:                   ASTM,
	ProtocolVersion:            "",
	Encoding:                   encoding.UTF8,
	LineSeparator:              lineseparator.LF,
	AutoDetectLineSeparator:    true,
	TimeZone:                   timezone.EuropeBerlin,
	EnforceSequenceNumberCheck: true,
	Notation:                   notation.Standard,
	DefaultDecimalPrecision:    3,
	RoundLastDecimal:           true,
	KeepShortDateTimeZone:      true,
	EscapeOutputStrings:        false,
	Delimiters:                 DefaultDelimitersASTM,
	TimeLocation:               nil,
}

var DefaultConfigurationHL7 = Configuration{
	Protocol:                   HL7,
	ProtocolVersion:            "",
	Encoding:                   encoding.UTF8,
	LineSeparator:              lineseparator.LF,
	AutoDetectLineSeparator:    true,
	TimeZone:                   timezone.EuropeBerlin,
	EnforceSequenceNumberCheck: true,
	Notation:                   notation.Standard,
	DefaultDecimalPrecision:    3,
	RoundLastDecimal:           true,
	KeepShortDateTimeZone:      true,
	EscapeOutputStrings:        false,
	Delimiters:                 DefaultDelimitersHL7,
	TimeLocation:               nil,
}

type Protocol string

const ASTM Protocol = "ASTM"
const HL7 Protocol = "HL7"

// Delimiters used in ASTM parsing
type Delimiters struct {
	Field        string
	Repeat       string
	Component    string
	SubComponent string
	Escape       string
}

var DefaultDelimitersASTM = Delimiters{
	Field:     `|`,
	Repeat:    `\`,
	Component: `^`,
	Escape:    `&`,
}
var DefaultDelimitersHL7 = Delimiters{
	Field:        `|`,
	Repeat:       `~`,
	Component:    `^`,
	SubComponent: `&`,
	Escape:       `\`,
}
