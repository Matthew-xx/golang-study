itcast@itcast-virtual-machine:~$ brctl help
never heard of command [help]
Usage: brctl [commands]
commands:
	addbr				<bridge>					add bridge
													添加网桥
	delbr				<bridge>					delete bridge
													删除网桥
	addif				<bridge> <device>			add interface to bridge
													增加网桥接口
	delif				<bridge> <device>			delete interface from bridge
													删除网桥接口
	hairpin				<bridge> <port> {on|off}	turn hairpin on/off
													打开或者关闭回传
	setageing			<bridge> <time>				set ageing time
													设置老化时间（时效）
	setbridgeprio		<bridge> <prio>				set bridge priority
													设置网桥优先级
	setfd				<bridge> <time>				set bridge forward delay
													设置桥转发延迟
	sethello			<bridge> <time>				set hello time
													设置时间间隔
	setmaxage			<bridge> <time>				set max message age
													设置该虚拟桥接口的报文最大周期
	setpathcost			<bridge> <port> <cost>		set path cost
													设置该虚拟桥接口的的端口成本
	setportprio			<bridge> <port> <prio>		set port priority
													设置该虚拟桥接口的端口优先级
	show				[ <bridge> ]				show a list of bridges
													显示网桥列表
	showmacs			<bridge>					show a list of mac addrs
													显示mac地址列表
	showstp				<bridge>					show bridge stp info
													显示网桥stp信息
	stp					<bridge> {on|off}			turn stp on/off
													打开或关闭stp（生成树协议）
