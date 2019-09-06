package logger

import (
	"github.com/sirupsen/logrus"
)

var logger *StandardLogger

func init() {
	logger = NewLogger()
}

type Event struct {
	id      int
	message string
}

type StandardLogger struct {
	*logrus.Logger
}

func NewLogger() *StandardLogger {
	var baseLogger = logrus.New()

	var standardLogger = &StandardLogger{baseLogger}

	standardLogger.Formatter = &logrus.TextFormatter{
		FullTimestamp: true,
	}
	// We could transform the errors into a JSON format, for external log SaaS tools
	// standardLogger.Formatter = &logrus.JSONFormatter{
	// 	PrettyPrint: true,
	// }

	return standardLogger
}

var (
	invalidArgMessage      = Event{1, "Invalid arg: %s"}
	invalidArgValueMessage = Event{2, "Invalid value for argument: %s: %v"}
	missingArgMessage      = Event{3, "Missing arg: %s"}
)

func Errorfn(fn string, err error) {
	logger.Errorf("[%s]: %v", fn, err)
}

func InvalidArg(argumentName string) {
	logger.Errorf(invalidArgMessage.message, argumentName)
}

func InvalidArgValue(argumentName string, argumentValue string) {
	logger.Errorf(invalidArgValueMessage.message, argumentName, argumentValue)
}

func MissingArg(argumentName string) {
	logger.Errorf(missingArgMessage.message, argumentName)
}

func Info(args ...interface{}) {
	logger.Infoln(args...)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Warn(args ...interface{}) {
	logger.Warnln(args...)
}

func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

func Panic(args ...interface{}) {
	logger.Panicln(args...)
}

func Panicf(format string, args ...interface{}) {
	logger.Panicf(format, args...)
}

func Error(args ...interface{}) {
	logger.Errorln(args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

func Fatal(args ...interface{}) {
	logger.Fatalln(args...)
}

func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}