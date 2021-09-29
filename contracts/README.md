# Smart Contracts

## Authority

The Authority smart contract is deployed by the government or any institution that has the authority to allow entities to create records in users credit score database.
It can create new credit score databases for an user with a given tax ID.
It can also give entities the right to be able to write to a user's credit score database.

## CreditScoreDB

The credit score database is created for each user.
It holds the records of the user to calculate the credit score.
Entities with the permission from the Authority can create new records or invalidate old ones in case of a mistake or error.

The records contain following fields:
| Name | Description | Encrypted |
| ------ | ------------- | -------- |
| Valid | Whether the record is still valid. A user can ask the authority to invalidate a record if they believe there has been a mistake or an error | No |
| Issuer | The issuer of this record. The address needs to be authorized by the authority | No |
| Timestamp | The time the record has been created | No |
| Reason | The reason for this record. This should be a very short probably standardized message | Yes |
| Score | The score of this record. Should be a standardized range e.g. 0-100 | Yes |
| Nonce | The nonce to create the score hash. Should be a random string of reasonable size | Yes |
| ScoreHash | The MiMC hash calculated by adding the unencrypted score to the unencrypted nonce | No |
