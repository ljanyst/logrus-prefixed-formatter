
Logrus Prefixed Formatter
=========================

This is a fork of the [Logrus Prefixed Formatter][1] which itself is a
[Logrus][2] formatter mainly based on original `logrus.TextFormatter` but with
a slightly modified colored output and support for custom color themes. We also
enable prefixing of log messages using either a field in the log entry called
"prefix" or an object called "logging-context" extracted from a context.

Usage
-----

Here is how it should be used:

```go
package main

import (
	log "github.com/sirupsen/logrus"
	prefixed "github.com/ljanyst/logrus-prefixed-formatter"
)

func init() {
	log.SetFormatter(&prefixed.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
		ForceFormatting: true,
	})
	log.SetLevel(log.DebugLevel)
}

func main() {
	log.WithFields(logrus.Fields{
		"prefix": "main",
		"animal": "walrus",
		"number": 8,
	}).Debug("Started observing beach")

	log.WithFields(logrus.Fields{
		"prefix":      "sensor",
		"temperature": -4,
	}).Info("Temperature changes")
}
```

And here's how it's used with contexts:

```go
package main

import (
	"context"
	"fmt"

	prefixed "github.com/ljanyst/logrus-prefixed-formatter"
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
	formatter := new(prefixed.TextFormatter)
	formatter.MinPrefixWidth = 25
	logrus.SetFormatter(formatter)
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	logrus.Info("Before context")
	ctx := &TestContext{}
	log := prefixed.Log(context.WithValue(context.Background(), "logging-context", ctx))
	log.Debugf("Started observing beach")
	ctx.Type = "foo"
	log.Debugf("A group of walrus emerges from the ocean")
	ctx.Foo = "bar"
	log.Warnf("The group's number increased tremendously!")
	ctx.Bar = 42
	log.Infof("Temperature changes")
}
```

API
---

`prefixed.TextFormatter` exposes the following fields and methods.

### Fields

 * `ForceColors bool` — set to true to bypass checking for a TTY before
   outputting colors.
 * `DisableColors bool` — force disabling colors. For a TTY colors are enabled
   by default.
 * `DisableUppercase bool` — set to true to turn off the conversion of the log
   level names to uppercase.
 * `ForceFormatting bool` — force formatted layout, even for non-TTY output.
 * `DisableTimestamp bool` — disable timestamp logging. Useful when output
   is redirected to logging system that already adds timestamps.
 * `FullTimestamp bool` — enable logging the full timestamp when a TTY is
   attached instead of just the time passed since beginning of execution.
 * `TimestampFormat string` — timestamp format to use for display when a full
   timestamp is printed.
 * `DisableSorting bool` — the fields are sorted by default for a consistent
   output. For applications that log extremely frequently and don't use the JSON
   formatter this may not be desired.
 * `QuoteEmptyFields bool` — wrap empty fields in quotes if true.
 * `QuoteCharacter string` — can be set to the override the default quoting
   character `"` with something else. For example: `'`, or `` ` ``.
 * `SpacePadding int` — pad msg field with spaces on the right for display.
   The value for this parameter will be the size of padding. Its default value
   is zero, which means no padding will be applied.
 * `MinPrefixWidth int` - the prefixes will be padded to be the length specified
   by this value. If the value is zero (the default), no padding will be
   applied.

### Methods

`SetColorScheme(colorScheme *prefixed.ColorScheme)`

Sets an alternative color scheme for colored output. `prefixed.ColorScheme`
struct supports the following fields:

 * `InfoLevelStyle string` — info level style.
 * `WarnLevelStyle string` — warn level style.
 * `ErrorLevelStyle string` — error style.
 * `FatalLevelStyle string` — fatal level style.
 * `PanicLevelStyle string` — panic level style.
 * `DebugLevelStyle string` — debug level style.
 * `PrefixStyle string` — prefix style.
 * `TimestampStyle string` — timestamp style.

Color styles should be specified using [mgutz/ansi][3] style syntax. For
example, here is the default theme:

```go
InfoLevelStyle:  "green",
WarnLevelStyle:  "yellow",
ErrorLevelStyle: "red",
FatalLevelStyle: "red",
PanicLevelStyle: "red",
DebugLevelStyle: "blue",
PrefixStyle:     "cyan",
TimestampStyle:  "black+h"
```

It's not necessary to specify all colors when changing color scheme if you want
to change just specific ones:

```go
formatter.SetColorScheme(&prefixed.ColorScheme{
    PrefixStyle:    "blue+b",
    TimestampStyle: "white+h",
})
```

`prefixed.Log` creates a logrus log wrapper that calls the message functions
on log entries created with context:

```go
func (l *LogWrap) Tracef(format string, args ...interface{}) {
	l.log.WithContext(l.ctx).Tracef(format, args...)
}
```

This removes the need for manually setting the context for every log entry
in a given scope.

License
-------

MIT

[1]: https://github.com/x-cray/logrus-prefixed-formatter
[2]: https://github.com/sirupsen/logrus
[3]: https://github.com/mgutz/ansi#style-format
