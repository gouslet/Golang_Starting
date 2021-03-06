# du4计算目录中文件占用的磁盘空间大小

du中大多数的工作由walkDir函数来完成，它使用dirent辅助函数来枚举目录中的条目

ioutil.ReadDir函数返回一个os.FileInfo类型的slice，针对单个文件同样的信息可以通过调用os.Stat函数获得

对于每一个子目录，walkDir递归调用自己，对于每一个文件，walkDir发送一条消息到fileSizes通道，消息为文件所占字节数

main函数使用两个goroutine，
- 后台goroutine调用walkDir遍历命令行上指定的每一个目录，最后关闭fileSizes通道
- 主goroutine计算从通道中接收的文件的大小的和，最后输出总数

主goroutine使用一个计时器每500ms定时产生事件，使用select等待一个关于文件大小的消息和定时器消息，接收到文件大小信息时更新数据，接收到定时器消息时输出当前数据

如果没有提供-v=true选项，则tick通道为nil，此时在select中相当于被禁用

为每个walkDir创建一个新的goroutine，使用sync.WaitGroup来为当前存活的walkDir计数，当计数器为0时，某个goroutine关闭fileSizes通道

为dirents函数使用计数信号量，以防止其打开太多文件

可以为程序添加取消机制
  - 创建一个取消通道，当其关闭时代表程序需要停止正在做的事情
  - 定义一个canceled函数，在其被调用的时候检测或轮询取消状态