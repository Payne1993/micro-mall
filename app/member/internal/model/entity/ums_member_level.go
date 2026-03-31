// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// UmsMemberLevel is the golang structure for table ums_member_level.
type UmsMemberLevel struct {
	Id                    int64   `json:"id"                    orm:"id"                      description:""`                   //
	Name                  string  `json:"name"                  orm:"name"                    description:""`                   //
	GrowthPoint           int     `json:"growthPoint"           orm:"growth_point"            description:""`                   //
	DefaultStatus         int     `json:"defaultStatus"         orm:"default_status"          description:"是否为默认等级：0->不是；1->是"` // 是否为默认等级：0->不是；1->是
	FreeFreightPoint      float64 `json:"freeFreightPoint"      orm:"free_freight_point"      description:"免运费标准"`              // 免运费标准
	CommentGrowthPoint    int     `json:"commentGrowthPoint"    orm:"comment_growth_point"    description:"每次评价获取的成长值"`         // 每次评价获取的成长值
	PriviledgeFreeFreight int     `json:"priviledgeFreeFreight" orm:"priviledge_free_freight" description:"是否有免邮特权"`            // 是否有免邮特权
	PriviledgeSignIn      int     `json:"priviledgeSignIn"      orm:"priviledge_sign_in"      description:"是否有签到特权"`            // 是否有签到特权
	PriviledgeComment     int     `json:"priviledgeComment"     orm:"priviledge_comment"      description:"是否有评论获奖励特权"`         // 是否有评论获奖励特权
	PriviledgePromotion   int     `json:"priviledgePromotion"   orm:"priviledge_promotion"    description:"是否有专享活动特权"`          // 是否有专享活动特权
	PriviledgeMemberPrice int     `json:"priviledgeMemberPrice" orm:"priviledge_member_price" description:"是否有会员价格特权"`          // 是否有会员价格特权
	PriviledgeBirthday    int     `json:"priviledgeBirthday"    orm:"priviledge_birthday"     description:"是否有生日特权"`            // 是否有生日特权
	Note                  string  `json:"note"                  orm:"note"                    description:""`                   //
}
