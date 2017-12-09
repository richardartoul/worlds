pragma solidity ^0.4.17;

import './Ownable.sol';

contract SingleMessage is Ownable {
  string public message;
  uint256 public priceInWei;

  function SingleMessage(string initialMessage, uint256 initialPriceInWei) public {
    message = initialMessage;
    priceInWei = initialPriceInWei;
  }

  function set(string newMessage) external payable {
    require(msg.value >= priceInWei);
    message = newMessage;
    priceInWei = priceInWei * 2;
  }

  function withdraw(address destination, uint256 amountInWei) external onlyOwner {
    require(this.balance >= amountInWei);
    destination.transfer(amountInWei);
  }
}