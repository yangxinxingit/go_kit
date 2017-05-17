package logickit

import (
	"fmt"
	"github.com/astaxie/beego"
	"unicode/utf8"
	"strconv"
)

const (
	// 公司类型
	COMPANY_CATEGORY_PLATFORM = 100 // 平台
	COMPANY_CATEGORY_SERVICE = 200 // 服务商
	COMPANY_CATEGORY_OPERATOR = 300 // 运营商
)

// 校验公司号
func VerificationCompanyId(companyId int) bool {
	if !(10000 <= companyId && companyId <= 99999) {
		beego.Error("公司号错误", fmt.Sprintf("公司号：%v", companyId))
		return false
	}
	return true
}

// 校验公司类型
func VerificationCompCategory(companyCategory int) bool {
	if !(companyCategory == COMPANY_CATEGORY_PLATFORM ||
		companyCategory == COMPANY_CATEGORY_SERVICE ||
		companyCategory == COMPANY_CATEGORY_OPERATOR) {
		beego.Error("公司类型错误", fmt.Sprintf("公司类型=%v", companyCategory))
		return false
	}
	return true
}

// 校验支付方式
func VerificationPayWay(payWay int) bool {
	if !(0 <= payWay&& payWay <= 11) {
		beego.Error("支付方式错误", fmt.Sprintf("支付方式=%v", payWay))
		return false
	}
	return true
}

// 校验售货机ID
func VerificationSvmId(svmId int) bool {
	if !(23456 <= svmId) {
		beego.Error("售货机ID错误", fmt.Sprintf("售货机ID：%v", svmId))
		return false
	}
	return true
}

// 校验售货机销售状态
func VerificationSvmSaleStatus(saleStatus int) bool {
	if !( saleStatus == 10 || saleStatus == 11) {
		beego.Error("校验售货机销售状态错误", fmt.Sprintf("售货机销售状态：%v", saleStatus))
		return false
	}
	return true
}

// 校验售货机运营状态
func VerificationSvmWorkStatus(workStatus int) bool {
	if !( workStatus == 10 || workStatus == 11 || workStatus == 12) {
		beego.Error("校验售货机运营状态错误", fmt.Sprintf("售货机运营状态：%v", workStatus))
		return false
	}
	return true
}

// 校验售货机经（纬）度
func VerificationSvmLonOrLat(lonOrLat string) bool {
	if ( utf8.RuneCountInString(lonOrLat) > 15 ) {
		beego.Error("售货机经度或维度长度错误", fmt.Sprintf("售货机经度或维度：%v", lonOrLat))
		return false
	} else if _, err := strconv.ParseFloat(lonOrLat, 10); err != nil {
		beego.Error("售货机经度或维度格式错误", fmt.Sprintf("售货机经度或维度：%v", lonOrLat))
		return false
	}
	return true
}

// 校验公司权限类型
func VerificationCompAccessCategory(companyCategory int) bool {
	if !(companyCategory == 0 ||
		companyCategory == COMPANY_CATEGORY_PLATFORM ||
		companyCategory == COMPANY_CATEGORY_SERVICE ||
		companyCategory == COMPANY_CATEGORY_OPERATOR) {
		beego.Error("公司权限类型错误", fmt.Sprintf("公司权限类型=%v", companyCategory))
		return false
	}
	return true
}

// 校验权限ID
func VerificationAccessId(accessId int) bool {
	if accessId < 100000 || accessId > 999999 {
		beego.Error("权限ID错误", fmt.Sprintf("权限ID：%v", accessId))
		return false
	}
	return true
}

// 校验权限父ID
func VerificationAccessPId(accessPId int) bool {
	if (accessPId < 100000 && accessPId != 0) || accessPId > 999999 {
		beego.Error("权限父ID错误", fmt.Sprintf("权限父ID：%v", accessPId))
		return false
	}
	return true
}

// 校验登录密码
func VerificationLoginPwd(pwd string) bool {
	if len(pwd) != utf8.RuneCountInString(pwd) || len(pwd) < 8 {
		beego.Error("密码格式错误", fmt.Sprintf("密码：%v", pwd))
		return false
	}
	return true
}

// 校验登录账号
func VerificationLoginName(loginName string) bool {
	if len(loginName) != utf8.RuneCountInString(loginName) || len(loginName) < 6 {
		beego.Error("账号格式错误", fmt.Sprintf("账号：%v", loginName))
		return false
	}
	return true
}

// 校验手机号
func VerificationPhone(phone string) bool {
	if !(11 <= len(phone) && len(phone) <= 20) || len(phone) != utf8.RuneCountInString(phone) {
		beego.Error("手机号码格式错误", fmt.Sprintf("手机号码：%v", phone))
		return false
	}
	return true
}

// 校验是否
func VerificationIsNo(isOrNo string) bool {
	if !(isOrNo == "Y" || isOrNo == "N") {
		beego.Error("Y/N格式错误", fmt.Sprintf("Y/N：%v", isOrNo))
		return false
	}
	return true
}

// 校验页码
func VerificationPageNumber(number int) bool {
	if !(0 <= number && number <= 1000000) {
		beego.Error("页码错误", fmt.Sprintf("页码：%v", number))
		return false
	}
	return true
}

// 校验页面展示数量
func VerificationPageSize(size int) bool {
	if !(0 <= size && size <= 50) {
		beego.Error("页面展示数目错误", fmt.Sprintf("页面展示数目：%v", size))
		return false
	}
	return true
}

// 校验普通折扣方案类型
func VerificationDiscountType(discountType int) bool {
	if !(discountType == 10 || discountType == 11 || discountType == 12 || discountType == 13) {
		beego.Error("方案类型错误", fmt.Sprintf("类型：%v", discountType))
		return false
	}
	return true
}

// 校验方案状态
func VerificationPlanStatus(status int) bool {
	if !(status == 10 || status == 11 || status == 12 || status == 13) {
		beego.Error("方案状态错误", fmt.Sprintf("方案状态：%v", status))
		return false
	}
	return true
}

// 校验虚拟商品方案类型
func VerificationVgoodsType(vgoodsType int) bool {
	// '扫码虚拟商品：11，套餐虚拟商品：12，惊喜虚拟商品：13'
	if !( vgoodsType == 11 || vgoodsType == 12 || vgoodsType == 13) {
		beego.Error("虚拟商品方案类型错误", fmt.Sprintf("虚拟商品方案类型：%v", vgoodsType))
		return false
	}
	return true
}

// 校验url
func VerificationUrl(url string) bool {
	if utf8.RuneCountInString(url) > 150 {
		beego.Error("url长度大于150", fmt.Sprintf("url长度：%v", utf8.RuneCountInString(url)))
		return false
	}
	return true
}