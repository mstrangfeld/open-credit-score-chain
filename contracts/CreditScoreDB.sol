// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.4.16 <0.9.0;

/**
 * @title CreditScoreDB
 * @dev Database for storing credit scores
 */
contract CreditScoreDB {
    address public owner; // Address of the owner of this database
    address public authority; // Address of the authority smart contract

    /**
     * @dev Represents an entry to the database
     */
    struct RecordEntry {
        bool valid; // Whether the entry is valid
        address issuer; // The issuer of the credit score entry
        uint256 timestamp; // Unix timestamp
        string reason; // Reason for the credit score entry (encrypted)
        string score; // Credit score (encrypted)
        string nonce; // Nonce for the generation of the scoreHash (encrypted)
        string scoreHash; // Hash of the credit score
    }

    RecordEntry[] public records; // Array of credit score entries

    /**
     * @dev Constructor
     * @param _owner Address of the owner of this database
     */
    constructor(address _owner) {
        owner = _owner;
        authority = msg.sender;
    }

    /**
     * @dev Adds a credit score entry to the database
     * @param _reason The reason for the credit score entry (encrypted)
     * @param _score The credit score (encrypted)
     */
    function createRecord(string calldata _reason, string calldata _score, string calldata _nonce, string calldata _scoreHash) public {
        // TODO: Call authority to check if the caller is authorized to add a record
        // TODO: Check if the score and the reason have been encrypted with the correct public key to avaoid unreadable data in the database
        records.push(RecordEntry(true, msg.sender, block.timestamp, _reason, _score, _nonce, _scoreHash));
    }

    /**
     * @dev Invalidates a credit score entry
     * @param _index Index of the credit score entry to invalidate
     */
    function invalidateRecord(uint256 _index) public {
        // TODO: Call authority to check if the caller is authorized to invalidate a record
        records[_index].valid = false;
    }
}
