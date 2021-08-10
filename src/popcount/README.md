# 种群计数
- 返回一个数字中被置位的个数，即在一个`uint64`的值中，值为`1`的个数
## v1 基本算法1
## v2 基本算法2
- 末位与`1`做与运算，计入结果，高位都变为`0`的时候即可停止迭代
## v3
- 表达式`x = x | (x + 1)`把x最右边的`0`位翻转为`1`位，直到所有的比特位都变成`1`为止（`x`值为`-1`）。这时迭代次数`n`就是`0`位的数目，`32-n`就是`1`位的数目。
## v4
- `x&(x-1)`可以清除`x`最右边的非零位，统计`x`变为`0`时经过变换的次数，即为`x`中`1`的位数
## v5
- 使用`init`函数来针对每一个可能的`8`位值预计算一个结果表`pc`，这样只需要将`8`个快查表的结果相加而不用进行`64`步的计算
- 