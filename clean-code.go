package cleancode

// Código Go Limpo

// Introdução ao código limpo: promove software legível e manutenível.

// Desenvolvimento Orientado a Testes
// Escreva testes frequentemente e refatore o código. Exemplo de uma função pequena e específica:
func Add(a int, b int) int {
    return a + b
}

// Convenções de Nomenclatura
// Nomes de funções devem ser específicos e descritivos. Exemplo de boas práticas na nomeação:
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

// Comentários
// Comentários devem explicar o "porquê" e não o "como". Exemplo ruim e bom:
for i := 0; i < 10; i++ {
    doSomething(i) // Comentário tutorial, desnecessário
}

// Melhor abordagem, explicando o "porquê":
for workerID := 0; workerID < 10; workerID++ {
    instantiateThread(workerID)
}

// Nomeação de Funções
// Nomes de funções devem ser descritivos e não excessivamente específicos. Exemplo de boa nomeação:
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

// Nomeação de Variáveis
// Nomes de variáveis devem ser mais específicos em escopos maiores. Exemplo de boas práticas:
func PrintBrandsInList(brands []BeerBrand) {
    for _, b := range brands { 
        fmt.Println(b)
    }
}

// Exemplo ruim de nomeação de variáveis:
func BeerBrandListToBeerList(b []BeerBrand) []Beer {
    var bl []Beer
    for _, beerBrand := range b {
        for _, beerBrandBeerName := range beerBrand {
            bl = append(bl, beerBrandBeerName)
        }
    }
    return bl
}

// Limpeza de Funções
// Manter funções curtas e legíveis. Exemplo de função refatorada para evitar "inferno da indentação":
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

// Assinaturas de Função
// Funções devem ter poucos parâmetros de entrada. Exemplo de uso de uma estrutura para simplificar:
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
