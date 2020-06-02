package util

import (
	"fmt"
	"gopkg.in/ini.v1"
	"os"
)

var ProxyConfigs map[string]string
type EnvConfig *os.File

func init() {
	// 因为保存在缓存里面的，所以修改了env要重启服务器
	ProxyConfigs=make(map[string]string)
	EnvConfig,err:= ini.Load("env")
	if err!=nil{
		fmt.Println(err)
		return
	}
	proxy,_:=EnvConfig.GetSection("proxy") // 假设是固定的分区
	if proxy!=nil{
		secs:=proxy.ChildSections() // 获取子分区
		for _,sec:=range secs{
			path,_:= sec.GetKey("path") // 获取固定key
			pass,_:= sec.GetKey("pass") // 获取固定key
			if path!=nil && pass !=nil{
				ProxyConfigs[path.Value()]=pass.Value()
			}
		}
	}
	//fmt.Println(ProxyConfigs)
}
