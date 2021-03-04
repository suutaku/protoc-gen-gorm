// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: example/feature_demo/demo_multi_file.proto

/*
Package example is a generated protocol buffer package.

It is generated from these files:
	example/feature_demo/demo_multi_file.proto
	example/feature_demo/demo_types.proto
	example/feature_demo/demo_service.proto
	example/feature_demo/demo_multi_file_service.proto

It has these top-level messages:
	ExternalChild
	BlogPost
	TestTypes
	TypeWithID
	MultiaccountTypeWithID
	MultiaccountTypeWithoutID
	APIOnlyType
	PrimaryUUIDType
	PrimaryStringType
	TestTag
	TestAssocHandlerDefault
	TestAssocHandlerReplace
	TestAssocHandlerClear
	TestAssocHandlerAppend
	TestTagAssociation
	PrimaryIncluded
	IntPoint
	CreateIntPointRequest
	CreateIntPointResponse
	ReadIntPointRequest
	ReadIntPointResponse
	UpdateIntPointRequest
	UpdateIntPointResponse
	UpdateSetIntPointRequest
	UpdateSetIntPointResponse
	DeleteIntPointRequest
	DeleteIntPointsRequest
	DeleteIntPointResponse
	ListIntPointResponse
	ListSomethingResponse
	Something
	ListIntPointRequest
	Circle
	ListCircleRequest
	ListCircleResponse
	ReadAccountRequest
	ReadBlogPostsResponse
*/
package example

import context "context"
import fmt "fmt"

import errors1 "github.com/suutaku/protoc-gen-gorm/errors"
import field_mask1 "google.golang.org/genproto/protobuf/field_mask"
import gateway1 "github.com/infobloxopen/atlas-app-toolkit/gateway"
import go_uuid1 "github.com/satori/go.uuid"
import gorm1 "github.com/jinzhu/gorm"
import gorm2 "github.com/infobloxopen/atlas-app-toolkit/gorm"

import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = fmt.Errorf
var _ = math.Inf

type ExternalChildORM struct {
	Id                  string
	PrimaryIncludedId   *go_uuid1.UUID
	PrimaryStringTypeId *string
	PrimaryUUIDTypeId   *go_uuid1.UUID
}

// TableName overrides the default tablename generated by GORM
func (ExternalChildORM) TableName() string {
	return "external_children"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *ExternalChild) ToORM(ctx context.Context) (ExternalChildORM, error) {
	to := ExternalChildORM{}
	var err error
	if prehook, ok := interface{}(m).(ExternalChildWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	if posthook, ok := interface{}(m).(ExternalChildWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *ExternalChildORM) ToPB(ctx context.Context) (ExternalChild, error) {
	to := ExternalChild{}
	var err error
	if prehook, ok := interface{}(m).(ExternalChildWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	if posthook, ok := interface{}(m).(ExternalChildWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type ExternalChild the arg will be the target, the caller the one being converted from

// ExternalChildBeforeToORM called before default ToORM code
type ExternalChildWithBeforeToORM interface {
	BeforeToORM(context.Context, *ExternalChildORM) error
}

// ExternalChildAfterToORM called after default ToORM code
type ExternalChildWithAfterToORM interface {
	AfterToORM(context.Context, *ExternalChildORM) error
}

// ExternalChildBeforeToPB called before default ToPB code
type ExternalChildWithBeforeToPB interface {
	BeforeToPB(context.Context, *ExternalChild) error
}

// ExternalChildAfterToPB called after default ToPB code
type ExternalChildWithAfterToPB interface {
	AfterToPB(context.Context, *ExternalChild) error
}

type BlogPostORM struct {
	Author string
	Id     uint64
	Title  string
}

// TableName overrides the default tablename generated by GORM
func (BlogPostORM) TableName() string {
	return "blog_posts"
}

// ToORM runs the BeforeToORM hook if present, converts the fields of this
// object to ORM format, runs the AfterToORM hook, then returns the ORM object
func (m *BlogPost) ToORM(ctx context.Context) (BlogPostORM, error) {
	to := BlogPostORM{}
	var err error
	if prehook, ok := interface{}(m).(BlogPostWithBeforeToORM); ok {
		if err = prehook.BeforeToORM(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.Title = m.Title
	to.Author = m.Author
	if posthook, ok := interface{}(m).(BlogPostWithAfterToORM); ok {
		err = posthook.AfterToORM(ctx, &to)
	}
	return to, err
}

// ToPB runs the BeforeToPB hook if present, converts the fields of this
// object to PB format, runs the AfterToPB hook, then returns the PB object
func (m *BlogPostORM) ToPB(ctx context.Context) (BlogPost, error) {
	to := BlogPost{}
	var err error
	if prehook, ok := interface{}(m).(BlogPostWithBeforeToPB); ok {
		if err = prehook.BeforeToPB(ctx, &to); err != nil {
			return to, err
		}
	}
	to.Id = m.Id
	to.Title = m.Title
	to.Author = m.Author
	if posthook, ok := interface{}(m).(BlogPostWithAfterToPB); ok {
		err = posthook.AfterToPB(ctx, &to)
	}
	return to, err
}

// The following are interfaces you can implement for special behavior during ORM/PB conversions
// of type BlogPost the arg will be the target, the caller the one being converted from

// BlogPostBeforeToORM called before default ToORM code
type BlogPostWithBeforeToORM interface {
	BeforeToORM(context.Context, *BlogPostORM) error
}

// BlogPostAfterToORM called after default ToORM code
type BlogPostWithAfterToORM interface {
	AfterToORM(context.Context, *BlogPostORM) error
}

// BlogPostBeforeToPB called before default ToPB code
type BlogPostWithBeforeToPB interface {
	BeforeToPB(context.Context, *BlogPost) error
}

// BlogPostAfterToPB called after default ToPB code
type BlogPostWithAfterToPB interface {
	AfterToPB(context.Context, *BlogPost) error
}

// DefaultCreateExternalChild executes a basic gorm create call
func DefaultCreateExternalChild(ctx context.Context, in *ExternalChild, db *gorm1.DB) (*ExternalChild, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ExternalChildORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ExternalChildORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type ExternalChildORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type ExternalChildORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm1.DB) error
}

// DefaultReadExternalChild executes a basic gorm read call
func DefaultReadExternalChild(ctx context.Context, in *ExternalChild, db *gorm1.DB) (*ExternalChild, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if ormObj.Id == "" {
		return nil, errors1.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(ExternalChildORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if db, err = gorm2.ApplyFieldSelection(ctx, db, nil, &ExternalChildORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ExternalChildORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := ExternalChildORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(ExternalChildORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type ExternalChildORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type ExternalChildORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type ExternalChildORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm1.DB) error
}

func DefaultDeleteExternalChild(ctx context.Context, in *ExternalChild, db *gorm1.DB) error {
	if in == nil {
		return errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Id == "" {
		return errors1.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(ExternalChildORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&ExternalChildORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(ExternalChildORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type ExternalChildORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type ExternalChildORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm1.DB) error
}

func DefaultDeleteExternalChildSet(ctx context.Context, in []*ExternalChild, db *gorm1.DB) error {
	if in == nil {
		return errors1.NilArgumentError
	}
	var err error
	keys := []string{}
	for _, obj := range in {
		ormObj, err := obj.ToORM(ctx)
		if err != nil {
			return err
		}
		if ormObj.Id == "" {
			return errors1.EmptyIdError
		}
		keys = append(keys, ormObj.Id)
	}
	if hook, ok := (interface{}(&ExternalChildORM{})).(ExternalChildORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("id in (?)", keys).Delete(&ExternalChildORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&ExternalChildORM{})).(ExternalChildORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type ExternalChildORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*ExternalChild, *gorm1.DB) (*gorm1.DB, error)
}
type ExternalChildORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*ExternalChild, *gorm1.DB) error
}

// DefaultStrictUpdateExternalChild clears / replaces / appends first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateExternalChild(ctx context.Context, in *ExternalChild, db *gorm1.DB) (*ExternalChild, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateExternalChild")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	var count int64
	lockedRow := &ExternalChildORM{}
	count = db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow).RowsAffected
	if hook, ok := interface{}(&ormObj).(ExternalChildORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(ExternalChildORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ExternalChildORMWithAfterStrictUpdateSave); ok {
		if err = hook.AfterStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		err = gateway1.SetCreated(ctx, "")
	}
	return &pbResponse, err
}

type ExternalChildORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type ExternalChildORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type ExternalChildORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm1.DB) error
}

// DefaultPatchExternalChild executes a basic gorm update call with patch behavior
func DefaultPatchExternalChild(ctx context.Context, in *ExternalChild, updateMask *field_mask1.FieldMask, db *gorm1.DB) (*ExternalChild, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	var pbObj ExternalChild
	var err error
	if hook, ok := interface{}(&pbObj).(ExternalChildWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := DefaultReadExternalChild(ctx, &ExternalChild{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(ExternalChildWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskExternalChild(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(ExternalChildWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateExternalChild(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(ExternalChildWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type ExternalChildWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *ExternalChild, *field_mask1.FieldMask, *gorm1.DB) (*gorm1.DB, error)
}
type ExternalChildWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *ExternalChild, *field_mask1.FieldMask, *gorm1.DB) (*gorm1.DB, error)
}
type ExternalChildWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *ExternalChild, *field_mask1.FieldMask, *gorm1.DB) (*gorm1.DB, error)
}
type ExternalChildWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *ExternalChild, *field_mask1.FieldMask, *gorm1.DB) error
}

// DefaultPatchSetExternalChild executes a bulk gorm update call with patch behavior
func DefaultPatchSetExternalChild(ctx context.Context, objects []*ExternalChild, updateMasks []*field_mask1.FieldMask, db *gorm1.DB) ([]*ExternalChild, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors1.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*ExternalChild, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchExternalChild(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskExternalChild patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskExternalChild(ctx context.Context, patchee *ExternalChild, patcher *ExternalChild, updateMask *field_mask1.FieldMask, prefix string, db *gorm1.DB) (*ExternalChild, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors1.NilArgumentError
	}
	var err error
	for _, f := range updateMask.Paths {
		if f == prefix+"Id" {
			patchee.Id = patcher.Id
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListExternalChild executes a gorm list call
func DefaultListExternalChild(ctx context.Context, db *gorm1.DB) ([]*ExternalChild, error) {
	in := ExternalChild{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ExternalChildORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm2.ApplyCollectionOperators(ctx, db, &ExternalChildORM{}, &ExternalChild{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ExternalChildORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("id")
	ormResponse := []ExternalChildORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(ExternalChildORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*ExternalChild{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type ExternalChildORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type ExternalChildORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type ExternalChildORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm1.DB, *[]ExternalChildORM) error
}

// DefaultCreateBlogPost executes a basic gorm create call
func DefaultCreateBlogPost(ctx context.Context, in *BlogPost, db *gorm1.DB) (*BlogPost, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(BlogPostORMWithBeforeCreate_); ok {
		if db, err = hook.BeforeCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Create(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(BlogPostORMWithAfterCreate_); ok {
		if err = hook.AfterCreate_(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	return &pbResponse, err
}

type BlogPostORMWithBeforeCreate_ interface {
	BeforeCreate_(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type BlogPostORMWithAfterCreate_ interface {
	AfterCreate_(context.Context, *gorm1.DB) error
}

// DefaultReadBlogPost executes a basic gorm read call
func DefaultReadBlogPost(ctx context.Context, in *BlogPost, db *gorm1.DB) (*BlogPost, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if ormObj.Id == 0 {
		return nil, errors1.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(BlogPostORMWithBeforeReadApplyQuery); ok {
		if db, err = hook.BeforeReadApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	if db, err = gorm2.ApplyFieldSelection(ctx, db, nil, &BlogPostORM{}); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(BlogPostORMWithBeforeReadFind); ok {
		if db, err = hook.BeforeReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	ormResponse := BlogPostORM{}
	if err = db.Where(&ormObj).First(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormResponse).(BlogPostORMWithAfterReadFind); ok {
		if err = hook.AfterReadFind(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormResponse.ToPB(ctx)
	return &pbResponse, err
}

type BlogPostORMWithBeforeReadApplyQuery interface {
	BeforeReadApplyQuery(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type BlogPostORMWithBeforeReadFind interface {
	BeforeReadFind(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type BlogPostORMWithAfterReadFind interface {
	AfterReadFind(context.Context, *gorm1.DB) error
}

func DefaultDeleteBlogPost(ctx context.Context, in *BlogPost, db *gorm1.DB) error {
	if in == nil {
		return errors1.NilArgumentError
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return err
	}
	if ormObj.Id == 0 {
		return errors1.EmptyIdError
	}
	if hook, ok := interface{}(&ormObj).(BlogPostORMWithBeforeDelete_); ok {
		if db, err = hook.BeforeDelete_(ctx, db); err != nil {
			return err
		}
	}
	err = db.Where(&ormObj).Delete(&BlogPostORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := interface{}(&ormObj).(BlogPostORMWithAfterDelete_); ok {
		err = hook.AfterDelete_(ctx, db)
	}
	return err
}

type BlogPostORMWithBeforeDelete_ interface {
	BeforeDelete_(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type BlogPostORMWithAfterDelete_ interface {
	AfterDelete_(context.Context, *gorm1.DB) error
}

func DefaultDeleteBlogPostSet(ctx context.Context, in []*BlogPost, db *gorm1.DB) error {
	if in == nil {
		return errors1.NilArgumentError
	}
	var err error
	keys := []uint64{}
	for _, obj := range in {
		ormObj, err := obj.ToORM(ctx)
		if err != nil {
			return err
		}
		if ormObj.Id == 0 {
			return errors1.EmptyIdError
		}
		keys = append(keys, ormObj.Id)
	}
	if hook, ok := (interface{}(&BlogPostORM{})).(BlogPostORMWithBeforeDeleteSet); ok {
		if db, err = hook.BeforeDeleteSet(ctx, in, db); err != nil {
			return err
		}
	}
	err = db.Where("id in (?)", keys).Delete(&BlogPostORM{}).Error
	if err != nil {
		return err
	}
	if hook, ok := (interface{}(&BlogPostORM{})).(BlogPostORMWithAfterDeleteSet); ok {
		err = hook.AfterDeleteSet(ctx, in, db)
	}
	return err
}

type BlogPostORMWithBeforeDeleteSet interface {
	BeforeDeleteSet(context.Context, []*BlogPost, *gorm1.DB) (*gorm1.DB, error)
}
type BlogPostORMWithAfterDeleteSet interface {
	AfterDeleteSet(context.Context, []*BlogPost, *gorm1.DB) error
}

// DefaultStrictUpdateBlogPost clears / replaces / appends first level 1:many children and then executes a gorm update call
func DefaultStrictUpdateBlogPost(ctx context.Context, in *BlogPost, db *gorm1.DB) (*BlogPost, error) {
	if in == nil {
		return nil, fmt.Errorf("Nil argument to DefaultStrictUpdateBlogPost")
	}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	var count int64
	lockedRow := &BlogPostORM{}
	count = db.Model(&ormObj).Set("gorm:query_option", "FOR UPDATE").Where("id=?", ormObj.Id).First(lockedRow).RowsAffected
	if hook, ok := interface{}(&ormObj).(BlogPostORMWithBeforeStrictUpdateCleanup); ok {
		if db, err = hook.BeforeStrictUpdateCleanup(ctx, db); err != nil {
			return nil, err
		}
	}
	if hook, ok := interface{}(&ormObj).(BlogPostORMWithBeforeStrictUpdateSave); ok {
		if db, err = hook.BeforeStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	if err = db.Save(&ormObj).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(BlogPostORMWithAfterStrictUpdateSave); ok {
		if err = hook.AfterStrictUpdateSave(ctx, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := ormObj.ToPB(ctx)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		err = gateway1.SetCreated(ctx, "")
	}
	return &pbResponse, err
}

type BlogPostORMWithBeforeStrictUpdateCleanup interface {
	BeforeStrictUpdateCleanup(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type BlogPostORMWithBeforeStrictUpdateSave interface {
	BeforeStrictUpdateSave(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type BlogPostORMWithAfterStrictUpdateSave interface {
	AfterStrictUpdateSave(context.Context, *gorm1.DB) error
}

// DefaultPatchBlogPost executes a basic gorm update call with patch behavior
func DefaultPatchBlogPost(ctx context.Context, in *BlogPost, updateMask *field_mask1.FieldMask, db *gorm1.DB) (*BlogPost, error) {
	if in == nil {
		return nil, errors1.NilArgumentError
	}
	var pbObj BlogPost
	var err error
	if hook, ok := interface{}(&pbObj).(BlogPostWithBeforePatchRead); ok {
		if db, err = hook.BeforePatchRead(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbReadRes, err := DefaultReadBlogPost(ctx, &BlogPost{Id: in.GetId()}, db)
	if err != nil {
		return nil, err
	}
	pbObj = *pbReadRes
	if hook, ok := interface{}(&pbObj).(BlogPostWithBeforePatchApplyFieldMask); ok {
		if db, err = hook.BeforePatchApplyFieldMask(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	if _, err := DefaultApplyFieldMaskBlogPost(ctx, &pbObj, in, updateMask, "", db); err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&pbObj).(BlogPostWithBeforePatchSave); ok {
		if db, err = hook.BeforePatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	pbResponse, err := DefaultStrictUpdateBlogPost(ctx, &pbObj, db)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(pbResponse).(BlogPostWithAfterPatchSave); ok {
		if err = hook.AfterPatchSave(ctx, in, updateMask, db); err != nil {
			return nil, err
		}
	}
	return pbResponse, nil
}

type BlogPostWithBeforePatchRead interface {
	BeforePatchRead(context.Context, *BlogPost, *field_mask1.FieldMask, *gorm1.DB) (*gorm1.DB, error)
}
type BlogPostWithBeforePatchApplyFieldMask interface {
	BeforePatchApplyFieldMask(context.Context, *BlogPost, *field_mask1.FieldMask, *gorm1.DB) (*gorm1.DB, error)
}
type BlogPostWithBeforePatchSave interface {
	BeforePatchSave(context.Context, *BlogPost, *field_mask1.FieldMask, *gorm1.DB) (*gorm1.DB, error)
}
type BlogPostWithAfterPatchSave interface {
	AfterPatchSave(context.Context, *BlogPost, *field_mask1.FieldMask, *gorm1.DB) error
}

// DefaultPatchSetBlogPost executes a bulk gorm update call with patch behavior
func DefaultPatchSetBlogPost(ctx context.Context, objects []*BlogPost, updateMasks []*field_mask1.FieldMask, db *gorm1.DB) ([]*BlogPost, error) {
	if len(objects) != len(updateMasks) {
		return nil, fmt.Errorf(errors1.BadRepeatedFieldMaskTpl, len(updateMasks), len(objects))
	}

	results := make([]*BlogPost, 0, len(objects))
	for i, patcher := range objects {
		pbResponse, err := DefaultPatchBlogPost(ctx, patcher, updateMasks[i], db)
		if err != nil {
			return nil, err
		}

		results = append(results, pbResponse)
	}

	return results, nil
}

// DefaultApplyFieldMaskBlogPost patches an pbObject with patcher according to a field mask.
func DefaultApplyFieldMaskBlogPost(ctx context.Context, patchee *BlogPost, patcher *BlogPost, updateMask *field_mask1.FieldMask, prefix string, db *gorm1.DB) (*BlogPost, error) {
	if patcher == nil {
		return nil, nil
	} else if patchee == nil {
		return nil, errors1.NilArgumentError
	}
	var err error
	for _, f := range updateMask.Paths {
		if f == prefix+"Id" {
			patchee.Id = patcher.Id
			continue
		}
		if f == prefix+"Title" {
			patchee.Title = patcher.Title
			continue
		}
		if f == prefix+"Author" {
			patchee.Author = patcher.Author
			continue
		}
	}
	if err != nil {
		return nil, err
	}
	return patchee, nil
}

// DefaultListBlogPost executes a gorm list call
func DefaultListBlogPost(ctx context.Context, db *gorm1.DB) ([]*BlogPost, error) {
	in := BlogPost{}
	ormObj, err := in.ToORM(ctx)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(BlogPostORMWithBeforeListApplyQuery); ok {
		if db, err = hook.BeforeListApplyQuery(ctx, db); err != nil {
			return nil, err
		}
	}
	db, err = gorm2.ApplyCollectionOperators(ctx, db, &BlogPostORM{}, &BlogPost{}, nil, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(BlogPostORMWithBeforeListFind); ok {
		if db, err = hook.BeforeListFind(ctx, db); err != nil {
			return nil, err
		}
	}
	db = db.Where(&ormObj)
	db = db.Order("id")
	ormResponse := []BlogPostORM{}
	if err := db.Find(&ormResponse).Error; err != nil {
		return nil, err
	}
	if hook, ok := interface{}(&ormObj).(BlogPostORMWithAfterListFind); ok {
		if err = hook.AfterListFind(ctx, db, &ormResponse); err != nil {
			return nil, err
		}
	}
	pbResponse := []*BlogPost{}
	for _, responseEntry := range ormResponse {
		temp, err := responseEntry.ToPB(ctx)
		if err != nil {
			return nil, err
		}
		pbResponse = append(pbResponse, &temp)
	}
	return pbResponse, nil
}

type BlogPostORMWithBeforeListApplyQuery interface {
	BeforeListApplyQuery(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type BlogPostORMWithBeforeListFind interface {
	BeforeListFind(context.Context, *gorm1.DB) (*gorm1.DB, error)
}
type BlogPostORMWithAfterListFind interface {
	AfterListFind(context.Context, *gorm1.DB, *[]BlogPostORM) error
}