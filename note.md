bitcoin and others use a databse called leveldb, which is a low level key value storage.

we need to consider how we want to store our blockchain data.

in the bitcoin core specification, there are two main groups of data, we have the blocks which are stored with metada
which describe all of the blocks on the chain and the chain state object which stores the state of a chain and all of the
current unspent transaction outputs and as well as few pieces of metadata.

with bitcoin, each block has its own separate file on the disk.


we used iterators to iterate through the blockchain backwards, starting from the latest block
