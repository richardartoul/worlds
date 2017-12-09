module.exports = function(deployer, network, accounts) {
  // Use deployer to state migration tasks.
};

async function deploy(deployer, accounts) {
  const address = accounts[0];
  await deployer.deploy(SingleMessage, "Hello world!", 1, address);
  contract = await SingleMessage.deployed();
  console.log("Deployed successfully!");
}
