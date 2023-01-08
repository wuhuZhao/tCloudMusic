package consul

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"

	"github.com/cloudwego/hertz/pkg/app/server/registry"
	"github.com/cloudwego/kitex/pkg/klog"
	consulapi "github.com/hashicorp/consul/api"
)

const consulAddress = "127.0.0.1:8500"

// impl Registry in hertz
var _ registry.Registry = (*ConsulRegistry)(nil)

type ConsulRegistry struct {
}

func (c *ConsulRegistry) Register(info *registry.Info) error {
	tags := []string{}
	for k, v := range info.Tags {
		tags = append(tags, fmt.Sprintf("%s=%s"), k, v)
	}
	return registerService(info.ServiceName, strings.Split(info.Addr.String(), ":")[0], strings.Split(info.Addr.String(), ":")[1], tags)
}

func (c *ConsulRegistry) Deregister(info *registry.Info) error {
	return nil
}

func registerService(serviceName, registraionAddress string, registrationPort string, tags []string) error {
	port, err := strconv.Atoi(registrationPort)
	if err != nil {
		klog.Errorf("convert port to int error: %v\n", err.Error())
		return err
	}
	config := consulapi.DefaultConfig()
	config.Address = consulAddress
	client, err := consulapi.NewClient(config)
	if err != nil {
		klog.Errorf("create consul client error: %v\n", err.Error())
		return err
	}
	registration := &consulapi.AgentServiceRegistration{
		ID:      strconv.Itoa(rand.Intn(math.MaxInt16)),
		Name:    serviceName,
		Port:    port,
		Tags:    tags,
		Address: registraionAddress,
	}
	check := &consulapi.AgentServiceCheck{}
	check.TCP = fmt.Sprintf("%s:%d", registraionAddress, registrationPort)
	check.Timeout = "5s"
	check.Interval = "5s"
	check.DeregisterCriticalServiceAfter = "60s"
	registration.Check = check

	if err := client.Agent().ServiceRegister(registration); err != nil {
		klog.Errorf("register to consul error: %v\n", err.Error())
		return err
	}
	return nil
}
