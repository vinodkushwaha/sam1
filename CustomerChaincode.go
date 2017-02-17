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
	CustomerDataObj.CustomerName.CUSTOMER_FIRST_NAME = args[0]
	CustomerDataObj.CustomerName.CUSTOMER_MIDDLE_NAME = args[1]
	CustomerDataObj.CustomerName.CUSTOMER_LAST_NAME   = args[2]
	CustomerDataObj.PAN_NUMBER = args[3]
	CustomerDataObj.AADHAR_NUMBER = args[4]
	CustomerDataObj.CUSTOMER_DOB = args[5]
	CustomerDataObj.CUSTOMER_RESIDENT_STATUS = args[6]
	CustomerDataObj.CUSTOMER_KYC_PROCESS_DATE = args[7]
	CustomerDataObj.CUSTOMER_KYC_FLAG = args[8]
	//Code for CustomerResidenceAddr Initialization
	CustomerDataObj.CustomerResidenceAddr.RESIDENCE_HOUSE_NO = args[9]
	CustomerDataObj.CustomerResidenceAddr.RESIDENCE_STREET_NAME = args[10]
	CustomerDataObj.CustomerResidenceAddr.RESIDENCE_AREA   = args[11]
	CustomerDataObj.CustomerResidenceAddr.RESIDENCE_CITY_CODE = args[12]
	CustomerDataObj.CustomerResidenceAddr.RESIDENCE_STATE = args[13]
	CustomerDataObj.CustomerResidenceAddr.RESIDENCE_COUNTRY   = args[14]
	//Code for CustomerPermanentAddr Initialization
	CustomerDataObj.CustomerPermanentAddr.PERMANENT_HOUSE_NO = args[15]
	CustomerDataObj.CustomerPermanentAddr.PERMANENT_STREET_NAME = args[16]
	CustomerDataObj.CustomerPermanentAddr.PERMANENT_AREA   = args[17]
	CustomerDataObj.CustomerPermanentAddr.PERMANENT_CITY_CODE = args[18]
	CustomerDataObj.CustomerPermanentAddr.PERMANENT_STATE = args[19]
	CustomerDataObj.CustomerPermanentAddr.PERMANENT_COUNTRY   = args[20]
	//Code for CustomerOfficeAddr Initialization
	CustomerDataObj.CustomerOfficeAddr.OFFICE_NAME = args[21]
	CustomerDataObj.CustomerOfficeAddr.OFFICE_STREET_NAME = args[22]
	CustomerDataObj.CustomerOfficeAddr.OFFICE_AREA   = args[23]
	CustomerDataObj.CustomerOfficeAddr.OFFICE_CITY_CODE = args[24]
	CustomerDataObj.CustomerOfficeAddr.PERMANENT_STATE = args[25]
	CustomerDataObj.CustomerOfficeAddr.OFFICE_COUNTRY   = args[26]
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

	var customer_name string // Entities
	var customer_id string
	var customer_dob string
	var err error
	var resAsBytes []byte

	if len(args) != 3 {
		return nil, errors.New("Incorrect number of arguments. Expecting 3 parameters to query")
	}

	customer_id = args[0]
	customer_name = args[1]
	customer_dob = args[2]

	resAsBytes, err = t.GetCustomerDetails(stub, customer_name, customer_id, customer_dob)

	fmt.Printf("Query Response:%s\n", resAsBytes)

	if err != nil {
		return nil, err
	}

	return resAsBytes, nil
}

func (t *CustomerChaincode)  GetCustomerDetails(stub shim.ChaincodeStubInterface, customer_name string, customer_id string, customer_dob string) ([]byte, error) {

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

	if customer_id == "" {
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
		
		fmt.Printf("Output from customer_id: %s\n", customer_id)
		fmt.Printf("Output from obj.CUSTOMER_ID: %s\n", obj.CUSTOMER_ID)
		fmt.Printf("Output from customer_name: %s\n", customer_name)
		fmt.Printf("Output from obj.CUSTOMER_NAME: %s\n", obj.CUSTOMER_NAME)
		fmt.Printf("Output from customer_dob: %s\n", customer_dob)
		fmt.Printf("Output from obj.CUSTOMER_DOB: %s\n", obj.CUSTOMER_DOB)
		if ((obj.CUSTOMER_ID) == customer_id){
			CustomerTxObjects1 = append(CustomerTxObjects1,obj)
			//requiredObj = obj
			objFound = true
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
