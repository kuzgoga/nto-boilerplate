package services

import "github.com/kuzgoga/nto-boilerplate/internal/dal"

type ReportService struct{}
type Report struct {
	RawProductAmount []string
	RawProductCount  int
}

func (report ReportService) GetResult(startDate int64, endDate int64) (result Report, err error) {
	rawProducts, err := dal.WorkResult.Preload(dal.WorkResult.Material).Where(dal.WorkResult.WorkDate.Gte(startDate), dal.WorkResult.WorkDate.Lte(endDate)).Distinct(dal.WorkResult.MaterialId).Find()
	rawProductsCount, _ := dal.WorkResult.Preload(dal.WorkResult.Material).Where(dal.WorkResult.WorkDate.Gte(startDate), dal.WorkResult.WorkDate.Lte(endDate)).Count()
	var rawProductsNames []string
	for _, rawProduct := range rawProducts {
		rawProductsNames = append(rawProductsNames, rawProduct.Material.Name)
	}
	if err != nil {
		return Report{}, err
	}
	return Report{
		RawProductAmount: rawProductsNames,
		RawProductCount:  int(rawProductsCount),
	}, nil
}
