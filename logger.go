// Copyright (c) 2024 The illium developers
// Use of this source code is governed by an MIT
// license that can be found in the LICENSE file.

package logger

import (
	"github.com/pterm/pterm"
)

// Logger is a custom logger that wraps multiple
// pterm.Loggers. This enables printing to multiple
// locations such as the terminal and to file, using
// different formats.
type Logger struct {
	Loggers    []*pterm.Logger
	Level      pterm.LogLevel
	ShowCaller bool
}

// WithCustomLogger appends the logger to the list of loggers.
// A log call propagates to all the stored loggers.
func (l Logger) WithCustomLogger(logger *pterm.Logger) *Logger {
	l.Loggers = append(l.Loggers, logger)
	return &l
}

// WithLevel sets the log level of the logger.
func (l Logger) WithLevel(level pterm.LogLevel) *Logger {
	for _, logger := range l.Loggers {
		logger.WithLevel(level)
	}
	return &l
}

// WithCaller enables or disables the caller.
func (l Logger) WithCaller(b ...bool) *Logger {
	for _, logger := range l.Loggers {
		logger.WithCaller(b...)
	}
	return &l
}

// Args converts any arguments to a slice of pterm.LoggerArgument.
func (l Logger) Args(args ...any) []pterm.LoggerArgument {
	var loggerArgs []pterm.LoggerArgument

	// args are in the format of: key, value, key, value, key, value, ...
	for i := 0; i < len(args); i += 2 {
		key := pterm.Sprint(args[i])
		value := args[i+1]

		loggerArgs = append(loggerArgs, pterm.LoggerArgument{
			Key:   key,
			Value: value,
		})
	}

	return loggerArgs
}

// ArgsFromMap converts a map to a slice of pterm.LoggerArgument.
func (l Logger) ArgsFromMap(m map[string]any) []pterm.LoggerArgument {
	var loggerArgs []pterm.LoggerArgument

	for k, v := range m {
		loggerArgs = append(loggerArgs, pterm.LoggerArgument{
			Key:   k,
			Value: v,
		})
	}

	return loggerArgs
}

// Trace prints a trace log.
func (l Logger) Trace(msg string, args ...[]pterm.LoggerArgument) {
	for _, logger := range l.Loggers {
		logger.Trace(msg, args...)
	}
}

// Debug prints a debug log.
func (l Logger) Debug(msg string, args ...[]pterm.LoggerArgument) {
	for _, logger := range l.Loggers {
		logger.Debug(msg, args...)
	}
}

// Info prints an info log.
func (l Logger) Info(msg string, args ...[]pterm.LoggerArgument) {
	for _, logger := range l.Loggers {
		logger.Info(msg, args...)
	}
}

// Warn prints a warning log.
func (l Logger) Warn(msg string, args ...[]pterm.LoggerArgument) {
	for _, logger := range l.Loggers {
		logger.Warn(msg, args...)
	}
}

// Error prints an error log.
func (l Logger) Error(msg string, args ...[]pterm.LoggerArgument) {
	for _, logger := range l.Loggers {
		logger.Error(msg, args...)
	}
}

// Fatal prints a fatal log and exits the program.
func (l Logger) Fatal(msg string, args ...[]pterm.LoggerArgument) {
	if len(l.Loggers) > 1 {
		for _, logger := range l.Loggers[1:] {
			logger.Error(msg, args...)
		}
	}
	if len(l.Loggers) > 0 {
		l.Loggers[0].Fatal(msg, args...)
	}
}
