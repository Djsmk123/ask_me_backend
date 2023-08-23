package firebaseutils

import "github.com/spf13/viper"

type FirebaseSdkConfig struct {
	Type                    string `json:"type"`
	ProjectId               string `json:"project"`
	PrivateKeyId            string `json:"private_key_id"`
	PrivateKey              string `json:"private_key"`
	ClientEmail             string `json:"client_email"`
	ClientId                string `json:"client_id"`
	AuthUri                 string `json:"auth_uri"`
	TokenUri                string `json:"token_uri"`
	AuthProviderX509CertUrl string `json:"auth_provider_x509_cert_url"`
	ClientX509CertUrl       string `json:"client_x509_cert_url"`
	UniverseDomain          string `json:"universe_domain"`
}

func ReadSdk() (*FirebaseSdkConfig, error) {
	viper.AddConfigPath(".")
	viper.SetConfigFile("firebase--admin-sdk.json")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	var sdk *FirebaseSdkConfig
	viper.Unmarshal(&sdk)

	return sdk, nil
}
