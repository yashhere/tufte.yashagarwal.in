+++
title = "About this blog's new design"
author = ["Yash Agarwal"]
date = 2020-04-01T10:54:48+05:30
categories = ["Updates"]
tags = ["theme"]
images = []
draft = false
+++

{{< load-photoswipe >}}

I was thinking about updating this blog's theme for quite some time. The only issue was that I was not willing to start working on a new theme from scratch. Frontend designing doesn't work *with* me anymore. None of the themes on the Hugo Themes [website](https://themes.gohugo.io/) suited my taste. I liked themes (especially those based on *tufte-css*) on some of the blogs that I follow, but either the themes were made for Jekyll, or their source was not public. In the meanwhile, I lost my interest in writing also. I believe that disinterest has some roots in this delay in changing the look of this blog. Weird, I know.

My requirements for a new theme were simplistic. The stylesheet should not be in SASS/SCSS, no NPM/Webpack/JS nonsense, the mobile version should be functioning with a *hamburger* icon for the menu, and the code should be open-source with a favorable license
 (I did not want to rewrite the theme completely).

Then the Coronavirus lockdown happened in India. Thankfully, I had anticipated this quite early and arrived at my home in Jaipur a couple of days before the announcement when everything was still under control. It gave me an ample amount of free time (outside of work hours) suddenly. By the way, today is the 11th day of my somewhat self-quarantine. Hopefully, I will be alright. :relieved:

I was looking for some themes, and then I found [this](https://scripter.co/) website. This website fitted all my conditions with some extra goodies (the integration of [Webmentions](https://www.w3.org/TR/webmention/) is one such feature). The only issue was that it came with its own share of problems. The author seems to have worked on this theme quite extensively, though I couldn't appreciate the code structure of the theme. It took me a couple of days to restructure the code, modified some CSS to my liking, and hurrah, everything was setup. The only issue is that the theme is broken on the Safari browser. I do not intend to fix that right now. I'll keep that for some future date.

You can see the original theme [here](https://gitlab.com/kaushalmodi/hugo-theme-refined), and my modified version [here](https://github.com/yashhere/refined-mod). It is still work in progress, and a lot can be improved in the structure of code. But for now, this is working for me. I will fix issues as and when they appear.


The Google page speed results are the best I have managed to get till now (although it is almost equivalent to the [original](https://scripter.co/) website, so I haven't contributed much to these scores).

{{< gallery >}}
  {{< figure link="/images/2020-04-01/desktop.png" src="/images/2020-04-01/desktop.png" alt="Desktop score">}}
  {{< figure link="/images/2020-04-01/mobile.png" src="/images/2020-04-01/mobile.png" alt="Mobile score">}}
{{< /gallery >}}

I am trying to switch from [Disqus](https://disqus.com/) to [Commento](https://commento.io/), but the pricing is an issue here. [Self-hosting](https://gitlab.com/commento/commento/) Commento is an option, but I am not sure if I am ready to maintain my own servers. If I go with self-hosting, I have a lot of things to host. But the cost is a significant obstruction. But, it is also true that I will not get the convenience of privacy for free. I need to think about it carefully.

One another trick that I found out was to enable `gzip` compression on pages served by Gitlab. The simple idea is to just gzip every page on the website in the same directory where the original HTML page is present. Gitlab will automatically pick the gzipped version and serve it to the client. You can see that in my `.gitlab-ci.yml` [file](https://gitlab.com/yashhere/yashhere.gitlab.io/-/blob/source/.gitlab-ci.yml#L21).

I was trying to create a staging environment (a Gitlab site, staging.yashagarwal.in, based on a branch of my Git repo, the master branch would serve the root domain) where I could test any changes to this website before pushing it to the actual site. But it does not seem possible with current Gitlab features. I do not want to move to Netlify or such third party services to achieve this, so for now, I will just keep myself limited to local testing.
