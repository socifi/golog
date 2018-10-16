# Go logging facility

Based on [Apex log](https://github.com/apex/log) with few tweaks. Unless otherwise noted (mainly in `pkg.go`), this package should be thread safe. Just for the record the architecture leaves almost all thread safety on handlers so if you are uncertain, consult the documentation and code of handlers you use.

## Functionality

### Autogenerated levels

This package contains python script for log level autogeneration. The script contains default values used in SOCIFI but can be easily tweaked, just alter level names and level codes (should be integer value, not a string) in file `gen_levels.py`

### Password hiding in dsns

If a field name contains `dsn` (case insensitive), this package will try to hide password in dsn string. Rationale behind this is that dsn might be very useful in troubleshooting some problems but logging passwords is very bad practice. This feature is however not intended to prevent irresponsibility of a user in any way. So there will be no checking for e.g. field names containing `password`.

### Simple logger initialization from config json

There is also initialization script which allows simple load of configuration needed by logger. This feature is now limited to json, text and elastic handlers and might change in the future to allow better interoperability with e.g. [Viper](https://github.com/spf13/viper)

## Handlers

- __cli__ – human-friendly CLI output
- __discard__ – discards all logs
- __es__ – Elasticsearch handler
- __graylog__ – Graylog handler
- __json__ – JSON output handler
- __kinesis__ – AWS Kinesis handler
- __level__ – level filter handler
- __logfmt__ – logfmt plain-text formatter
- __memory__ – in-memory handler for tests
- __multi__ – fan-out to multiple handlers
- __papertrail__ – Papertrail handler
- __text__ – human-friendly colored output
- __delta__ – outputs the delta between log calls and spinner

