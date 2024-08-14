// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Burnable.sol";

contract FarawayCollection is ERC721, ERC721URIStorage, ERC721Burnable {
    event TokenMinted(address collection, address recipient, uint256 tokenId, string tokenUri);

    uint256 private _nextTokenId;

    constructor(string memory _name, string memory _symbol) ERC721(_name, _symbol) {}

    function _safeMint(address _to, string memory _uri) internal {
        uint256 tokenId = _nextTokenId++;
        _safeMint(_to, tokenId);
        _setTokenURI(tokenId, _uri);
        emit TokenMinted(address(this), _to, tokenId, _uri);
    }

    function mint(string calldata _uri) external {
        _safeMint(msg.sender, _uri);
    }

    function tokenURI(uint256 tokenId)
        public
        view
        override(ERC721, ERC721URIStorage)
        returns (string memory)
    {
        return super.tokenURI(tokenId);
    }

    function supportsInterface(bytes4 interfaceId)
        public
        view
        override(ERC721, ERC721URIStorage)
        returns (bool)
    {
        return super.supportsInterface(interfaceId);
    }
}