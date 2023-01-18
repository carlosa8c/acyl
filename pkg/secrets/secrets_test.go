package secrets

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/Pluto-tv/acyl/pkg/config"
	"github.com/Pluto-tv/pvc"
)

const (
	testSecretsJSONfile = "testdata/fakesecrets.json"
	testSecretPrefix    = "secret/app/test-acyl/"
)

var testGithubConfig = &config.GithubConfig{}
var testSlackConfig = &config.SlackConfig{}
var testServerConfig = &config.ServerConfig{}
var testPGConfig = &config.PGConfig{}

func readSecretsJSON(t *testing.T) map[string]string {
	d, err := os.ReadFile(testSecretsJSONfile)
	if err != nil {
		t.Fatalf("error reading secrets json file: %v", err)
	}
	sm := map[string]string{}
	err = json.Unmarshal(d, &sm)
	if err != nil {
		t.Fatalf("error unmarshaling secrets json: %v", err)
	}
	return sm
}

func TestSecretsPopulateAllSecrets(t *testing.T) {
	sm := readSecretsJSON(t)
	sc, err := pvc.NewSecretsClient(pvc.WithJSONFileBackend(testSecretsJSONfile), pvc.WithMapping(testSecretPrefix+"{{ .ID }}"))
	if err != nil {
		t.Fatalf("error getting secrets client: %v", err)
	}
	psf := PVCSecretsFetcher{sc: sc}
	err = psf.PopulateAllSecrets(testGithubConfig, testSlackConfig, testServerConfig, testPGConfig)
	if err != nil {
		t.Fatalf("should have succeeded: %v", err)
	}
	if testGithubConfig.HookSecret != sm[testSecretPrefix+githubHookSecret] {
		t.Fatalf("bad value for github hook secret: %v", testGithubConfig.HookSecret)
	}
	if testSlackConfig.Token != sm[testSecretPrefix+slackToken] {
		t.Fatalf("bad value for slack token: %v", testSlackConfig.Token)
	}
	if testPGConfig.PostgresURI != sm[testSecretPrefix+dbURI] {
		t.Fatalf("bad value for db uri: %v", testPGConfig.PostgresURI)
	}
}
