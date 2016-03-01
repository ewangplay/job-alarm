package alarm

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var t_alarm *Alarm

func TestMain(m *testing.M) {
	f := func(desc string) error {
		fmt.Println(desc)
		return nil
	}

	t_alarm = NewAlarm(f)

	os.Exit(m.Run())
}

func TestAlert(t *testing.T) {
	t_alarm.Alert("保存数据库失败")
}

func TestMonitor(t *testing.T) {
	deadline := time.Now().Add(time.Second * 2)
	err := t_alarm.SetDeadlineAlert("test", deadline, "用户数据没有提交")
	if err != nil {
		t.Errorf("set deadline alert fail: %v", err)
	}

	time.Sleep(time.Second * 4)

	err = t_alarm.CancelDeadlineAlert("test")
	if err != nil {
		t.Errorf("cancel deadline alert fail: %v", err)
	}
}
