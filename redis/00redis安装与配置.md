# redis的安装与配置

## 补linux基础

### linux下软件安装

1. 源码安装

从官网下载源码压缩包（一般为tar文件），自行编译，可能需要下载编译器，编译时间较长。好处是安装灵活，自由。

安装后，根据个人习惯，需要自己把配置文件，复制到系统合适的位置

2. 软件源安装

通过yum或者aptget类的命令，从提供软件源的站点中，获取想要的已经针对本平台编译好的软件。优点：安装简单方便快捷，缺点是安装受限，可选项少。

    apt-get install redis-server

安装完成后，安装文件会根据包管理者的配置散放在系统各处。通常按照linux系统文件夹的定义安排

### linux常见文件夹

文档一般在 /usr/share
可执行文件 /usr/bin
配置文件 /etc
lib文件 /usr/lib

## ubuntu下redis安装

使用软件源的方式进行安装

### 安装

使用以下命令开始获取软件：

    sudo apt-get update
    sudo apt-get install redis-server

安装成功后，默认以服务形式自动开启

### 本地连接redis

    redis-cli

### 启动与关闭

#### apt安装的启动与关闭

使用apt安装的redis会自动创建linux服务,故可以使用以下方式启动与关闭，建议。
特点是不会被kill杀死进程

启动

    service redis-server start [ 配置文件路径：通常为/etc/redis/redis.conf ]

关闭

    service redis-server stop

#### 源码安装的启动与关闭

启动： 找到可执行文件直接运行

    ./redis-server [ 配置文件路径：通常为/etc/redis/redis.conf ]
此种运行方式的特点是，可以被kill杀死进程。

关闭：此种方式只要杀死进程即可关闭

    kill -9 (pid)
    进入redis-cli后或者通过客户端传递关闭命令：
    shutdown
注意以上两种方式，只有在redis以进程形式运行时才有效 。若在redis以服务运行的情况使用，结果会让进程以新pid继续运行。


#### 守护进程的启动与关闭

这是不同于以上两个启动关闭方式的特定：它定义启动之后是否转入后台运行

需要注意：可以修改配置文件daemonize 为yes,以守护进程的方式运行，防止开启之后，不会卡死在当前的命令行界面。如果关闭，此进程会的生命周期绑定在当前会话窗口上。

### redis配置

#### 配置bind

bind是指安装redis的主机，允许通过该主机连接的的网络端口列表。
一般情况下：除了配置127.0.0.1之外，还应配置本机网卡的ip

#### 配置主从，集群

后面具体讲。