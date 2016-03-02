package alarm

type Alert interface {
    Enable() error
    Disable() error
    GetStatus() int
}
