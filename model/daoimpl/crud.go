package daoimpl

import "fmt"

func buildInsert(dataList map[string]interface{},tableName string) (string){
	cond := ""
	vals := ""
	for k,v := range dataList {
		cond = cond + k
		cond = cond + ","
		switch vv := v.(type) {
		case string:
			vals = vals + fmt.Sprintf("'%s'", vv)
		case int64:
			vals = vals + fmt.Sprintf("%d", vv)
		case bool:
			vals = vals + fmt.Sprintf("%v", vv)
		}
		vals = vals + ","
	}
	cond = cond[0 : len(cond)-1]
	vals = vals[0:len(vals)-1]
	cond = fmt.Sprintf("insert into %s(%s) values (%v)", tableName,cond,vals)

	return cond
}

func buildUpdate(data, conds map[string]interface{}, tableName string) string{
	cond := ""
	vals := ""
	for k,v := range data {
		cond = cond + k
		switch vv := v.(type) {
		case string:
			cond = cond + fmt.Sprintf(" = '%s'",vv)
		case int64:
			cond = cond + fmt.Sprintf(" = %d",vv)
		case bool:
			cond = cond + fmt.Sprintf(" = %v",vv)
		}
		cond = cond + ","
	}
	cond = cond[0 : len(cond)-1]

	for k,v := range conds {
		vals = vals + k
		switch vv := v.(type) {
		case string:
			vals = vals + fmt.Sprintf(" = '%s'",vv)
		case int64:
			vals = vals + fmt.Sprintf(" = %d",vv)
		case bool:
			vals = vals + fmt.Sprintf(" = %v",vv)
		}
		vals = vals + " and"
	}
	vals = vals[0 : len(cond)-4]

	cond = fmt.Sprintf("update %s set %s where %s",tableName,cond,vals)
	return cond
}

func buildDelete(data, conds map[string]interface{}, tableName string)string{
	cond := ""
	vals := ""
	for k,v := range data {
		cond = cond + k
		switch vv := v.(type) {
		case string:
			cond = cond + fmt.Sprintf(" = '%s'",vv)
		case int64:
			cond = cond + fmt.Sprintf(" = %d",vv)
		case bool:
			cond = cond + fmt.Sprintf(" = %v",vv)
		}
		cond = cond + ","
	}
	cond = cond[0 : len(cond)-1]

	for k,v := range conds {
		vals = vals + k
		switch vv := v.(type) {
		case string:
			vals = vals + fmt.Sprintf(" = '%s'",vv)
		case int64:
			vals = vals + fmt.Sprintf(" = %d",vv)
		case bool:
			vals = vals + fmt.Sprintf(" = %v",vv)
		}
		vals = vals + " and"
	}
	vals = vals[0 : len(cond)-4]

	cond = fmt.Sprintf("update %s set %s where %s",tableName,cond,vals)
	return cond
}

func buildList(conds map[string]interface{},tableName string)string{
	vals := ""

	for k,v := range conds {
		vals = vals + k
		switch vv := v.(type) {
		case string:
			vals = vals + fmt.Sprintf(" = '%s'",vv)
		case int64:
			vals = vals + fmt.Sprintf(" = %d",vv)
		case bool:
			vals = vals + fmt.Sprintf(" = %v",vv)
		}
		vals = vals + " and"
	}
	vals = vals[0 : len(vals)-4]

	vals = fmt.Sprintf("select * from %s where %s",tableName,vals)

	return vals
}