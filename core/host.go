package core

const (
	Host      = "https://api.beget.com/api/" // prefix of api endpoint
	UserAgent = "beget-go-server"            // name of user agent
)

var testHost = ""

func SetTestHost(host string) {
	testHost = host
}

type Mode uint64

const (
	Test Mode = iota
	Prod
)

var mode Mode = Prod

func SetMode(md Mode) {
	mode = md
}

func GetMode() Mode {
	return mode
}
