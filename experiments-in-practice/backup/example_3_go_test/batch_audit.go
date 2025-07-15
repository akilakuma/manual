package main

/*

func userSet(event batch.EventData, userInfo batch.UserObject) (batch.PaySet, bool) {
	var audit_live, audit_ball, audit_complex, audit_normal string
	var paySet batch.PaySet
	currency := userInfo.Ret[0].Currency
	level, err := getUserLevel(event)
	if err {
		return paySet, err
	}
	id, err := api{}.getLevelCurrency(level, currency)
	if err {
		fmt.Printf("getLevelCurrency err: %+v", paySet)
		batchLog("ERROR - getLevelCurrency", 0, event, "")
		var d batch.PreserveData
		d.EventData = event
		p := batch.ErrInfo{
			Type:       "API",
			Err_Func:   "Redo_BatchDeposit",
			ReceiveEnd: "GET /api/level/{levelId}/currency",
		}
		preserveError(event, d, p)
		return paySet, err
	}
	paymentChargeId := strconv.Itoa(id)
	if paymentChargeId == "0" {
		paymentChargeId, err = api{}.getDomainPaySet(event.HallID, currency)
		if err {
			fmt.Printf("getDomainPaySet err: %+v", paySet)
			batchLog("ERROR - getDomainPaySet", 0, event, "")
			var d batch.PreserveData
			d.EventData = event
			p := batch.ErrInfo{
				Type:       "API",
				Err_Func:   "Redo_BatchDeposit",
				ReceiveEnd: "GET /api/domain/{domain}/payment_charge",
			}
			preserveError(event, d, p)
			return paySet, err
		}
	}
	deposit_company, err := api{}.getDepositCompany(paymentChargeId)
	if err {
		fmt.Printf("getDepositCompany err: %+v", paySet)
		batchLog("ERROR - getDepositCompany", 0, event, "")
		var d batch.PreserveData
		d.EventData = event
		p := batch.ErrInfo{
			Type:       "API",
			Err_Func:   "Redo_BatchDeposit",
			ReceiveEnd: "GET /api/payment_charge/{paymentChargeId}/deposit_company",
		}
		preserveError(event, d, p)
		return paySet, err
	}
	ret := deposit_company["ret"].(map[string]interface{})
	if ret["audit_live"] == true {
		audit_live = "Y"
	} else {
		audit_live = "N"
	}
	if ret["audit_ball"] == true {
		audit_ball = "Y"
	} else {
		audit_ball = "N"
	}
	if ret["audit_complex"] == true {
		audit_complex = "Y"
	} else {
		audit_complex = "N"
	}
	if ret["audit_normal"] == true {
		audit_normal = "Y"
	} else {
		audit_normal = "N"
	}
	paySet.SpType_Co = ret["discount"].(float64)
	paySet.LIVE_OPEN_Co = audit_live
	paySet.Audit_LIVE_Co, _ = strconv.ParseFloat(ret["audit_live_amount"].(string), 64)
	paySet.BALL_OPEN_Co = audit_ball
	paySet.Audit_BALL_Co, _ = strconv.ParseFloat(ret["audit_ball_amount"].(string), 64)
	paySet.COMPLEX_OPEN_Co = audit_complex
	paySet.Audit_COMPLEX_Co, _ = strconv.ParseFloat(ret["audit_complex_amount"].(string), 64)
	paySet.SURPLUS_RATE_Co, _ = strconv.ParseFloat(ret["audit_discount_amount"].(string), 64)
	paySet.NORMALITY_OPEN_Co = audit_normal
	paySet.Audit_NORMALITY_Co, _ = strconv.ParseFloat(ret["audit_normal_amount"].(string), 64)
	paySet.SpRate_Co, _ = strconv.ParseFloat(ret["discount_percent"].(string), 64)
	paySet.TOLERANCE_Co, _ = strconv.ParseFloat(ret["audit_loosen"].(string), 64)
	paySet.CHARGE_RATE_Co, _ = strconv.ParseFloat(ret["audit_administrative"].(string), 64)
	paySet.Daily_Discount_Limit, _ = strconv.ParseFloat(ret["daily_discount_limit"].(string), 64)
	return paySet, false
}

*/
