package config

type AuthConfig struct {
	Guards map[string]GuardConfig
}

type GuardConfig struct {
	Driver   string
	Provider string
}

func LoadAuthConfig() *AuthConfig {
	return &AuthConfig{
		Guards: map[string]GuardConfig{
			"web": {
				Driver:   "session",
				Provider: "users",
			},
			"api": {
				Driver:   "token",
				Provider: "users",
			},
		},
	}

}
