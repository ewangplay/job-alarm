package alarm 

import (
    "testing"
)

func TestAlert(t *testing.T) {
    alarm := &Alarm{}
    
    alarm.Alert("保存数据库失败")
}

