package mgquery

import (
	"affiliate/internal/constants"
	"affiliate/internal/format"
	"affiliate/internal/util"
	"fmt"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/thoas/go-funk"
	"go.mongodb.org/mongo-driver/bson"
)

type CommonQuery struct {
	Page                     int64
	Limit                    int64
	SortInterface            interface{}
	Keyword                  string
	SortStr                  string
	PartnerID                string
	affiliateID              string
	Status                   string
	Active                   string
	FromAt                   time.Time
	ToAt                     time.Time
	Source                   string
	Type                     string
	StatusOrder              string
	FromCreatedAt            string
	ToCreatedAt              string
	TypePackage              string
	Code                     string
	Arrange                  string
	OrderFromAt              string
	OrderToAt                string
	CashRevenueFromDate      string
	CashRevenueToDate        string
	TotalCashRevenueFromDate string
	TotalCashRevenueToDate   string
	TotalOrdersFromDate      string
	TotalOrdersToDate        string
	OrderStaticsFromDate     string
	OrderStaticsToDate       string
	PakageMonthlyType        string
	ParentNameId             string
	FromPrice                int
	ToPrice                  int
}

// AssignFromPriceToPrice ...
func (q *CommonQuery) AssignFromPriceToPrice(cond *bson.M) {
	priceCond := bson.M{}

	if q.FromPrice > 0 {
		priceCond["$gte"] = int32(q.FromPrice)
	}

	if q.ToPrice > 0 {
		priceCond["$lte"] = int32(q.ToPrice)
	}

	if len(priceCond) > 0 {
		(*cond)["price"] = priceCond
	}
}

// AssignFromToCreatedAt ...
func (q *CommonQuery) AssignFromToCreatedAt(cond *bson.M) {

	//2023-12-04T17:00:00.000Z
	fromUpdatedAt := util.TimeParseISODate(q.FromCreatedAt)
	toUpdatedAt := util.TimeParseISODate(q.ToCreatedAt)
	if toUpdatedAt.Unix() > fromUpdatedAt.Unix() {
		fmt.Println("fromUpdatedAt", fromUpdatedAt, "toUpdatedAt", toUpdatedAt)
		(*cond)["createdAt"] = bson.M{
			"$gte": fromUpdatedAt,
			"$lte": toUpdatedAt,
		}
	}
}

// AssignPackageMonthlyType ...
func (q *CommonQuery) AssignPackageMonthlyType(cond *bson.M) {
	if q.PakageMonthlyType != "" {
		(*cond)["packageMonthlyType"] = q.PakageMonthlyType
	}
}

// AssiTypePackage
func (q *CommonQuery) AssignTypePackage(cond *bson.M) {
	if q.TypePackage != "" {
		(*cond)["type"] = q.TypePackage
	}
}

// AssignStatus ...
func (q *CommonQuery) AssignStatus(cond *bson.M) {
	if q.Status != "" {
		(*cond)["status"] = q.Status
	}
}

// AssignKeyValueString ...
func (q CommonQuery) AssignKeyValueString(cond *bson.M, key, value string) {
	if !funk.Contains([]string{"", "all"}, value) {
		(*cond)[key] = value
	}
}

// AssignKeyValueObjectID ...
func (q CommonQuery) AssignKeyValueObjectID(cond *bson.M, key, value string) {
	if value != "" && !util.ConvertStringToObjectID(value).IsZero() {
		(*cond)[key] = util.ConvertStringToObjectID(value)
	}
}

// AssignKeyword ...
func (q *CommonQuery) AssignKeyword(cond *bson.M) {
	if q.Keyword != "" {
		q.Keyword = format.NonAccentVietnamese(q.Keyword)
		(*cond)["searchString"] = format.SearchString(strings.Trim(q.Keyword, " "))
	}
}

// AssignStatusOrder ...
func (q *CommonQuery) AssignStatusOrder(cond *bson.M) {
	if q.StatusOrder != "" {
		(*cond)["status"] = q.StatusOrder
	}
}

// AssignPartnerID ...
func (q *CommonQuery) AssignPartnerID(cond *bson.M) {
	if q.PartnerID != "" {
		q.AssignKeyValueObjectID(cond, "partnerID", q.PartnerID)
	}
}

// AssignaffiliateID ...
func (q *CommonQuery) AssignaffiliateID(cond *bson.M) {
	if q.affiliateID != "" {
		(*cond)["affiliateID"] = q.affiliateID
	}
}

// AssignType
func (q *CommonQuery) AssignType(cond *bson.M) {
	if q.Type != "" {
		(*cond)["type"] = q.Type
	}
}

// AssignSortStr ...
func (q *CommonQuery) AssignSortStr() {
	if q.SortStr != "" {
		switch q.SortStr {
		case constants.SortNewest:
			q.SortInterface = bson.D{
				{"createdAt", -1},
			}
		case constants.SortOldest:
			q.SortInterface = bson.D{
				{"createdAt", 1},
			}
		}
	}
}

// AssignParentNameId ...
func (q *CommonQuery) AssignParentNameId(cond *bson.M) {
	if q.ParentNameId != "" {
		(*cond)["parentNameId"] = q.ParentNameId
	}
}

// AssignSortStr ...
func (q *CommonQuery) GetSortInt() int64 {
	if q.SortStr != "" {
		switch q.SortStr {
		case constants.SortNewest:
			return -1
		case constants.SortOldest:
			return 1
		}
	}

	return -1
}

// AssignActive ...
func (q *CommonQuery) AssignActive(cond *bson.M) {
	if q.Active != "" {
		if q.Active == "true" {
			(*cond)["active"] = true
		}
		if q.Active == "false" {
			(*cond)["active"] = false
		}
	}
}

// GetFindOptsUsingPage ...
func (q *CommonQuery) GetFindOptsUsingPage() *options.FindOptions {
	opts := options.Find()
	if q.Limit > 0 {
		opts.SetLimit(q.Limit).SetSkip(q.Page * q.Limit)
	}

	if q.SortInterface != nil {
		opts.SetSort(q.SortInterface)
	}
	return opts
}

// GetFindOptsUsingTimestamp ...
func (q *CommonQuery) GetFindOptsUsingTimestamp() *options.FindOptions {
	opts := options.Find()
	if q.Limit > 0 {
		opts.SetLimit(q.Limit)
	}
	if q.SortInterface != nil {
		opts.SetSort(q.SortInterface)
	}
	return opts
}

// GetFindOptionsUsingSort ...
func (q *CommonQuery) GetFindOptionsUsingSort() *options.FindOptions {
	return options.Find().SetSort(q.SortInterface)
}
