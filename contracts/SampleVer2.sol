// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

import './SampleInterface.sol';

contract SampleVer2 is SampleInterface {
  uint256 i = 0;

  function setValue(uint v) external {
    setValueExtention(v);
  }

  function getValue() external view returns (uint256) {
    return getValueExtention();
  }

  function setValueExtention(uint v) private {
    i = v * 2;
  }
  function getValueExtention() private view returns (uint256) {
    return i * 4;
  }
}
