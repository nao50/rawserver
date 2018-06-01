## setup
sendbox
```
$ ip route add 10.0.11.0/24 via 10.0.10.20 dev eth0
```

recvbox
```
$ ip route add 10.0.10.0/24 via 10.0.11.10 dev eth0
```
## analyse
strace
```
$ go build rawserver.go && strace -e 'trace=!pselect6,futex,sched_yield' ./rawserver
```

result
```
openat(AT_FDCWD, "/proc/sys/net/core/somaxconn", O_RDONLY|O_CLOEXEC) = 3
epoll_create1(EPOLL_CLOEXEC)            = 4
epoll_ctl(4, EPOLL_CTL_ADD, 3, {EPOLLIN|EPOLLOUT|EPOLLRDHUP|EPOLLET, {u32=2174217984, u64=140245741334272}}) = 0
fcntl(3, F_GETFL)                       = 0x8000 (flags O_RDONLY|O_LARGEFILE)
fcntl(3, F_SETFL, O_RDONLY|O_NONBLOCK|O_LARGEFILE) = 0
read(3, "128\n", 65536)                 = 4
read(3, "", 65532)                      = 0
epoll_ctl(4, EPOLL_CTL_DEL, 3, 0xc42004fd4c) = 0
close(3)                                = 0
write(1, "\n===== syscall.Socket() =====\n", 30
===== syscall.Socket() =====
) = 30
socket(AF_PACKET, SOCK_DGRAM, 768)      = 3
socket(AF_NETLINK, SOCK_RAW, NETLINK_ROUTE) = 5
bind(5, {sa_family=AF_NETLINK, nl_pid=0, nl_groups=00000000}, 12) = 0
sendto(5, {{len=17, type=0x12 /* NLMSG_??? */, flags=NLM_F_REQUEST|0x300, seq=1, pid=0}, "\0"}, 17, 0, {sa_family=AF_NETLINK, nl_pid=0, nl_groups=00000000}, 12) = 17
recvfrom(5, [{{len=1180, type=0x10 /* NLMSG_??? */, flags=NLM_F_MULTI, seq=1, pid=72}, "\0\0\4\3\1\0\0\0I\0\1\0\0\0\0\0\7\0\3\0lo\0\0\10\0\r\0\1\0\0\0"...}, {{len=1216, type=0x10 /* NLMSG_??? */, flags=NLM_F_MULTI, seq=1, pid=72}, "\0\0\1\0\214\1\0\0C\20\1\0\0\0\0\0\t\0\3\0eth1\0\0\0\0\10\0\r\0"...}, {{len=1216, type=0x10 /* NLMSG_??? */, flags=NLM_F_MULTI, seq=1, pid=72}, "\0\0\1\0\216\1\0\0C\20\1\0\0\0\0\0\t\0\3\0eth0\0\0\0\0\10\0\r\0"...}], 4096, 0, {sa_family=AF_NETLINK, nl_pid=0, nl_groups=00000000}, [112->12]) = 3612
getsockname(5, {sa_family=AF_NETLINK, nl_pid=72, nl_groups=00000000}, [112->12]) = 0
getsockname(5, {sa_family=AF_NETLINK, nl_pid=72, nl_groups=00000000}, [112->12]) = 0
getsockname(5, {sa_family=AF_NETLINK, nl_pid=72, nl_groups=00000000}, [112->12]) = 0
recvfrom(5, {{len=20, type=NLMSG_DONE, flags=NLM_F_MULTI, seq=1, pid=72}, "\0\0\0\0"}, 4096, 0, {sa_family=AF_NETLINK, nl_pid=0, nl_groups=00000000}, [112->12]) = 20
getsockname(5, {sa_family=AF_NETLINK, nl_pid=72, nl_groups=00000000}, [112->12]) = 0
close(5)                                = 0
write(1, "\n===== syscall.SockaddrLinklayer"..., 41
===== syscall.SockaddrLinklayer() =====
) = 41
write(1, "\n===== syscall.Bind() =====\n", 28
===== syscall.Bind() =====
) = 28
bind(3, {sa_family=AF_PACKET, sll_protocol=htons(ETH_P_ALL), sll_ifindex=if_nametoindex("eth1"), sll_hatype=ARPHRD_NETROM, sll_pkttype=PACKET_HOST, sll_halen=6, sll_addr=[0x2, 0x42, 0xa, 00, 0xa, 0x14]}, 20) = 0
write(1, "\n===== syscall.Socket() =====\n", 30
===== syscall.Socket() =====
) = 30
socket(AF_PACKET, SOCK_DGRAM, 768)      = 5
socket(AF_NETLINK, SOCK_RAW, NETLINK_ROUTE) = 6
bind(6, {sa_family=AF_NETLINK, nl_pid=0, nl_groups=00000000}, 12) = 0
sendto(6, {{len=17, type=0x12 /* NLMSG_??? */, flags=NLM_F_REQUEST|0x300, seq=1, pid=0}, "\0"}, 17, 0, {sa_family=AF_NETLINK, nl_pid=0, nl_groups=00000000}, 12) = 17
recvfrom(6, [{{len=1180, type=0x10 /* NLMSG_??? */, flags=NLM_F_MULTI, seq=1, pid=72}, "\0\0\4\3\1\0\0\0I\0\1\0\0\0\0\0\7\0\3\0lo\0\0\10\0\r\0\1\0\0\0"...}, {{len=1216, type=0x10 /* NLMSG_??? */, flags=NLM_F_MULTI, seq=1, pid=72}, "\0\0\1\0\214\1\0\0C\20\1\0\0\0\0\0\t\0\3\0eth1\0\0\0\0\10\0\r\0"...}, {{len=1216, type=0x10 /* NLMSG_??? */, flags=NLM_F_MULTI, seq=1, pid=72}, "\0\0\1\0\216\1\0\0C\20\1\0\0\0\0\0\t\0\3\0eth0\0\0\0\0\10\0\r\0"...}], 4096, 0, {sa_family=AF_NETLINK, nl_pid=0, nl_groups=00000000}, [112->12]) = 3612
getsockname(6, {sa_family=AF_NETLINK, nl_pid=72, nl_groups=00000000}, [112->12]) = 0
getsockname(6, {sa_family=AF_NETLINK, nl_pid=72, nl_groups=00000000}, [112->12]) = 0
getsockname(6, {sa_family=AF_NETLINK, nl_pid=72, nl_groups=00000000}, [112->12]) = 0
recvfrom(6, {{len=20, type=NLMSG_DONE, flags=NLM_F_MULTI, seq=1, pid=72}, "\0\0\0\0"}, 4096, 0, {sa_family=AF_NETLINK, nl_pid=0, nl_groups=00000000}, [112->12]) = 20
getsockname(6, {sa_family=AF_NETLINK, nl_pid=72, nl_groups=00000000}, [112->12]) = 0
close(6)                                = 0
write(1, "\n===== syscall.SockaddrLinklayer"..., 41
===== syscall.SockaddrLinklayer() =====
) = 41
write(1, "\n===== syscall.Bind() =====\n", 28
===== syscall.Bind() =====
) = 28
bind(5, {sa_family=AF_PACKET, sll_protocol=htons(ETH_P_ALL), sll_ifindex=if_nametoindex("eth0"), sll_hatype=ARPHRD_NETROM, sll_pkttype=PACKET_HOST, sll_halen=6, sll_addr=[0x2, 0x42, 0xa, 00, 0xb, 0xa]}, 20) = 0
write(1, "Starting raw server...\n", 23Starting raw server...
) = 23
write(1, "\n===== syscall.Recvfrom() =====\n", 32
===== syscall.Recvfrom() =====
) = 32
recvfrom(3, 0xc42009e000, 8214, 0, 0xc42004fcf0, [112]) = ? ERESTARTSYS (To be restarted if SA_RESTART is set)
--- SIGWINCH {si_signo=SIGWINCH, si_code=SI_KERNEL} ---
rt_sigreturn({mask=[]})                 = 45
recvfrom(3, 0xc42009e000, 8214, 0, 0xc42004fcf0, [112]) = ? ERESTARTSYS (To be restarted if SA_RESTART is set)
--- SIGWINCH {si_signo=SIGWINCH, si_code=SI_KERNEL} ---
rt_sigreturn({mask=[]})                 = 45
recvfrom(3, 0xc42009e000, 8214, 0, 0xc42004fcf0, [112]) = ? ERESTARTSYS (To be restarted if SA_RESTART is set)
```