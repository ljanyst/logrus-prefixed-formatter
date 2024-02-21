package main

import (
	"context"
	"fmt"

	"github.com/ljanyst/pre"
	"github.com/sirupsen/logrus"
)

type TestContext struct {
	Type string
	Foo  string
	Bar  int
}

func (t *TestContext) String() string {
	return fmt.Sprintf("%-5s f: %-10s b: %3d", t.Type, t.Foo, t.Bar)
}

func init() {
	formatter := new(pre.TextFormatter)
	formatter.MinPrefixWidth = 25
	logrus.SetFormatter(formatter)
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	logrus.Info("Before context")
	ctx := &TestContext{}
	log := pre.Log(context.WithValue(context.Background(), "logging-context", ctx))
	log.Debugf("Started observing beach")
	ctx.Type = "foo"
	log.Debugf("A group of walrus emerges from the ocean")
	ctx.Foo = "bar"
	log.Warnf("The group's number increased tremendously!")
	ctx.Bar = 42
	log.Infof("Temperature changes")
}
