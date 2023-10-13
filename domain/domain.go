package domain

type ModelDB struct {
	Id          int    `json:"id"`
	ModelDBId   string `json:"model_db_Id"`
	ModelDBName string `json:"model_db_name"`
	Libs        string `json:"libs"`
	Describe    string `json:"describe"`
	Version     int    `json:"version"`
	Status      string `json:"status"`
}

type ModelSchema struct {
	Id               int    `json:"id"`
	ModelDBId        string `json:"model_db_id"`
	SchemaID         string `json:"schema_id"`
	SchemaName       string `json:"schema_name"`
	SchemaEntityName string `json:"schema_entity_name"`
	SchemaType       string `json:"schema_type"`
	StorageType      string `json:"storage_type"`
	Describe         string `json:"describe"`
}

type ModelCollection struct {
	Id                   int    `json:"id"`
	SchemaId             string `json:"schema_id"`
	CollectionId         string `json:"collection_id"`
	CollectionName       string `json:"collection_name"`
	CollectionEntityName string `json:"collection_entity_name"`
	Rules                string `json:"rules"`
	ConfigInfo           string `json:"config_info"`
}

type CollectionField struct {
	Id              int    `json:"id"`
	CollectionId    string `json:"collection_id"`
	FieldId         string `json:"field_id"`
	FieldName       string `json:"field_name"`
	FieldEntityName string `json:"field_entity_name"`
	ConfigInfo      string `json:"config_info"`
}

type ModelDBDescribe struct {
	ModelDB    ModelDB          `json:"model_db"`
	Schemas    []SchemaDescribe `json:"schemas"`
	ConnectStr string           `json:"connect_str"`
}

type SchemaDescribe struct {
	Schema      ModelSchema          `json:"schema"`
	Collections []CollectionDescribe `json:"collections"`
}

type CollectionDescribe struct {
	Collection ModelCollection   `json:"collection"`
	Fields     []CollectionField `json:"fields"`
}

type CreateResp struct {
	ErrCode    int             `json:"err_code"`
	ErrMessage string          `json:"err_message"`
	Result     ModelDBDescribe `json:"result"`
}

type EmitEventReq struct {
	ApplicationId string `json:"application_id"`
	InstanceId    string `json:"instance_id"`
	ModelDBId     string `json:"model_db_id"`
	EventName     string `json:"event_name"`
	Source        string `json:"source"` // 事件源，一般为触发此事件的函数名
	Params        any    `json:"params"` // 触发事件的参数，json字符串
}

type ModelVersionDiff struct {
	Id          string `json:"id"`
	ModelDBId   string `json:"model_db_id"`
	DiffType    string `json:"diff_type"`
	Type        string `json:"type"`
	DiffContent string `json:"diff_content"`
	Version     int    `json:"version"`
	Remark      string `json:"remark"`
}
