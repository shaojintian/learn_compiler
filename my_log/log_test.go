package my_log

import (
	"errors"
	"testing"
)

func TestLog(t *testing.T) {
	err := errors.New("阿瓦达无多")
	LogPrint(err)
	LogFatal(err)
}
