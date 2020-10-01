package config

import "github.com/timoth-y/kicksware-api/service-common/core/meta"

type CommonConfig struct {
	Host               string `yaml:"host"`
	HostName           string `yaml:"hostname"`
	ContentType        string `yaml:"contentType"`
	UsedDB             string `yaml:"usedDB"`
	InnerServiceFormat string `yaml:"innerServiceFormat"`
}

type SecurityConfig struct {
	TLSCertificate     *meta.TLSCertificate `yaml:"tlsCertificate"`
}

type AuthConfig struct {
	PublicKeyPath  string               `yaml:"publicKeyPath"`
	AuthEndpoint   string               `yaml:"authEndpoint"`
	TLSCertificate *meta.TLSCertificate `yaml:"tlsCertificate"`
}

type DataStoreConfig struct {
	URL              string `yaml:"URL"`
	TLS              *meta.TLSCertificate `yaml:"TLS"`
	Database   string `yaml:"database"`
	Collection string `yaml:"collection"`
	Login      string `yaml:"login"`
	Password   string `yaml:"password"`
	Timeout    int    `yaml:"timeout"`
}