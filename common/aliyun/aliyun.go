package aliyun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
)

func GenerateStsInfo(roleArn string, regin string, accessKeyId string, accessKeySecret string) (*sts.AssumeRoleResponse, error) {
	client, err := sts.NewClientWithAccessKey(regin, accessKeyId, accessKeySecret)
	if err != nil {
		return nil, err
	}
	request := sts.CreateAssumeRoleRequest()
	request.Scheme = "https"
	request.RoleArn = roleArn
	request.RoleSessionName = "libra"
	response, err := client.AssumeRole(request)
	return response, nil
}

type PolicyToken struct {
	Expire          string `json:"expire"`
	Signature       string `json:"signature"`
	Policy          string `json:"policy"`
	Dir             string `json:"dir"`
	AccessId        string `json:"accessId"`
	Host            string `json:"host"`
	AccessKeySecret string `json:"accessKeySecret"`
	SecurityToken   string `json:"securityToken"`
	Expiration      string `json:"expiration"`
}

type ConfigStruct struct {
	Expiration string     `json:"expiration"`
	Conditions [][]string `json:"conditions"`
}

//func AppAliyunPolicy(c *gin.Context) {
//	uploadDir := c.DefaultQuery("dir", "user-dir/")
//	//获取token2中的accessKeyId,accessKeySecret
//	resp := generateStsInfo()
//
//	now := time.Now().Unix()
//	expire_end := now + expire_time
//	var tokenExpire = getGmtIso8601(expire_end)
//
//	//create post policy json
//	var config ConfigStruct
//	config.Expiration = tokenExpire
//	var condition []string
//	condition = append(condition, "starts-with")
//	condition = append(condition, "$key")
//	condition = append(condition, uploadDir)
//	config.Conditions = append(config.Conditions, condition)
//
//	//calucate signature
//	result, err := json.Marshal(config)
//	debyte := base64.StdEncoding.EncodeToString(result)
//	h := hmac.New(func() hash.Hash { return sha1.New() }, []byte(*resp.Credentials.AccessKeySecret))
//	io.WriteString(h, debyte)
//	signedStr := base64.StdEncoding.EncodeToString(h.Sum(nil))
//
//	policyToken := &PolicyToken{}
//	policyToken.AccessKeyId = *resp.Credentials.AccessKeyId
//	policyToken.AccessKeySecret = *resp.Credentials.AccessKeySecret
//	policyToken.SecurityToken = *resp.Credentials.SecurityToken
//	policyToken.Expiration = *resp.Credentials.Expiration
//	policyToken.Expire = expire_end
//	policyToken.Signature = string(signedStr)
//	policyToken.Directory = uploadDir
//	policyToken.Policy = string(debyte)
//	policyToken.UploadOssUrl = UploadOssUrl
//	if err != nil {
//		fmt.Println("json err:", err)
//	}
//	c.JSON(http.StatusOK, gin.H{
//		"code": 1,
//		"data": policyToken,
//	})
//	return
//}
