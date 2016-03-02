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

func TestSetDeadlineAlert(t *testing.T) {
	deadline := time.Now().Add(time.Second * 3)
	err := t_alarm.SetDeadlineAlert("module3", "订单数据没有提交", deadline)
	if err != nil {
		t.Errorf("set deadline alert fail: %v", err)
	}

	time.Sleep(time.Second * 2)
	t_alarm.PrintMonitors()
}

func TestUnsetDeadlineAlert(t *testing.T) {
	err := t_alarm.UnsetDeadlineAlert("module3")
	if err != nil {
		t.Errorf("unset deadline alert fail: %v", err)
	}

	t_alarm.PrintMonitors()
}

func TestAddTimerAlert(t *testing.T) {
    timer := time.Date(2016, time.March, 2, 11, 30, 0, 0, time.Local)
	err := t_alarm.AddTimerAlert("module4", "活动数据没有提交", timer)
	if err != nil {
		t.Errorf("add timer alert fail: %v", err)
	}
	t_alarm.PrintMonitors()
}

func TestEnableTimerAlert(t *testing.T) {
	err := t_alarm.EnableTimerAlert("module4")
	if err != nil {
		t.Errorf("enable timer alert fail: %v", err)
	}
	t_alarm.PrintMonitors()
}

func TestDisableTimerAlert(t *testing.T) {
	err := t_alarm.DisableTimerAlert("module4")
	if err != nil {
		t.Errorf("disable timer alert fail: %v", err)
	}
	t_alarm.PrintMonitors()
}

func TestRemoveTimerAlert(t *testing.T) {
	err := t_alarm.RemoveTimerAlert("module4")
	if err != nil {
		t.Errorf("remove timer alert fail: %v", err)
	}
	t_alarm.PrintMonitors()
}

func TestSetTimerAlert(t *testing.T) {
    timer := time.Date(2016, time.March, 2, 11, 30, 0, 0, time.Local)
	err := t_alarm.SetTimerAlert("module5", "渠道数据没有提交", timer)
	if err != nil {
		t.Errorf("set timer alert fail: %v", err)
	}
	t_alarm.PrintMonitors()

}

func TestUnsetTimerAlert(t *testing.T) {
	err := t_alarm.UnsetTimerAlert("module5")
	if err != nil {
		t.Errorf("set timer alert fail: %v", err)
	}
	t_alarm.PrintMonitors()
}
