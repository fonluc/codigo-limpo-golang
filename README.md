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

- [Código Go Limpo](#código-go-limpo)
  - [Prefácio: Por que Escrever Código Limpo?](#prefácio-por-que-escrever-código-limpo)
    - [Sumário](#sumário)
    - [Introdução ao Código Limpo](#introdução-ao-código-limpo)
    - [Desenvolvimento Orientado a Testes](#desenvolvimento-orientado-a-testes)
    - [Convenções de Nomenclatura](#convenções-de-nomenclatura)
    - [Comentários](#comentários)
    - [Nomeação de Funções](#nomeação-de-funções)
    - [Nomeação de Variáveis](#nomeação-de-variáveis)
    - [Limpeza de Funções](#limpeza-de-funções)
    - [Comprimento da Função](#comprimento-da-função)
    - [Assinaturas de Função](#assinaturas-de-função)
    - [Declaração de Variáveis](#declaração-de-variáveis)
    - [Go Limpo](#go-limpo)
    - [Valores Retornados](#valores-retornados)
    - [Valores Nil](#valores-nil)

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

### Comentários

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

### Nomeação de Funções

Vamos agora abordar as convenções de nomeação de funções. A regra geral aqui é realmente simples: quanto mais específica a função, mais geral deve ser seu nome. Em outras palavras, queremos começar com um nome de função muito amplo e curto, como `Run` ou `Parse`, que descreve a funcionalidade geral. Vamos imaginar que estamos criando um analisador de configuração. Seguindo essa convenção de nomenclatura, nosso nível superior de abstração pode ser algo como o seguinte:

```go
func main() {
    configpath := flag.String("config-path", "", "caminho do arquivo de configuração")
    flag.Parse()

    config, err := configuration.Parse(*configpath)
    
    ...
}
```

Focaremos na nomeação da função `Parse`. Apesar de o nome desta função ser muito curto e geral, está bastante claro o que ela tenta alcançar.

Quando vamos um nível mais profundo, a nomeação das nossas funções se torna um pouco mais específica:

```go
func Parse(filepath string) (Config, error) {
    switch fileExtension(filepath) {
    case "json":
        return parseJSON(filepath)
    case "yaml":
        return parseYAML(filepath)
    case "toml":
        return parseTOML(filepath)
    default:
        return Config{}, ErrUnknownFileExtension
    }
}
```

Aqui, distinguimos claramente as chamadas de funções aninhadas de seu pai sem ser excessivamente específico. Isso permite que cada chamada de função aninhada faça sentido por si só, bem como no contexto do pai. Por outro lado, se tivéssemos nomeado a função `parseJSON` como `json`, ela não poderia se sustentar por conta própria. A funcionalidade se perderia no nome, e não poderíamos mais dizer se essa função está analisando, criando ou convertendo JSON.

Observe que `fileExtension` é um pouco mais específico. No entanto, isso ocorre porque sua funcionalidade é de fato bastante específica por natureza:

```go
func fileExtension(filepath string) string {
    segments := strings.Split(filepath, ".")
    return segments[len(segments)-1]
}
```

Esse tipo de progressão lógica nos nomes das funções — de um alto nível de abstração para um mais baixo e específico — torna o código mais fácil de seguir e ler. Considere a alternativa: se nosso nível mais alto de abstração for muito específico, acabaremos com um nome que tenta cobrir todas as bases, como `DetermineFileExtensionAndParseConfigurationFile`. Isso é horrivelmente difícil de ler; estamos tentando ser excessivamente específicos muito cedo e acabamos confundindo o leitor, apesar de tentar ser claro!

### Nomeação de Variáveis

Curiosamente, o oposto é verdadeiro para variáveis. Ao contrário das funções, nossas variáveis devem ser nomeadas de forma mais específica à medida que nos aprofundamos em escopos aninhados.

Você não deve nomear suas variáveis com base em seus tipos, assim como não nomearia seus animais de estimação como 'cachorro' ou 'gato'. – Dave Cheney

Por que nossos nomes de variáveis devem se tornar menos específicos à medida que viajamos mais fundo no escopo de uma função? Simplificando, à medida que o escopo de uma variável se torna menor, fica cada vez mais claro para o leitor o que essa variável representa, eliminando a necessidade de nomes específicos. No exemplo da função `fileExtension` anterior, poderíamos até encurtar o nome da variável `segments` para `s`, se quisermos. O contexto da variável é tão claro que não é necessário explicá-lo mais com nomes de variáveis mais longos. Outro bom exemplo disso é em loops `for` aninhados:

```go
func PrintBrandsInList(brands []BeerBrand) {
    for _, b := range brands { 
        fmt.Println(b)
    }
}
```

No exemplo acima, o escopo da variável `b` é tão pequeno que não precisamos gastar energia extra lembrando o que exatamente ela representa. No entanto, como o escopo de `brands` é um pouco maior, ajuda que seja mais específico. Ao expandir o escopo da variável na função abaixo, essa distinção se torna ainda mais evidente:

```go
func BeerBrandListToBeerList(beerBrands []BeerBrand) []Beer {
    var beerList []Beer
    for _, brand := range beerBrands {
        for _, beer := range brand {
            beerList = append(beerList, beer)
        }
    }
    return beerList
}
```

Ótimo! Esta função é fácil de ler. Agora, vamos aplicar a lógica oposta (ou seja, errada) ao nomear nossas variáveis:

```go
func BeerBrandListToBeerList(b []BeerBrand) []Beer {
    var bl []Beer
    for _, beerBrand := range b {
        for _, beerBrandBeerName := range beerBrand {
            bl = append(bl, beerBrandBeerName)
        }
    }
    return bl
}
```

Embora seja possível descobrir o que essa função está fazendo, a brevidade excessiva dos nomes das variáveis torna difícil seguir a lógica conforme viajamos mais fundo. Isso pode facilmente se transformar em uma confusão total, pois estamos misturando nomes de variáveis curtos e longos de forma inconsistente.

### Limpeza de Funções

Agora que conhecemos algumas boas práticas para nomear nossas variáveis e funções, bem como para esclarecer nosso código com comentários, vamos explorar algumas especificidades de como podemos refatorar funções para torná-las mais limpas.

### Comprimento da Função  
Qual deve ser o tamanho de uma função? Menor do que isso! – Robert C. Martin

Ao escrever código limpo, nosso objetivo principal é tornar o código facilmente digerível. A forma mais eficaz de fazer isso é manter nossas funções o mais curtas possível. É importante entender que não fazemos isso apenas para evitar a duplicação de código. A razão mais importante é melhorar a compreensão do código.

Pode ajudar olhar para a descrição de uma função de forma muito geral para entender melhor:

```go
fn GetItem:
    - analisar entrada JSON para o ID do pedido
    - obter usuário do contexto
    - verificar se o usuário tem a função apropriada
    - obter pedido do banco de dados
```

Ao escrever funções curtas (que geralmente têm de 5 a 8 linhas em Go), podemos criar código que lê quase tão naturalmente quanto a descrição acima:

```go
var (
    NullItem = Item{}
    ErrInsufficientPrivileges = errors.New("usuário não tem privilégios suficientes")
)

func GetItem(ctx context.Context, json []byte) (Item, error) {
    order, err := NewItemFromJSON(json)
    if err != nil {
        return NullItem, err
    }
    if !GetUserFromContext(ctx).IsAdmin() {
        return NullItem, ErrInsufficientPrivileges
    }
    return db.GetItem(order.ItemID)
}
```

Usar funções menores também elimina outro hábito terrível de escrever código: o inferno da indentação. O inferno da indentação geralmente ocorre quando uma cadeia de instruções `if` é descuidadamente aninhada em uma função. Isso torna muito difícil para os humanos entenderem o fluxo do código e deve ser eliminado sempre que for detectado. O inferno da indentação é particularmente comum ao trabalhar com `interface{}` e ao usar casting de tipo:

```go
func GetItem(extension string) (Item, error) {
    if refIface, ok := db.ReferenceCache.Get(extension); ok {
        if ref, ok := refIface.(string); ok {
            if itemIface, ok := db.ItemCache.Get(ref); ok {
                if item, ok := itemIface.(Item); ok {
                    if item.Active {
                        return Item, nil
                    } else {
                        return EmptyItem, errors.New("nenhum item ativo encontrado no cache")
                    }
                } else {
                    return EmptyItem, errors.New("não foi possível fazer cast da interface de cache para Item")
                }
            } else {
                return EmptyItem, errors.New("extensão não encontrada na referência do cache")
            }
        } else {
            return EmptyItem, errors.New("não foi possível fazer cast da interface de referência do cache para Item")
        }
    }
    return EmptyItem, errors.New("referência não encontrada no cache")
}
```

Primeiro, o inferno da indentação torna difícil para outros desenvolvedores entenderem o fluxo do seu código. Em segundo lugar, se a lógica em nossas instruções `if` se expandir, será exponencialmente mais difícil descobrir qual instrução retorna o quê (e garantir que todos os caminhos retornem algum valor). Outro problema é que essa profundidade de aninhamento de declarações condicionais força o leitor a rolar frequentemente e acompanhar muitos estados lógicos na cabeça. Isso também torna mais difícil testar o código e encontrar bugs, pois há muitas possibilidades diferentes aninhadas que você tem que considerar.

O inferno da indentação pode resultar em fadiga do leitor se um desenvolvedor tiver que analisar constantemente código difícil de manejar como o exemplo acima. Naturalmente, isso é algo que queremos evitar a todo custo.

Então, como limpamos essa função? Felizmente, é bastante simples. Em nossa primeira tentativa, vamos tentar garantir que estamos retornando um erro o mais rápido possível. Em vez de aninhar os `if` e `else`, queremos "empurrar nosso código para a esquerda", por assim dizer. Veja:

```go
func GetItem(extension string) (Item, error) {
    refIface, ok := db.ReferenceCache.Get(extension)
    if !ok {
        return EmptyItem, errors.New("referência não encontrada no cache")
    }

    ref, ok := refIface.(string)
    if !ok {
        // retornar erro de cast na referência 
    }

    itemIface, ok := db.ItemCache.Get(ref)
    if !ok {
        // retornar nenhum item encontrado no cache pela referência
    }

    item, ok := itemIface.(Item)
    if !ok {
        // retornar erro de cast na interface do item
    }

    if !item.Active {
        // retornar nenhum item ativo
    }

    return Item, nil
}
```

Depois de concluir nossa primeira tentativa de refatoração da função, podemos prosseguir para dividir a função em funções menores. Aqui está uma boa regra: Se o padrão `value, err :=` é repetido mais de uma vez em uma função, isso indica que podemos dividir a lógica do nosso código em partes menores:

```go
func GetItem(extension string) (Item, error) {
    ref, ok := getReference(extension)
    if !ok {
        return EmptyItem, ErrReferenceNotFound
    }
    return getItemByReference(ref)
}

func getReference(extension string) (string, bool) {
    refIface, ok := db.ReferenceCache.Get(extension)
    if !ok {
        return "", false
    }
    return refIface.(string), true
}

func getItemByReference(reference string) (Item, error) {
    item, ok := getItemFromCache(reference)
    if !item.Active || !ok {
        return EmptyItem, ErrItemNotFound
    }
    return item, nil
}

func getItemFromCache(reference string) (Item, bool) {
    if itemIface, ok := db.ItemCache.Get(reference); ok {
        return itemIface.(Item), true
    }
    return Item{}, false
}
```

Como mencionado anteriormente, o inferno da indentação pode dificultar o teste do nosso código. Quando dividimos nossa função GetItem em várias funções auxiliares, tornamos mais fácil rastrear bugs ao testar nosso código. Ao contrário da versão original, que consistia em várias instruções `if` no mesmo escopo, a versão refatorada de GetItem tem apenas dois caminhos de ramificação que precisamos considerar. As funções auxiliares também são curtas e digeríveis, tornando-as mais fáceis de ler.

Nota: Para código de produção, deve-se elaborar ainda mais o código retornando erros em vez de valores booleanos. Isso facilita a compreensão de onde o erro está originando. No entanto, como estas são apenas funções de exemplo, retornar valores booleanos será suficiente por agora. Exemplos de retorno de erros de forma mais explícita serão explicados em mais detalhes mais adiante.

Observe que a limpeza da função GetItem resultou em mais linhas de código no total. No entanto, o código agora está muito mais fácil de ler. Está organizado em uma estrutura em camadas, onde podemos ignorar "camadas" que não nos interessam e simplesmente descascar aquelas que queremos examinar. Isso facilita a compreensão da funcionalidade de baixo nível, pois só precisamos ler talvez de 3 a 5 linhas por vez.

Este exemplo ilustra que não podemos medir a limpeza do nosso código pelo número de linhas que ele usa. A primeira versão do código era certamente muito mais curta. No entanto, era artificialmente curta e muito difícil de ler. Na maioria dos casos, limpar o código inicialmente expandirá a base de código existente em termos de número de linhas. Mas isso é altamente preferível à alternativa de ter uma lógica confusa e bagunçada. Se você estiver em dúvida sobre isso, considere como você se sente em relação à seguinte função, que faz exatamente a mesma coisa que o nosso código, mas usa apenas duas linhas:

```go
func GetItemIfActive(extension string) (Item, error) {
    if refIface, ok := db.ReferenceCache.Get(extension); ok {
        if ref, ok := refIface.(string); ok {
            if itemIface, ok := db.ItemCache.Get(ref); ok {
                if item, ok := itemIface.(Item); ok {
                    if item.Active {
                        return item, nil
                    }
                }
            }
        }
    }
    return EmptyItem, errors.New("referência não encontrada no cache")
}
```

### Assinaturas de Função  
Criar uma boa estrutura de nomeação de função torna mais fácil ler e entender a intenção do código. Como vimos acima, fazer nossas funções mais curtas ajuda a entender a lógica da função. A última parte de limpar nossas funções envolve entender o contexto da entrada da função. Com isso vem outra regra fácil de seguir: Assinaturas de função devem conter apenas um ou dois parâmetros de entrada. Em alguns casos excepcionais, três podem ser aceitáveis, mas é aqui que devemos começar a considerar uma refatoração. Assim como a regra de que nossas funções devem ter apenas 5–8 linhas, isso pode parecer bastante extremo no início. No entanto, sinto que esta regra é muito mais fácil de justificar.

Pegue a seguinte função do tutorial de introdução do RabbitMQ à sua biblioteca Go:

```go
q, err := ch.QueueDeclare(
  "hello", // nome
  false,   // durável
  false,   // deletar quando não utilizado
  false,   // exclusivo
  false,   // sem espera
  nil,     // argumentos
)
```

A função QueueDeclare aceita seis parâmetros de entrada,

 o que é bastante. Com algum esforço, é possível entender o que esse código faz graças aos comentários. No entanto, os comentários são, na verdade, parte do problema—como mencionado anteriormente, eles devem ser substituídos por código descritivo sempre que possível. Afinal, não há nada que nos impeça de invocar a função QueueDeclare sem comentários:

```go
q, err := ch.QueueDeclare("hello", false, false, false, false, nil)
```

Agora, sem olhar para a versão comentada, tente lembrar o que os quarto e quinto argumentos `false` representam. É impossível, certo? Você inevitavelmente esquecerá em algum momento. Isso pode levar a erros custosos e bugs difíceis de corrigir. Os erros podem até ocorrer através de comentários incorretos—imagine rotular o parâmetro de entrada errado. Corrigir esse erro será insuportavelmente difícil, especialmente quando a familiaridade com o código tiver se deteriorado ao longo do tempo ou era baixa para começar. Portanto, recomenda-se substituir esses parâmetros de entrada por uma estrutura `Options`:

```go
type QueueOptions struct {
    Name string
    Durable bool
    DeleteOnExit bool
    Exclusive bool
    NoWait bool
    Arguments []interface{}
}

q, err := ch.QueueDeclare(QueueOptions{
    Name: "hello",
    Durable: false,
    DeleteOnExit: false,
    Exclusive: false,
    NoWait: false,
    Arguments: nil,
})
```

Isso resolve dois problemas: uso incorreto de comentários e rotulagem acidental incorreta das variáveis. Claro, ainda podemos confundir propriedades com o valor errado, mas nesses casos, será muito mais fácil determinar onde está nosso erro dentro do código. A ordenação das propriedades também não importa mais, portanto, a ordenação incorreta dos valores de entrada não é mais uma preocupação. O último benefício adicional dessa técnica é que podemos usar nossa estrutura QueueOptions para inferir os valores padrão dos parâmetros de entrada da nossa função. Quando estruturas em Go são declaradas, todas as propriedades são inicializadas com seu valor padrão. Isso significa que nossa opção QueueDeclare pode na verdade ser invocada da seguinte maneira:

```go
q, err := ch.QueueDeclare(QueueOptions{
    Name: "hello",
})
```

Os outros valores são inicializados com seu valor padrão de `false` (exceto para `Arguments`, que como uma interface tem um valor padrão de `nil`). Não só estamos muito mais seguros com essa abordagem, mas também somos muito mais claros com nossas intenções. Nesse caso, poderíamos realmente escrever menos código. Isso é um ganho geral para todos no projeto.

Uma nota final sobre isso: Não é sempre possível mudar a assinatura de uma função. Neste caso, por exemplo, não temos controle sobre a assinatura da função QueueDeclare porque ela é da biblioteca RabbitMQ. Não é nosso código, então não podemos alterá-lo. No entanto, podemos envolver essas funções para adequá-las aos nossos propósitos:

```go
type RMQChannel struct {
    channel *amqp.Channel
}

func (rmqch *RMQChannel) QueueDeclare(opts QueueOptions) (Queue, error) {
    return rmqch.channel.QueueDeclare(
        opts.Name,
        opts.Durable,
        opts.DeleteOnExit,
        opts.Exclusive,
        opts.NoWait,
        opts.Arguments,
    )
}
```

Basicamente, criamos uma nova estrutura chamada RMQChannel que contém o tipo `amqp.Channel`, que tem o método QueueDeclare. Em seguida, criamos nossa própria versão desse método, que basicamente apenas chama a versão antiga da função da biblioteca RabbitMQ. Nosso novo método tem todas as vantagens descritas anteriormente, e conseguimos isso sem ter que alterar nenhum código na biblioteca RabbitMQ.

Usaremos essa ideia de envolver funções para introduzir código mais limpo e seguro mais adiante ao discutir `interface{}`.

**Escopo de Variáveis**  
Agora, vamos dar um passo atrás e revisar a ideia de escrever funções menores. Isso tem outro efeito colateral agradável que não cobrimos no capítulo anterior: escrever funções menores pode tipicamente eliminar a dependência de variáveis mutáveis que vazam para o escopo global.

Variáveis globais são problemáticas e não pertencem a código limpo; elas tornam muito difícil para os programadores entenderem o estado atual de uma variável. Se uma variável é global e mutável, então por definição, seu valor pode ser alterado por qualquer parte da base de código. Em nenhum momento você pode garantir que essa variável terá um valor específico... E isso é uma dor de cabeça para todos. Este é mais um exemplo de um problema trivial que é exacerbado quando a base de código se expande.

Vamos ver um exemplo curto de como variáveis não globais com um grande escopo podem causar problemas. Essas variáveis também introduzem o problema do sombreamento de variáveis, como demonstrado no código retirado de um artigo intitulado "Golang scope issue":

```go
func doComplex() (string, error) {
    return "Success", nil
}

func main() {
    var val string
    num := 32

    switch num {
    case 16:
        // não fazer nada
    case 32:
        val, err := doComplex()
        if err != nil {
            panic(err)
        }
        if val == "" {
            // fazer algo mais
        }
    case 64:
        // não fazer nada
    }

    fmt.Println(val)
}
```

Qual é o problema com este código? À primeira vista, parece que o valor da variável `val` deve ser impresso como "Success" ao final da função `main`. Infelizmente, não é o caso. A razão para isso está na linha seguinte:

```go
val, err := doComplex()
```

Isso declara uma nova variável `val` no escopo do caso `32` do switch e não tem relação com a variável declarada na primeira linha de `main`. Claro, pode-se argumentar que a sintaxe de Go é um pouco complicada, o que eu não discordo necessariamente, mas há um problema muito pior em questão. A declaração de `var val string` como uma variável mutável e de escopo amplo é completamente desnecessária. Se fizermos uma refatoração muito simples, não teremos mais esse problema:

```go
func getStringResult(num int) (string, error) {
    switch num {
    case 16:
        // não fazer nada
    case 32:
        return doComplex()
    case 64:
        // não fazer nada
    }
    return "", nil
}

func main() {
    val, err := getStringResult(32)
    if err != nil {
        panic(err)
    }
    if val == "" {
        // fazer algo mais
    }
    fmt.Println(val)
}
```

Após nossa refatoração, `val` não é mais modificado, e o escopo foi reduzido. Novamente, lembre-se de que essas funções são muito simples. Uma vez que esse estilo de código se torna parte de sistemas maiores e mais complexos, pode ser impossível descobrir por que os erros estão ocorrendo. Não queremos que isso aconteça—não só porque geralmente não gostamos de erros de software, mas também porque é desrespeitoso para nossos colegas e para nós mesmos; estamos potencialmente desperdiçando o tempo uns dos outros tendo que depurar esse tipo de código. Os desenvolvedores precisam assumir a responsabilidade por seu próprio código, em vez de culpar esses problemas na sintaxe de declaração de variáveis de uma linguagem específica como Go.

A propósito, se a parte `// fazer algo mais` é outra tentativa de modificar a variável `val`, devemos extrair essa lógica para uma função autossuficiente, assim como a parte anterior. Dessa forma, em vez de expandir o escopo mutável de nossas variáveis, podemos simplesmente retornar um novo valor:

```go
func getVal(num int) (string, error) {
    val, err := getStringResult(num)
    if err != nil {
        return "", err
    }
    if val == "" {
        return NewValue() // função fictícia
    }
    return val, err
}

func main() {
    val, err := getVal(32)
    if err != nil {
        panic(err)
    }
    fmt.Println(val)
}
```

### Declaração de Variáveis

Além de evitar problemas com escopo e mutabilidade de variáveis, podemos melhorar a legibilidade declarando variáveis o mais próximo possível de seu uso. Em programação C, é comum ver a seguinte abordagem para declarar variáveis:

```go
func main() {
  var err error
  var items []Item
  var sender, receiver chan Item
  
  items = store.GetItems()
  sender = make(chan Item)
  receiver = make(chan Item)
  
  for _, item := range items {
    ...
  }
}
```

Isso sofre do mesmo sintoma descrito em nossa discussão sobre escopo de variáveis. Mesmo que essas variáveis possam não ser realmente reatribuídas em nenhum ponto, esse estilo de codificação mantém os leitores em alerta, de maneira errada. Assim como a memória do computador, a memória de curto prazo do nosso cérebro tem uma capacidade limitada. Ter que acompanhar quais variáveis são mutáveis e se um determinado fragmento de código vai ou não alterá-las torna mais difícil entender o que o código está fazendo. Descobrir o valor retornado eventualmente pode ser um pesadelo. Portanto, para facilitar isso para nossos leitores (e para nós mesmos no futuro), é recomendável declarar variáveis o mais próximo possível de seu uso:

```go
func main() {
	var sender chan Item
	sender = make(chan Item)

	go func() {
		for {
			select {
			case item := <-sender:
				// faça algo
			}
		}
	}()
}
```

No entanto, podemos fazer ainda melhor invocando a função diretamente após sua declaração. Isso torna muito mais claro que a lógica da função está associada à variável declarada:

```go
func main() {
  sender := func() chan Item {
    channel := make(chan Item)
    go func() {
      for {
        select { ... }
      }
    }()
    return channel
  }
}
```

E, voltando ao início, podemos mover a função anônima para torná-la uma função nomeada:

```go
func main() {
  sender := NewSenderChannel()
}

func NewSenderChannel() chan Item {
  channel := make(chan Item)
  go func() {
    for {
      select { ... }
    }
  }()
  return channel
}
```

Ainda está claro que estamos declarando uma variável, e a lógica associada ao canal retornado é simples, ao contrário do primeiro exemplo. Isso facilita a navegação pelo código e a compreensão do papel de cada variável.

Claro, isso não impede que possamos modificar nossa variável `sender`. Não há como declarar um struct const ou variáveis estáticas em Go. Isso significa que teremos que nos restringir de modificar essa variável em um ponto posterior do código.

NOTA: A palavra-chave `const` existe, mas é limitada a tipos primitivos apenas.

Uma maneira de contornar isso pode, ao menos, limitar a mutabilidade de uma variável ao nível do pacote. O truque envolve criar uma estrutura com a variável como uma propriedade privada. Essa propriedade privada só é acessível através de outros métodos fornecidos por essa estrutura de encapsulamento. Expandindo nosso exemplo de canal, isso seria algo assim:

```go
type Sender struct {
  sender chan Item
}

func NewSender() *Sender {
  return &Sender{
    sender: NewSenderChannel(),
  }
}

func (s *Sender) Send(item Item) {
  s.sender <- item
}
```

Agora garantimos que a propriedade `sender` do nosso struct `Sender` nunca seja modificada—pelo menos não de fora do pacote. Até o momento, esta é a única maneira de criar variáveis não primitivas publicamente imutáveis. É um pouco verboso, mas realmente vale o esforço para garantir que não acabemos com bugs estranhos resultantes de modificações acidentais de variáveis.

```go
func main() {
  sender := NewSender()
  sender.Send(&Item{})
}
```

Olhando para o exemplo acima, fica claro como isso também simplifica o uso de nosso pacote. Esse modo de ocultar a implementação é benéfico não apenas para os mantenedores do pacote, mas também para os usuários. Agora, ao inicializar e usar a estrutura `Sender`, não há preocupação com sua implementação. Isso abre uma arquitetura muito mais flexível. Como nossos usuários não estão preocupados com a implementação, estamos livres para alterá-la a qualquer momento, já que reduzimos o ponto de contato que os usuários têm com o pacote. Se não quisermos mais usar uma implementação de canal em nosso pacote, podemos facilmente mudar isso sem quebrar o uso do método `Send` (desde que mantenhamos a assinatura atual da função).

NOTA: Há uma explicação fantástica sobre como lidar com a abstração em bibliotecas de cliente, retirada da palestra AWS re:Invent 2017: Embracing Change without Breaking the World (DEV319).

### Go Limpo

Esta seção foca menos nos aspectos genéricos de escrever código Go limpo e mais nos específicos, com ênfase nos princípios subjacentes de código limpo.

### Valores Retornados

*Retornando Erros Definidos*

Vamos começar com uma maneira mais limpa de retornar erros. Como discutimos anteriormente, nosso principal objetivo ao escrever código limpo é garantir a legibilidade, testabilidade e manutenção do código. A técnica para retornar erros que discutiremos aqui atingirá todos esses objetivos com muito pouco esforço.

Vamos considerar a maneira normal de retornar um erro personalizado. Este é um exemplo hipotético retirado de uma implementação de mapa thread-safe que chamamos de `Store`:

```go
package smelly

func (store *Store) GetItem(id string) (Item, error) {
    store.mtx.Lock()
    defer store.mtx.Unlock()

    item, ok := store.items[id]
    if !ok {
        return Item{}, errors.New("item could not be found in the store") 
    }
    return item, nil
}
```

Não há nada inerentemente ruim sobre esta função quando a consideramos isoladamente. Olhamos para o mapa `items` do nosso struct `Store` para ver se já temos um item com o id fornecido. Se tivermos, retornamos; caso contrário, retornamos um erro. Bastante padrão. Então, qual é o problema com retornar erros personalizados como valores de string? Bem, vamos ver o que acontece quando usamos esta função dentro de outro pacote:

```go
func GetItemHandler(w http.ResponseWriter, r *http.Request) {
    item, err := smelly.GetItem("123")
    if err != nil {
        if err.Error() == "item could not be found in the store" {
            http.Error(w, err.Error(), http.StatusNotFound)
            return
        }
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    } 
    json.NewEncoder(w).Encode(item)
}
```

Isso não é tão ruim. No entanto, há um problema flagrante: Um erro em Go é simplesmente uma interface que implementa uma função (`Error()`) retornando uma string; assim, estamos agora codificando em nosso código o código de erro esperado, o que não é ideal. Essa string codificada é conhecida como uma string mágica. E seu principal problema é a flexibilidade: Se em algum momento decidirmos mudar o valor da string usado para representar um erro, nosso código quebrará (suavemente) a menos que o atualizemos em muitos lugares diferentes. Nosso código está fortemente acoplado—depende dessa string mágica específica e da suposição de que ela nunca mudará à medida que a base de código cresce.

Uma situação ainda pior ocorreria se um cliente usasse nosso pacote em seu próprio código. Imagine que decidimos atualizar nosso pacote e mudamos a string que representa um erro—o software do cliente agora quebraria repentinamente. Isso é algo que queremos evitar. Felizmente, a solução é muito simples:

```go
package clean

var (
    NullItem = Item{}

    ErrItemNotFound = errors.New("item could not be found in the store") 
)

func (store *Store) GetItem(id string) (Item, error) {
    store.mtx.Lock()
    defer store.mtx.Unlock()

    item, ok := store.items[id]
    if !ok {
        return NullItem, ErrItemNotFound
    }
    return item, nil
}
```

Ao representar o erro como uma variável (`ErrItemNotFound`), garantimos que qualquer pessoa que use este pacote possa verificar contra a variável em vez da string real que ela retorna:

```go
func GetItemHandler(w http.ResponseWriter, r *http.Request) {
    item, err := clean.GetItem("123")
    if err != nil {
        if errors.Is(err, clean.ErrItemNotFound) {
           http.Error(w, err.Error(), http.StatusNotFound)
            return
        }
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    } 
    json.NewEncoder(w).Encode(item)
}
```

Isso é muito mais agradável e também mais seguro. Alguns diriam que é mais fácil de ler também. No caso de uma mensagem de erro mais verbosa, certamente seria preferível para um desenvolvedor ler `ErrItemNotFound` do que um romance sobre por que um determinado erro foi retornado.

Essa abordagem não se limita a erros e pode ser usada para outros valores retornados. Como exemplo, também estamos retornando um `NullItem` em vez de `Item{}` como fizemos antes. Existem muitos cenários em que pode ser preferível retornar um objeto definido, em vez de inicializá-lo no retorno.

Retornar valores `NullItem` padrão como fizemos nos exemplos anteriores também pode ser mais seguro em certos casos. Por exemplo, um usuário de nosso

 pacote pode esquecer de verificar erros e acabar inicializando uma variável que aponta para um struct vazio contendo um valor padrão de `nil` como um ou mais valores de propriedade. Ao tentar acessar esse valor `nil` mais tarde no código, o software do cliente entraria em pânico. No entanto, ao retornarmos nosso valor padrão personalizado, podemos garantir que todos os valores que de outra forma seriam `nil` sejam inicializados. Assim, garantimos que não causamos pânicos no software dos nossos usuários.

Isso também nos beneficia. Considere o seguinte: Se quisermos alcançar a mesma segurança sem retornar um valor padrão, teríamos que mudar nosso código em todos os lugares onde retornamos esse tipo de valor vazio. No entanto, com nossa abordagem de valor padrão, agora só precisamos alterar nosso código em um único lugar:

```go
var NullItem = Item{
    itemMap: map[string]Item{},
}
```

NOTA: Em muitos cenários, invocar um pânico será realmente preferível para indicar que está faltando uma verificação de erro.

NOTA: Cada propriedade de interface em Go tem um valor padrão de `nil`. Isso significa que isso é útil para qualquer struct que tenha uma propriedade de interface. Isso também é verdadeiro para structs que contêm canais, mapas e slices, que também podem ter um valor `nil`.

*Retornando Erros Dinâmicos*

Há certamente alguns cenários onde retornar uma variável de erro pode não ser viável. Em casos onde a informação em erros personalizados é dinâmica, se quisermos descrever eventos de erro mais especificamente, não podemos mais definir e retornar nossos erros estáticos. Aqui está um exemplo:

```go
func (store *Store) GetItem(id string) (Item, error) {
    store.mtx.Lock()
    defer store.mtx.Unlock()

    item, ok := store.items[id]
    if !ok {
        return NullItem, fmt.Errorf("Could not find item with ID: %s", id)
    }
    return item, nil
}
```

Então, o que fazer? Não há um método bem definido ou padrão para lidar e retornar esses tipos de erros dinâmicos. Minha preferência pessoal é retornar uma nova interface, com um pouco de funcionalidade adicional:

```go
type ErrorDetails interface {
    Error() string
    Type() string
}

type errDetails struct {
    errtype error
    details interface{}
}

func NewErrorDetails(err error, details ...interface{}) ErrorDetails {
    return &errDetails{
        errtype: err,
        details: details,
    }
}

func (err *errDetails) Error() string {
    return fmt.Sprintf("%v: %v", err.errtype, err.details)
}

func (err *errDetails) Type() error {
    return err.errtype
}
```

Essa nova estrutura de dados ainda funciona como nosso erro padrão. Podemos ainda compará-la com `nil` já que é uma implementação de interface, e ainda podemos chamar `.Error()` nela, então não quebrará implementações existentes. No entanto, a vantagem é que agora podemos verificar o tipo do erro como podíamos anteriormente, apesar de nosso erro agora conter detalhes dinâmicos:

```go
func (store *Store) GetItem(id string) (Item, error) {
    store.mtx.Lock()
    defer store.mtx.Unlock()

    item, ok := store.items[id]
    if !ok {
        return NullItem, NewErrorDetails(
            ErrItemNotFound,
            fmt.Sprintf("could not find item with id: %s", id))
    }
    return item, nil
}
```

E nossa função de manipulador HTTP pode então ser refatorada para verificar um erro específico novamente:

```go
func GetItemHandler(w http.ResponseWriter, r *http.Request) {
    item, err := clean.GetItem("123")
    if err != nil {
        if errors.Is(err.Type(), clean.ErrItemNotFound) {
            http.Error(w, err.Error(), http.StatusNotFound)
            return
        }
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    } 
    json.NewEncoder(w).Encode(item)
}
```

### Valores Nil

Um aspecto controverso de Go é a adição de `nil`. Esse valor corresponde ao valor `NULL` em C e é essencialmente um ponteiro não inicializado. Já vimos alguns dos problemas que `nil` pode causar, mas para resumir: As coisas quebram quando você tenta acessar métodos ou propriedades de um valor `nil`. Portanto, é recomendável evitar retornar um valor `nil sempre que possível. Dessa forma, os usuários do nosso código são menos propensos a acessar valores `nil` acidentalmente.

Existem outros cenários em que é comum encontrar valores `nil` que podem causar algum sofrimento desnecessário. Um exemplo disso é inicializar incorretamente um struct (como no exemplo abaixo), o que pode levar a ele conter propriedades `nil`. Se acessadas, essas propriedades `nil` causarão um pânico.

```go
type App struct {
   Cache *KVCache
}

type KVCache struct {
  mtx sync.RWMutex
  store map[string]string
}

func (cache *KVCache) Add(key, value string) {
  cache.mtx.Lock()
  defer cache.mtx.Unlock()
  
  cache.store[key] = value
}
```

Este código está absolutamente correto. No entanto, o perigo é que nosso `App` pode ser inicializado incorretamente, sem inicializar a propriedade `Cache`. Se o seguinte código for invocado, nossa aplicação entrará em pânico:

```go
app := App{}
app.Cache.Add("panic", "now")
```

A propriedade `Cache` nunca foi inicializada e, portanto, é um ponteiro `nil`. Assim, invocar o método `Add` como fizemos aqui causará um pânico, com a seguinte mensagem:

```
panic: runtime error: invalid memory address or nil pointer dereference
```

Em vez disso, podemos transformar a propriedade `Cache` de nossa estrutura `App` em uma propriedade privada e criar um método tipo getter para acessá-la. Isso nos dá mais controle sobre o que estamos retornando; especificamente, garante que não estamos retornando um valor `nil`:

```go
type App struct {
   cache *KVCache
}

func (app *App) Cache() *KVCache {
  if app.cache == nil {
      app.cache = NewKVCache()
   }
   return app.cache
}
```

O código que anteriormente causava pânico agora será refatorado para o seguinte:

```go
app := App{}
app.Cache().Add("panic", "now")
```

Isso garante que os usuários de nosso pacote não tenham que se preocupar com a implementação e se estão usando nosso pacote de maneira insegura. Tudo o que eles precisam se preocupar é em escrever seu próprio código limpo.

NOTA: Existem outros métodos para alcançar um resultado seguro semelhante. No entanto, acredito que este é o método mais direto.