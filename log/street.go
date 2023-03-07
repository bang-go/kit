package log

func (l *Logger) Debug(msg string, fields ...Field) {
	l.Debug(msg, fields...)
}

func (l *Logger) Info(msg string, fields ...Field) {
	l.Info(msg, fields...)
}
func (l *Logger) Warn(msg string, fields ...Field) {
	l.Warn(msg, fields...)
}
func (l *Logger) Error(msg string, fields ...Field) {
	l.Error(msg, fields...)
}
func (l *Logger) DPanic(msg string, fields ...Field) {
	l.DPanic(msg, fields...)
}
func (l *Logger) Panic(msg string, fields ...Field) {
	l.Panic(msg, fields...)
}
func (l *Logger) Fatal(msg string, fields ...Field) {
	l.Fatal(msg, fields...)
}
func (l *Logger) Sync() error {
	return l.Sync()
}
