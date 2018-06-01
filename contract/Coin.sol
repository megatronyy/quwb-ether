pragma solidity ^0.4.21;
/*
    1、访问权限 public，internal，pravate，external
    2、属性默认权限为inernal，只有public类型的属性才可能供外部访问。
       internal和pravate类型的属性只能在合约内部使用。
    3、函数的权限默认是public类型，public类型的函数可供外表访问，
       而internal和pravate类型的函数不能通过指针（this表示当前智能合约的指针）访问。
       哪怕在内容通过this访问都会报错。在合约内部访问，可以直接访问函数。
    4、function (<parameter types>) {internal|external} [pure|constant|view|payable] [returns (<return types>)]
*/
contract Coin {
    // 关键字“public”让这些变量可以从外部读取
    address public minter;
    mapping (address => uint) public balances;

    // 轻客户端可以通过事件针对变化作出高效的反应
    event Sent(address from, address to, uint amount);

    // 这是构造函数，只有当合约创建时运行
    function Coin() public {
        minter = msg.sender;
    }

    function mint(address receiver, uint amount) public {
        if (msg.sender != minter) return;
        balances[receiver] += amount;
    }

    function send(address receiver, uint amount) public {
        if (balances[msg.sender] < amount) return;
        balances[msg.sender] -= amount;
        balances[receiver] += amount;
        emit Sent(msg.sender, receiver, amount);
    }
}
