package drds

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

// CheckRdsExpandStatus invokes the drds.CheckRdsExpandStatus API synchronously
// api document: https://help.aliyun.com/api/drds/checkrdsexpandstatus.html
func (client *Client) CheckRdsExpandStatus(request *CheckRdsExpandStatusRequest) (response *CheckRdsExpandStatusResponse, err error) {
	response = CreateCheckRdsExpandStatusResponse()
	err = client.DoAction(request, response)
	return
}

// CheckRdsExpandStatusWithChan invokes the drds.CheckRdsExpandStatus API asynchronously
// api document: https://help.aliyun.com/api/drds/checkrdsexpandstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckRdsExpandStatusWithChan(request *CheckRdsExpandStatusRequest) (<-chan *CheckRdsExpandStatusResponse, <-chan error) {
	responseChan := make(chan *CheckRdsExpandStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckRdsExpandStatus(request)
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

// CheckRdsExpandStatusWithCallback invokes the drds.CheckRdsExpandStatus API asynchronously
// api document: https://help.aliyun.com/api/drds/checkrdsexpandstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckRdsExpandStatusWithCallback(request *CheckRdsExpandStatusRequest, callback func(response *CheckRdsExpandStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckRdsExpandStatusResponse
		var err error
		defer close(result)
		response, err = client.CheckRdsExpandStatus(request)
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

// CheckRdsExpandStatusRequest is the request struct for api CheckRdsExpandStatus
type CheckRdsExpandStatusRequest struct {
	*requests.RpcRequest
	InstanceList   *[]string `position:"Query" name:"InstanceList"  type:"Repeated"`
	DbName         string    `position:"Query" name:"DbName"`
	DrdsInstanceId string    `position:"Query" name:"DrdsInstanceId"`
}

// CheckRdsExpandStatusResponse is the response struct for api CheckRdsExpandStatus
type CheckRdsExpandStatusResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateCheckRdsExpandStatusRequest creates a request to invoke CheckRdsExpandStatus API
func CreateCheckRdsExpandStatusRequest() (request *CheckRdsExpandStatusRequest) {
	request = &CheckRdsExpandStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "CheckRdsExpandStatus", "drds", "openAPI")
	return
}

// CreateCheckRdsExpandStatusResponse creates a response to parse from CheckRdsExpandStatus response
func CreateCheckRdsExpandStatusResponse() (response *CheckRdsExpandStatusResponse) {
	response = &CheckRdsExpandStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
