# QLMessage
work

思路：
1.命令行工具实现使用urfave/cli
2.爬虫任务使用队列+channel+select监听的方法去完成，定位元素使用正则表达式
3.做的是一个并发类型的爬虫，如果继续改进的话，可以使用grpc做分布式
