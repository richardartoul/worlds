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
    },
    production: {
      host: "localhost",
      port: 8543,
      network_id: "1",
      gas: 4700000,
      from: "0x52ce8b05cc449f9bbb6b95c8b501e3d294869372"
    },
    infuraProduction: {
      host: "https://mainnet.infura.io/<REDACTED_ACCESS_TOKEN>",
      network_id: "1",
      gas: 4700000,
      from: "0x52ce8b05cc449f9bbb6b95c8b501e3d294869372"
    }
  }
};
