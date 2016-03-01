package alarm

import (
    "time"
    "sync"
)

const (
    STATUS_INIT = 0
    STATUS_RUNNING = 1
    STATUS_CANCEL = 2
)

type Monitor struct {
    name string
    deadline time.Time
    alarm_msg string
    cb_alarm func(string)error
    status int
    wg *sync.WaitGroup
}

func NewMonitor(name string, deadline time.Time, alarm_msg string, cb func(string)error) *Monitor {
    monitor := &Monitor{}
    monitor.name = name
    monitor.deadline = deadline
    monitor.alarm_msg = alarm_msg 
    monitor.cb_alarm = cb
    monitor.status = STATUS_INIT
    monitor.wg = &sync.WaitGroup{}
    return monitor
}

func (this *Monitor) Start() error {
    this.status = STATUS_RUNNING
    this.wg.Add(1)
    
    go func() error {
        for {
            if this.status == STATUS_CANCEL {
                this.status = STATUS_INIT
                this.wg.Done()
                return nil
            }
            
            now := time.Now().Unix()
            if now < this.deadline.Unix() {
                time.Sleep(time.Second)
            } else {
                this.status = STATUS_INIT
                this.wg.Done()
                break
            }
        }

        return this.cb_alarm(this.alarm_msg)
    }()
    
    return nil
}

func (this *Monitor) Stop() error {
    if this.status == STATUS_RUNNING {
        this.status = STATUS_CANCEL
        this.wg.Wait()
    }

    return nil
}
