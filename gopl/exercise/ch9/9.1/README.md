# 9.1
向bank程序中增加一个函数Withdraw(amount int)bool。结果应当反映交易成功还是余额不足而失败。函数发送到监控goroutine的消息应当包含取款金额和一个新的通道，该通道用于监控goroutine把布尔型的结果发送回Withdraw函数