package service

import (
	"bytes"
	"github.com/longjoy/blog/dao/db"
	"github.com/longjoy/blog/model"
	"strconv"
	"strings"
)

func GetSkuListByProduct(pid int64)(sbp []*model.SkuByProductSpec, err error)  {
	productSku, err := db.SkuListByProduct(pid)
	saleType := db.GetSaleTypeById(pid)
	//初始化一个空切片，存放规格id
	spec := make([]string,0)
	for _, v := range productSku{
		spec = append(spec, strings.Split(v.SpecsIds,",")...)
	}

	//通过specid 查询specname
	specObjs, err := db.GetSpecsById(spec)
	if err != nil {
		return
	}

	for _, sku := range productSku{
		sp :=&model.SkuByProductSpec{}
		sp.SaleType = saleType
		sp.SpecsIds = sku.SpecsIds
		sp.CostPrice = sku.CostPrice
		sp.CreatedAt = sku.CreatedAt
		sp.CreatedTime = sku.CreatedAt.Format("2006-01-02 15:03:04")
		sp.OriginPrice = sku.OriginPrice
		sp.Price = sku.Price
		sp.Remain = sku.Remain
		sp.Skuid = sku.Skuid
		sp.Sold = sku.Sold
		sp.Stock = sku.Stock
		if sku.SpecsIds != "" {
			specIds := strings.Split(sku.SpecsIds,",")
			nameSlice := []string{}
			for _,specId := range specIds{
				specIdInt, err := strconv.ParseInt(specId, 10, 64)
				if err != nil {
					continue
				}

				var specName bytes.Buffer
				for _, specObj := range specObjs {
					if specIdInt == specObj.Id {
						specName.WriteString(specObj.SpecsName)
						specName.WriteString(":")
						specName.WriteString(specObj.SpecsValueName)
						name := specName.String()
						//log.Printf("name=",name)
						nameSlice = append(nameSlice, name)
						//log.Printf("nameSlice=",nameSlice)
					}
				}
				sp.Name = strings.Join(nameSlice, ",")
			}
		}else {
			sp.Name = ""
		}
		sbp = append(sbp, sp)
	}
	return
}
