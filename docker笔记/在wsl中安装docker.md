# 在WSL中安装docker（支持守护进程daemon）

## WSL介绍

Linux子系统(Windows Subsystem for Linux ,WSL)不是虚拟机，只是一个模拟环境，因而并不能完全实现Linux中的各种操作，而且性能也不如一般的虚拟机。得益于此，WSL与windows交互非常方便，消耗非常低，没有启动时间，用来做轻量开发或测试很适合。

网上关于如何在WSL中安装Docker有很多教程，大多时间较早，都面临着需要安装Docker for Windows的问题，这就需要运行虚拟机，对于我资源吃紧的笔记本还是有些负担，而且对于没有Win10 Pro（Hyper-V）的用户也不方便安装。Win10创意者更新后，在WSL中运行Docker Engine成为可能，当然这个功能正在逐步完善中，当前并不能支持docker全部指令。

本机环境
Win10 Pro 1903


## 安装WSL

Microsoft Store==>搜索Linux==>安装Ubuntu18.04 

## WSL中安装Docker

启动WSL控制台，执行以下指令

    sudo apt update
    sudo apt install docker.io
    sudo usermod -aG docker $USER
随后再以管理员启动WSL控制台，执行

    sudo cgroupfs-mount
    sudo service docker start
    sudo chmod -R 777 /var/run/docker.sock

## 测试安装结果

    $ docker version

    Client: Docker Engine - Community
    Version:           19.03.1
    API version:       1.40
    Go version:        go1.12.5
    Git commit:        74b1e89
    Built:             Thu Jul 25 21:21:05 2019
    OS/Arch:           linux/amd64
    Experimental:      false

    Server: Docker Engine - Community
    Engine:
    Version:          19.03.1
    API version:      1.40 (minimum version 1.12)
    Go version:       go1.12.5
    Git commit:       74b1e89
    Built:            Thu Jul 25 21:19:41 2019
    OS/Arch:          linux/amd64
    Experimental:     false
    containerd:
    Version:          1.2.6
    GitCommit:        894b81a4b802e4eb2a91d1ce216b8817763c29fb
    runc:
    Version:          1.0.0-rc8
    GitCommit:        425e105d5a03fabd737a126ad93d62a9eeede87f
    docker-init:
    Version:          0.18.0
    GitCommit:        fec3683

显示如上信息，就表示可以正常使用了，需要注意的是每次电脑重启后先执行cgroupfs-mount再启动docker服务，目前并不完善，期待巨硬能做的更好。

## 启动docker

每次打开wsl需要重新启动

    sudo apt update
    sudo apt install docker.io

如果不想每次都书这个命令，可自己写个shell脚本，并设置，管理员权限开机启动
