package astm

import (
	"github.com/blutspende/bloodlab-common/encoding"
	"github.com/blutspende/go-astm/v3/functions"
	"github.com/blutspende/go-astm/v3/models/astmmodels"
	"github.com/blutspende/go-astm/v3/utils"
)

func Marshal(sourceStruct interface{}, config *astmmodels.Configuration) (result [][]byte, err error) {
	// Init configuration
	err = utils.InitConfig(config)
	if err != nil {
		return nil, err
	}
	// Build the lines from the source structure
	lines, err := functions.BuildStruct(sourceStruct, 1, 0, config)
	if err != nil {
		return nil, err
	}
	// Convert UTF8 string array to encoding
	result, err = encoding.ConvertArrayFromUtf8ToEncoding(lines, config.Encoding)
	if err != nil {
		return nil, err
	}
	// Return the result and no error if everything went well
	return result, nil
}
