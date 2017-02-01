package push

import (
	"net/http"

	"github.com/denverdino/aliyungo/common"
)

//高级推送参数
type PushArgs struct {
	/*----基础参数----*/
	//AppKey信息
	AppKey                         int64
	/*----推送目标----*/
	//推送目标
	Target                         string
	//根据Target来设定，多个值使用逗号分隔，最多支持100个。
	TargetValue                    string
	//设备类型
	DeviceType                     string
	/*----推送配置----*/
	PushType                       string
	//Android消息标题,Android通知标题,iOS消息标题
	Title                          string
	//Android消息内容,Android通知内容,iOS消息内容
	Body                           string
	//[iOS通知内容]
	Summary                        string
	/*----下述配置仅作用于iOS通知任务----*/
	//[iOS通知声音]
	IOSMusic                       string `ArgName:"iOSMusic"`
	//[iOS应用图标右上角角标]
	IOSBadge                       string `ArgName:"iOSBadge"`
	//[iOS通知标题（iOS 10+通知显示标题]
	IOSTitle                       string `ArgName:"iOSTitle"`
	//[开启iOS静默通知]
	IOSSilentNotification          string `ArgName:"iOSSilentNotification"`
	//[iOS通知副标题（iOS 10+）]
	IOSSubtitle                    string `ArgName:"iOSSubtitle"`
	//[设定iOS通知Category（iOS 10+）]
	IOSNotificationCategory        string `ArgName:"iOSNotificationCategory"`
	//[是否使能iOS通知扩展处理（iOS 10+）]
	IOSMutableContent              string `ArgName:"iOSMutableContent"`
	//[iOS通知的扩展属性]
	IOSExtParameters               string `ArgName:"iOSExtParameters"`
	//[环境信息]
	IOSApnsEnv                     string `ArgName:"iOSApnsEnv"`
	//[推送时设备不在线则这条推送会做为通知]
	IOSRemind                      string `ArgName:"iOSRemind"`
	//[iOS消息转通知时使用的iOS通知内容]
	IOSRemindBody                  string `ArgName:"iOSRemindBody"`
	/*----下述配置仅作用于Android通知任务----*/
	//[Android通知声音]
	AndroidMusic                   string
	//[点击通知后动作]
	AndroidOpenType                string
	//通知的提醒方式
	AndroidNotifyType              string
	//[设置该参数后启动小米托管弹窗功能]
	AndroidXiaoMiActivity          string
	//[小米托管弹窗模式下Title内容]
	AndroidXiaoMiNotifyTitle       string
	//[小米托管弹窗模式下Body内容]
	AndroidXiaoMiNotifyBody        string
	//[设定通知打开的activity]
	AndroidActivity                string
	//[Android收到推送后打开对应的url]
	AndroidOpenUrl                 string
	//[Android自定义通知栏样式]
	AndroidNotificationBarType     int
	//[Android通知在通知栏展示时排列位置的优先级]
	AndroidNotificationBarPriority int
	//[设定通知的扩展属性]
	AndroidExtParameters           string
	/*----推送控制----*/
	//[用于定时发送]
	PushTime                       string
	//[离线消息/通知是否保存]
	StoreOffline                   string
	//[离线消息/通知的过期时间]
	ExpireTime                     string
}

func (this *Client) Push(args *PushArgs) error {
	return this.InvokeByAnyMethod(http.MethodPost, Push, "", args, &common.Response{})
}