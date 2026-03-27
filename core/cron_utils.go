package core

// Implement the cron logger interface
type CronUtils struct {
	Logger Logger
}

func NewCronUtils(l Logger) *CronUtils {
	return &CronUtils{Logger: l}
}

func (c *CronUtils) Info(msg string, keysAndValues ...any) {
	c.Logger.Debug("cron update", append(keysAndValues, "cron", msg)...)
}

func (c *CronUtils) Error(err error, msg string, keysAndValues ...interface{}) {
	c.Logger.Error("cron error", append(keysAndValues, "cron", msg, "error", err)...)
}
