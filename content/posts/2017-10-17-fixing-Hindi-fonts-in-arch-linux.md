+++
date = "2017-10-17T21:55:53+05:30"
title = "Fixing Hindi Fonts in Arch Linux"
Categories = ["Arch Linux"]
type = "posts"
+++

When viewing Hindi content any browser in Arch Linux, the rendering looks weird.

![https://hi.wikipedia.org/s/2lu5](/images/before.png)

Doesn't look good right! I'll try to fix this issue in this post. You might need to install the appropriate font support in Arch Linux before applying this fix. The appropriate package for installing Indic Language support is <code>ttf-indic-otf</code>.

Now go to <code>/usr/share/fonts/TTF</code> and and take the backup of two fonts <code>FreeSans.ttf</code> and <code>FreeSerif.ttf</code>. Now delete these two fonts from the directory. Restart the browser and see the difference.

![https://hi.wikipedia.org/s/2lu5](/images/after.png)

This bug is reported in [this](https://bugs.launchpad.net/ubuntu/+source/chromium-browser/+bug/856736) bug report in 2011. I don't know why nobody has fixed it yet. Or maybe I was not able to find the proper solution. This workaround is also given in same bug report.

Hope this helps. :)