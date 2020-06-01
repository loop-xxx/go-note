package session

type Session interface{
	Set(string, interface{})
	Del(string)
	Get(string) (interface{}, bool)
}

type Manager interface {
	GetSessionByID(string) (Session, bool)
	NewSession(uint) (string, Session)
	Done() error
}