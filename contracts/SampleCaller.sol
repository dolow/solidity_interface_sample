// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

import './SampleInterface.sol';

contract SampleCaller {
  address sampleAddress = address(0x0);

  function setSampleAddress(address a) public {
    sampleAddress = a;
  }

  function setSampleValue(uint256 i) public {
    SampleInterface sample = SampleInterface(sampleAddress);
    sample.setValue(i);
  }

  function getSampleValue() public view returns (uint256) {
    SampleInterface sample = SampleInterface(sampleAddress);
    return sample.getValue();
  }
}
