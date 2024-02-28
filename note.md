a wallet is made up of two keys, the `private` and `public` key.

the private key itself is the identifier for each of the accounts inside of the blockchain, so each provate key we generate needs to be completely random and unique
from the other priate keys, otherwise multiple users can access the same account.

the public key on the other hand is a key that we can actually give to other users and we use this public key to derive the address which is the address that we used to send
and receive data in the blockchain.

we derive the public key from the private key, and then we derive the address from the public key.


`ecdsa` stands for eliptical curve digital signing algorithm



we take the public key and run it through `sha256` hash and we run the result into `ripemd160` hashing algorithm to get the value called the `public key hash`,
we then take the `public key hash` and pass it through `sha256` twice, take of the 4 bytes of the resulting hash and then we get the checksum value.

we also take the original `public key hash` and pass it down, and we also have a `version` value, we concatinate these 3 values and pass it to `base58` encoder to create the `address`

the checksum is important for verifying and signing transactions, the version value indicates where in our blockchain this address resides.
