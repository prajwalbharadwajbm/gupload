package interceptor

var errors = map[string]string{
	"GUPLD001": "Invalid request body",
	"GUPLD002": "Internal Server Error",

	"GUPLD101": "Unable to register user",
	"GUPLD102": "Invalid username or password",
	"GUPLD103": "Username must be between 3 and 30 characters",
	"GUPLD104": "Password must be at least 8 characters long",
	"GUPLD105": "Password cannot be your username",
	"GUPLD106": "Failed to generate authentication token",
	"GUPLD107": "Authentication Error",
	"GUPLD108": "Invalid User",

	"GUPLD201": "Unable to get file",
	"GUPLD202": "Unable to upload file",
	"GUPLD203": "Storage Quota Full",
}
