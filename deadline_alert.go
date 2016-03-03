package alarm

import (
    "time"
    "sync"
)

type DeadlineAlert struct {
    name string
    alarm_msg string
    duration time.Duration
    cb_alarm func(string, string)error
    status int
    wg *sync.WaitGroup
}

func NewDeadlineAlert(name string, alarm_msg string, duration time.Duration, cb func(string, string)error) *DeadlineAlert {
    alert := &DeadlineAlert{}
    alert.name = name
    alert.alarm_msg = alarm_msg 
    alert.duration = duration
    alert.cb_alarm = cb
    alert.status = ALERT_STATUS_INIT
    alert.wg = &sync.WaitGroup{}
    return alert
}

func (this *DeadlineAlert) Enable() error {
    this.status = ALERT_STATUS_ENABLE
    this.wg.Add(1)

    enable_time := time.Now()
    
    go func() error {
        for {
            if this.status == ALERT_STATUS_DISABLE {
                this.wg.Done()
                return nil
            }
            
            if time.Since(enable_time) < this.duration {
                time.Sleep(time.Second)
            } else {
                this.status = ALERT_STATUS_DISABLE
                this.wg.Done()
                break
            }
        }

        return this.cb_alarm(this.name, this.alarm_msg)
    }()
    
    return nil
}

func (this *DeadlineAlert) Disable() error {
    if this.status == ALERT_STATUS_ENABLE {
        this.status = ALERT_STATUS_DISABLE
        this.wg.Wait()
    }

    return nil
}

func (this *DeadlineAlert) GetStatus() int {
    return this.status
}
