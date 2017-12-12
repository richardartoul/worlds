if (typeof web3 !== 'undefined') {
  web3 = new Web3(web3.currentProvider);
} else {
  // set the provider you want from Web3.providers
  web3 = new Web3(new Web3.providers.HttpProvider("http://localhost:8545"));
}

var singleMessageABI = [
  {
    "constant": true,
    "inputs": [],
    "name": "priceInWei",
    "outputs": [
      {
        "name": "",
        "type": "uint256"
      }
    ],
    "payable": false,
    "stateMutability": "view",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "newMessage",
        "type": "string"
      }
    ],
    "name": "set",
    "outputs": [],
    "payable": true,
    "stateMutability": "payable",
    "type": "function"
  },
  {
    "constant": true,
    "inputs": [],
    "name": "owner",
    "outputs": [
      {
        "name": "",
        "type": "address"
      }
    ],
    "payable": false,
    "stateMutability": "view",
    "type": "function"
  },
  {
    "constant": true,
    "inputs": [],
    "name": "maxLength",
    "outputs": [
      {
        "name": "",
        "type": "uint256"
      }
    ],
    "payable": false,
    "stateMutability": "view",
    "type": "function"
  },
  {
    "constant": true,
    "inputs": [],
    "name": "message",
    "outputs": [
      {
        "name": "",
        "type": "string"
      }
    ],
    "payable": false,
    "stateMutability": "view",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "newOwner",
        "type": "address"
      }
    ],
    "name": "transferOwnership",
    "outputs": [],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "constant": false,
    "inputs": [
      {
        "name": "destination",
        "type": "address"
      },
      {
        "name": "amountInWei",
        "type": "uint256"
      }
    ],
    "name": "withdraw",
    "outputs": [],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "function"
  },
  {
    "inputs": [
      {
        "name": "initialMessage",
        "type": "string"
      },
      {
        "name": "initialPriceInWei",
        "type": "uint256"
      },
      {
        "name": "maxLengthArg",
        "type": "uint256"
      }
    ],
    "payable": false,
    "stateMutability": "nonpayable",
    "type": "constructor"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": false,
        "name": "message",
        "type": "string"
      },
      {
        "indexed": false,
        "name": "priceInWei",
        "type": "uint256"
      },
      {
        "indexed": false,
        "name": "newPriceInWei",
        "type": "uint256"
      },
      {
        "indexed": false,
        "name": "payer",
        "type": "address"
      }
    ],
    "name": "MessageSet",
    "type": "event"
  },
  {
    "anonymous": false,
    "inputs": [
      {
        "indexed": true,
        "name": "previousOwner",
        "type": "address"
      },
      {
        "indexed": true,
        "name": "newOwner",
        "type": "address"
      }
    ],
    "name": "OwnershipTransferred",
    "type": "event"
  }
];

var singleMessage = web3.eth.contract(singleMessageABI);
var contract = singleMessage.at("0x6c273582e2f2f34dc0442eb00883397f18cb4a2e");

function autoUpdate() {
  getPrice(function(error, price) {
    if (error) {
      return handleErrorSilent(error);
    }

    getMessage(function(error, message) {
      if (error) {
        return handleErrorSilent(error);
      }

      document.getElementById("message").innerText = message;
      document.getElementById("price").innerText = price;
    });
  })
}

window.setInterval(autoUpdate, 5000);

function getBalance(cb) {
  web3.eth.getBalance(web3.eth.coinbase, function(error, balance) {
    if (error) {
      return cb(error, null);
    }
    return cb(null, balance.toNumber());
  });
}

function getPrice(cb) {
  contract.priceInWei(function(error, price) {
    if (error) {
      return cb(error, null);
    }

    cb(null, price.toNumber());
  });
}

function getMessage(cb) {
  contract.message(function(error, message) {
    return cb(error, message);
  });
}

function purchase(newMessage, price, cb) {
  contract.set(newMessage, {value: price}, function(error, message) {
    return cb(error, message)
  });
}

function onClickPurchase(cb) {
  getBalance(function (error, balance) {
    if (error) {
      return handleError(error);
    }
    getPrice(function(error, price) {
      if (error) {
        return handleError(error);
      }
      
      if (balance < price) {
        return alert("Sorry peasant, you can't afford to update this message!");
      }

      var newMessage = window.prompt("What would you like to change the message to?");
      if (!newMessage) {
        return alert("Sorry, you have to type in something!");
      }

      purchase(newMessage, price, function(error, message) {
        if (error) {
          return handleError(error);
        }

        alert("Success! If your transaction succeeds the page will update shortly!");
      });
    });
  });
}

document.getElementById("purchase").addEventListener("click", onClickPurchase);

function handleError(error) {
  console.log(error);
  alert("Sorry, something went wrong!");
}

function handleErrorSilent(error) {
  console.log(error);
}