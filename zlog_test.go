package zlog

import (
	"testing"
)

func TestZLog(*testing.T) {
	logger.Debug("hello", Field("Sex", "Girl"))
}
