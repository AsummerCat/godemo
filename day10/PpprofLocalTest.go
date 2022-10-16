package main

import (
	"fmt"
	"os"
	"runtime/pprof"
)

/*
*
内置资源采集工具
比如CPU资源等

runtime/pprof 采集工具型应用运行数据进行分析  (本地使用)
net/http/pprof 采集服务型应用运行数据进行分析 (线上服务使用)

默认10ms进行采集一次堆栈信息,CPU和内存资源
*/
func main() {
	runtimeTest()
}

/*
本地数据分析
*/
func runtimeTest() {
	file, err := os.Create("./CPU.pprof")
	if err != nil {
		fmt.Println("创建采集文件失败")
	}
	//1.开始采集CPU信息
	pprof.StartCPUProfile(file)

	//2.停止采集CPU信息
	defer pprof.StopCPUProfile()

	//3.堆栈信息也写入文件中
	pprof.WriteHeapProfile(file)

	file.Close()

}
