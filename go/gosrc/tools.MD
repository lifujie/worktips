# 查看程序入口地址
  - readelf -h  main
    > 入口点地址
  - objdump -f  main 
    > 起始地址
# 根据入口地址找到对应的函数
  - nm main | grep 4498a0
