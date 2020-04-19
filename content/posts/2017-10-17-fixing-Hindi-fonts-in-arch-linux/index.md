+++
title = "Fixing Hindi Fonts in Arch Linux"
author = ["Yash Agarwal"]
date = 2017-10-17T21:55:53+05:30
images = []
tags = ["Arch Linux"]
categories = ["Hacks"]
draft = false
+++



When viewing Hindi content in any browser in Arch Linux, the rendering looks weird.

{{< figure src="/images/posts/2017-10-17/before.png" alt="before applying the fix" caption="before applying the fix" caption-position="bottom" caption-effect="appear">}}

It doesn't look good, right! I'll try to fix this issue in this post. You might need to install the appropriate [font support](https://wiki.archlinux.org/index.php/fonts) in Arch Linux before applying this fix. The suitable package for installing Indic Language support is [ttf-indic-otf](https://www.archlinux.org/packages/extra/any/ttf-indic-otf/).

Now go to `/usr/share/fonts/TTF` and take the backup of two fonts `FreeSans.ttf` and `FreeSerif.ttf`. Now delete these two fonts from the directory. Restart the browser and see the difference.

{{< figure src="/images/posts/2017-10-17/after.png" alt="after applying the fix" caption="after applying the fix" caption-position="bottom" caption-effect="appear">}}

This bug is reported in [this](https://bugs.launchpad.net/ubuntu/+source/chromium-browser/+bug/856736) bug report in 2011. I don't know why nobody has fixed it yet. Or maybe I was not able to find the proper solution. This workaround is also given in the same bug report.

I hope this helps. :)
