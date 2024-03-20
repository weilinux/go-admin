package model

// Optimi: 这里还应该有一个可以进行数据排序的参数和处理
func Paginate(entity Entity, page, limit int) (interface{}, int64) {
	offset := (page - 1) * limit

	data := entity.Take(db, limit, offset)
	total := entity.Count(db)

	return data, total
}
