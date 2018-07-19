package log

// singletons ftw?
var Log Interface = &Logger{
	Handler: HandlerFunc(handleStdLog),
	Level:   InfoLevel,
}

// SetHandler sets the handler. This is not thread-safe.
// The default handler outputs to the stdlib log.
func SetHandler(h Handler) {
	if logger, ok := Log.(*Logger); ok {
		logger.Handler = h
	}
}

//SetAdditionalFields sets the env,project and hostname fields into log.
func SetAdditionalFields(fields Fielder) *Entry {
	return Log.SetAdditionalFields(fields)
}

// SetLevel sets the log level. This is not thread-safe.
func SetLevel(l Level) {
	if logger, ok := Log.(*Logger); ok {
		logger.Level = l
	}
}

// SetLevelFromString sets the log level from a string, panicing when invalid. This is not thread-safe.
func SetLevelFromString(s string) {
	if logger, ok := Log.(*Logger); ok {
		logger.Level = MustParseLevel(s)
	}
}

// WithFields returns a new entry with `fields` set.
func WithFields(fields Fielder) *Entry {
	return Log.WithFields(fields)
}

// WithField returns a new entry with the `key` and `value` set.
func WithField(key string, value interface{}) *Entry {
	return Log.WithField(key, value)
}

// WithError returns a new entry with the "error" set to `err`.
func WithError(err error) *Entry {
	return Log.WithError(err)
}

/********* Pkg wide simple *********/
// Debug level message.
func Debug(msg string) {
	Log.Debug(msg)
}

// Info level message.
func Info(msg string) {
	Log.Info(msg)
}

// Notice level message.
func Notice(msg string) {
	Log.Notice(msg)
}

// Warn level message.
func Warn(msg string) {
	Log.Warn(msg)
}

// Error level message.
func Error(msg string) {
	Log.Error(msg)
}

// Critical level message.
func Critical(msg string) {
	Log.Critical(msg)
}

// Alert level message.
func Alert(msg string) {
	Log.Alert(msg)
}

// Emergency level message.
func Emergency(msg string) {
	Log.Emergency(msg)
}

// Fatal level message.
func Fatal(msg string) {
	Log.Fatal(msg)
}

/********* Pkg wide formated *********/
// Debugf level message.
func Debugf(msg string, v ...interface{}) {
	Log.Debugf(msg, v...)
}

// Infof level message.
func Infof(msg string, v ...interface{}) {
	Log.Infof(msg, v...)
}

// Noticef level message.
func Noticef(msg string, v ...interface{}) {
	Log.Noticef(msg, v...)
}

// Warnf level message.
func Warnf(msg string, v ...interface{}) {
	Log.Warnf(msg, v...)
}

// Errorf level message.
func Errorf(msg string, v ...interface{}) {
	Log.Errorf(msg, v...)
}

// Criticalf level message.
func Criticalf(msg string, v ...interface{}) {
	Log.Criticalf(msg, v...)
}

// Alertf level message.
func Alertf(msg string, v ...interface{}) {
	Log.Alertf(msg, v...)
}

// Emergencyf level message.
func Emergencyf(msg string, v ...interface{}) {
	Log.Emergencyf(msg, v...)
}

// Fatalf level message.
func Fatalf(msg string, v ...interface{}) {
	Log.Fatalf(msg, v...)
}

// Trace returns a new entry with a Stop method to fire off
// a corresponding completion log, useful with defer.
func Trace(msg string) *Entry {
	return Log.Trace(msg)
}
