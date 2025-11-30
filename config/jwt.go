package config

var jwtSecret []byte

// InitJWTSecret initializes the JWT secret from environment
func InitJWTSecret() {
	jwtSecret = []byte(GetEnv("JWT_SECRET", "default_key"))
}

// GetJWTSecret returns the JWT secret
func GetJWTSecret() []byte {
	if jwtSecret == nil {
		InitJWTSecret()
	}
	return jwtSecret
}
