package alarm

import (
    "time"
)

type Monitor struct {
    name string
    start_time time.Time
    stop_time time.Time
    desc string
    cb_alarm func(string)error
}

func NewMonitor(name string, start_time, stop_time time.Time, desc string, cb func(string)error) *Monitor {
    monitor := &Monitor{}
    monitor.name = name
    monitor.start_time = start_time
    monitor.stop_time = stop_time
    monitor.desc = desc 
    monitor.cb_alarm = cb
    return monitor
}

func (this *Monitor) Start() error {
    return nil
}

func (this *Monitor) Stop() error {
    return nil
}