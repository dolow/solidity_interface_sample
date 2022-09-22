const SampleCaller = artifacts.require("SampleCaller");

module.exports = function (deployer) {
  deployer.deploy(SampleCaller);
};
