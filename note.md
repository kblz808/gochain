transactions are made up of inputs and outputs.

since a blockchain is open in public database, we dont want to store sensetive information
inside of the blockchain.

there are no accounts, balance, address, coin, senders and receivers.
we only have inputs and outputs, and then we can derive all the other things from the inputs and outputs.


## outputs
each output has a value assigned to it, which is the value in tokens, which is signed and locked inside of the output.

the `PubKey` is the value needed to unlock the token inside of the `Value` field.

on bitcoin, the `PubKey` is derived through a complicated scripting language called `script`

for now we're just going to use the user address like `jack`, then the `PubKey` field will be `jack`, and that will allow the user to unlock the tokens insid the output.

outputs are indivisible, meaning we cannot reference a part of an output, eg:- if we give $10 to a $5 product, the cashier cant just rip the $10 in half and give us the other half back,
instead they have to give you back a $5 bill.

so if there are 10 tokens inside of our output, we need to create new outputs, one with 5 tokens inside of it and another one with 5 tokens inside of it.


## inputs
are just references to previous outputs.

the `ID` field references the transaction that the output is inside of.

we also need to know the index where this output appears using the `Out` field.

the `Sig` field is similar to the `PubKey` field in `TXOutput`, the `Sig` is a script which provides the data which is used in the outputs pub key.
for now the field is just going to be the users account, eg:- if we have an account named `jack`, then we're gonna push it into the `Sig` as well as the `PubKey`.
so the `jack` can then unlock the output which is being referenced by the input.


## coinbase
in our chain, we have our genesis block, in that block we also have the first transaction, this is whats called a `coinbase` transaction.

in this transaction, we only have one input and one output, and the input inside of it rather than referencing an older output, just references an empty output.
if doesnt also store a signiture, instead it stores a bunch of data.

the `coinbase` also has whats called a `subsidy`, which is a reward attached to it.
this reward is released to a single account when that individual `mines` the `coinbase`.

so when we create our blockchain, if an acount named `jack` mines this block, then `jack` will get 100 tokens.


## -
if `CanUnlock` and `CanBeUnlocked` methods come back as `true`, then that means that the account which is the `data`, owns the information inside of the output, or it owns the information
inside of the output which is referenced by the input.


we now need to add transactions in our block instead of the `Data` field.

> note that each block needs to have atleast one transaction inside of it, and each block can have many different transactions inside of it as well


unspent transactions are ones that have outputs which are not referenced by othe inputs.

if an output hasn't been spent, that means that those tokens still exist for certain user.
so by counting all the unspent outputs that are assigned to a certain user, we can find out how many tokens are assigned to that user.


## test

gochain createblockchain -address alex

gochain printchain

gochain getbalance -address alex

gochain send -from alex -to skipper -amount 50
