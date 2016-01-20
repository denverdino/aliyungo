package ram

import (
	"github.com/denverdino/aliyungo/common"
)

/*
	All common struct
*/

const (
	Active   State = "Active"
	Inactive State = "Inactive"
)

/*
	AccountAlias
	类型：String
	必须：是
	描述：指定云账号的别名, 长度限制为3-63个字符
	限制：^[a-z0-9](([a-z0-9]|-(?!-))*[a-z0-9])?$
*/
type AccountAlias string

type UserQueryRequest struct {
	UserName string
}

type User struct {
	UserId        string `json:"UserId,omitempty"`
	UserName      string `json:"UserName"`
	DisplayName   string `json:"DisplayName,omitempty"`
	MobilePhone   string `json:"MobilePhone,omitempty"`
	Email         string `json:"Email,omitempty"`
	Comments      string `json:"Comments,omitempty"`
	CreateDate    string `json:"CreateDate,omitempty"`
	UpdateDate    string `json:"UpdateDate,omitempty"`
	LastLoginDate string `json:"LastLoginDate,omitempty"`
}

type LoginProfile struct {
}

type MFADevice struct {
}

type VirtualMFADevice struct {
}

type AccessKey struct {
	AccessKeyId     string
	AccessKeySecret string
	Status          State
	CreateDate      string
}

type Group struct {
}

type Role struct {
}

type Policy struct {
	PolicyName      string
	PolicyType      string
	Description     string
	DefaultVersion  string
	CreateDate      string
	UpdateDate      string
	AttachmentCount int64
}

type PolicyDocument struct {
	Statement []PolicyItem
	Version   string
}

type PolicyItem struct {
	Action   string
	Effect   string
	Resource string
}

/*
	"PasswordPolicy": {
        "MinimumPasswordLength": 12,
        "RequireLowercaseCharacters": true,
        "RequireUppercaseCharacters": true,
        "RequireNumbers": true,
        "RequireSymbols": true
    }
*/

type PasswordPolicy struct {
	MinimumPasswordLength      int8 `json:"MinimumPasswordLength"`
	RequireLowercaseCharacters bool `json:"RequireLowercaseCharacters"`
	RequireUppercaseCharacters bool `json:"RequireUppercaseCharacters"`
	RequireNumbers             bool `json:"RequireNumbers"`
	RequireSymbols             bool `json:"RequireSymbols"`
}

type RamCommonResponse struct {
	common.Response
}
