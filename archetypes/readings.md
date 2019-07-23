---
{{ $alternate_title := replace .TranslationBaseName "-" " " | humanize | title }}
date: "{{ .Date }}"
title: "{{ slicestr $alternate_title 11 }}"
categories:
  - Reviews
draft: true
---
