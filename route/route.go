package route

import (
	"templategoapi/controller"
	"templategoapi/db"
	"templategoapi/middlewares"

	"github.com/gin-gonic/gin"
)

// MemberRoute create route
func NewRoute(r *gin.Engine, resource *db.Resource) {

	route := r.Group("/api")
	api := r.Group("/api")
	api.Use(middlewares.AuthToken(resource))
	api.Use(middlewares.AuthRequired())

	route.POST("/login", controller.Login(resource))

	//for test
	route.GET("/undeletealluser", controller.UndeleteAllUser(resource))

	route.POST("/upload-img/:path", controller.HandleUpload(resource))
	//route.GET("/IsLogin", controller.IsLogin(resource))
	{
		api.GET("/", controller.Test(resource))
		api.GET("/Logs", controller.GetAllLog(resource))
		api.POST("/Log", controller.GetLogByAccount(resource))
		api.GET("/AllProgramer", controller.GetAllrogramer(resource))

		api.GET("/GetUser", controller.GetUser(resource))
		api.GET("/logout", controller.LogOut(resource))
		api.GET("/showAllUser", controller.GetAllUser(resource))
		api.POST("/CheckDP", controller.CheckDefaultPassword(resource))
		api.PUT("/ChangeDP", controller.ChangeDefaultPassword(resource))
		api.POST("/CreateAccount", controller.CreateAccount(resource))
		api.GET("/IsLogin", controller.IsLogin(resource))
		api.POST("/showOneUser", controller.GetOneUser(resource))
		api.DELETE("/deleteOneUser", controller.DeleteOneUser(resource))
		api.PUT("/updateUser", controller.UpdateOneUser(resource))
		api.PUT("/updateProfile", controller.UpdateProfile(resource))

		api.POST("/active", controller.ActiveUser(resource))
		api.POST("/inactive", controller.InActiveUser(resource))
		api.PUT("/ChangePassword", controller.ChangePassword(resource))

		/////// Ticket
		api.POST("/CreateTicket", controller.CreateTicket(resource))
		api.PUT("/updateTicketCS", controller.CsUpdateTicket(resource))
		api.GET("/ShowAllTicket", controller.GetAllTicket(resource))
		api.PUT("/updateTicket", controller.UserUpdateTicket(resource))
		// api.PUT("/ChangeTicketStatus", controller.ChangeTicketStatus(resource))
		api.DELETE("/deleteOneTicket", controller.DeleteOneTicket(resource))
		api.DELETE("/CancelOneTicket", controller.CancelOneTicket(resource))
		api.GET("/showPendingTicket", controller.GetPendingTicket(resource))
		api.GET("/showCancelTicket", controller.GetCancelTicket(resource))
		api.GET("/showDeleteTicket", controller.GetDeleteTicket(resource))
		api.GET("/showSuccessTicket", controller.GetSuccessTicket(resource))
		api.GET("/showInfixTicket", controller.GetInfixTicket(resource))
		api.POST("/SearchTicket", controller.SearchTicket(resource))
		api.PUT("/PassCase", controller.PassToProgramer(resource))
		api.POST("/TicketPG", controller.GetAllTicketProgramer(resource))
		api.PUT("/TicketSuccess", controller.SuccessOneTicket(resource))
		api.POST("/GetTicket", controller.GetOneTicket(resource))
		api.POST("/TicketUser", controller.GetAllTicketByAccount(resource))
		api.POST("/showPendingTicket", controller.GetPendingTicketByAccount(resource))
		api.POST("/showInfixTicket", controller.GetInfixTicketByAccount(resource))
		api.POST("/showSuccessTicket", controller.GetSuccessTicketByAccount(resource))
		api.POST("/showCancelTicket", controller.GetCancelTicketByAccount(resource))
		api.POST("/Reporting", controller.GetPendingTicketCS(resource))

		// Dashboard
		api.GET("/AGGYMD", controller.AggregateYMD(resource))
		api.GET("/AGGYM", controller.AggregateYM(resource))
		api.GET("/AGGY", controller.AggregateY(resource))
		api.GET("/AGGW", controller.AggregateWeek(resource))
		api.GET("/AGGCountPendingYM", controller.AggregateCountPendingYM(resource))
		api.GET("/AGGCountInfixYM", controller.AggregateCountInfixYM(resource))
		api.GET("/AGGCountPendingYMD", controller.AggregateCountPendingYMD(resource))
		api.GET("/AGGCountInfixYMD", controller.AggregateCountInfixYMD(resource))
		api.GET("/AGGCountSuccessYM", controller.AggregateCountSuccessYM(resource))
		api.GET("/AGGCountSuccessYMD", controller.AggregateCountSuccessYMD(resource))
		api.GET("/Ticket_Count_All", controller.AggregateTicket_Count_All(resource))
		api.GET("/Type_Ticket_Count_YMD", controller.Aggregate_TypeTicket_Count_YMD(resource))
		api.GET("/TypeTicket_Count_YM", controller.Aggregate_TypeTicket_Count_YM(resource))
		api.GET("/TypeTicket_Count_Y", controller.Aggregate_TypeTicket_Count_Y(resource))
		api.GET("/TypeTicket_Count_All", controller.Aggregate_TypeTicket_Count_All(resource))

		/////// Product
		api.POST("/CreateProduct", controller.CreateProduct(resource))
		api.PUT("/UpdateProduct", controller.UpdateOneProduct(resource))
		api.DELETE("/DeleteProduct", controller.DeleteOneProduct(resource))
		api.GET("/ShowProduct", controller.GetAllProduct(resource))
		api.GET("/DropdownShowProduct", controller.DropdownGetAllProduct(resource))
		/////// Type Ticket
		api.POST("/CreateTypeTicket", controller.CreateTypeTicket(resource))
		api.PUT("/UpdateTypeTicket", controller.UpdateOneTypeTicket(resource))
		api.DELETE("/DeleteTypeTicket", controller.DeleteOneTypeTicket(resource))
		api.GET("/ShowTypeTicket", controller.GetAllTypeTicket(resource))
		api.GET("/DropdownShowTypeTicket", controller.DropdownGetAllTypeTicket(resource))
		/////// Tag Ticket
		api.POST("/CreateTagTicket", controller.CreateTagTicket(resource))
		api.PUT("/UpdateTagTicket", controller.UpdateOneTagTicket(resource))
		api.DELETE("/DeleteTagTicket", controller.DeleteOneTagTicket(resource))
		api.GET("/ShowTagTicket", controller.GetAllTagTicket(resource))
		api.GET("/DropdownShowTagTicket", controller.DropdownGetAllTagTicket(resource))
		/////// Role
		api.POST("/CreateRole", controller.CreateRole(resource))
		api.PUT("/UpdateRole", controller.UpdateOneRole(resource))
		api.DELETE("/DeleteRole", controller.DeleteOneRole(resource))
		api.GET("/ShowRole", controller.GetAllRole(resource))
		api.POST("/ShowRoleByRole_name", controller.GetRoleByRole_name(resource))
		api.PUT("/UpdateRoleByRole_name", controller.UpdateOneRoleByRole_name(resource))
		api.GET("/DropdownShowRole", controller.DropdownGetAllRole(resource))

		/////// Forum
		api.POST("/CreateKnowledge", controller.CreateKnowledge(resource))
		api.GET("/ShowKnowledge", controller.GetAllKnowledge(resource))
		api.PUT("/UpdateKnowledge", controller.UpdateKnowledge(resource))
		api.DELETE("/DeleteKnowledge", controller.DeleteOneKnowledge(resource))
		api.POST("/SearchKnowledge", controller.SearchKnowledge(resource))
		api.POST("/GetOneKnowledge", controller.GetOneKnowledge(resource))
		api.GET("/ShowSolution", controller.KnowledgeFromTicket(resource))

		/////// Customer
		api.POST("/CreateCustomer", controller.CreateCustomer(resource))
		api.GET("/ShowCustomer", controller.GetAllCustomer(resource))
		api.PUT("/UpdateCustomer", controller.UpdateOneCustomer(resource))
		api.DELETE("/DeleteCustomer", controller.DeleteOneCustomer(resource))
		api.POST("/showOneCustomer", controller.GetOneCustomer(resource))
		api.GET("/DropdownShowCustomer", controller.DropdownGetAllCustomer(resource))

		//TEST
		api.GET("/TestPagination", controller.GetAllUserTEST(resource))
	}

}
