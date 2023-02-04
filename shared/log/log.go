package log

import (
	"context"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"

	utilctx "SANDBOX-TASHA-ISSUER-SERVICE-BE/shared/util/context"
)

type (
	LoggerWithContext interface {
		Info(context.Context, ...interface{})
		Infof(context.Context, string, ...interface{})
		Debug(context.Context, ...interface{})
		Debugf(context.Context, string, ...interface{})
		Error(context.Context, ...interface{})
		Errorf(context.Context, string, ...interface{})
		Warning(context.Context, ...interface{})
		Warningf(context.Context, string, ...interface{})
		Fatal(context.Context, ...interface{})
		Fatalf(context.Context, string, ...interface{})
		Print(context.Context, ...interface{})
		Printf(context.Context, string, ...interface{})
		Instance(context.Context) interface{}
	}

	loggerWithContext struct {
		instance *logrus.Logger
	}
)

func NewLogger(logger *logrus.Logger) (LoggerWithContext, error) {
	if logger == nil {
		return nil, errors.New("logger can't be nil")
	}
	return &loggerWithContext{
		instance: logger,
	}, nil
}

func (l *loggerWithContext) appendRequestID(ctx context.Context, v ...interface{}) []interface{} {
	args := make([]interface{}, 0, len(v)+1)
	args = append(args, fmt.Sprintf("[%s] ", utilctx.MandatoryRequest(ctx).RequestID))
	return append(args, v...)
}

func (l *loggerWithContext) appendRequestIDtoMsg(ctx context.Context, s string) string {
	return fmt.Sprintf("[%s] %s", utilctx.MandatoryRequest(ctx).RequestID, s)
}

func (l *loggerWithContext) Info(ctx context.Context, v ...interface{}) {
	l.instance.Info(l.appendRequestID(ctx, v...)...)
}

func (l *loggerWithContext) Infof(ctx context.Context, s string, v ...interface{}) {
	l.instance.Infof(l.appendRequestIDtoMsg(ctx, s), v...)
}

func (l *loggerWithContext) Debug(ctx context.Context, v ...interface{}) {
	l.instance.Debug(l.appendRequestID(ctx, v...)...)
}

func (l *loggerWithContext) Debugf(ctx context.Context, s string, v ...interface{}) {
	l.instance.Debugf(l.appendRequestIDtoMsg(ctx, s), v...)
}

func (l *loggerWithContext) Error(ctx context.Context, v ...interface{}) {
	l.instance.Error(l.appendRequestID(ctx, v...)...)
}

func (l *loggerWithContext) Errorf(ctx context.Context, s string, v ...interface{}) {
	l.instance.Errorf(l.appendRequestIDtoMsg(ctx, s), v...)
}

func (l *loggerWithContext) Warning(ctx context.Context, v ...interface{}) {
	l.instance.Warning(l.appendRequestID(ctx, v...)...)
}

func (l *loggerWithContext) Warningf(ctx context.Context, s string, v ...interface{}) {
	l.instance.Warningf(l.appendRequestIDtoMsg(ctx, s), v...)
}

func (l *loggerWithContext) Fatal(ctx context.Context, v ...interface{}) {
	l.instance.Fatal(l.appendRequestID(ctx, v...)...)
}

func (l *loggerWithContext) Fatalf(ctx context.Context, s string, v ...interface{}) {
	l.instance.Fatalf(l.appendRequestIDtoMsg(ctx, s), v...)
}

func (l *loggerWithContext) Print(ctx context.Context, v ...interface{}) {
	l.instance.Print(l.appendRequestID(ctx, v...)...)
}

func (l *loggerWithContext) Printf(ctx context.Context, s string, v ...interface{}) {
	l.instance.Printf(l.appendRequestIDtoMsg(ctx, s), v...)
}

func (l *loggerWithContext) Instance(ctx context.Context) interface{} {
	return l.instance
}
