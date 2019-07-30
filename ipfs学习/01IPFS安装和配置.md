# 安装和配置

## 下载

    wget https://dist.ipfs.io/go-ipfs/v0.4.21/go-ipfs_v0.4.21_linux-amd64.tar.gz

## 安装

    $ tar xvfz go-ipfs.tar.gz
    $ cd go-ipfs
    $ ./install.sh

## 创建本地项目节点

    ipfs init
    注意创建的节点是用户级别的，这是与git的不同之处，git是通过init创建文件夹级别的仓库，切换文件夹就切换到不同的仓库。而ipfs通过init初始化出来的.ipfs文件夹只存在于家目录，ipfs切换文件夹不会切换节点，
## 欢迎页

    ipfs cat /ipfs/QmS4ustL54uo8FzR9455qaxZwuMiUhyvMcX9Ba8nUH4uVv/readme
