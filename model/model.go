package model

import (
	"time"

	"gorm.io/gorm"
)

type Register struct {
	Tenent TenantMaster `json:"tenent,omitempty"`
	User   UserMaster   `json:"user,omitempty"`
}

type TenantMaster struct {
	gorm.Model
	Name       string       `gorm:"column:name" json:"name,omitempty"`
	Status     string       `gorm:"column:status" json:"status,omitempty"` // Active / block / Lock
	Domain     string       `gorm:"domain;index:idx_domain,unique" json:"domain,omitempty"`
	ExpiryDate time.Time    `gorm:"column:expiry_date" json:"expiry_date,omitempty"`
	Users      []UserMaster `json:"users,omitempty"`
}

// User AS in User / Device
type UserMaster struct {
	gorm.Model
	TenantMasterID uint         `gorm:"column:tenant_master_id;index:idx_member,unique,priority:1" json:"tenant_master_id,omitempty"`
	Name           string       `gorm:"column:name" json:"name,omitempty"`
	Username       string       `gorm:"column:username;index:idx_member,unique,priority:2" json:"username,omitempty"`
	ExpiryDate     time.Time    `gorm:"column:expiry_date" json:"expiry_date,omitempty"`
	UserType       string       `gorm:"column:user_type" json:"user_type,omitempty"`
	Password       string       `gorm:"column:password" json:"password,omitempty"`
	Address        string       `gorm:"column:address" json:"address,omitempty"`
	Contacts       []Contact    `json:"contacts,omitempty"`
	DeviceInfos    []DeviceInfo `json:"device_infos,omitempty"`
}

type Contact struct {
	gorm.Model
	UserMasterID uint   `gorm:"column:user_master_id" json:"user_master_id,omitempty"`
	Types        string `gorm:"column:types" json:"types,omitempty"`
	Value        string `gorm:"column:value" json:"value,omitempty"`
	Status       string `gorm:"column:status" json:"status,omitempty"` // Active / Block / Lock
}

type Login struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Key      string `json:"key,omitempty"`
}

type DeviceInfo struct {
	gorm.Model
	UserMasterID uint      `gorm:"column:user_master_id" json:"user_master_id,omitempty"`
	Hostname     string    `gorm:"column:hostname" json:"hostname,omitempty"`
	DeviceKey    string    `gorm:"column:device_key;index:idx_keys,unique,priority:1" json:"device_key,omitempty"`
	Types        string    `gorm:"types" json:"types,omitempty"`
	OS           string    `gorm:"os" json:"os,omitempty"`
	Status       string    `gorm:"status" json:"status,omitempty"`
	ExpiryDate   time.Time `gorm:"expiry_date" json:"expiry_date,omitempty"`
}
