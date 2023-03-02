package log

func Debug(msg string, fields ...Field) {
	l.Debug(msg, fields...)
}
func Info(msg string, fields ...Field) {
	l.Info(msg, fields...)
}
func Warn(msg string, fields ...Field) {
	l.Warn(msg, fields...)
}
func Error(msg string, fields ...Field) {
	l.Error(msg, fields...)
}
func DPanic(msg string, fields ...Field) {
	l.DPanic(msg, fields...)
}
func Panic(msg string, fields ...Field) {
	l.Panic(msg, fields...)
}
func Fatal(msg string, fields ...Field) {
	l.Fatal(msg, fields...)
}
func Sync() error {
	return l.Sync()
}
