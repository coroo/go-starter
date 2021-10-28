package usecases

import (
	"fmt"
	"os"
	"strconv"
	"crypto/sha1"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"net/http"
	entity "github.com/coroo/go-starter/app/entity"
	utils "github.com/coroo/go-starter/app/utils"
	repositories "github.com/coroo/go-starter/app/repositories"
	// "github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type PaymentMethodService interface {
	SavePaymentMethod(entity.PaymentMethod) (int, error)
	UpdatePaymentMethod(entity.PaymentMethod) error
	DeletePaymentMethod(entity.PaymentMethod) error
	GetAllPaymentMethods(total_premium string, status string) []entity.PaymentMethodWithPremium
	GetPaymentMethod(id string) []entity.PaymentMethod
	GetPaymentMethodByCode(code string) entity.PaymentMethod
	GenerateVaSignature(code string, proposal_group_number string) string
	ConnectionTest() *http.Response
}

type paymentMethodService struct {
	repositories repositories.PaymentMethodRepository
}

type fastpay struct {
	userId string
	userPass string
	merchantId string
	merchantName string
	signature string
}

func (fastpayConstruct *fastpay) construct() {
	fastpayConstruct.userId = os.Getenv("FASPAY_DEBIT_USER_ID")
	fastpayConstruct.userPass = os.Getenv("FASPAY_DEBIT_PASSWORD")
	fastpayConstruct.merchantName = os.Getenv("FASPAY_DEBIT_MERCHANT_NAME")
	fastpayConstruct.merchantId = os.Getenv("FASPAY_DEBIT_MERCHANT_ID")
	md5Hash := md5.Sum([]byte(fastpayConstruct.userId + fastpayConstruct.userPass))
   	md5String := hex.EncodeToString(md5Hash[:])
    sha := sha1.New()
    sha.Write([]byte(md5String))
	encryptedSha1 := sha.Sum(nil)
	fastpayConstruct.signature = fmt.Sprintf("%x", encryptedSha1)
}

func NewPaymentMethodService(repository repositories.PaymentMethodRepository) PaymentMethodService {
	return &paymentMethodService{
		repositories: repository,
	}
}

func (usecases *paymentMethodService) GetAllPaymentMethods(total_premium string, status string) []entity.PaymentMethodWithPremium {
	// get all payment method + rates
	allPaymentMethod := usecases.repositories.GetAllPaymentMethods(status)
	AllPaymentMethodWithTotalPremium := []entity.PaymentMethodWithPremium{}
	// change total premium query from string (A) into integer (i)
	tempPremium, err := strconv.Atoi(total_premium)
	if err != nil {
		fmt.Println(err)
	}
	// loop all payment to add the total payment and fee
	for _, value := range allPaymentMethod{
		// adding data to schema
		data := entity.PaymentMethodWithPremium{}
		data.ID = value.ID
		data.Code = value.Code
		data.InitPaymentCode = value.InitPaymentCode
		data.RenewalPaymentCode = value.RenewalPaymentCode
		data.FastpayCode = value.FastpayCode
		data.BankCode = value.BankCode
		data.Name = value.Name
		data.PaymentLogo = value.PaymentLogo
		data.Status = value.Status
		data.Spec = value.Spec
		data.CreatedAt = value.CreatedAt
		data.UpdatedAt = value.UpdatedAt
		// total payment calculation
		if (total_premium != ""){
			for _, valueRate := range value.PaymentMethodRate {
				// checking transaction range
				if valueRate.MinTransaction <= tempPremium && valueRate.MaxTransaction >= tempPremium {
					// total premium
					data.TotalPremium = tempPremium
					// calculate total fee
					data.Fee = utils.EvaluateStringToFormula(valueRate.FormulaFee, tempPremium)
					// total payment
					data.TotalPayment = data.TotalPremium + data.Fee
				}
			}
		}
		AllPaymentMethodWithTotalPremium = append(AllPaymentMethodWithTotalPremium, data) 
	}
	return AllPaymentMethodWithTotalPremium
}

func (usecases *paymentMethodService) GetPaymentMethod(id string) []entity.PaymentMethod {
	return usecases.repositories.GetPaymentMethod(id)
}

func (usecases *paymentMethodService) GetPaymentMethodByCode(code string) entity.PaymentMethod {
	return usecases.repositories.GetPaymentMethodByCode(code)
}

func (usecases *paymentMethodService) SavePaymentMethod(paymentMethod entity.PaymentMethod) (int, error) {
	return usecases.repositories.SavePaymentMethod(paymentMethod)
}

func (usecases *paymentMethodService) UpdatePaymentMethod(paymentMethod entity.PaymentMethod) error {
	return usecases.repositories.UpdatePaymentMethod(paymentMethod)
}

func (usecases *paymentMethodService) DeletePaymentMethod(paymentMethod entity.PaymentMethod) error {
	return usecases.repositories.DeletePaymentMethod(paymentMethod)
}

func (usecases *paymentMethodService) GenerateVaSignature(code string, proposal_group_number string) string {
	getPaymentMethodDetail := usecases.repositories.GetActivePaymentMethodByCode(code)
	fmt.Println(proposal_group_number)
	fmt.Println(code)
	fmt.Println(getPaymentMethodDetail.ID)
	fmt.Println(os.Getenv("FASPAY_VA_MERCHANT_CODE"))
	proposalGroupNumber := proposal_group_number+"00"
	generateVa := getPaymentMethodDetail.BankCode + os.Getenv("FASPAY_VA_MERCHANT_CODE") + fmt.Sprintf("%09d", utils.StringToInt(proposalGroupNumber))

	return generateVa
}

func (usecases *paymentMethodService) ConnectionTest() *http.Response{
	// construct fastpay data
	fastpayData := new(fastpay)
	fastpayData.construct()

	// request fastpay data
	fmt.Println(fastpayData.userId)
	fmt.Println(fastpayData.userPass)
	fmt.Println(fastpayData.merchantId)
	fmt.Println(fastpayData.merchantName)
	fmt.Println(fastpayData.signature)
	jsonMap := map[string]string {
		"request"     		: "Daftar Payment Channel",
		"merchant_id"     	: os.Getenv("FASPAY_DEBIT_MERCHANT_ID"),
		"merchant"    		: os.Getenv("FASPAY_DEBIT_MERCHANT_NAME"),
		"signature"    		: os.Getenv("FASPAY_DEBIT_MERCHANT_ID"),
	}
	jsonValue, _ := json.Marshal(jsonMap)
	result, err := utils.CreateHttpRequest("POST", os.Getenv("FASPAY_DEBIT_CHANNEL_URL"), jsonValue)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
	return result
}