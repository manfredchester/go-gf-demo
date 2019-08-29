package main

import (
	"github.com/emicklei/go-restful"
)

const (
	PATH    = "/api"
	PATH_V1 = "/v1"
)

// restful 服务初始化
var ws = new(restful.WebService)

// restful容器初始化
var wc = restful.NewContainer()

// 获取配置restful容器
func GetRouterContainer() *restful.Container {
	ws.Consumes(restful.MIME_XML, restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON, restful.MIME_XML)
	ws.Path(PATH)
	routerConf()

	for _, route := range ws.Routes() {
		// logs.Info("route url:%s > %s", route.Method, route.Path)
	}

	wc.Add(ws)
	return wc
}
func routerConf() {
	cloudBillRouterConf()
	cmdbRouterConf()
	enumRouterConf()
}

func cloudBillRouterConf() {
	rootPath := PATH_V1 + "/cloudbill/"

	methodPath := "QueryAccountDetail"
	ws.Route(ws.POST(rootPath + methodPath).To(cloudBill.QueryAccountDetail))

	methodPath = "QueryAccount"
	ws.Route(ws.POST(rootPath + methodPath).To(cloudBill.QueryAccount))

	methodPath = "NewAccount"
	ws.Route(ws.POST(rootPath + methodPath).To(cloudBill.NewAccount))

	methodPath = "UpdateAccount"
	ws.Route(ws.POST(rootPath + methodPath).To(cloudBill.UpdateAccount))

	methodPath = "DeleteAccount"
	ws.Route(ws.POST(rootPath + methodPath).To(cloudBill.DeleteAccount))

	methodPath = "QueryCpp"
	ws.Route(ws.POST(rootPath + methodPath).To(cloudBill.QueryCpp))

	methodPath = "NewCpp"
	ws.Route(ws.POST(rootPath + methodPath).To(cloudBill.NewCpp))

	methodPath = "UpdateCpp"
	ws.Route(ws.POST(rootPath + methodPath).To(cloudBill.UpdateCpp))

	methodPath = "DeleteCpp"
	ws.Route(ws.POST(rootPath + methodPath).To(cloudBill.DeleteCpp))

	methodPath = "CheckAccount"
	ws.Route(ws.POST(rootPath + methodPath).To(cloudBill.CheckAccount))

	methodPath = "BillAnalysis"
	ws.Route(ws.POST(rootPath + methodPath).To(cloudBill.BillAnalysis))

	methodPath = "BillImport"
	ws.Route(ws.POST(rootPath + methodPath).To(cloudBill.BillImport))

	methodPath = "IsBillSuper"
	ws.Route(ws.POST(rootPath + methodPath).To(cloudBill.IsSuper))

	methodPath = "BillDetail"
	ws.Route(ws.POST(rootPath + methodPath).To(cloudBill.BillDetail))

	methodPath = "BillDelete"
	ws.Route(ws.POST(rootPath + methodPath).To(cloudBill.BillDelete))

	methodPath = "InitDataTime"
	ws.Route(ws.POST(rootPath + methodPath).To(cloudBill.InitDataTime))

	methodPath = "HasLinkApp"
	ws.Route(ws.GET(rootPath + methodPath).To(cloudBill.HasLinkApp))
}

func cmdbRouterConf() {
	rootPath := PATH_V1 + "/cmdb/"
	cornPath := PATH_V1 + "/cmdbCron/"
	resPath := PATH_V1 + "/cmdbRes/"
	appPath := PATH_V1 + "/app/"

	methodPath := "CloudInfo"
	ws.Route(ws.POST(rootPath + methodPath).To(cmdb.CloudInfo))
	// 查询云类型获取（修改）
	methodPath = "CloudInfoV2"
	ws.Route(ws.POST(rootPath + methodPath).To(cmdb.CloudInfoV2))
	// cmdb主页
	methodPath = "dashbord"
	ws.Route(ws.GET(rootPath + methodPath).To(cmdb.Dashbord))
	methodPath = "Share"
	ws.Route(ws.POST(rootPath + methodPath).To(cmdb.Share))
	methodPath = "Delete"
	ws.Route(ws.POST(rootPath + methodPath).To(cmdb.Delete))
	methodPath = "Query"
	ws.Route(ws.POST(rootPath + methodPath).To(cmdb.Query))
	methodPath = "QueryV2"
	ws.Route(ws.POST(rootPath + methodPath).To(cmdb.QueryV2))
	//从应用删除
	methodPath = "AllocatedAppDel"
	ws.Route(ws.POST(rootPath + methodPath).To(cmdb.AllocatedAppDel))
	//应用删除列表
	methodPath = "AllocatedAppList"
	ws.Route(ws.GET(rootPath + methodPath).To(cmdb.AllocatedAppList))
	//分配到应用
	methodPath = "AllocatedApp"
	ws.Route(ws.POST(rootPath + methodPath).To(cmdb.AllocatedAppPost))
	//分配应用列表
	methodPath = "AppList"
	ws.Route(ws.GET(rootPath + methodPath).To(cmdb.AppList))
	//导出
	methodPath = "Export"
	ws.Route(ws.POST(rootPath + methodPath).To(cmdb.Export))

	methodPath = "Pause"
	ws.Route(ws.POST(cornPath + methodPath).To(cmdbCron.Pause))
	methodPath = "Restore"
	ws.Route(ws.POST(cornPath + methodPath).To(cmdbCron.Restore))
	methodPath = "Execute"
	ws.Route(ws.POST(cornPath + methodPath).To(cmdbCron.Execute))
	methodPath = "LogList"
	ws.Route(ws.POST(cornPath + methodPath).To(cmdbCron.LogList))
	methodPath = "LogDetail"
	ws.Route(ws.POST(cornPath + methodPath).To(cmdbCron.LogDetail))
	methodPath = "List"
	ws.Route(ws.POST(cornPath + methodPath).To(cmdbCron.List))

	methodPath = "ResList"
	ws.Route(ws.POST(resPath + methodPath).To(cmdbRes.ResList))

	methodPath = "ResListV2"
	ws.Route(ws.POST(resPath + methodPath).To(cmdbRes.ResListV2))

	methodPath = "ResTurn"
	ws.Route(ws.POST(resPath + methodPath).To(cmdbRes.ResTurn))

	methodPath = "ResField"
	ws.Route(ws.POST(resPath + methodPath).To(cmdbRes.ResField))
	methodPath = "ResFieldUnique"
	ws.Route(ws.POST(resPath + methodPath).To(cmdbRes.ResFieldUnique))
	methodPath = "ResFieldGroup"
	ws.Route(ws.POST(resPath + methodPath).To(cmdbRes.ResFieldGroup))

	// app path
	methodPath = "CloudInfo"
	ws.Route(ws.POST(appPath + methodPath).To(appRes.CloudInfo))
	methodPath = "Query"
	ws.Route(ws.POST(appPath + methodPath).To(appRes.Query))
}

func enumRouterConf() {
	rootPath := PATH_V1 + "/enum/"

	methodPath := "getenum"
	ws.Route(ws.GET(rootPath + methodPath).To(enumApi.Getenum))

	methodPath = "timestamp"
	ws.Route(ws.GET(rootPath + methodPath).To(enumApi.Timestamp))
}
