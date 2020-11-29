package types


type OrderCollectionType string

const (
	OrderCollectionAny     OrderCollectionType = ""
	OrderCollectionActive  OrderCollectionType = "active"
	OrderCollectionArchive OrderCollectionType = "archive"
)