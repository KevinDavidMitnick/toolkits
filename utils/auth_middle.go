// Copyright 2017 Xiaomi, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utils

import (
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
)

func CORS() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		context.Writer.Header().Set("Access-Control-Max-Age", "86400")
		context.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE, PATCH")
		context.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Apitoken")
		context.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		context.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if context.Request.Method == "OPTIONS" {
			context.AbortWithStatus(200)
		} else {
			context.Next()
		}
	}
}

func SetParaDefault(o interface{}) {
	v := reflect.ValueOf(o)

	if v.Kind() == reflect.Ptr && !v.Elem().CanSet() { //判断是否为指针类型 元素是否可以修改
		return
	} else {
		v = v.Elem() //实际取得的对象
	}

	//判断字段是否存在
	f := v.FieldByName("Name")
	if f.IsValid() {
		if f.Kind() == reflect.String {
			f.SetString(".*")
		}
	}

	email := v.FieldByName("Email")
	if email.IsValid() {
		if email.Kind() == reflect.String {
			email.SetString(".*")
		}
	}

	phone := v.FieldByName("Phone")
	if phone.IsValid() {
		if phone.Kind() == reflect.String {
			phone.SetString(".*")
		}
	}
	createUser := v.FieldByName("Create_User")
	if createUser.IsValid() {
		if createUser.Kind() == reflect.String {
			createUser.SetString(".+")
		}
	}
	createRole := v.FieldByName("Create_Role")
	if createRole.IsValid() {
		if createRole.Kind() == reflect.String {
			createRole.SetString(".+")
		}
	}

	//IsSelfService
	isSelfService := v.FieldByName("IsSelfService")
	if isSelfService.IsValid() {
		if isSelfService.Kind() == reflect.Int {
			isSelfService.SetInt(0)
		}
	}

	//SelfServiceName/	StartTime
	selfServiceName := v.FieldByName("SelfServiceName")
	if selfServiceName.IsValid() {
		if selfServiceName.Kind() == reflect.String {
			selfServiceName.SetString(".+")
		}
	}

	//	EndTime
	//	HostGroupName
	//	DataCenter
	//	Metrics
	//	MachineRoom
	//	Endpoints
	//	ProcessStatus
	//	Priority
	//	PageSize
	//	Status
	//	PageNumber
	//Name        string `json:"name"`
	//Email       string `json:"email"`
	//Phone       string `json:"phone"`
	//Create_User string `json:"createUser"`
	//Create_Role string `json:"createRole"`
	//GrpName      string `json:"name"`
	//TplName      string `json:"tplName"`
	//ResourceType string `json:"resourceType"`
	//Is_SelfUser  int    `json:"isSelfuser"`
	//SelfUserName string `json:"selfuserName"`
	//Create_User  string `json:"createUser"`
	//Create_Role  string `json:"createRole"`
	grpName := v.FieldByName("GrpName")
	if grpName.IsValid() {
		if grpName.Kind() == reflect.String {
			grpName.SetString(".+")
		}
	}

	tplName := v.FieldByName("TplName")
	if tplName.IsValid() {
		if tplName.Kind() == reflect.String {
			tplName.SetString(".+")
		}
	}

	resourceName := v.FieldByName("Resource_name")
	if resourceName.IsValid() {
		if resourceName.Kind() == reflect.String {
			resourceName.SetString(".+")
		}
	}

	selfName := v.FieldByName("SelfUserName")
	if selfName.IsValid() {
		if selfName.Kind() == reflect.String {
			selfName.SetString(".+")
		}
	}

	st := v.FieldByName("StartTime")
	if st.IsValid() {
		if st.Kind() == reflect.Int64 {
			st.SetInt(0)
		}
	}

	et := v.FieldByName("EndTime")
	if et.IsValid() {
		if et.Kind() == reflect.Int64 {
			et.SetInt(time.Now().Unix())
		}
	}

	//pageNumber
	pn := v.FieldByName("PageNumber")
	if pn.IsValid() {
		if pn.Kind() == reflect.Int {
			pn.SetInt(1)
		}
	}

	//pageSize
	ps := v.FieldByName("PageSize")
	if ps.IsValid() {
		if ps.Kind() == reflect.Int {
			ps.SetInt(10)
		}
	}

	//priority
	po := v.FieldByName("Priority")
	if po.IsValid() {
		if po.Kind() == reflect.Int {
			po.SetInt(-1)
		}
	}

	//alarmStype
	alarmStyle := v.FieldByName("AlarmStyle")
	if alarmStyle.IsValid() {
		if alarmStyle.Kind() == reflect.String {
			alarmStyle.SetString(".+")
		}
	}
	searchUser := v.FieldByName("SearchUser")
	if searchUser.IsValid() {
		if searchUser.Kind() == reflect.String {
			searchUser.SetString(".+")
		}
	}
	searchRole := v.FieldByName("SearchRole")
	if searchRole.IsValid() {
		if searchRole.Kind() == reflect.String {
			searchRole.SetString(".+")
		}
	}

	tag := v.FieldByName("Tag")
	if tag.IsValid() {
		if tag.Kind() == reflect.String {
			tag.SetString(".*")
		}
	}

	tp := v.FieldByName("Type")
	if tp.IsValid() {
		if tp.Kind() == reflect.String {
			tp.SetString(".*")
		}
	}

	group := v.FieldByName("Group")
	if group.IsValid() {
		if group.Kind() == reflect.String {
			group.SetString(".*")
		}
	}

	groupName := v.FieldByName("GroupName")
	if groupName.IsValid() {
		if groupName.Kind() == reflect.String {
			groupName.SetString(".*")
		}
	}

	statusAlarm := v.FieldByName("StatusAlarm")
	if statusAlarm.IsValid() {
		if statusAlarm.Kind() == reflect.Int {
			statusAlarm.SetInt(3)
		}
	}
	levelAlarm := v.FieldByName("LevelAlarm")
	if levelAlarm.IsValid() {
		if levelAlarm.Kind() == reflect.Int {
			levelAlarm.SetInt(3)
		}
	}
	return
}
