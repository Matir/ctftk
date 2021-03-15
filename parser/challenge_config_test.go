package parser

import (
	"testing"

	"github.com/Matir/ctftk/testdata"
)

func TestChallengeParserValid(t *testing.T) {
	valid := []string{
		"challengeconfigs/http.yml",
		"challengeconfigs/tcp.yml",
	}
	for _, fname := range valid {
		func() {
			fp := testdata.MustOpenTestdataFile(fname)
			defer fp.Close()
			cfg, err := ReadChallengeConfig(fp, "test")
			if err != nil {
				t.Fatalf("Reading %s, got unexpected error: %s", fname, err)
			}
			if cfg == nil {
				t.Fatalf("Reading %s, expected config, got nil.", fname)
			}
		}()
	}
}

func TestChallengeParserInvalid(t *testing.T) {
	configs := map[string]string{
		"challengeconfigs/missing_name.yml": "Name is required",
		"challengeconfigs/missing_type.yml": "Invalid challenge type",
		"challengeconfigs/bad_type.yml":     "Invalid challenge type",
		"challengeconfigs/extra_field.yml":  "field unknown_field not found",
	}
	for fname, errExpected := range configs {
		func() {
			fp := testdata.MustOpenTestdataFile(fname)
			defer fp.Close()
			cfg, err := ReadChallengeConfig(fp, "test")
			if err == nil {
				t.Fatalf("Reading %s, expected error, got nil.", fname)
			}
			if _, ok := err.(ConfigError); !ok {
				t.Fatalf("Reading %s, expected config error, got %v", fname, err)
			}
			testdata.ExpectErrorContains(t, err, errExpected)
			if cfg != nil {
				t.Fatalf("Reading %s, expected no config!", fname)
			}
		}()
	}
}
