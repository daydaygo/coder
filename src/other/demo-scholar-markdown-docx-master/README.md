# demo-scholar-markdown-docx

```sh
brew install pandoc
pip install pandoc-fignos
pandoc --filter pandoc-fignos -C --bibliography=myref.bib --csl=chinese-gb7714-2005-numeric.csl demo.md -o demo.docx

pandoc a.docx -f docx -t markdown a.md
```

- <https://pandoc.org/installing.html#macos>
- 无法使用 `-t pdf -o demo.pdf` 转 pdf 文件
