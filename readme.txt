shred program:

cd shred

build the program:
go build shred.go

two example files are provided for your convenience:
testfile1.txt
testfile2.bin

shred and delete the files:
./shred testfile1.txt
./shred testfile2.bin


bootable linux image via QEMU:

cd qemu 

bash script.sh

the script will install any missing gcc and qemu requirements so it will ask for your sudo password

run the script and it prints out Hello World every 2 seconds.

I ran out of time to do more than what is included.  I hope you like it.  Thanks.


tested on Ubuntu 22.04 LTS see output of uname -a below:
chet@chet-u:~/Dropbox/aTest$ uname -a
Linux chet-u 5.19.0-38-generic #39~22.04.1-Ubuntu SMP PREEMPT_DYNAMIC Fri Mar 17 21:16:15 UTC 2 x86_64 x86_64 x86_64 GNU/Linux
