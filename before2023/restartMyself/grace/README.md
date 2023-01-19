使用以下命令杀死进程，进程会被从新拉活
```
kill  -SIGUSR2 15275
```

# 后续思考和实验
* 平滑重启是否能完成平滑更新？
* 平滑重启过程中是否会丢请求？丢请求的比例和概率是多少？

# 参考资料
## Golang学习--平滑重启
https://www.cnblogs.com/CraryPrimitiveMan/p/8560839.html

<img src='./grace.jpeg'/>
