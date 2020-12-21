package mysql

//func (db *DB) GetResult(column, table string, where map[string]interface{}) (ar *ActiveRecord) {
//	ar = new(ActiveRecord)
//	ar.Reset()
//	ar.tablePrefix = db.Config.TablePrefix
//	ar.tablePrefixSqlIdentifier = db.Config.TablePrefixSqlIdentifier
//
//	ar.Select(column).
//		From(table).
//		Where(where)
//
//	return
//}

//func (db *DB) GetResult(column, table string, where map[string]interface{}) (ar *ActiveRecord) {
//}

func (ar *ActiveRecord) getQueryResult() (rs *ResultSet, err error) {
	rs, err = ar.currentDBSession.Query(ar)

	return
}
func (ar *ActiveRecord) ExecQueryResult() (rs *ResultSet, err error) {
	rs, err = ar.currentDBSession.Exec(ar)

	return
}

func (ar *ActiveRecord) Find(table string, where map[string]interface{}) (rs *ResultSet, err error) {

	rs, err = ar.From(table).
		Wheres(where).
		getQueryResult()

	return
}

func (ar *ActiveRecord) FindAll(table string) (rs *ResultSet, err error) {

	rs, err = ar.From(table).
		getQueryResult()

	return
}

func (ar *ActiveRecord) FindOne(table string, column, value string) (rs *ResultSet, err error) {

	rs, err = ar.From(table).
		Where(column, value).
		Limit(1).
		getQueryResult()

	return
}

func (ar *ActiveRecord) HasRow(table string, column, value string) (bool, error) {

	rs, err := ar.Select(column).
		From(table).
		Where(column, value).
		Limit(1).
		getQueryResult()

	if rs.Len() > 0 {
		return true, err
	}

	return false, err
}

func (ar *ActiveRecord) Where(column, value string) *ActiveRecord {
	if len(column) > 0 {
		ar.WhereWrap(map[string]interface{}{
			column: value,
		}, "AND", "")
	}
	return ar
}

//func (ar *ActiveRecord) Column(name, value string) *ActiveRecord {
//	if len(name) > 0 {
//		ar.WhereWrap(where, "AND", "")
//	}
//	return ar
//}

func (ar *ActiveRecord) Insert(table string, data map[string]interface{}) (rs *ResultSet, err error) {
	ar.sqlType = "insert"
	ar.arInsert = data
	ar.From(table)

	rs, err = ar.ExecQueryResult()

	return
}
func (ar *ActiveRecord) Replace(table string, data map[string]interface{}) (rs *ResultSet, err error) {
	ar.sqlType = "replace"
	ar.arInsert = data
	ar.From(table)

	rs, err = ar.ExecQueryResult()
	return
}

func (ar *ActiveRecord) InsertBatch(table string, data []map[string]interface{}) (rs *ResultSet, err error) {
	ar.sqlType = "insertBatch"
	ar.arInsertBatch = data
	ar.From(table)

	rs, err = ar.ExecQueryResult()
	return
}
func (ar *ActiveRecord) ReplaceBatch(table string, data []map[string]interface{}) (rs *ResultSet, err error) {
	ar.InsertBatch(table, data)
	ar.sqlType = "replaceBatch"

	rs, err = ar.ExecQueryResult()
	return
}

func (ar *ActiveRecord) Delete(table string, where map[string]interface{}) (rs *ResultSet, err error) {
	ar.From(table)
	ar.Wheres(where)
	ar.sqlType = "delete"

	rs, err = ar.ExecQueryResult()
	return
}
func (ar *ActiveRecord) Update(table string, data, where map[string]interface{}) (rs *ResultSet, err error) {
	ar.From(table)
	ar.Wheres(where)
	for k, v := range data {
		if isBool(v) {
			value := 0
			if v.(bool) {
				value = 1
			}
			ar.Set(k, value)
		} else if v == nil {
			ar.SetNoWrap(k, "NULL")
		} else {
			ar.Set(k, v)
		}
	}

	rs, err = ar.ExecQueryResult()
	return
}

func (ar *ActiveRecord) UpdateBatch(table string, values []map[string]interface{}, whereColumn []string) (rs *ResultSet, err error) {
	ar.From(table)
	ar.sqlType = "updateBatch"
	ar.arUpdateBatch = []interface{}{values, whereColumn}
	if len(values) > 0 {
		for _, whereCol := range whereColumn {
			ids := []interface{}{}
			for _, val := range values {
				ids = append(ids, val[whereCol])
			}
			ar.Wheres(map[string]interface{}{whereCol: ids})
		}
	}

	rs, err = ar.ExecQueryResult()
	return
}
