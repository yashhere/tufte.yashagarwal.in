---
date: "2018-01-21T11:23:23+05:30"
title: "My Own Configuration Manager"
categories:
  - Linux
---

I have been using Linux since I was in my second year of undergraduate. My experience with the dotfiles(configuration files) also goes back to the same time.

For the uninformed, in Linux, it is common to configure a lot of settings, configurations, or themes within dotfiles. Dotfiles are files in a Linux user's home directory that begin with a dot or a full-stop character. This dot indicates to the operating system, that these files are used to store the settings of programs like `vim` or shells like `bash` or `fish` to name a few.

In the beginning, I was keeping a manual backup of my dotfiles by copying them from time to time in a git repository. But the list soon started getting huge, that it became very difficult for me to keep track of the changes. Then I moved to symlinks. I started symlinking all the dotfiles in my git repository to their usual locations. This setup works perfectly fine, but as my collection of dotfiles increased, It became very cumbersome for me to symlink each and every dotfile manually.

I also tried few tools built for this particular purpose. Some of them are `vcsh`, `mr`, `GNU Stow`. These tools work just fine, but I was not willing to learn new tools just for maintaining my dotfiles. At last, I decided to write my own tool to solve this problem. This way, there will not be any external dependency and this tool will also become part of my dotfiles.

### Design
The tool, in itself, is inspired by the UNIX philosophy of keeping configuration files for settings of the programs. This configuration system has a JSON formatted dotfile.

[Here](https://github.com/yashhere/ConMan) is the source code of the configuration system. Let's look at the file structure of the repository.
```text
|-- .backups
|   |-- 08-01-2018-15:47
|   |-- 08-01-2018-19:30
|   |-- ......
|-- configure.py
|-- current_status
|-- dotfiles
|   |-- dunst
|   |-- gtk-3.0
|   |-- i3
|   |-- ......
|-- dotfiles.json
|-- LICENSE
`-- README.md
```

During the initial setup, you need to edit the `dotfiles.json` file to suit your setup. A relevant section of the JSON file is given below.
```json
{
  "pre": [
    {
      "name": "cloning repository",
      "command": "git",
      "subcommand": "clone",
      "argument": "https://github.com/yashhere/dotfiles.git"
    }
  ],
  "linking": [
      {
        "name": "bashrc",
        "src": "dotfiles/.bashrc",
        "dest": ".bashrc"
      },
      {
        "name": "bash_profile",
        "src": "dotfiles/.bash_profile",
        "dest": ".bash_profile"
      },
      {
        "name": "profile",
        "src": "dotfiles/.profile",
        "dest": ".profile"
      },
      {
        "name": "i3",
        "src": "dotfiles/i3",
        "dest": ".config/i3"
      },
  ]
}
```

As can be seen, the JSON file has an array variable `linking` which can be used to set the paths for each configuration file and folders. The `configure.py` script also requires a dotfiles folder to be present in the current directory. The folder can be created manually or if it is already version controlled on GitHub, then the script can clone it. For that, the `pre` section in the `dotfiles.json` need to be edited.

Your dotfiles and config folders go inside the `dotfiles` folder. You just need to copy all your current configurations to this folder to get started.

So, how does the script know, where a file or folder will be linked to. Simple, you just need to edit the `dotfiles.json` file. For example, if you want to setup configurations of i3 to its exact location(which is, `$HOME/.config/i3`), then you need to create a new JSON object in the `linking` array, like this.
```json
{
  "name": "i3",
  "src": "dotfiles/i3",
  "dest": ".config/i3"
}
```

Here the `name` is used to identify the configuration file, the `src` parameter is the location of your config file/folder in the dotfiles directory and the `dest` parameter is the final destination of the file/folder. Keen observers would notice that I have not given `$HOME` anywhere. It is understood that the configuration will go to the current user's home directory. So the `dest` is relative to the user's home directory and `src` is relative to the directory from which the `configure.py` script is executed.

And you are done! All you have to do is to run `configure.py` and all your dotfiles and folders will be symlinked to their rightful place.

The `current_status` file saves all the symlink locations that are being managed by the script, for your easy reference and to debug any error.

### Behind the Scenes
A lot to cool things happen behind the scenes. The script will check if any previous symlink exists at the given `dest` location. It removes any symlinks to avoid redundancy. If the dest already has any dotfile or folder, then it backs it up in the `.backups` under current date and time before being symlinked to avoid any potential data loss.


Hope the article was useful. Cheers :)
