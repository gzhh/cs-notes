## 日期时间和时区
### Linux系统修改时间与时区

linux系统时钟有两个，一个是硬件时钟，即BIOS时间，就是我们进行CMOS设置时看到的时间，另一个是系统时钟，是linux系统Kernel时间。当Linux启动时，系统Kernel会去读取硬件时钟的设置，然后系统时钟就会独立于硬件运作

时间设置及时钟同步的命令使用方法:

- --- date -s 06/18/14

将时间设置为14点20分50秒

- --- date -s 14:20:50

将时间设置为2014年6月18日14点16分30秒（MMDDhhmmYYYY.ss）

- ---date 0618141614.30

hwclock/clock 命令查看、设置硬件时间查看系统硬件时钟hwclock --show 或者clock --show设置硬件时间 hwclock --set --date="06/18/14 14:55" （月/日/年时:分:秒）或者# clock --set --date="06/18/14 14:55" （月/日/年时:分:秒）

同步系统及硬件时钟。

下图中可以看到硬件和系统时钟相差半小时。可以使用hwclock或者clock进行同步，

硬件时钟与系统时钟同步：

# hwclock --hctosys 或者 # clock --hctosys hc代表硬件时间，sys代表系统时间，即用硬件时钟同步系统时钟

系统时钟和硬件时钟同步：

# hwclock --systohc或者# clock --systohc 即用系统时钟同步硬件时钟
