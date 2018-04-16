package main

import (
	"fmt"
	"log"

	"github.com/zpatrick/go-config"
)

func main() {
	iniFile := config.NewINIFile("config.ini")

	c := config.NewConfig([]config.Provider{iniFile})

	// Optional
	c.Validate = func(settings map[string]string) error {
		// var val string
		if _, ok := settings["global.timeout"]; !ok {
			return fmt.Errorf("Required setting 'global.timeout' not set!")
		}
		// log.Printf(" we check the global.timeout value : %s \n ", val)
		return nil
	}

	if err := c.Load(); err != nil {
		log.Fatal(err)
	}
	// <section>.<item>
	// 都是字符串需要使用转换方法：	c.Int(), c.Float(), c.Bool()
	timeout, err := c.Int("global.timeout")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(timeout)

	// 不存在的默认配置
	someKey, err := c.StringOr("local.someKey", "some default val")
	if err != nil {
		log.Println(err)
	}
	fmt.Println(someKey)

	// 获取所有的设置
	settings, err := c.Settings()
	if err != nil {
		log.Fatalln(err)
	}

	for key, val := range settings {
		fmt.Printf("%s = %s \n", key, val)
	}

}
