# BibleConsole [![wercker status](https://app.wercker.com/status/9b8bacfc036334dd3004d106e2ebb40e/s "wercker status")](https://app.wercker.com/project/bykey/9b8bacfc036334dd3004d106e2ebb40e) [![Gobuild Download](http://gobuild.io/badge/github.com/weslleyandrade/bibleConsole/downloads.svg)](http://gobuild.io/github.com/weslleyandrade/bibleConsole) [![GoWalker](http://img.shields.io/badge/doc-gowalker-blue.svg?style=flat)](https://gowalker.org/github.com/weslleyandrade/bibleConsole)

BibleConsole é um projeto open source que permite o consumo da Bíblia Sagrada via terminal de uma maneira fácil, rápida e dinâmica. O projeto nasceu de dois amigos onde possuem duas paixões em comum: a palavra de Deus e desenvolvimento de software. O foco do projeto é permitir que qualquer pessoa tenha acesso a Bília através de uma interface de linha de comando, tendo a possibilidade de consultar (através de provedores) versículos e capítulos da Bíblia através dos nomes e números. Caso tiver interesse fique a vontade para contribuir. :)

# Exemplo de uso

    [fmamud@go ~]# bibleConsole -l 1 -c 1 -v 1
    No princípio Deus criou os céus e a terra.

    [fmamud@go ~]# bibleConsole -q "romanos 1:16-17"
    16 Não me envergonho do evangelho, porque é o poder de Deus para a salvação de todo aquele que crê: primeiro do judeu, depois do grego.
    17 Porque no evangelho é revelada a justiça de Deus, uma justiça que do princípio ao fim é pela fé, como está escrito: "O justo viverá pela fé".
