package api

import (
	"context"
	"douyin/app/model"
	"douyin/app/service"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tencentyun/cos-go-sdk-v5"
)

// 获取视频流
func Getfeed(c *gin.Context) {
	//定义返回值
	feedResponse := &model.Feed{}

	//获取参数
	latestTime := c.Query("latest_time")
	if latestTime == "0" {
		latestTime = strconv.FormatInt(time.Now().UnixNano()/1e6, 10)
	}
	v1, err := strconv.ParseInt(latestTime, 10, 64)
	token := c.Query("token")
	//获取视频流
	videoList, err := service.Getfeed(v1, token)
	if err != nil {
		fmt.Println("获取视频流失败：%v", err)
		statusMsg := "获取视频流失败"
		feedResponse.StatusCode = -1
		feedResponse.StatusMsg = &statusMsg
		c.JSON(http.StatusOK, feedResponse)
		return
	}

	//返回值
	statusMsg := "请求成功"
	nextTime := time.Now().UnixNano() / 1e6

	if len(videoList) != 0 {
		lastVideo := videoList[len(videoList)-1:]
		nextTime = lastVideo[0].PublicTime

	}
	fmt.Println("nextTime：", nextTime)

	feedResponse.StatusCode = 0
	feedResponse.StatusMsg = &statusMsg
	feedResponse.VideoList = videoList[:]
	feedResponse.NextTime = &nextTime
	fmt.Println(feedResponse)
	c.JSON(http.StatusOK, feedResponse)
}

//发布
func GetPublicList(c *gin.Context) {
	//定义返回值
	feedResponse := &model.PublishListResp{}

	//获取参数
	userId := c.Query("user_id")
	token := c.Query("token")
	//获取视频流
	videoList, err := service.PublishList(userId, token)
	if err != nil {
		fmt.Println("获取发布列表失败：%v", err)
		statusMsg := "获取发布列表失败"
		feedResponse.StatusCode = -1
		feedResponse.StatusMsg = &statusMsg
		c.JSON(http.StatusOK, feedResponse)
		return
	}

	//返回值
	statusMsg := "请求成功"

	feedResponse.StatusCode = 0
	feedResponse.StatusMsg = &statusMsg
	feedResponse.VideoList = videoList[:]

	c.JSON(http.StatusOK, feedResponse)
}

//发布
func PublishAction(c *gin.Context) {

	//获取参数
	token := c.Query("token")
	title := c.Query("title")

	file, _ := c.FormFile("date")
	fmt.Println(token, title, file)
	// 上传文件到指定的路径
	dst := "./test"
	c.SaveUploadedFile(file, dst)

	// 存储桶名称，由bucketname-appid 组成，appid必须填入，可以在COS控制台查看存储桶名称。 https://console.cloud.tencent.com/cos5/bucket
	// 替换为用户的 region，存储桶region可以在COS控制台“存储桶概览”查看 https://console.cloud.tencent.com/ ，关于地域的详情见 https://cloud.tencent.com/document/product/436/6224 。
	u, _ := url.Parse("https://peng-1311860297.cos.ap-beijing.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	client := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			// 通过环境变量获取密钥
			// 环境变量 SECRETID 表示用户的 SecretId，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretID: "AKIDqUs3ze3B2btue6HBn8D2XlhMlumNimY4",
			// 环境变量 SECRETKEY 表示用户的 SecretKey，登录访问管理控制台查看密钥，https://console.cloud.tencent.com/cam/capi
			SecretKey: "GFxNWOzJM9ftra1C6j7E4dvwg8KFsEI6",
		},
	})

	// Case1 使用 Put 上传对象
	// key := "exampleobject"
	// f, err := os.Open("../test")
	opt := &cos.ObjectPutOptions{
		ObjectPutHeaderOptions: &cos.ObjectPutHeaderOptions{
			ContentType: "text/html",
		},
		ACLHeaderOptions: &cos.ACLHeaderOptions{
			// 如果不是必要操作，建议上传文件时不要给单个文件设置权限，避免达到限制。若不设置默认继承桶的权限。
			XCosACL: "private",
		},
	}
	// _, err = client.Object.Put(context.Background(), key, f, opt)
	// if err != nil {
	// 	panic(err)
	// }

	// Case 2 使用 PUtFromFile 上传本地文件到COS
	key := "exampleobject.mp4"
	filepath := "./test"
	_, err := client.Object.PutFromFile(context.Background(), key, filepath, opt)
	if err != nil {
		panic(err)
	}

	// Case 3 上传 0 字节文件, 设置输入流长度为 0
	// key := "exampleobject"
	// _, err := client.Object.Put(context.Background(), key, file., nil)
	// if err != nil {
	// 	// ERROR
	// }

}
