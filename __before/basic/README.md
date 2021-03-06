### 基础知识

[CodeJam](https://codingcompetitions.withgoogle.com/codejam)

### 一道题：
* [Beautiful Numbers](https://code.google.com/codejam/contest/5264487/dashboard#s=p1)
    * 1, 11, 111 .... 是 beautiful 的
    * 3 -> 2进制 -> 11
    * 13 -> 3进制 -> 111 (√)
        * 到过来算（111）-> 1*3*3 + 1*3 + 1 = 13 (进位制)
        * 13 % 3 = 1(低位)，13 / 3 = 4
        * 4 % 3 = 1，4 / 3 = 1
        * 1 % 3 = 1，1 / 3 = 0
    * 13 -> 12进制 -> 11 (X)

#### 操作系统

* 进程(隔离)
    * 线程
    * 内存（逻辑内存）
    * 文件/网络句柄
* 线程(真正运行)
    * 栈（调用堆栈）,从主线程的入口 main 函数调用，每次调用把所有的参数和返回地址压到栈里面
    * PC (program counter)：下一条执行指令地址，放在内存
    * TLS (Thread Local Storage)：线程独立的内存
* 存储
    * 寄存器 > 缓存 > 内存 > 硬盘
* 寻址空间
    * 32位 -> 4G
    * 64位 -> ~10^19 Bytes
    * 64位JVM -> 可使用更大内存，需重新编译

寻址 int n = *p; -> MOV EAX,[EBX]

指针p --> 逻辑内存，进程独立 2^32 或 2^64 ---> 物理内存（分页，虚拟内存）---> 寄存器


#### 网络

数据链路层->网络层->传输层->应用层

* 网络传输
    * 不可靠
        * 丢包，重复包
        * 出错
        * 乱序
    * 不安全
        * IS THIS LINE SECURE？
        * 中间人攻击
        * 窃取
        * 篡改

* 例：滑动窗口（增加线路的吞吐量，按照工程化的思维理解）
    * TCP 协议中使用
    * 为维持发送方 / 接收方缓冲区
        * 发送 & Ack
    * 丢 Ack (一定是按照顺序 Ack )
        * 超时重传
* 题：
    [阿里巴巴有相距1500km的机房A和B，现有100GB数据需要通过...](https://www.nowcoder.com/questionTerminal/97a4bed9cf644832bbb8dec72afccfa8)

#### 关系性数据库

* 基于关系代数理论
* 缺点：表结构不直观，实现复杂，速度慢
* 优点：健壮性高，社区庞大