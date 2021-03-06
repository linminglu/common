//author saplmm

package db

import (
	"casino_common/common/consts"
	"casino_common/common/log"
	"casino_common/utils/redisUtils"
	"errors"
	"github.com/name5566/leaf/db/mongodb"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"reflect"
	"sync"
	"time"
)

func init() {
	//初始化dialc
	mongoConfig.dialc = map[string]*mongodb.DialContext{}
}

type BaseMode interface {
	GetId() int32
}

type BaseModeu32 interface {
	GetId() uint32
}

var mongoConfig struct {
	ip                   string
	logIp                string
	port                 int
	dbname               string
	DB_ENSURECOUNTER_KEY string
	dialc                map[string]*mongodb.DialContext
	SessionNum           int
}

func Oninit(ip string, logIp string, dbname string, key string) {
	log.D("初始化mongoDb的信息  ip[%v] logIp[%v] dbname[%v] key[%v]", ip, logIp, dbname, key)
	mongoConfig.ip = ip
	mongoConfig.logIp = logIp
	mongoConfig.dbname = dbname
	mongoConfig.DB_ENSURECOUNTER_KEY = key
	mongoConfig.SessionNum = 100
	dial_context, _ := mongodb.Dial(mongoConfig.ip, mongoConfig.SessionNum)
	mongoConfig.dialc[mongoConfig.ip] = dial_context
}

func GetDBName() string {
	return mongoConfig.dbname
}

//活的链接
func GetMongoConn() (*mongodb.DialContext, error) {
	if dial_context, ok := mongoConfig.dialc[mongoConfig.ip]; ok {
		return dial_context, nil
	}
	dial_context, err := mongodb.Dial(mongoConfig.ip, mongoConfig.SessionNum)
	if err != nil {
		return nil, err
	}
	mongoConfig.dialc[mongoConfig.ip] = dial_context
	return dial_context, nil
}

//获得一个连接，通过连接地址
func GetMongoConnByDialAddr(dialAddr string) (*mongodb.DialContext, error) {
	if dialAddr == "" {
		dialAddr = mongoConfig.ip
	}
	if dial_context, ok := mongoConfig.dialc[dialAddr]; ok {
		return dial_context, nil
	}
	dial_context, err := mongodb.Dial(dialAddr, mongoConfig.SessionNum)
	if err != nil {
		return nil, err
	}
	mongoConfig.dialc[dialAddr] = dial_context
	return dial_context, nil
}

//保存数据（dialAddr为空则连接默认数据库）
func InsertMgoData(dbt string, data interface{}) error {
	//得到连接
	c, err := GetMongoConn()
	if err != nil {
		return err
	}
	//defer c.Close()

	// 获取回话 session
	s := c.Ref()
	defer c.UnRef(s)

	error := s.DB(mongoConfig.dbname).C(dbt).Insert(data)
	return error
}

//批量保存数据（dialAddr为空则连接默认数据库）
func InsertMgoDatas(dbt string, datas []interface{}) (err error, count int) {
	c, err := GetMongoConn()
	if err != nil {
		return err, -1
	}
	//defer c.Close()

	// 获取回话 session
	s := c.Ref()
	defer c.UnRef(s)

	for _, data := range datas {

		err = s.DB(mongoConfig.dbname).C(dbt).Insert(data)
		if err == nil {
			count++
		}
	}
	return nil, count
}

//更新数据通过_id来更新
func UpdateMgoData(dbt string, data BaseMode) error {
	c, err := GetMongoConn()
	if err != nil {
		return err
	}
	//defer c.Close()

	// 获取回话 session
	s := c.Ref()
	defer c.UnRef(s)

	error := s.DB(mongoConfig.dbname).C(dbt).Update(bson.M{"id": data.GetId()}, data)
	return error
}

func UpdateMgoDataU32(dbt string, data BaseModeu32) error {
	c, err := GetMongoConn()
	if err != nil {
		return err
	}
	//defer c.Close()

	// 获取回话 session
	s := c.Ref()
	defer c.UnRef(s)

	error := s.DB(mongoConfig.dbname).C(dbt).Update(bson.M{"id": data.GetId()}, data)
	return error
}

var next_redis_seq_lock sync.Mutex

//得到序列号（新版从redis中取）
func GetNextSeq(dbt string) (int32, error) {
	//加锁防止id重复及其他异常
	next_redis_seq_lock.Lock()
	defer next_redis_seq_lock.Unlock()

	rkey := redisUtils.K_STRING(consts.RKEY_SEQ_ID_KEY, dbt)

	rval := redisUtils.GetInt64(rkey)

	//如果不存在该键值，则从数据库中查询
	if rval == 0 {
		db_id, err := GetNextSeqByMongo(dbt)
		//如果数据库中不存在，则直接返回错误
		if err != nil {
			return db_id, err
		}
		//如果数据库中存在则将数据set进redis中并返回
		redisUtils.SetInt64(rkey, int64(db_id))
		return db_id, err
	}

	//返回redis中存的值
	new_val := rval + 1
	redisUtils.SetInt64(rkey, new_val)

	return int32(new_val), nil
}

//获取自增ID
func GetNextIncrementID(dbt string, rdb string) (int32, error) {
	//加锁防止id重复及其他异常
	next_redis_seq_lock.Lock()
	defer next_redis_seq_lock.Unlock()

	rkey := redisUtils.K_STRING(rdb, dbt)

	rval := redisUtils.GetInt64(rkey)

	//如果不存在该键值，则从数据库中查询
	if rval == 0 {
		db_id, err := GetNextSeqByMongo(dbt)
		//如果数据库中不存在，则直接返回错误
		if err != nil {
			return db_id, err
		}
		//如果数据库中存在则将数据set进redis中并返回
		redisUtils.SetInt64(rkey, int64(db_id))
		return db_id, err
	}

	//返回redis中存的值
	new_val := rval + 1
	redisUtils.SetInt64(rkey, new_val)

	return int32(new_val), nil
}

//通过数据库-得到自增键（旧版从mmongo中取）
func GetNextSeqByMongo(dbt string) (int32, error) {
	//连接数据库
	c, err := GetMongoConn()
	if err != nil {
		return 0, err
	}
	//defer c.Close()

	//获取session
	s := c.Ref()
	defer c.UnRef(s)
	id, err := c.NextSeq(mongoConfig.dbname, dbt, mongoConfig.DB_ENSURECOUNTER_KEY)
	return int32(id), err
}

//查询一个list
//如果连接默认数据库，则dialAddr填空字符串即可。
func QueryByDialAddr(dialAddr string, f func(*mgo.Database)) {
	c, err := GetMongoConnByDialAddr(dialAddr)

	if err != nil {
		log.E("QueryByDialAddr(%s) err:%v", dialAddr, err)
		return
	}
	//defer c.Close()

	//获取session
	s := c.Ref()
	defer c.UnRef(s)
	f(s.DB(mongoConfig.dbname))
}

//如果连接默认数据库，则dialAddr填空字符串即可。
func Query(f func(*mgo.Database)) {
	c, err := GetMongoConn()

	if err != nil {
		log.E("Query() err:%v", err)
		return
	}
	//defer c.Close()

	//获取session
	s := c.Ref()
	defer c.UnRef(s)
	f(s.DB(mongoConfig.dbname))
}

//通过k==v查询一条数据
func FindByKV(dialAddr string, dbt string, key string, value interface{}, obj interface{}) error {
	var err error = errors.New("connect err.")
	QueryByDialAddr(dialAddr, func(mgo *mgo.Database) {
		err = mgo.C(dbt).Find(bson.M{
			key: value,
		}).One(obj)
	})
	return err
}

//通过k==v查询多条数据
func SelectByKV(dialAddr string, dbt string, key string, value interface{}, obj interface{}) error {
	var err error = errors.New("connect err.")
	QueryByDialAddr(dialAddr, func(mgo *mgo.Database) {
		err = mgo.C(dbt).Find(bson.M{
			key: value,
		}).All(obj)
	})
	return err
}

func CloseMGO() {
	log.T("mgo close...")
	for _, dial_context := range mongoConfig.dialc {
		dial_context.Close()
	}
}

//模拟类似ThinkPHP的M()方法
type Collection struct {
	TableName string
	DialAddr  string
}

//主数据库
//args[0]:TableName args[1]:DialAddr
func C(tableName string) *Collection {
	return &Collection{
		TableName: tableName,
		DialAddr:  mongoConfig.ip,
	}
}

//日志数据库
func Log(tableName string) *Collection {
	return &Collection{
		TableName: tableName,
		DialAddr:  mongoConfig.logIp,
	}
}

//查询单条
func (c *Collection) Find(query interface{}, result interface{}) error {
	if c == nil {
		return errors.New("collection is nil.")
	}
	var err error = errors.New("connect err.")
	QueryByDialAddr(c.DialAddr, func(mgo *mgo.Database) {
		err = mgo.C(c.TableName).Find(query).One(result)
	})
	return err
}

//查询多条
func (c *Collection) FindAll(query interface{}, result interface{}) error {
	if c == nil {
		return errors.New("collection is nil.")
	}
	var err error = errors.New("connect err.")
	QueryByDialAddr(c.DialAddr, func(mgo *mgo.Database) {
		err = mgo.C(c.TableName).Find(query).All(result)
	})
	return err
}

//查询分页
func (c *Collection) Page(query interface{}, result interface{}, sort string, page int, limit int) (err error, count int) {
	if c == nil {
		return errors.New("collection is nil."), 0
	}
	err = errors.New("connect err.")
	QueryByDialAddr(c.DialAddr, func(mgo *mgo.Database) {
		q := mgo.C(c.TableName).Find(query)
		//统计数量
		count, err = q.Count()

		if sort != "" {
			q = q.Sort(sort)
		}
		if limit > 0 && page >= 1 {
			q = q.Skip((page - 1) * limit).Limit(limit)
		}
		err = q.All(result)
	})
	return err, count
}

//更新一条
func (c *Collection) Update(query interface{}, update interface{}) error {
	if c == nil {
		return errors.New("collection is nil.")
	}
	var err error = errors.New("connect err.")
	QueryByDialAddr(c.DialAddr, func(mgo *mgo.Database) {
		err = mgo.C(c.TableName).Update(query, update)
	})
	return err
}

//更新多条
func (c *Collection) UpdateAll(query interface{}, update interface{}) (change_info *mgo.ChangeInfo, err error) {
	if c == nil {
		return nil, errors.New("collection is nil.")
	}
	err = errors.New("connect err.")
	QueryByDialAddr(c.DialAddr, func(mgo *mgo.Database) {
		change_info, err = mgo.C(c.TableName).UpdateAll(query, update)
	})
	return
}

//查不到则插入
func (c *Collection) Upsert(query interface{}, update interface{}) error {
	if c == nil {
		return errors.New("collection is nil.")
	}
	var err error = errors.New("connect err.")
	QueryByDialAddr(c.DialAddr, func(mgo *mgo.Database) {
		_, err = mgo.C(c.TableName).Upsert(query, update)
	})
	return err
}

//插入一条或多条
func (c *Collection) Insert(doc ...interface{}) error {
	if c == nil {
		return errors.New("collection is nil.")
	}
	var err error = errors.New("connect err.")
	QueryByDialAddr(c.DialAddr, func(mgo *mgo.Database) {
		err = mgo.C(c.TableName).Insert(doc...)
	})
	return err
}

//一次性插入多条数据
func (c *Collection) InsertAll(doc_slice interface{}) (error, int) {
	if c == nil {
		return errors.New("collection is nil."), 0
	}
	if reflect.TypeOf(doc_slice).Kind() != reflect.Slice {
		return errors.New("doc_slice must be a slice."), 0
	}
	val := reflect.ValueOf(doc_slice)
	doc_arr := []interface{}{}
	doc_len := val.Len()
	for i := 0; i < doc_len; i++ {
		doc_arr = append(doc_arr, val.Index(i).Interface())
	}
	return c.Insert(doc_arr...), doc_len
}

//删除单条数据
func (c *Collection) Remove(query interface{}) error {
	if c == nil {
		return errors.New("collection is nil.")
	}
	var err error = errors.New("connect err.")
	QueryByDialAddr(c.DialAddr, func(mgo *mgo.Database) {
		err = mgo.C(c.TableName).Remove(query)
	})
	return err
}

//删除多条数据
func (c *Collection) RemoveAll(query interface{}) (change_info *mgo.ChangeInfo, err error) {
	if c == nil {
		return nil, errors.New("collection is nil.")
	}
	err = errors.New("connect err.")
	QueryByDialAddr(c.DialAddr, func(mgo *mgo.Database) {
		change_info, err = mgo.C(c.TableName).RemoveAll(query)
	})
	return
}

//计数
func (c *Collection) Count(query interface{}) (count int, err error) {
	if c == nil {
		return 0, errors.New("collection is nil.")
	}
	err = errors.New("connect err.")
	QueryByDialAddr(c.DialAddr, func(mgo *mgo.Database) {
		count, err = mgo.C(c.TableName).Find(query).Count()
	})
	return
}

//管道
func (c *Collection) PipeAll(query interface{}, result interface{}) (err error) {
	if c == nil {
		return errors.New("collection is nil.")
	}
	err = errors.New("connect err.")
	QueryByDialAddr(c.DialAddr, func(mgo *mgo.Database) {
		pipe := mgo.C(c.TableName).Pipe(query)
		err = pipe.All(result)
	})
	return
}

//管道
func (c *Collection) Pipe(query interface{}, result interface{}) (err error) {
	if c == nil {
		return errors.New("collection is nil.")
	}
	err = errors.New("connect err.")
	QueryByDialAddr(c.DialAddr, func(mgo *mgo.Database) {
		pipe := mgo.C(c.TableName).Pipe(query)
		err = pipe.One(result)
	})
	return
}

//时间区间统计器
type DateRangeCounter []struct {
	Sum  float64
	Time time.Time
}

func (data *DateRangeCounter) GetTimeJsonArr() string {
	return ""
}

/**
每日数据统计
list, err := db.C("t_test_group").RangeDayCount(bson.M{}, "$time", "$sum", "$score")
*/
func (c *Collection) RangeDayCount(query interface{}, time_field string, count_type string, count_field string) (result DateRangeCounter, err error) {
	err = c.PipeAll([]bson.M{
		bson.M{
			"$match": query,
		},
		bson.M{
			"$group": bson.M{
				"_id": bson.M{
					"year":  bson.M{"$year": time_field},
					"month": bson.M{"$month": time_field},
					"day":   bson.M{"$dayOfMonth": time_field},
				},
				"sum":  bson.M{count_type: count_field},
				"time": bson.M{"$first": time_field},
			},
		},
		bson.M{
			"$sort": bson.M{
				"time": 1,
			},
		},
	}, &result)
	return result, err
}

//删除表
func (c *Collection) Drop() (err error) {
	if c == nil {
		return errors.New("collection is nil.")
	}
	err = errors.New("connect err.")
	QueryByDialAddr(c.DialAddr, func(mgo *mgo.Database) {
		err = mgo.C(c.TableName).DropCollection()
	})
	return
}

type ObjIds struct {
	Id bson.ObjectId `bson:"_id"`
}

//查询id列表
func (c *Collection) FindAllId(query interface{}) (list []bson.ObjectId, err error) {
	if c == nil {
		return list, errors.New("collection is nil.")
	}
	err = errors.New("connect err.")
	res := []ObjIds{}
	QueryByDialAddr(c.DialAddr, func(mgo *mgo.Database) {
		err = mgo.C(c.TableName).Find(query).All(&res)
	})
	if err != nil {
		return list, err
	}
	for _, id := range res {
		list = append(list, id.Id)
	}
	return list, nil
}
