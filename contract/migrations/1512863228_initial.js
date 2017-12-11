var SingleMessage = artifacts.require('../contracts/SingleMessage.sol');

module.exports = function(deployer, network, accounts) {
  return liveDeploy(deployer, accounts);
};

async function liveDeploy(deployer, accounts) {
  const address = accounts[0];
  const initialMessage = "Hello world!";
  const initialPriceInWei = 1;
  const maxLength = 200;

  console.log("Contract arguments: ", {
    initialMessage: initialMessage,
    initialPriceInWei: initialPriceInWei,
    maxLength: maxLength,
    address: address
  });

  return deployer.deploy(SingleMessage, initialMessage, initialPriceInWei, maxLength, address).then(async() => {
    const contract = await SingleMessage.deployed();
    const message = await contract.message.call();
    const priceInWei = await contract.priceInWei.call();
    const maxLength = await contract.maxLength.call();

    console.log("public vars from contract: ", {
      message: message,
      priceInWei: priceInWei,
      maxLength: maxLength
    });
  });
}
