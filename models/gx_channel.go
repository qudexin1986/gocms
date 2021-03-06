package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"reflect"
	"strings"
)

type GxChannel struct {
	Id           int    `orm:"column(id);auto"`
	Pid          int16  `orm:"column(pid)"`
	Status       int8   `orm:"column(status)"`
	Cname        string `orm:"column(cname);size(20)"`
	Ctpl         string `orm:"column(ctpl);size(20)"`
	Ctitle       string `orm:"column(ctitle);size(50)"`
	Ckeywords    string `orm:"column(ckeywords);size(255)"`
	Cdescription string `orm:"column(cdescription);size(255)"`
}

func init() {
	orm.RegisterModel(new(GxChannel))
	orm.RegisterDataBase("default", "mysql", "root:1987123@/gocms?charset=utf8")
}

// AddGxChannel insert a new GxChannel into database and returns
// last inserted Id on success.
func AddGxChannel(m *GxChannel) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetGxChannelById retrieves GxChannel by Id. Returns error if
// Id doesn't exist
func GetGxChannelById(id int) (v *GxChannel, err error) {
	o := orm.NewOrm()
	v = &GxChannel{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllGxChannel retrieves all GxChannel matches certain condition. Returns empty list if
// no records exist
func GetAllGxChannel(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(GxChannel))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []GxChannel
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateGxChannel updates GxChannel by Id and returns error if
// the record to be updated doesn't exist
func UpdateGxChannelById(m *GxChannel) (err error) {
	o := orm.NewOrm()
	v := GxChannel{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteGxChannel deletes GxChannel by Id and returns error if
// the record to be deleted doesn't exist
func DeleteGxChannel(id int) (err error) {
	o := orm.NewOrm()
	v := GxChannel{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&GxChannel{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
