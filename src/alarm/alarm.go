package alarm

import (
	"fmt"
	"time"
)

type Alarm struct {
	cb_alarm func(string) error
	monitors map[string]*Monitor
}

func NewAlarm(cb func(alarm_msg string) error) *Alarm {
	alarm := &Alarm{}
	alarm.cb_alarm = cb
	alarm.monitors = make(map[string]*Monitor, 0)
	return alarm
}

func (this *Alarm) Alert(alarm_msg string) error {
	return this.cb_alarm(alarm_msg)
}

func (this *Alarm) SetDeadlineAlert(module string, deadline time.Time, alarm_msg string) error {
	var err error
	var monitor *Monitor

	monitor, ok := this.monitors[module]
	if ok && monitor != nil {
		return fmt.Errorf("monitor %v already exists", module)
	}

	monitor = NewMonitor(module, deadline, alarm_msg, this.cb_alarm)

	err = monitor.Start()
	if err != nil {
		return err
	}

	this.monitors[module] = monitor

	return nil
}

func (this *Alarm) CancelDeadlineAlert(module string) error {
	var err error
	var monitor *Monitor

	monitor, ok := this.monitors[module]
	if !ok {
		return fmt.Errorf("monitor %v not found", module)
	}

	if monitor == nil {
		return fmt.Errorf("monitor %v invalid", module)
	}

	err = monitor.Stop()
	if err != nil {
		return err
	}

	delete(this.monitors, module)

	return nil
}
