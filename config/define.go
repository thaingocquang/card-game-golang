package config

// Database ...
type Database struct {
	Uri      string
	Name     string
	TestName string
}

// Jwt ...
type Jwt struct {
	SecretKey string
}

// ENV ...
type ENV struct {
	// AppPort ...
	AppPort string

	// Database ...
	Database Database

	// Jwt
	Jwt Jwt
}
