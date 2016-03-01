报警器
===================================

* 实现一个go包用来在特定条件下触发报警机制
* 该go包提供触发报警的api
* 该ga包提供关闭报警的api
* 使用TDD的方式来实现

## 需求整理
* 对于一些程序中的错误，需要直接发送报警
* 对于周期性执行的操作，如果在周期内没有执行，那么触发报警
* 周期性报警检测可以关闭

## 如何使用
```
git clone https://github.com/ewangplay/job-alarm.git
cd job-alarm
export GOPATH=/absolute/path/to/job-alarm # replace with your absolute path
go test alarm # run tests
```