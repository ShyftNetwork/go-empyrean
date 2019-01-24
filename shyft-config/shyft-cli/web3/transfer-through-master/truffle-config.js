module.exports = {
  // See <http://truffleframework.com/docs/advanced/configuration>
  // to customize your Truffle configuration!
    solc: {
        optimizer: {
            enabled: true,
            runs: 200
        }
    },
    networks: {
    development: {
      host: "localhost",
      port: 8545,
      netword_id: "*",
      gas: 8000000
    }
  },
};
