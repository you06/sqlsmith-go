package sqlsmith


// DataGenerator defines data generator
type DataGenerator struct {
	total int
	batch int
	curr int
	sqlsmith *SQLSmith
}

// GenData returns data generator
func (s *SQLSmith) GenData(total, batch int) *DataGenerator {
	return &DataGenerator{
		total,
		batch,
		0,
		s,
	}
}

// Next returns data batch
func (d *DataGenerator) Next() []string {
	if d.curr >= d.total {
		return []string{}
	} else if d.curr < d.total - d.batch {
		sqls, _ := d.sqlsmith.BatchData(d.batch, d.batch)
		d.curr += d.batch
		return sqls
	}
	sqls, _ := d.sqlsmith.BatchData(d.total - d.curr, d.batch)
	d.curr += (d.total - d.curr)
	return sqls
}
