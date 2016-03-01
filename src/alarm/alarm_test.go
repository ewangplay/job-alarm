package alarm 

import (
    "testing"
    "fmt"
    "time"
)

func TestAlert(t *testing.T) {
    f := func (desc string) error {
        fmt.Println(desc)
        return nil
    }
    
    alarm := NewAlarm(f)
    
    alarm.Alert("保存数据库失败")
}

func TestMonitor(t *testing.T) {
    f := func (desc string) error {
        fmt.Println(desc)
        return nil
    }
    
    alarm := NewAlarm(f)
    
    start_time := time.Now() 
    stop_time := start_time.Add(24 * time.Hour)
    err := alarm.AddMonitor("test", start_time, stop_time, "用户数据没有提交")
    if err != nil {
        t.Errorf("add monitor fail: %v", err)
    }
    
    err = alarm.RemoveMonitor("test")
    if err != nil {
        t.Errorf("remove monitor fail: %v", err)
    }
}