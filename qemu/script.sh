#!/bin/bash
echo "installing build tools"
sudo apt-get install build-essential libncurses-dev bison flex libssl-dev libelf-dev qemu-kvm -y
mkdir mykernel
cd mykernel
wget https://cdn.kernel.org/pub/linux/kernel/v5.x/linux-5.15.106.tar.xz
tar xf linux-5.15.106.tar.xz
cd linux-5.15.106
make defconfig
make -j8
cp arch/x86/boot/bzImage ../vmlinuz
cd ..
cat > init.c <<- EOF
#include <stdio.h>
#include <unistd.h>

int main() {
  sleep(1);
  while (1) {
    printf("Hello World\n");
    sleep(2);
  }
}
EOF
gcc -static -o init init.c
echo ./init | cpio -o --format=newc > initramfs
qemu-system-x86_64 -kernel vmlinuz -initrd initramfs -append console=ttyS0 -nographic
