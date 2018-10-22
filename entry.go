package log

import (
	"fmt"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

// assert interface compliance.
var _ Interface = (*Entry)(nil)

// Now returns the current time.
var Now = time.Now

// Entry represents a single log entry.
type Entry struct {
	Logger    *Logger   `json:"-"`
	Fields    Fields    `json:"context"`
	Level     int       `json:"level"`
	LevelName string    `json:"level_name"`
	Timestamp time.Time `json:"timestamp"`
	Message   string    `json:"message"`
	start     time.Time
	fields    []Fields
	Env       string `json:"env"`
	Project   string `json:"project"`
	Hostname  string `json:"hostname"`
}

// NewEntry returns a new entry for `log`.
func NewEntry(log *Logger) *Entry {
	return &Entry{
		Logger: log,
	}
}

// WithFields returns a new entry with `fields` set.
func (e *Entry) WithFields(fields Fielder) *Entry {
	f := []Fields{}
	f = append(f, e.fields...)
	f = append(f, fields.Fields())
	return &Entry{
		Logger:   e.Logger,
		fields:   f,
		Hostname: e.Hostname,
		Env:      e.Env,
		Project:  e.Project,
	}
}

// SetEnvProject returns new instance with important info set up
func (e *Entry) SetEnvProject(env string, project string) *Entry {
	e.Hostname, _ = os.Hostname()
	e.Env = env
	e.Project = project
	return &Entry{
		Hostname: e.Hostname,
		Env:      e.Env,
		Project:  e.Project,
		Logger:   e.Logger,
	}
}

// WithField returns a new entry with the `key` and `value` set.
func (e *Entry) WithField(k string, v interface{}) *Entry {
	return e.WithFields(Fields{k: v})
}

// WithError returns a new entry with the "error" set to `err`.
// The given error may implement .Fielder, if it does the method
// will add all its `.Fields()` into the returned entry.
func (e *Entry) WithError(err error) *Entry {
	var errorMap = map[string]interface{}{
		"trace":   string(debug.Stack()),
		"message": err.Error(),
	}
	ctx := e.WithField("error", errorMap)
	if f, ok := err.(Fielder); ok {
		ctx = ctx.WithFields(f.Fields())
	}
	return ctx
}

/********* Entry simple *********/
/* WARNING: This part is autogenerated! Do not change! */

// Debug level message.
func (e *Entry) Debug(msg string) {
	e.Logger.log(DebugLevel, e, msg)
}

// Info level message.
func (e *Entry) Info(msg string) {
	e.Logger.log(InfoLevel, e, msg)
}

// Notice level message.
func (e *Entry) Notice(msg string) {
	e.Logger.log(NoticeLevel, e, msg)
}

// Warn level message.
func (e *Entry) Warn(msg string) {
	e.Logger.log(WarnLevel, e, msg)
}

// Error level message.
func (e *Entry) Error(msg string) {
	e.Logger.log(ErrorLevel, e, msg)
}

// Critical level message.
func (e *Entry) Critical(msg string) {
	e.Logger.log(CriticalLevel, e, msg)
}

// Alert level message.
func (e *Entry) Alert(msg string) {
	e.Logger.log(AlertLevel, e, msg)
}

// Fatal level message.
func (e *Entry) Fatal(msg string) {
	e.Logger.log(FatalLevel, e, msg)
}

// Emergency level message.
func (e *Entry) Emergency(msg string) {
	e.Logger.log(EmergencyLevel, e, msg)
}

/* END OF WARNING */
/********* End Entry simple *********/

/********* Entry formated *********/
/* WARNING: This part is autogenerated! Do not change! */

// Debugf level formatted message.
func (e *Entry) Debugf(msg string, v ...interface{}) {
	e.Debug(fmt.Sprintf(msg, v...))
}

// Infof level formatted message.
func (e *Entry) Infof(msg string, v ...interface{}) {
	e.Info(fmt.Sprintf(msg, v...))
}

// Noticef level formatted message.
func (e *Entry) Noticef(msg string, v ...interface{}) {
	e.Notice(fmt.Sprintf(msg, v...))
}

// Warnf level formatted message.
func (e *Entry) Warnf(msg string, v ...interface{}) {
	e.Warn(fmt.Sprintf(msg, v...))
}

// Errorf level formatted message.
func (e *Entry) Errorf(msg string, v ...interface{}) {
	e.Error(fmt.Sprintf(msg, v...))
}

// Criticalf level formatted message.
func (e *Entry) Criticalf(msg string, v ...interface{}) {
	e.Critical(fmt.Sprintf(msg, v...))
}

// Alertf level formatted message.
func (e *Entry) Alertf(msg string, v ...interface{}) {
	e.Alert(fmt.Sprintf(msg, v...))
}

// Fatalf level formatted message.
func (e *Entry) Fatalf(msg string, v ...interface{}) {
	e.Fatal(fmt.Sprintf(msg, v...))
}

// Emergencyf level formatted message.
func (e *Entry) Emergencyf(msg string, v ...interface{}) {
	e.Emergency(fmt.Sprintf(msg, v...))
}

/* END OF WARNING */
/********* End Entry formated *********/

// Trace returns a new entry with a Stop method to fire off
// a corresponding completion log, useful with defer.
func (e *Entry) Trace(msg string) *Entry {
	e.Info(msg)
	v := e.WithFields(e.Fields)
	v.Message = msg
	v.start = time.Now()
	return v
}

// Stop should be used with Trace, to fire off the completion message. When
// an `err` is passed the "error" field is set, and the log level is error.
func (e *Entry) Stop(err *error) {
	if err == nil || *err == nil {
		e.WithField("duration", time.Since(e.start)).Info(e.Message)
	} else {
		e.WithField("duration", time.Since(e.start)).WithError(*err).Error(e.Message)
	}
}

// mergedFields returns the fields list collapsed into a single map.
func (e *Entry) mergedFields() Fields {
	f := Fields{}

	for _, fields := range e.fields {
		for k, v := range fields {
			for _, hook := range hooks {
				if hook.Check(strings.ToLower(k)) {
					v = hook.Sanitize(v)
				}
			}
			f[k] = v
		}
	}

	return f
}

// finalize returns a copy of the Entry with Fields merged.
func (e *Entry) finalize(level Level, msg string) *Entry {
	return &Entry{
		Logger:    e.Logger,
		Fields:    e.mergedFields(),
		Level:     level.Int(),
		LevelName: level.String(),
		Message:   msg,
		Timestamp: Now(),
		Env:       e.Env,
		Project:   e.Project,
		Hostname:  e.Hostname,
	}
}
