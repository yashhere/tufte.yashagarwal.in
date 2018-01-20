---
date: "2018-01-20T11:23:23+05:30"
title: "My Own Configuration Manager"
categories:
  - Linux
---

I have been using Linux since I was in my second year of undergraduate. My experience with the dotfiles(configuration files) also goes back to the same time.

For the uninformed, in Linux, it is common to configure a lot of settings, configurations, or themes within dotfiles. Dotfiles are files in a Linux user's home directory that begin with a dot or a full-stop character. This dot indicates to the operating system, that these files are used to store the settings of programs like `vim` or shells like `bash` or `fish` to name a few.

In the beginning, I was keeping manual backup of my dotfiles by copying them from time to time in a git repository. But the list soon started getting huge, that it became very difficult for me to keep track of the changes. Then I moved to symlinks. I started symlinking all the dotfiles in my git repository to their usual locations. This setup works perfectly fine, but as my collection of dotfiles increased, It became very cumbersome for me to symlink each and every dotfile manually.

I also tried few tools built for this particular purpose. Some of them are `vcsh`, `mr`, `GNU Stow`. These tools works just fine, but I was not willing to learn a new tools just for maintaining my dotfiles. Atlast, I decided to write my own tool to solve this problem. This way, their will not be any external dependency and this tool will also become part of my dotfiles.

### Design
The tool has a python script and a configuration file. The configuration file
