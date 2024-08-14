// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { FarawayCollection } from "./FarawayCollection.sol";

contract FarawayCollectionFactory {
    event CollectionCreated(address collection, string name, string symbol);

    function createCollection(string calldata _name, string calldata _symbol) external {
        FarawayCollection collection = new FarawayCollection(_name, _symbol);
        emit CollectionCreated(address(collection), _name, _symbol);
    }
}