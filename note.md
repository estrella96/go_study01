# goland

## 1 基础知识

### 1.1 关键字

- 关键字 小写
    break
    default
    func
    interface
    select
    case
    defer
    go
    map
    struct
    chan
    else
    goto
    package
    switch
    const
    fallthrough
    if
    range
    type
    continue
    for
    import
    return
    var
- 结构
    1. package定义包名
    2. 只有main包可以有main函数
    3. import导入非main包
    4. const定义常量
    5. var定义变量
    6. type+struct/interface 定义结构或接口
    7. func定义函数和方法
 - 导入外部包
```go
package main
import "fmt"
import "os"
import(
    "fmt"
    "os"
)
//别名
import f "fmt"
//省略调用
import . "fmt"
```

- 对外部包的可见性 适用于所有对象 
    - 首字母小写 私有
    - 首字母大写 公开
 
### 1.2 基本类型

1. bool
    - 1位
    - true/flase
2. int/uint(无符号整型)
    - 32/64 位
3. int8/uint8
    - 1字节
    - -128~+127  0~255
4. byte
    uint8的别名
5. int16/uint16
    - 2字节
    - -32768~32767 0-65535
6. int32/uint32 rune-int32(Unicode编码)
    - 4字节
7. int64/uint64
    - 8字节
8. float32/float64
    - 4/8字节
    - 精确到7/15小数位
9. complex64/complex128 复数
    - 8/16字节
10. uintptr 指针类型
    保存指针的32或64位整型
11. 其他
    数组 array `var ipv4 [4]uint8=[4]uint8{192,168,0,1}`
    结构 struct
    字符串 string
12. 引用类型 非传统引用
    切片 slice 对数组的包装
    类哈希表 map `var ipSwitches=map[String]bool{}  //string,bool`
    管道 chan
13. 接口型 interface
14. 函数型 func

### 1.3 类型的零值与别名

- 类型零值
    类型的默认值
    int 0
    bool false
    float64 0.0
    string ""
    结构 空结构
    引用类型 nil(空指针)

- 类型别名 可以用中文 但不建议
    `type Int int32` 
    1. 为某个类型赋予新的合法名称
    2. 使用源类型作为底层类型，定义了一个全新的类型
    3. 新类型与源类型之间需要类型转化 属于相互兼容的类型
    4. 新类型可以拥有属于自己的一套方法
    5. 新类型不会继承源类型的方法，需要类型转换后才可以调用
 
### 1.4 变量的声明与赋值
 
- 单个变量的声明和赋值
    1. 声明 var <变量名称> <变量类型>
    2. 赋值 <变量名称> = <表达式>
    3. 声明并赋值 var与:重复
        var <变量名称> [变量类型] = <表达式>
        <变量名称> := <表达式>  (类型推断)
    4. 抛弃赋值
        (函数返回值下划线)
        _ = <表达式>
- 多个变量的声明和赋值
    1. var()
    2. 并行赋值 变量类型相同
        var <变量1>,<变量2> [变量类型]=<表达式1>, <表达式2>
        <变量1>,<变量2> := <表达式1>, <表达式2> 
        a,b,c:=1.1,2,3
- 类型转换
    1. 所有类型转换都要显式声明
    2. 类型转换必须发生在两种互相兼容的类型之间
    3. 格式
        <变量2> [:]= <类型2>(<变量1>)
        `var a=float64=1.1
        var b int
        b=int(a)
        c:=int(a)`
   
### 1.5 操作符
    ||
    &&
    ==
    !=
    |
    ^
    %
    << >>
    <- 接收操作
  
### 1.6 流程控制

- if
    if...else if...else
- switch
- for
- defer
    延迟调用指定的函数
    defer 调用函数的表达式
    外围函数语句正常执行完毕 延迟函数都执行完毕外围函数才结束
    延迟函数执行完毕 外围函数才return
    延迟函数中使用外部变量需要参数引入
    多个延迟函数调用的执行顺序与defer语句顺序相反 底层是栈
- panic和recover
    1. panic
        报告运行期间的致命错误
        停止当前的控制流程并引发一个运行时恐慌
    2. recover
        拦截运行时恐慌的内建函数
        使当前程序从恐慌状态中恢复并重新获得流程控制权
        返回一个interface{}类型的结果 

```go
package main
import (
	"errors"
"fmt"
)
func main() {
    
	//recover 
    defer func(){
        if p:=recover();p!=nil{
            fmt.Printf("Recovered panic: %s\n",p)
        }
    }()

	//go run hello.go
	var number int
    if 100<number{
        number++
    } else{
        number--
    }

    switch number{
    default:
        fmt.Println("unknown")
    case 1:
        fmt.Println("hello")
    case 2,3,4:
        fmt.Println("www")
    }
    
    for i:=0;i<100;i++{
        number++
    }
    
    ints :=[]int{1,2,3,4,5}
    for i,d :=range ints{
        fmt.Printf("Index: %d, Value :%d\n",i,d)
    }
    
	//panic函数
    outerFunc()


}

func outerFunc(){
    innerFunc()
}

func innerFunc(){
    panic(errors.New("An intended fatal error!"))}
```


   
## 2 web

### 2.1 表单 web.go

- 表单 Form
    客户端提交数据到服务器
    get post
- 文件上传
    上传表单
    获取上传的文件数据
    文件保存
 
 
## 3 代码规范 

### 3.1 一般规范

- Comment Sentences
    声明的注释是句子
    以事物开头.结尾
    ```go
    // A Request represents a request to run a command.
    type Request struct { ...
    
    // Encode writes the JSON encoding of req to w.
    func Encode(w io.Writer, req *Request) { ...
    ```
- Declaring Empty Slices 声明空的切片
    `var t []string` 可以避免分配内存
    不用
    `t := []string{}`
- Doc Comments 文档注释
    1. 所有顶级的导出名称都应具有文档注释，非平凡的未导出类型或函数声明也应具有文档注释
- Don't Panic
    对于正常的错误处理不要用panic
    使用多个error和返回值
- Error Strings
    错误字符串小写 不需要以标点符号结束
    因为错误信息可能在句子中间打印
- Handle Errors 
    不要用_丢弃错误
    处理错误并返回
    或者必须情况下panic
- Imports
    成组分段
    ```go
    package main
    
    import (
        "fmt"
        "hash/adler32"
        "os"
        
        "appengine/foo"
        "appengine/user"
        
        "code.google.com/p/x/y"
        "github.com/foo/bar"
    )
    ```
- Import Dot
    import . 可以用在tests中，不能是在被测试的包中
    ```go
    package foo_test
    
    import (
        "bar/testutil" // also imports "foo"
        . "foo"
    )
    ```
- Indent Error Flow
    使普通代码路径保持最小缩进，并缩进错误处理，并首先对其进行处理。 
    ```go
    if err != nil {
        // error handling
        return // or continue, etc.
    }
    // normal code
    ```
- Initialisms
    缩写词大写 ServeHTTP URL appleID
- Line Length
    不限制 但避免太长的行
- Mixed Caps
    未导出的常量maxLength
- Named Result Parameters 结果参数命名
    返回值较少可以不声明类型
    ```go
    func (n *Node) Parent1() *Node
    func (n *Node) Parent2() (*Node, error)
    ```
- Package Comments    
    在包上面 无空行
- Package Names
    在包中可以省略包名
- Pass Values 传参
    不要仅将指针作为函数参数传递
    如果函数在整个过程中仅将其参数x称为* x，则该参数不应为指针。 
    常见的实例包括传递指向字符串的指针（* string）或指向接口值的指针（* io.Reader，值本身都是固定大小，可以直接传递
    此建议不适用于大型结构，甚至不适用于可能增长的小型结构。
- Receiver Names
    方法接收变量
    一或两个单词缩写 如c cl 为Client
    不要使用me this self
- Receiver Type
    值还是指针接收 value pointer
    1. receiver是map,func,chan 不使用pointer
    2. slice方法没有重新切片 不使用pointer
    3. 方法改变了receiver 使用pointer
    4. 包含同步片段的结构 使用pointer
    5. 大的struct或array pointer更有效
- Useful Test Failures
    test要显示有用的错误信息    
- Variable Names
    变量名要短

### 3.2 常见规范

- Package
    1. 若干.go文件
    2. go文件内部顺序
        包声明
        包导入
        常量定义
        全局变量定义
        类型定义
        方法/函数定义
    2. import规范
        除测试怒允许使用 . 
        不允许使用别名，同时导入2个相同的包名除外
        使用小括号 包名要另起一行
        标准库 程序内部 第三方 顺序引入包
- Naming
    package名和目录名保持一致，简短有意义，小写，单数
    变量、函数、结构体、接口统一采用驼峰法则，避免使用_
    源码文件命名，小写
    常量声明 驼峰法 需要导出的首字母大写 不需要导出的首字母小写
    专有名词 全部大写或全部小写
- Error Handling
    错误处理或向上抛出
    除了影响系统启动等关键步骤可以panic外 其他避免使用panic
    避免用_忽略错误
    统一处理错误 
- Project Layout
          