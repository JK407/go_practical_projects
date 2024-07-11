package main

import (
	"bytes"
	"encoding/binary"
	"flag"
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

var (
	timeout      int64                  //  超时时间
	size         int                    //  发送缓冲区大小
	count        int                    //  要发送的回显请求数
	typ          uint8              = 8 //回送请求8
	code         uint8              = 0 //  回送请求0
	sendCount    int                    //  发送次数
	successCount int                    //  成功次数
	failCount    int                    //  失败次数
	tsList       = make([]int64, 0)     //  时间列表
)

// ICMP
// @Description: ICMP
// @Author Oberl-Fitzgerald 2024-07-11 15:43:38
type ICMP struct {
	Type       uint8  // Type @Description: 类型
	Code       uint8  // Code @Description: 代码
	Checksum   uint16 // Checksum @Description: 校验和
	Identifier uint16 // Identifier @Description: 标识符
	Sequence   uint16 // Sequence @Description: 序列号
}

func main() {
	//  获取命令行参数
	getCommandArgs()
	//  最后一个参数是ip地址
	desIp := os.Args[len(os.Args)-1]
	//  连接目标ip，icmp协议，超时时间timeout
	conn, err := net.DialTimeout("ip:icmp", desIp, time.Duration(timeout)*time.Millisecond)
	if err != nil {
		log.Fatal(err)
		return
	}
	//  延迟关闭连接
	defer conn.Close()
	fmt.Printf("正在 Ping %s[%s] 具有 %d 字节的数据:\n", desIp, conn.RemoteAddr(), size)
	//  循环发送数据
	for i := 0; i < count; i++ {
		sendCount++
		//  获取当前时间
		t1 := time.Now()
		//  创建一个大小为size的字节数组
		data := make([]byte, size)
		//  创建icmp
		icmp := &ICMP{
			Type: typ,
			Code: code,
			//  校验和
			Checksum: 0,
			//  标识符
			Identifier: 1,
			//  序列号
			Sequence: 1,
		}
		//  创建一个buffer
		var buffer bytes.Buffer
		//  写入icmp,bigEndian是大端序,即高位字节存放在内存的低地址;littleEndian是小端序,即高位字节存放在内存的高地址
		binary.Write(&buffer, binary.BigEndian, icmp)
		//  写入data
		buffer.Write(data)
		//  将buffer转换为字节数组
		icmpData := buffer.Bytes()
		//  计算校验和
		checkSum := CheckSum(icmpData)
		//  将校验和写入icmpData,icmpData[2]是高位，icmpData[3]是低位
		icmpData[2] = byte(checkSum >> 8)
		icmpData[3] = byte(checkSum)
		//  设置超时时间
		conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Millisecond))
		//  发送数据
		_, err = conn.Write(icmpData)
		if err != nil {
			failCount++
			log.Fatal(err)
		}
		//  创建一个大小为1024的字节数组接收数据
		buf := make([]byte, 1024)
		//  读取数据
		n, err := conn.Read(buf)
		if err != nil {
			failCount++
			log.Fatal(err)
		}
		successCount++
		//  获取时间差
		ts := time.Since(t1).Milliseconds()
		//  添加到时间列表
		tsList = append(tsList, ts)
		fmt.Printf("来自 %d.%d.%d.%d 的回复: 字节=%d 时间=%dms TTL=%d\n", buf[12], buf[13], buf[14], buf[15], n-28, ts, buf[8])
		time.Sleep(1 * time.Second)
	}
	fmt.Printf("%s 的 Ping 统计信息:\n", conn.RemoteAddr())
	fmt.Printf("    数据包: 已发送 = %d，已接收 = %d，丢失 = %d（%d%% 丢失），\n", sendCount, successCount, failCount, failCount/sendCount*100)
	fmt.Printf("往返行程的估计时间（以毫秒为单位）:\n")
	fmt.Printf("    最短 = %dms，最长 = %dms，平均 = %dms\n", Min(tsList), Max(tsList), Avg(tsList, sendCount))
}

// getCommandArgs
// @Description 获取命令行参数
// @Author Oberl-Fitzgerald 2024-07-11 15:44:22
func getCommandArgs() {
	flag.Int64Var(&timeout, "w", 1000, "等待每次回复的超时时间，单位毫秒")
	flag.IntVar(&size, "l", 32, "发送缓冲区大小，单位字节")
	flag.IntVar(&count, "n", 4, "要发送的回显请求数")
	//  解析命令行参数
	flag.Parse()
}

// CheckSum
// @Description 计算校验和
// @Author Oberl-Fitzgerald 2024-07-11 15:44:29
// @Param  data []byte
// @Return uint16
func CheckSum(data []byte) uint16 {
	var sum uint32
	var length = len(data)
	var index int
	for length > 1 {
		//  将data[index]的高8位和data[index+1]的低8位相加
		sum += uint32(data[index])<<8 + uint32(data[index+1])
		index += 2
		length -= 2
	}
	//  如果length为奇数，则将最后一个字节的高8位和低8位相加
	if length > 0 {
		sum += uint32(data[index])
	}
	//  将sum的高16位和低16位相加
	sum += sum >> 16
	//  取反
	return uint16(^sum)
}

// Max
// @Description 求列表最大值
// @Author Oberl-Fitzgerald 2024-07-11 15:44:55
// @Param  l []int64
// @Return int64
func Max(l []int64) int64 {
	max := l[0]
	for _, v := range l {
		if v > max {
			max = v
		}
	}
	return max
}

// Min
// @Description 求列表最小值
// @Author Oberl-Fitzgerald 2024-07-11 15:45:06
// @Param  l []int64
// @Return int64
func Min(l []int64) int64 {
	min := l[0]
	for _, v := range l {
		if v < min {
			min = v
		}
	}
	return min
}

// Avg
// @Description  求列表平均值
// @Author Oberl-Fitzgerald 2024-07-11 15:45:27
// @Param  l []int64
// @Param  n int
// @Return int64
func Avg(l []int64, n int) int64 {
	var sum int64
	for _, v := range l {
		sum += v
	}
	return sum / int64(n)
}
