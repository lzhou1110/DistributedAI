pragma solidity ^0.4.0;

contract bidcomputing{

    address public seller;

    uint256 public Payment;

    uint public status;

    mapping(address => bool) lock;
    mapping(address => bool) relock;

    enum State { Created, Inactive }
    State public state;

    function bidcomputing(
        string OperationSchemasAddress,
        string ComputingAddress,
        string ComputerAttributesAddress,
        uint _Payment
    ){
        seller = msg.sender;
        state = State.Created;
        Payment = _Payment;

        Bidcomputing(
            OperationSchemasAddress,
            ComputingAddress,
            ComputerAttributesAddress,
            _Payment,
            this
        );
    }
    //inState(State.Created)
    function askTraining(string DataSchemaAddress,string MetadataAddress,string ModelAddress,string StrategyAddress,string ComputionAddress)
         payable returns(uint)
    {
        lock[msg.sender] = true;

        // don't need event ,computor use buyer's address to comfirm

       // AskTraining(DataSchemaAddress,MetadataAddress,ModelAddress,StrategyAddress,ComputionAddress);
        status=200;
        return status;
    }

    function isbaid(address buyer) returns(bool){
        return lock[buyer];
    }

    function uploadTrainRes (
        string ResultAddress,
        string TransactionDetailAddress
    )
        returns(uint256)
    {
        UploadTrainRes (
            ResultAddress,
            TransactionDetailAddress
        );
        return 200;
    }

    function confirmresult()
    {
        relock[msg.sender] = true;
        seller.transfer(Payment);
    }


    function abort()
        onlySeller
        inState(State.Created)
    {
        Aborted();
        state = State.Inactive;
        seller.transfer(this.balance);
    }

    function() payable public {
        FallbackCalled(msg.sender, msg.value);
    }

    modifier onlySeller() {
        require(msg.sender == seller);
        _;
    }

    modifier condition(bool _condition) {
        require(_condition);
        _;
    }


    modifier inState(State _state) {
        require(state == _state);
        _;
    }

    modifier baid(address payer){
        require(isbaid(payer));
        _;
    }

    event FallbackCalled(address from, uint256 amount);

    event Bidcomputing(
        string OperationSchemasAddress,
        string ComputingAddress,
        string ComputerAttributesAddress,
        uint256 _Payment,
        address _this
    );

    event UploadTrainRes (
        string ResultAddress,
        string TransactionDetailAddress
    );


    event AskTraining (
        string DataSchemaAddress,
        string MetadataAddress,
        string ModelAddress,
        string StrategyAddress,
        string ComputionAddress
    );

    event Aborted();


}
