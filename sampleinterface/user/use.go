package user

// 使う側で interface を定義する
type module interface {
	Hello() string
}

func Greet(m module) string {
	return m.Hello()
}
