pragma solidity ^0.4.17;

contract CryptoContract{

    mapping (address => uint256) public balanceOf;
    string _tokenName;

    function CryptoContract(string _tokenName) public {
        _tokenName = _tokenName;
    }

    function Transfer(address _recipeintAddress, uint256 value) payable public returns (uint256) {
            balanceOf[_recipeintAddress] += msg.value;
            return msg.sender.balance;
    }
    
    function moneyTrnasfer(address _recipeintAddress, uint256 value) public {
            _recipeintAddress.transfer(value);  
            balanceOf[_recipeintAddress] -= value;
    }
    
    function getBalanceOfAccount(address _recipeintAddress) public constant returns (uint256) {
        return balanceOf[_recipeintAddress];
    }
    
    function getBalance() public constant returns (uint256) {
        return address(this).balance;
    }
    
}