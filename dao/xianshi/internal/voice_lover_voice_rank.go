// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"context"
	"database/sql"
	"errors"
	"reflect"
	"strings"
	"time"

	"github.com/gogf/gf/container/gvar"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
	"github.com/gogf/gf/util/gconv"

	//	"github.com/olaola-chat/slp-proto/model"
	pb "github.com/olaola-chat/slp-proto/gen_pb/db/xianshi"
)

// VoiceLoverVoiceRankDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type VoiceLoverVoiceRankDao struct {
	gmvc.M
	DB      gdb.DB
	Table   string
	Columns voiceLoverVoiceRankColumns
}

// VoiceLoverVoiceRankColumns defines and stores column names for table voice_lover_voice_rank.
type voiceLoverVoiceRankColumns struct {
	ID         string //
	ActivityID string // 活动id
	AudioID    string // 作品id
	LikeNum    string // 点赞数
	CreateTime string //
	UpdateTime string //
	SentAward  string // 是否已发奖励 0-未发放 1-已发放
}

var (
	// VoiceLoverVoiceRank is globally public accessible object for table voice_lover_voice_rank operations.
	VoiceLoverVoiceRank = VoiceLoverVoiceRankDao{
		M:     g.DB("xianshi").Model("voice_lover_voice_rank").Safe(),
		DB:    g.DB("xianshi"),
		Table: "voice_lover_voice_rank",
		Columns: voiceLoverVoiceRankColumns{
			ID:         "id",
			ActivityID: "activity_id",
			AudioID:    "audio_id",
			LikeNum:    "like_num",
			CreateTime: "create_time",
			UpdateTime: "update_time",
			SentAward:  "sent_award",
		},
	}
)

// Ctx is a chaining function, which creates and returns a new DB that is a shallow copy
// of current DB object and with given context in it.
// Note that this returned DB object can be used only once, so do not assign it to
// a global or package variable for long using.
func (d *VoiceLoverVoiceRankDao) Ctx(ctx context.Context) *VoiceLoverVoiceRankDao {
	return &VoiceLoverVoiceRankDao{M: d.M.Ctx(ctx)}
}

// As sets an alias name for current table.
func (d *VoiceLoverVoiceRankDao) As(as string) *VoiceLoverVoiceRankDao {
	return &VoiceLoverVoiceRankDao{M: d.M.As(as)}
}

// TX sets the transaction for current operation.
func (d *VoiceLoverVoiceRankDao) TX(tx *gdb.TX) *VoiceLoverVoiceRankDao {
	return &VoiceLoverVoiceRankDao{M: d.M.TX(tx)}
}

// Master marks the following operation on master node.
func (d *VoiceLoverVoiceRankDao) Master() *VoiceLoverVoiceRankDao {
	return &VoiceLoverVoiceRankDao{M: d.M.Master()}
}

// Slave marks the following operation on slave node.
// Note that it makes sense only if there's any slave node configured.
func (d *VoiceLoverVoiceRankDao) Slave() *VoiceLoverVoiceRankDao {
	return &VoiceLoverVoiceRankDao{M: d.M.Slave()}
}

// Args sets custom arguments for model operation.
func (d *VoiceLoverVoiceRankDao) Args(args ...interface{}) *VoiceLoverVoiceRankDao {
	return &VoiceLoverVoiceRankDao{M: d.M.Args(args...)}
}

// LeftJoin does "LEFT JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").LeftJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").LeftJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *VoiceLoverVoiceRankDao) LeftJoin(table ...string) *VoiceLoverVoiceRankDao {
	return &VoiceLoverVoiceRankDao{M: d.M.LeftJoin(table...)}
}

// RightJoin does "RIGHT JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").RightJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").RightJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *VoiceLoverVoiceRankDao) RightJoin(table ...string) *VoiceLoverVoiceRankDao {
	return &VoiceLoverVoiceRankDao{M: d.M.RightJoin(table...)}
}

// InnerJoin does "INNER JOIN ... ON ..." statement on the model.
// The parameter <table> can be joined table and its joined condition,
// and also with its alias name, like:
// Table("user").InnerJoin("user_detail", "user_detail.uid=user.uid")
// Table("user", "u").InnerJoin("user_detail", "ud", "ud.uid=u.uid")
func (d *VoiceLoverVoiceRankDao) InnerJoin(table ...string) *VoiceLoverVoiceRankDao {
	return &VoiceLoverVoiceRankDao{M: d.M.InnerJoin(table...)}
}

// Fields sets the operation fields of the model, multiple fields joined using char ','.
// The parameter <fieldNamesOrMapStruct> can be type of string/map/*map/struct/*struct.
func (d *VoiceLoverVoiceRankDao) Fields(fieldNamesOrMapStruct ...interface{}) *VoiceLoverVoiceRankDao {
	return &VoiceLoverVoiceRankDao{M: d.M.Fields(fieldNamesOrMapStruct...)}
}

// FieldsEx sets the excluded operation fields of the model, multiple fields joined using char ','.
// The parameter <fieldNamesOrMapStruct> can be type of string/map/*map/struct/*struct.
func (d *VoiceLoverVoiceRankDao) FieldsEx(fieldNamesOrMapStruct ...interface{}) *VoiceLoverVoiceRankDao {
	return &VoiceLoverVoiceRankDao{M: d.M.FieldsEx(fieldNamesOrMapStruct...)}
}

// Option sets the extra operation option for the model.
func (d *VoiceLoverVoiceRankDao) Option(option int) *VoiceLoverVoiceRankDao {
	return &VoiceLoverVoiceRankDao{M: d.M.Option(option)}
}

// OmitEmpty sets OPTION_OMITEMPTY option for the model, which automatically filers
// the data and where attributes for empty values.
func (d *VoiceLoverVoiceRankDao) OmitEmpty() *VoiceLoverVoiceRankDao {
	return &VoiceLoverVoiceRankDao{M: d.M.OmitEmpty()}
}

// Filter marks filtering the fields which does not exist in the fields of the operated table.
func (d *VoiceLoverVoiceRankDao) Filter() *VoiceLoverVoiceRankDao {
	return &VoiceLoverVoiceRankDao{M: d.M.Filter()}
}

// Where sets the condition statement for the model. The parameter <where> can be type of
// string/map/gmap/slice/struct/*struct, etc. Note that, if it's called more than one times,
// multiple conditions will be joined into where statement using "AND".
// Eg:
// Where("uid=10000")
// Where("uid", 10000)
// Where("money>? AND name like ?", 99999, "vip_%")
// Where("uid", 1).Where("name", "john")
// Where("status IN (?)", g.Slice{1,2,3})
// Where("age IN(?,?)", 18, 50)
// Where(User{ Id : 1, UserName : "john"})
func (d *VoiceLoverVoiceRankDao) Where(where interface{}, args ...interface{}) *VoiceLoverVoiceRankDao {
	return &VoiceLoverVoiceRankDao{M: d.M.Where(where, args...)}
}

// WherePri does the same logic as M.Where except that if the parameter <where>
// is a single condition like int/string/float/slice, it treats the condition as the primary
// key value. That is, if primary key is "id" and given <where> parameter as "123", the
// WherePri function treats the condition as "id=123", but M.Where treats the condition
// as string "123".
func (d *VoiceLoverVoiceRankDao) WherePri(where interface{}, args ...interface{}) *VoiceLoverVoiceRankDao {
	return &VoiceLoverVoiceRankDao{M: d.M.WherePri(where, args...)}
}

// And adds "AND" condition to the where statement.
func (d *VoiceLoverVoiceRankDao) And(where interface{}, args ...interface{}) *VoiceLoverVoiceRankDao {
	return &VoiceLoverVoiceRankDao{M: d.M.And(where, args...)}
}

// Or adds "OR" condition to the where statement.
func (d *VoiceLoverVoiceRankDao) Or(where interface{}, args ...interface{}) *VoiceLoverVoiceRankDao {
	return &VoiceLoverVoiceRankDao{M: d.M.Or(where, args...)}
}

// Group sets the "GROUP BY" statement for the model.
func (d *VoiceLoverVoiceRankDao) Group(groupBy string) *VoiceLoverVoiceRankDao {
	return &VoiceLoverVoiceRankDao{M: d.M.Group(groupBy)}
}

// Order sets the "ORDER BY" statement for the model.
func (d *VoiceLoverVoiceRankDao) Order(orderBy ...string) *VoiceLoverVoiceRankDao {
	return &VoiceLoverVoiceRankDao{M: d.M.Order(orderBy...)}
}

// Limit sets the "LIMIT" statement for the model.
// The parameter <limit> can be either one or two number, if passed two number is passed,
// it then sets "LIMIT limit[0],limit[1]" statement for the model, or else it sets "LIMIT limit[0]"
// statement.
func (d *VoiceLoverVoiceRankDao) Limit(limit ...int) *VoiceLoverVoiceRankDao {
	return &VoiceLoverVoiceRankDao{M: d.M.Limit(limit...)}
}

// Offset sets the "OFFSET" statement for the model.
// It only makes sense for some databases like SQLServer, PostgreSQL, etc.
func (d *VoiceLoverVoiceRankDao) Offset(offset int) *VoiceLoverVoiceRankDao {
	return &VoiceLoverVoiceRankDao{M: d.M.Offset(offset)}
}

// Page sets the paging number for the model.
// The parameter <page> is started from 1 for paging.
// Note that, it differs that the Limit function start from 0 for "LIMIT" statement.
func (d *VoiceLoverVoiceRankDao) Page(page, limit int) *VoiceLoverVoiceRankDao {
	return &VoiceLoverVoiceRankDao{M: d.M.Page(page, limit)}
}

// Batch sets the batch operation number for the model.
func (d *VoiceLoverVoiceRankDao) Batch(batch int) *VoiceLoverVoiceRankDao {
	return &VoiceLoverVoiceRankDao{M: d.M.Batch(batch)}
}

// Cache sets the cache feature for the model. It caches the result of the sql, which means
// if there's another same sql request, it just reads and returns the result from cache, it
// but not committed and executed into the database.
//
// If the parameter <duration> < 0, which means it clear the cache with given <name>.
// If the parameter <duration> = 0, which means it never expires.
// If the parameter <duration> > 0, which means it expires after <duration>.
//
// The optional parameter <name> is used to bind a name to the cache, which means you can later
// control the cache like changing the <duration> or clearing the cache with specified <name>.
//
// Note that, the cache feature is disabled if the model is operating on a transaction.
func (d *VoiceLoverVoiceRankDao) Cache(duration time.Duration, name ...string) *VoiceLoverVoiceRankDao {
	return &VoiceLoverVoiceRankDao{M: d.M.Cache(duration, name...)}
}

// Data sets the operation data for the model.
// The parameter <data> can be type of string/map/gmap/slice/struct/*struct, etc.
// Eg:
// Data("uid=10000")
// Data("uid", 10000)
// Data(g.Map{"uid": 10000, "name":"john"})
// Data(g.Slice{g.Map{"uid": 10000, "name":"john"}, g.Map{"uid": 20000, "name":"smith"})
func (d *VoiceLoverVoiceRankDao) Data(data ...interface{}) *VoiceLoverVoiceRankDao {
	return &VoiceLoverVoiceRankDao{M: d.M.Data(data...)}
}

// All does "SELECT FROM ..." statement for the model.
// It retrieves the records from table and returns the result as []*pb.EntityVoiceLoverVoiceRank.
// It returns nil if there's no record retrieved with the given conditions from table.
//
// The optional parameter <where> is the same as the parameter of M.Where function,
// see M.Where.
func (d *VoiceLoverVoiceRankDao) All(where ...interface{}) ([]*pb.EntityVoiceLoverVoiceRank, error) {
	all, err := d.M.All(where...)
	if err != nil {
		return nil, err
	}
	var entities []*pb.EntityVoiceLoverVoiceRank
	//if err = all.Structs(&entities); err != nil && err != sql.ErrNoRows {
	//	return nil, err
	//}
	if err = d.scan(all, &entities); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entities, nil
}

// One retrieves one record from table and returns the result as *pb.EntityVoiceLoverVoiceRank.
// It returns nil if there's no record retrieved with the given conditions from table.
//
// The optional parameter <where> is the same as the parameter of M.Where function,
// see M.Where.
func (d *VoiceLoverVoiceRankDao) One(where ...interface{}) (*pb.EntityVoiceLoverVoiceRank, error) {
	one, err := d.M.One(where...)
	if err != nil {
		return nil, err
	}
	var entity *pb.EntityVoiceLoverVoiceRank
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entity, nil
}

// FindOne retrieves and returns a single Record by M.WherePri and M.One.
// Also see M.WherePri and M.One.
func (d *VoiceLoverVoiceRankDao) FindOne(where ...interface{}) (*pb.EntityVoiceLoverVoiceRank, error) {
	one, err := d.M.FindOne(where...)
	if err != nil {
		return nil, err
	}
	var entity *pb.EntityVoiceLoverVoiceRank
	if err = one.Struct(&entity); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entity, nil
}

// FindAll retrieves and returns Result by by M.WherePri and M.All.
// Also see M.WherePri and M.All.
func (d *VoiceLoverVoiceRankDao) FindAll(where ...interface{}) ([]*pb.EntityVoiceLoverVoiceRank, error) {
	all, err := d.M.FindAll(where...)
	if err != nil {
		return nil, err
	}
	var entities []*pb.EntityVoiceLoverVoiceRank
	//if err = all.Structs(&entities); err != nil && err != sql.ErrNoRows {
	//	return nil, err
	//}
	if err = d.scan(all, &entities); err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	return entities, nil
}

// Struct retrieves one record from table and converts it into given struct.
// The parameter <pointer> should be type of *struct/**struct. If type **struct is given,
// it can create the struct internally during converting.
//
// The optional parameter <where> is the same as the parameter of Model.Where function,
// see Model.Where.
//
// Note that it returns sql.ErrNoRows if there's no record retrieved with the given conditions
// from table and <pointer> is not nil.
//
// Eg:
// user := new(User)
// err  := dao.User.Where("id", 1).Struct(user)
//
// user := (*User)(nil)
// err  := dao.User.Where("id", 1).Struct(&user)
func (d *VoiceLoverVoiceRankDao) Struct(pointer interface{}, where ...interface{}) error {
	return d.M.Struct(pointer, where...)
}

// Structs retrieves records from table and converts them into given struct slice.
// The parameter <pointer> should be type of *[]struct/*[]*struct. It can create and fill the struct
// slice internally during converting.
//
// The optional parameter <where> is the same as the parameter of Model.Where function,
// see Model.Where.
//
// Note that it returns sql.ErrNoRows if there's no record retrieved with the given conditions
// from table and <pointer> is not empty.
//
// Eg:
// users := ([]User)(nil)
// err   := dao.User.Structs(&users)
//
// users := ([]*User)(nil)
// err   := dao.User.Structs(&users)
func (d *VoiceLoverVoiceRankDao) Structs(pointer interface{}, where ...interface{}) error {
	return d.M.Structs(pointer, where...)
}

// Scan automatically calls Struct or Structs function according to the type of parameter <pointer>.
// It calls function Struct if <pointer> is type of *struct/**struct.
// It calls function Structs if <pointer> is type of *[]struct/*[]*struct.
//
// The optional parameter <where> is the same as the parameter of Model.Where function,
// see Model.Where.
//
// Note that it returns sql.ErrNoRows if there's no record retrieved and given pointer is not empty or nil.
//
// Eg:
// user  := new(User)
// err   := dao.User.Where("id", 1).Scan(user)
//
// user  := (*User)(nil)
// err   := dao.User.Where("id", 1).Scan(&user)
//
// users := ([]User)(nil)
// err   := dao.User.Scan(&users)
//
// users := ([]*User)(nil)
// err   := dao.User.Scan(&users)
func (d *VoiceLoverVoiceRankDao) Scan(pointer interface{}, where ...interface{}) error {
	return d.M.Scan(pointer, where...)
}

// Chunk iterates the table with given size and callback function.
func (d *VoiceLoverVoiceRankDao) Chunk(limit int, callback func(entities []*pb.EntityVoiceLoverVoiceRank, err error) bool) {
	d.M.Chunk(limit, func(result gdb.Result, err error) bool {
		var entities []*pb.EntityVoiceLoverVoiceRank
		err = result.Structs(&entities)
		if err == sql.ErrNoRows {
			return false
		}
		return callback(entities, err)
	})
}

// LockUpdate sets the lock for update for current operation.
func (d *VoiceLoverVoiceRankDao) LockUpdate() *VoiceLoverVoiceRankDao {
	return &VoiceLoverVoiceRankDao{M: d.M.LockUpdate()}
}

// LockShared sets the lock in share mode for current operation.
func (d *VoiceLoverVoiceRankDao) LockShared() *VoiceLoverVoiceRankDao {
	return &VoiceLoverVoiceRankDao{M: d.M.LockShared()}
}

// Unscoped enables/disables the soft deleting feature.
func (d *VoiceLoverVoiceRankDao) Unscoped() *VoiceLoverVoiceRankDao {
	return &VoiceLoverVoiceRankDao{M: d.M.Unscoped()}
}

func (d *VoiceLoverVoiceRankDao) scan(res gdb.Result, in interface{}) error {
	length := res.Len()
	if length == 0 {
		return sql.ErrNoRows
	}
	v := reflect.ValueOf(in).Elem()
	newv := reflect.MakeSlice(v.Type(), 0, length)
	v.Set(newv)
	v.SetLen(length)

	index := 0
	var fieldToTag []string
	for i := 0; i < length; i++ {
		k := v.Type().Elem().Elem()
		newObj := reflect.New(k)
		if i == 0 {
			fieldToTag = d.getFieldTag(newObj)
		}
		err := d.mapping(res[i], newObj, fieldToTag)
		if err != nil {
			return err
		}
		v.Index(index).Set(newObj)
		index++
	}
	v.SetLen(index)
	return nil
}

func (d *VoiceLoverVoiceRankDao) getFieldTag(v reflect.Value) []string {
	t := v.Type()
	typ := t.Elem()
	val := v.Elem()
	vTypeOfT := val.Type()
	length := val.NumField()
	fieldToTag := make([]string, length)
	for j := 0; j < length; j++ {
		name := vTypeOfT.Field(j).Name
		if name[0] < 'A' || name[0] > 'Z' {
			continue
		}
		value := val.Field(j)
		tagVal := typ.Field(j).Tag.Get("orm")
		if !value.CanSet() || len(tagVal) == 0 {
			continue
		}
		index := strings.Index(tagVal, ",")
		if index > -1 {
			fieldToTag[j] = tagVal[0:index]
		} else {
			fieldToTag[j] = tagVal
		}
	}
	return fieldToTag
}

func (d *VoiceLoverVoiceRankDao) mapping(m map[string]*gvar.Var, v reflect.Value, fieldToTag []string) error {
	val := v.Elem()
	length := val.NumField()
	for i := 0; i < length; i++ {
		tag := fieldToTag[i]
		if len(tag) == 0 {
			continue
		}

		meta, ok := m[tag]
		if !ok {
			continue
		}
		value := val.Field(i)
		kind := value.Kind()
		switch kind {
		case reflect.String:
			value.SetString(meta.String())
			break
		case reflect.Float32, reflect.Float64:
			value.SetFloat(gconv.Float64(meta.String()))
			break
		case reflect.Int64, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
			value.SetInt(meta.Int64())
			break
		case reflect.Uint64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32:
			value.SetUint(meta.Uint64())
			break
		case reflect.Bool:
			value.SetBool(meta.Bool())
			break
		default:
			return errors.New("unknown kind " + kind.String())
		}
	}
	return nil
}
