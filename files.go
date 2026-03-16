package fastlog

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"time"
)

func writeSliceToWriter(s any, w io.Writer) {
	if reflect.TypeOf(s).Kind() == reflect.Slice {
		for _, v := range s.([]any) {
			if reflect.TypeOf(v).Kind() == reflect.Slice {
				writeSliceToWriter(v, w)
			} else {
				fmt.Fprint(w, v)
			}
		}
	} else {
		fmt.Fprint(w, s)
	}
}

// TODO: Maybe open file once, and close it when logger finish, not sure tho
func writeLogToFile(level string, args ...any) {
	logFile, err := os.OpenFile(startOfLogger.Format(time.RFC3339)+".log", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		DefaultLogger.Error("Couldn't write to "+startOfLogger.String()+".log : ", err.Error())
		return
	}
	writeSliceToWriter(fmt.Sprintf("[%s] %s: ", startOfLogger.Format(time.DateTime), level), logFile)
	for _, arg := range args {
		writeSliceToWriter(arg, logFile)
	}
	fmt.Fprintln(logFile)
	logFile.Close()
}
