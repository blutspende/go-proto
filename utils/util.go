package utils

import (
	"errors"

	"github.com/blutspende/go-astm/v3/models/astmmodels"
)

func NewDefaultConfiguration(protocol astmmodels.Protocol) *astmmodels.Configuration {
	var newConfig astmmodels.Configuration
	switch protocol {
	case astmmodels.ASTM:
		newConfig = astmmodels.DefaultConfigurationASTM
	case astmmodels.HL7:
		newConfig = astmmodels.DefaultConfigurationHL7
	default:
		return nil
	}
	return &newConfig
}

func InitConfig(config *astmmodels.Configuration) (err error) {
	if config == nil {
		return errors.New("nil configuration")
	}
	if config.Protocol != astmmodels.ASTM && config.Protocol != astmmodels.HL7 {
		return errors.New("invalid protocol")
	}
	if config.Delimiters.Field == "" ||
		config.Delimiters.Repeat == "" ||
		config.Delimiters.Component == "" ||
		config.Delimiters.Escape == "" ||
		(config.Protocol == astmmodels.HL7 && config.Delimiters.SubComponent == "") {
		return errors.New("missing delimiter(s)")
	}
	config.TimeLocation, err = config.TimeZone.GetLocation()
	return err
}
