package alarm

import (
	"fmt"
	"time"
)

type Alarm struct {
	cb_alarm func(string, string) error
	monitors map[string]*Monitor
}

func NewAlarm(cb func(string, string) error) *Alarm {
	alarm := &Alarm{}
	alarm.cb_alarm = cb
	alarm.monitors = make(map[string]*Monitor, 0)
	return alarm
}

func (this *Alarm) Alert(module, alarm_msg string) error {
	return this.cb_alarm(module, alarm_msg)
}

func (this *Alarm) SetDeadlineAlert(module string, alarm_msg string, deadline time.Time) error {
	var err error
	var monitor *Monitor

	monitor, ok := this.monitors[module]
	if ok && monitor != nil {
        if monitor.IsEnabled() {
            monitor.Stop()
        }
	}

	monitor = NewMonitor(module, alarm_msg, this.cb_alarm)

	err = monitor.Start(deadline)
	if err != nil {
		return err
	}

	this.monitors[module] = monitor

	return nil
}

func (this *Alarm) UnsetDeadlineAlert(module string) error {
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

func (this *Alarm) AddDeadlineAlert(module, alarm_msg string) error {
	var monitor *Monitor

	monitor, ok := this.monitors[module]
	if ok && monitor != nil {
        return fmt.Errorf("monitor %v already exists", module)
	}

	monitor = NewMonitor(module, alarm_msg, this.cb_alarm)

	this.monitors[module] = monitor

    return nil 
}

func (this *Alarm) EnableDeadlineAlert(module string, deadline time.Time) error {
    var monitor *Monitor

	monitor, ok := this.monitors[module]
	if !ok || monitor == nil {
        return fmt.Errorf("monitor %v not found", module)
	}

    if monitor.IsEnabled() {
        monitor.Stop()
    }

    return monitor.Start(deadline)
}

func (this *Alarm) DisableDeadlineAlert(module string) error {
    var monitor *Monitor

	monitor, ok := this.monitors[module]
	if !ok || monitor == nil {
        return fmt.Errorf("monitor %v not found", module)
	}

    if monitor.IsEnabled() {
        monitor.Stop()
    }

    return nil 
}

func (this *Alarm) RemoveDeadlineAlert(module string) error {
	var monitor *Monitor

	monitor, ok := this.monitors[module]
	if !ok || monitor == nil {
		return fmt.Errorf("monitor %v not found", module)
	}

    if monitor.IsEnabled() {
        monitor.Stop()
    }

	delete(this.monitors, module)

    return nil 
}

func (this *Alarm) PrintMonitors() {
    for module, monitor := range this.monitors {
        fmt.Printf(">> %v\n", module)
        fmt.Println(">> ===============================")
        fmt.Printf(">> Alarm Msg: %v\n", monitor.alarm_msg)
        fmt.Printf(">> Status: %v\n", monitor.status)
        fmt.Println()
    }
}
