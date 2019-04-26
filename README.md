# docker 初始化工具
主要用于 docker-compose 相关初始化，删除不用的网络并生成新网络，并生成 .env 文件。



## 命令使用
**使用 -h 查看帮助**

### init
初始化网络并生成 .env 文件。
#### .env 文件
自动分析 docker-compose.yml 下面的变量，主要包含以下几种：
- `${*_PORT}` 宿主机端口变量，自动搜索系统可用端口（1000-65535），并按照数目生成到 .env 文件
- `${NETWORK_NAME}` 网络名，网络处理生成

#### 网络处理
获取可用端口之后，首先执行：

> $ docker network prune -f

删除所有没有使用的网络，然后生成一个新的网络名以供这个项目使用：

> docker network create -d bridge `${NETWORK_NAME}`

如果获取了可用端口，`${NETWORK_NAME}` 格式为 `docker-*_*`，* 代表端口号；如果不需要在宿主机暴露端口，则不生成网络名，项目启动自行随机生成。