# 关闭所有相关进程
- ps aux|grep epmd
- ps aux|grep erl
> 否则在启动多节点的时候会出现端口占用情况
# 关闭plugins ，否则会抢占15672端口导致其他节点无法启动
- rabbitmq-plugins disable rabbitmq_management
# 关闭节点
- rabbitmqctl stop
# 启动多节点
- RABBITMQ_NODE_PORT=5672 RABBITMQ_NODENAME=rabbit  rabbitmq-server -detached
# 启动其他节点
- RABBITMQ_NODE_PORT=5673 RABBITMQ_NODENAME=rabbit_1  rabbitmq-server -detached
- RABBITMQ_NODE_PORT=5674 RABBITMQ_NODENAME=rabbit_2  rabbitmq-server -detached
# 多节点加入集群
- sudo rabbitmqctl -n rabbit_1@localhost stop_app
- sudo rabbitmqctl -n rabbit_1@localhost reset
- sudo rabbitmqctl -n rabbit_1@localhost join_cluster rabbit@localhost
- sudo rabbitmqctl -n rabbit_1@localhost start_app
- sudo rabbitmqctl -n rabbit_1@localhost cluster_status
- sudo rabbitmqctl -n rabbit_2@localhost stop_app
- sudo rabbitmqctl -n rabbit_2@localhost reset
- sudo rabbitmqctl -n rabbit_2@localhost join_cluster rabbit@localhost
- sudo rabbitmqctl -n rabbit_2@localhost start_app
- sudo rabbitmqctl -n rabbit_2@localhost cluster_status
# 启动页面管理程序
- rabbitmq-plugins enable rabbitmq_management
# 添加用户&设置角色&设置操作权限
- rabbitmqctl add_user admin admin
- rabbitmqctl set_user_tags admin administrator
- rabbitmqctl set_permissions -p / admin ".*" ".*" ".*"