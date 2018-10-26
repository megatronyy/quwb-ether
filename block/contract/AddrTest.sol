pragma solidity ^0.4.0;

contract AddrTest {
    event logdata(bytes data);

    function() payable {
        logdata(msg.data);
    }
    function AddrTest(){

    }
}
