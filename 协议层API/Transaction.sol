pragma solidity ^0.4.11;

import "./biddata.sol";
import "./bidcomputing.sol";

contract Transaction {
    
    /*
    mapping(address => uint256) store; 
    mapping(address => bool[3]) role;   // 0 is dator, 1 is modelor, 2 is computor
    */
    
    biddata[] biddatas;
    bidcomputing[] bidcomputings;
    
    // save & lock 100max Transaction
    struct lock{
        uint256 sum;
        uint256[100] q;
        uint8 index;
        uint256[100] createtime;
    }
    mapping (address => lock) locks;
    
    /*
    struct bidDataInfo{
        string DataSchemaAddress;
        string MataDataAddress;
        string PaymentAddress;
    }
    
    struct bidComputingInfo{
        string OperationSchemasAddress;
        string ComputingAddress;
        string PaymentAddress;
        string ComputerAttributesAddress;
    }
    
    struct askForDataInfo{
        string KeyAddress;
        string DataSchemaAddress;
        string TransactionDetailAddress;
    }
    
    struct TransactionInfo{
        string DataSchemaAddress;
        string MetadataAddress;
        string ModelAddress;
        string StrategyAddress;
        string ComputionAddress;
    }
    
    struct uploadTrainResInfo{
        string ResultAddress;
        string TransactionDetailAddress;
    }
    */
    
    /*
    function changerole(bool dator, bool modelor, bool computor){
        role[msg.sender][0] = dator;
        role[msg.sender][1] = modelor;
        role[msg.sender][2] = computor;
    }
    
    
    function storeeth() public payable overflow{
        store[msg.sender] += msg.value;
        Store(msg.sender, msg.value);
    }
    
    
    function withdraw(uint256 ethnum) returns (bool) {
        uint256 amount = store[msg.sender];
        if(amount >= ethnum){
            store[msg.sender] -= ethnum;
            if(!msg.sender.send(ethnum)){
                store[msg.sender] = amount;
                return false;
            }
            Withdrew(msg.sender, ethnum);
            return true;
        }
        return false;
    }
    
    
    function read() returns (uint256){
        return store[msg.sender];
    }
    */
    
    /*
    function pay(address bid, uint256 ethnum)  public returns (bool){
        uint256 amount = store[msg.sender];
        if(amount >= ethnum){
            store[msg.sender] -= ethnum;
            if(!bid.transfer(ethnum)){          // is right?
                store[msg.sender] = amount;
                return false;
            }
            event Pay(msg.sender, ethnum);
            return true;
        }
        return false;
    }
    */
    
    
    // keep all the ether sent to this address
    function() payable public {
        FallbackCalled(msg.sender, msg.value);
    }


    function bidData (
        string DataSchemaAddress,
        string MataDataAddress,
        uint256 Payment
    )public returns(uint256){
        uint256 len = biddatas.push(new biddata(DataSchemaAddress, MataDataAddress, Payment));
        
        return len;     // ???? how to err
    }
    
    
    function bidComputing (
        string OperationSchemasAddress,
        string ComputingAddress,
        uint256 Payment,
        string ComputerAttributesAddress
    ) returns (uint256){
        
        uint256 len = bidcomputings.push(new bidcomputing(OperationSchemasAddress, ComputingAddress, ComputerAttributesAddress, Payment));
        
        return len;     // ???? how
    }
    
    
    /*
    function askTraining (
        string DataSchemaAddress,
        string MetadataAddress,
        string ModelAddress,
        string StrategyAddress,
        string ComputionAddress
    ) public payable returns(uint256) {
        if(store[msg.sender] >= msg.value){
            store[msg.sender] -= msg.value;
            locks[msg.sender].sum += msg.value;
            locks[msg.sender].q[locks[msg.sender].index % 100] = msg.value;
            locks[msg.sender].createtime[locks[msg.sender].index % 100] = now;
            locks[msg.sender].index = (locks[msg.sender]index + 1) % 100;
            AskTraining (
                string DataSchemaAddress,
                string MetadataAddress,
                string ModelAddress,
                string StrategyAddress,
                string ComputionAddress
             );
            return 201;
        }
        return 500;
    }
    */
    


    // avoid overflow
    
    /*
    modifier overflow{
        require(msg.value + store[msg.sender] >= store[msg.sender]);
        _;
    }
    */
    
    event Store(address from, uint256 amount);
    
    event Withdrew(address to, uint256 amount);
    
    event FallbackCalled(address from, uint256 amount);
    
    event BidData(string DataSchemaAddress, string MataDataAddress, string PaymentAddress);
    
    event BidComputing(
        string OperationSchemasAddress,
        string ComputingAddress,
        string PaymentAddress,
        string ComputerAttributesAddress
    );
    
    event AskTraining (
        string DataSchemaAddress,
        string MetadataAddress,
        string ModelAddress,
        string StrategyAddress,
        string ComputionAddress
    );
    
    event AskForData (
        string KeyAddress,
        string DataSchemaAddress,
        string TransactionDetailAddress
    );
    
    
}