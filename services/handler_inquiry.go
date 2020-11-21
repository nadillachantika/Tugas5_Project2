package services

import (
	"context"
	"database/sql"
	"fmt"

	cm "pnp/Framework/git/order/common"

	_ "github.com/go-sql-driver/mysql"
)

func (PaymentService) InquiryHandler(ctx context.Context, req cm.InquiryRequest) (res cm.InquiryResponse) {

	defer panicRecovery()

	host := cm.Config.Connection.Host
	port := cm.Config.Connection.Port
	user := cm.Config.Connection.User
	pass := cm.Config.Connection.Password
	data := cm.Config.Connection.Database

	var mySQL = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v", user, pass, host, port, data)

	db, err = sql.Open("mysql", mySQL)

	if err != nil {
		panic(err.Error())
	}

	var inquiryResponse cm.InquiryResponse
	var list cm.PaymentChannel

	sql := `SELECT
			TrxID,
			IFNULL(Merchant,'')Merchant,
			IFNULL(BillNo,'')BillNo,
			IFNULL(PaymentStatusCOde,'')PaymentStatusCOde,
			IFNULL(PaymentStatusDesc,'')PaymentStatusDesc

			FROM transaksi WHERE MerchantID = ?`
	result, err := db.Query(sql, req.MerchantID)

	defer result.Close()

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {
		err := result.Scan(&list.PgCode, &list.PgName, &list.BillNo, &list.PaymentStatusCode, &list.PaymentStatusDesc)

		if err != nil {
			panic(err.Error())
		}

		inquiryResponse.PaymentChannel = append(inquiryResponse.PaymentChannel, list)

	}

	// if inquiryResponse.MerchantID != "" {
	// 	res = inquiryResponse
	// 	res.Response = req.Request
	// 	res.ResponseCode = strconv.Itoa(http.StatusOK)
	// 	res.ResponseDesc = "Sukses ambil data"
	// } else {
	// 	res.ResponseCode = strconv.Itoa(http.StatusNotFound)
	// 	res.ResponseDesc = "Gagal ambil data"
	// }

	// return
	inquiryResponse.Merchant = req.Merchant
	inquiryResponse.MerchantID = req.MerchantID

	res = inquiryResponse

	return

}
