// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// UmsMemberReceiveAddress is the golang structure for table ums_member_receive_address.
type UmsMemberReceiveAddress struct {
	Id            int64  `json:"id"            orm:"id"             description:""`         //
	MemberId      int64  `json:"memberId"      orm:"member_id"      description:""`         //
	Name          string `json:"name"          orm:"name"           description:"收货人名称"`    // 收货人名称
	PhoneNumber   string `json:"phoneNumber"   orm:"phone_number"   description:""`         //
	DefaultStatus int    `json:"defaultStatus" orm:"default_status" description:"是否为默认"`    // 是否为默认
	PostCode      string `json:"postCode"      orm:"post_code"      description:"邮政编码"`     // 邮政编码
	Province      string `json:"province"      orm:"province"       description:"省份/直辖市"`   // 省份/直辖市
	City          string `json:"city"          orm:"city"           description:"城市"`       // 城市
	Region        string `json:"region"        orm:"region"         description:"区"`        // 区
	DetailAddress string `json:"detailAddress" orm:"detail_address" description:"详细地址(街道)"` // 详细地址(街道)
}
