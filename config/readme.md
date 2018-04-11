共享的配置文件在这里了

## 测试环境下错误的配置路径

参考： [golang-tests-and-working-directory](https://stackoverflow.com/questions/23847003/golang-tests-and-working-directory)

几个不错的方案：
~~~go

package sample

import (
    "testing"
    "runtime"
)

func TestGetFilename(t *testing.T) {
    _, filename, _, _ := runtime.Caller(0)
    fmt.Println("Current test filename: " + filename)
}

// --------------------------------------- ---------- ---------- ----------  
// 			## 要求你每次手动设置工作目录

package blah_test

import (
    "flag"
    "fmt"
    "os"
    "testing"
)

var (
    cwd_arg = flag.String("cwd", "", "set cwd")
)

func init() {
    flag.Parse()
    if *cwd_arg != "" {
        if err := os.Chdir(*cwd_arg); err != nil {
            fmt.Println("Chdir error:", err)
        }
    }
}

func TestBlah(t *testing.T) {
    t.Errorf("cwd: %+q", *cwd_arg)
}

// --------------------------------------- ---------- ---------- ----------  
//              假设配置文件是在项目目录下的 然后利用查找即可找到

wd, _ := os.Getwd()
for !strings.HasSuffix(wd, "<yourProjectDirName>") {
    wd = filepath.Dir(wd)
}

raw, err := ioutil.ReadFile(fmt.Sprintf("%s/src/conf/conf.dev.json", wd))

~~~

##  支持多环境配置
根据环境变量 选择加载指定的配置文件
可以参考这个做法： https://github.com/brainattica/golang-jwt-authentication-api-sample/blob/master/settings/settings.go