package str2objs

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func getKeyAndValue(str string) (key, value string, err error) {
	err = errors.New("str invaild")
	item := strings.Split(str, ":")
	if len(item) == 2 {
		err = nil
		key = item[0]
		value = item[1]
	}

	return
}

//Str2obj ...
func Str2obj(text string, pobj interface{}) (err error) {
	err = errors.New("parameter invaild")
	pv := reflect.ValueOf(pobj)
	if pv.Kind() == reflect.Ptr {
		v := pv.Elem()
		if v.Kind() == reflect.Struct {
			err = errors.New("text invaild")
			if strings.HasPrefix(text, v.Type().Name()) {
				start := strings.IndexByte(text, '{')
				end := strings.IndexByte(text, '}')
				if start != -1 && end != -1 && end > start {
					err = nil
					text = text[start+1 : end]
					for _, subtxt := range strings.Split(text, ",") {
						if key, value, err := getKeyAndValue(strings.TrimSpace(subtxt)); err == nil {
							if _, isExist := v.Type().FieldByName(key); isExist {
								//fmt.Printf("name:%s index:%d type:%v json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("json"))
								//字段首字母小写, 不对外导出可以进行读操作, 不能进行写操作
								fmt.Println(v.FieldByName(key).String())
								//如果字段对外不导出, 则此处发生panic
								v.FieldByName(key).SetString(value)
							}
						}
					}
				}
			}
		}
	}

	return
}
