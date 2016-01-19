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

type User struct {
	UserId      string
	UserName    string
	DisplayName string
	MobilePhone string
	Email       string
	Comments    string
	CreateDate  string
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
}

type PolicyVersion struct {
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
