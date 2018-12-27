## 需要安装的插件
  - go-outline
  - go-symbols
  - guru
  - gorename
  - goreturns
  - golint

## 手动安装
国内因为墙的原因，无法自动下载， 手动更新
```bash
git clone https://github.com/golang/tools.git 
go install github.com/golang/lint/golint
go install sourcegraph.com/sqs/goreturns
go install  golang.org/x/tools/cmd/gorename
go install  github.com/newhook/go-symbols
go install  golang.org/x/tools/cmd/guru
```


## 参考文档  
- [VS Code1.4 搭建Golang的开发调试环境](https://blog.csdn.net/u010019717/article/details/52138833)


## golang基本类型 
```golang
bool
string
int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr
byte // uint8 的别名
rune // int32 的别名

float32 float64

complex64 complex128
```


## goroutine 疑惑 

- goroutines.go 线程结束，主程序才会退出
- coolshell-goroutines.go 主程序结束线程退出

没看出有啥区别
### 参考资料

- [Go 语言简介（上）— 特性](https://coolshell.cn/articles/8460.html)
- [Go 语言简介（下）— 特性](https://coolshell.cn/articles/8489.html)
- [GO指南](http://go-tour-zh.appspot.com)
- [gobyexample](https://gobyexample.com/)
- [golang-api神器](https://gowalker.org/fmt)