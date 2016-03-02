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

func TestAddDeadlineAlert(t *testing.T) {
	duration := time.Second * 2
	err := t_alarm.AddDeadlineAlert("module2", "用户数据没有提交", duration)
	if err != nil {
		t.Errorf("add deadline alert fail: %v", err)
	}

    status := t_alarm.GetAlertStatus("module2")
    if status != ALERT_STATUS_INIT {
		t.Errorf("module2 alert's status should be INIT")
    }
}

func TestEnableDeadlineAlert(t *testing.T) {
	err := t_alarm.EnableAlert("module2")
	if err != nil {
		t.Errorf("enable deadline alert fail: %v", err)
	}

    status := t_alarm.GetAlertStatus("module2")
    if status != ALERT_STATUS_ENABLE {
		t.Errorf("module2 alert's status should be ENABLE")
    }

	time.Sleep(time.Second * 3)
}

func TestDisableDeadlineAlert(t *testing.T) {
	err := t_alarm.DisableAlert("module2")
	if err != nil {
		t.Errorf("disable deadline alert fail: %v", err)
	}

    status := t_alarm.GetAlertStatus("module2")
    if status != ALERT_STATUS_DISABLE {
		t.Errorf("module2 alert's status should be DISABLE")
    }
}

func TestRemoveDeadlineAlert(t *testing.T) {
	err := t_alarm.RemoveAlert("module2")
	if err != nil {
		t.Errorf("remove deadline alert fail: %v", err)
	}

    status := t_alarm.GetAlertStatus("module2")
    if status != ALERT_STATUS_NONE {
		t.Errorf("module2 alert's status should be NONE")
    }
}

func TestSetDeadlineAlert(t *testing.T) {
	duration := time.Second * 3
	err := t_alarm.SetDeadlineAlert("module3", "订单数据没有提交", duration)
	if err != nil {
		t.Errorf("set deadline alert fail: %v", err)
	}

    status := t_alarm.GetAlertStatus("module3")
    if status != ALERT_STATUS_ENABLE {
		t.Errorf("module2 alert's status should be ENABLE")
    }

	time.Sleep(time.Second * 2)
}

func TestUnsetDeadlineAlert(t *testing.T) {
	err := t_alarm.UnsetAlert("module3")
	if err != nil {
		t.Errorf("unset deadline alert fail: %v", err)
	}

    status := t_alarm.GetAlertStatus("module3")
    if status != ALERT_STATUS_NONE {
		t.Errorf("module2 alert's status should be NONE")
    }
}

func TestAddTimerAlert(t *testing.T) {
    timer := time.Date(2016, time.March, 2, 11, 30, 0, 0, time.Local)
	err := t_alarm.AddTimerAlert("module4", "活动数据没有提交", timer)
	if err != nil {
		t.Errorf("add timer alert fail: %v", err)
	}

    status := t_alarm.GetAlertStatus("module4")
    if status != ALERT_STATUS_INIT {
		t.Errorf("module2 alert's status should be INIT")
    }
}

func TestEnableTimerAlert(t *testing.T) {
	err := t_alarm.EnableAlert("module4")
	if err != nil {
		t.Errorf("enable timer alert fail: %v", err)
	}

    status := t_alarm.GetAlertStatus("module4")
    if status != ALERT_STATUS_ENABLE {
		t.Errorf("module2 alert's status should be ENABLE")
    }
}

func TestDisableTimerAlert(t *testing.T) {
	err := t_alarm.DisableAlert("module4")
	if err != nil {
		t.Errorf("disable timer alert fail: %v", err)
	}

    status := t_alarm.GetAlertStatus("module4")
    if status != ALERT_STATUS_DISABLE {
		t.Errorf("module2 alert's status should be DISABLE")
    }
}

func TestRemoveTimerAlert(t *testing.T) {
	err := t_alarm.RemoveAlert("module4")
	if err != nil {
		t.Errorf("remove timer alert fail: %v", err)
	}

    status := t_alarm.GetAlertStatus("module4")
    if status != ALERT_STATUS_NONE {
		t.Errorf("module2 alert's status should be NONE")
    }
}

func TestSetTimerAlert(t *testing.T) {
    timer := time.Date(2016, time.March, 2, 11, 30, 0, 0, time.Local)
	err := t_alarm.SetTimerAlert("module5", "渠道数据没有提交", timer)
	if err != nil {
		t.Errorf("set timer alert fail: %v", err)
	}

    status := t_alarm.GetAlertStatus("module5")
    if status != ALERT_STATUS_ENABLE {
		t.Errorf("module2 alert's status should be ENABLE")
    }
}

func TestUnsetTimerAlert(t *testing.T) {
	err := t_alarm.UnsetAlert("module5")
	if err != nil {
		t.Errorf("set timer alert fail: %v", err)
	}

    status := t_alarm.GetAlertStatus("module5")
    if status != ALERT_STATUS_NONE {
		t.Errorf("module2 alert's status should be NONE")
    }
}
