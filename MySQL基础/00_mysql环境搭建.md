# mysql环境搭建

## ubuntu使用apt安装

MySQL的安装非常简单，更新包索引，安装mysql-server包，最后运行安全设置脚本。

    sudo apt update
    sudo apt install mysql-server
    sudo mysql_secure_installation

本文将介绍如何在Ubuntu18.04上安装MySQL 5.7版本。但是，如果你想从一个已经存在的版本升级到5.7则不适用本文

Ubuntu 18.04， 在APT包库中，默认情况下只有最新版的MySQL。本文写作是MySQL的版本是5.7。

### 开始安装

为了安装MySQL，需要在你的服务器上通过apt更新包索引：

    sudo apt update

然后按照默认软件包:

    sudo apt install mysql-server

执行上面命令将会安装MySQL，但并不会让你设置密码或者做任何其它配置。因为，这样会使你的安装不安全，我们将在下一步解决该问题。

### 配置MySQL

对于全新安装，你应该运行包含的安全脚本。该脚本改变一些诸如远程root登录和简单用户等不安全的缺省选项。在老版本的MySQL中，你还需要手动初始化数据目录，但现在将被自动完成。

运行安全脚本：

    sudo mysql_secure_installation

该脚本将通过一系列的提示帮你完成MySQL安装安全选项的变更。第一个提示将询问你是否愿意安装密码检测插件，该插件用来测试你设置的MySQL密码的强壮性。无论你如何选择，下一个提示是让你设置MySQL root用户的密码。回车，然后需要确认你输入的密码。

从这开始，后续所有问题可以输入Y或者回车，采用默认配置即可。这将移除一些匿名用户和测试数据库，并且禁用远程root登录。同时，将加载这些新规则以使您做的变更能够在MySQL立刻生效。

初始化MySQL数据目录，在5.7.6之前的版本需要使用mysql_install_db， 5.7.6及之后的版本使用mysqld --initialize进行初始化。如果您通过步骤1描述的Debian包安装的MySQL，数据目录将被自动初始化，您不需要做任何事情。如果你试着运行这个命令，您将看到如下错误提示信息：

    mysqld: Can't create directory '/var/lib/mysql/' (Errcode: 17 - File exists). . .2018-04-23T13:48:00.572066Z 0 [ERROR] Aborting

需要注意的是，虽然你设置了MySQL服务root用户的密码，但当通过MySQL终端登录时并不能通过密码认证登录。如果您愿意，可以通过步骤3进行设置。

### 调整用户认证和权限（步骤3，可选）

在Ubuntu系统中MySQL 5.7及之后的版本，MySQL的root用户被默认设置成通过auth_socket插件进行认证，而不是通过密码。在很多情况下，这些配置可以使系统更加的安全和可靠，但如果允许外部程序（例如phpMyAdmin）访问时，这将是事情变得非常复杂。

为了能够以root用户通过密码的方式连接MySQL，你需要将其认证方式从 auth_socket 方式变更为mysql_native_password。进行该设置，通过终端打开MySQL的提示符：

    sudo mysql

下一步，通过如下命令检查您的MySQL系统每个用户的认证方式：

    SELECT user,authentication_string,plugin,host FROM mysql.user;
    
    Output
    
    +------------------+-------------------------------------------+-----------------------+-----------+| user | authentication_string | plugin | host |+------------------+-------------------------------------------+-----------------------+-----------+| root | | auth_socket | localhost || mysql.session | *THISISNOTAVALIDPASSWORDTHATCANBEUSEDHERE | mysql_native_password | localhost || mysql.sys | *THISISNOTAVALIDPASSWORDTHATCANBEUSEDHERE | mysql_native_password | localhost || debian-sys-maint | *CC744277A401A7D25BE1CA89AFF17BF607F876FF | mysql_native_password | localhost |+------------------+-------------------------------------------+-----------------------+-----------+4 rows in set (0.00 sec)

本例中，您可以看到实际上root用户通过auth_socket插件的方式进行认证。要将root用户设置为通过密码认证，运行如下ALTER USER命令。务必将密码设置为高强度的密码，需要注意的是该操作将改变您在步骤2中设置的密码：

    ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'password';

然后运行 FLUSH PRIVILEGES， 这将让数据库系统重新加载授权表是变更生效：

    FLUSH PRIVILEGES;

检查每个用户的授权方法，确认root用户不再使用auth_socket插件进行认证。

    SELECT user,authentication_string,plugin,host FROM mysql.user;
    
    Output
    
    +------------------+-------------------------------------------+-----------------------+-----------+| user | authentication_string | plugin | host |+------------------+-------------------------------------------+-----------------------+-----------+| root | *3636DACC8616D997782ADD0839F92C1571D6D78F | mysql_native_password | localhost || mysql.session | *THISISNOTAVALIDPASSWORDTHATCANBEUSEDHERE | mysql_native_password | localhost || mysql.sys | *THISISNOTAVALIDPASSWORDTHATCANBEUSEDHERE | mysql_native_password | localhost || debian-sys-maint | *CC744277A401A7D25BE1CA89AFF17BF607F876FF | mysql_native_password | localhost |+------------------+-------------------------------------------+-----------------------+-----------+4 rows in set (0.00 sec)

在本例中，您可以看到MySQL的root用户现在是通过密码的方式进行认证。一旦确认服务器上配置挣钱，你可以退出MySQL终端：

    exit

或者，有些人可能会发现，使用专用用户连接到MySQL更适合他们的工作流。要创建这样的用户，请再次打开mysql 终端：

    sudo mysql

注意： 如果您启用的root用户的密码认证，您需要通过不同的命令登录MySQL终端。如上所述，将以常规用户权限运行MySQL客户端。只能通过认证以获得管理员权限。

    mysql -u root -p

如下，创建一个新用户，并设置强密码：

    CREATE USER 'sammy'@'localhost' IDENTIFIED BY 'password';

然后，授予新用户合适的权限。例如，授予新用户访问数据库中所有表的权限，及添加、变更和移除用户的权限，通过如下命令即可：

    GRANT ALL PRIVILEGES ON *.* TO 'sammy'@'localhost' WITH GRANT OPTION;

需要注意的是，这时您不再需要运行FLUSH PRIVILEGES命令。只有通过 INSERT, UPDATE或者DELETE命令的方式变更授权表的时候才需要该命令。由于您创建了一个新用户，而不是改变一个已经存在的用户，因此FLUSH PRIVILEGES并不是必须要运行的。

退出MySQL终端：

    exit

## 无法远程连接的解决

### 添加1个用户专门用户远程连接

     grant all privileges on *.* to beego@'%' identified by '123456';

### 关闭防火墙

    sudo ufw disable

### 修改配置文件的bind

一般在etc/mysql文件夹下

my.cnf里面没有bind-address项，有可能在这里：
    
    /etc/mysql/mysql.conf.d/mysqld.cnf