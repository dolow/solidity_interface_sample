// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

import './SampleInterface.sol';

contract SampleVer1 is SampleInterface {
  uint256 i = 0;

  function setValue(uint v) external {
    i = v;
  }

  function getValue() external view returns (uint256) {
    return i;
  }
}
