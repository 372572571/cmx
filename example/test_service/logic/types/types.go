package types

import (
	"strconv"
)

type defaultGroupSwitch struct{}

var _defaultGroupSwitch = defaultGroupSwitch{}

type EDefaultGroupSwitch int32

// 开启
const DefaultGroupSwitch_On EDefaultGroupSwitch = 1

// 关闭
const DefaultGroupSwitch_Off EDefaultGroupSwitch = 0

func (e EDefaultGroupSwitch) desc() map[int32]string {
	return map[int32]string{
		1: "开启",
		0: "关闭",
	}
}
func (e EDefaultGroupSwitch) keys() map[int32]string {
	return map[int32]string{
		1: "on",
		0: "off",
	}
}
func (e EDefaultGroupSwitch) ZH() string  { return e.desc()[e.I3()] }
func (e EDefaultGroupSwitch) I3() int32   { return int32(e) }
func (e EDefaultGroupSwitch) UI3() uint32 { return uint32(e) }
func (e EDefaultGroupSwitch) I() int      { return int(e) }
func (e EDefaultGroupSwitch) S() string   { return strconv.Itoa(e.I()) }
func (e EDefaultGroupSwitch) K() string   { return e.keys()[e.I3()] }

var DefaultGroupSwitch func() defaultGroupSwitch = func() defaultGroupSwitch {
	return _defaultGroupSwitch
}

func (e defaultGroupSwitch) nameToValueMap() map[string]int32 {
	return map[string]int32{
		"on":  1,
		"off": 0,
	}
}
func (e defaultGroupSwitch) valueToNameMap() map[int32]string {
	return map[int32]string{
		1: "on",
		0: "off",
	}
}

func (e defaultGroupSwitch) Value(key string) int32 {
	return e.nameToValueMap()[key]
}

func (e defaultGroupSwitch) Key(value int32) string {
	return e.valueToNameMap()[value]
}

// 开启
func (e defaultGroupSwitch) OnKey() string {
	return "on"
}

// 关闭// 关闭
func (e defaultGroupSwitch) OffKey() string {
	return "off"
}

// 开启
func (e defaultGroupSwitch) OnValue() int32 {
	return 1
}

func (e defaultGroupSwitch) OnValString() string {
	return "1"
}

// 关闭// 关闭
func (e defaultGroupSwitch) OffValue() int32 {
	return 0
}

func (e defaultGroupSwitch) OffValString() string {
	return "0"
}

type defaultGroupCurd struct{}

var _defaultGroupCurd = defaultGroupCurd{}

type EDefaultGroupCurd int32

// 新增
const DefaultGroupCurd_Create EDefaultGroupCurd = 1

// 更新
const DefaultGroupCurd_Update EDefaultGroupCurd = 2

// 删除
const DefaultGroupCurd_Delete EDefaultGroupCurd = 3

// 查询
const DefaultGroupCurd_Read EDefaultGroupCurd = 4

func (e EDefaultGroupCurd) desc() map[int32]string {
	return map[int32]string{
		1: "新增",
		2: "更新",
		3: "删除",
		4: "查询",
	}
}
func (e EDefaultGroupCurd) keys() map[int32]string {
	return map[int32]string{
		1: "create",
		2: "update",
		3: "delete",
		4: "read",
	}
}
func (e EDefaultGroupCurd) ZH() string  { return e.desc()[e.I3()] }
func (e EDefaultGroupCurd) I3() int32   { return int32(e) }
func (e EDefaultGroupCurd) UI3() uint32 { return uint32(e) }
func (e EDefaultGroupCurd) I() int      { return int(e) }
func (e EDefaultGroupCurd) S() string   { return strconv.Itoa(e.I()) }
func (e EDefaultGroupCurd) K() string   { return e.keys()[e.I3()] }

var DefaultGroupCurd func() defaultGroupCurd = func() defaultGroupCurd {
	return _defaultGroupCurd
}

func (e defaultGroupCurd) nameToValueMap() map[string]int32 {
	return map[string]int32{
		"create": 1,
		"update": 2,
		"delete": 3,
		"read":   4,
	}
}
func (e defaultGroupCurd) valueToNameMap() map[int32]string {
	return map[int32]string{
		1: "create",
		2: "update",
		3: "delete",
		4: "read",
	}
}

func (e defaultGroupCurd) Value(key string) int32 {
	return e.nameToValueMap()[key]
}

func (e defaultGroupCurd) Key(value int32) string {
	return e.valueToNameMap()[value]
}

// 新增// 新增
func (e defaultGroupCurd) CreateKey() string {
	return "create"
}

// 更新// 更新
func (e defaultGroupCurd) UpdateKey() string {
	return "update"
}

// 删除// 删除
func (e defaultGroupCurd) DeleteKey() string {
	return "delete"
}

// 查询// 查询
func (e defaultGroupCurd) ReadKey() string {
	return "read"
}

// 新增// 新增
func (e defaultGroupCurd) CreateValue() int32 {
	return 1
}

func (e defaultGroupCurd) CreateValString() string {
	return "1"
}

// 更新// 更新
func (e defaultGroupCurd) UpdateValue() int32 {
	return 2
}

func (e defaultGroupCurd) UpdateValString() string {
	return "2"
}

// 删除// 删除
func (e defaultGroupCurd) DeleteValue() int32 {
	return 3
}

func (e defaultGroupCurd) DeleteValString() string {
	return "3"
}

// 查询// 查询
func (e defaultGroupCurd) ReadValue() int32 {
	return 4
}

func (e defaultGroupCurd) ReadValString() string {
	return "4"
}

type defaultGroupAuth struct{}

var _defaultGroupAuth = defaultGroupAuth{}

type EDefaultGroupAuth int32

// 授权码
const DefaultGroupAuth_AuthorizationCode EDefaultGroupAuth = 1

// 刷新token
const DefaultGroupAuth_RefreshToken EDefaultGroupAuth = 2

func (e EDefaultGroupAuth) desc() map[int32]string {
	return map[int32]string{
		1: "授权码",
		2: "刷新token",
	}
}
func (e EDefaultGroupAuth) keys() map[int32]string {
	return map[int32]string{
		1: "authorization_code",
		2: "refresh_token",
	}
}
func (e EDefaultGroupAuth) ZH() string  { return e.desc()[e.I3()] }
func (e EDefaultGroupAuth) I3() int32   { return int32(e) }
func (e EDefaultGroupAuth) UI3() uint32 { return uint32(e) }
func (e EDefaultGroupAuth) I() int      { return int(e) }
func (e EDefaultGroupAuth) S() string   { return strconv.Itoa(e.I()) }
func (e EDefaultGroupAuth) K() string   { return e.keys()[e.I3()] }

var DefaultGroupAuth func() defaultGroupAuth = func() defaultGroupAuth {
	return _defaultGroupAuth
}

func (e defaultGroupAuth) nameToValueMap() map[string]int32 {
	return map[string]int32{
		"authorization_code": 1,
		"refresh_token":      2,
	}
}
func (e defaultGroupAuth) valueToNameMap() map[int32]string {
	return map[int32]string{
		1: "authorization_code",
		2: "refresh_token",
	}
}

func (e defaultGroupAuth) Value(key string) int32 {
	return e.nameToValueMap()[key]
}

func (e defaultGroupAuth) Key(value int32) string {
	return e.valueToNameMap()[value]
}

// 授权码// 授权码
func (e defaultGroupAuth) AuthorizationCodeKey() string {
	return "authorization_code"
}

// 刷新token// 刷新token
func (e defaultGroupAuth) RefreshTokenKey() string {
	return "refresh_token"
}

// 授权码// 授权码
func (e defaultGroupAuth) AuthorizationCodeValue() int32 {
	return 1
}

func (e defaultGroupAuth) AuthorizationCodeValString() string {
	return "1"
}

// 刷新token// 刷新token
func (e defaultGroupAuth) RefreshTokenValue() int32 {
	return 2
}

func (e defaultGroupAuth) RefreshTokenValString() string {
	return "2"
}

type defaultGroupBoolean struct{}

var _defaultGroupBoolean = defaultGroupBoolean{}

type EDefaultGroupBoolean int32

// true
const DefaultGroupBoolean_T EDefaultGroupBoolean = 1

// false
const DefaultGroupBoolean_F EDefaultGroupBoolean = 0

func (e EDefaultGroupBoolean) desc() map[int32]string {
	return map[int32]string{
		1: "true",
		0: "false",
	}
}
func (e EDefaultGroupBoolean) keys() map[int32]string {
	return map[int32]string{
		1: "T",
		0: "F",
	}
}
func (e EDefaultGroupBoolean) ZH() string  { return e.desc()[e.I3()] }
func (e EDefaultGroupBoolean) I3() int32   { return int32(e) }
func (e EDefaultGroupBoolean) UI3() uint32 { return uint32(e) }
func (e EDefaultGroupBoolean) I() int      { return int(e) }
func (e EDefaultGroupBoolean) S() string   { return strconv.Itoa(e.I()) }
func (e EDefaultGroupBoolean) K() string   { return e.keys()[e.I3()] }

var DefaultGroupBoolean func() defaultGroupBoolean = func() defaultGroupBoolean {
	return _defaultGroupBoolean
}

func (e defaultGroupBoolean) nameToValueMap() map[string]int32 {
	return map[string]int32{
		"T": 1,
		"F": 0,
	}
}
func (e defaultGroupBoolean) valueToNameMap() map[int32]string {
	return map[int32]string{
		1: "T",
		0: "F",
	}
}

func (e defaultGroupBoolean) Value(key string) int32 {
	return e.nameToValueMap()[key]
}

func (e defaultGroupBoolean) Key(value int32) string {
	return e.valueToNameMap()[value]
}

// true// true
func (e defaultGroupBoolean) TKey() string {
	return "T"
}

// false// false
func (e defaultGroupBoolean) FKey() string {
	return "F"
}

// true// true
func (e defaultGroupBoolean) TValue() int32 {
	return 1
}

func (e defaultGroupBoolean) TValString() string {
	return "1"
}

// false// false
func (e defaultGroupBoolean) FValue() int32 {
	return 0
}

func (e defaultGroupBoolean) FValString() string {
	return "0"
}
