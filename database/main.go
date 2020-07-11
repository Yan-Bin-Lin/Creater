package database

import (
	"database/sql"
	"errors"
)

var (
	ERR_TASK_FAIL     = errors.New("Fail to affect row")
	ERR_NAME_CONFLICT = errors.New("NAME CONFLICT")
	ERR_PARAMETER     = errors.New("PARAMETER WRONG")

	//projAs     = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}
	ownerCol   = []string{"owner.oid", "owner.nickname", "owner.uniquename", "owner.single", "owner.createtime", "owner.updatetime"}
	subprojCol = []string{"proj.pid", "proj.name", "proj.description"}
	blogCol    = []string{"blog.bid", "blog.name", "blog.like", "blog.hate", "blog.viewtime", "blog.super", "blog.belong", "blog.number", "blog.description", "blog.type", "blog.createtime", "blog.updatetime"}
)

func GetRoot(page string) (blogs []BlogList, err error) {
	err = db.SQL("call get_root(?)", page).Find(&blogs)
	if err != nil {
		return nil, err
	} else if len(blogs) == 0 {
		return nil, nil
	}
	return blogs, nil
}

// return owner page
func GetOwner(owner string) (*OwnerOut, error) {
	ownerData := &OwnerOut{}
	has, err := db.SQL("call get_owner(?)", owner).Get(ownerData)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, nil
	}
	return ownerData, nil
}

// Update an owner
func CreateOwner(uid, uniquename, nickname, descript string) error {
	return checkAffect(db.Exec("call create_owner(?, ?, ?, ?)", uid, uniquename, nickname, descript))
}

// Update an owner if no need to update uniquename, newuniname should be ""
func UpdateOwner(uid, oid, uniquename, newuniname, nickname, descript string) error {
	return checkResult(db.Exec("call update_owner(?, ?, ?, ?, ?, ?)", uid, oid, uniquename, newuniname, nickname, descript))
}

// delect an owner
func DelOwner(oid, uid, owner string) error {
	return checkAffect(db.Exec("call del_owner(?, ?, ?)", oid, uid, owner))
}

/*
// return project by name
func GetProject(url string) ([]ProjOut, error) {
	projDatas := make([]ProjOut, 0)
	err := db.SQL("call get_project(?)", url).Find(&projDatas)
	if err != nil {
		return nil, err
	} else if len(projDatas) == 0 {
		return nil, nil
	}
	return projDatas, nil
}

// update or insert a project
func PutProject(owner string, proj []string, oid, superid, super_url, pid, descript, url string) error {
	if len(proj) == 1 {
		return checkAffect(db.Exec("call put_root_project(?, ?, ?, ?, ?, ?)", owner, oid, pid, proj[0], descript, url))
	} else {
		return checkAffect(db.Exec("call put_sub_project(?, ?, ?, ?, ?, ?, ?, ?)", owner, oid, superid, super_url, pid, proj[len(proj)-1], descript, url))
	}
}

// delete project
func DelProject(oid, pid, urlpath string) error {
	return checkAffect(db.Exec("call del_project(?, ?, ?, ?)", oid, pid, urlpath))
}
*/
// get blog with owner super project and category data
func GetBlog(path string) (*BlogProj, error) {
	blogData := &BlogProj{}
	has, err := db.SQL("call get_blog(?)", path).Get(blogData)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, nil
	}
	return blogData, nil
}

// Update an blog
func CreateBlog(oid, superid, blog, descript, typeid, superUrl string) error {
	return checkAffect(db.Exec("call create_blog(?, ?, ?, ?, ?, ?)", oid, superid, blog, descript, typeid, superUrl))
}

// Update an owner if no need to update uniquename, newuniname should be ""
func UpdateBlog(oid, superid, newsuperid, bid, blog, newblog, descript, originUrl, newsuperUrl string) error {
	return checkResult(db.Exec("call update_blog(?, ?, ?, ?, ?, ?, ?, ?, ?)", oid, superid, newsuperid, bid, blog, newblog, descript, originUrl, newsuperUrl))
}

// delete a blog
func DelBlog(oid, bid, url string) error {
	return checkAffect(db.Exec("call del_blog(?, ?, ?)", oid, bid, url))
}

/*
// combine join of dynamic project name
func combineProjectsSQL(sb *sqlbuilder.SelectBuilder, proj []string) {
	sb.Join("project "+projAs[0], projAs[0]+".oid = owner.oid", sb.Equal(projAs[0]+".name", proj[0]))
	for i, l := 1, len(proj); i < l; i++ {
		sb.Join("project "+projAs[i], projAs[i]+".super = "+projAs[i-1]+".pid", sb.Equal(projAs[i]+".name", proj[i]))
	}
}

// return sql of get blog or project
func combineWorkSQL(sb *sqlbuilder.SelectBuilder, owner string, proj []string, blog string) {
	// select from owner
	sb.From("owner")
	// join project sql
	combineProjectsSQL(sb, proj)
	if blog != "" {
		// join blog SQL
		sb.Join("blog", "blog.super", sb.Equal("blog.name", blog))
	}
	sb.Where(sb.Equal("owner.uniquename", owner))
}

// return sql of selec column of getting blog or project
// if check is false, return more sql for client side
func getWorkSQL(check bool, owner string, proj []string, blog string, col ...string) (string, []interface{}) {
	sb := sqlbuilder.NewSelectBuilder()

	sb.Select(col...)
	combineWorkSQL(sb, owner, proj, blog)
	if !check {
		if blog == "" {
			// join sub project and blog
			sb.JoinWithOption(sqlbuilder.LeftJoin, "blog", projAs[len(proj)-1]+".pid = "+"blog.super")
			if len(proj) < len(projAs) {
				sb.JoinWithOption(sqlbuilder.LeftJoin, "project proj", projAs[len(proj)-1]+".pid = "+"proj.super")
			}
		}
		sb.JoinWithOption(sqlbuilder.LeftJoin, "category", "blog.belong = "+"category.cid")
		sb.Limit(20)
	}
	return sb.Build()
}

func GetBlogByName(owner string, proj []string, blog string) (*BlogProj, error) {
	blogData := &BlogProj{}
	log.Debug("", blogData)
	sql, args := getWorkSQL(false, owner, proj, blog, "owner.*", "blog.*", "category.name")
	has, err := db.SQL(sql, args...).Get(blogData)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, nil
	}
	return blogData, nil
}
*/

// check affect row is > 0 or not
func checkAffect(res sql.Result, err error) error {
	if err == nil {
		count, err := res.RowsAffected()
		if err == nil {
			if count > 0 {
				// affect success
				return nil
			} else {
				// affect fail
				return ERR_TASK_FAIL
			}
		}
		return err
	}
	return err
}

// check database return error or not
func checkResult(res sql.Result, err error) error {
	if err == nil {
		return nil
	}
	return err
}
