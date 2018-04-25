pragma solidity ^0.4.0;
/*
    1、访问权限 public，internal，pravate，external
    2、属性默认权限为inernal，只有public类型的属性才可能供外部访问。
       internal和pravate类型的属性只能在合约内部使用。
    3、函数的权限默认是public类型，public类型的函数可供外表访问，
       而internal和pravate类型的函数不能通过指针（this表示当前智能合约的指针）访问。
       哪怕在内容通过this访问都会报错。在合约内部访问，可以直接访问函数。
*/
contract Coin {
    //关键字“public”使变量能从合约外部访问
    address public minter;
    mapping (address => uint) public balances;

    //事件让轻客户端能高效的对变化做出反应
    event Sent(address from, address to, int amount);

    //这个构造函数的代码仅仅只在合约创建的时候被运行
    function Coin(){
        minter = msg.sender;
    }
    //铸币
    function mint(address receiver, uint amount) {
        if (msg.sender != minter) return;
        balances[receiver] += amount;
    }
    //发送货币
    function send(address receiver, uint amount) {
        if (balances[msg.sender] < amount) return;
        balances[msg.sender] -= amount;
        balances[receiver] += amount;
        Sent(msg.sender, receiver, amount);
    }
}
