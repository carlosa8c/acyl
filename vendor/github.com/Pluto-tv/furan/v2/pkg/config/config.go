package config

/*
"secret" struct tags are secret ids for PVC auto-fill
*/

type VaultConfig struct {
	Addr        string
	Token       string
	TokenAuth   bool
	K8sAuth     bool
	K8sJWTPath  string
	K8sAuthPath string
	K8sRole     string
}

type GitHubConfig struct {
	Token string `secret:"GITHUB_TOKEN"`
}

type QuayConfig struct {
	Token string `secret:"QUAY_TOKEN"`
}

// AWSConfig contains all information needed to access AWS services
// AWS credentials scoped to S3 Read/Write & ECR Push/Pull only
type AWSConfig struct {
	Region           string
	CacheBucket      string
	CacheKeyPrefix   string
	AccessKeyID      string `secret:"AWS_ACCESS_KEY_ID"`
	SecretAccessKey  string `secret:"AWS_SECRET_ACCESS_KEY"`
	EnableECR        bool
	ECRRegistryHosts []string
}

type DBConfig struct {
	PostgresURI             string `secret:"DB_URI"`
	CredentialEncryptionKey []byte `secret:"DB_CREDENTIAL_ENCRYPTION_KEY"`
	CredEncKeyArray         [32]byte
}

type APMConfig struct {
	Addr             string
	App, Environment string
	APM, Profiling   bool
}

type ServerConfig struct {
	HTTPSAddr string
	GRPCAddr  string
	SeedKey   string
}
