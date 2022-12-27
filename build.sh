#!/usr/bin/env bash
# 编译脚本，直接bash build即可把相关静态扫描规则编译为二进制文件，mac下可直接运行，windows下可在git bash下运行
# 编译成二进制文件后，直接在想要被扫描的项目根目录下运行 xxx ./... 即可，例如 xxx/output/bin/loopgoroutinecheck ./...
go mod vendor

mkdir -p output/bin
for dir_name in $(ls ./cmd)
do
  go build -mod=vendor -o output/bin/"$dir_name" cmd/"$dir_name"/main.go
done
chmod +x output/bin/*