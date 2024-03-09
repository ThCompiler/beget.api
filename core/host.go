package core

const (
	Host      = "https://api.beget.com/api/" // prefix of api endpoint
	UserAgent = "beget-go-server"            // name of user agent
)

var testHost = ""

// SetTestHost sets the Beget API host for testing.
func SetTestHost(host string) {
	testHost = host
}

// Mode is the mode of operation of the package.
// If test mode is set, the package will use the test host as the host for the request to the Beget API.
type Mode uint64

const (
	Test Mode = iota // Test mode
	Prod             // Prod mode
)

var mode = Prod

// SetMode sets the mode of operation of the package.
func SetMode(md Mode) {
	mode = md
}

// GetMode returns the mode of operation of the package.
func GetMode() Mode {
	return mode
}
