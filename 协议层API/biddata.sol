pragma solidity ^0.4.0;

contract biddata{

    address public seller;

    uint public Payment;

    bool public status;

    mapping(address => bool) lock;

    enum State { Created, Inactive }
    State public state;

    function biddata(
        string DataSchemaAddress,
        string MataDataAddress,
        uint256 _Payment
    ){
        seller = msg.sender;
        state = State.Created;
        Payment = _Payment;

        BidData(
            DataSchemaAddress,
            MataDataAddress,
            Payment,
            this
        );
    }
    //inState(State.Created) enoughmoney
    function buydata()  payable returns(bool)
    {
        lock[msg.sender] = true;

        // don't need event ,computor use buyer's address to comfirm
        status=true;
        return status;
    }


    function isbaid(address buyer) returns(bool){
        return lock[buyer];
    }

    // wait for discussing
    function askfordata(
        string KeyAddress,
        string DataSchemaAddress,
        string computionAddress,
        string TransactionDetailAddress
    )
        // baid(address payer)
        returns(uint256)
    {
        return 200;
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

    event FallbackCalled(address from, uint256 amount);

    modifier onlySeller() {
        require(msg.sender == seller);
        _;
    }

    modifier condition(bool _condition) {
        require(_condition);
        _;
    }

    modifier enoughmoney(){
        require(msg.value >= Payment);
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

    event BidData(
        string DataSchemaAddress,
        string MataDataAddress,
        uint256 Payment,
        address _this
    );

    event Aborted();

}
