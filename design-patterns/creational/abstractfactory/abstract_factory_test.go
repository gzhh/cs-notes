package abstractfactory

import "testing"

func handleMainAndDetail(factory DAOFactory) {
	factory.CreateOrderMainDAO().SaveOrderMain()
	factory.CreateOrderDetailDAO().SaveOrderDetail()
}

func TestRDBFactory(t *testing.T) {
	var factory DAOFactory
	factory = &RDBDAOFactory{}
	handleMainAndDetail(factory)
}

func TestXMLFactory(t *testing.T) {
	var factory DAOFactory
	factory = &XMLDAOFactory{}
	handleMainAndDetail(factory)
}
