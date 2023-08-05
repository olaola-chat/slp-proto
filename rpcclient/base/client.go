package base

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"os"
	"strings"
	"time"

	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
)

// IsDev 表明当前系统是不是dev
// 不要在init初始化函数中使用
var IsDev bool = false
var RunMode string = "prod"

func init() {
	//通过机器来检测是不是alpha，这样来统一所有的配置和部署
	mode := g.Cfg().GetString("server.RunMode")
	alphaHosts := g.Cfg().GetStrings("server.AlphaHosts")
	if alphaHosts != nil && mode == "prod" {
		host, err := os.Hostname()
		if err == nil {
			host = strings.ToLower(host)
			for _, alpha := range alphaHosts {
				if host == strings.ToLower(alpha) {
					//是alpha服务器
					mode = "alpha"
				}
			}
		}
	}

	IsDev = mode == "dev" || len(mode) == 0
	RunMode = mode

	fmt.Println("server run with", RunMode)
}

// NewClient 根据服务名字和序列化方式创建rpc client
func NewClient(serviceName string, serializeType protocol.SerializeType) (client.XClient, error) {
	discovery, err := NewClientDiscover(serviceName)
	if err != nil {
		return nil, err
	}
	option := client.DefaultOption
	//隔离alpha和prod的访问
	mode := RunMode
	if len(mode) > 0 {
		option.Group = mode
	}
	option.Heartbeat = true
	option.HeartbeatInterval = time.Second
	option.MaxWaitForHeartbeat = time.Second * 1
	option.ConnectTimeout = time.Millisecond * 100 //设定建立连接超时时间为100ms
	option.IdleTimeout = time.Second * 10          // 10秒空闲
	option.SerializeType = serializeType
	xclient := client.NewXClient(serviceName, client.Failover, client.RandomSelect, discovery, option)
	pc := client.NewPluginContainer()
	pc.Add(&client.OpenTracingPlugin{})
	xclient.SetPlugins(pc)
	return xclient, nil
}

// NewClientUseJSON 根据服务名字创建rpc json client
func NewClientUseJSON(serviceName string) (client.XClient, error) {
	return NewClient(serviceName, protocol.JSON)
}

// NewClientUsePb 根据服务名字创建rpc pb client
func NewClientUsePb(serviceName string) (client.XClient, error) {
	return NewClient(serviceName, protocol.ProtoBuffer)
}

// NewClientUseMp 根据服务名字创建rpc msgpack client
func NewClientUseMp(serviceName string) (client.XClient, error) {
	return NewClient(serviceName, protocol.MsgPack)
}
