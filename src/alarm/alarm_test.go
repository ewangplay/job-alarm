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
    job_monitor := alarm.NewMonitor(start_time, stop_time, "用户数据没有提交").Start()
    
    job_monitor.Stop()
}