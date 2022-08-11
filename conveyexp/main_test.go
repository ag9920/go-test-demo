package conveyexp

import (
	"os"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMain(m *testing.M) {
	// convey在TestMain场景下的入口
	SuppressConsoleStatistics()
	result := m.Run()
	// convey在TestMain场景下的结果打印
	PrintConsoleStatistics()
	os.Exit(result)
}
