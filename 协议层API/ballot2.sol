pragma solidity ^0.4.0;

contract HelloWorld{
    uint public balance;
    string public data;
    string[] public ipfsHashs;
    mapping(string=>string)name2hash;
    function update(uint amount) payable returns (address, uint){
        balance += amount;
        return (msg.sender, balance);
    }
    function sayHello() payable returns (string){
        return "hello world";
    }
    function add2Ipfspool(string name,string ipfsHash)payable returns(string){
        name2hash[name]=ipfsHash;
        if(ipfsHashs.push(ipfsHash)>0)
            return ("ok");
        return ("fail");
    }
    function getIpfsHashByPool(string name) constant returns(string){
        data=name2hash[name];
        return name2hash[name];
    }
}
