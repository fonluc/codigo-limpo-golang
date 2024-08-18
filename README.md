# Código Go Limpo

![comic](assets/clean-code-comic.jpeg)

*Traduzido de: "https://github.com/Pungyeon/clean-go-article"*

## Prefácio: Por que Escrever Código Limpo?

Este documento é uma referência para a comunidade Go, que visa ajudar os desenvolvedores a escreverem códigos mais limpos. Seja trabalhando em um projeto pessoal ou como parte de uma equipe maior, escrever código limpo é uma habilidade importante. Estabelecer bons paradigmas e padrões consistentes e acessíveis para a escrita de código limpo pode ajudar a evitar que os desenvolvedores percam horas tentando entender seu próprio trabalho (ou o de outros).

*Nós não lemos código, nós o decodificamos – Peter Seibel*

Como desenvolvedores, às vezes somos tentados a escrever código de uma maneira que seja conveniente no momento, sem levar em conta as melhores práticas; isso torna revisões de código e testes mais difíceis. Em certo sentido, estamos codificando e, ao fazer isso, dificultando a decodificação de nosso trabalho por outros. Mas queremos que nosso código seja utilizável, legível e manutenível. E isso exige codificar da maneira certa, não da maneira fácil.

Este documento começa com uma introdução simples e curta aos fundamentos da escrita de código limpo. Mais tarde, discutiremos exemplos concretos de refatoração específicos para Go.

*Uma breve palavra sobre o gofmt*

Gostaria de dedicar algumas frases para esclarecer minha opinião sobre o gofmt, porque há muitas coisas com as quais não concordo em relação a essa ferramenta. Prefiro snake case ao invés de camel case, e gosto bastante que minhas variáveis constantes sejam em maiúsculas. E, naturalmente, também tenho muitas opiniões sobre o posicionamento das chaves. Dito isso, o gofmt nos permite ter um padrão comum para escrever código Go, e isso é uma coisa ótima. Como desenvolvedor, posso certamente apreciar que os programadores Go possam se sentir um pouco restritos pelo gofmt, especialmente se discordarem de algumas de suas regras. Mas, na minha opinião, um código homogêneo é mais importante do que ter total liberdade expressiva.

### Sumário

- [Introdução ao Código Limpo](#Introdução ao Código Limpo)
- [Desenvolvimento Orientado a Testes] (#Desenvolvimento Orientado a Testes)
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

### Desenvolvimento Orientado a Testes

O desenvolvimento orientado a testes é a prática de testar seu código frequentemente ao longo de ciclos de desenvolvimento curtos ou sprints. Isso contribui para a limpeza do código ao convidar os desenvolvedores a questionar a funcionalidade e o propósito de seu código. Para facilitar os testes, os desenvolvedores são incentivados a escrever funções curtas que fazem apenas uma coisa. Por exemplo, é consideravelmente mais fácil testar (e entender) uma função de 4 linhas do que uma de 40.

O desenvolvimento orientado a testes consiste no seguinte ciclo:

1. Escrever (ou executar) um teste
2. Se o teste falhar, fazer com que ele passe
3. Refatorar seu código conforme necessário
4. Repetir

Testar e refatorar estão entrelaçados nesse processo. À medida que você refatora seu código para torná-lo mais compreensível ou manutenível, é necessário testar suas mudanças minuciosamente para garantir que você não alterou o comportamento de suas funções. Isso pode ser extremamente útil à medida que a base de código cresce.

### Convenções de Nomenclatura

#### Comentários

Gostaria de primeiro abordar o tópico de comentar código, que é uma prática essencial, mas tende a ser mal aplicada. Comentários desnecessários podem indicar problemas com o código subjacente, como o uso de convenções de nomenclatura ruins. No entanto, se um comentário específico é "necessário" ou não é um pouco subjetivo e depende de quão legível o código foi escrito. Por exemplo, a lógica de um código bem escrito pode ainda ser tão complexa que requer um comentário para esclarecer o que está acontecendo. Nesse caso, pode-se argumentar que o comentário é útil e, portanto, necessário.

No Go, de acordo com o gofmt, todas as variáveis e funções públicas devem ser anotadas. Acho isso absolutamente aceitável, pois nos dá regras consistentes para documentar nosso código. No entanto, sempre quero distinguir entre comentários que possibilitam a documentação gerada automaticamente e todos os outros comentários. Comentários de anotação, para documentação, devem ser escritos como documentação — devem estar em um nível alto de abstração e preocupar-se o mínimo possível com a implementação lógica do código.

Digo isso porque há outras maneiras de explicar o código e garantir que ele esteja sendo escrito de forma compreensível e expressiva. Se o código não for nenhuma dessas coisas, algumas pessoas acham aceitável introduzir um comentário explicando a lógica convoluta. Infelizmente, isso não ajuda muito. Para começar, a maioria das pessoas simplesmente não lerá os comentários, pois eles tendem a ser muito intrusivos na experiência de revisão de código. Além disso, como você pode imaginar, um desenvolvedor não ficará muito feliz se for forçado a revisar um código pouco claro repleto de comentários. Quanto menos as pessoas tiverem que ler para entender o que seu código está fazendo, melhor será.

Vamos dar um passo para trás e olhar alguns exemplos concretos. Veja como você não deve comentar seu código:

```go
// iterar sobre o intervalo de 0 a 9 
// e invocar a função doSomething
// para cada iteração
for i := 0; i < 10; i++ {
  doSomething(i)
}
```

Isso é o que eu gosto de chamar de comentário tutorial; é bastante comum em tutoriais, que frequentemente explicam a funcionalidade de baixo nível de uma linguagem (ou programação em geral). Embora esses comentários possam ser úteis para iniciantes, são absolutamente inúteis em código de produção. Espera-se que não estejamos colaborando com programadores que não entendem algo tão simples quanto uma construção de loop quando começam a trabalhar em uma equipe de desenvolvimento. Como programadores, não deveríamos ter que ler o comentário para entender o que está acontecendo — sabemos que estamos iterando sobre o intervalo de 0 a 9 porque podemos simplesmente ler o código. Daí o provérbio:

*"Documente o porquê, não o como." – Venkat Subramaniam*

Seguindo essa lógica, agora podemos alterar nosso comentário para explicar por que estamos iterando do intervalo de 0 a 9:

```go
// instanciar 10 threads para lidar com a carga de trabalho futura
for i := 0; i < 10; i++ {
  doSomething(i)
}
```

Agora entendemos por que temos um loop e podemos saber o que estamos fazendo apenas lendo o código... Mais ou menos.

Isso ainda não é o que eu consideraria código limpo. O comentário é preocupante porque provavelmente não deveria ser necessário expressar tal explicação em prosa, assumindo que o código esteja bem escrito (o que não está). Tecnicamente, ainda estamos dizendo o que estamos fazendo, não por que estamos fazendo isso. Podemos expressar facilmente esse "o que" diretamente em nosso código usando nomes mais significativos:

```go
for workerID := 0; workerID < 10; workerID++ {
  instantiateThread(workerID)
}
```

Com apenas algumas mudanças em nossos nomes de variáveis e funções, conseguimos explicar o que estamos fazendo diretamente no nosso código. Isso é muito mais claro para o leitor porque ele não terá que ler o comentário e depois mapear a prosa para o código. Em vez disso, eles podem simplesmente ler o código para entender o que está fazendo.

Claro, este foi um exemplo relativamente trivial. Escrever código claro e expressivo infelizmente não é sempre tão fácil; pode se tornar cada vez mais difícil à medida que a base de código cresce em complexidade. Quanto mais você praticar escrever comentários com essa mentalidade e evitar explicar o que você está fazendo, mais limpo seu código se tornará.