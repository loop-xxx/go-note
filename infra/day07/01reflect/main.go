package main

import (
	"fmt"
	"reflect"
	"strconv"
)

type mysqlInfo struct {
	Address string `loop:"addr"`
	User    string `loop:"user"`
	Passwd  string `loop:"passwd"`
	Port    uint32 `loop:"p"`
	Max     uint32 `loop:"m"`
}

type ridesInfo struct {
	Address string
	User    string
	Passwd  string
	Port    uint32
	Test    bool
}
type databaseInfo struct {
	mysqlInfo `loop:"mysql"`
	ridesInfo `loop:"rides"`
}

var data = map[string]map[string]string{
	"mysql": map[string]string{
		"addr":   "192.161.0.1",
		"user":   "root",
		"passwd": "toor",
		"p":      "2333",
		"m":      "1024",
		"test":   "true",
	},
	"rides": map[string]string{
		"Address": "192.161.0.2",
		"User":    "loop",
		"Passwd":  "root",
		"Port":    "666",
		"Test":    "true",
	},
}

func loadSubConfig(value map[string]string, objt reflect.Type, objv reflect.Value) {
itemloop:
	for key, item := range value {
		for j := int(0); j < objt.NumField(); j++ {
			if objt.Field(j).Tag.Get("loop") == key || objt.Field(j).Name == key {
				//fmt.Println(pt.Elem().Field(i).Type.Field(j).Tag.Get("loop"), pt.Elem().Field(i).Type.Field(j).Name)
				switch objv.Field(j).Kind() {
				case reflect.String:
					objv.Field(j).SetString(item)
				case reflect.Uint32:
					if number, err := strconv.ParseUint(item, 10, 64); err == nil {
						objv.Field(j).SetUint(number)
					} else {
						fmt.Printf("[WARING] type error %s can't cast to int\n", item)
					}
				case reflect.Bool:
					if bo, err := strconv.ParseBool(item); err == nil {
						objv.Field(j).SetBool(bo)
					} else {
						fmt.Printf("[WARING] type error %s can't cast to bool\n", item)
					}
				case reflect.Int, reflect.Uint, reflect.Int64:
					fmt.Println("emmmm~")
				default:
					fmt.Println("sory!")
				}
				continue itemloop
			}
		}
	}
}

func loadCofig(data map[string]map[string]string, obj interface{}) (loaderr error) {
	//kind为Ptr 的Value,Tyep都可以调用Elem()函数, Value.Elem() 返回Value类型(指针指向的值), Type.Elem()返回Type类型(指针是什么类型的指针)
	pt := reflect.TypeOf(obj)
	pv := reflect.ValueOf(obj)
	loaderr = fmt.Errorf("%#v must be struct pointer", obj)
	if pt.Kind() == reflect.Ptr && pt.Elem().Kind() == reflect.Struct {
		loaderr = nil
	valueloop:
		for key, value := range data {
			//kind为Struct 的Value,Type都可以调用Field()函数, Value.Field(i) 返回Value类型(字段的值) Type.Field(i) 返回StructField类型
			for i := int(0); i < pt.Elem().NumField(); i++ {
				if pt.Elem().Field(i).Tag.Get("loop") == key || pt.Elem().Field(i).Name == key {
					loadSubConfig(value, pt.Elem().Field(i).Type, pv.Elem().Field(i))
					/**
					itemloop:
						for key, item := range value {
							for j := int(0); j < pv.Elem().Field(i).NumField(); j++ {
								if pt.Elem().Field(i).Type.Field(j).Tag.Get("loop") == key || pt.Elem().Field(i).Type.Field(j).Name == key {
									//fmt.Println(pt.Elem().Field(i).Type.Field(j).Tag.Get("loop"), pt.Elem().Field(i).Type.Field(j).Name)
									switch pv.Elem().Field(i).Field(j).Kind() {
									case reflect.String:
										pv.Elem().Field(i).Field(j).SetString(item)
									case reflect.Uint32:
										if number, err := strconv.ParseInt(item, 10, 64); err == nil {
											pv.Elem().Field(i).Field(j).SetUint(uint64(number))
										} else {
											fmt.Printf("[WARING] type error %s can't cast int", item)
										}
									case reflect.Bool:
										if bo, err := strconv.ParseBool(item); err == nil {
											pv.Elem().Field(i).Field(j).SetBool(bo)
										} else {
											fmt.Printf("[WARING] type error %s can't cast bool", item)
										}
									case reflect.Int, reflect.Uint, reflect.Int64:
										fmt.Println("emmmm~")
									default:
										fmt.Println("sory!")
									}
									continue itemloop
								}
							}
						}
					*/
					continue valueloop
				}
			}
		}

	}
	return
}

func main() {
	var variable databaseInfo
	if err := loadCofig(data, &variable); err == nil {
		fmt.Println(variable)
	} else {
		fmt.Printf("[ERROR] loadCofig:%v\n", err)
	}
}
