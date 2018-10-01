package main

import (
	"fmt"
	"log"
	"math/big"
  "strings"
	"os"


	//"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
  	"github.com/ethereum/go-ethereum/ethclient"
  	//"github.com/ethereum/go-ethereum/common"
	//"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common"
	//"context"
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
const key=`{"address":"397f448c0c94b1abd0bb158effcf1de5b2bcb1ae","crypto":{"cipher":"aes-128-ctr","ciphertext":"d2cc2ba89e13228413b8a5306e842d89a1fb1ee5a791ce380f77c1e643a8e69a","cipherparams":{"iv":"931caef142f18681a425256a7cd19271"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"a6ae8fa0feb58067268c73f4afceee4fe71f07c8b54569a25f9c3adf352fb5c9"},"mac":"2a31b952d391a304db46e851f8447b9b49a2eb2e70a0d18915465a107757146c"},"id":"191f0e6f-b5c9-4fc7-99cb-013dfc101acc","version":3}`
//const key=`{"address":"020301d472a3533f3785f62601a9304a87a52788","crypto":{"cipher":"aes-128-ctr","ciphertext":"140b2a13b21859a10c246b4f9d6353abac97d7113a183b15e81b5a14892a86a6","cipherparams":{"iv":"c61cde641a9a91cb6fba1b1e6b195485"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"bc28d2d2241b4cc1f5ee4694a21408771bb156d4364c3b4479bedcfd63c1d1b8"},"mac":"bcfec911fe2ebdfefaba966b97e5d21aafe081f189cf8c4860b62e261a5f7121"},"id":"55831c46-3118-48b2-a250-e17fbaf1373a","version":3}`
//const contractAddress="0x17ee0ae98e3dfa57bf7a878e5a30d8e7b3dee52c"

func main() {



	fileName:=os.Args[1]
	ipfsHash_des:=os.Args[2]
	contractAddress:=os.Args[3]



  //const address :=contractAddress
  //const address2="0xb0e6867aae16331236e5954804a51b6a0d2b2cbd" //bidcomputing
  //const address3="0x246eda5ecf5349e8aed2f50033a4ed9af7a827b4" //biddata
  //const key = "//Users//huyifan//go-ethereum//build//bin//data/UTC--2018-09-17T12-53-35.996167402Z--020301d472a3533f3785f62601a9304a87a52788"

  //conn, err := ethclient.Dial("/Users/huyifan/go-ethereum/build/bin/data/geth.ipc")
  conn, err := ethclient.Dial("/home/ec2-user/eth/go-ethereum/build/bin/chain/geth.ipc")
  fmt.Println("connect to local geth node...",conn)
  if err != nil {
      log.Fatalf("could not connect to local node: %v", err)
  }


	fmt.Println("get the auth.....")
	 auth, err := bind.NewTransactor(strings.NewReader(key), "abc")
	 if err != nil {
	 	 log.Fatalf("could not create auth: %v", err)
	 }
	 alloc := make(core.GenesisAlloc)
	 alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(1337000000000)}
	sim := backends.NewSimulatedBackend(alloc,100000000)

	 fmt.Println("get the contract object...")

  //ballot2 contract
  token, err := NewMain2(common.HexToAddress(contractAddress), conn)

	//_,_,token,err:=DeployMain2(auth,conn)
   if err != nil {
       log.Fatalf("Failed to instantiate a Token contract: %v", err)
   }
   fmt.Println("contract token======>:",token)
   //fmt.Println("get the auth.....")
   //bidcomputing contract
   /*token2, err := NewMain3(common.HexToAddress(address2), conn)
    if err != nil {
        log.Fatalf("Failed to instantiate a Token contract: %v", err)
    }
    fmt.Println("contract token2======>:",token2)


    //biddata contract
    token3, err := NewMain4(common.HexToAddress(address), conn)
     if err != nil {
         log.Fatalf("Failed to instantiate a Token contract: %v", err)
     }
     fmt.Println("contract token3======>:",token3)*/


   // fmt.Println("token:=====>",token)
  _,err=token.Add2Ipfspool(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasLimit: 238162,
		Value:    big.NewInt(10),
	},fileName,ipfsHash_des)

   //fmt.Println("returnedData======>:",returnedData)
   if err != nil {
       log.Fatalf("add Ipfs Hash to pool err:%v", err)
   }
  //fmt.Printf("returned Value is:%s\n", returnedData)


  	fmt.Println("Mining...")
	//sim.Commit()


  info, _ := token.GetIpfsHashByPool(&bind.CallOpts{Pending:true}, fileName)
	fmt.Println("File's Ipfs hash is: %v\n", info)
  //sim.Commit()

	returnedData,err:=token.Data(nil)
  //returnedData,err:=token.Data(&bind.CallOpts{})
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

 returnedBalance,err:=token.Balance(&bind.CallOpts{Pending:true})
 if err!=nil{
   log.Fatalf("get balance err:%v", err)
 }
 fmt.Printf("the balance is %d: ",returnedBalance)


//sim.Commit()


  /*fmt.Println("=====test bidcomputing=====")

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




  fmt.Println("=====test biddata=======")

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
