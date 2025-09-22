# <p style="text-align: center;"> Go 学习笔记 </p>

## 一、go 项目的典型结构：
### 典型项目：
- root
  - Lincese
  - Makefile
  - README.md
  - cmd
    - app1
      - main.go
    - app2
      - main.go
  - go.mod
  - go.sum
  - pkg
    - lib1
      - lib1.go
    - lib2
      - lib2.go
  - vender

### 典型库的项目结构
- root
  - Lincese
  - Makefile
  - README.md
  - go.mod
  - go.sum
  - lig.go
  - lib1
      - lib1.go
    - lib2
      - lib2.go


## 二、go常见特效和功能

### 包的命名
- 以小写形式的单词命名
- 包名尽量与包导入路径（import path）的最后一个路径分段保持一致；如golang.org/x/text/ending的最后路径分段是encoding，那么该路径下包名就应该为encoding。

### 变量/类型/函数和方法
- 建议采用驼峰命名法（词与词之间不使用任何符号）
- 命名风格保持一致，简洁

## iota实现枚举
- Go const语法提供了“隐式重复前一个非空表达式的机制：
   ```go
    const(
        Apple,Banana=11,12
        Strawberry,Grape
        Pear, Watermelon
    )
    ```
- 后两行没有显式给予初始值，Go编译器将为其隐式使用第一行的表达式，上述代码等价于：
 
   ```go
    const(
        Apple,Banana=11,12
        Strawberry,Grape=11,12
        Pear, Watermelon=11,12
    )
  ```

- iota的含义
  - iota是go语言预定义的标识符，它表示的是const声明块中每个常量所处位置在块中的偏移值（从零开始）。
  - 在同一行出现多个iota，其值也是一样的。
  - iota是无类型常量，可以与任何形式的类型转换操作
  - 典型的使用：
  ```go
    const(
        _ =iota   (0)
        monday    (1)
        wensday   (2)
        tuesday=iota  (3)
    )

## 尽量定义零值可用的类型

### go语言的原生类型零值
- Go语言每个原生类型都有其默认值，这个默认值就是这个类型的零值：
    所有整数类型：0
    浮点数：0.0
    bool：false
    字符串类型：“”
    指针/interface/slice/channel/map/functon：nil

### 理解零值可用
- 没有初始化的slice也可操作
  ```go
  var zeroSlice []int
  zeroSlice=append(zeroSilce,1)
 没有初始化的slice的值是nil，但append操作不会因此异常

- 通过nil指针调用方法
  ```go
    func main(){
    var p *net.TCPAddr
    fmt.Println(p)
    }

### 使用复合字面值作为初值构造器
    
    声明变量不赋值也不是好习惯，对于结构题初始化会很麻烦，复合字面值可以解决该问题
    例如
        var s myStruct
        s.name="myname"
        s.age=12
    这种赋值会很繁琐

- 建议使用的复合字面值初始化
  ```go
  s：=myStruct{"myname",12}
  a:=[]int{1,4,1,22,22}
  m:=map[int]strng{1:"mymap1",2:"mymap2"}
  - 复合字面值有两部分组成：
    - 类型：myStruct/[]int/map[int]string
    - 由{}包裹的字面值

- field:value形式的复合字面值初值构造器
  ```go
  p:=mystruct{age:12,myname:"lzs"}
  
  - 数组和切片的复合字面值：
    field：value--> index:value
    num：=【23】int{1:2,22:1}
  - map的复合字面值：key：vale
  map ：=map[int]string{
    1:"valueof1"
  }
  ```

## 认识slice
```go
type slice struct{
    array unsafe.pointer
    len int
    cap int
}
```
slice实际上是底层数组的<mark>描述符</mark>.

### slice的基本使用方法
```go
方式1:
var myqueue =[5]int{1,3,6,7,9}
myslice:=myqueue[1,5]
方式2:
make([]int,len,cap)
如：make([]int,4,8)
方式3:
s1:=myslice[1,3]

注：slice或数组index都是从0开始，slice的引用方式为s[low,high]，实际的最大的index=high-1
```
### slice的高级特效--动态扩张
append可以动态将元素加到数组末尾
注：append会导致cap动态变化，一旦我们追加的元素达到cap上限，go底层会动态分配新的数组，并将原数组的值拷贝过去，并按照一定策略调整cap，同时slice与原数组解除绑定关系，旧的数组由垃圾回收器回收掉。


## map的使用
  - **map的基本属性火注意事项：**
    - map是go语言提供的一种抽象数据类型，它表示一组无序的键值对；
    - map对value没有限制，但对key是有要求的，key的类型需要严格定义支持“==”和“！=”操作；
    - 对为零值的map操作会产生panic，所以map必须进行初始化才能使用，初始化map除了前面提到的复合字面值初始化；
    - map也会出现slice的扩容情况，所以使用cap初始化map是最佳实践；
    - 不要依赖map的遍历顺序；
    - 不要尝试获取map中元素vale的地址；
    - **map不是线程安全的，不支持并发；**
  
  
- **map的基本操作：make初始化**
    ```go
    make(type,cap),如：
        mymap:=make(map[int]string,10)
    map与slice都是<mark>引用类型</mark>
    ```
- **map的基本操作：增加或修改map元素值**
  - map[key]=value
    - 如果key存在则更新value，如果key不存在则插入该key和value
- **map的基本操作：获取数据个数**
  - len(map)
- **map基本操作：查找和读取** 
  - "comma ok"惯用法来进行查找
    ```go
    _,ok:=mymap["key1"]
    if ok{
        //todo 
    }
    这里只关心是否存在，通过忽略符"_"忽略value
    **注：<mark>comma ok是map获取value的最佳实践</mark>**
    ```
- **map的基本操作：删除数据**
  - delete(mymap[key])  
    ```go
        注：即使key不存在，go也不会导致panic
    ```
- **map的基本操作：遍历数据**
    - 使用for range语句遍历map中的数据
    ```go
    for k,v:=range mymap{
        fmt.Println("map's value is %d",v)
    }
    注：每次遍历map的顺序都是不一致的。
    ```

## string的使用
- **go中string的特点：**
  ```go
    type stringStruct struct{
        str unsafe.Pointer
        len int
    }
  ```
  - string一旦声明，无论是常量或者变量，在整个生命周期内都不能改变
  - 即使使用slice，go也会将原string拷贝一份，再进行slice操作，slice的操作不会影响原string对象；
  - 即使通过获取string的地址，粗暴的想修改string内容，go在运行态也会报SIGBUS错误
  - string零值可用
    ```go
        var str string
        fmt.Println(str)  //不会出错
    ```
  - 获取string的时间复杂度是O(1)级别
  - string支持+/+=操作
    ```go
    s:="mystring1"
    s=s+"string2"
    s+="string3"
    ```
  - string支持==/！=/>=/<=/>/<的操作
   ```mermaid
        graph LR
        checkLen["检查字符长度"]
        checkAdd["检查对象地址"]
        checkCon["检查对象内容"]
        checkLen-->checkAdd-->checkCon
    ```
    - 长度不一致，直接返回结果
    - 检测对象一致，直接返回结果
    - 如果前面都通过，则详细比对string内容
  - string对非ASCII字符提供原生支持(rune是unicode编码的类型)
  
  ## 理解go语言的包导入
    Go语言是使用包（package）作为基本单元来组织源码的，可以说一个Go程序就是由一些包链接在一起构建而成的。

  ## 包级别变量声明语句中的表达式求值顺序
    ```go
        var (
          a=b+c
          b=func1()
          c=func1()
          d=3
        )

        func func1(){
          d++
          return d
        }
        
        func main(){
          fmt.Priintln(a,b,c,d)
        }
    ```
    初始化的过程：
     ```mermaid
         graph LR
         a-->b-->c-->d
    ```
   -  **从上到下“ready initialization“元素进行初始化，每轮找到第一个，然后再反复直到全部初始化完成；**

   - **每行表达式的求值顺序则按从左到右直行**
   - **Switch/select语句中的表达式求值属于“懒式加载”**，也就是从左到右，一旦遇到满足条件的值，后续的求值就忽略，提高性能；
   - fallthrough 将执行权直接转移到下一个case执行语句中，会略过case的求值和判断语句。
  
## go编译打包过程

  ```mermaid
  graph LR
  库文件-->goBuild[go build]-->点a[临时生成 .a文件]-->goinstall[go install]-->pkg[.a 文件会放到$GOPATH/pkg/目录下]

  main[执行文件]-->build-->link-->可执行文件
  ```
  - 标准库包的源文件在\$GOROOT/src下，对应的.a文件放在\$GOROOT/pkg/darwin_amd64或linux_amd64

  - go install 过程中link -L的顺序决定了go连接搜索的.a文件的顺序。
  
## 路径和包名
  从编译过程可以知道路径包含两部分：
    - 基础搜索路径
      - 所有包的搜索路径都包含\$GOROOT/src目录
      - >go 1.13版本有两种模式：
        - 金典gopath模式(GO111MODULE=off):$GOPATH/src
        - module-aware模式下(GO111MODULE=on):$GOPATH/pkg/mod
    - 包导入路径
      - 包导入路径最后一部分依然是目录名，不是包名，只是go一般习惯是该路径名称与包名保持一致
      - 