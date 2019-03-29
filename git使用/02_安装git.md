# 安装git

最早Git是在Linux上开发的，很长一段时间内，Git也只能在Linux和Unix系统上跑。不过，慢慢地有人把它移植到了Windows上。现在，Git可以在Linux、Unix、Mac和Windows

## linux下安装

### 判断git是否安装

    git

### Debian系列（含ubuntu系列）安装

    sudo apt-get install git

## 在Mac OS X上安装Git

如果你正在使用Mac做开发，有两种安装Git的方法。
  
1. 安装homebrew，然后通过homebrew安装Git，具体方法请参考homebrew的文档：<http://brew.sh/。>

2. 推荐此方法。就是直接从AppStore安装Xcode，Xcode集成了Git，不过默认没有安装，你需要运行Xcode，选择菜单“Xcode”->“Preferences”，在弹出窗口中找到“Downloads”，选择“Command Line Tools”，点“Install”就可以完成安装了。

## 在Windows上安装Git

在Windows上使用Git，可以从Git官网直接下载安装程序，然后按默认选项安装即可。

安装完成后，在开始菜单里找到“Git”->“Git Bash”，蹦出一个类似命令行窗口的东西，就说明Git安装成功！