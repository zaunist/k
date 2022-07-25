# k

> 当访问不同版本的 k8s 的时候，官方的建议是 kubectl 也要跟着版本走，所以在需要操作多个跨大版本的集群时，切换 kubectl 就比较麻烦，这个工具也是因此诞生。

K 是一个用于管理 kubectl 版本的工具，可以使用它下载对应版本的 kubectl，支持使用 k use 切换至想要的版本。

## 用法

k install --version=v1.23.0  安装 v1.23.0 版本的 kubectl

k use --version=v1.22.0  切换为 v1.22.0 版本的 kubectl

k ls  --列出当前主机上已有的 kubectl 版本

待支持的命令： uninstall、clean、remote
