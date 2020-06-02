package cms

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

// DescribeEventRuleTargetList invokes the cms.DescribeEventRuleTargetList API synchronously
// api document: https://help.aliyun.com/api/cms/describeeventruletargetlist.html
func (client *Client) DescribeEventRuleTargetList(request *DescribeEventRuleTargetListRequest) (response *DescribeEventRuleTargetListResponse, err error) {
	response = CreateDescribeEventRuleTargetListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeEventRuleTargetListWithChan invokes the cms.DescribeEventRuleTargetList API asynchronously
// api document: https://help.aliyun.com/api/cms/describeeventruletargetlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeEventRuleTargetListWithChan(request *DescribeEventRuleTargetListRequest) (<-chan *DescribeEventRuleTargetListResponse, <-chan error) {
	responseChan := make(chan *DescribeEventRuleTargetListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeEventRuleTargetList(request)
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

// DescribeEventRuleTargetListWithCallback invokes the cms.DescribeEventRuleTargetList API asynchronously
// api document: https://help.aliyun.com/api/cms/describeeventruletargetlist.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeEventRuleTargetListWithCallback(request *DescribeEventRuleTargetListRequest, callback func(response *DescribeEventRuleTargetListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeEventRuleTargetListResponse
		var err error
		defer close(result)
		response, err = client.DescribeEventRuleTargetList(request)
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

// DescribeEventRuleTargetListRequest is the request struct for api DescribeEventRuleTargetList
type DescribeEventRuleTargetListRequest struct {
	*requests.RpcRequest
	RuleName string `position:"Query" name:"RuleName"`
}

// DescribeEventRuleTargetListResponse is the response struct for api DescribeEventRuleTargetList
type DescribeEventRuleTargetListResponse struct {
	*responses.BaseResponse
	Code              string            `json:"Code" xml:"Code"`
	Message           string            `json:"Message" xml:"Message"`
	RequestId         string            `json:"RequestId" xml:"RequestId"`
	ContactParameters ContactParameters `json:"ContactParameters" xml:"ContactParameters"`
	FcParameters      FcParameters      `json:"FcParameters" xml:"FcParameters"`
	MnsParameters     MnsParameters     `json:"MnsParameters" xml:"MnsParameters"`
	WebhookParameters WebhookParameters `json:"WebhookParameters" xml:"WebhookParameters"`
	SlsParameters     SlsParameters     `json:"SlsParameters" xml:"SlsParameters"`
}

// CreateDescribeEventRuleTargetListRequest creates a request to invoke DescribeEventRuleTargetList API
func CreateDescribeEventRuleTargetListRequest() (request *DescribeEventRuleTargetListRequest) {
	request = &DescribeEventRuleTargetListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cms", "2019-01-01", "DescribeEventRuleTargetList", "cms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeEventRuleTargetListResponse creates a response to parse from DescribeEventRuleTargetList response
func CreateDescribeEventRuleTargetListResponse() (response *DescribeEventRuleTargetListResponse) {
	response = &DescribeEventRuleTargetListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
