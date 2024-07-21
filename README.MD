# qssh

`qssh` 是一个用于快速连接和管理 SSH 主机的命令行工具，支持批量 SSH 批量脚本执行。

## 功能

- 快速连接到 SSH 主机
- 管理 SSH 主机配置
- 批量执行命令和脚本

## 安装

### 从源代码构建

要从源代码构建 `qssh`，请确保您已安装 Go 语言环境。

```sh
git clone https://github.com/aiyijing/qssh.git
cd qssh
go build -o qssh
```

## Example

### 连接到 SSH 主机


```sh
./qssh ssh root@192.168.1.1 -p 22 -k ~/.ssh/id_rsa
```
### 批量执行命令

```sh
./qssh add root@192.168.1.1  -k /root/.ssh/id_rsa
./qssh add root@192.168.1.2  -k /root/.ssh/id_rsa
./qssh run "uname -r"
[0] 192.168.88.101
5.15.0-78-generic

[1] 192.168.88.119
5.15.0-78-generic
```