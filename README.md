# tufte.yashagarwal.in

This repository hosts the code for the design of my blog from April 2020 to July 2023. You can visit the website at [tufte.yashagarwal.in](https://tufte.yashagarwal.in/).

## Building
I used a static site generator `Hugo` for building this blog. The design is a modified version of [tufte-css](https://github.com/edwardtufte/tufte-css), with some added features from my side.
To build the site, run the following commands -
```bash
git clone --recursive https://github.com/yashhere/tufte.yashagarwal.in
cd tufte.yashagarwal.in
yarn install
yarn build
```

After `make`, the site is available under the `public` directory.

To use Hugo's built-in `serve` feature, run the following command -
```bash
hugo serve
```
Now, the website is available at [http://localhost:1313/](http://localhost:1313/).

## Elsewhere
### Highlights
I maintain a website to save my books and the content and highlights I mark while reading books on my Kindle Paperwhite. It serves as a backup in case [goodreads](https://www.goodreads.com/yashhere) decides to shut down someday suddenly. The website can be seen at [highlights.yashagarwal.in](https://highlights.yashagarwal.in/) and the code is hosted at [highlights](https://github.com/yashhere/highlights).

### Old blogs
I also host some of the older versions of my blog. The purpose is to make the process smooth if I decide to go back someday.

One from 2016 can be seen at [jekyll.yashagarwal.in](https://jekyll.yashagarwal.in/), and the code is available at [Jekyll-Blog](https://github.com/yashhere/Jekyll-Blog).

Another one is more recent ([2017](https://github.com/yashhere/tufte.yashagarwal.in/commit/42165108f3cc94cc9c60ae07c1981eb6d5d06751)-[2020](https://github.com/yashhere/tufte.yashagarwal.in/commit/bb7139ba1d4a96127b7a6c7718d5239d7eb35068)). It is live at [minimo.yashagarwal.in](https://minimo.yashagarwal.in/) and the code is hosted at [minimo](https://github.com/yashhere/tufte.yashagarwal.in/tree/minimo).
