package config

type AuthConfig struct {
	SecretKey string
	User      string
	Password  string
}

func LoadAuthConfig() *AuthConfig {
	return &AuthConfig{
		SecretKey: "secretKey",
		User:      "admin",
		Password:  "password",
	}
}
