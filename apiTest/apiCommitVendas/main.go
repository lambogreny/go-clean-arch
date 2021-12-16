package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	url := "http://172.30.0.77:8881/api/v1/sales"

	var initialPhone int = 989831579591

	for i := 0; i < 50; i++ {

		payloadString := fmt.Sprintf("{ \n   \"client\":{\n      \"spc\":{\n         \"status\":1,\n         \"text\":\"Retorno do motor de regra\"\n      },\n      \"personalData\":{\n         \"fullName\":\"Teste Augusyopçsd[a ''' adsda asdlç!! (\",\n         \"email\":\"TesteDoEmail1231@'email\",\n         \"rg\":\"222222222\",\n         \"document\":\"34029121802\",\n         \"birthDate\":\"1900-01-01\",\n         \"motherName\":\"Nome da mae\"\n      },\n      \"contract\":{\n         \"address\":{\n            \"billingCity\":\"Salto\",\n            \"billingState\":\"SP\",\n            \"billingStreet\":\"rua de 31111111111111111111111111111\",\n            \"billingDistrict\":\"Bairro de testes\",\n            \"billingPostalCode\":\"11111111\",\n            \"billingHouseNumber\":29,\n            \"billingComplement\":\"Complementando o endereco de testes\",\n          \n\t\t\t\t\t  \"installationCity\":\"Salto de Pirapora\",\n            \"installationState\":\"SP\",\n            \"installationStreet\":\"dasda\",\n            \"installationDistrict\":\"Bairro de testes 2\",\n            \"installationPostalCode\":\"22222222\",\n            \"installationHouseNumber\":4,\n            \"installationComplement\":\"Complementando o endereco de instalacao\",\n            \"country\": \"Augusto\"\n         },\n         \"installationInfo\":{\n            \"amountInstallments\":\"10\",\n            \"amountAccessionInstallments\":1,\n            \"migratedProvider\":0,\n            \"loyaltyContract\":1,\n            \"receiveInvoice\":0\n         },\n         \"geolocation\":{\n            \"lat\":\"-22.824564654101465\",\n            \"lng\":\"-47.269521772525685\"\n         },\n         \"phone\":{\n            \"phoneType\":0,\n            \"phone\":\"%d\"\n         },\n         \"plans\":{\n            \"planName\":\"F-H.15M.150G\",\n            \"additionalPlans\":[\n               1\n            ],\n            \"sellerId\":789,\n            \"promotionId\": 7\n         },\n         \"tipo_contrato\": 1\n      },\n      \"scheduling\":{\n         \"date\":\"2021-12-23 06:00:00\",\n         \"serviceOrderTypeId\":\"33a33091-5a90-47a3-ab8c-7658b54fe44b\"\n      },\n      \"installationInfra\":{\n        \"ctoId\": \"SCT-163\"\n      }\n   }\n}", initialPhone)

		payload := strings.NewReader(payloadString)

		req, _ := http.NewRequest("POST", url, payload)

		req.Header.Add("Content-Type", "application/json")

		res, _ := http.DefaultClient.Do(req)

		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)

		fmt.Println(string(body))

		//fmt.Println(reflect.TypeOf(payloadString))
		//fmt.Println(payloadString)
		initialPhone += 1

	}

}
