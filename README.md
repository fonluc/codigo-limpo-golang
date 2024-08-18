# Código Go Limpo

![comic](assets/clean-code-comic.jpeg)

## Prefácio: Por que Escrever Código Limpo?

Este documento é uma referência para a comunidade Go, que visa ajudar os desenvolvedores a escreverem códigos mais limpos. Seja trabalhando em um projeto pessoal ou como parte de uma equipe maior, escrever código limpo é uma habilidade importante. Estabelecer bons paradigmas e padrões consistentes e acessíveis para a escrita de código limpo pode ajudar a evitar que os desenvolvedores percam horas tentando entender seu próprio trabalho (ou o de outros).

*Nós não lemos código, nós o decodificamos – Peter Seibel*

Como desenvolvedores, às vezes somos tentados a escrever código de uma maneira que seja conveniente no momento, sem levar em conta as melhores práticas; isso torna revisões de código e testes mais difíceis. Em certo sentido, estamos codificando e, ao fazer isso, dificultando a decodificação de nosso trabalho por outros. Mas queremos que nosso código seja utilizável, legível e manutenível. E isso exige codificar da maneira certa, não da maneira fácil.

Este documento começa com uma introdução simples e curta aos fundamentos da escrita de código limpo. Mais tarde, discutiremos exemplos concretos de refatoração específicos para Go.

*Uma breve palavra sobre o gofmt*

Gostaria de dedicar algumas frases para esclarecer minha opinião sobre o gofmt, porque há muitas coisas com as quais não concordo em relação a essa ferramenta. Prefiro snake case ao invés de camel case, e gosto bastante que minhas variáveis constantes sejam em maiúsculas. E, naturalmente, também tenho muitas opiniões sobre o posicionamento das chaves. Dito isso, o gofmt nos permite ter um padrão comum para escrever código Go, e isso é uma coisa ótima. Como desenvolvedor, posso certamente apreciar que os programadores Go possam se sentir um pouco restritos pelo gofmt, especialmente se discordarem de algumas de suas regras. Mas, na minha opinião, um código homogêneo é mais importante do que ter total liberdade expressiva.

### Sumário

- Introdução ao Código Limpo
- Desenvolvimento Orientado a Testes
- Convenções de Nomenclatura
- Comentários
- Nomeação de Funções
- Nomeação de Variáveis
- Limpando Funções
- Comprimento de Funções
- Assinaturas de Funções
- Escopo de Variáveis
- Declaração de Variáveis
- Go Limpo
- Valores de Retorno
- Retornando Erros Definidos
- Retornando Erros Dinâmicos
- Ponteiros em Go
- Closures São Ponteiros de Função
- Interfaces em Go
- A interface vazia {}
- Resumo

### Introdução ao Código Limpo

Código limpo é o conceito pragmático de promover software legível e manutenível. Código limpo estabelece confiança na base de código e ajuda a minimizar as chances de bugs descuidados serem introduzidos. Também ajuda os desenvolvedores a manterem sua agilidade, que normalmente diminui à medida que a base de código se expande devido ao aumento do risco de introdução de bugs.