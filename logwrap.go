package pre

import (
	"context"

	"github.com/sirupsen/logrus"
)

type LogWrap struct {
	log *logrus.Logger
	ctx context.Context
}

func Log(ctx context.Context) *LogWrap {
	return &LogWrap{logrus.StandardLogger(), ctx}
}

func LogCustom(log *logrus.Logger, ctx context.Context) *LogWrap {
	return &LogWrap{log, ctx}
}

func (l *LogWrap) Tracef(format string, args ...interface{}) {
	l.log.WithContext(l.ctx).Tracef(format, args...)
}

func (l *LogWrap) Debugf(format string, args ...interface{}) {
	l.log.WithContext(l.ctx).Debugf(format, args...)
}

func (l *LogWrap) Printf(format string, args ...interface{}) {
	l.log.WithContext(l.ctx).Printf(format, args...)
}

func (l *LogWrap) Infof(format string, args ...interface{}) {
	l.log.WithContext(l.ctx).Infof(format, args...)
}

func (l *LogWrap) Warnf(format string, args ...interface{}) {
	l.log.WithContext(l.ctx).Warnf(format, args...)
}

func (l *LogWrap) Warningf(format string, args ...interface{}) {
	l.log.WithContext(l.ctx).Warningf(format, args...)
}

func (l *LogWrap) Errorf(format string, args ...interface{}) {
	l.log.WithContext(l.ctx).Errorf(format, args...)
}

func (l *LogWrap) Panicf(format string, args ...interface{}) {
	l.log.WithContext(l.ctx).Panicf(format, args...)
}

func (l *LogWrap) Fatalf(format string, args ...interface{}) {
	l.log.WithContext(l.ctx).Fatalf(format, args...)
}
