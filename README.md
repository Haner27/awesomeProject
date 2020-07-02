1.安装protobuf:
cd /opt
rm -rf protobuf-all-3.12.3.tar.gz
wget https://github.com/protocolbuffers/protobuf/releases/download/v3.12.3/protobuf-all-3.12.3.tar.gz
tar -xvf protobuf-all-3.12.3.tar.gz
cd protobuf-3.12.3
./configure --prefix=/usr/local/protobuf
make&make install
设置全局变量

2.安装proto-gen-rpcx插件：
go get github.com/rpcxio/protoc-gen-gogorpcx
编译其中一个插件放在$GOPATH/bin下

3.生成rpcx服务依赖代码
protoc --gofast_out=plugins=rpcx:. proto/helloword/helloword.proto

4.启动依赖服务etcd, zipkin
工程下执行脚本(前提是安装docker):
scripts/start/etcd.sh
scripts/start/zipkin.sh
// zipkin启动后打开http://127.0.0.1:9411/zipkin/traces/5bd824d00f49768b查看链路调用信息

5.run:
命令行分别启动2个服务
go run main.go
go run main2.go
执行一个client测试调用结果
go run client.go


