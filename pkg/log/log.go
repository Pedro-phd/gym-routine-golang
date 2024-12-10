package log

import "go.uber.org/zap"

type CustomLog struct {
	logger *zap.SugaredLogger
}

func InitLogger() *CustomLog {
	logger, _ := zap.NewProduction()

	return &CustomLog{logger: logger.Sugar()}
}

func (c *CustomLog) Sync() {
	c.logger.Sync()
}

func (c *CustomLog) Success(msg string, args ...interface{}) {
	c.logger.Infow("✅ "+msg, args...)
}

func (c *CustomLog) Info(msg string, args ...interface{}) {
	c.logger.Infow("ℹ️ "+msg, args...)
}

func (c *CustomLog) Error(msg string, args ...interface{}) {
	c.logger.Errorw("❌ "+msg, args...)
}

func (c *CustomLog) Warn(msg string, args ...interface{}) {
	c.logger.Warnw("⚠️ "+msg, args...)
}

func (c *CustomLog) Fatal(msg string, args ...interface{}) {
	c.logger.Fatalw("🆘 "+msg, args...)
}

func (c *CustomLog) Debug(msg string, args ...interface{}) {
	c.logger.Debugw("🐞 "+msg, args...)
}
