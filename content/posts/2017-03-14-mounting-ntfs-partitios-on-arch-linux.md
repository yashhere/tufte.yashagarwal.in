---
categories:
- Linux
- Hacks
- Configuration
- Arch Linux
- Polkit
- NTFS
date: '2017-03-14T10:30:36+00:00'
title: Mounting NTFS partitions on Arch Linux
type: posts
---

Yesterday I installed Arch Linux once again. A clean, bloat free desktop with Budgie Desktop environment with some must have open source tools. Everything worked fine except WiFi and some minor bugs in Budgie(I don't know whether it is a bug in Budgie or just a wrong setting). I also faced the problem of mounting Windows NTFS volumes on user's wish. Arch Linux wiki has details about how to auto mount partitions on start-up, but I had a tough time to find out what exactly needs to be done to simulate the behavior of Ubuntu-like distribution on the mounting of NTFS drives. I got a hint from Arch Linux Wiki about Polkit configuration setting which can be used to allow a standard user to mount partitions. Here is a solution which I searched on various Arch Linux Community pages.

You will need to install <code>ntfs-3g</code>, <code>polkit</code> and <code>udisks2</code> to use this code. Please refer to [Arch Wiki](https://wiki.archlinux.org/index.php/Udisks).

I think when using Gentoo, you will also need to compile the support for NTFS file system in Kernel also. Please see [here](https://wiki.gentoo.org/wiki/NTFS).

Add the following code to <code>/etc/polkit-1/rules.d/10-udisks2.rules</code>

<pre>
// Allow udisks2 to mount devices without authentication for users in the "wheel" group.
polkit.addRule(function(action, subject) {
    if ((action.id == "org.freedesktop.udisks2.filesystem-mount-system" ||
        action.id == "org.freedesktop.udisks2.filesystem-mount") &&
        subject.isInGroup("wheel")) {
            return polkit.Result.YES;
        }
    }
);

polkit.addRule(function(action, subject) {
   if ((action.id == "org.freedesktop.udisks.filesystem-mount-system-internal") && 
        subject.isInGroup("wheel")) {
            return polkit.Result.YES;
        }
    }
);
</pre>

Now you will be able to mount NTFS partition without any problem. :)