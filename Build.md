### go linux 环境安装 (ubuntu)

```
sudo add-apt-repository ppa:longsleep/golang-backports
sudo apt update
sudo apt install golang-go

```

### 编译

```

## 下载依赖包到vendor
go mod vendor
## 指定依赖包进行编译
go build -mod=vendor
## 编译生成到指定目录
go build -mod=vendor -o build/linux-amd64

```

### plugin

## 实现方法

```

func Map(input *monstachemap.MapperPluginInput) (output *monstachemap.MapperPluginOutput, err error)

func Filter(input *monstachemap.MapperPluginInput) (keep bool, err error)

func Pipeline(ns string, changeStream bool) (stages []interface{}, err error)

func Process(input*monstachemap.ProcessPluginInput)  (err error)


```

## 编译

```

go build -buildmode=plugin -mod=vendor -o plugin/plugin.so plugin/plugin.go


```
