#!/usr/bin/env python3

func_names = ["debug", "info", "notice", "warn", "error", "critical", "alert", "emergency", "fatal"]

#########################
# Entry simple messages #
#########################
print("/********* Entry simple *********/")
for func in func_names:
	f = func.title()
	print("""// %s level message.
func (e *Entry) %s(msg string) {
	e.Logger.log(%sLevel, e, msg)
}
""" % (f, f, f))

###########################
# Entry formated messages #
###########################
print("/********* Entry formated *********/")
for func in func_names:
	f = func.title()
	print("""// %sf level formatted message.
func (e *Entry) %sf(msg string, v ...interface{}) {
	e.%s(fmt.Sprintf(msg, v...))
}
""" % (f, f, f))


#######################
# Log simple messages #
#######################
print("/********* Log simple *********/")
for func in func_names:
	f = func.title()
	print("""// %s level message.
func (l *Logger) %s(msg string) {
	NewEntry(l).%s(msg)
}
""" % (f, f, f))

#########################
# Log formated messages #
#########################
print("/********* Log formated *********/")
for func in func_names:
	f = func.title()
	print("""// %sf level formatted message.
func (l *Logger) %sf(msg string, v ...interface{}) {
	NewEntry(l).%sf(msg, v...)
}
""" % (f, f, f))

############################
# Pkg wide simple messages #
############################
print("/********* Pkg wide simple *********/")
for func in func_names:
	f = func.title()
	print("""// %s level message.
func %s(msg string) {
	Log.%s(msg)
}
""" % (f, f, f))

##############################
# Pkg wide formated messages #
##############################
print("/********* Pkg wide formated *********/")
for func in func_names:
	f = func.title()
	print("""// %sf level message.
func %sf(msg string, v ...interface{}) {
	Log.%sf(msg, v...)
}
""" % (f, f, f))

#############################
# Interface simple messages #
#############################
print("/********* Interface simple *********/")
for func in func_names:
	f = func.title()
	print("""	%s(msg string)""" % (f))

###############################
# Interface formated messages #
###############################
print("/********* Interface formated *********/")
for func in func_names:
	f = func.title()
	print("""	%sf(msg string, v ...interface{})""" % (f))

