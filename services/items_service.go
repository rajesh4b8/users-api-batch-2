// Dummy service
package services

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsService struct {
}

type itemsServiceInterface interface {
	GetItem()
	CreateItem()
	DeleteItem()
}

func (s *itemsService) GetItem() {

}

func (s *itemsService) CreateItem() {

}

func (s *itemsService) DeleteItem() {

}
