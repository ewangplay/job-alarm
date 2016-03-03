package alarm

import (
    "time"
    "sync"
)

type TimerAlert struct {
	name      string
	alarm_msg string
	timer     time.Time
	cb_alarm  func(string, string) error
	status    int
	wg        *sync.WaitGroup
}

func NewTimerAlert(name string, alarm_msg string, timer time.Time, cb func(string, string) error) *TimerAlert {
	alert := &TimerAlert{}
	alert.name = name
	alert.alarm_msg = alarm_msg
	alert.timer = timer
	alert.cb_alarm = cb
	alert.status = ALERT_STATUS_INIT
	alert.wg = &sync.WaitGroup{}
	return alert
}

func (this *TimerAlert) Enable() error {
	this.status = ALERT_STATUS_ENABLE
	this.wg.Add(1)

	go func() error {
		for {
			if this.status == ALERT_STATUS_DISABLE {
				this.wg.Done()
				return nil
			}

            now := time.Now()

            if this.timer.Year() > 0 && this.timer.Month() > 0 && this.timer.Day() > 0 {
                if now.Year() < this.timer.Year() {
                    time.Sleep(time.Second)
                    continue
                }

                if now.Month() < this.timer.Month() {
                    time.Sleep(time.Second)
                    continue
                }

                if now.Day() < this.timer.Day() {
                    time.Sleep(time.Second)
                    continue
                }
            }

            if this.timer.Hour() > 0 {
                if now.Hour() < this.timer.Hour() {
                    time.Sleep(time.Second)
                    continue
                }
            }

            if this.timer.Minute() > 0 {
                if now.Minute() < this.timer.Minute() {
                    time.Sleep(time.Second)
                    continue
                }
            }

            if this.timer.Second() > 0 {
                if now.Second() < this.timer.Second() {
                    time.Sleep(time.Second)
                    continue
                }
            }

            this.status = ALERT_STATUS_DISABLE
            this.wg.Done()
            break
		}

		return this.cb_alarm(this.name, this.alarm_msg)
	}()

	return nil
}

func (this *TimerAlert) Disable() error {
	if this.status == ALERT_STATUS_ENABLE {
		this.status = ALERT_STATUS_DISABLE
		this.wg.Wait()
	}

	return nil
}

func (this *TimerAlert) GetStatus() int {
	return this.status
}
