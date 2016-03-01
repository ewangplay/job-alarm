package alarm

import (
    "time"
    "sync"
)

type Monitor struct {
    name string
    stop_time time.Time
    desc string
    cb_alarm func(string)error
    status int
    wg *sync.WaitGroup
}

func NewMonitor(name string, stop_time time.Time, desc string, cb func(string)error) *Monitor {
    monitor := &Monitor{}
    monitor.name = name
    monitor.stop_time = stop_time
    monitor.desc = desc 
    monitor.cb_alarm = cb
    monitor.status = 0
    monitor.wg = &sync.WaitGroup{}
    return monitor
}

func (this *Monitor) Start() error {
    this.status = 1
    this.wg.Add(1)
    
    go func() error {
        for {
            if this.status == 2 {
                this.status = 0
                this.wg.Done()
                return nil
            }
            
            now := time.Now().Unix()
            if now < this.stop_time.Unix() {
                time.Sleep(time.Second)
            } else {
                this.status = 0
                this.wg.Done()
                break
            }
        }

        return this.cb_alarm(this.desc)
    }()
    
    return nil
}

func (this *Monitor) Stop() error {
    if this.status == 1 {
        this.status = 2
        this.wg.Wait()
    }

    return nil
}