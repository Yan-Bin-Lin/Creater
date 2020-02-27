package database

import (
	"time"
)

type User struct {
	Uid        int       `json:"uid" xorm:"not null pk autoincr INT(11) 'uid'"`
	Username   string    `json:"username" xorm:"not null comment('account number') VARCHAR(40) 'username'"`
	Password   string    `json:"password" xorm:"not null CHAR(40) 'password'"`
	Email      string    `json:"email" xorm:"not null VARCHAR(40) 'email'"`
	Createtime time.Time `json:"ucreatetime" xorm:"default 'current_timestamp()' comment('REGIST TIME') DATETIME 'createtime'"`
}

type Owner struct {
	Oid        int       `json:"oid" xorm:"not null pk autoincr INT(11) 'oid'"`
	Uid        int       `json:"ouid" xorm:"not null INT(11) 'uid'"`
	Nickname   string    `json:"nickname" xorm:"not null VARCHAR(50) 'nickname'"`
	Uniquename string    `json:"uniquename" xorm:"not null comment('id of User-defined') VARCHAR(50) 'uniquename'"`
	Createtime time.Time `json:"ocreatetime" xorm:"default 'current_timestamp()' DATETIME 'createtime'"`
	Updatetime time.Time `json:"oupdatetime" xorm:"default 'current_timestamp()' comment('last login time') DATETIME 'updatetime'"`
}

type Project struct {
	Pid         int       `json:"pid" xorm:"not null pk autoincr INT(11) 'pid'"`
	Oid         int       `json:"poid" xorm:"not null INT(11) 'oid'"`
	Name        string    `json:"pname" xorm:"not null VARCHAR(50) 'name'"`
	Like        int       `json:"plike" xorm:"not null default 0 INT(11) 'like'"`
	Hate        int       `json:"phate" xorm:"not null default 0 INT(11) 'hate'"`
	Viewtime    int       `json:"pviewtime" xorm:"not null default 0 INT(11) 'viewtime'"`
	Super       int       `json:"psuper" xorm:"INT(11) 'super'"`
	Description string    `json:"pdescription" xorm:"comment('abstract') VARCHAR(255) 'description'"`
	Urlpath     string    `json:"urlpath" xorm:"not null CHAR(44) 'urlpath'"`
	Createtime  time.Time `json:"pcreatetime" xorm:"default 'current_timestamp()' DATETIME 'createtime'"`
	Updatetime  time.Time `json:"pupdatetime" xorm:"default 'current_timestamp()' DATETIME 'updatetime'"`
}

type Category struct {
	Cid    int    `json:"cid" xorm:"not null pk autoincr INT(11) 'id'"`
	Super  int    `json:"csuper" xorm:"not null INT(11) 'super'"`
	Number int    `json:"cnumber" xorm:"not null default 1 INT(11) 'number'"`
	Name   string `json:"cname" xorm:"not null VARCHAR(50) 'name'"`
}

type Blogtype struct {
	Typeid int    `json:"typeid" xorm:"not null pk autoincr TINYINT(4) 'typeid'"`
	Name   string `json:"name" xorm:"not null VARCHAR(20) 'name'"`
}

type Blog struct {
	Bid         int       `json:"bid" xorm:"not null pk autoincr INT(11) 'id'"`
	Name        string    `json:"bname" xorm:"not null VARCHAR(50) 'name'"`
	Like        int       `json:"blike" xorm:"not null default 0 INT(11) 'like'"`
	Hate        int       `json:"bhate" xorm:"not null default 0 INT(11) 'hate'"`
	Viewtime    int       `json:"bviewtime" xorm:"not null default 0 INT(11) 'viewtime'"`
	Super       int       `json:"bsuper" xorm:"INT(11) 'super'"`
	Cid         int       `json:"bbelong" xorm:"INT(11) 'cid'"`
	Number      int       `json:"bnumber" xorm:"not null default 1 INT(11) 'number'"`
	Description string    `json:"bdescription" xorm:"comment('abstract') VARCHAR(255) 'description'"`
	Type        int       `json:"type" xorm:"not null TINYINT(4) 'type'"`
	Filepath    string    `json:"filepath" xorm:"not null CHAR(44) 'filepath'"`
	Createtime  time.Time `json:"bcreatetime" xorm:"default 'current_timestamp()' DATETIME 'createtime'"`
	Updatetime  time.Time `json:"bupdatetime" xorm:"default 'current_timestamp()' DATETIME 'updatetime'"`
}

type Article struct {
	Name  string `json:"name" xorm:"not null pk CHAR(44) 'name'"`
	Super int    `json:"super" xorm:"not null INT(11) 'super'"`
}

/* parse to html response */
type UserOut struct {
	Uid      int    `json:"uid" xorm:"not null pk autoincr INT(11) 'uid'"`
	Username string `json:"username" xorm:"not null comment('account number') VARCHAR(40) 'username'"`
	Owner    `xorm:"extends"`
}

type OwnerOut struct {
	Owner   `xorm:"extends"`
	Project `xorm:"extends"`
	Blog    `xorm:"extends"`
	Catname string `json:"cname" xorm:"not null VARCHAR(50) 'name'"`
}

type ProjOut struct {
	Owner          `xorm:"extends"`
	SuperPid       int    `json:"superpid"xorm:"not null pk autoincr INT(11) 'pid'"`
	SuperName      string `json:"supername" xorm:"not null VARCHAR(50) 'name'"`
	Project        `xorm:"extends"`
	SubPid         int    `json:"subpid" xorm:"not null pk autoincr INT(11) 'pid'"`
	SubName        string `json:"subpname" xorm:"not null VARCHAR(50) 'name'"`
	SubDescription string `json:"subpdescription" xorm:"comment('abstract') VARCHAR(255) 'description'"`
	Blog           `xorm:"extends"`
	Catname        string `json:"cname" xorm:"not null VARCHAR(50) 'name'"`
}

type BlogOut struct {
	Owner    `xorm:"extends"`
	ProjName string `json:"projname" xorm:"not null VARCHAR(50) 'name'"`
	Blog     `xorm:"extends"`
	Catname  string `json:"cname" xorm:"not null VARCHAR(50) 'name'"`
}
