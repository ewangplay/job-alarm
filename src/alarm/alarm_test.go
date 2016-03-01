package alarm

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var t_alarm *Alarm

func TestMain(m *testing.M) {
	f := func(module, alarm_msg string) error {
		fmt.Printf("%v: %v\n", module, alarm_msg)
		return nil
	}

	t_alarm = NewAlarm(f)

	os.Exit(m.Run())
}

func TestAlert(t *testing.T) {
	t_alarm.Alert("module1", "保存数据库失败")
}

func TestDeadlineAlert(t *testing.T) {
	deadline := time.Now().Add(time.Second * 2)
	err := t_alarm.SetDeadlineAlert("module2", "用户数据没有提交", deadline)
	if err != nil {
		t.Errorf("set deadline alert fail: %v", err)
	}

	time.Sleep(time.Second * 3)
}

func TestCancelDeadlineAlert(t *testing.T) {
	deadline := time.Now().Add(time.Second * 2)
	err := t_alarm.SetDeadlineAlert("module3", "用户数据没有提交", deadline)
	if err != nil {
		t.Errorf("set deadline alert fail: %v", err)
	}

	time.Sleep(time.Second * 1)

	err = t_alarm.CancelDeadlineAlert("module3")
	if err != nil {
		t.Errorf("cancel deadline alert fail: %v", err)
	}
}
