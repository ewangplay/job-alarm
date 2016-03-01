package alarm

type Alarm struct {
    cb_alarm func(string)error
}

func NewAlarm(cb func(desc string)error) *Alarm {
    alarm := &Alarm{}
    alarm.cb_alarm = cb
    return alarm
}

func (this *Alarm) Alert(desc string) error {
    return nil
}