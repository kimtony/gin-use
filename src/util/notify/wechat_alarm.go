package notify

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
	"gin-use/src/global"
	"gin-use/configs"


)

const KEY_DEV_OPS string = "dev_ops" // 报警处理人

var alarm *WcAlarm
var alarmOnce = new(sync.Once)

func GetWcAlarm() *WcAlarm {
	alarmOnce.Do(func() {
		alarm = NewWcAlarm()
	})
	return alarm
}

const MAX_UINT64_NUM = ^uint64(0)

// 报警器
type WcAlarm struct {
	addrs         map[string][]string // 报警 地址:维护人
	totalTimes    uint64              // 总报警次数
	currTimes     int                 // 今天报警次数
	lastResetTime time.Time           // 上一次重置时间
	maxNumPerDay  int                 // 每天最多报警次数
}

func NewWcAlarm() *WcAlarm {
	back := &WcAlarm{
		addrs:         make(map[string][]string),
		totalTimes:    0,
		currTimes:     0,
		lastResetTime: time.Now(),
		maxNumPerDay:  100,
	}
	return back
}

/** 报警信息
POST https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=53ea6495-b457-4ef8-9199-98afb277023c
Content-Type: application/json
{
    "msgtype": "text",
    "text": {
        "content": "这是一条警报测试@所有人"
    }
}
*/
func (a *WcAlarm) Warn(fmtStr string, args ...interface{}) {
	if time.Now().After(a.lastResetTime.Add(time.Duration(24) * time.Hour)) {
		a.Reset()
	}
	if a.currTimes >= a.maxNumPerDay {
		return
	}
	msg := ""
	if args != nil {
		msg = fmt.Sprintf(fmtStr, args...)
	} else {
		msg = fmtStr
	}
	msg = configs.MachineId() + ": " + msg

	w := &Warn{Typ: "text", Text: &WarnText{Content: msg}}
	data, err := json.Marshal(w)
	if err != nil {
		return
	}
	for addr, _ := range a.addrs {
		a.doWarning(addr, data)
		if a.totalTimes == MAX_UINT64_NUM {
			a.totalTimes = 0
		}
		a.totalTimes++
	}
	a.currTimes++
}

func (a *WcAlarm) Reset() {
	a.lastResetTime = time.Now()
	a.currTimes = 0
}

func (a *WcAlarm) SetMaxAlarm(n int) {
	a.maxNumPerDay = n
}

func (a *WcAlarm) AddAddr(addr string, ops ...string) {
	if len(ops) == 0 {
		a.addrs[addr] = []string{}
	} else {
		a.addrs[addr] = ops
	}
}

func (a *WcAlarm) doWarning(addr string, data []byte) {
	buf := bytes.NewBuffer(data)
	req, err := http.NewRequest("POST", addr, buf)
	if err != nil {
		global.Logger.Errorf("create warn request head fail, err:%v", err)
		return
	}
	req.Header.Set("Content-Type", "application/json;charset=UTF-8")
	cli := http.Client{}
	rsp, err := cli.Do(req.WithContext(context.TODO()))
	if err != nil {
		global.Logger.Errorf("post warn fail, err:%v", err)
		return
	}
	rsp.Body.Close()
}

type Warn struct {
	Typ  string    `json:"msgtype"`
	Text *WarnText `json:"text"`
}
type WarnText struct {
	Content string `json:"content"`
}
