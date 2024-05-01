# netstaksim

A Go program using BFS to simulate the most basic possible network I can think of, modelled as a simple undirected graph.

There's no IP address, addressing is just integers - each host has an address; and of course every host has to have a unique address (enforced by the hashmap in the AddNode method on the graph - but it may be fun to break that!).

A message (analogous to a packet) has a source, destination and content.

The way I deliver a message from one host to another is just BFS, so we get shortest path by default.

I also added a "trace route" function which just accumulates the path that the message took along the path, not the search, and output it :)
Go Playground link: <https://goplay.tools/snippet/XOTKWKIew4F>
