package consul

import (
	"fmt"
	"gin-use/configs"
	"gin-use/src/global"
	"log"
	"math/rand"
	"strconv"

	consulapi "github.com/hashicorp/consul/api"
	"go.uber.org/zap"
)

var (
	//consul地址
	ConsulAddr = configs.ConsulAddr()
	//consul服务id 用项目名+ip
	registrationId = fmt.Sprintf("%s:%s", configs.ProjectName(), configs.GetLocalIp()[0])
	//consul健康检查
	apiHealth = fmt.Sprintf("%s:%s/api/health", configs.ProjectHost(), configs.ProjectPort())
)

// 注册服务到consul
func Register() {
	// 创建连接consul服务配置
	config := consulapi.DefaultConfig()
	config.Address = ConsulAddr
	client, err := consulapi.NewClient(config)
	if err != nil {
		global.Logger.Error("consul client error : ", zap.Any("err", err))
	}

	// 创建注册到consul的服务到
	registration := new(consulapi.AgentServiceRegistration)
	//consul服务id 用项目名+ip后3个数
	registration.ID = registrationId
	registration.Name = configs.ProjectName()
	registration.Port, _ = strconv.Atoi(configs.ProjectPort())
	registration.Tags = []string{""}
	registration.Address = configs.GetLocalIp()[0] //ip

	// 增加consul健康检查回调函数
	check := new(consulapi.AgentServiceCheck)
	check.HTTP = apiHealth // 健康检查地址
	check.Timeout = "5s"
	check.Interval = "5s"                        // 健康检查间隔
	check.DeregisterCriticalServiceAfter = "30s" // 故障检查失败30s后 consul自动将注册服务删除
	registration.Check = check

	// 注册服务到consul
	err = client.Agent().ServiceRegister(registration)
}

// 取消consul注册的服务
func DeRegister() {
	// 创建连接consul服务配置
	config := consulapi.DefaultConfig()
	config.Address = ConsulAddr
	client, err := consulapi.NewClient(config)
	if err != nil {
		global.Logger.Error("consul client error : ", zap.Any("err", err))
	}
	client.Agent().ServiceDeregister(registrationId)
}

/**
 * @summary 从consul中发现服务
 * @param serviceName 服务名
 * @param serviceTag 服务标签
 */
func FindServer(serviceName string, serviceTag string) string {
	var lastIndex uint64
	// 创建连接consul服务配置
	config := consulapi.DefaultConfig()
	config.Address = ConsulAddr
	client, err := consulapi.NewClient(config)
	if err != nil {
		global.Logger.Error("consul client error : ", zap.Any("err", err))
	}
	// 获取指定service
	services, metainfo, err := client.Health().Service(serviceName, serviceTag, false, &consulapi.QueryOptions{
		WaitIndex: lastIndex, // 同步点，这个调用将一直阻塞，直到有新的更新
	})

	if err != nil {
		global.Logger.Error("error retrieving instances from Consul: %v", zap.Any("err", err))
	}
	lastIndex = metainfo.LastIndex

	var AddressList []string = []string{}

	for _, service := range services {
		AddressList = append(AddressList, fmt.Sprintf("%s:%d", service.Service.Address, service.Service.Port))
	}

	//有可能多个ip,负载均衡随机返回其中一个
	return AddressList[rand.Intn(len(AddressList))]
}

//服务健康检查
func CheckHeath() {
	// 创建连接consul服务配置
	config := consulapi.DefaultConfig()
	config.Address = ConsulAddr
	client, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatal("consul client error : ", zap.Any("err", err))
	}
	// 健康检查
	a, b, _ := client.Agent().AgentHealthServiceByID(registrationId)
	fmt.Println("健康检查:", a, b)
}

//返回http地址
func FindHttpServer(serviceName string, serviceTag string) string {
	ip := FindServer(serviceName, serviceTag)
	return fmt.Sprintf("http://%s", ip)
}
