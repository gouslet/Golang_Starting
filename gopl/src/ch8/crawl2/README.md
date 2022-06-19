# 用信号量限制并发

无限制的并发通常并不是好主意，因为系统中总有限制因素，例如，对于计算型应用的CPU的核数，对于磁盘I/O操作磁头和磁盘的个数，下载流所使用的网络带宽，或者web服务本身的容量

解决办法是根据资源可用情况限制并发的个数，以匹配合适的并发度


可以使用容量为n的缓冲通道来建立一个并发原语，称为计数信号量

对于缓冲通道中的n个空闲位置，每个代表一个令牌，持有者可以执行

通过发送一个值到通道中来领取令牌，从通道中接收一个值来释放令牌，创建一个新的空闲位置


这保证了在没有接收操作的时候，最多同时有n个发送

crawl2中使用struct{}类型的通道，它所占用的空间为0。并使用令牌的获取和释放操作来包括对links.Extract函数的调用，这样保证最多同时20个调用可以进行

计数器n跟踪发送到人物列表中的任务数量，
- 每次一个条目发送到人物列表时，递增变量n，第一次递增是在发送初始化命令行参数之前，第二次递增是在每次启动一个新的爬取goroutine时
- 当主循环中n减少到0时，再没有任务需要完成
