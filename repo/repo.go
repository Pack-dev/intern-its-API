package repo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"templategoapi/db"

	"github.com/fatih/color"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func databaseConnect() (*db.Resource, error) {
	resource, err := db.CreateResource()
	if err != nil {
		color.Red("Connection database failure, Please check connection")
		color.Cyan(err.Error())
		logrus.Error(err)
		lineNotifyAlert(err)
		return resource, err
	}
	return resource, nil
}

func CreateStatement(resource *db.Resource, collection string, payload interface{}) (*mongo.InsertOneResult, error) {
	// resource, err := databaseConnect()
	// defer resource.Close()
	// if err != nil {
	// return nil, err
	// }
	ctx, cancel := db.InitContext()
	defer cancel()
	obj, err := resource.DB.Collection(collection).InsertOne(ctx, payload)
	if err != nil {
		return obj, err
	}
	// if obj == nil {
	// 	// err := errors.New("INSERT FAIL")
	// 	return obj, err
	// }
	return obj, nil
}

func UpdateOneStatement(resource *db.Resource, collection string, filter interface{}, set interface{}) (*mongo.UpdateResult, error) {
	// resource, err := databaseConnect()
	// defer resource.Close()
	// if err != nil {
	// return err
	// }
	ctx, cancel := db.InitContext()
	defer cancel()
	obj, err := resource.DB.Collection(collection).UpdateOne(ctx, filter, set)
	if err != nil {
		err := errors.New("Update fail, syntax error")
		return obj, err
	}
	if obj == nil {
		err := errors.New("Update fail, syntax error")
		return obj, err
	}
	// if obj.MatchedCount == 0 {
	// 	err := errors.New("Data mismatched")
	// 	return err
	// }
	// if obj.ModifiedCount == 0 {
	// 	err := errors.New("Matched but Nothing changed")
	// 	return err
	// }
	return obj, nil
}

func UpdateManyStatement(resource *db.Resource, collection string, filter interface{}, set interface{}) error {
	// resource, err := databaseConnect()
	// defer resource.Close()
	// if err != nil {
	// return err
	// }
	ctx, cancel := db.InitContext()
	defer cancel()
	obj, err := resource.DB.Collection(collection).UpdateMany(ctx, filter, set)
	if err != nil {
		err := errors.New("Update fail, syntax error")
		return err
	}
	if obj == nil {
		err := errors.New("Update fail, syntax error")
		return err
	}
	// if obj.MatchedCount == 0 {
	// 	err := errors.New("Data mismatched")
	// 	return err
	// }
	// if obj.ModifiedCount == 0 {
	// 	err := errors.New("Matched but Nothing changed")
	// 	return err
	// }
	return nil
}

func GetOneStatement(resource *db.Resource, collection string, filter interface{}, filterOption interface{}, data interface{}) error {
	// resource, err := databaseConnect()
	// defer resource.Close()
	// if err != nil {
	// return err
	// }
	ctx, cancel := db.InitContext()
	option := options.FindOne()
	option.SetSort(filterOption)
	defer cancel()
	err := resource.DB.Collection(collection).FindOne(ctx, filter, option).Decode(data)
	if err != nil {
		return err
	}
	return nil
}

func GetManyStatement(resource *db.Resource, collection string, filter interface{}, filterOption interface{}, data interface{}) error {
	// resource, err := databaseConnect()
	// defer resource.Close()
	// if err != nil {
	// return err
	// }
	ctx, cancel := db.InitContext()
	option := options.Find()
	option.SetSort(filterOption)
	defer cancel()
	obj, err := resource.DB.Collection(collection).Find(ctx, filter, option)
	if err != nil {
		return err
	}
	err = obj.All(ctx, data)
	if err != nil {
		return err
	}
	return nil
}
func GetManyStatementLimit(resource *db.Resource, collection string, filter interface{}, filterOption interface{}, limit int64, data interface{}) error {

	ctx, cancel := db.InitContext()
	option := options.Find()
	option.SetSort(filterOption)
	option.SetLimit(limit)
	defer cancel()
	obj, err := resource.DB.Collection(collection).Find(ctx, filter, option)
	if err != nil {
		return err
	}
	err = obj.All(ctx, data)
	if err != nil {
		return err
	}
	return nil
}
func GetManyStatementLimitOne(resource *db.Resource, collection string, filter interface{}, filterOption interface{}, data interface{}) error {
	// resource, err := databaseConnect()
	// defer resource.Close()
	// if err != nil {
	// return err
	// }
	ctx, cancel := db.InitContext()
	option := options.Find()
	option.SetSort(filterOption)
	option.SetLimit(1)
	defer cancel()
	obj, err := resource.DB.Collection(collection).Find(ctx, filter, option)
	if err != nil {
		return err
	}
	err = obj.All(ctx, data)
	if err != nil {
		return err
	}
	return nil
}
func CountStatement(resource *db.Resource, collection string, filter interface{}, data *int64) error {
	// resource, err := databaseConnect()
	// defer resource.Close()
	// if err != nil {
	// return err
	// }
	ctx, cancel := db.InitContext()
	defer cancel()
	obj, err := resource.DB.Collection(collection).CountDocuments(ctx, filter)
	if err != nil {
		return err
	}
	*data = obj
	return nil
}
func AggregateStatement(resource *db.Resource, collection string, filter interface{}, data interface{}) error {
	// resource, err := databaseConnect()
	// defer resource.Close()
	// if err != nil {
	// return err
	// }
	ctx, cancel := db.InitContext()
	defer cancel()
	obj, err := resource.DB.Collection(collection).Aggregate(ctx, filter)
	if err != nil {
		return err
	}
	err = obj.All(ctx, data)
	if err != nil {
		return err
	}
	return nil
}

func DeleteOneStatement(resource *db.Resource, collection string, filter interface{}) (*mongo.DeleteResult, error) {
	ctx, cancel := db.InitContext()
	defer cancel()
	obj, err := resource.DB.Collection(collection).DeleteOne(ctx, filter)
	if err != nil {
		return obj, err
	}
	return obj, nil
}

func CreateManyStatement(resource *db.Resource, collection string, payload []interface{}) (*mongo.InsertManyResult, error) {

	ctx, cancel := db.InitContext()
	defer cancel()
	obj, err := resource.DB.Collection(collection).InsertMany(ctx, payload)
	if err != nil {
		return obj, err
	}

	if obj == nil {
		err := errors.New("INSERT FAIL")
		return obj, err
	}
	return obj, nil
}

func lineNotifyAlert(msg error) error {
	type reslinenoti struct {
		Status  int    `json:"status"`
		Message string `json:"message"`
	}

	var response reslinenoti
	accesstoken := "pKahNdB8I8ifXmaeekmD66EbvXSgKMF5hBYUt5bgiNa"
	payload := fmt.Sprintf(
		"message= \n  API-SERVICE %s Error Connect DB \n  Error : %s \n", os.Getenv("SERVICE"), msg)

	URL := "https://notify-api.line.me/api/notify"
	Header := map[string][]string{
		"Authorization": {"Bearer " + accesstoken},
		"Content-Type":  {"application/x-www-form-urlencoded"},
	}
	if err := externalCALL(URL, "POST", Header, payload, &response); err != nil {
		return nil
	}
	return nil
}

func externalCALL(URL string, method string, headers map[string][]string, bodyPayload string, obj interface{}) error {
	payload := strings.NewReader(bodyPayload)
	client := &http.Client{}
	req, err := http.NewRequest(method, URL, payload)
	if err != nil {
		return err
	}
	req.Header = headers
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	json.Unmarshal(body, obj)
	if res.StatusCode != 200 {
		return errors.New("CANNOT CALL API")
	}
	if err != nil {
		return err
	}
	if err := json.Unmarshal(body, obj); err != nil {
		return err
	}
	return nil
}
func GetManyStatementLimitPagination(resource *db.Resource, collection string, filter interface{}, filterOption interface{}, limit int64, page int64, data interface{}) (totalPages int64, count int64, err error) {
	ctx, cancel := db.InitContext()
	defer cancel()
	count, err = resource.DB.Collection(collection).CountDocuments(ctx, filter)
	if err != nil {
		return 0, 0, err
	}
	option := options.Find()
	option.SetSort(filterOption)
	option.SetLimit(limit)
	skip := (page - 1) * limit
	if skip < 0 {
		skip = 0
	}
	option.SetSkip(skip)
	obj, err := resource.DB.Collection(collection).Find(ctx, filter, option)
	if err != nil {
		return 0, 0, err
	}
	err = obj.All(ctx, data)
	if err != nil {
		return 0, 0, err
	}
	totalPages = count / limit
	if count%limit != 0 {
		totalPages++
	}
	return totalPages, count, nil
}
