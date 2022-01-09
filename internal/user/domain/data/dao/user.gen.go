// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"github.com/dapr-ddd-action/internal/user/domain/data/po"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"
)

func newUser(db *gorm.DB) user {
	_user := user{}

	_user.userDo.UseDB(db)
	_user.userDo.UseModel(&po.User{})

	tableName := _user.userDo.TableName()
	_user.ALL = field.NewField(tableName, "*")
	_user.ID = field.NewInt64(tableName, "id")
	_user.UserID = field.NewInt64(tableName, "user_id")
	_user.Age = field.NewInt32(tableName, "age")
	_user.UserName = field.NewString(tableName, "user_name")
	_user.Password = field.NewString(tableName, "password")
	_user.Email = field.NewString(tableName, "email")
	_user.Phone = field.NewString(tableName, "phone")
	_user.RoleName = field.NewString(tableName, "role_name")
	_user.CreatedAt = field.NewTime(tableName, "created_at")
	_user.UpdatedAt = field.NewTime(tableName, "updated_at")

	_user.fillFieldMap()

	return _user
}

type user struct {
	userDo userDo

	ALL       field.Field
	ID        field.Int64
	UserID    field.Int64
	Age       field.Int32
	UserName  field.String
	Password  field.String
	Email     field.String
	Phone     field.String
	RoleName  field.String
	CreatedAt field.Time
	UpdatedAt field.Time

	fieldMap map[string]field.Expr
}

func (u user) As(alias string) *user {
	u.userDo.DO = *(u.userDo.As(alias).(*gen.DO))

	u.ALL = field.NewField(alias, "*")
	u.ID = field.NewInt64(alias, "id")
	u.UserID = field.NewInt64(alias, "user_id")
	u.Age = field.NewInt32(alias, "age")
	u.UserName = field.NewString(alias, "user_name")
	u.Password = field.NewString(alias, "password")
	u.Email = field.NewString(alias, "email")
	u.Phone = field.NewString(alias, "phone")
	u.RoleName = field.NewString(alias, "role_name")
	u.CreatedAt = field.NewTime(alias, "created_at")
	u.UpdatedAt = field.NewTime(alias, "updated_at")

	u.fillFieldMap()

	return &u
}

func (u *user) WithContext(ctx context.Context) *userDo { return u.userDo.WithContext(ctx) }

func (u user) TableName() string { return u.userDo.TableName() }

func (u *user) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := u.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	return _f.(field.OrderExpr), true
}

func (u *user) fillFieldMap() {
	u.fieldMap = make(map[string]field.Expr, 10)
	u.fieldMap["id"] = u.ID
	u.fieldMap["user_id"] = u.UserID
	u.fieldMap["age"] = u.Age
	u.fieldMap["user_name"] = u.UserName
	u.fieldMap["password"] = u.Password
	u.fieldMap["email"] = u.Email
	u.fieldMap["phone"] = u.Phone
	u.fieldMap["role_name"] = u.RoleName
	u.fieldMap["created_at"] = u.CreatedAt
	u.fieldMap["updated_at"] = u.UpdatedAt
}

func (u user) clone(db *gorm.DB) user {
	u.userDo.ReplaceDB(db)
	return u
}

type userDo struct{ gen.DO }

func (u userDo) Debug() *userDo {
	return u.withDO(u.DO.Debug())
}

func (u userDo) WithContext(ctx context.Context) *userDo {
	return u.withDO(u.DO.WithContext(ctx))
}

func (u userDo) Clauses(conds ...clause.Expression) *userDo {
	return u.withDO(u.DO.Clauses(conds...))
}

func (u userDo) Not(conds ...gen.Condition) *userDo {
	return u.withDO(u.DO.Not(conds...))
}

func (u userDo) Or(conds ...gen.Condition) *userDo {
	return u.withDO(u.DO.Or(conds...))
}

func (u userDo) Select(conds ...field.Expr) *userDo {
	return u.withDO(u.DO.Select(conds...))
}

func (u userDo) Where(conds ...gen.Condition) *userDo {
	return u.withDO(u.DO.Where(conds...))
}

func (u userDo) Order(conds ...field.Expr) *userDo {
	return u.withDO(u.DO.Order(conds...))
}

func (u userDo) Distinct(cols ...field.Expr) *userDo {
	return u.withDO(u.DO.Distinct(cols...))
}

func (u userDo) Omit(cols ...field.Expr) *userDo {
	return u.withDO(u.DO.Omit(cols...))
}

func (u userDo) Join(table schema.Tabler, on ...field.Expr) *userDo {
	return u.withDO(u.DO.Join(table, on...))
}

func (u userDo) LeftJoin(table schema.Tabler, on ...field.Expr) *userDo {
	return u.withDO(u.DO.LeftJoin(table, on...))
}

func (u userDo) RightJoin(table schema.Tabler, on ...field.Expr) *userDo {
	return u.withDO(u.DO.RightJoin(table, on...))
}

func (u userDo) Group(cols ...field.Expr) *userDo {
	return u.withDO(u.DO.Group(cols...))
}

func (u userDo) Having(conds ...gen.Condition) *userDo {
	return u.withDO(u.DO.Having(conds...))
}

func (u userDo) Limit(limit int) *userDo {
	return u.withDO(u.DO.Limit(limit))
}

func (u userDo) Offset(offset int) *userDo {
	return u.withDO(u.DO.Offset(offset))
}

func (u userDo) Scopes(funcs ...func(gen.Dao) gen.Dao) *userDo {
	return u.withDO(u.DO.Scopes(funcs...))
}

func (u userDo) Unscoped() *userDo {
	return u.withDO(u.DO.Unscoped())
}

func (u userDo) Create(values ...*po.User) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Create(values)
}

func (u userDo) CreateInBatches(values []*po.User, batchSize int) error {
	return u.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (u userDo) Save(values ...*po.User) error {
	if len(values) == 0 {
		return nil
	}
	return u.DO.Save(values)
}

func (u userDo) First() (*po.User, error) {
	if result, err := u.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*po.User), nil
	}
}

func (u userDo) Take() (*po.User, error) {
	if result, err := u.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*po.User), nil
	}
}

func (u userDo) Last() (*po.User, error) {
	if result, err := u.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*po.User), nil
	}
}

func (u userDo) Find() ([]*po.User, error) {
	result, err := u.DO.Find()
	return result.([]*po.User), err
}

func (u userDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*po.User, err error) {
	buf := make([]*po.User, 0, batchSize)
	err = u.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (u userDo) FindInBatches(result *[]*po.User, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return u.DO.FindInBatches(result, batchSize, fc)
}

func (u userDo) Attrs(attrs ...field.AssignExpr) *userDo {
	return u.withDO(u.DO.Attrs(attrs...))
}

func (u userDo) Assign(attrs ...field.AssignExpr) *userDo {
	return u.withDO(u.DO.Assign(attrs...))
}

func (u userDo) Joins(field field.RelationField) *userDo {
	return u.withDO(u.DO.Joins(field))
}

func (u userDo) Preload(field field.RelationField) *userDo {
	return u.withDO(u.DO.Preload(field))
}

func (u userDo) FirstOrInit() (*po.User, error) {
	if result, err := u.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*po.User), nil
	}
}

func (u userDo) FirstOrCreate() (*po.User, error) {
	if result, err := u.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*po.User), nil
	}
}

func (u userDo) FindByPage(offset int, limit int) (result []*po.User, count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	if limit <= 0 {
		return
	}

	result, err = u.Offset(offset).Limit(limit).Find()
	return
}

func (u userDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = u.Count()
	if err != nil {
		return
	}

	err = u.Offset(offset).Limit(limit).Scan(result)
	return
}

func (u *userDo) withDO(do gen.Dao) *userDo {
	u.DO = *do.(*gen.DO)
	return u
}
