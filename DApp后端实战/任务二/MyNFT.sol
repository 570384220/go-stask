// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

// ✅ 使用 OpenZeppelin v5 导入路径
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract MyNFT is ERC721URIStorage, Ownable {
    uint256 private _nextTokenId = 1;

    // ✅ 构造函数需要显式传入 initialOwner
    constructor(address initialOwner)
        ERC721("My Image NFT", "MINFTAA")
        Ownable(initialOwner) // ✅ 必须加这一行
    {}

    /// @notice 铸造 NFT
    /// @param recipient 接收者地址
    /// @param tokenURI 元数据 JSON 文件的 IPFS 链接
    function mintNFT(address recipient, string memory tokenURI)
        public
        onlyOwner
        returns (uint256)
    {
        uint256 tokenId = _nextTokenId++;
        _safeMint(recipient, tokenId);
        _setTokenURI(tokenId, tokenURI);
        return tokenId;
    }
}
