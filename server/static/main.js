//animation
var currentAnime = null;
var isMetamaskInstalled = false;
var singleMessage = null;
var contract = null;

$(document).ready(function() {
  restartAnimation();

  if (typeof web3 !== 'undefined') {
    web3 = new Web3(web3.currentProvider);
    isMetamaskInstalled = true;
  }

  singleMessage = web3.eth.contract(singleMessageABI);
  contract = singleMessage.at("0x15d3122103c5c17ed791fd5a3dba847ecfd6037e");
  window.setInterval(autoUpdate, 5000);
  document.getElementById("purchase").addEventListener("click", onClickPurchase);  
});

function restartAnimation() {
  $('.ml9 .letters').each(function () {
    $(this).html($(this).text().replace(/([^\x00-\x80]|\w)/g, "<span class='letter'>$&</span>"));
  });

  currentAnime = anime.timeline({ loop: true })
    .add({
      targets: '.ml9 .letter',
      scale: [0, 1],
      duration: 1500,
      elasticity: 600,
      delay: function (el, i) {
        return 45 * (i + 1)
      }
    }).add({
      targets: '.ml9',
      opacity: 0,
      duration: 1000,
      easing: "easeOutExpo",
      delay: 1000
    });
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

function autoUpdate() {
  getPrice(function(error, price) {
    if (error) {
      return handleErrorSilent(error);
    }

    getMessage(function(error, message) {
      if (error) {
        return handleErrorSilent(error);
      }

      var currentMessage = document.getElementById("message").innerText;
      if (currentMessage !== message) {
        document.getElementById("message").innerText = message;        
        // Wrap each letter in a span again
        $('.ml9 .letters').each(function () {
          $(this).html($(this).text().replace(/([^\x00-\x80]|\w)/g, "<span class='letter'>$&</span>"));
        });
        if (currentAnime) {
          currentAnime.pause();
        }
        restartAnimation();
      }
      document.getElementById("price").innerText = web3.fromWei(price, 'ether');
    });
  });
}

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
  if (!isMetamaskInstalled) {
    alert("You need Metamask to become the Biggest G, click the link below to get it.")
  }
  getBalance(function (error, balance) {
    if (error) {
      return handleError(error);
    }
    getPrice(function(error, price) {
      if (error) {
        return handleError(error);
      }
      
      if (balance < price) {
        return alert("Sorry, you're not ballin' enough to become the Biggest G (need more Ether)");
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

function handleError(error) {
  console.log(error);
  alert("Sorry, something went wrong!");
}

function handleErrorSilent(error) {
  console.log(error);
}