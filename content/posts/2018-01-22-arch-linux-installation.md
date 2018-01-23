---
date: "2018-01-26T17:20:12+05:30"
title: "Arch Linux Installation Guide"
categories:
  -
tags:
  -
draft: true
---
Arch Linux is a Linux distribution known for its not-so-beginner-friendly command line installer, no-ready-to-use system after installation and requirement for above average knowledge of command line. However, Arch Linux allows me to setup a system in my desired state in shortest possible time with least effort. This is why, I keep coming back to Arch Linux even after some of its annoyances.


This guide is written primarily for my reference, as someone who has installed Arch Linux several times, I still can't remember all the installation steps perfectly. Most of the steps have been taken from [Arch wiki](https://wiki.archlinux.org/index.php/installation_guide) and should work on other setups also.



**Note**: All the commands are run in root shell unless otherwise specified.

###   1. Check your network connection

If you are behind a captive portal, use links to open browser and login into your network. For WiFi connections, use `wifi-menu`. LAN connections should not require any setup. The boot environment should automatically detect any wired network. After connecting, test your connection by pinging any website:
```bash
ping -c 5 google.com
```


###   2. Partition the disks

If Windows 8 or above is already installed on your machine, then your hard disk is probably using GPT partitioning scheme. In that case, use `gdisk` to partition your hard disk. If you use `fdisk` on  a GPT partitioned HDD, there is a possibility of data loss.

My preferred setup is to have one root partition and one home partition, and use EFI partition created by Windows to install boot-loader. The root and home partition will be formatted using `ext4` file-system and the EFI partition should be formatted using `FAT32` file-system.

For this guide, I am assuming that the EFI partition is `sda1`, root partition is `sda9` and home partition is `sda10`.

Now to format the partitions with `ext4` file-system:
```bash
mkfs.ext4 /dev/sda9
mkfs.ext4 /dev/sda10
```

###   3. Mount the partitions
Now mount the root partition (`sda9` in this case) to `/mnt`
```bash
mount /dev/sda9 /mnt
```

If you have created any other partitions in previous steps, mount them at appropriate locations.
```bash
mkdir /mnt/home
mount /dev/sda10 /mnt/home

mkdir /mnt/boot
mount /dev/sda1 /mnt/boot
```

###   4. Install the base file-system
To install the base system and some development tools, issue the following command.
```bash
pacstrap /mnt base base-devel
```

This will take a while to download and install. After it finishes, it will give you a bare-bone Arch Linux system with just the tools required to run a Linux distribution, no other software is installed.

###   5. Generate `/etc/fstab`
The `/etc/fstab` file stores the information about file systems of partitions and how to mount the partitions on system boot up. To generate this file, issue the following command:
```bash
genfstab -U /mnt >> /mnt/etc/fstab
```
If you prefer to use partition labels(`sda1`, `sda9` etc.) instead of UUID, then use `-L` flag in place of `-U`.

###   6. `chroot` into the system

From the [Arch wiki](https://wiki.archlinux.org/index.php/Change_root):

  > Chroot is an operation that changes the apparent root directory for the current running process and their children. A program that is run in such a modified environment cannot access files and commands outside that environmental directory tree. This modified environment is called a `chroot` jail.

At this step, we will go into the root of the newly installed system at `/mnt` and pretend as if we are logged into this system.
```bash
arch-chroot /mnt
```

###   7. Setup the time zone, locale and hostname
Browse the `/use/share/zoneinfo` directory to find your location entries. My location is India, so I will use this command.
```bash
ln -sf /usr/share/zoneinfo/Asia/Kolkata /etc/localtime
```

To set the hardware clock:
```bash
hwclock --systohc
```

To set the locale for your system, open the `/etc/locale.gen` file and uncomment your language. or run the following command for `en_US.UTF-8 UTF-8`.
```bash
LANG=C perl -i -pe 's/#(en_US.UTF)/$1/' /etc/locale.gen
```
Now generate the localizations with
```bash
locale-gen
```
Then set the `LANG` variable in `/etc/locale.conf` accordingly, or run the following command:
```bash
localectl set-locale LANG="en_US.UTF-8"
```

To set the hostname for your machine:
```bash
hostnamectl set-hostname your-host-name
```

To allow other machines to address the host by name, it is necessary to edit the `/etc/hosts` file to look like this:
```bash
127.0.0.1	localhost.localdomain	      localhost
::1	        localhost.localdomain	      localhost
127.0.1.1 	your-host-name.localdomain    your-host-name
```

###   8. Create user account

Before creating user account, set passwd for `root` account
```bash
passwd
```

Now create a local account for your user
```bash
useradd -m -G wheel -s /bin/bash your-user-name
```
This will setup your user account, create a home directory for your user, set the default shell to `bash` and add your user to `wheel` group, which is necessary to do to gain `sudo` access in later steps.

Set password for your user.
```bash
passwd your-user-name
```

###   9. Enable `sudo` access
This allows you to use root privileges without using root account. To enable this, first open `/etc/sudoers` file
```bash
nano /etc/sudoers
```

Now uncomment the following line to enable root privilege for all the users inside wheel group:
```bash
# %wheel ALL=(ALL) ALL
```

Now you can safely disable root account
```bash
passwd -l root

# login into your user account
su your-user-name
```
From this point onwards, it is necessary to append `sudo` to any command that requires `root` privileges.

###   10. Install bootloader
My preferred bootloader of choice is `grub`. To install `grub`, we need to install following packages.
```bash
pacman -S grub efibootmgr
```

Now install grub with following command.
```bash
grub-install --target=x86_64-efi --efi-directory=/boot --bootloader-id=arch
```
Here `--efi-directory` is the folder where the EFI partiton is mounted [step 3](#) and `--bootloader-id` is the label that will appear in your UEFI boot menu entry.

This particular step is specific to my system hardware, you might not need to run this step. I need to add `pci=nommconf` to my kernel boot parameters in `/etc/default/grub`, otherwise `tty` prints error messages continuously.

Now run to generate grub configuration file.
```bash
grub-mkconfig -o /boot/grub/grub.cfg
```

If you encounter any errors related to `lvm` during installation of grub, then follow these steps.
```bash
# come out of chroot
exit
mkdir /mnt/hostrun
mount --bind /run /mnt/hostrun

# back to chroot
arch-chroot /mnt
mkdir /run/lvm
mount --bind /hostrun/lvm /run/lvm
```

Now you can install `grub` without any errors.

###   11. Configure the network
By default, your current system cannot connect to network in current state. For wireless networking, install the following packages. I prefer to use [`NetworkManager`](https://wiki.archlinux.org/index.php/NetworkManager) for my network management, even when I am not using GNOME.
```bash
sudo pacman -S iw wpa_supplicant dialog networkmanager network-manager-applet dhclient
```

`NetworkManager` supports basic DHCP configuration. For full support, I have installed `dhclient`. `NetworkManager` also supports automatic wired connection detection and comes with a curses based tool `nmtui` to setup wireless connection.

To enable NetworkManager to start at system startup
```bash
sudo systemctl enable NetworkManager.service
```

###   12. Reboot now
If you had performed the `lvm` troubleshooting steps during grub install, Then
```bash
umount /run/lvm
```

Now exit from chroot by typing exit in shell. Unmount all the mounted partitions with:
```bash
umount -R /mnt
```
Finally reboot your machine by typing `reboot`.


###   13. Install general utilities

##### Install a terminal based browser
In case, you can't login to captive portal after restart, you will require a browser to enter your login credentials.
```bash
sudo pacman -S elinks w3m
```

##### Install X server.
```bash
sudo pacman -S xorg
```
This will install minimal X desktop environment required for testing with fonts.

##### Enable multilib repository for 32-bit package support
To enable multilib repository, uncomment the [multilib] section in `/etc/pacman.conf`.
```bash
[multilib]
Include = /etc/pacman.d/mirrorlist
```

Now upgrade your system.
```bash
sudo pacman -Syyu
```
##### Install video and touchpad drivers
```bash
sudo pacman -S xf86-video-intel xf86-input synaptics
```

##### Install `pacaur` to install packages from AUR
```bash
sudo pacman -S expac yajl --noconfirm
cd /tmp
gpg --recv-keys --keyserver hkp://pgp.mit.edu:80 1EB2638FF56C0C53
curl -o PKGBUILD https://aur.archlinux.org/cgit/aur.git/plain/PKGBUILD?h=cower
makepkg -i PKGBUILD --noconfirm
curl -o PKGBUILD https://aur.archlinux.org/cgit/aur.git/plain/PKGBUILD?h=pacaur
makepkg -i PKGBUILD --noconfirm
cd
```

##### Install graphical browsers
```bash
pacaur -S firefox chromium
```

##### Install text editors
```bash
pacaur -S sublime-text-dev atom-editor-git visual-studio-code-bin neovim neovim-drop-in leafpad
```

##### Setup LAMP stack
1. Install `apache` server
```bash
sudo pacman -S apache
```

  * Make your user-directory available to apache server
  ```bash
  mkdir ~/public_html
  chmod o+x ~
  chmod o+x ~/public_html
  chmod -R o+r ~/public_html
  ```

  * To enable virtualhosts, uncomment the following line in `/etc/httpd/conf/httpd.conf`
        ```bash
        Include conf/extra/httpd-vhosts.conf
        ```
        and add your virtualhost configuration in following file.
        ```bash
        sudo vim /etc/httpd/conf/extra/httpd-vhosts.conf
        ```
        To test the virtual hosts on you local machine, add the virtual names to your `/etc/hosts` file.


2. Install `PHP`:
```bash
sudo pacman -S php php-apache
```
  * Use PHP with apache:

        Open `/etc/httpd/conf/httpd.conf` and uncomment following line.
        ```bash
        LoadModule mpm_prefork_module modules/mod_mpm_prefork.so
        ```
        and comment the following installation_guide
        ```bash
        #LoadModule mpm_event_module modules/mod_mpm_event.so
        ```

        Now add these lines to `/etc/httpd/conf/httpd.conf`:

        * Add these at the end of `LoadModule` section.
        ```bash
        LoadModule php7_module modules/libphp7.so
        AddHandler php7-script .php
        ```

        * Place this at the end of the Include list:
        ```bash
        Include conf/extra/php7_module.conf
        ```

3. Install MySQL server
```bash
sudo pacman -S mariadb
```

      * Initialize the MariaDB data directory prior to starting the service. To do so, run:

          ```bash
          sudo mysql_install_db --user=mysql --basedir=/usr --datadir=/var/lib/mysql
          ```

      * Then issue the commands to start the database server

          ```bash
          sudo systemctl enable mariadb.service
          sudo systemctl start mariadb.service
          ```

      * To apply recommended security settings to your database, run

          ```bash
          sudo mysql_secure_installation
          ```

4. Install `phpmyadmin`
```
sudo pacman -S phpmyadmin php-mcrypt
```
  * Enable `mysqli`, `mcrypt`, `zip` and `bz2` extensions in `/etc/php/php.ini`.
  * Create the apache configuration file `/etc/httpd/conf/extra/phpmyadmin.conf`

          ```bash
          Alias /phpmyadmin "/usr/share/webapps/phpMyAdmin"
          <Directory "/usr/share/webapps/phpMyAdmin">
              DirectoryIndex index.php
              AllowOverride All
              Options FollowSymlinks
              Require all granted
          </Directory>
          ```
    Then include following in `/etc/httpd/conf/httpd.conf`
    ```bash
      # phpMyAdmin configuration
      Include conf/extra/phpmyadmin.conf
    ```



Now restart httpd service to apply settings.
```bash
sudo systemctl restart httpd
```
Once all these steps are done, your LAMP stack should be working.


##### Setup power management
Install tlp and some of its optional dependencies
```bash
sudo pacman -S tlp tlp-rdw bash-completion ethtool lsb-release smartmontools
```

Then enable tlp services
```bash
sudo systemctl enable tlp.service
sudo systemctl enable tlp-sleep.service

# mask some services for tlp to work properly
sudo systemctl mask systemd-rfkill.service
sudo systemctl mask systemd-rfkill.socket
```
##### Install i3 and other tools

These all tools are part of my i3 config with exception of the theme related packages. So installing them here will help me later while setting up the i3.
```bash
pacaur -S i3 rofi polybar xautolock powerline-fonts-git i3lock-fancy-git compton scrot feh dunst unclutter xfce4-power-manager numlockx lxappearance arc-gtk-theme gtk-engine-murrine gnome-themes-standard
```


##### Install some other common tools
```bash
sudo pacman -S vlc openssh npm imagemagick git
```


