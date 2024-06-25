package pmg

import (
	"encoding/json"
	"sort"

	"github.com/thoas/go-funk"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetAppIDFromHex ...
func GetAppIDFromHex(s string) primitive.ObjectID {
	id, _ := primitive.ObjectIDFromHex(s)
	return id
}

// ConvertObjectIDsToStrings ...
func ConvertObjectIDsToStrings(ids []primitive.ObjectID) []string {
	return funk.Map(ids, func(item primitive.ObjectID) string {
		return item.Hex()
	}).([]string)
}

// UniqObjectIds ...
func UniqObjectIds(ids []primitive.ObjectID) []primitive.ObjectID {
	idStr := ConvertObjectIDsToStrings(ids)
	idStr = funk.UniqString(idStr)
	return ConvertStringsToObjectIDs(idStr)
}

// ConvertStringsToObjectIDs ...
func ConvertStringsToObjectIDs(strValues []string) []primitive.ObjectID {
	return funk.Map(strValues, func(item string) primitive.ObjectID {
		return GetAppIDFromHex(item)
	}).([]primitive.ObjectID)
}

// ConvertStringToObjectID ...
func ConvertStringToObjectID(id string) primitive.ObjectID {
	ID, _ := primitive.ObjectIDFromHex(id)
	return ID
}

// SortObjectId ...
func SortObjectId(ids []primitive.ObjectID) []primitive.ObjectID {
	sort.Slice(ids, func(i, j int) bool {
		return ids[i].Timestamp().Before(ids[j].Timestamp())
	})
	return ids
}

// ConvertInterfacesToObjectIDs ...
func ConvertInterfacesToObjectIDs(values []interface{}) []primitive.ObjectID {
	s := make([]primitive.ObjectID, len(values))
	for i, v := range values {
		s[i] = v.(primitive.ObjectID)
	}
	return s
}

// ConvertInterfacesToStrings ...
func ConvertInterfacesToStrings(values []interface{}) []string {
	s := make([]string, len(values))
	for i, v := range values {
		s[i] = v.(string)
	}
	return s
}

// ConvertStringsToInterfaces ...
func ConvertStringsToInterfaces(values []string) []interface{} {
	s := make([]interface{}, len(values))
	for _, v := range values {
		s = append(s, v)
	}
	return s
}

// ConvertToJSONString ...
func ConvertToJSONString(data interface{}) string {
	dataByte, _ := json.Marshal(data)
	return string(dataByte)
}
