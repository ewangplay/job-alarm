package alarm

import (
    "time"
)

type Alarm struct {
    cb_alarm func(string)error
}

func NewAlarm(cb func(desc string)error) *Alarm {
    alarm := &Alarm{}
    alarm.cb_alarm = cb
    return alarm
}

func (this *Alarm) Alert(desc string) error {
    return this.cb_alarm(desc)
}

func (this *Alarm) NewMonitor(start_time, stop_time time.Time, desc string) *Monitor {
    return nil
}