package parser

import (
	"testing"

	"github.com/Matir/ctftk/testdata"
)

func TestParserMain(t *testing.T) {
	t.Logf("Testdata: %s", testdata.GetTestdataDir())
}

func TestValidConfigLoads(t *testing.T) {
	fp := testdata.MustOpenTestdataFile("rootconfig/config.yml")
	defer fp.Close()
	cfg, err := ReadRootConfig(fp)
	if err != nil {
		t.Fatalf("Attempted to load valid config, got error: %s", err.Error())
	}
	if cfg == nil {
		t.Fatalf("Attempted to load valid config, but got nil config.")
	}
}

func TestInvalidConfigs(t *testing.T) {
	badConfigs := []string{
		"rootconfig/config_noversion.yml",
	}
	for _, cfgPath := range badConfigs {
		// This function is here so each defer will be closed within that iteration
		// of the loop.
		func() {
			fp := testdata.MustOpenTestdataFile(cfgPath)
			defer fp.Close()
			cfg, err := ReadRootConfig(fp)
			if err == nil {
				t.Fatalf("Attempted to load invalid config, but got no error!")
			}
			if cfg != nil {
				t.Fatalf("Attempted to load invalid config, but got a config back.")
			}
			if _, ok := err.(ConfigError); !ok {
				t.Fatalf("Loading an invalid config, got error, but not of type ConfigError: %v", err)
			}
		}()
	}
}
