# goland

## 1 基础知识

### 1.1 语法

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
6. int32/uint32
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
    数组 array
    结构 struct
    字符串 string
12. 引用类型 非传统引用
    切片 slice
    类哈希表 map
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
        
## 2 web

### 2.1 表单 web.go

- 表单 Form
    客户端提交数据到服务器
    get post
- 文件上传
    上传表单
    获取上传的文件数据
    文件保存