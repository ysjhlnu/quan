package asset

import (
	"database/sql"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"time"
)

var db *gorm.DB

type Hosts struct {
	ID         uint `gorm:"primarykey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  sql.NullTime `gorm:"index" json:"-"`
	Env       string `json:"env" form:"env" gorm:"column:env;comment:环境;type:varchar(50);size:50;"`
	Instanceid string `json:"instanceid" form:"instanceid" gorm:"column:instanceid;comment:实例ID;type:varchar(50);size:50;"`
	Name       string `json:"name" form:"name" gorm:"column:name;comment:实例名字;type:varchar(50);size:50;"`
	Privateip  string `json:"privateip" form:"privateip" gorm:"column:privateip;comment:内网IP;type:varchar(50);size:50;"`
	Status     string `json:"status" form:"status" gorm:"column:status;comment:运行状态;type:varchar(50);size:50;"`
	Type       string `json:"type" form:"type" gorm:"column:type;comment:实例模板;type:varchar(50);size:50;"`
	Region       string `json:"region" form:"region" gorm:"column:region;comment:地区;type:varchar(50);size:50;"`
	Ps         string `json:"ps" form:"ps" gorm:"column:ps;comment:备注;type:varchar(50);size:50;"`
}

func (Hosts) TableName() string {
	return "hosts"
}

func GetInstances(region map[string]string, newInstanceList []string) ([]string, error) {

	for ps, v := range region {
		fmt.Println(ps,v)
		sess, _ := session.NewSession(&aws.Config{
			Credentials: credentials.NewStaticCredentials("1111111111111111111", "222222222222222222", ""),
			Region:      aws.String(v),
		})
		svc := ec2.New(sess)
		result, err := svc.DescribeInstances(nil)
		if err != nil {
			return nil, err
		}

		for _, r := range result.Reservations {
			//fmt.Println(r)
			for _, i := range r.Instances {
				var a Hosts
				newInstanceList = append(newInstanceList, *i.InstanceId)
				err = db.Where("instanceid = ?", *i.InstanceId).Where("deleted_at is null").Unscoped().First(&a).Error
				if a.Instanceid != "" {
					asset := Hosts{
						Instanceid: *i.InstanceId,
						Type:       *i.InstanceType,
						Privateip:  *i.PrivateIpAddress,
						Status:     *i.State.Name,
						Name:       *i.Tags[0].Value,
						CreatedAt:  *i.LaunchTime,
						Env: ps,
						Region: v,
					}
					db.Model(&Hosts{}).Where("instanceid = ?", i.InstanceId).Updates(asset)
				} else {
					asset := Hosts{
						Instanceid: *i.InstanceId,
						Type:       *i.InstanceType,
						Privateip:  *i.PrivateIpAddress,
						Status:     *i.State.Name,
						Name:       *i.Tags[0].Value,
						CreatedAt:  *i.LaunchTime,
						Env: ps,
						Region: v,
					}
					err = db.Create(&asset).Error
				}
			}
		}
	}

	return newInstanceList, nil
}


func GetAliyunInstances( newInstanceList []string) ([]string, error) {

	region := []string{"cn-beijing","cn-hongkong"}

	for _, v := range region {
		fmt.Println("测试",v)
		client, _ := ecs.NewClientWithAccessKey(v, "1111111111111111", "22222222222222222222222")
		request := ecs.CreateDescribeInstancesRequest()
		response, _ := client.DescribeInstances(request)

		r := response.Instances

		for _, i := range r.Instance {
			var a Hosts
			newInstanceList = append(newInstanceList, i.InstanceId)
			_ = db.Where("instanceid = ?", i.InstanceId).Where("deleted_at is null").Unscoped().First(&a).Error
			if a.Instanceid != "" {
				local, _ := time.LoadLocation("Asia/Shanghai")
				tt,_ := time.ParseInLocation("2006-01-02T15:04Z", i.CreationTime,local)
				asset := Hosts{
					Instanceid: i.InstanceId,
					Type:       i.InstanceType,
					Privateip:  i.NetworkInterfaces.NetworkInterface[0].PrimaryIpAddress,
					Status:     i.Status,
					Name:       i.InstanceName,
					CreatedAt:  tt,
					Env:  "测试",
					Region: v,
				}
				db.Model(&Hosts{}).Where("instanceid = ?", i.InstanceId).Updates(asset)
			} else {
				local, _ := time.LoadLocation("Asia/Shanghai")
				tt,_ := time.ParseInLocation("2006-01-02T15:04Z", i.CreationTime,local)
				asset := Hosts{
					Instanceid: i.InstanceId,
					Type:       i.InstanceType,
					Privateip:  i.NetworkInterfaces.NetworkInterface[0].PrimaryIpAddress,
					Status:     i.Status,
					Name:       i.InstanceName,
					CreatedAt:  tt,
					Env: "测试",
					Region: v,
				}
				_ = db.Create(&asset).Error
			}
		}

	}

	return newInstanceList, nil
}


func AssetHostAwsUpdate() {

	var err error
	db, _ = gorm.Open("mysql", fmt.Sprintf("root:123456@tcp(192.168.111.129)/quan?charset=utf8&parseTime=True&loc=Local"))
	var newInstanceList []string


	regionMap := map[string]string{"测试":"ap-east-1","线上":"ap-northeast-1"}

	newInstanceList, err = GetInstances(regionMap, newInstanceList)
	if err != nil {
		fmt.Println("Got an error retrieving information about your Amazon EC2 instances:")
		return
	}



	newInstanceList, err = GetAliyunInstances(newInstanceList)
	if err != nil {
		fmt.Println("Got an error retrieving information about your Amazon EC2 instances:")
		return
	}


	var oldInstanceList []string
	var oldAsset []Hosts
	err = db.Where("deleted_at is  null").Unscoped().Find(&oldAsset).Error
	for _, v := range oldAsset {
		oldInstanceList = append(oldInstanceList, v.Instanceid)
	}

	fm := make(map[string]int)
	for i, v := range newInstanceList {
		fm[v] = i
	}
	for _, v := range oldInstanceList {
		if _, ok := fm[v]; ok {
		} else {
			asset := Hosts{
				Instanceid: v,
				DeletedAt:  sql.NullTime{time.Now(),true},
			}
			db.Model(&Hosts{}).Where("instanceid = ?", v).Updates(asset)
		}
	}

}

func main()  {
	AssetHostAwsUpdate()
}
