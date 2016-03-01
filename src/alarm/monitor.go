package alarm

import (
    "time"
)

type Monitor struct {
    name string
    start_time time.Time
    stop_time time.Time
    desc string
}

func NewMonitor(name string, start_time, stop_time time.Time, desc string) *Monitor {
    monitor := &Monitor{}
    monitor.name = name
    monitor.start_time = start_time
    monitor.stop_time = stop_time
    monitor.desc = desc 
    return monitor
}

func (this *Monitor) Start() *Monitor {
    return nil
}

func (this *Monitor) Stop() *Monitor {
    return nil
}