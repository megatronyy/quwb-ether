pragma solidity ^0.4.0;

contract SimpleStorage {
    uint storeData;
    address myAddress = this;

    function set(uint x){
        storeData = x;
    }

    function get() constant returns (uint retVal){
        return storeData;
    }
}
