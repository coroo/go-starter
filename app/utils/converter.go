package utils

import (
	"strconv"
	"gopkg.in/Knetic/govaluate.v2"
	"golang.org/x/crypto/bcrypt"
)

func StringToInt(req string) int {
	res, _ := strconv.Atoi(req)
	return res
}

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
func EvaluateStringToFormula(formula string, parameter int) int {
	// add formula
	expression, _ := govaluate.NewEvaluableExpression(formula);
	// create parameter to add to formula
	parameters := make(map[string]interface{}, 8)
	parameters["premi"] = parameter;
	// compile
	result, _ := expression.Evaluate(parameters);
	// add result into data
	return int(result.(float64))
}
