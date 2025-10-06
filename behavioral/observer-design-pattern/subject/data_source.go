package subject

type DataSource struct {
	Subject
	values []int
}

func NewDataSource() *DataSource {
	return &DataSource{}
}

func (d *DataSource) GetValues() []int {
	return d.values
}

func (d *DataSource) SetValues(v []int) {
	d.values = v
	d.NotifyObserver()
}
