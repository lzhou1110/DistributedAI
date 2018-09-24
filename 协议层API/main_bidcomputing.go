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
//remote account
/*const key = `
{
  "address":"4199aa2326c257ccda5de1507e599cce17cd264f",
  "crypto":{
    "cipher":"aes-128-ctr",
    "ciphertext":"e42edf56411abb7ee8b4c15ee118c8bdfc7be2ad0eb655454a38cef3b602652f",
    "cipherparams":{"iv":"d8454d46ce2b1c381fc21fc11d6f5bc4"},
    "kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"3b3b910e90aa8e6c1de8875cc0a5f262c7704894931becbfe5603402adf4478f"},
    "mac":"8592a2a3bb207a921741239e8f689d72c5686126dabbdafa4efa72c6c9b8d065"},
    "id":"b0b0bf4a-af7c-4375-936b-7d033959e434","version":3
  }
`*/
//local accounts
const key=`{"address":"a26ef2ded2fad7f9462d2d60e928a27bb96db252","crypto":{"cipher":"aes-128-ctr","ciphertext":"fda2b3b00ed1855f1859f9dc4ed46d0f93460ef6158d76a25c2ee118970dc3b4","cipherparams":{"iv":"f3ef060a5cd06a3c4a881dbc06c53cfd"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"555d6950bbc67c9e412e7aed5ceaef12335b889515b12e9957d15093042f56ae"},"mac":"8bb8d2e157b494b295781bbc1b0b0633bcb517b44e3430a1ae1839899ae75305"},"id":"487ba13d-61e9-494b-aeb0-e73d3ba9c6f7","version":3}`




func main() {
  //const address ="0xd50ebd8c3f659b65aa642b21f268df538fb2a2bc"
  const address2="0xb8168f7cb6891f5e386358c199e7e68d98c636a7" //bidcomputing
  //const address3="0x246eda5ecf5349e8aed2f50033a4ed9af7a827b4" //biddata
  //const key = "//Users//huyifan//go-ethereum//build//bin//data/UTC--2018-09-17T12-53-35.996167402Z--020301d472a3533f3785f62601a9304a87a52788"
  conn, err := ethclient.Dial("/root/eth/go-ethereum/build/bin/chain/geth.ipc")
  //conn, err := ethclient.Dial("/Users/huyifan/go-ethereum/build/bin/data/geth.ipc")
  fmt.Println("connect to local geth node...",conn)
  if err != nil {
      log.Fatalf("could not connect to local node: %v", err)
  }
  fmt.Println("get the contract object...")
  //ballot2 contract
  /*token, err := NewMain2(common.HexToAddress(address), conn)
   if err != nil {
       log.Fatalf("Failed to instantiate a Token contract: %v", err)
   }
   fmt.Println("contract token======>:",token)*/
   //fmt.Println("get the auth.....")
   //bidcomputing contract
   token2, err := NewMain3(common.HexToAddress(address2), conn)
    if err != nil {
        log.Fatalf("Failed to instantiate a Token contract: %v", err)
    }
    fmt.Println("contract token2======>:",token2)


    //biddata contract
    /*token3, err := NewMain4(common.HexToAddress(address), conn)
     if err != nil {
         log.Fatalf("Failed to instantiate a Token contract: %v", err)
     }
     fmt.Println("contract token3======>:",token3)

     fmt.Println("get the auth.....")*/
   auth, err := bind.NewTransactor(strings.NewReader(key), "abc")
   if err != nil {
       log.Fatalf("could not create auth: %v", err)
   }
    alloc := make(core.GenesisAlloc)
    alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(1337000000000)}
    sim := backends.NewSimulatedBackend(alloc,100000000)
   // fmt.Println("token:=====>",token)
  /*_,err=token.Add2Ipfspool(&bind.TransactOpts{
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

  info, _ := token.GetIpfsHashByPool(&bind.CallOpts{}, "filename")
	fmt.Println("File's Ipfs hash is: %v\n", info)
  //sim.Commit()


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
  fmt.Println("Mining...")
 sim.Commit()
 returnedBalance,err:=token.Balance(&bind.CallOpts{})
 if err!=nil{
   log.Fatalf("get balance err:%v", err)
 }
 fmt.Println("the balance is %d: ",returnedBalance)


//sim.Commit()*/


  fmt.Println("=====test bidcomputing=====")

  token2.AskTraining(&bind.TransactOpts{
   From:     auth.From,
   Signer:   auth.Signer,
   GasLimit: 238162,
   Value:    big.NewInt(10),
 },"DataSchemaAddress","MetadataAddress","ModelAddress","StrategyAddress","ComputionAddress")
 //fmt.Println("test bidcomputing data: %v\n", returnTraningData)
 sim.Commit()

 info_res1,_:=token2.Status(&bind.CallOpts{})
fmt.Println("test bidcomuting askTraing result: ", info_res1)




  /*fmt.Println("=====test biddata=======")

  token3.Buydata(&bind.TransactOpts{
   From:     auth.From,
   Signer:   auth.Signer,
   GasLimit: 238162,
   Value:    big.NewInt(10),
  })
  sim.Commit()
  info_res2,_:=token3.Status(&bind.CallOpts{})
  fmt.Println("test biddata buydata result:  ", info_res2)*/

  //fmt.Println("test biddata data: %v\n", returnedBuyData)

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
