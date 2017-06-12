/* Test for storing ZipprCoins */

package main

import (

	 "fmt"
	"errors"

        "github.com/hyperledger/fabric/core/chaincode/shim"

)


type SimpleChaincode struct {

}

type ZipprCoinOwnership struct {
	Username  	string `json:"username"`
	ZipprCoins	float64 `json:"zippr_coins"`
}


// ============================================================================================================================
// Main
// ============================================================================================================================
func main() {
        err := shim.Start(new(SimpleChaincode))
        if err != nil {
                fmt.Printf("Error starting Simple chaincode - %s", err)
        }
}



// Init resets all the things
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
        if len(args) != 1 {
                return nil, errors.New("Incorrect number of arguments. Expecting 1")
        }

        err := stub.PutState("hello_world_zippr_coin", []byte(args[0]))
        if err != nil {
         return nil, err
        }

        return nil, nil
}


// Invoke is our entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
        fmt.Println("invoke is running " + function)

        // Handle different functions
        if function == "init" {                                                                                                 //initialize the chaincode state, used as reset
                return t.Init(stub, "init", args)
          }
//        } else if function == "write" {
//                return t.write(stub, args)
//        }
        fmt.Println("invoke did not find func: " + function)                                    //error

        return nil, errors.New("Received unknown function invocation: " + function)
}


// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
        fmt.Println("query is running " + function)

        // Handle different functions
/*
        if function == "read" {                                                                                 //read a variable
                fmt.Println("hi there " + function)                                             //error
                return t.read(stub,args);
        }
*/
        fmt.Println("query did not find func: " + function)                                             //error

        return nil, errors.New("Received unknown function query: " + function)
}
