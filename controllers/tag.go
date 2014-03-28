package controllers

import (
	"github.com/beego/admin/src/rbac"
	m "github.com/wangbinxiang/codesave/models"
)

type TagController struct {
	rbac.CommonController
}

const TagPageSize int64 = 20

func (this *TagController) Index() {
	if this.IsAjax() {
		page, _ := this.GetInt("page")
		pageSize, _ := this.GetInt("rows")
		if pageSize <= 0 {
			pageSize = TagPageSize
		}

		tagInfos, _, count, err := m.GetTagLabelList(page, pageSize, true)
		if err != nil {
			this.Rsp(false, err.Error())
		} else {
			this.Data["json"] = &map[string]interface{}{"total": count, "rows": &tagInfos}
			this.ServeJson()
		}
	} else {
		this.TplNames = this.GetTemplatetype() + "/cs/tag.tpl"
	}
}

func (this *TagController) Add() {
	tagLabel := m.TagLabel{}

	if err := this.ParseForm(&tagLabel); err != nil {
		this.Rsp(false, err.Error())
		return
	}

	var id int64
	var err error

	Tid, _ := this.GetInt("Id")
	if Tid > 0 {
		id, err = m.UpdateTagLabel(&tagLabel)
	} else {
		id, err = m.AddTagLabel(&tagLabel)
	}
	if err == nil && id > 0 {
		this.Rsp(true, "Success")
		return
	} else {
		this.Rsp(false, err.Error())
		return
	}
}
