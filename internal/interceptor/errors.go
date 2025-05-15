package interceptor

var errors = map[string]string{
	"GUPLD001": "Invalid request body",

	"GUPLD101": "Unable to register user",
	"GUPLD102": "Invalid username or password",
	"GUPLD103": "Username must be between 3 and 30 characters",
	"GUPLD104": "Password must be at least 8 characters long",
	"GUPLD105": "Password cannot be your username",
}
