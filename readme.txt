go的各种包管理，在国内有加速效果

github地址：https://github.com/gpmgo/gopm
官方地址：https://gopm.io/
文档路径：https://github.com/gpmgo/docs/tree/master/zh-CN

// cmd:go get -u github.com/gpmgo/gopm

安装：
go get -v -u github.com/gpmgo/gopm

//如果不行请手动安装
go get github.com/gpmgo/gopm
go install github.com/gpmgo/gopm
go build github.com/gpmgo/gopm

使用：
# 查看当前工程依赖
gopm list
# 显示依赖详细信息
gopm list -v
# 列出文件依赖
gopm list -t [file]
# 拉取依赖到缓存目录
gopm get -r xxx
# 仅下载当前指定的包
gopm get -d xxx
# 拉取依赖到$GOPATH
gopm get -g xxx
# 检查更新所有包
gopm get -u xxx
# 拉取到当前所在目录
gopm get -l xxx
# 运行当前目录程序
gopm run
# 生成当前工程的 gopmfile 文件用于包管理
gopm gen -v
# 根据当前项目 gopmfile 链接依赖并执行 go install
gopm install -v
# 更新当前依赖
gopm update -v
# 清理临时文件
gopm clean
# 编译到当前目录
gopm bin

gopm 下载存放缓存目录 $USER/.gopm/repos
.gopmfile
gopmfile 需放在项目根目录下，名称为 .gopmfile
这个文件可以通过生成
gopm gen

文件格式为
    [target]
    path = github.com/gpmgo/gopm

    [deps]
    github.com/codegangsta/cli = branch:master

    [res]
    include = conf|etc|public|scripts|templates

说明：
target -> path 指示项目名称或导入路径。
deps 节包含了特殊（非最新）版本的依赖。
res 在执行 gopm bin 命令时自动打包的资源。

包版本
有五种可能的包版本组合：
空白：表示使用最新版本的依赖进行构建
/path/to/my/project：绝对或者相对的文件路径，例如：d:\projects\xorm
branch:<value>：固定分支，例如：branch:master
tag:<value>：指定标签，例如：tag:v0.9.0
commit:<value>：某个提交，例如：commit:6ffffe9 一般来说只需要 SHA 的前 7 个字母就可以确定一个提交


NAME:
   Gopm - Go Package Manager

USAGE:
   Gopm [global options] command [command options] [arguments...]

VERSION:
   0.8.8.0307 Beta

COMMANDS:
   list         list all dependencies of current project
   gen          generate a gopmfile for current Go project
   get          fetch remote package(s) and dependencies
   bin          download and link dependencies and build binary
   config       configure gopm settings
   run          link dependencies and go run
   test         link dependencies and go test
   build        link dependencies and go build
   install      link dependencies and go install
   clean        clean all temporary files
   update       check and update gopm resources including itself
   help, h      Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --noterm, -n         disable color output
   --strict, -s         strict mode
   --debug, -d          debug mode
   --help, -h           show help
   --version, -v        print the version

example:
gopm get -g -v -u golang.org/x/tools/cmd/goimports


//自动发现网页编码库:
下载：https://github.com/gpmgo/gopm

gopm get -g -v golang.org/x/text
go install golang.org/x/text
将网页的gbk乱码解决
自动发现网页是gbk编码还是utf8编码
gopm get -g -v golang.org/x/net/html

	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"



"golang.org/x/net/html/charset"
"golang.org/x/text/encoding"
"golang.org/x/text/transform"


//获得城市名称方法
1：使用css选择器
2：使用xpath
3：使用正则表达式来实现


//安装redis
在“终端”进行安装
输入：go get -u github.com/go-redis/redis

go get -v github.com/go-redis/redis


各自的优缺点
encoding/json, 官方自带的, 文档最多, 易用性差, 性能差
go-simplejson, gabs, jason等衍生包, 简单且易用, 易于阅读, 便于维护, 但性能最差
easyjson, ffjson此类包, 适合固定结构的json, 易用性一般, 维护成本高, 性能特别好
jsonparser 适合动态和固定结构的json, 简单且易用, 维护成本低, 性能极好
以性能的高低排序: jsonparser > easyjson > encoding/json > go-simplejson, gabs, jason

性能测试, 可见jsonparser的github
buger/jsonparser