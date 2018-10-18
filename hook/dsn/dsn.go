package dsn

import(
	"github.com/socifi/go-logging-facility"
	"regexp"
	"strings"
)

type dsn struct{}

var dsnHook dsn

func init() {
	log.RegisterSanitizeHook(dsnHook)
}

// Check checks if dsn is in the string
func (dsn) Check(s string) bool {
	if strings.Contains(strings.ToLower(s), "dsn") {
		return true
	}
	return false
}

// Sanitize tries its best to replace all passwords in a string which has dsn format
func (dsn) Sanitize(v interface{}) (interface{}) {
	s, ok := v.(string)
	if !ok {
		return v
	}

	r := regexp.MustCompile(`(.*?://)*(.+?:)(.+?)(@)`)
	if r.MatchString(s) {
		s = r.ReplaceAllString(s, "${1}${2}***${4}")
	}

	r = regexp.MustCompile(`([\s;])(p|password|pwd|pass|P|PASSWORD|PWD|PASS|Password|Pwd|Pass)([=:])(.*?)([\s;])`)
	if r.MatchString(s) {
		s = r.ReplaceAllString(s, "${1}${2}${3}***${5}")
	}

	return s
}
