
package main
 
import (
                "errors"
                "fmt"
               
                "github.com/hyperledger/fabric/core/chaincode/shim"   
)
 
// transaction will implement the processes
type SBITransaction struct {
}
 
//setTransactionInitiationMap
func StoreTransactionDetInBC(stub shim.ChaincodeStubInterface,bytesread []byte) error {
 
    fmt.Printf("SBIMap : Entering into StoreTransactionDetInBC function")
    err = stub.PutState("TransactionItemMap", bytesread)
                if err != nil {
                                fmt.Printf("SBIMap : Failed to set the TransactionItemMap %v\n", err)
                                return errors.New("SBIMap : Failed to set the TransactionItemMap")
                }
    return nil
}
 
 
// Init sets up the chaincode
func (t *SBITransaction) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
                fmt.Println("SBIMap : Inside INIT for test chaincode")
                return nil, nil
}
 
// Query the chaincode
func (t *SBITransaction) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
                                fmt.Println("SBIMap : Query")
                    return nil, nil
}
 
// Invoke the function in the chaincode
func (t *SBITransaction) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
                fmt.Println("SBIMap : Invoke")
                return nil, nil
}
 
func main() {
                err := shim.Start(new(SBITransaction))
                if err != nil {
                                fmt.Println("SBIMap : Could not start SBITransaction")
                } else {
                                fmt.Println("SBIMap : SBITransaction successfully started")
                }
 
}
