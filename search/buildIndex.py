from bs4 import BeautifulSoup
from markdown import markdown

f = open("2016-04-10-searching-the-goal.md").readlines()
# html = markdown(f)
# text = ''.join(BeautifulSoup(html, "lxml").findAll(text=True))

content = ''
for line in f:
	html = markdown(line)
	text = ''.join(BeautifulSoup(html, "lxml").findAll(text=True))

	if text.startswith("+++"):
		print(text)
	content = content + ''.join(BeautifulSoup(html, "lxml").findAll(text=True))

print(content)