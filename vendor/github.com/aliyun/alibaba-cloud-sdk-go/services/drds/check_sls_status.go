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

// CheckSlsStatus invokes the drds.CheckSlsStatus API synchronously
// api document: https://help.aliyun.com/api/drds/checkslsstatus.html
func (client *Client) CheckSlsStatus(request *CheckSlsStatusRequest) (response *CheckSlsStatusResponse, err error) {
	response = CreateCheckSlsStatusResponse()
	err = client.DoAction(request, response)
	return
}

// CheckSlsStatusWithChan invokes the drds.CheckSlsStatus API asynchronously
// api document: https://help.aliyun.com/api/drds/checkslsstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckSlsStatusWithChan(request *CheckSlsStatusRequest) (<-chan *CheckSlsStatusResponse, <-chan error) {
	responseChan := make(chan *CheckSlsStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckSlsStatus(request)
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

// CheckSlsStatusWithCallback invokes the drds.CheckSlsStatus API asynchronously
// api document: https://help.aliyun.com/api/drds/checkslsstatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckSlsStatusWithCallback(request *CheckSlsStatusRequest, callback func(response *CheckSlsStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckSlsStatusResponse
		var err error
		defer close(result)
		response, err = client.CheckSlsStatus(request)
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

// CheckSlsStatusRequest is the request struct for api CheckSlsStatus
type CheckSlsStatusRequest struct {
	*requests.RpcRequest
}

// CheckSlsStatusResponse is the response struct for api CheckSlsStatus
type CheckSlsStatusResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Status    string `json:"Status" xml:"Status"`
}

// CreateCheckSlsStatusRequest creates a request to invoke CheckSlsStatus API
func CreateCheckSlsStatusRequest() (request *CheckSlsStatusRequest) {
	request = &CheckSlsStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "CheckSlsStatus", "drds", "openAPI")
	return
}

// CreateCheckSlsStatusResponse creates a response to parse from CheckSlsStatus response
func CreateCheckSlsStatusResponse() (response *CheckSlsStatusResponse) {
	response = &CheckSlsStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
