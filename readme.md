# netstaksim

A Go program using BFS to simulate the most basic possible network I can think of, modelled as a simple undirected graph.

There's no IP address, addressing is just integers - each host has an address; and of course every host has to have a unique address (enforced by the hashmap in the AddNode method on the graph - but it may be fun to break that!).

A message (analogous to a packet) has a source, destination and content.

The way I deliver a message from one host to another is just BFS, so we get shortest path by default.

I also added a "trace route" function which just accumulates the path that the message took along the path, not the search, and output it :)

Go Playground link: <https://goplay.tools/snippet/XOTKWKIew4F>

```go
    bfsDeliverMessage(msg, sourceHost)
    // at node {2}
    // at node {1}
    // at node {3}
    // at node {4}
    // at node {5}
    // at node {6}
    // at node {7}
    // at node {8}
    // [8] receiveved message from [2]: Hello, World!
    // Path taken aka route trace: 2 1 8 
```
