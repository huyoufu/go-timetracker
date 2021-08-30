package timetracker

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestPrefix(t *testing.T) {
	s := prefix(0)
	fmt.Println(s)
}
func TestHuman(t *testing.T) {
	var time int64 = 1*3600*1000 + 5*60*1000 + 45*1000 + 110
	human(time)
	fmt.Println(human(time))
}

func TestBusiness(t *testing.T) {

	tracker := NewTimeTracker("查询商品信息")
	tracker.StepStart("第一步查询商品信息")
	time.Sleep(time.Duration(int64(rand.Intn(100) * 1000 * 10000)))
	tracker.StepEnd()
	tracker.StepStart("第二步查询商品关联的用户信息")
	time.Sleep(time.Duration(int64(rand.Intn(100) * 1000 * 10000)))
	tracker.StepEnd()
	tracker.StepStart("第三步查询商品关联的产地信息")
	tracker.StepStart("第三步查询商品关联的产地信息之省份信息")
	time.Sleep(time.Duration(int64(rand.Intn(100) * 1000 * 10000)))
	tracker.StepEnd()
	tracker.StepStart("第三步查询商品关联的产地信息之城市信息")
	time.Sleep(time.Duration(int64(rand.Intn(100) * 1000 * 10000)))
	tracker.StepEnd()
	tracker.StepEnd()
	tracker.Close()

	tracker.PrintBeautiful()
}
