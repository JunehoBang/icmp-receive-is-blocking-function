# icmp-receive-is-blocking-function

This repository was created to test whether the `icmp.ReadFrom()` function is 
indeed a blocking function.

As can be tested by the code, the function is blocking function.
To make the code to arrive at the termination, the local machine has to receive an icmp message
so that the blocking state is cleared
