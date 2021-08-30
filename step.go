package timetracker

import (
	"fmt"
	"log"
	"sync/atomic"
	"time"
)

type StepInfo struct {
	//步骤注释/描述信息
	cmt string
	//开始时间
	start int64
	//结束时间
	end int64
	//花费时间
	time int64
	//步骤在整个步骤里的等级
	level int32
	//子步骤
	children []*StepInfo
	//父步骤
	parent *StepInfo
}

type TimerTrack struct {
	//步骤信息根
	root *StepInfo
	//当前信息指针
	current *StepInfo
	//步骤数
	stepNum int32
	//是否以毫秒输出
	isMs bool
	//是否关闭
	close bool
}

var defaultTimerTrack *TimerTrack = &TimerTrack{
	root: &StepInfo{},
	isMs: true,
}

// NewTimeTracker 创建一个timeTracker
func NewTimeTracker(cmt string) *TimerTrack {
	step := &StepInfo{
		cmt:   cmt,
		start: time.Now().UnixNano(),
		level: 1,
	}
	t := &TimerTrack{
		stepNum: 0,
		root:    step,
		current: step,
		isMs:    true,
	}
	return t
}
func (t *TimerTrack) Close() {
	root := t.root
	root.end = time.Now().UnixNano()
	root.time = (root.end - root.start) / 1e6
	t.close = true
}

func (t *TimerTrack) StepStart(cmt string) *TimerTrack {
	if t.close {
		log.Fatalf("the timeTracker:%s has been closed", t.root.cmt)
	}
	step := &StepInfo{
		cmt:    cmt,
		start:  time.Now().UnixNano(),
		parent: t.current,
		level:  t.current.level + 1,
	}

	//将新开的这步骤 设置为父步骤的子步骤
	t.current.children = append(t.current.children, step)

	//并且将 当前步骤设置为 当前步骤
	t.current = step

	//添加总步骤数目
	atomic.AddInt32(&t.stepNum, 1)
	//支持链式编程
	return t
}
func (t *TimerTrack) StepEnd() *TimerTrack {
	current := t.current
	current.end = time.Now().UnixNano()
	current.time = (current.end - current.start) / 1e6

	//将当前步骤 上移
	t.current = current.parent
	//支持链式编程
	return t
}

func (t *TimerTrack) PrintBeautiful() {

	//formatter :="2006-01-02 15:04:05.999999999 -0700 MST";
	//formatter := "2006-01-02 15:04:05.999 -0700 MST"
	//fmt.Println(formatter)

	///从根元素开始
	print(t.root)

}
func print(info *StepInfo) {

	prefix := prefix(info.level)
	fmt.Printf("%s %s %s\n", prefix, info.cmt, human(info.time))
	for _, child := range info.children {
		print(child)
	}
}
func human(t int64) string {
	//将时间转换为 人类可识别的字眼
	//这里我们转换最大值为 小时
	ms := t % 1000
	//留下的整数为秒
	s := t / 1000
	//留下的整数为分
	m := s / 60
	//留下的整数为时
	h := m / 60

	s = s % 60
	m = m % 60

	return fmt.Sprintf("%dh %dm %ds %dms ", h, m, s, ms)
}
func prefix(level int32) string {
	result := []byte{}
	for i := 0; i < int(level+1); i++ {
		result = append(result, byte('-'))
		result = append(result, byte('-'))
		result = append(result, byte('-'))
	}
	return string(result)
}
