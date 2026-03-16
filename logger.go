package fastlog

import (
	"reflect"
	"time"

	scolor "github.com/SnowsSky/scolor/v2/pkg"
)

type Logger struct {
	WriteToFile bool
	WithDate    bool
	Colorize    bool
}

var startOfLogger = time.Now()
var DefaultLogger = NewLogger(false, true, true)

func writeSliceWithScolor(s any) {
	if reflect.TypeOf(s).Kind() == reflect.Slice {
		for _, v := range s.([]any) {
			if reflect.TypeOf(v).Kind() == reflect.Slice {
				writeSliceWithScolor(v)
			} else {
				scolor.White.DisplayText(v)
			}
		}
	} else {
		scolor.White.DisplayText(s)
	}
}

/*
Create a new logger.
Persistent means it will write to a "logs" folder
WithDate will display the date in the terminal
Disabling Colorize will disable colors
*/
func NewLogger(Persistent bool, WithDate bool, Colorize bool) *Logger {
	return &Logger{
		WriteToFile: Persistent,
		WithDate:    WithDate,
		Colorize:    Colorize,
	}
}

/*
Print an error message, and exits with code -1
*/
func (l *Logger) Error(args ...any) {
	if l.WithDate {
		scolor.White.DisplayText(time.Now().Format(time.DateTime), " ")
	}
	if l.Colorize {
		scolor.DisplayText(scolor.BgRed, "ERROR")
	} else {
		scolor.White.DisplayText("ERROR")
	}
	scolor.White.DisplayText(": ")
	for _, arg := range args {
		writeSliceWithScolor(arg)
	}
	println()
	if l.WriteToFile {
		writeLogToFile("ERROR", args)
	}
}

func (l *Logger) Warn(args ...any) {
	if l.WithDate {
		scolor.White.DisplayText(time.Now().Format(time.DateTime), " ")
	}
	if l.Colorize {
		scolor.DisplayText(scolor.BgYellow, "WARNING")
	} else {
		scolor.White.DisplayText("WARNING")
	}
	scolor.White.DisplayText(": ")
	for _, arg := range args {
		writeSliceWithScolor(arg)
	}
	println()
	if l.WriteToFile {
		writeLogToFile("WARNING", args)
	}
}

func (l *Logger) Success(args ...any) {
	if l.WithDate {
		scolor.White.DisplayText(time.Now().Format(time.DateTime), " ")
	}
	if l.Colorize {
		scolor.DisplayText(scolor.BgGreen, "SUCCESS")
	} else {
		scolor.White.DisplayText("SUCCESS")
	}
	scolor.White.DisplayText(": ")
	for _, arg := range args {
		writeSliceWithScolor(arg)
	}
	println()
	if l.WriteToFile {
		writeLogToFile("SUCCESS", args)
	}
}

func (l *Logger) Info(args ...any) {
	if l.WithDate {
		scolor.White.DisplayText(time.Now().Format(time.DateTime), " ")
	}
	if l.Colorize {
		scolor.DisplayText(scolor.BgBlue, "INFO")
	} else {
		scolor.White.DisplayText("INFO")
	}
	scolor.White.DisplayText(": ")
	for _, arg := range args {
		writeSliceWithScolor(arg)
	}
	println()
	if l.WriteToFile {
		writeLogToFile("INFO", args)
	}
}
