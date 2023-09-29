package function

import (
	"context"
	"encoding/json"
	"fmt"
	characters "github.com/nanikjava/comicstype/json/characters"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"project/github/comics/dbhandler/models"
	"project/github/comics/dbhandler/mongo"
	"project/github/comics/utils"
	"reflect"
)

func CharactersFunction(m *mongo.Mongo, coltype string, body []byte) error {
	//convert to JSON object first
	var result = characters.Results{}

	if err := json.Unmarshal(body, &result); err != nil {
		log.Println("Error converting json :", err)
		return err
	}

	//func (m *Mongo) Read(ctx context.Context, col string, req *models.ReadRequest) (int64, interface{}, map[string]map[string]string, *models.SQLMetaData, error) {
	_, intface, _, _, err := m.Read(context.Background(),
		coltype,
		&models.ReadRequest{
			Find: map[string]interface{}{
				"id": result.ID,
			},
			Operation: utils.One,
			Options:   nil,
		},
	)

	// convert into struct
	if intface != nil {
		jsonData, err := json.Marshal(intface)
		if err != nil {
			fmt.Println("Error:", err)
		}

		var convertToResult = characters.Results{}
		if err := json.Unmarshal(jsonData, &convertToResult); err != nil {
			fmt.Println("Error:", err)
		}

		isEqual := reflect.DeepEqual(convertToResult, result)

		if isEqual {
			fmt.Println("The structs are equal.")
		} else {
			fmt.Println("The structs are not equal.")
		}
	}
	//
	//if err != nil {
	//	log.Fatal(err)
	//}

	var bsonDocument bson.M
	err = bson.UnmarshalExtJSON(body, false, &bsonDocument)
	if err != nil {
		return err
	}

	_, err = m.Create(context.Background(), coltype, &models.CreateRequest{
		Document:  bsonDocument,
		Operation: utils.One,
		IsBatch:   false,
	})

	return err
}
