package alarm

import (
	"fmt"
	"time"
)

const (
	ALERT_STATUS_NONE = iota
	ALERT_STATUS_INIT
	ALERT_STATUS_ENABLE
	ALERT_STATUS_DISABLE
)

type Alarm struct {
	cb_alarm func(string, string) error
	alerts map[string]Alert
}

func NewAlarm(cb func(string, string) error) *Alarm {
	alarm := &Alarm{}
	alarm.cb_alarm = cb
	alarm.alerts = make(map[string]Alert, 0)
	return alarm
}

func (this *Alarm) Alert(module, alarm_msg string) error {
	return this.cb_alarm(module, alarm_msg)
}

func (this *Alarm) SetDeadlineAlert(module string, alarm_msg string, duration time.Duration) error {
	var err error
	var alert Alert

	alert, ok := this.alerts[module]
	if ok && alert != nil {
		alert.Disable()
	}

	alert = NewDeadlineAlert(module, alarm_msg, duration, this.cb_alarm)

	err = alert.Enable()
	if err != nil {
		return err
	}

	this.alerts[module] = alert

	return nil
}

func (this *Alarm) UnsetAlert(module string) error {
	var err error
	var alert Alert

	alert, ok := this.alerts[module]
	if !ok {
		return fmt.Errorf("alert %v not found", module)
	}

	if alert == nil {
		return fmt.Errorf("alert %v invalid", module)
	}

	err = alert.Disable()
	if err != nil {
		return err
	}

	delete(this.alerts, module)

	return nil
}

func (this *Alarm) AddDeadlineAlert(module, alarm_msg string, duration time.Duration) error {
	var alert Alert

	alert, ok := this.alerts[module]
	if ok && alert != nil {
		return fmt.Errorf("alert %v already exists", module)
	}

	alert = NewDeadlineAlert(module, alarm_msg, duration, this.cb_alarm)

	this.alerts[module] = alert

	return nil
}

func (this *Alarm) EnableAlert(module string) error {
	var alert Alert

	alert, ok := this.alerts[module]
	if !ok || alert == nil {
		return fmt.Errorf("alert %v not found", module)
	}

	alert.Disable()

	return alert.Enable()
}

func (this *Alarm) DisableAlert(module string) error {
	var alert Alert

	alert, ok := this.alerts[module]
	if !ok || alert == nil {
		return fmt.Errorf("alert %v not found", module)
	}

	alert.Disable()

	return nil
}

func (this *Alarm) RemoveAlert(module string) error {
	var alert Alert

	alert, ok := this.alerts[module]
	if !ok || alert == nil {
		return fmt.Errorf("alert %v not found", module)
	}

	alert.Disable()

	delete(this.alerts, module)

	return nil
}

func (this *Alarm) SetTimerAlert(module, alarm_msg string, timer time.Time) error {
	var err error
	var alert Alert

	alert, ok := this.alerts[module]
	if ok && alert != nil {
		alert.Disable()
	}

	alert = NewTimerAlert(module, alarm_msg, timer, this.cb_alarm)

	err = alert.Enable()
	if err != nil {
		return err
	}

	this.alerts[module] = alert

	return nil
}

func (this *Alarm) AddTimerAlert(module, alarm_msg string, timer time.Time) error {
	var alert Alert

	alert, ok := this.alerts[module]
	if ok && alert != nil {
		return fmt.Errorf("alert %v already exists", module)
	}

	alert = NewTimerAlert(module, alarm_msg, timer, this.cb_alarm)

	this.alerts[module] = alert

	return nil
}

func (this *Alarm) GetAlertStatus(module string) int {

	alert, ok := this.alerts[module]
	if !ok || alert == nil {
		return ALERT_STATUS_NONE
	}

	return alert.GetStatus()
}

func (this *Alarm) PrintAlerts() {
	for module, alert := range this.alerts {
		fmt.Println(">> =====================================================")
		fmt.Printf(">> Module: %v\n", module)
		fmt.Printf(">> Status: %v\n", alert.GetStatus())
		fmt.Println()
	}
}
