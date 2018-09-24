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
	"address":"a26ef2ded2fad7f9462d2d60e928a27bb96db252","crypto":{"cipher":"aes-128-ctr","ciphertext":"fda2b3b00ed1855f1859f9dc4ed46d0f93460ef6158d76a25c2ee118970dc3b4","cipherparams":{"iv":"f3ef060a5cd06a3c4a881dbc06c53cfd"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"555d6950bbc67c9e412e7aed5ceaef12335b889515b12e9957d15093042f56ae"},"mac":"8bb8d2e157b494b295781bbc1b0b0633bcb517b44e3430a1ae1839899ae75305"},"id":"487ba13d-61e9-494b-aeb0-e73d3ba9c6f7","version":3}
`

func main() {
  const address = "0x0d3f9573b01c55084705f1b455f3b6978a7d6933"
  //const key = "//Users//huyifan//go-ethereum//build//bin//data/UTC--2018-09-17T12-53-35.996167402Z--020301d472a3533f3785f62601a9304a87a52788"

  conn, err := ethclient.Dial("/root/eth/go-ethereum/build/bin/chain/geth.ipc")
  fmt.Println("connect to local geth node...",conn)
  if err != nil {
      log.Fatalf("could not connect to local node: %v", err)
  }
  fmt.Println("get the contract object...")
  token, err := NewMain2(common.HexToAddress(address), conn)
   if err != nil {
       log.Fatalf("Failed to instantiate a Token contract: %v", err)
   }
   fmt.Println("contract token======>:",token)
   fmt.Println("get the auth.....")
   auth, err := bind.NewTransactor(strings.NewReader(key), "abc")
   if err != nil {
       log.Fatalf("could not create auth: %v", err)
   }
    alloc := make(core.GenesisAlloc)
    alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(1337000000000)}
    sim := backends.NewSimulatedBackend(alloc,1000000000)
   // fmt.Println("token:=====>",token)
  _,err=token.Add2Ipfspool(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasLimit: 238162,
		Value:    big.NewInt(10),
	},"filename","ipfsHashValue")

   //fmt.Println("returnedData======>:",returnedData)
   if err != nil {
       log.Fatalf("add Ipfs Hash to pool err:%v", err)
   }
  //fmt.Printf("returned Value is:%s\n", returnedData)


  fmt.Println("Mining...")
	sim.Commit()



  returnedData,err:=token.Data(&bind.CallOpts{})
  if err!=nil {
    log.Fatalf("get data err:%v", err)
  }

  fmt.Println("IPFS Hash of uploaded file: ", returnedData)

  token.Update(&bind.TransactOpts{
   From:     auth.From,
   Signer:   auth.Signer,
   GasLimit: 238162,
   Value:    big.NewInt(10),
 },big.NewInt(102))
 sim.Commit()
 returnedBalance,err:=token.Balance(&bind.CallOpts{})
 if err!=nil{
   log.Fatalf("get balance err:%v", err)
 }
 fmt.Println("the balance is %d: ",returnedBalance)

  info, _ := token.GetIpfsHashByPool(&bind.CallOpts{}, "filename")
	fmt.Println("File's Ipfs hash is: %v\n", info)

	//key, _ := crypto.GenerateKey()
	//auth := bind.NewKeyedTransactor(key)


//
// 	// deploy contract
// 	/*addr, _, contract, err := DeployWinnerTakesAll(auth, sim, big.NewInt(10), big.NewInt(time.Now().Add(2*time.Minute).Unix()), big.NewInt(time.Now().Add(5*time.Minute).Unix()))
// 	if err != nil {
// 		log.Fatalf("could not deploy contract: %v", err)
// 	}*/
//   fmt.Printf("Instantiating contract at address %s...\n", auth.From.String())
//   addr:=common.HexToAddress(address)
//   contract,err:=NewMain(addr,sim)
//   if err !=nil{
//     log.Fatalf("could not found contract:")
//   }
//
// 	// interact with contract
// 	fmt.Printf("Contract found to %s\n", addr.String())
// //	deadlineCampaign, _ := contract.DeadlineCampaign(nil)
//     sayHelloCompagin,_:=contract.SayHello(&bind.TransactOpts{
// 		From:     auth.From,
// 		Signer:   auth.Signer,
// 		GasLimit: 2381623,
// 		Value:    big.NewInt(10),
// 	})
// 	//fmt.Printf("Pre-mining Campaign Deadline: %s\n", deadlineCampaign)
//   fmt.Printf("pre-mining Campagin sayHello: %s\n",sayHelloCompagin)
// 	fmt.Println("Mining...")
// 	// simulate mining
// 	sim.Commit()

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
