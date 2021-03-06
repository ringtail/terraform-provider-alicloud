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

// Item is a nested struct in emr response
type Item struct {
	Name                   string                                 `json:"Name" xml:"Name"`
	EntityGroupId          string                                 `json:"EntityGroupId" xml:"EntityGroupId"`
	Operation              string                                 `json:"Operation" xml:"Operation"`
	CreateTime             int64                                  `json:"CreateTime" xml:"CreateTime"`
	AliyunUserId           string                                 `json:"AliyunUserId" xml:"AliyunUserId"`
	QueryId                string                                 `json:"QueryId" xml:"QueryId"`
	Index                  int                                    `json:"Index" xml:"Index"`
	SourceType             string                                 `json:"SourceType" xml:"SourceType"`
	Content                string                                 `json:"Content" xml:"Content"`
	InstanceId             string                                 `json:"InstanceId" xml:"InstanceId"`
	UserId                 string                                 `json:"UserId" xml:"UserId"`
	Md5                    string                                 `json:"Md5" xml:"Md5"`
	RootPath               string                                 `json:"RootPath" xml:"RootPath"`
	OperatorId             string                                 `json:"OperatorId" xml:"OperatorId"`
	Id                     string                                 `json:"Id" xml:"Id"`
	Description            string                                 `json:"Description" xml:"Description"`
	KnoxSyncStatus         string                                 `json:"KnoxSyncStatus" xml:"KnoxSyncStatus"`
	KerberosStutus         string                                 `json:"KerberosStutus" xml:"KerberosStutus"`
	QueryName              string                                 `json:"QueryName" xml:"QueryName"`
	TarFileName            string                                 `json:"TarFileName" xml:"TarFileName"`
	ClusterId              string                                 `json:"ClusterId" xml:"ClusterId"`
	Error                  string                                 `json:"Error" xml:"Error"`
	StorePath              string                                 `json:"StorePath" xml:"StorePath"`
	Ts                     int64                                  `json:"Ts" xml:"Ts"`
	BackupPlanId           string                                 `json:"BackupPlanId" xml:"BackupPlanId"`
	BackupMethodType       string                                 `json:"BackupMethodType" xml:"BackupMethodType"`
	GmtCreate              int64                                  `json:"GmtCreate" xml:"GmtCreate"`
	MetadataType           string                                 `json:"MetadataType" xml:"MetadataType"`
	EntityId               string                                 `json:"EntityId" xml:"EntityId"`
	RunId                  string                                 `json:"RunId" xml:"RunId"`
	EntityType             string                                 `json:"EntityType" xml:"EntityType"`
	EventId                string                                 `json:"EventId" xml:"EventId"`
	GmtModified            int64                                  `json:"GmtModified" xml:"GmtModified"`
	Enable                 bool                                   `json:"Enable" xml:"Enable"`
	UserName               string                                 `json:"UserName" xml:"UserName"`
	Status                 string                                 `json:"Status" xml:"Status"`
	LinuxSyncStatus        string                                 `json:"LinuxSyncStatus" xml:"LinuxSyncStatus"`
	AlertUserGroupIdList   AlertUserGroupIdList                   `json:"AlertUserGroupIdList" xml:"AlertUserGroupIdList"`
	AlertDingDingGroupList AlertDingDingGroupListInDescribeETLJob `json:"AlertDingDingGroupList" xml:"AlertDingDingGroupList"`
	MetadataInfo           MetadataInfo                           `json:"MetadataInfo" xml:"MetadataInfo"`
	RoleDTOList            RoleDTOListInPageListResourceUsers     `json:"RoleDTOList" xml:"RoleDTOList"`
}
