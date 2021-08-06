package db

import (
	"github.com/jmoiron/sqlx"
	"github.com/longjoy/blog/model"
	"log"
)

func GetSpecsById(specIds []string) (spec []*model.Spec, err error)  {
	sqlStr, args, err := sqlx.In("select id, specs_name, specs_value_name from shop_specs where  id in(?) ", specIds)
	if err != nil {
		log.Printf("spec_id_in_err=", err)
		return
	}
	err = DBWeb.Select(&spec, sqlStr, args...)
	if err != nil {
		log.Printf("spec_id_err=", err)
		return
	}
	return
}
