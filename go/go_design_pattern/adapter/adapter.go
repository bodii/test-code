package adapter

import "fmt"

// ICreateServer create server interface
type ICreateServer interface {
	CreateServer(cup, mem float64) error
}

// AWSClient aws sdk
type AWSClient struct{}

// RunInstance run instance
func (c *AWSClient) RunInstance(cup, mem float64) error {
	fmt.Printf("aws client run success, cpu: %f, mem: %f", cup, mem)
	return nil
}

// AwsClientAdapter aws client adapter
type AwsClientAdapter struct {
	Client AWSClient
}

// CreateServer create aws server
type (a *AwsClientAdapter) CreateServer(cup, mem float64) error {
	a.Client.RunInstance(cup, mem)
	return nil
}

// AliyunClient aliyun sdk
type AliyunClient struct{}

// CreateServer aliyun client
func (a *AliyunClient) CreateServer(cup, mem int) error {
	fmt.Printf("aliyun client run success, cpu: %d, mem: %d", cup, mem)
	return nil
}

// AliyunClientAdapter aliyun client adapter
type AliyunClientAdapter struct {
	Client AliyunClient
}

// CreateServer CreateServer function
func (a *AliyunClientAdapter) CreateServer(cup, mem float64) error {
	a.Client.CreateServer(int(cup), int(mem))
	return nil
}
