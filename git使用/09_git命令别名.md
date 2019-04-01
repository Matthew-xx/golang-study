# git别名

可以把常用的git命令通过配置文件设置别名。

    git config --global alias.co checkout
    git config --global alias.br branch
    git config --global alias.ci commit
    git config --global alias.st status

设置完后git status 可以写成 git st

## 让一些命令更易理解

### 撤销暂存

> 把文件从暂存区撤回到未暂存状态

        git reset HEAD <file>

可以直接修改此命令为unstage

    git config --global alias.unstage 'reset HEAD'

### 显示最后一条提交日志

原命令

    git log -1
配置

    git config --global alias.last 'log -1'
于是，可以使用命令

    git last

### 让日志已彩色格式化输出（装逼必备而且实用！）

    git config --global alias.lg "log --color --graph --pretty=format:'%Cred%h%Creset -%C(auto)%d%Creset %s %Cgreen(%cr) %C(bold blue)<%an>%Creset' --abbrev-commit"

## 修改原理

以上所有修改都是通过修改配置文件来实现的。

1. --global意味着全局修改，对整个用户起作用。去掉只对当前仓库有影响
2. 全局修改的配置文件路径在。C:\Users\\[username]\\.gitconfig
3. 仓库范围修改配置文件在仓库目录下的\\.git\config
4. 可以直接修改配置文件达到上面的效果。