package depends

import (
	"context"
	"time"
)

type ServiceAData struct {
	Message string
	Query string
}
type ServiceA struct {
	data ServiceAData
}

func NewServiceA() *ServiceA {
	return &ServiceA{}
}

func (s *ServiceA) Name() string {
	return "service_a"
}

func (s * ServiceA) Run(ctx context.Context, dc *DataContainer) error {
	time.Sleep(10 * time.Millisecond)
	s.data = ServiceAData{
		Message:"I am service a",
		Query: ctx.Value("q").(string),
	}
	return nil
}

func (s * ServiceA) Decode(receiver interface{}) error {
	if _, ok := receiver.(*ServiceAData); ok {
		*(receiver.(*ServiceAData)) = s.data
	}
	return nil
}


type ServiceBData struct {
	Message string
	Query string
}
type ServiceB struct {
	data ServiceBData
}

func NewServiceB() *ServiceB {
	return &ServiceB{}
}

func (s *ServiceB) Name() string {
	return "service_b"
}

func (s * ServiceB) Run(ctx context.Context, dc *DataContainer) error {
	time.Sleep(20 * time.Millisecond)
	s.data = ServiceBData{
		Message:"I am service b",
		Query: ctx.Value("q").(string),
	}
	return nil
}

func (s * ServiceB) Decode(receiver interface{}) error {
	if _, ok := receiver.(*ServiceBData); ok {
		*(receiver.(*ServiceBData)) = s.data
	}
	return nil
}


type ServiceCData struct {
	Message string
	Query string
}
type ServiceC struct {
	data ServiceCData
}

func NewServiceC() *ServiceC {
	return &ServiceC{}
}

func (s *ServiceC) Name() string {
	return "service_c"
}

func (s * ServiceC) Run(ctx context.Context, dc *DataContainer) error {
	time.Sleep(15 * time.Millisecond)
	s.data = ServiceCData{
		Message:"I am service c",
		Query: ctx.Value("q").(string),
	}
	return nil
}

func (s * ServiceC) Decode(receiver interface{}) error {
	if _, ok := receiver.(*ServiceCData); ok {
		*(receiver.(*ServiceCData)) = s.data
	}
	return nil
}