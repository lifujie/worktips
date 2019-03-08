# 下载镜像
- 镜像VBoxGuestAdditions_5.2.12.iso 
- 下载地址 http://download.virtualbox.org/virtualbox
# 添加镜像
- virtualbox管理器->机器列表中的具体机器(比如centos)->设置->存储(添加镜像)
# 安装镜像
- 挂载镜像
  - mkdir /home/cdrom
  - mount -t auto /dev/cdrom /home/cdrom
  > 这样就可以在/home/cdrom目录下看到VBoxGuestAdditions_5.2.12.iso 文件的内容了
- 执行
  - sh ./VBoxLinuxAdditions.run
  - reboot
# 挂载共享盘
- mount -t vboxsf go /root/go
# 问题
- /sbin/mount.vboxsf: mounting failed with the error: No such Device
  - 解决方案
    - yum update
    - yum install kernel
    - yum install kernel-devel
    - yum install kernel-headers
    - reboot
    - cd /opt/VBoxGuestAdditions-*/init 
    - ./vboxadd setup
    - reboot 
