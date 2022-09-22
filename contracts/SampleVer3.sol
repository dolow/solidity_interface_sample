// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

import './SampleInterface.sol';

contract SampleVer3 is SampleInterface {
  uint256 i = 0;

  function setValue(uint v) external {
    setValueExtention(v + 1);
  }

  function getValue() external view returns (uint256) {
    return getValueExtention() + 1;
  }

  function setValueExtention(uint v) private {
    i = v * 2;
  }
  function getValueExtention() private view returns (uint256) {
    return i * 4;
  }
}
