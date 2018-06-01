##setup
sendbox
```
$ ip route add 10.0.11.0/24 via 10.0.10.20 dev eth0
```

recvbox
```
$ ip route add 10.0.10.0/24 via 10.0.11.10 dev eth0
```
