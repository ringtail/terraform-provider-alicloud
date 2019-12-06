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

// CheckSqlAuditEnableStatus invokes the drds.CheckSqlAuditEnableStatus API synchronously
// api document: https://help.aliyun.com/api/drds/checksqlauditenablestatus.html
func (client *Client) CheckSqlAuditEnableStatus(request *CheckSqlAuditEnableStatusRequest) (response *CheckSqlAuditEnableStatusResponse, err error) {
	response = CreateCheckSqlAuditEnableStatusResponse()
	err = client.DoAction(request, response)
	return
}

// CheckSqlAuditEnableStatusWithChan invokes the drds.CheckSqlAuditEnableStatus API asynchronously
// api document: https://help.aliyun.com/api/drds/checksqlauditenablestatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckSqlAuditEnableStatusWithChan(request *CheckSqlAuditEnableStatusRequest) (<-chan *CheckSqlAuditEnableStatusResponse, <-chan error) {
	responseChan := make(chan *CheckSqlAuditEnableStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CheckSqlAuditEnableStatus(request)
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

// CheckSqlAuditEnableStatusWithCallback invokes the drds.CheckSqlAuditEnableStatus API asynchronously
// api document: https://help.aliyun.com/api/drds/checksqlauditenablestatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CheckSqlAuditEnableStatusWithCallback(request *CheckSqlAuditEnableStatusRequest, callback func(response *CheckSqlAuditEnableStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CheckSqlAuditEnableStatusResponse
		var err error
		defer close(result)
		response, err = client.CheckSqlAuditEnableStatus(request)
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

// CheckSqlAuditEnableStatusRequest is the request struct for api CheckSqlAuditEnableStatus
type CheckSqlAuditEnableStatusRequest struct {
	*requests.RpcRequest
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
	DbName         string `position:"Query" name:"DbName"`
}

// CheckSqlAuditEnableStatusResponse is the response struct for api CheckSqlAuditEnableStatus
type CheckSqlAuditEnableStatusResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Status    string `json:"Status" xml:"Status"`
}

// CreateCheckSqlAuditEnableStatusRequest creates a request to invoke CheckSqlAuditEnableStatus API
func CreateCheckSqlAuditEnableStatusRequest() (request *CheckSqlAuditEnableStatusRequest) {
	request = &CheckSqlAuditEnableStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2019-01-23", "CheckSqlAuditEnableStatus", "Drds", "openAPI")
	return
}

// CreateCheckSqlAuditEnableStatusResponse creates a response to parse from CheckSqlAuditEnableStatus response
func CreateCheckSqlAuditEnableStatusResponse() (response *CheckSqlAuditEnableStatusResponse) {
	response = &CheckSqlAuditEnableStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
