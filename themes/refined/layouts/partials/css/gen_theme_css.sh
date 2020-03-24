#!/usr/bin/env bash
# Time-stamp: <2018-08-01 16:40:40 kmodi>

# Pygments
# https://github.com/richleland/pygments-css
gen_pygments_css () {
    pygmentize -S "${1}" -f html -a .highlight > "${1}"_pygments.css
}
gen_pygments_css trac

# Go Chroma
# https://gohugo.io/content-management/syntax-highlighting/#generate-syntax-highlighter-css
gen_chroma_css () {
    hugo gen chromastyles --style="${1}" > "${1}"_chroma.css
}
gen_chroma_css trac
gen_chroma_css github
