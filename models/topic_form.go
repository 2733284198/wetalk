// Copyright 2013 wetalk authors
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package models

import (
	"github.com/astaxie/beego/validation"

	"github.com/beego/wetalk/utils"
)

type TopicAdminForm struct {
	Create    bool   `form:"-"`
	Id        int    `form:"-"`
	Name      string `valid:"Required;MaxSize(30)"`
	Intro     string `form:"type(textarea)" valid:"Required"`
	Slug      string `valid:"Required;MaxSize(100)"`
	Followers int    ``
	Order     int    ``
}

func (form *TopicAdminForm) Valid(v *validation.Validation) {
	qs := Topics()

	if CheckIsExist(qs, "Name", form.Name, form.Id) {
		v.SetError("Name", "Field value need unique")
	}

	if CheckIsExist(qs, "Slug", form.Slug, form.Id) {
		v.SetError("Slug", "Field value need unique")
	}
}

func (form *TopicAdminForm) SetFromTopic(topic *Topic) {
	utils.SetFormValues(topic, form)
}

func (form *TopicAdminForm) SetToTopic(topic *Topic) {
	utils.SetFormValues(form, topic, "Id")
}

type CategoryAdminForm struct {
	Create bool   `form:"-"`
	Id     int    `form:"-"`
	Name   string `valid:"Required;MaxSize(30)"`
	Slug   string `valid:"Required;MaxSize(100)"`
	Order  int    ``
}

func (form *CategoryAdminForm) Valid(v *validation.Validation) {
	qs := Categories()

	if CheckIsExist(qs, "Name", form.Name, form.Id) {
		v.SetError("Name", "Field value need unique")
	}

	if CheckIsExist(qs, "Slug", form.Slug, form.Id) {
		v.SetError("Slug", "Field value need unique")
	}
}

func (form *CategoryAdminForm) SetFromCategory(cat *Category) {
	utils.SetFormValues(cat, form)
}

func (form *CategoryAdminForm) SetToCategory(cat *Category) {
	utils.SetFormValues(form, cat, "Id")
}