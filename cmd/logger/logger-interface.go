package logger

import (
	"github.com/fatih/color"
	"github.com/guumaster/logsymbols"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

// Event stores messages to log later, from our standard interface
type Event struct {
	id      int
	message string
}

// StandardLogger enforces specific log message formats
type StandardLogger struct {
	*logrus.Logger
}

// NewLogger initializes the standard logger
func NewLogger() *StandardLogger {
	var baseLogger = logrus.New()

	var standardLogger = &StandardLogger{baseLogger}

	//standardLogger.Formatter = &logrus.StdLogger{}
	//standardLogger.Formatter = &logrus.JSONFormatter{}
	standardLogger.Formatter = &logrus.TextFormatter{
		DisableColors:   false,
		TimestampFormat: "01-02-01-2006 15:04:05.9999Z07:00",
		//TimestampFormat : "01-02-01-2006 15:04:05.999999999Z07:00",
		FullTimestamp: true,
		//ForceFormatting: true,
	}
	return standardLogger
}

// Declare variables to store log messages as new Events
var (
	invalidArgMessage      = Event{1, "Invalid arg: %s"}
	invalidArgValueMessage = Event{2, "Invalid value for argument: %s: %v"}
	missingArgMessage      = Event{3, "Missing arg: %s"}
)

// InvalidArg is a standard error message
func (l *StandardLogger) InvalidArg(argumentName string) {
	l.Errorf(invalidArgMessage.message, argumentName)
}

// InvalidArgValue is a standard error message
func (l *StandardLogger) InvalidArgValue(argumentName string, argumentValue string) {
	l.Errorf(invalidArgValueMessage.message, argumentName, argumentValue)
}

// MissingArg is a standard error message
func (l *StandardLogger) MissingArg(argumentName string) {
	l.Errorf(missingArgMessage.message, argumentName)
}

func (l *StandardLogger) PrettyHttpRequestServer(r *http.Request) {
	success := logsymbols.Success

	l.Info("<-----------------------" + success + " Http Request " + success + "---------------------------------<")
	l.Info("Call Method:", r.Method, "| Path:", r.RequestURI)
	l.Info("Header :", r.Header)
	l.Info("Body :", r.Body)
	l.Info("<-----------------------" + success + " Http Request " + success + "---------------------------------<")

}
func (l *StandardLogger) PrettyHttpRequestClient(resp *http.Response) {
	ok := logsymbols.Ok

	l.Info("<-----------------------" + ok + " Client: Http Response " + ok + "---------------------------------<")
	l.Info("Call Method:", resp.Request.Method, "| Path:", resp.Request.RequestURI)
	l.Info("Header :", resp.Header)
	l.Info("Body :", resp.Body)
	l.Info("<-----------------------" + ok + " Client: Http Response " + ok + "---------------------------------<")

}
func (l *StandardLogger) PrintTracingTransport(duration time.Duration, connDuration time.Duration, reqDuration time.Duration) {
	ok := logsymbols.Success
	yellow := color.New(color.FgYellow).SprintFunc()
	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()

	l.Info(">------> "+ok+" Request statistic ( Duration: ", red(duration), " ConnectionDuration : ", green(connDuration), " RequestDuration : ", yellow(reqDuration), ") "+ok+" >------>")

}
