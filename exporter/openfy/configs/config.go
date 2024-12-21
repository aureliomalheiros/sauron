package configs

var Config = struct {
	APIURL 		string
	HTTPPort 	string
}{
	APIURL: "http://api.open-notify.org/astros.json",
	HTTPPort: "7777",
}
