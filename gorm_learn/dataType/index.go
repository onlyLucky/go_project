package dataType

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"gorm.io/gorm"
)

func DataTypeFunc(DB *gorm.DB) {
	/* // 1.建表
	crateTableFunc(DB)
	// 2.添加和查询
	addDataFunc(DB)
	// 3.枚举
	enum1Func()
	enum2Func()
	enum3Func()
	enum4Func()
	enum5Func() */
}

/* 
1.存储结构体
自定义的数据类型必须实现 Scanner 和 Valuer 接口，以便让 GORM 知道如何将该类型接收、保存到数据库
*/
type Info struct {
	Status string `json:"status"`
	Addr string `json:"addr"`
	Age int `json:"age"`
}

// Scan从数据种读取出来
func (i *Info) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("Failed to unmarshal JSONB value:", value))
	}
	info := Info{}
	err := json.Unmarshal(bytes,&info)
	*i = info
	fmt.Println("Info:=====",err,i)
	return err
}
// Value 存入数据库
func (i Info) Value() (driver.Value,error) {
	return json.Marshal(i)
}
type User struct{
	ID uint   `gorm:"primaryKey;"`
	Name string
	Info Info `gorm:"type: json"`
}

func crateTableFunc(DB *gorm.DB){
	DB.Migrator().DropTable(&User{})
	DB.AutoMigrate(&User{})
}

/* 2.添加和查询 */
func addDataFunc(DB *gorm.DB){
	DB.Create(&User{
		Name: "枫枫",
		Info: Info{
			Status: "牛逼",
			Addr:   "成都市",
			Age:    21,
		},
	})
	var user User
	DB.Take(&user)
	fmt.Println(user)
}

/* 2.枚举 */

func enum1Func(){
	type Host struct {
		ID     uint
		Name   string
		Status string
	}
	host := Host{}
	if host.Status == "Running"{
		fmt.Println("在线")
	}
	if host.Status == "Except" {
    fmt.Println("异常")
  }
  if host.Status == "OffLine" {
    fmt.Println("离线")
  }
}

func enum2Func(){
	type Host struct {
		ID     uint
		Name   string
		Status string
	}
	
	const (
		Running = "Running"
		Except = "Except"
		OffLine = "OffLine"
	)
	host := Host{}
  if host.Status == Running {
    fmt.Println("在线")
  }
  if host.Status == Except {
    fmt.Println("异常")
  }
  if host.Status == OffLine {
    fmt.Println("离线")
  }
}

func enum3Func(){
	type Host struct {
		ID     uint
		Name   string
		Status int
	}
	
	const (
		Running = 1
		Except  = 2
		OffLine = 3
	)

	host := Host{}
  if host.Status == Running {
    fmt.Println("在线")
  }
  if host.Status == Except {
    fmt.Println("异常")
  }
  if host.Status == OffLine {
    fmt.Println("离线")
  }
}

type Host struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Status int    `json:"status"`
}
const (
  Running = 1
  Except  = 2
  OffLine  = 3
)
func (h Host) MarshalJSON() ([]byte, error) {
	var status string
	switch h.Status {
	case Running:
		status = "Running"
	case Except:
		status = "Except"
	case OffLine :
		status = "OffLine"
	}
	return json.Marshal(&struct {
		ID     uint   `json:"id"`
		Name   string `json:"name"`
		Status string `json:"status"`
	}{
		ID:     h.ID,
		Name:   h.Name,
		Status: status,
	})
}
func enum4Func(){
	host := Host{1, "枫枫", Running}
  data, _ := json.Marshal(host)
  fmt.Println(string(data)) // {"id":1,"name":"枫枫","status":"Running"}
}

type Weekday int

const (
  Sunday    Weekday = iota + 1 // EnumIndex = 1
  Monday                       // EnumIndex = 2
  Tuesday                      // EnumIndex = 3
  Wednesday                    // EnumIndex = 4
  Thursday                     // EnumIndex = 5
  Friday                       // EnumIndex = 6
  Saturday                     // EnumIndex = 7
)

var WeekStringList = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
var WeekTypeList = []Weekday{Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday}

// String 转字符串
func (w Weekday) String() string {
  return WeekStringList[w-1]
}

// MarshalJSON 自定义类型转换为json
func (w Weekday) MarshalJSON() ([]byte, error) {
  return json.Marshal(w.String())
}

// EnumIndex 自定义类型转原始类型
func (w Weekday) EnumIndex() int {
  return int(w)
}

// ParseWeekDay 字符串转自定义类型
func ParseWeekDay(week string) Weekday {
  for i, i2 := range WeekStringList {
    if week == i2 {
      return WeekTypeList[i]
    }
  }
  return Monday
}

// ParseIntWeekDay 数字转自定义类型
func ParseIntWeekDay(week int) Weekday {
  return Weekday(week)
}

type DayInfo struct {
  Weekday Weekday   `json:"weekday"`
  Date    time.Time `json:"date"`
}

func enum5Func(){
	w := Sunday
  fmt.Println(w)
  dayInfo := DayInfo{Weekday: Sunday, Date: time.Now()}
  data, err := json.Marshal(dayInfo)
  fmt.Println(string(data), err)
  week := ParseWeekDay("Sunday")
  fmt.Println(week)
  week = ParseIntWeekDay(2)
  fmt.Println(week)
}