### 关闭首次登陆“Are you sure you want to continue connecting (yes/no)?“
 - 修改ssh配置文件
    1. vi /etc/ssh/ssh_config
    2. 将#StrictHostKeyChecking ask 修改为
       StrictHostKeyChecking no
    3. 重启服务
       service sshd reload
 - 命令行指定
    1. ssh -o StrictHostKeyChecking=no
