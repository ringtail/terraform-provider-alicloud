package rds

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeDedicatedHostAttribute invokes the rds.DescribeDedicatedHostAttribute API synchronously
// api document: https://help.aliyun.com/api/rds/describededicatedhostattribute.html
func (client *Client) DescribeDedicatedHostAttribute(request *DescribeDedicatedHostAttributeRequest) (response *DescribeDedicatedHostAttributeResponse, err error) {
	response = CreateDescribeDedicatedHostAttributeResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDedicatedHostAttributeWithChan invokes the rds.DescribeDedicatedHostAttribute API asynchronously
// api document: https://help.aliyun.com/api/rds/describededicatedhostattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDedicatedHostAttributeWithChan(request *DescribeDedicatedHostAttributeRequest) (<-chan *DescribeDedicatedHostAttributeResponse, <-chan error) {
	responseChan := make(chan *DescribeDedicatedHostAttributeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDedicatedHostAttribute(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeDedicatedHostAttributeWithCallback invokes the rds.DescribeDedicatedHostAttribute API asynchronously
// api document: https://help.aliyun.com/api/rds/describededicatedhostattribute.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeDedicatedHostAttributeWithCallback(request *DescribeDedicatedHostAttributeRequest, callback func(response *DescribeDedicatedHostAttributeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDedicatedHostAttributeResponse
		var err error
		defer close(result)
		response, err = client.DescribeDedicatedHostAttribute(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeDedicatedHostAttributeRequest is the request struct for api DescribeDedicatedHostAttribute
type DescribeDedicatedHostAttributeRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	DedicatedHostId      string           `position:"Query" name:"DedicatedHostId"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	DedicatedHostGroupId string           `position:"Query" name:"DedicatedHostGroupId"`
}

// DescribeDedicatedHostAttributeResponse is the response struct for api DescribeDedicatedHostAttribute
type DescribeDedicatedHostAttributeResponse struct {
	*responses.BaseResponse
	RequestId              string `json:"RequestId" xml:"RequestId"`
	DedicatedHostGroupId   string `json:"DedicatedHostGroupId" xml:"DedicatedHostGroupId"`
	DedicatedHostId        string `json:"DedicatedHostId" xml:"DedicatedHostId"`
	RegionId               string `json:"RegionId" xml:"RegionId"`
	ZoneId                 string `json:"ZoneId" xml:"ZoneId"`
	VPCId                  string `json:"VPCId" xml:"VPCId"`
	VSwitchId              string `json:"VSwitchId" xml:"VSwitchId"`
	IPAddress              string `json:"IPAddress" xml:"IPAddress"`
	HostName               string `json:"HostName" xml:"HostName"`
	HostStatus             string `json:"HostStatus" xml:"HostStatus"`
	HostClass              string `json:"HostClass" xml:"HostClass"`
	HostCPU                int    `json:"HostCPU" xml:"HostCPU"`
	HostMem                int    `json:"HostMem" xml:"HostMem"`
	HostStorage            int    `json:"HostStorage" xml:"HostStorage"`
	CPUAllocationRatio     string `json:"CPUAllocationRatio" xml:"CPUAllocationRatio"`
	MemAllocationRatio     string `json:"MemAllocationRatio" xml:"MemAllocationRatio"`
	DiskAllocationRatio    string `json:"DiskAllocationRatio" xml:"DiskAllocationRatio"`
	InstanceNumber         int    `json:"InstanceNumber" xml:"InstanceNumber"`
	InstanceNumberMaster   int    `json:"InstanceNumberMaster" xml:"InstanceNumberMaster"`
	InstanceNumberSlave    int    `json:"InstanceNumberSlave" xml:"InstanceNumberSlave"`
	InstanceNumberROMaster int    `json:"InstanceNumberROMaster" xml:"InstanceNumberROMaster"`
	InstanceNumberROSlave  int    `json:"InstanceNumberROSlave" xml:"InstanceNumberROSlave"`
	CreatedTime            string `json:"CreatedTime" xml:"CreatedTime"`
	ExpiredTime            string `json:"ExpiredTime" xml:"ExpiredTime"`
	AutoRenew              string `json:"AutoRenew" xml:"AutoRenew"`
	AllocationStatus       string `json:"AllocationStatus" xml:"AllocationStatus"`
	CpuUsed                string `json:"CpuUsed" xml:"CpuUsed"`
	MemoryUsed             string `json:"MemoryUsed" xml:"MemoryUsed"`
	StorageUsed            string `json:"StorageUsed" xml:"StorageUsed"`
}

// CreateDescribeDedicatedHostAttributeRequest creates a request to invoke DescribeDedicatedHostAttribute API
func CreateDescribeDedicatedHostAttributeRequest() (request *DescribeDedicatedHostAttributeRequest) {
	request = &DescribeDedicatedHostAttributeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeDedicatedHostAttribute", "", "")
	return
}

// CreateDescribeDedicatedHostAttributeResponse creates a response to parse from DescribeDedicatedHostAttribute response
func CreateDescribeDedicatedHostAttributeResponse() (response *DescribeDedicatedHostAttributeResponse) {
	response = &DescribeDedicatedHostAttributeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
