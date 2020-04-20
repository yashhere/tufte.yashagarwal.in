{{ $alternate_title := replace .Name "-" " " | humanize | title }}
+++
title = "{{ slicestr $alternate_title 11 }}"
author = ["Yash Agarwal"]
date = {{ .Date }}
draft = true
+++
