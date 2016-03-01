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
    t_alarm.PrintMonitors()
}

func TestAddDeadlineAlert(t *testing.T) {
	err := t_alarm.AddDeadlineAlert("module2", "用户数据没有提交")
	if err != nil {
		t.Errorf("add deadline alert fail: %v", err)
	}
    t_alarm.PrintMonitors()
}

func TestEnableDeadlineAlert1(t *testing.T) {
	deadline := time.Now().Add(time.Second * 2)
	err := t_alarm.EnableDeadlineAlert("module2", deadline)
	if err != nil {
		t.Errorf("enable deadline alert fail: %v", err)
	}

    t_alarm.PrintMonitors()

	time.Sleep(time.Second * 3)
}

func TestEnableDeadlineAlert2(t *testing.T) {
	deadline := time.Now().Add(time.Second * 3)
	err := t_alarm.EnableDeadlineAlert("module2", deadline)
	if err != nil {
		t.Errorf("enable deadline alert fail: %v", err)
	}

    t_alarm.PrintMonitors()

	time.Sleep(time.Second * 2)
}

func TestDisableDeadlineAlert(t *testing.T) {
	err := t_alarm.DisableDeadlineAlert("module2")
	if err != nil {
		t.Errorf("disable deadline alert fail: %v", err)
	}
    t_alarm.PrintMonitors()
}

func TestRemoveDeadlineAlert(t *testing.T) {
	err := t_alarm.RemoveDeadlineAlert("module2")
	if err != nil {
		t.Errorf("remove deadline alert fail: %v", err)
	}
    t_alarm.PrintMonitors()
}

