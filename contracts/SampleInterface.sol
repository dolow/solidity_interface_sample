// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

interface SampleInterface {
  function setValue(uint v) external;
  function getValue() external view returns (uint256);
}
