package alarm 

import (
    "testing"
    "fmt"
)

func TestAlert(t *testing.T) {
    f := func (desc string) error {
        fmt.Println(desc)
        return nil
    }
    
    alarm := NewAlarm(f)
    
    alarm.Alert("保存数据库失败")
}

