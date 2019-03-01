#go-package-plantuml

## 修改
+ 默认通过环境变量获取GOPATH
+ 输出目标参数功能修正
+ 兼容Windows路径
+ 输出全部改为英文

## Install

```
go get github.com\Pingze-github\go-package-plantuml
go build github.com\Pingze-github\go-package-plantuml
```

## Usage

Parse a golang package into a .puml file:
```
go-package-plantuml  --gopath $GOPATH --codedir $GOPATH\src\github.com\micro\go-micro --outputfile $GOPATH\src\github.com\micro\go-micro.puml
```

Compile the .puml file into a .svg image:
```
java -jar C:\jars\plantuml.jar -tsvg go-micro.puml
```

(Download plantuml.jar [here](http://plantuml.com/en/download))

You will got this:

[go-micro.svg](https://github.com/Pingze-github/go-package-plantuml/tree/master/images/go-micro.svg)