package yundun_bastionhost

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

// UntagResourcesSystemTags invokes the yundun_bastionhost.UntagResourcesSystemTags API synchronously
// api document: https://help.aliyun.com/api/yundun-bastionhost/untagresourcessystemtags.html
func (client *Client) UntagResourcesSystemTags(request *UntagResourcesSystemTagsRequest) (response *UntagResourcesSystemTagsResponse, err error) {
	response = CreateUntagResourcesSystemTagsResponse()
	err = client.DoAction(request, response)
	return
}

// UntagResourcesSystemTagsWithChan invokes the yundun_bastionhost.UntagResourcesSystemTags API asynchronously
// api document: https://help.aliyun.com/api/yundun-bastionhost/untagresourcessystemtags.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UntagResourcesSystemTagsWithChan(request *UntagResourcesSystemTagsRequest) (<-chan *UntagResourcesSystemTagsResponse, <-chan error) {
	responseChan := make(chan *UntagResourcesSystemTagsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UntagResourcesSystemTags(request)
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

// UntagResourcesSystemTagsWithCallback invokes the yundun_bastionhost.UntagResourcesSystemTags API asynchronously
// api document: https://help.aliyun.com/api/yundun-bastionhost/untagresourcessystemtags.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UntagResourcesSystemTagsWithCallback(request *UntagResourcesSystemTagsRequest, callback func(response *UntagResourcesSystemTagsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UntagResourcesSystemTagsResponse
		var err error
		defer close(result)
		response, err = client.UntagResourcesSystemTags(request)
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

// UntagResourcesSystemTagsRequest is the request struct for api UntagResourcesSystemTags
type UntagResourcesSystemTagsRequest struct {
	*requests.RpcRequest
	All          requests.Boolean `position:"Query" name:"All"`
	ResourceId   *[]string        `position:"Query" name:"ResourceId"  type:"Repeated"`
	ResourceType string           `position:"Query" name:"ResourceType"`
	SourceIp     string           `position:"Query" name:"SourceIp"`
	TagKey       *[]string        `position:"Query" name:"TagKey"  type:"Repeated"`
}

// UntagResourcesSystemTagsResponse is the response struct for api UntagResourcesSystemTags
type UntagResourcesSystemTagsResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUntagResourcesSystemTagsRequest creates a request to invoke UntagResourcesSystemTags API
func CreateUntagResourcesSystemTagsRequest() (request *UntagResourcesSystemTagsRequest) {
	request = &UntagResourcesSystemTagsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Yundun-bastionhost", "2018-10-10", "UntagResourcesSystemTags", "bastionhost", "openAPI")
	return
}

// CreateUntagResourcesSystemTagsResponse creates a response to parse from UntagResourcesSystemTags response
func CreateUntagResourcesSystemTagsResponse() (response *UntagResourcesSystemTagsResponse) {
	response = &UntagResourcesSystemTagsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
