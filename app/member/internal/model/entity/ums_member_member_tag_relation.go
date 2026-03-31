// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// UmsMemberMemberTagRelation is the golang structure for table ums_member_member_tag_relation.
type UmsMemberMemberTagRelation struct {
	Id       int64 `json:"id"       orm:"id"        description:""` //
	MemberId int64 `json:"memberId" orm:"member_id" description:""` //
	TagId    int64 `json:"tagId"    orm:"tag_id"    description:""` //
}
