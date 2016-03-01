package alarm

import (
    "fmt"
    "time"
)

type Alarm struct {
    cb_alarm func(string)error
    monitors map[string]*Monitor
}

func NewAlarm(cb func(desc string)error) *Alarm {
    alarm := &Alarm{}
    alarm.cb_alarm = cb
    alarm.monitors = make(map[string]*Monitor, 0)
    return alarm
}

func (this *Alarm) Alert(desc string) error {
    return this.cb_alarm(desc)
}

func (this *Alarm) AddMonitor(name string, start_time, stop_time time.Time, desc string) error {
    var err error
    var monitor *Monitor
    
    monitor, ok := this.monitors[name]
    if ok && monitor != nil {
        return fmt.Errorf("monitor %v already exists", name)
    }
    
    monitor = NewMonitor(name, start_time, stop_time, desc, this.cb_alarm)
    
    err = monitor.Start()
    if err != nil {
        return err
    }
    
    this.monitors[name] = monitor
    
    return nil
}

func (this *Alarm) RemoveMonitor(name string) error{
    var err error
    var monitor *Monitor
    
    monitor, ok := this.monitors[name]
    if !ok {
        return fmt.Errorf("monitor %v not found", name)
    }
    
    if monitor == nil {
        return fmt.Errorf("monitor %v invalid", name)
    }
    
    err = monitor.Stop()
    if err != nil {
        return err
    }
    
    delete(this.monitors, name)
    
    return nil
}