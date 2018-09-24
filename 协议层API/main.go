package main

import (
	"fmt"
	"log"
	"math/big"
  "strings"



	//"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
  "github.com/ethereum/go-ethereum/ethclient"
  "github.com/ethereum/go-ethereum/common"
	//"github.com/ethereum/go-ethereum/crypto"
)

const key = `
{
  "address":"020301d472a3533f3785f62601a9304a87a52788",
  "crypto":{
    "cipher":"aes-128-ctr","ciphertext":"140b2a13b21859a10c246b4f9d6353abac97d7113a183b15e81b5a14892a86a6",
    "cipherparams":{"iv":"c61cde641a9a91cb6fba1b1e6b195485"},
    "kdf":"scrypt",
    "kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"bc28d2d2241b4cc1f5ee4694a21408771bb156d4364c3b4479bedcfd63c1d1b8"},
    "mac":"bcfec911fe2ebdfefaba966b97e5d21aafe081f189cf8c4860b62e261a5f7121"
  },
    "id":"55831c46-3118-48b2-a250-e17fbaf1373a",
    "version":3
  }
`

func main() {
  const address = "0x04daa338cb2bda10d86181d1a830c9735ec6f2aa"
  //const key = "//Users//huyifan//go-ethereum//build//bin//data/UTC--2018-09-17T12-53-35.996167402Z--020301d472a3533f3785f62601a9304a87a52788"

  conn, err := ethclient.Dial("/Users/huyifan/go-ethereum/build/bin/data/geth.ipc")
  fmt.Println("connect to local geth node...",conn)
  if err != nil {
      log.Fatalf("could not connect to local node: %v", err)
  }
  fmt.Println("get the contract object...")
  token, err := NewMain(common.HexToAddress(address), conn)
   if err != nil {
       log.Fatalf("Failed to instantiate a Token contract: %v", err)
   }
   fmt.Println("contract token======>:",token)
   fmt.Println("get the auth.....")
   auth, err := bind.NewTransactor(strings.NewReader(key), "abc")
   if err != nil {
       log.Fatalf("could not create auth: %v", err)
   }
   // fmt.Println("token:=====>",token)
   returnedData, err := token.SayHello(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasLimit: 238162,
		Value:    big.NewInt(10),
	})

   //fmt.Println("returnedData======>:",returnedData)
   if err != nil {
       log.Fatalf("invoke method err:%v", err)
   }
  fmt.Printf("returned Value is:%s\n", returnedData)



	//key, _ := crypto.GenerateKey()
	//auth := bind.NewKeyedTransactor(key)

	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(1337000000000)}
	sim := backends.NewSimulatedBackend(alloc,100000000000)

	// deploy contract
	/*addr, _, contract, err := DeployWinnerTakesAll(auth, sim, big.NewInt(10), big.NewInt(time.Now().Add(2*time.Minute).Unix()), big.NewInt(time.Now().Add(5*time.Minute).Unix()))
	if err != nil {
		log.Fatalf("could not deploy contract: %v", err)
	}*/
  fmt.Printf("Instantiating contract at address %s...\n", auth.From.String())
  addr:=common.HexToAddress(address)
  contract,err:=NewMain(addr,sim)
  if err !=nil{
    log.Fatalf("could not found contract:")
  }

	// interact with contract
	fmt.Printf("Contract found to %s\n", addr.String())
//	deadlineCampaign, _ := contract.DeadlineCampaign(nil)
    sayHelloCompagin,_:=contract.SayHello(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasLimit: 2381623,
		Value:    big.NewInt(10),
	})
	//fmt.Printf("Pre-mining Campaign Deadline: %s\n", deadlineCampaign)
  fmt.Printf("pre-mining Campagin sayHello: %s\n",sayHelloCompagin)
	fmt.Println("Mining...")
	// simulate mining
	sim.Commit()

//after_sayHelloCompagin,_:=contract.SayHello(auth)
	//postDeadlineCampaign, _ := contract.DeadlineCampaign(nil)
	//fmt.Printf("Post-mining Campaign Deadline: %s\n", time.Unix(postDeadlineCampaign.Int64(), 0))

	// create a project
/*	numOfProjects, _ := contract.NumberOfProjects(nil)
	fmt.Printf("Number of Projects before: %d\n", numOfProjects)

	fmt.Println("Adding new project...")
	contract.SubmitProject(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasLimit: 2381623,
		Value:    big.NewInt(10),
	}, "test project", "http://www.example.com")

	fmt.Println("Mining...")
	sim.Commit()

	numOfProjects, _ = contract.NumberOfProjects(nil)
	fmt.Printf("Number of Projects after: %d\n", numOfProjects)
	info, _ := contract.GetProjectInfo(nil, auth.From)
	fmt.Printf("Project Info: %v\n", info)

	// instantiate deployed contract
	fmt.Printf("Instantiating contract at address %s...\n", auth.From.String())
	instContract, err := NewWinnerTakesAll(addr, sim)
	if err != nil {
		log.Fatalf("could not instantiate contract: %v", err)
	}
	numOfProjects, _ = instContract.NumberOfProjects(nil)
	fmt.Printf("Number of Projects of instantiated Contract: %d\n", numOfProjects)*/
}
