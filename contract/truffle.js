module.exports = {
  networks: {
    development: {
      host: "localhost",
      port: 8545,
      network_id: "*",
      gas: 4700000
    },
    test: {
      host: "localhost",
      port: 9545,
      network_id: "*"
    }
  }
};
