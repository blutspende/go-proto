package e2e

import (
	"bytes"
	"testing"

	"github.com/blutspende/go-astm/v3/models/astmmodels"
	"github.com/blutspende/go-astm/v3/utils"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
)

// Configuration struct for tests
var config *astmmodels.Configuration

// Reset config to default values
func teardown() {
	config = utils.NewDefaultConfiguration(astmmodels.ASTM)
	_ = utils.InitConfig(config)
}

// Setup default config and run all tests
func TestMain(m *testing.M) {
	// Set up configuration
	teardown()
	// Run all tests
	m.Run()
}

// Encoding helper function
func helperEncode(charmap *charmap.Charmap, data []byte) []byte {
	e := charmap.NewEncoder()
	var b bytes.Buffer
	writer := transform.NewWriter(&b, e)
	writer.Write(data)
	resultdata := b.Bytes()
	writer.Close()
	return resultdata
}
