package flash

import (
	"encoding/gob"
	"github.com/klein/go-mvc/pkg/session"
)

type Flashes map[string]interface{}

//存入会话数据的key

var flashKey = "_flashes"

func init()  {
	//在 gorilla/session 中存储的map 和 struct 数据
	//需要提前注册 gob 方便后续 gob 序列化编码和解码
	gob.Register(Flashes{})
}

//info 添加 Info 类型的消息提示
func Info(message string)  {
	addFlash("info", message)
}

//Warning 添加 Warning 类型的消息提示
func Warning(message string)  {
	addFlash("warning", message)
}

//Success 添加 Success 类型的消息提示
func Success(message string)  {
	addFlash("success", message)
}

//Danger 添加 Danger 类型的消息提示
func Danger(message string)  {
	addFlash("danger", message)
}

//获取所有消息
func All() Flashes {
	val := session.Get(flashKey)
	//读取是必须做类型检测
	flashMessages, ok := val.(Flashes)
	if !ok {
		return nil
	}
	//读取即销毁，直接删除
	session.Forget(flashKey)
	return flashMessages
}

//私有方法 新增一条提示
func addFlash(key string, message string)  {
	flashes := Flashes{}
	flashes[key] = message
	session.Set(flashKey, flashes)
}