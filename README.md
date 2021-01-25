# hostgator-challenge
This is my repository for hosting the Hostgator Challenge.
=======
# HostGator Challenging

HostGator Latin America

## Installation

Clone the repository: 
```bash
git clone git@github.com:guther/hostgator-challenge.git
```
Enter the directory and do the command:

```bash
docker-compose up
```
  

## Usage

This will create the necessary containers for the application to work.

At the end, the application will be active at the address [http://localhost:6060](http://localhost:6060), however it is necessary to authenticate beforehand. Otherwise, you will receive error 401 (Not authenticated).

Authenticate sending a POST request to [http://localhost:6060/login](http://localhost:6060/login) along with the user data:

`body: '{" password ":"@#$RF@!718"," username ":" admin "}`

You will receive the token (JWT).
Use this token to perform requests.

## Requisitos de Implementação

- JWT
- Cache com MySQL
- RESTful
- Golang (bonus :-) .. using Gin framework!!! :P
- Private repository on GitHub
- Several commits during the development
- Unit Tests
- Database Schema using migrations
- CI Pipeline using Jenkins

Agora eu irei passar por cada item acima e explicar como a aplicação foi desenvolvida. Okay?

Ao tocar conhecimento dos itens da aplicação pesquisei como poderia facilitar a vida do meu avaliador, de tal forma que no arranque da aplicação já fosse possível ter tudo configurado. 
E acredite, essa aplicação é assim.

# JWT

Como já informei, a aplicação usa JWT para efetuar a autenticação.
![JWT](https://raw.githubusercontent.com/guther/hostgator-challenge/dev/images/postmantoken.jpg?token=AEWA2JO67DZSSGUL4TUSFPDABX67E)
O token está configurado para explicar em 48 horas e tomou-se cuidado para usar os códigos corretos do HTTP, tal como o retorno do código 401 em vez do 403, pois trata-se de uma autenticação, cuja autorização está concedida mediante login. 

# Cache 

Usei o MySQL como solicitado. O esquema do banco de dados foi feito usando migration, e embora tenha optado por não usar nenhum framework de ORM, a base de dados está consistente e tolerante a falhas. Isto é: (Agora vai uma dica de teste para vocês fazerem com os meus concorrentes) Quando o esquema da base de dados já existe e você muda a senha do banco, ocorre o erro de autenticação que é o velho conhecido dos desenvolvedores, né?! Mas, e quando você já altera a senha do banco antes mesmo da migração ocorrer? Hummmm... é um outro erro. E a aplicação tem que ser tolerante a isso para verificar o momento que poderá criar a base de dados e o  funcionamento do cache passará a ocorrer.
![Schema](https://raw.githubusercontent.com/guther/hostgator-challenge/dev/images/migrations.jpg?token=AEWA2JPFPRPR5AN6AHEABYDABYAKO)

# RESTful
Com as boas práticas, não usei nomes de verbos nas rotas e usei os códigos HTTP de acordo com o desejado. Não tive necessidade de tratar os métodos PUT e DELETE, pois a aplicação só possui GET e POST, conforme enunciado.

# Golang
Eu nunca tinha usado o Golang na minha vida! Acredita? E não tenho medo algum de dizer isso, é super sério! Mas, como todos nós sabemos que quem domina os paradigmas aprende rápido (e eu já fui monitor de disciplina de Algoritmo na faculdade), foi exatamente isso que aconteceu. Confesso que amei o funcionamento do panic, defer e recover e somado ao uso do middleware gin.Recovery() ficou sensacional.
Usei o framework Gin nessa aplicação.
Aprendi muito nesses 5 dias, desde que recebi essa prova na Terça-Feira. 

# Private repository on GitHub
Aqui é mais uma dica de avaliação dos meu concorrentes, pois com o repositório privado aumenta o nível de dificuldade devido a necessidade de troca de informação entre os containers e o host. Há complicação de SSL, problema de keychain com o .ssh/id_rsa, etc. Só não vale retirar o repositório do modo privado e nem expor as credenciais para todo mundo ver. (Tokens? pessoais?! Não!)

# Several commits during the development
Eu acho que fiz bastante commit para 5 dias, hein?
Os commits estão aí, todos com mensagens e cada um com sua feature.

# Unit Tests
Esse aqui foi uma das melhores partes! O Golang ajuda muito a criar testes unitários com o ``` go test ```.
O Jenkins ficou responsável por executar os testes e podemos visualizar o Pipeline em execução.

![pipeline](https://raw.githubusercontent.com/guther/hostgator-challenge/dev/images/pipeline.jpg?token=AEWA2JNSFH47EQAX64QBHI3ABYFLI)

Esse job também não é criado manualmente, é automático.
Usei o plugin Job DSL para fazer isso. Veja o formato SDL, é outro formato.

![jsdl](https://raw.githubusercontent.com/guther/hostgator-challenge/dev/images/jdsl.jpg?token=AEWA2JMRY7PCL7XP5U6GIJ3ABYFXI)

Esse código sdl cria o **job-generator**, que é apenas uma seed para criar os reais jobs no Jenkins, como o job CI, por exemplo.

# Jenkins

Para pular todas as etapas de configuração do **Jenkins**, usei a prática do **JCasC - Jenkins Configuration as Code** e não temos que passar pela criação de usuário e senha, criação de jobs, plugins, etc. Ao acessar a página [http://localhost:8080](http://localhost:8080) você já cai na tela de login. 
![cockerfile](https://raw.githubusercontent.com/guther/hostgator-challenge/dev/images/dockerfile.jpg?token=AEWA2JMDL3YZVSTO6A57TMDABYCCY)
Para obter isso, usei o jenkins-plugin-cli direto no Dockerfile em conjunto com a configuração do casc.yaml compartilhada pra imagem do Jenkins no momento do up do Composer.
![casc](https://raw.githubusercontent.com/guther/hostgator-challenge/dev/images/casc.yaml.jpg?token=AEWA2JNWNH42ORAOEEJCXQTABYCKI)


Informe as credenciais: **admin** e **password**.

Então você irá ver dois jobs criados no momento da inicialização.
![Jobs](https://raw.githubusercontent.com/guther/hostgator-challenge/dev/images/jobs.jpg?token=AEWA2JNX64OBW22PHXLQLWLABYARY)
O CI é o job responsável pelo Pipeline da Integração Contínua com as etapas de testes. Ele realiza o clone do repositório, roda os testes e armazena o log.


## License
[MIT](https://choosealicense.com/licenses/mit/)
