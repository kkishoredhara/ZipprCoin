/* Test for storing ZipprCoins */

package main

import (

	"errors"
	"encoding/json"
	"encoding/binary"
        "fmt"
        "strconv"

	"math"

	"github.com/hyperledger/fabric/core/chaincode/shim"

)


type SimpleChaincode struct {

}

type ZipprCoinTotal struct {
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
        } else if function == "write" {
                return t.write(stub, args)
        }
        fmt.Println("invoke did not find func: " + function)                                    //error

        return nil, errors.New("Received unknown function invocation: " + function)
}


// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
        fmt.Println("query is running " + function)

        // Handle different functions
        if function == "read" {                                                                                 //read a variable
                fmt.Println("hi there " + function)                                             //error
                return t.read(stub,args);
        }
        fmt.Println("query did not find func: " + function)                                             //error

        return nil, errors.New("Received unknown function query: " + function)
}


func (t *SimpleChaincode) write(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
 
  var name, jsonResp string
  var err error
  var zippr_coin_Obj ZipprCoinTotal

  fmt.Println("running write()")

  if len(args) != 2 {
        return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the variable and value to set")
  }

  name = args[0]
 
  stored_zippr_coins_Asbytes, err := stub.GetState(name)

  if err != nil {
	jsonResp = "{\"Error\":\"Failed to get coins for " + name + "\"}"
                return nil, errors.New(jsonResp)
  }

  json.Unmarshal(stored_zippr_coins_Asbytes, &zippr_coin_Obj)


  zippr_coin_addValue, err := strconv.ParseFloat(args[1], 64)
  if err != nil {
      return nil, errors.New("zippr_coin_addValue: argument to write is not a floating number")
  }

  zippr_coin_Obj.ZipprCoins = zippr_coin_Obj.ZipprCoins + zippr_coin_addValue

  jsonAsbytes, _ := json.Marshal(zippr_coin_Obj) 
  err = stub.PutState(name, jsonAsbytes)
  if err != nil {
   return nil, err
  }
  return nil, nil
}


func (t *SimpleChaincode) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
        var name, jsonResp string
        var err error
        var zippr_coin_Obj ZipprCoinTotal

        if len(args) != 1{
                return nil, errors.New("Incorrect number of arguments. Expecting name of the var to query")
        }

        name = args[0]
        valAsbytes, err := stub.GetState(name)

        if err != nil {
                jsonResp = "{\"Error\":\"Failed to get state for " + name + "\"}"
                return nil, errors.New(jsonResp)
        }

       json.Unmarshal(valAsbytes, &zippr_coin_Obj)
      
       return ([]byte(float64ToByte(zippr_coin_Obj.ZipprCoins))), nil
}


func float64ToByte(f float64) []byte {
   var buf [8]byte
   binary.BigEndian.PutUint64(buf[:], math.Float64bits(f))
   return buf[:]
}
