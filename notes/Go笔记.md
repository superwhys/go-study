# Go

### 一、准备

笔记博客地址：https://www.liwenzhou.com/posts/Go/go_menu/



### 二、HelloWorld

1. 程序

![image-20210807222244007](pic/image-20210807222244007.png)

2. 编译

   go build

   ​	在当前目录下： `go build` 

   ​	在其他目录下：`go build hello`

   ​								go编译器会去 `GOPATH`的`src`目录下查找你要编译的`hello`项目

   ​	我们还可以使用`-o`参数来指定编译后得到的可执行文件的名字。

   ​	`go build -o heiheihei.exe`

   

   go run

   ​	直接执行go文件

   


   go install

   ​	分为两步：

   ​			1.先编译得到一个可执行文件

   ​			2.将可执行文件拷贝到bin目录下

### 三、跨平台编译

默认我们`go build`的可执行文件都是当前操作系统可执行的文件，如果我想在windows下编译一个linux下可执行文件，那需要怎么做呢？

只需要指定目标操作系统的平台和处理器架构即可，例如Windows平台cmd下按如下方式指定环境变量。

```bash
SET CGO_ENABLED=0  // 禁用CGO
SET GOOS=linux  // 目标平台是linux
SET GOARCH=amd64  // 目标处理器架构是amd64
```

注意：如果你使用的是PowerShell终端，那么设置环境变量的语法为`$ENV:CGO_ENABLED=0`。

然后再执行`go build`命令，得到的就是能够在Linux平台运行的可执行文件了。

Mac 下编译 Linux 和 Windows平台 64位 可执行程序：

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
```

Linux 下编译 Mac 和 Windows 平台64位可执行程序：

```bash
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
```

Windows下编译Mac平台64位可执行程序：

```bash
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build
```



### 四、go module

`go module`是Go1.11版本之后官方推出的版本管理工具，并且从Go1.13版本开始，`go module`将是Go语言默认的依赖管理工具。

#### A. GO111MODULE

要启用`go module`支持首先要设置环境变量`GO111MODULE`，通过它可以开启或关闭模块支持，它有三个可选值：`off`、`on`、`auto`，默认值是`auto`。

1. `GO111MODULE=off`禁用模块支持，编译时会从`GOPATH`和`vendor`文件夹中查找包。
2. `GO111MODULE=on`启用模块支持，编译时会忽略`GOPATH`和`vendor`文件夹，只根据 `go.mod`下载依赖。
3. `GO111MODULE=auto`，当项目在`$GOPATH/src`外且项目根目录有`go.mod`文件时，开启模块支持。

简单来说，设置`GO111MODULE=on`之后就可以使用`go module`了，以后就没有必要在GOPATH中创建项目了，并且还能够很好的管理项目依赖的第三方包信息。

使用 go module 管理依赖后会在项目根目录下生成两个文件`go.mod`和`go.sum`。

#### B. GOPROXY

Go1.11之后设置GOPROXY命令为：

```bash
export GOPROXY=https://goproxy.cn
```

Go1.13之后`GOPROXY`默认值为`https://proxy.golang.org`，在国内是无法访问的，所以十分建议大家设置GOPROXY，这里我推荐使用[goproxy.cn](https://studygolang.com/topics/10014)。

```bash
go env -w GOPROXY=https://goproxy.cn,direct
```

#### C. go mod命令

常用的`go mod`命令如下：

```
go mod download    下载依赖的module到本地cache（默认为$GOPATH/pkg/mod目录）
go mod edit        编辑go.mod文件
go mod graph       打印模块依赖图
go mod init        初始化当前文件夹, 创建go.mod文件
go mod tidy        增加缺少的module，删除无用的module
go mod vendor      将依赖复制到vendor下
go mod verify      校验依赖
go mod why         解释为什么需要依赖
```

#### D. go.mod

go.mod文件记录了项目所有的依赖信息，其结构大致如下：

```sh
module github.com/Q1mi/studygo/blogger

go 1.12

require (
	github.com/DeanThompson/ginpprof v0.0.0-20190408063150-3be636683586
	github.com/gin-gonic/gin v1.4.0
	github.com/go-sql-driver/mysql v1.4.1
	github.com/jmoiron/sqlx v1.2.0
	github.com/satori/go.uuid v1.2.0
	google.golang.org/appengine v1.6.1 // indirect
)
```

其中，

- `module`用来定义包名
- `require`用来定义依赖包及版本
- `indirect`表示间接引用

##### 依赖的版本

go mod支持语义化版本号，比如`go get foo@v1.2.3`，也可以跟git的分支或tag，比如`go get foo@master`，当然也可以跟git提交哈希，比如`go get foo@e3702bed2`。关于依赖的版本支持以下几种格式：

```go
gopkg.in/tomb.v1 v1.0.0-20141024135613-dd632973f1e7
gopkg.in/vmihailenco/msgpack.v2 v2.9.1
gopkg.in/yaml.v2 <=v2.2.1
github.com/tatsushid/go-fastping v0.0.0-20160109021039-d7bb493dee3e
latest
```

##### replace

在国内访问`golang.org/x`的各个包都需要翻墙，你可以在`go.mod`中使用replace替换成`github`上对应的库。

```go
replace (
	golang.org/x/crypto v0.0.0-20180820150726-614d502a4dac => github.com/golang/crypto v0.0.0-20180820150726-614d502a4dac
	golang.org/x/net v0.0.0-20180821023952-922f4815f713 => github.com/golang/net v0.0.0-20180826012351-8a410e7b638d
	golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0
)
```



#### E. go get

在项目中执行`go get`命令可以下载依赖包，并且还可以指定下载的版本。

1. 运行`go get -u`将会升级到最新的次要版本或者修订版本(x.y.z, z是修订版本号， y是次要版本号)
2. 运行`go get -u=patch`将会升级到最新的修订版本
3. 运行`go get package@version`将会升级到指定的版本号version

如果下载所有依赖可以使用`go mod download`命令。 

#### F. 整理依赖

我们在代码中删除依赖代码后，相关的依赖库并不会在`go.mod`文件中自动移除。这种情况下我们可以使用`go mod tidy`命令更新`go.mod`中的依赖关系。

#### G. go mod edit

##### 格式化

因为我们可以手动修改go.mod文件，所以有些时候需要格式化该文件。Go提供了一下命令：

```bash
go mod edit -fmt
```

##### 添加依赖项

```bash
go mod edit -require=golang.org/x/text
```

##### 移除依赖项

如果只是想修改`go.mod`文件中的内容，那么可以运行`go mod edit -droprequire=package path`，比如要在`go.mod`中移除`golang.org/x/text`包，可以使用如下命令：

```bash
go mod edit -droprequire=golang.org/x/text
```

关于`go mod edit`的更多用法可以通过`go help mod edit`查看。



#### H. 在项目中使用go module

##### 既有项目

如果需要对一个已经存在的项目启用`go module`，可以按照以下步骤操作：

1. 在项目目录下执行`go mod init`，生成一个`go.mod`文件。
2. 执行`go get`，查找并记录当前项目的依赖，同时生成一个`go.sum`记录每个依赖库的版本和哈希值。

##### 新项目

对于一个新创建的项目，我们可以在项目文件夹下按照以下步骤操作：

1. 执行`go mod init 项目名`命令，在当前项目文件夹下创建一个`go.mod`文件。
2. 手动编辑`go.mod`中的require依赖项或执行`go get`自动发现、维护依赖。

#### I. go module导入本地包

`go module`是Go1.11版本之后官方推出的版本管理工具，并且从`Go1.13`版本开始，`go module`将是Go语言默认的依赖管理工具。到今天`Go1.14`版本推出之后`Go modules` 功能已经被正式推荐在生产环境下使用了。

这几天已经有很多教程讲解如何使用`go module`，以及如何使用`go module`导入gitlab私有仓库，我这里就不再啰嗦了。但是最近我发现很多小伙伴在群里问如何使用`go module`导入本地包，作为初学者大家刚开始接触package的时候肯定都是先在本地创建一个包，然后本地调用一下，然后就被卡住了。。。

这里就详细介绍下如何使用`go module`导入本地包。

##### 前提

假设我们现在有`moduledemo`和`mypackage`两个包，其中`moduledemo`包中会导入`mypackage`包并使用它的`New`方法。

`mypackage/mypackage.go`内容如下：

```go
package mypackage

import "fmt"

func New(){
	fmt.Println("mypackage.New")
}
```

我们现在分两种情况讨论：

##### 在同一个项目下

**注意**：在一个项目（project）下我们是可以定义多个包（package）的。

目录结构

现在的情况是，我们在`moduledemo/main.go`中调用了`mypackage`这个包。

```bash
moduledemo
├── go.mod
├── main.go
└── mypackage
    └── mypackage.go
```

导入包

这个时候，我们需要在`moduledemo/go.mod`中按如下定义：

```go
module moduledemo

go 1.14
```

然后在`moduledemo/main.go`中按如下方式导入`mypackage`

```go
package main

import (
	"fmt"
	"moduledemo/mypackage"  // 导入同一项目下的mypackage包
)
func main() {
	mypackage.New()
	fmt.Println("main")
}
```

举个例子

举一反三，假设我们现在有文件目录结构如下：

```bash
└── bubble
    ├── dao
    │   └── mysql.go
    ├── go.mod
    └── main.go
```

其中`bubble/go.mod`内容如下：

```go
module github.com/q1mi/bubble

go 1.14
```

`bubble/dao/mysql.go`内容如下：

```go
package dao

import "fmt"

func New(){
	fmt.Println("mypackage.New")
}
```

`bubble/main.go`内容如下：

```go
package main

import (
	"fmt"
	"github.com/q1mi/bubble/dao"
)
func main() {
	dao.New()
	fmt.Println("main")
}
```

##### 不在同一个项目下

目录结构

```bash
├── moduledemo
│   ├── go.mod
│   └── main.go
└── mypackage
    ├── go.mod
    └── mypackage.go
```

导入包

这个时候，`mypackage`也需要进行module初始化，即拥有一个属于自己的`go.mod`文件，内容如下：

```go
module mypackage

go 1.14
```

然后我们在`moduledemo/main.go`中按如下方式导入：

```go
import (
	"fmt"
	"mypackage"
)
func main() {
	mypackage.New()
	fmt.Println("main")
}
```

因为这两个包不在同一个项目路径下，你想要导入本地包，并且这些包也没有发布到远程的github或其他代码仓库地址。这个时候我们就需要在`go.mod`文件中使用`replace`指令。

在调用方也就是`moduledemo/go.mod`中按如下方式指定使用相对路径来寻找`mypackage`这个包。

```go
module moduledemo

go 1.14


require "mypackage" v0.0.0
replace "mypackage" => "../mypackage"
```

举个例子

最后我们再举个例子巩固下上面的内容。

我们现在有文件目录结构如下：

```bash
├── p1
│   ├── go.mod
│   └── main.go
└── p2
    ├── go.mod
    └── p2.go
```

`p1/main.go`中想要导入`p2.go`中定义的函数。

`p2/go.mod`内容如下：

```go
module liwenzhou.com/q1mi/p2

go 1.14
```

`p1/main.go`中按如下方式导入

```go
import (
	"fmt"
	"liwenzhou.com/q1mi/p2"
)
func main() {
	p2.New()
	fmt.Println("main")
}
```

因为我并没有把`liwenzhou.com/q1mi/p2`这个包上传到`liwenzhou.com`这个网站，我们只是想导入本地的包，这个时候就需要用到`replace`这个指令了。

`p1/go.mod`内容如下：

```go
module github.com/q1mi/p1

go 1.14


require "liwenzhou.com/q1mi/p2" v0.0.0
replace "liwenzhou.com/q1mi/p2" => "../p2"
```