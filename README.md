
### Design and implement “Word of Wisdom” tcp server.

• TCP server should be protected from DDOS attacks with the Proof of Work (https://en.wikipedia.org/wiki/Proof_of_work),
the challenge-response protocol should be used.

• The choice of the POW algorithm should be explained.

• After Proof Of Work verification, server should send one of the quotes from “word of wisdom” book or any other
collection of the quotes.

• Docker file should be provided both for the server and for the client that solves the POW challenge

### Run the server and client separately

```
 make run-server
```

then in another terminal, run the following command to start the client

```
 make run-client
```

clean up the server and client, stop and remove the docker images

```
 make clean
```