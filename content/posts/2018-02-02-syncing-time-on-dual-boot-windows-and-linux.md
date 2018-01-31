---
date: "2018-02-02T00:32:51+05:30"
title: "Syncing Time on Windows & GNU/Linux Dual Boot Setups"
categories:
    - linux
    - System Administration
tags:
    - configuration
    - Arch Linux
draft: true
---

This post is going to be one among those, that I have written for my own reference. Whatever I am going to mention in this post, is not new. Everything has already been said and written many times on many websites and Linux forums.

So I will start with explaining the problem. When you try to dual boot your machine to run both GNU/Linux and Windows operating systems, you might have noticed that the time is not same in both the operating systems. It is generally one operating system showing the correct time, and other one showing wrong time. It happens because Microsoft Windows thinks that the hardware clock (CMOS clock or BIOS clock) of the machine is using localtime (depends on your current time zone) and hence it doesn't do anything and shows you the same time. But most GNU/Linux operating systems (Ubuntu, Arch Linux, etc.) thinks that the hardware clock is set to track UTC time. Hence the mismatch in the time happens. For example, assume that the current real time is `10:22:51` and the hardware clock time is set to `10:22:51`. Windows will interpret this time as local time and show `10:22:51` but \*nix based systems will show `15:52:51` because they will interpret this time as UTC time. Of course, the above example is true if we assume time zone as India, which is `+05:30` from UTC.

This issue can be fixed either from Windows or from GNU/Linux OS. I prefer to adjust the behavior of Windows to use UTC time. It is much more convenient to use when traveling between different time zones. Please note that this method might not work or cause instability with older versions of Windows OS. I have tried this fix on Windows 10, and it works without any issues.

Open an Administrator Command Prompt by pressing `⊞ + x`, then type `a`. This method of opening Administrator Command Prompt does not work on Windows 7.

Now execute the following command:
```
reg add "HKEY_LOCAL_MACHINE\System\CurrentControlSet\Control\TimeZoneInformation" /v RealTimeIsUniversal /d 1 /t REG_DWORD /f
```

Windows Time Service which keeps the clock in Windows OS accurate will still write the localtime to the Real-time clock (RTC) regardless of the registry settings on shutdown. So I prefer to disable the Windows Time Service.
```
sc config w32time start= disabled
```

Now you might need to change the time in your BIOS to UTC time. Although that depends whether your Windows OS was showing correct time before applying the above modifications. If yes, then changing BIOS time to UTC will make sure that both Windows and GNU/Linux convert hardware clock to localtime.

*****
##### References
1. [Multiple Boot Systems Time Conflicts](https://help.ubuntu.com/community/UbuntuTime#Multiple_Boot_Systems_Time_Conflicts)
2. [UTC in Windows](https://wiki.archlinux.org/index.php/Time#UTC_in_Windows)
