# go 基础
## 优点
自带gc。
静态编译，编译好后，扔服务器直接运行。
语法层支持并发，和拥有同步并发的channel类型，使并发开发变得非常方便。
超级简单的交叉编译，仅需更改环境变量。
## 变量范围
    1）声明在函数内部，是函数的本地值，类似private
    2）声明在函数外部，是对当前包可见(包内所有.go文件都可见)的全局值，类似protect
    3）声明在函数外部且首字母大写是所有包可见的全局值,类似public
## 编译调试
GOPATH=工程根目录；其下应创建src，pkg，bin目录。
bin目录中用于生成可执行文件，pkg目录中用于生成.a文件；
系统编译时 go install abc_name时，系统会到GOPATH的src目录中寻找abc_name目录，然后编译其下的go文件；
golang中的import name，实际是到GOPATH中去寻找name.a, 使用时是该name.a的源码中声明的package 名字。
对于main方法，只能在bin目录下运行 go build path_tomain.go; 可以用-o参数指出输出文件名。
go build 可以加编译参数 -gcflags "-N -l"
### 编译调试包
```shell
go build examplepackage
go test examplepackage
go install examplepackage
```
在pkg文件夹中有一个操作系统文件夹，里面会有examplepackage.a文件
### 编译主程序
```shell
go build goproject.go
```
生成goproject文件
```shell
./goproject
```
运行文件
## 一些小tips
### 分配内存
make 分配内存，返回Type本身（用于slice,map,channel）
new 分配内存，返回指向Type的指针、
### err接口
只要实现了Error()函数，返回值为String的都实现了err接口
### init函数
每个包可以拥有多个init函数
包的每个源文件也可以拥有多个init函数
不同包的init函数按照包导入的依赖关系决定该初始化函数的执行顺序
init 函数在main函数执行之前，自动被调用
### 下划线
```go
import _ "github.com/go-sql-driver/mysql"
```
这句话不直接使用mysql包，只是执行一下这个包的init函数，把mysql的驱动注册到sql包里，然后程序里就可以使用sql包来访问mysql数据库了。
### iota
常量计数器，const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)
使用_跳过某些值
```go
    const (
            n1 = iota //0
            n2        //1
            _
            n4        //3
        )
```

iota声明中间插队
```go
    const (
            n1 = iota //0
            n2 = 100  //100
            n3 = iota //2
            n4        //3
        )
    const n5 = iota //0
```
1左移10位就是1024
```go
    const (
            _  = iota
            KB = 1 << (10 * iota)
            MB = 1 << (10 * iota)
            GB = 1 << (10 * iota)
            TB = 1 << (10 * iota)
            PB = 1 << (10 * iota)
        )
```

多个iota定义在一行
```go
    const (
            a, b = iota + 1, iota + 2 //1,2
            c, d                      //2,3
            e, f                      //3,4
        )
```
### rune
```go
    // 遍历字符串
    func traversalString() {
        s := "pprof.cn博客"
        for i := 0; i < len(s); i++ { //byte
            fmt.Printf("%v(%c) ", s[i], s[i])
        }
        fmt.Println()
        for _, r := range s { //rune
            fmt.Printf("%v(%c) ", r, r)
        }
        fmt.Println()
    }
```
输出：
    112(p) 112(p) 114(r) 111(o) 102(f) 46(.) 99(c) 110(n) 229(å) 141() 154() 229(å) 174(®) 162(¢)
    112(p) 112(p) 114(r) 111(o) 102(f) 46(.) 99(c) 110(n) 21338(博) 23458(客)

```go
s1 := "hello"
byteS1 := []byte(s1)

s2 := "博客"
runeS2 := []rune(s2)
```