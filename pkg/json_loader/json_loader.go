package json_loader

//TODO: Tightly coupled to product entity, refactor this later
import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/gmalheiro/playground-golang-clean-arch/internal/domain/entity"
)

type StorageObj struct {
	Products map[int]*entity.Product
	Set      map[string]bool
}

type tempProduct struct {
	ID          int
	Name        string
	Quantity    int
	CodeValue   string
	IsPublished bool
	Expiration  string
	Price       float64
}

func jsonToInterface(path string) ([]tempProduct, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var products []tempProduct
	if err = json.Unmarshal(file, &products); err != nil {
		return nil, err
	}
	return products, nil
}

func stringToDate(productExpiration string) time.Time {
	layout := "02/01/2006"
	date, err := time.Parse(layout, productExpiration)
	if err != nil {
		log.Fatal(err.Error())
	}
	return date
}

func Read(path string) (StorageObj, error) {
	tempProducts, err := jsonToInterface(path)
	if err != nil {
		log.Fatal(err.Error())
	}
	var products = make(map[int]*entity.Product)
	var set = make(map[string]bool)
	for _, tempProduct := range tempProducts {
		date := stringToDate(tempProduct.Expiration)
		product := entity.Product{
			ID:          tempProduct.ID,
			Name:        tempProduct.Name,
			Quantity:    tempProduct.Quantity,
			CodeValue:   tempProduct.CodeValue,
			IsPublished: tempProduct.IsPublished,
			Expiration:  date,
			Price:       tempProduct.Price,
		}
		products[tempProduct.ID] = &product
		set[tempProduct.CodeValue] = true
	}
	stObj := StorageObj{Products: products, Set: set}
	return stObj, nil
}
