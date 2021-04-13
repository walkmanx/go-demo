package greeting

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	// 如果name没有值，则返回一个错误信息
	if name == "" {
		return "", errors.New("Empty name")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}
	return messages, nil
}

// Go在程序启动时，全局变量初始化后自动执行init函数
func init() {
	rand.Seed(time.Now().UnixNano())
}

// 随即返回一个欢迎信息格式
func randomFormat() string {
	formats := []string{
		"Hi,%v. Welcome!",
		"Great to see you,%v!",
		"Hail,%v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}
