package zlog

import (
	"testing"
)

func TestDebug(*testing.T) {
	Debug("Debug", Field("Test", "test Debug"))
}

func TestInfo(*testing.T) {
	Info("Info", Field("Test", "test Info"))
}

func TestWarn(*testing.T) {
	Warn("Warn", Field("Test", "test Warn"))
}

func TestError(*testing.T) {
	Error("Error", Field("Test", "test Error"))
}

func TestDPanic(*testing.T) {
	DPanic("DPanic", Field("Test", "test DPanic"))
}

func TestFatal(*testing.T) {
	Fatal("Fatal", Field("Test", "test Fatal"))
}

func TestPanic(*testing.T) {
	Panic("Panic", Field("Test", "test Panic"))
}
