#!/usr/bin/env bash
# Time-stamp: <2018-07-12 10:32:45 kmodi>

# https://github.com/be5invis/Iosevka/issues/238#issuecomment-351527918
# pyftsubset - https://github.com/fonttools/fonttools
# Name Identifiers (name-ID): https://developer.apple.com/fonts/TrueType-Reference-Manual/RM06/Chap6name.html
# Creating subset font: https://michaeljherold.com/2016/05/04/creating-a-subset-font.html

# Using the --with-zopfli option required first doing "pip install --user zopfli"
# Using pyftsubset for woff2 required first doing "pip install --user brotli"

# Use http://torinak.com/font/lsfont.html to verify the generated fonts

set -euo pipefail # http://redsymbol.net/articles/unofficial-bash-strict-mode
IFS=$'\n\t'

# glyphs_file="./glyphs.txt" # Source: https://github.com/be5invis/Iosevka/issues/238#issuecomment-384515046
# common_args="--name-IDs+=0,4,6 --text-file=${glyphs_file}"

# Unicode table: https://unicode-table.com/en/
#   U+0020-007E : Basic Latin : https://unicode-table.com/en/blocks/basic-latin/
#   U+00A7-00BE : Latin 1 Supplement: https://unicode-table.com/en/blocks/latin-1-supplement/
#   U+00C0-00FF : more from Latin 1 Supplement, needed to cover non-English names of commentors like "Frédéric"
#   U+2000-205E : General Punctuation : https://unicode-table.com/en/blocks/general-punctuation/ (includes zero width space, curly quotes, etc.)
#   U+25A0,U+25CB : Unicode chars BLACK SQUARE and WHITE CIRCLE for unordered list bullets
#   U+2B89,U+2B8B : Unicode chars UPWARDS BLACK CIRCLED WHITE ARROW and DOWNWARDS BLACK CIRCLED WHITE ARROW for the top/bottom jump links

print_glyph_count () {
    font_file="${1}"
    # echo ${font_file}
    # https://github.com/fonttools/fonttools/issues/1294#issuecomment-404485282
    count=$(ttx -q -o - -t GlyphOrder "${font_file}" | grep -c '<GlyphID id')
    echo "Exported ${count} glyphs to ${font_file}"
}

run_pyftsubset () {
    font="${1}"
    unicodes="${2}"
    # out_dir_base="../../static/fonts"
    # out_dir="${out_dir_base}/${font}/subset"
    out_dir="../../static/fonts/${font}/subset"
    mkdir -p "${out_dir}"

    common_args="--layout-features='' --unicodes=${unicodes}"

    echo "${font}: Generating subset WOFF files .."
    for file in ${font}/ttf/*
    do
        # Get basename without extension from ${file}
        # https://stackoverflow.com/a/2664746/1219634
        tmp="${file##*/}"
        basename_no_ext="${tmp%.*}"
        # echo "file = ${file}"
        # echo "basename_no_ext = ${basename_no_ext}"
        # https://stackoverflow.com/a/407334/1219634
        if [[ $file == *.ttf ]] # The ${font}/ttf HAS to contain .ttf font files.
        then
            font_ext="woff"
            out_file="${out_dir}/${basename_no_ext}.${font_ext}"
            eval "pyftsubset ${file} ${common_args} --flavor=${font_ext} --with-zopfli --output-file=${out_file}"
            print_glyph_count "${out_file}"
        fi
    done

    echo "${font}: Generating subset WOFF2 files .."
    for file in ${font}/ttf/*
    do
        tmp="${file##*/}"
        basename_no_ext="${tmp%.*}"
        if [[ $file == *.ttf ]] # The ${font}/ttf HAS to contain .ttf font files.
        then
            font_ext="woff2"
            out_file="${out_dir}/${basename_no_ext}.${font_ext}"
            eval "pyftsubset ${file} ${common_args} --flavor=${font_ext} --output-file=${out_file}"
            print_glyph_count "${out_file}"
        fi
    done
}

run_pyftsubset "libre-baskerville/2012"              "U+0020-007E,U+00A7-00BE,U+2000-205E,U+00C0-00FF"
run_pyftsubset "andada/2013"                         "U+0020-007E,U+00A7-00BE,U+2000-205E"
run_pyftsubset "merriweather/2.002"                  "U+0020-007E,U+00A7-00BE,U+2000-205E,U+00C0-00FF"
run_pyftsubset "source-sans-pro/2.020R-ro-1.075R-it" "U+0020-007E,U+00A7-00BE,U+2000-205E"
run_pyftsubset "linux-libertine/5.3.0"               "U+0020-007E,U+00A7-00BE,U+2000-205E"
run_pyftsubset "iosevka/1.14.1"                      "U+0020-007E,U+00A7-00BE,U+2000-205E"
run_pyftsubset "symbola/v11.0"                       "U+25A0,U+25CB,U+2B89,U+2B8B"

echo "Done!"
