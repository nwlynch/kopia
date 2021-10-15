// Package zaplogutil provides reusable utilities for working with ZAP logger.
package zaplogutil

import (
	"time"

	"go.uber.org/zap/zapcore"

	"github.com/kopia/kopia/internal/clock"
)

// PreciseTimeEncoder encodes the time as RFC3389 with 6 digits of sub-second precision.
var PreciseTimeEncoder = zapcore.TimeEncoderOfLayout("2006-01-02T15:04:05.000000Z07:00")

type theClock struct{}

func (c theClock) Now() time.Time                         { return clock.Now() }
func (c theClock) NewTicker(d time.Duration) *time.Ticker { return time.NewTicker(d) }

// Clock isn aimplementation of zapcore.Clock that uses clock.Now().
var Clock zapcore.Clock = theClock{}

// TimezoneAdjust returns zapcore.TimeEncoder that adjusts the time to either UTC or local time before logging.
func TimezoneAdjust(inner zapcore.TimeEncoder, isLocal bool) zapcore.TimeEncoder {
	if isLocal {
		return func(t time.Time, pae zapcore.PrimitiveArrayEncoder) {
			inner(t.Local(), pae)
		}
	}

	return func(t time.Time, pae zapcore.PrimitiveArrayEncoder) {
		inner(t.UTC(), pae)
	}
}