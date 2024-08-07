package types

import ()

type defaultGroupCurd struct{}

var _defaultGroupCurd = defaultGroupCurd{}
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

type defaultGroupSwitch struct{}

var _defaultGroupSwitch = defaultGroupSwitch{}
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
