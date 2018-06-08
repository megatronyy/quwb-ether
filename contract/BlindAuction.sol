pragma solidity ^0.4.21;
/**
秘密竞拍（盲拍）
之前的公开拍卖接下来将被扩展为一个秘密竞拍。 秘密竞拍的好处是在投标结束前不会有时间压力。
在一个透明的计算平台上进行秘密竞拍听起来像是自相矛盾，但密码学可以实现它。
在投标期间 ，投标人实际上并没有发送她的出价，而只是发送一个哈希版本的出价。
由于目前几乎不可能找到两个（足够长的）值，其哈希值是相等的，因此投标人可通过该方式提交报价。 在投标结束后，投标人必须公开他们的出价：他们不加密的发送他们的出价，合约检查出价的哈希值是否与投标期间提供的相同。

另一个挑战是如何使拍卖同时做到 绑定和秘密 : 唯一能阻止投标者在她赢得拍卖后不付款的方式是，
让她将钱连同出价一起发出。 但由于资金转移在 以太坊Ethereum 中不能被隐藏，
因此任何人都可以看到转移的资金。

下面的合约通过接受任何大于最高出价的值来解决这个问题。 当然，因为这只能在披露阶段进行检查，
有些出价可能是 无效 的， 并且，这是故意的(与高出价一起，
它甚至提供了一个明确的标志来标识无效的出价): 投标人可以通过设置几个或高或低的无效出价来
迷惑竞争对手。
 **/
contract BlindAuction {
    struct Bid {
        bytes32 blindedBid;
        uint deposit;
    }

    address public beneficiary;
    uint public biddingEnd;
    uint public revealEnd;
    bool public ended;

    mapping(address => Bid[]) public bids;

    address public highestBidder;
    uint public highestBid;

    // 可以取回的之前的出价
    mapping(address => uint) pendingReturns;

    event AuctionEnded(address winner, uint highestBid);

    /// 使用 modifier 可以更便捷的校验函数的入参。
    /// `onlyBefore` 会被用于后面的 `bid` 函数：
    /// 新的函数体是由 modifier 本身的函数体，并用原函数体替换 `_;` 语句来组成的
    modifier onlyBefore(uint _time){
        require(now < _time);
        _;
    }
    modifier onlyAfter(uint _time){
        require(now>_time);
        _;
    }

    function BlindAuction(){

    }
}
