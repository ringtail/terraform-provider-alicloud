package emr

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

// ResolveETLJobSqlSchema invokes the emr.ResolveETLJobSqlSchema API synchronously
// api document: https://help.aliyun.com/api/emr/resolveetljobsqlschema.html
func (client *Client) ResolveETLJobSqlSchema(request *ResolveETLJobSqlSchemaRequest) (response *ResolveETLJobSqlSchemaResponse, err error) {
	response = CreateResolveETLJobSqlSchemaResponse()
	err = client.DoAction(request, response)
	return
}

// ResolveETLJobSqlSchemaWithChan invokes the emr.ResolveETLJobSqlSchema API asynchronously
// api document: https://help.aliyun.com/api/emr/resolveetljobsqlschema.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ResolveETLJobSqlSchemaWithChan(request *ResolveETLJobSqlSchemaRequest) (<-chan *ResolveETLJobSqlSchemaResponse, <-chan error) {
	responseChan := make(chan *ResolveETLJobSqlSchemaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ResolveETLJobSqlSchema(request)
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

// ResolveETLJobSqlSchemaWithCallback invokes the emr.ResolveETLJobSqlSchema API asynchronously
// api document: https://help.aliyun.com/api/emr/resolveetljobsqlschema.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ResolveETLJobSqlSchemaWithCallback(request *ResolveETLJobSqlSchemaRequest, callback func(response *ResolveETLJobSqlSchemaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ResolveETLJobSqlSchemaResponse
		var err error
		defer close(result)
		response, err = client.ResolveETLJobSqlSchema(request)
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

// ResolveETLJobSqlSchemaRequest is the request struct for api ResolveETLJobSqlSchema
type ResolveETLJobSqlSchemaRequest struct {
	*requests.RpcRequest
	StageName       string           `position:"Query" name:"StageName"`
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	EtlJobId        string           `position:"Query" name:"EtlJobId"`
	DataSourceId    string           `position:"Query" name:"DataSourceId"`
	Sql             string           `position:"Query" name:"Sql"`
}

// ResolveETLJobSqlSchemaResponse is the response struct for api ResolveETLJobSqlSchema
type ResolveETLJobSqlSchemaResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	ResolveId string `json:"ResolveId" xml:"ResolveId"`
}

// CreateResolveETLJobSqlSchemaRequest creates a request to invoke ResolveETLJobSqlSchema API
func CreateResolveETLJobSqlSchemaRequest() (request *ResolveETLJobSqlSchemaRequest) {
	request = &ResolveETLJobSqlSchemaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Emr", "2016-04-08", "ResolveETLJobSqlSchema", "emr", "openAPI")
	return
}

// CreateResolveETLJobSqlSchemaResponse creates a response to parse from ResolveETLJobSqlSchema response
func CreateResolveETLJobSqlSchemaResponse() (response *ResolveETLJobSqlSchemaResponse) {
	response = &ResolveETLJobSqlSchemaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
