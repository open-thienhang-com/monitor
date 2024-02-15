package product

import (
	"encoding/json"
	"fmt"
	"reflect"

	"mono.thienhang.com/pkg/database"
	"mono.thienhang.com/pkg/plugins/api"
)

type Model struct {
	api.Model   `json:"-"`
	ID          int64   `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Unit        string  `json:"unit"`
	Category    string  `json:"category"`
	Brand       string  `json:"brand"`
	Sale_status string  `json:"sale_status"`
	Sales_type  string  `json:"sales_type"`
}

func (p *Model) ToByte() ([]byte, error) {
	bytes, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}

func (t Model) SetConn(con database.Connection) Model {
	t.Conn = con
	return t
}

// Find return a default user model of given id.
func (t Model) Find(id interface{}) *Model {
	item, _ := t.Table(t.TableName).Find(id)
	x, _ := mapToProduct(item)
	fmt.Println(item, x)
	return &x
}

func mapToProduct(data map[string]interface{}) (Model, error) {
	result := Model{}

	// Get the type of the model
	productType := reflect.TypeOf(result)

	// Iterate over the map and set field values
	for key, value := range data {
		fmt.Println(key)
		// Find the corresponding field in the struct
		field, found := productType.FieldByName(key)
		fmt.Println(field)
		if !found {
			return result, fmt.Errorf("Field %s not found in the model", key)
		}

		// Check if the field type is compatible with the map value type
		if field.Type.Kind() != reflect.TypeOf(value).Kind() {
			return result, fmt.Errorf("Type mismatch for field %s", key)
		}

		// Set the field value in the result struct
		reflect.ValueOf(&result).Elem().FieldByName(key).Set(reflect.ValueOf(value))
	}

	return result, nil
}

// IsEmpty check the model is empty or not.
func (t Model) IsEmpty() bool {
	return t.ID == int64(0)
}

// FindByUserName return a default user model of given name.
func (t Model) FindByUserName(username interface{}) Model {
	item, _ := t.Table(t.TableName).Where("username", "=", username).First()
	x, _ := mapToProduct(item)
	fmt.Println(item, x)
	return x
}

// Update update the user model.
func (t Model) Update(username, password, name, avatar string, isUpdateAvatar bool) (int64, error) {

	// fieldValues := dialect.H{
	// 	"username":   username,
	// 	"name":       name,
	// 	"updated_at": time.Now().Format("2006-01-02 15:04:05"),
	// }

	// if avatar == "" || isUpdateAvatar {
	// 	fieldValues["avatar"] = avatar
	// }

	// if password != "" {
	// 	fieldValues["password"] = password
	// }

	// // return t.WithTx(t.Tx).Table(t.TableName).
	// // 	Where("id", "=", t.Id).
	// // 	Update(fieldValues)
	return 0, nil
}
