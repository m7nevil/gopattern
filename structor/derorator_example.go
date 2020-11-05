package structor

import "encoding/base64"

type DataSource interface {
	WriteData(data string)
	ReadData() string
}

// 数据源装饰器基类
type DataSourceDecorator struct {
	wrappee DataSource
}
func NewDataSourceDecorator(ds DataSource) *DataSourceDecorator {
	return &DataSourceDecorator{wrappee: ds}
}
func (bd *DataSourceDecorator) WriteData(data string) {
	bd.wrappee.WriteData(data)
}
func (bd *DataSourceDecorator) ReadData() string {
	return bd.wrappee.ReadData()
}


// 加密装饰器
type EncryptionDecorator struct {
	*DataSourceDecorator
}
func NewEncryptionDecorator(ds DataSource) *EncryptionDecorator{
	return &EncryptionDecorator{NewDataSourceDecorator(ds)}
}
// 数据据加密
func (cd *EncryptionDecorator) WriteData(data string) {
	encodeStr := base64.StdEncoding.EncodeToString([]byte(data))
	cd.DataSourceDecorator.WriteData(encodeStr)
}
//func (cd *EncryptionDecorator) ReadData() interface{}  {
//	return cd.ReadData()
//}

// 前缀装饰器
type PrefixDecorator struct {
	*DataSourceDecorator
}
func NewPrefixDecorator(ds DataSource) *PrefixDecorator {
	return &PrefixDecorator{NewDataSourceDecorator(ds)}
}
func (pd *PrefixDecorator) WriteData(data string)  {
	pd.DataSourceDecorator.WriteData("脑壳帝就这么叼：" + data)
}


// 数据源
type TextDataSource struct {
	Text string
}
func NewTextDataSource(text string) *TextDataSource {
	return &TextDataSource{Text: text}
}
func (tds *TextDataSource) WriteData(data string) {
	tds.Text = data
}
func (tds *TextDataSource) ReadData() string {
	return tds.Text
}

