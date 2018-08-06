package lib

// источник данных
type DataSourcer interface {
	// сделать запрос и получить данные
	MakeRequest(Point, Point) (Data, error)
}
