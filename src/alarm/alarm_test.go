package alarm

import (
	"fmt"
	"os"
	"testing"
	"time"
)

var t_cb func(string,string) error
var t_alarm *Alarm

func TestMain(m *testing.M) {
	t_cb = func(module, alarm_msg string) error {
		fmt.Printf("%v: %v\n", module, alarm_msg)
		return nil
	}

	t_alarm = NewAlarm(t_cb)

	os.Exit(m.Run())
}

func TestAlert(t *testing.T) {
	t_alarm.Alert("module1", "保存数据库失败")
}

func TestDeadlineAlert1(t *testing.T) {
    var err error 
    var status int

    //Add alert
	duration := time.Second * 2
    alert := NewDeadlineAlert("module2", "用户数据没有提交", duration, t_cb)
	err = t_alarm.AddAlert("module2", alert)
	if err != nil {
		t.Errorf("add deadline alert fail: %v", err)
	}

    status = t_alarm.GetAlertStatus("module2")
    if status != ALERT_STATUS_INIT {
		t.Errorf("module2 alert's status should be INIT")
    }

    //Enable alert
	err = t_alarm.EnableAlert("module2")
	if err != nil {
		t.Errorf("enable deadline alert fail: %v", err)
	}

    status = t_alarm.GetAlertStatus("module2")
    if status != ALERT_STATUS_ENABLE {
		t.Errorf("module2 alert's status should be ENABLE")
    }

    //Sleep 
	time.Sleep(time.Second * 3)

    //Disable alert
	err = t_alarm.DisableAlert("module2")
	if err != nil {
		t.Errorf("disable deadline alert fail: %v", err)
	}

    status = t_alarm.GetAlertStatus("module2")
    if status != ALERT_STATUS_DISABLE {
		t.Errorf("module2 alert's status should be DISABLE")
    }

    //Remove alert
	err = t_alarm.RemoveAlert("module2")
	if err != nil {
		t.Errorf("remove deadline alert fail: %v", err)
	}

    status = t_alarm.GetAlertStatus("module2")
    if status != ALERT_STATUS_NONE {
		t.Errorf("module2 alert's status should be NONE")
    }
}

func TestDeadlineAlert2(t *testing.T) {
    var err error 
    var status int

    //Add alert
	duration := time.Second * 3
    alert := NewDeadlineAlert("module3", "用户数据没有提交", duration, t_cb)
	err = t_alarm.AddAlert("module3", alert)
	if err != nil {
		t.Errorf("add deadline alert fail: %v", err)
	}

    status = t_alarm.GetAlertStatus("module3")
    if status != ALERT_STATUS_INIT {
		t.Errorf("module3 alert's status should be INIT")
    }

    //Enable alert
	err = t_alarm.EnableAlert("module3")
	if err != nil {
		t.Errorf("enable deadline alert fail: %v", err)
	}

    status = t_alarm.GetAlertStatus("module3")
    if status != ALERT_STATUS_ENABLE {
		t.Errorf("module3 alert's status should be ENABLE")
    }

    //Sleep 
	time.Sleep(time.Second * 2)

    //Disable alert
	err = t_alarm.DisableAlert("module3")
	if err != nil {
		t.Errorf("disable deadline alert fail: %v", err)
	}

    status = t_alarm.GetAlertStatus("module3")
    if status != ALERT_STATUS_DISABLE {
		t.Errorf("module3 alert's status should be DISABLE")
    }

    //Remove alert
	err = t_alarm.RemoveAlert("module3")
	if err != nil {
		t.Errorf("remove deadline alert fail: %v", err)
	}

    status = t_alarm.GetAlertStatus("module3")
    if status != ALERT_STATUS_NONE {
		t.Errorf("module3 alert's status should be NONE")
    }
}

func TestTimerAlert1(t *testing.T) {
    var err error
    var status int

    //Add alert
    timer := time.Now().Add(time.Second * 3)
    alert := NewTimerAlert("module4", "活动数据没有提交", timer, t_cb)
	err = t_alarm.AddAlert("module4", alert)
	if err != nil {
		t.Errorf("add timer alert fail: %v", err)
	}

    status = t_alarm.GetAlertStatus("module4")
    if status != ALERT_STATUS_INIT {
		t.Errorf("module2 alert's status should be INIT")
    }

    //Enable alert
	err = t_alarm.EnableAlert("module4")
	if err != nil {
		t.Errorf("enable timer alert fail: %v", err)
	}

    status = t_alarm.GetAlertStatus("module4")
    if status != ALERT_STATUS_ENABLE {
		t.Errorf("module2 alert's status should be ENABLE")
    }

    //Sleep
    time.Sleep(time.Second * 2)

    //Disable alert
	err = t_alarm.DisableAlert("module4")
	if err != nil {
		t.Errorf("disable timer alert fail: %v", err)
	}

    status = t_alarm.GetAlertStatus("module4")
    if status != ALERT_STATUS_DISABLE {
		t.Errorf("module2 alert's status should be DISABLE")
    }

    //Remove alert
	err = t_alarm.RemoveAlert("module4")
	if err != nil {
		t.Errorf("remove timer alert fail: %v", err)
	}

    status = t_alarm.GetAlertStatus("module4")
    if status != ALERT_STATUS_NONE {
		t.Errorf("module2 alert's status should be NONE")
    }
}

func TestTimerAlert2(t *testing.T) {
    var err error
    var status int

    //Add alert
    timer := time.Now().Add(time.Second * 2)
    alert := NewTimerAlert("module5", "活动数据没有提交", timer, t_cb)
	err = t_alarm.AddAlert("module5", alert)
	if err != nil {
		t.Errorf("add timer alert fail: %v", err)
	}

    status = t_alarm.GetAlertStatus("module5")
    if status != ALERT_STATUS_INIT {
		t.Errorf("module2 alert's status should be INIT")
    }

    //Enable alert
	err = t_alarm.EnableAlert("module5")
	if err != nil {
		t.Errorf("enable timer alert fail: %v", err)
	}

    status = t_alarm.GetAlertStatus("module5")
    if status != ALERT_STATUS_ENABLE {
		t.Errorf("module2 alert's status should be ENABLE")
    }

    //Sleep
    time.Sleep(time.Second * 3)

    //Disable alert
	err = t_alarm.DisableAlert("module5")
	if err != nil {
		t.Errorf("disable timer alert fail: %v", err)
	}

    status = t_alarm.GetAlertStatus("module5")
    if status != ALERT_STATUS_DISABLE {
		t.Errorf("module2 alert's status should be DISABLE")
    }

    //Remove alert
	err = t_alarm.RemoveAlert("module5")
	if err != nil {
		t.Errorf("remove timer alert fail: %v", err)
	}

    status = t_alarm.GetAlertStatus("module5")
    if status != ALERT_STATUS_NONE {
		t.Errorf("module2 alert's status should be NONE")
    }
}

func TestTimerAlert3(t *testing.T) {
    var err error
    var status int

    timer := time.Date(0, 0, 0, 17, 58, 0, 0, time.Local)
    fmt.Printf("Timer: %v\n", timer.Local())

    //Add alert
    alert := NewTimerAlert("module6", "活动数据没有提交", timer, t_cb)
	err = t_alarm.AddAlert("module6", alert)
	if err != nil {
		t.Errorf("add timer alert fail: %v", err)
	}

    status = t_alarm.GetAlertStatus("module6")
    if status != ALERT_STATUS_INIT {
		t.Errorf("module2 alert's status should be INIT")
    }

    //Enable alert
	err = t_alarm.EnableAlert("module6")
	if err != nil {
		t.Errorf("enable timer alert fail: %v", err)
	}

    status = t_alarm.GetAlertStatus("module6")
    if status != ALERT_STATUS_ENABLE {
		t.Errorf("module2 alert's status should be ENABLE")
    }

    //Sleep
    time.Sleep(time.Second * 3)

    //Disable alert
	err = t_alarm.DisableAlert("module6")
	if err != nil {
		t.Errorf("disable timer alert fail: %v", err)
	}

    status = t_alarm.GetAlertStatus("module6")
    if status != ALERT_STATUS_DISABLE {
		t.Errorf("module2 alert's status should be DISABLE")
    }

    //Remove alert
	err = t_alarm.RemoveAlert("module6")
	if err != nil {
		t.Errorf("remove timer alert fail: %v", err)
	}

    status = t_alarm.GetAlertStatus("module6")
    if status != ALERT_STATUS_NONE {
		t.Errorf("module2 alert's status should be NONE")
    }
}

