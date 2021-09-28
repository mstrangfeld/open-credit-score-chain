# Smart Contracts

## Authority

The Authority smart contract is deployed by the government or any institution that has the authority to allow entities to create records in users credit score database.
It can create new credit score databases for an user with a given tax ID.
It can also allow entities to be able to write to user's credit score database.

## CreditScoreDB

The credit score database is created for each user.
It holds the records of to calculate the credit score.
Entities with the permission from the Authority smart contract can create new records.
Other entities with the permission from the authority can invalidate records.
