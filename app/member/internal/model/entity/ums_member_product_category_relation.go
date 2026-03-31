// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// UmsMemberProductCategoryRelation is the golang structure for table ums_member_product_category_relation.
type UmsMemberProductCategoryRelation struct {
	Id                int64 `json:"id"                orm:"id"                  description:""` //
	MemberId          int64 `json:"memberId"          orm:"member_id"           description:""` //
	ProductCategoryId int64 `json:"productCategoryId" orm:"product_category_id" description:""` //
}
