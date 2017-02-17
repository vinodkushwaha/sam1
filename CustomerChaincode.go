package main

import (
	"errors"
	"fmt"
	//"strconv"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	//"github.com/golang/protobuf/ptypes/timestamp"
)

// Customer Chaincode implementation
type CustomerChaincode struct {
}

var customerIndexTxStr = "_customerIndexTxStr"

type CustomerDoc struct {
    DOCUMENT_NAME string `json:"DOCUMENT_NAME"`
	DOCUMENT_STRING string `json:"DOCUMENT_STRING"`
}
type CustomerResidenceAddr struct {
    RESIDENCE_HOUSE_NO  string `json:"RESIDENCE_HOUSE_NO"`
	RESIDENCE_STREET_NAME string `json:"RESIDENCE_STREET_NAME"`
	RESIDENCE_AREA  string `json:"RESIDENCE_AREA"`
	RESIDENCE_CITY_CODE string `json:"RESIDENCE_CITY_CODE"`
	RESIDENCE_STATE  string `json:"RESIDENCE_STATE"`
	RESIDENCE_COUNTRY string `json:"RESIDENCE_COUNTRY"`
}
type CustomerPermanentAddr struct {
    PERMANENT_HOUSE_NO  string `json:"PERMANENT_HOUSE_NO"`
	PERMANENT_STREET_NAME string `json:"PERMANENT_STREET_NAME"`
	PERMANENT_AREA  string `json:"PERMANENT_AREA"`
	PERMANENT_CITY_CODE string `json:"PERMANENT_CITY_CODE"`
	PERMANENT_STATE  string `json:"PERMANENT_STATE"`
	PERMANENT_COUNTRY string `json:"PERMANENT_COUNTRY"`
}
type CustomerOfficeAddr struct {
    OFFICE_NAME  string `json:"OFFICE_NAME"`
	OFFICE_STREET_NAME string `json:"OFFICE_STREET_NAME"`
	OFFICE_AREA  string `json:"OFFICE_AREA"`
	OFFICE_CITY_CODE string `json:"OFFICE_CITY_CODE"`
	OFFICE_STATE  string `json:"OFFICE_STATE"`
	OFFICE_COUNTRY string `json:"OFFICE_COUNTRY"`
}
type CustomerName struct{
    CUSTOMER_FIRST_NAME  string `json:"CUSTOMER_FIRST_NAME"`
	CUSTOMER_MIDDLE_NAME string `json:"CUSTOMER_MIDDLE_NAME"`
	CUSTOMER_LAST_NAME  string `json:"CUSTOMER_LAST_NAME"`
}

type CustomerData struct{
	CUSTOMER_NAME CustomerName
	PAN_NUMBER string `json:"PAN_NUMBER"`
	AADHAR_NUMBER string `json:"AADHAR_NUMBER"`
	CUSTOMER_DOB string `json:"CUSTOMER_DOB"`
	CUSTOMER_RESIDENT_STATUS string `json:"RESIDENT_STATUS"`
	CUSTOMER_KYC_PROCESS_DATE string `json:"CUSTOMER_KYC_PROCESS_DATE"`
	CUSTOMER_KYC_FLAG string `json:"CUSTOMER_KYC_FLAG"`
	CUSTOMER_RESIDENCE_ADDR CustomerResidenceAddr
	CUSTOMER_PERMANENT_ADDR CustomerPermanentAddr
	CUSTOMER_OFFICE_ADDR CustomerOfficeAddr
	CUSTOMER_DOC []CustomerDoc
	}


func (t *CustomerChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	var err error
	// Initialize the chaincode

	fmt.Printf("Deployment of Customer ChainCode is completed\n")

	var emptyCustomerTxs []CustomerData
	jsonAsBytes, _ := json.Marshal(emptyCustomerTxs)
	err = stub.PutState(customerIndexTxStr, jsonAsBytes)
	if err != nil {
		return nil, err
	}


	return nil, nil
}

// Add customer data for the policy
func (t *CustomerChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if function == customerIndexTxStr {
		return t.RegisterCustomer(stub, args)
	}
	return nil, nil
}

func (t *CustomerChaincode)  RegisterCustomer(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var CustomerDataObj CustomerData
	var CustomerDataList []CustomerData
	var err error
   	fmt.Printf("********pankaj CUSTOMER_DOC:%d\n", len(args))
	
	if len(args) < 4 {
		return nil, errors.New("Incorrect number of arguments. Need 4 arguments")
	}

	// Initialize the chaincode
	
	//Code for Name Initialization
	CustomerDataObj.CUSTOMER_NAME.CUSTOMER_FIRST_NAME = args[0]
	CustomerDataObj.CUSTOMER_NAME.CUSTOMER_MIDDLE_NAME = args[1]
	CustomerDataObj.CUSTOMER_NAME.CUSTOMER_LAST_NAME   = args[2]
	CustomerDataObj.PAN_NUMBER = args[3]
	CustomerDataObj.AADHAR_NUMBER = args[4]
	CustomerDataObj.CUSTOMER_DOB = args[5]
	CustomerDataObj.CUSTOMER_RESIDENT_STATUS = args[6]
	CustomerDataObj.CUSTOMER_KYC_PROCESS_DATE = args[7]
	CustomerDataObj.CUSTOMER_KYC_FLAG = args[8]
	//Code for CustomerResidenceAddr Initialization
	CustomerDataObj.CUSTOMER_RESIDENCE_ADDR.RESIDENCE_HOUSE_NO = args[9]
	CustomerDataObj.CUSTOMER_RESIDENCE_ADDR.RESIDENCE_STREET_NAME = args[10]
	CustomerDataObj.CUSTOMER_RESIDENCE_ADDR.RESIDENCE_AREA   = args[11]
	CustomerDataObj.CUSTOMER_RESIDENCE_ADDR.RESIDENCE_CITY_CODE = args[12]
	CustomerDataObj.CUSTOMER_RESIDENCE_ADDR.RESIDENCE_STATE = args[13]
	CustomerDataObj.CUSTOMER_RESIDENCE_ADDR.RESIDENCE_COUNTRY   = args[14]
	//Code for CustomerPermanentAddr Initialization
	CustomerDataObj.CUSTOMER_PERMANENT_ADDR.PERMANENT_HOUSE_NO = args[15]
	CustomerDataObj.CUSTOMER_PERMANENT_ADDR.PERMANENT_STREET_NAME = args[16]
	CustomerDataObj.CUSTOMER_PERMANENT_ADDR.PERMANENT_AREA   = args[17]
	CustomerDataObj.CUSTOMER_PERMANENT_ADDR.PERMANENT_CITY_CODE = args[18]
	CustomerDataObj.CUSTOMER_PERMANENT_ADDR.PERMANENT_STATE = args[19]
	CustomerDataObj.CUSTOMER_PERMANENT_ADDR.PERMANENT_COUNTRY   = args[20]
	//Code for CustomerOfficeAddr Initialization
	CustomerDataObj.CUSTOMER_OFFICE_ADDR.OFFICE_NAME = args[21]
	CustomerDataObj.CUSTOMER_OFFICE_ADDR.OFFICE_STREET_NAME = args[22]
	CustomerDataObj.CUSTOMER_OFFICE_ADDR.OFFICE_AREA   = args[23]
	CustomerDataObj.CUSTOMER_OFFICE_ADDR.OFFICE_CITY_CODE = args[24]
	CustomerDataObj.CUSTOMER_OFFICE_ADDR.OFFICE_STATE = args[25]
	CustomerDataObj.CUSTOMER_OFFICE_ADDR.OFFICE_COUNTRY   = args[26]
	//Code for the Document Process	
	fmt.Printf("********pankaj CUSTOMER_DOC:%s\n", args[4])
	var number_of_docs int
	number_of_docs = (len(args)-27)/2
	var CustomerDocObjects1 []CustomerDoc
	for i := 0; i < number_of_docs; i++ {
		var CustomerDocObj CustomerDoc
		fmt.Printf("********pankaj CustomerDocObj[i].DOCUMENT_NAMEC:%d\n",i)
		fmt.Printf("********pankaj CustomerDocObj[i].DOCUMENT_NAMEC:%d\n",number_of_docs)
		//CustomerDocObj[i] := CustomerDoc{DOCUMENT_NAME: args[27+(i*2)], DOCUMENT_STRING: args[27+(i*2)]}
		CustomerDocObj.DOCUMENT_NAME = args[27+(i*2)]
		//fmt.Printf("********pankaj CustomerDocObj[i].DOCUMENT_NAMEC:%s\n", CustomerDocObj[i].DOCUMENT_NAME)
		CustomerDocObj.DOCUMENT_STRING = args[28+(i*2)]
		CustomerDocObjects1 = append(CustomerDocObjects1,CustomerDocObj)
	}
	
	CustomerDataObj.CUSTOMER_DOC = CustomerDocObjects1
	
	customerTxsAsBytes, err := stub.GetState(customerIndexTxStr)
	if err != nil {
		return nil, errors.New("Failed to get customer transactions")
	}
	json.Unmarshal(customerTxsAsBytes, &CustomerDataList)

	CustomerDataList = append(CustomerDataList, CustomerDataObj)
	jsonAsBytes, _ := json.Marshal(CustomerDataList)

	err = stub.PutState(customerIndexTxStr, jsonAsBytes)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// Query callback representing the query of a chaincode
func (t *CustomerChaincode) Query(stub shim.ChaincodeStubInterface,function string, args []string) ([]byte, error) {

	var PAN_NUMBER string // Entities
	var AADHAR_NUMBER string
	var err error
	var resAsBytes []byte

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 3 parameters to query")
	}

	PAN_NUMBER = args[0]
	AADHAR_NUMBER = args[1]
	
	resAsBytes, err = t.GetCustomerDetails(stub, PAN_NUMBER, AADHAR_NUMBER)

	fmt.Printf("Query Response:%s\n", resAsBytes)

	if err != nil {
		return nil, err
	}

	return resAsBytes, nil
}

func (t *CustomerChaincode)  GetCustomerDetails(stub shim.ChaincodeStubInterface, PAN_NUMBER string, AADHAR_NUMBER string) ([]byte, error) {

	//var requiredObj CustomerData
	var objFound bool
	CustomerTxsAsBytes, err := stub.GetState(customerIndexTxStr)
	if err != nil {
		return nil, errors.New("Failed to get Customer Records")
	}
	var CustomerTxObjects []CustomerData
	var CustomerTxObjects1 []CustomerData
	json.Unmarshal(CustomerTxsAsBytes, &CustomerTxObjects)
	length := len(CustomerTxObjects)
	fmt.Printf("Output from chaincode: %s\n", CustomerTxsAsBytes)

	if PAN_NUMBER == "" && AADHAR_NUMBER == ""{
		res, err := json.Marshal(CustomerTxObjects)
		if err != nil {
		return nil, errors.New("Failed to Marshal the required Obj")
		}
		return res, nil
	}

	objFound = false
	// iterate
	for i := 0; i < length; i++ {
		obj := CustomerTxObjects[i]
		//if ((customer_id == obj.CUSTOMER_ID) && (customer_name == obj.CUSTOMER_NAME) && (customer_dob == obj.CUSTOMER_DOB)) 
		
	if (PAN_NUMBER != ""){
		if ((obj.PAN_NUMBER) == PAN_NUMBER){
			CustomerTxObjects1 = append(CustomerTxObjects1,obj)
			//requiredObj = obj
			objFound = true
			break;
		}
	}else {
		if ((obj.AADHAR_NUMBER) == AADHAR_NUMBER){
			CustomerTxObjects1 = append(CustomerTxObjects1,obj)
			//requiredObj = obj
			objFound = true
			break;
		}
	}
	}

	if objFound {
		res, err := json.Marshal(CustomerTxObjects1)
		if err != nil {
		return nil, errors.New("Failed to Marshal the required Obj")
		}
		return res, nil
	} else {
		res, err := json.Marshal("No Data found")
		if err != nil {
		return nil, errors.New("Failed to Marshal the required Obj")
		}
		return res, nil
	}
}

func main() {
	err := shim.Start(new(CustomerChaincode))
	if err != nil {
		fmt.Printf("Error starting Customer Simple chaincode: %s", err)
	}
}
