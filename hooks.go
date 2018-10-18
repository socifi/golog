package log

// Hook is an interface which has two functions.
// It is intended to be used internaly on map[string]interface{}
// Check checks a map string for a condition and if true is returned, Sanitize is called with a relevant value 
type Hook interface {
	Check(string) bool
	Sanitize(interface{}) interface{}
}

var hooks []Hook

// RegisterSanitizeHook adds a hook interface to internal queue for entry processing
func RegisterSanitizeHook(h Hook) {
	hooks = append(hooks, h)
}
