# ubunut+ssqt5

## 生成PAC文件

默认qt5开启后是全局代理

### 安装pip

    apt-get install python-pip

### 安装GenPAC

GenPAC 是基于gfwlist的代理自动配置（Proxy Auto-config）文件生成工具，支持自定义规则。

    sudo pip install genpac
    pip install --upgrade genpac

genpac 的详细使用说明见 GitHub - Wiki：https://github.com/JinnLynn/GenPAC

### 下载gfwlist生成pac文件 

此处生成的文件名为autoproxy.pac

    genpac --pac-proxy "SOCKS5 127.0.0.1:1080" --gfwlist-proxy="SOCKS5 127.0.0.1:1080" --gfwlist-	url=https://raw.githubusercontent.com/gfwlist/gfwlist/master/gfwlist.txt --output="autoproxy.pac"


四、配置PAC文件
上述两种方式生成的.pac文件的内容其实是一样的
点击：System settings > Network > Network Proxy，选择 Method 为 Automatic，设置 Configuration URL 为 autoproxy.pac 文件的路径，点击 Apply System Wide。
格式如：file:///home/{user}/autoproxy.pac


## 或者使用在线pac

    https://raw.githubusercontent.com/petronny/gfwlist2pac/master/gfwlist.pac


## bash shell走socks5代理

bash上配置代理是export两个环境变量

    export http_proxy=socks5://127.0.0.1:1080
    export https_proxy=socks5://127.0.0.1:1080

    注意这时候从这个bash启动的所有软件都是自带这两个环境变量的。如果需要vscode使用代理，直接使用code命令启动即可

## 某软件单独走代理如：VSCODE

    code --https_proxy=socks5://127.0.0.1:1080 --http_proxy=socks5://127.0.0.1:1080
    此方法没有奏效
