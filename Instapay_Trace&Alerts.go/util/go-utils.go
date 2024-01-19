package util

import (
	"instapay/db"
	"instapay/model/bah"
)

func GetServiceEP(service, environment string) string {
	serviceRoute := &bah.ServiceRoute{}
	db.DB.Raw("SELECT * FROM rbi_instapay.get_service(?,?)", service, environment).Scan(serviceRoute)
	return serviceRoute.ServiceUrl
}
