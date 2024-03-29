+++
title = "Wanna get insulted by sudo"
author = ["Yash Agarwal"]
date = 2016-04-03T12:20:36+00:00
images = []
tags = ["sudo"]
categories = ["Hacks"]
draft = false
+++

  You might have tried many Linux easters eggs for fun, but you are going to love this hack which makes `sudo` insult you.

  Confused what I am talking about? Here, take a look at this gif to get an idea of how `sudo` can insult you for typing in the incorrect password.

![A terminal screen](/images/posts/2016-04-03/insult.gif "Illustration of the insult")

Now you might be thinking, why would anyone want to take insult? Afterall, nobody likes being insulted. For me, it is just another way to have fun with Linux, and anyway, this is way better than the plain "You entered a wrong password" error message. So let's learn how to do this.

## Enable insults in sudo
You can enable the insults feature in `sudo` by modifying the `sudo` configuration file. To open the `sudo` configuration file, launch a terminal and type the following command.

{{< highlight bash >}}
sudo visudo
{{< /highlight >}}

It will open `/etc/sudoers` configuration file in the terminal, in vim text editor if you have configured it as your default editor. In distros like Ubuntu, it will be opened in nano. Now you will have to find the section where the defaults are listed. Most probably you will find it at the top. Now find the line that starts with `Defaults` and append the word `insults` to the end of the line (any addition to the line is comma separated). If this line is not present then add the following line to the section

{{< highlight bash >}}
Defaults insults
{{< /highlight >}}

(Always use `visudo` as it has a self-check system which will save you from messing up things)

Now save the file. If you are using vim, then use `Ctrl+X` to save the file and quit the editor and if you are using nano then use `Ctrl+X` to leave the editor. At the time of quitting, it will ask you if you want to save the changes. To keep the changes, press `Y`.

![Sample sudoers file](/images/posts/2016-04-03/sudoers.png "Sample sudoers file")

Once you have saved the file, go to terminal and type the following command to clear the old password from `sudo`'s cache.

{{< highlight bash >}}
sudo -k
{{< /highlight >}}

That's all. Use any command with `sudo`. Deliberately type a wrong password and enjoy abusing ...
