// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.4.16 <0.9.0;

/**
 * @title Authority
 * @dev The authority contract is used to manage the mapping of taxIDs to their corresponding credit score databases.
 * It is also used to manage the entities that are allowed to write to the credit score database.
 */
contract Authority {
    address public root; // The root of trust for this authority

    mapping(string => address) private creditDbs; // A mapping of taxIDs to their corresponding credit score database
    mapping(address => bool) public recordWriters; // A mapping of all addresses that are allowed to create records in the user's credit score database

    /**
     * @dev Constructor
     */
    constructor() {
        require(msg.sender != address(0), "Root address cannot be the zero address");
        root = msg.sender;
    }

    /**
     * @param _taxID The tax ID of the user
     * @return The credit score database of the user with the given taxID
     */
    function getDB(string calldata _taxID) public view returns (address) {
        return creditDbs[_taxID];
    }

    /**
     * @dev Creates a new credit score database for the user with the given taxID
     * @param _taxID The tax ID of the user
     * @param _addr The address of the credit score database smart contract
     */
    function createDB(string calldata _taxID, address _addr) public {
        creditDbs[_taxID] = _addr;
    }

    /**
     * @dev Authorize the given address to create records in the credit score database of the users
     * @param _addr The address to authorize
     */
    function authorizeRecordWriter(address _addr) public {
        recordWriters[_addr] = true;
    }

    /**
     * @dev Unauthorize the given address to create records in the credit score database of the users
     * @param _addr The address to unauthorize
     */
    function unauthorizeRecordWriter(address _addr) public {
        recordWriters[_addr] = false;
    }

    /**
     * @dev Check if the given address is authorized to create records in the credit score database of the users
     * @param _addr The address to check
     */
    function isAllowedToCreateRecords(address _addr) public view returns (bool) {
        return recordWriters[_addr];
    }
}
