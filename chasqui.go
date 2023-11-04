package chasqui

type CriteriaAdapter interface {
	ToMongo(criteria Criteria) interface{}
	ToES(criteria Criteria) string
	ToMySQL(criteria Criteria) string
	ToPostgreSQL(criteria Criteria) string
	ToSQLServer(criteria Criteria) string
	ToOracle(criteria Criteria) string
}

var CriteriaAdapterInstance CriteriaAdapter = &CriteriaAdapterImpl{}

type CriteriaAdapterImpl struct {
}

func (ca *CriteriaAdapterImpl) ToMongo(c Criteria) interface{} {
	return nil
}
func (ca *CriteriaAdapterImpl) ToES(criteria Criteria) string {

	return ""
}
func (ca *CriteriaAdapterImpl) ToMySQL(criteria Criteria) string {
	return ""
}
func (ca *CriteriaAdapterImpl) ToPostgreSQL(criteria Criteria) string {
	return ""
}
func (ca *CriteriaAdapterImpl) ToSQLServer(criteria Criteria) string {
	return ""
}
func (ca *CriteriaAdapterImpl) ToOracle(criteria Criteria) string {
	return ""
}
