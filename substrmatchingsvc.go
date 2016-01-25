package substrmatchingsvc

import (
	"fmt"
	"github.com/blue-saber/fastmatching"
	"github.com/gin-gonic/gin"
	"net/http"
)

type InsertRequest struct {
	Keyword string `form:"keyword" json:"keyword" binding:"required"`
	Value   int32  `form:"value" json:"value" binding:"required"`
}

type RetrieveReq struct {
	Keyword string `form:"keyword" json:"keyword" binding:"required"`
}

type SubstrMatchingService struct {
	fmsvc *fastmatching.IFastMatching `@Autowired:"*"`
}

func (svc *SubstrMatchingService) Setfmsvc(f interface{}) {
	if origional, ok := f.(fastmatching.IFastMatching); ok {
		svc.fmsvc = &origional
	}
}

func (svc *SubstrMatchingService) PathPrefix() string {
	return "keyword"
}

func (svc *SubstrMatchingService) GetParam() string {
	return "param"
}

func (svc *SubstrMatchingService) InReleaseMode() bool {
	return false
}

func (svc *SubstrMatchingService) GetAll(c *gin.Context) {
	var req RetrieveReq

	if c.BindJSON(&req) == nil {
		result := (*svc.fmsvc).RetrieveData(req.Keyword)
		c.JSON(http.StatusOK, gin.H{"status": "success", "result": result})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "failed"})
	}
}
func (svc *SubstrMatchingService) DoGet(c *gin.Context) {
	param := c.Param("param")
	c.String(http.StatusOK, "You want %s", param)
}
func (svc *SubstrMatchingService) DoPost(c *gin.Context) {
	var req InsertRequest

	fmt.Printf("Host: %s, User: %s\n",
		c.Request.Host,
		c.Request.UserAgent())

	if c.BindJSON(&req) == nil {
		if (*svc.fmsvc).RegistData(req.Keyword, req.Value) {
			c.JSON(http.StatusOK, gin.H{"status": "success"})
		} else {
			c.JSON(http.StatusOK, gin.H{"status": "failed"})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{"status": "failed"})
	}
}
func (svc *SubstrMatchingService) DoPut(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "failed", "errmsg": "Not implement yet!"})
}

func (svc *SubstrMatchingService) DoDelete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "failed", "errmsg": "Not implement yet!"})
}

func (svc *SubstrMatchingService) DeleteAll(c *gin.Context) {
	(*svc.fmsvc).Clear()
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
