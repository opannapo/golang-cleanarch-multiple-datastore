package endpoints

import (
	super "app/app/base"
	"app/app/v1/apis/param"
	"app/app/v1/injection"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserEndpoint struct {
	services *injection.ServiceInjection
}

func (instance *UserEndpoint) getUser(c *gin.Context) {
	mysqlUser := instance.services.MysqlUserService
	req := c.Request.URL.Query()
	id, _ := strconv.Atoi(req.Get("id"))
	result, err := mysqlUser.GetUser(id)
	if err != nil {
		super.OutFailed(c, 0, err.Error())
	} else {
		super.OutOk(c, result)
	}
}

func (instance *UserEndpoint) getUsers(c *gin.Context) {
	mysqlUser := instance.services.MysqlUserService
	data, err := mysqlUser.GetUsers()
	//data, err := instance.FirestoreUserService.GetUsers()
	if err != nil {
		super.OutFailed(c, 0, err.Error())
	} else {
		super.OutOk(c, data)
	}
}

func (instance *UserEndpoint) addUser(c *gin.Context) {
	var p param.UserCreate
	err := c.ShouldBindJSON(&p)
	if err != nil {
		super.OutFailed(c, http.StatusBadRequest, err.Error())
		return
	}

	/*mysqlUser := instance.services.MysqlUserService
	mysqlTopic := instance.services.MysqlTopicTypeService
	mysqlUserFollowingTopic := instance.services.MysqlUserFollowingTopicService
	firestoreUser := instance.services.FirestoreUserService
	firestoreTopicType := instance.services.FirestoreTopicTypeService

	//Check Topic Type master table
	var topicTypeTmpToCreate []*entities.TopicType
	var topicTypeTmpExist []*entities.TopicType
	for i := range p.FollowingTopic {
		label := *p.FollowingTopic[i]
		topicExist, err := mysqlTopic.GetOneByLabel(label)
		if err != nil {
			tmp := entities.TopicType{
				Id:    0,
				Label: label}
			topicTypeTmpToCreate = append(topicTypeTmpToCreate, &tmp)
		} else {
			topicTypeTmpExist = append(topicTypeTmpExist, &topicExist)
		}
	}

	err, txInsertTopic := mysqlTopic.Inserts(topicTypeTmpToCreate)
	if err != nil {
		txInsertTopic.Rollback()
		super.OutFailed(c, 500, err.Error())
		return
	} else {
		topicTypeTmpExist = append(topicTypeTmpExist, topicTypeTmpToCreate...)
	}

	//Insert User
	err, txInsertUser := mysqlUser.AddUser(&p)
	if err != nil {
		txInsertTopic.Rollback()
		txInsertUser.Rollback()
		super.OutFailed(c, 500, err.Error())
		return
	}

	//Generate UserFollowingTopic & insert
	var tmpUserFollowingTopic []*entities.UserFollowingTopic
	for i := range topicTypeTmpExist {
		tmp := entities.UserFollowingTopic{
			UserId:      p.User.Id,
			TopicTypeId: topicTypeTmpExist[i].Id,
		}
		fmt.Println(tmp)
		tmpUserFollowingTopic = append(tmpUserFollowingTopic, &tmp)
		fmt.Println(tmpUserFollowingTopic)
	}
	err, txInsertUserFollowingTopic := mysqlUserFollowingTopic.Inserts(tmpUserFollowingTopic)
	if err != nil {
		txInsertTopic.Rollback()
		txInsertUser.Rollback()
		txInsertUserFollowingTopic.Rollback()
		super.OutFailed(c, 500, err.Error())
		return
	}

	//Insert FirestoreServicesInjected User
	err, _ = firestoreUser.AddUser(&p)
	if err != nil {
		txInsertTopic.Rollback()
		txInsertUser.Rollback()
		txInsertUserFollowingTopic.Rollback()
		super.OutFailed(c, 500, err.Error())
		return
	}

	err, _ = firestoreTopicType.Inserts(topicTypeTmpExist)
	if err != nil {
		txInsertTopic.Rollback()
		txInsertUser.Rollback()
		txInsertUserFollowingTopic.Rollback()
		super.OutFailed(c, 500, err.Error())
		return
	}

	txInsertTopic.Commit()
	txInsertUser.Commit()
	txInsertUserFollowingTopic.Commit()*/
	err = instance.services.MysqlUserService.AddUser(&p)
	if err != nil {
		super.OutFailed(c, 500, err.Error())
		return
	}
	super.OutOk(c, p)
}

func (instance *UserEndpoint) updateUser(c *gin.Context) {

}

func (instance *UserEndpoint) deleteUser(c *gin.Context) {

}

func NewUserEndpoint(g *gin.RouterGroup, services *injection.ServiceInjection) {
	instance := &UserEndpoint{
		services: services,
	}
	g.GET("user", instance.getUser)
	g.GET("users", instance.getUsers)
	g.POST("user/add", instance.addUser)
	g.POST("user/update", instance.updateUser)
	g.POST("user/delete", instance.deleteUser)
}
