package errors

import (
	"baikai/leetcode/a_project/8/i18n"
)

const (
	ErrorInvalidParams                 = 400
	SUCCESS                            = 200
	ERROR                              = 500
	ErrorFailedCaptchaAction           = 10007
	ErrorMailConfigurationDoesNotExist = 10008
	ErrorConnectTimeout                = 10009
	ErrorInvalidParams1                = 10005
	ErrorAction                        = 10010
	ErrorSystemDoesNotOpenRegistration = 10011
	ErrorSysBasicNotExist              = 10012
	ErrorRecipientNotExist             = 10013
	ErrorDBInfo                        = 10014
	ErrorSsh                           = 10015
	ErrorGsOm                          = 10016
)

// user_svc
const (
	ErrorConflict                    = 11000
	ErrorNotAcceptable               = 11016
	ErrorUnAuthorized                = 11024
	ErrorTokenTimeout                = 11025
	ErrorTokenAuthFailed             = 11026
	ErrorCaptchaTimeout              = 11006
	ErrorExistEmail                  = 11002
	ErrorEmailRateLimit              = 11008
	ErrorSendEmailFailed             = 11017
	ErrorUsernameAndPasswordNotMatch = 11020
	ErrorUserNotExist                = 11003
	ErrorCaptchaNotMatch             = 11007
	ErrorFailedGetUserInfo           = 11040
	ErrorFailedOperateTimeOut        = 11041 //
	ErrorWrongPassword               = 11004
	ErrorUserPhoneExist              = 11044
	ErrorUserNickname                = 11045
	ErrorAccountIsNotAvailable       = 11046
	ErrorLoginInvalid                = 11047
	ErrorTestSendEmail               = 11048
)

// 团队 12000-12999
const (
	ErrorInviteExpired              = 12001
	ErrorExistGroupName             = 12007
	ErrorOther                      = 12999
	ErrorExistOtherUser             = 12005
	ErrorAddUserRoleFailed          = 11039
	ErrorTimeOut                    = 12008
	ErrorUserRestrict               = 12013
	ErrorExist                      = 12011
	ErrorUserIsOwner                = 12021
	ErrorGroupName                  = 12016
	ErrorInviteColNotExist          = 12028
	ErrorUserDontExist              = 12022
	ErrorExistName                  = 15011
	ErrorNotHaveOwnerRole           = 12989
	ErrorInvalidParamsDefaultRole   = 12988
	ErrorAllFailed                  = 12987
	ErrorInviteAlreadyAccepted      = 12002
	ErrorAlreadyInvited             = 12012
	ErrorSameAccountCanOnlyCloudKey = 13010
	ErrorAccessKeyExist             = 13008
	ErrorNotExist                   = 13003
	ErrorAccessKeyInvalid           = 13002
	ErrorUserHasSpecificRole        = 12059
)

const (
	ErrorConnectionDB        = 14000
	ErrorOpenDB              = 14001
	ErrorDuplicateLabelKey   = 14002
	ErrorDbNotExist          = 14003
	ErrorConnectionDB1       = 14004
	ErrorTagNotExist         = 14005
	ErrorDbHostIPExist       = 14006
	ErrorTagExist            = 14007
	ErrorParsingFailed       = 14008
	ErrorUpdateFailed        = 14009
	ErrorCheckUserNameFailed = 14010
	ErrorTestContent         = 14011
	ErrorUpSteamErr          = 14012
	ErrorPublisher           = 14013
	ErrorPlatform            = 14014
	ErrorVersion             = 14015
	ErrorDbVersion           = 14016
	ErrorIPIsExist           = 14017
	ErrorDbNumError          = 14018
	ErrorWhetherPython       = 14019
	ErrorWhetherLibaio       = 14020
	ErrorIsNotExistCluster   = 14021
	ErrorInstallIngDoNotDel  = 14022
	ErrorNodeIsNotExist      = 14023
	ErrorIsNotExistUnzip     = 14024
	ErrorStartMogHaFail      = 14025
	ErrorVipFail             = 14026
	ErrorPingListFail        = 14027
	ErrorArpingFail          = 14028
	ErrorNotSameCluster      = 14029
	ErrorZoneSame            = 14030
	ErrorPublisherFail       = 14031
	ErrorClusterUnavailable  = 14032
	ErrorClusterNotPrimary   = 14033
	ErrorCascadeStandby      = 14034
	ErrorZoneHavePri         = 14035
	ErrorOperatingTimeout    = 14036
)

var codeToStr = map[int]map[i18n.Language]string{
	SUCCESS: {
		i18n.LangZh: "成功",
		i18n.LangEn: "success",
	},
	ERROR: {
		i18n.LangZh: "系统异常，请联系管理员",
		i18n.LangEn: "The system is abnormal, please contact the administrator",
	},
	ErrorCaptchaTimeout: {
		i18n.LangZh: "验证码已过期",
		i18n.LangEn: "Verification code has expired",
	},
	ErrorTokenAuthFailed: {
		i18n.LangZh: "token认证失败",
		i18n.LangEn: "token authentication failed",
	},
	ErrorTokenTimeout: {
		i18n.LangZh: "token超时",
		i18n.LangEn: "token timeout",
	},
	ErrorExistEmail: {
		i18n.LangZh: "邮箱已存在",
		i18n.LangEn: "Email already exists",
	},
	ErrorNotAcceptable: {
		i18n.LangZh: "用户无权限操作此团队",
		i18n.LangEn: "user_svc does not have permission to operate this team",
	},
	ErrorUnAuthorized: {
		i18n.LangZh: "未授权",
		i18n.LangEn: "unauthorized",
	},
	ErrorExistGroupName: {
		i18n.LangZh: "团队名已存在",
		i18n.LangEn: "The team name already exists",
	},
	ErrorExistOtherUser: {
		i18n.LangZh: "团队存在其他用户",
		i18n.LangEn: "There are other users on the team",
	},
	ErrorAddUserRoleFailed: {
		i18n.LangZh: "添加用户角色失败,请联系团队管理员进行手工添加",
		i18n.LangEn: "Failed to add user_svc role, please contact the team administrator to add manually",
	},
	ErrorTimeOut: {
		i18n.LangZh: "超时",
		i18n.LangEn: "time out",
	},
	ErrorUserRestrict: {
		i18n.LangZh: "用户无权限",
		i18n.LangEn: "User has no permissions",
	},
	ErrorExist: {
		i18n.LangZh: "已加入",
		i18n.LangEn: "joined",
	},
	ErrorUserIsOwner: {
		i18n.LangZh: "不可操作此账号",
		i18n.LangEn: "This account cannot be operated",
	},
	ErrorGroupName: {
		i18n.LangZh: "团队名字不能为空",
		i18n.LangEn: "Team name cannot be empty",
	},
	ErrorInviteColNotExist: {
		i18n.LangZh: "邀请记录不存在或者已经被删除",
		i18n.LangEn: "The invitation record does not exist or has been deleted",
	},
	ErrorUserDontExist: {
		i18n.LangZh: "用户不存在",
		i18n.LangEn: "User does not exist",
	},
	ErrorExistName: {
		i18n.LangZh: "名称已存在",
		i18n.LangEn: "Name already exists",
	},
	ErrorInvalidParams: {
		i18n.LangZh: "请求参数错误",
		i18n.LangEn: "Request parameter error",
	},
	ErrorNotHaveOwnerRole: {
		i18n.LangZh: "无法添加所有者角色",
		i18n.LangEn: "can not add owner role",
	},
	ErrorInvalidParamsDefaultRole: {
		i18n.LangZh: "请求参数错误 default role",
		i18n.LangEn: "Request parameter error default role",
	},
	ErrorAllFailed: {
		i18n.LangZh: "全部失败",
		i18n.LangEn: "all failed",
	},
	ErrorInviteAlreadyAccepted: {
		i18n.LangZh: "邀请已接受",
		i18n.LangEn: "Invitation accepted",
	},
	ErrorAlreadyInvited: {
		i18n.LangZh: "该用户已被邀请",
		i18n.LangEn: "This user_svc has been invited",
	},
	ErrorSameAccountCanOnlyCloudKey: {
		i18n.LangZh: "该账号已存在",
		i18n.LangEn: "The account already exists",
	},
	ErrorAccessKeyExist: {
		i18n.LangZh: "云账号已存在",
		i18n.LangEn: "Cloud Key already exists",
	},
	ErrorNotExist: {
		i18n.LangZh: "不在该组，无法查看",
		i18n.LangEn: "Not in this group, unable to view",
	},
	ErrorAccessKeyInvalid: {
		i18n.LangZh: "输入的密钥无效",
		i18n.LangEn: "Invalid key entered",
	},
	ErrorUserHasSpecificRole: {
		i18n.LangEn: "invalid role err",
		i18n.LangZh: "无效的角色",
	},
	ErrorFailedCaptchaAction: {
		i18n.LangEn: "Incorrect verification code type",
		i18n.LangZh: "验证码类型不正确",
	},
	ErrorEmailRateLimit: {
		i18n.LangZh: "发送验证码频率过快",
		i18n.LangEn: "Send verification code too fast",
	},
	ErrorSendEmailFailed: {
		i18n.LangZh: "邮件发送失败，请确认邮箱存在",
		i18n.LangEn: "Failed to send the mail, please confirm the mailbox exists",
	},
	ErrorConflict: {
		i18n.LangZh: "条目已存在",
		i18n.LangEn: "Entry already exists",
	},
	ErrorMailConfigurationDoesNotExist: {
		i18n.LangZh: "邮件配置不存在",
		i18n.LangEn: "Error Mail Configuration Does Not Exist",
	},
	ErrorUsernameAndPasswordNotMatch: {
		i18n.LangEn: "Incorrect account name or password",
		i18n.LangZh: "账号或密码错误",
	},
	ErrorUserNotExist: {
		i18n.LangZh: "用户不存在",
		i18n.LangEn: "User does not exist",
	},
	ErrorCaptchaNotMatch: {
		i18n.LangZh: "验证码不匹配",
		i18n.LangEn: "Verification code does not match",
	},
	ErrorFailedGetUserInfo: {
		i18n.LangZh: "获取用户信息失败",
		i18n.LangEn: "Failed to obtain user information",
	},
	ErrorConnectTimeout: {
		i18n.LangZh: "连接超时",
		i18n.LangEn: "Connection timed out",
	},
	ErrorWrongPassword: {
		i18n.LangZh: "密码错误",
		i18n.LangEn: "wrong password",
	},
	ErrorInvalidParams1: {
		i18n.LangZh: "请求参数错误:%v",
		i18n.LangEn: "Request parameter error :%v",
	},
	ErrorUserPhoneExist: {
		i18n.LangZh: "用户手机号已存在",
		i18n.LangEn: "Duplicate user phone number",
	},
	ErrorConnectionDB: {
		i18n.LangZh: "测试连接数据库失败",
		i18n.LangEn: "Test connection to the database failed",
	},
	ErrorOpenDB: {
		i18n.LangZh: "打开数据库失败",
		i18n.LangEn: "Failed to open database",
	},
	ErrorDuplicateLabelKey: {
		i18n.LangZh: "标签键不能重复",
		i18n.LangEn: "Label key cannot be repeated",
	},
	ErrorDbNotExist: {
		i18n.LangZh: "数据库不存在",
		i18n.LangEn: "Database does not exist",
	},
	ErrorConnectionDB1: {
		i18n.LangZh: "数据库连接失败%v",
		i18n.LangEn: "Test connection to the database failed %v",
	},
	ErrorTagNotExist: {
		i18n.LangZh: "标签不存在",
		i18n.LangEn: "Label does not exist",
	},
	ErrorAction: {
		i18n.LangZh: "action 不存在",
		i18n.LangEn: "action does not exist",
	},
	ErrorDbHostIPExist: {
		i18n.LangZh: "IP 端口已存在",
		i18n.LangEn: "IP Port exist",
	},
	ErrorTagExist: {
		i18n.LangZh: "标签重复",
		i18n.LangEn: "Duplicate label",
	},
	ErrorUserNickname: {
		i18n.LangZh: "用户昵称不合法",
		i18n.LangEn: "User nickname is illegal",
	},
	ErrorAccountIsNotAvailable: {
		i18n.LangZh: "该账号不可用,请联系管理员",
		i18n.LangEn: "The account is not available, please contact the administrator",
	},
	ErrorLoginInvalid: {
		i18n.LangZh: "登录失效",
		i18n.LangEn: "Login invalid",
	},
	ErrorTestSendEmail: {
		i18n.LangZh: "测试发送邮件失败",
		i18n.LangEn: "Test sending mail failed",
	},
	ErrorSystemDoesNotOpenRegistration: {
		i18n.LangZh: "当前系统未开放注册功能",
		i18n.LangEn: "The current system does not open the registration function",
	},
	ErrorSysBasicNotExist: {
		i18n.LangZh: "该配置不存在",
		i18n.LangEn: "The configuration does not exist",
	},
	ErrorRecipientNotExist: {
		i18n.LangZh: "收件人不存在",
		i18n.LangEn: "Recipient does not exist",
	},
	ErrorParsingFailed: {
		i18n.LangZh: "解析失败",
		i18n.LangEn: "Parsing failed",
	},
	ErrorUpdateFailed: {
		i18n.LangZh: "更新密码类型失败",
		i18n.LangEn: "Failed to update password type",
	},
	ErrorCheckUserNameFailed: {
		i18n.LangZh: "检查用户失败",
		i18n.LangEn: "Check user failed",
	},
	ErrorDBInfo: {
		i18n.LangZh: "数据库错误:%v",
		i18n.LangEn: "Database err is :%v",
	},
	ErrorSsh: {
		i18n.LangZh: "ssh 连接失败 err : %v",
		i18n.LangEn: "ssh connection failed err is :%v",
	},
	ErrorGsOm: {
		i18n.LangZh: "连接失败 : %v",
		i18n.LangEn: "connection failed err is :%v",
	},

	ErrorTestContent: {
		i18n.LangZh: "连接服务器失败",
		i18n.LangEn: "connection failure",
	},
	ErrorUpSteamErr: {
		i18n.LangZh: "上游库选择有误",
		i18n.LangEn: "Wrong upstream library selection",
	},
	ErrorPublisher: {
		i18n.LangZh: "发行商不同",
		i18n.LangEn: "Different publishers",
	},
	ErrorPlatform: {
		i18n.LangZh: "平台不同",
		i18n.LangEn: "Different platforms",
	},
	ErrorVersion: {
		i18n.LangZh: "主机版本不同",
		i18n.LangEn: "Different version",
	},
	ErrorDbVersion: {
		i18n.LangZh: "暂不支持该数据库版本",
		i18n.LangEn: "This database version is not currently supported",
	},
	ErrorIPIsExist: {
		i18n.LangZh: "Host IP重复",
		i18n.LangEn: "Duplicate Host IP",
	},
	ErrorDbNumError: {
		i18n.LangZh: "数据库节点不能超过9个",
		i18n.LangEn: "No more than 9 database nodes",
	},
	ErrorWhetherPython: {
		i18n.LangZh: "服务器中没有python依赖",
		i18n.LangEn: "there is no python environment in the environment",
	},
	ErrorWhetherLibaio: {
		i18n.LangZh: "服务器中没有Libaio依赖",
		i18n.LangEn: "No Libaio package in the environment",
	},
	ErrorIsNotExistCluster: {
		i18n.LangZh: "没有该集群",
		i18n.LangEn: "No such cluster",
	},
	ErrorInstallIngDoNotDel: {
		i18n.LangZh: "安装中不允许操作",
		i18n.LangEn: "Operation not allowed during installation",
	},
	ErrorNodeIsNotExist: {
		i18n.LangZh: "节点不存在",
		i18n.LangEn: "Node does not exist",
	},
	ErrorFailedOperateTimeOut: {
		i18n.LangZh: "用户操作时间过长，已超时",
		i18n.LangEn: "User operation time is too long and has timed out",
	},
	ErrorIsNotExistUnzip: {
		i18n.LangZh: "unzip 不存在,无法解压",
		i18n.LangEn: "unzip does not exist and cannot be unzipped",
	},
	ErrorStartMogHaFail: {
		i18n.LangZh: "启动mogHA失败",
		i18n.LangEn: "Failed to start mogHA",
	},
	ErrorZoneSame: {
		i18n.LangZh: "zone 不能重复",
		i18n.LangEn: "zone cannot be repeated",
	},
	ErrorVipFail: {
		i18n.LangEn: "Vip must be the same in the same region",
		i18n.LangZh: "相同地域下Vip必须相同",
	},
	ErrorPingListFail: {
		i18n.LangZh: "相同地域下PingList必须相同",
		i18n.LangEn: "PingList must be the same in the same region",
	},
	ErrorArpingFail: {
		i18n.LangZh: "相同地域下Arping必须相同",
		i18n.LangEn: "Arping must be the same in the same region",
	},
	ErrorNotSameCluster: {
		i18n.LangZh: "添加节点不在同一集群",
		i18n.LangEn: "The added node is not in the same cluster",
	},
	ErrorPublisherFail: {
		i18n.LangZh: "获取系统发行商错误",
		i18n.LangEn: "Get system publisher error",
	},
	ErrorClusterUnavailable: {
		i18n.LangZh: "集群不可用",
		i18n.LangEn: "Cluster unavailable",
	},
	ErrorClusterNotPrimary: {
		i18n.LangZh: "添加集群中没有主库",
		i18n.LangEn: "There is no main library in the add cluster",
	},
	ErrorCascadeStandby: {
		i18n.LangZh: "级联备库不能单独在一个zone",
		i18n.LangEn: "The cascaded standby database cannot be in a zone alone",
	},
	ErrorZoneHavePri: {
		i18n.LangZh: "该zone下未找到级联备库的上游节点",
		i18n.LangEn: "The upstream node of the cascaded standby database was not found in this zone",
	},
	ErrorOperatingTimeout: {
		i18n.LangZh: "操作超时",
		i18n.LangEn: "operation timeout",
	},
}
