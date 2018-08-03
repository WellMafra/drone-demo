
# Executando o drone
Esse documento possui instruções de como configurar Drone usando docker-compose.
As configurações do drone estão dentro da pasta **drone-ci**

## Configurações das variaveis de ambiente
* Renomeie o arquivo _.env-example_ para _.env_
* Altere os valores conforme suas configurações

## Configurando banco de dados
Siga essas instruções apenas se você deseja usar um banco de dados que não é o padrão do drone, no caso _sllite_

* Primeiro configure o postgres para permitir conexão remota: [link](https://blog.bigbinary.com/2016/01/23/configure-postgresql-to-allow-remote-connection.html)
* Você deve criar o banco, pois o drone não cria automaticamente
* `create table drone`
* No arquivo *.env*, você não deve colocar o ip do banco como _localhost_ ou _127.0.0.1_, pois a rede do docker não é a mesma que a da sua máquina. Leia mais: [Docker network](https://docs.docker.com/compose/networking/)
* Você deve colocar o ip da sua máquina local ou da máquina onde está instalado o banco de dados.

## Configurando acesso externo do drone
Siga essas instruções apenas se você deseja rodar o drone na sua máquina local

* Para você autenticar o drone com o github, você precisa ter um ip fixo ou uma forma de acesso externo onde o github possa acessar a sua máquina.
* Uma das formas é criando _proxy reverso_, escrevi um artigo ensinando o passo-a-passo: [ngRok](https://medium.com/@WellMafra/ngrok-crie-t%C3%BAneis-para-o-seu-localhost-b9142011b972)
* Pegue o link gerado e coloque na variável _DRONE_HOST_ do arquivo _.env_

## Iniciando os containers 
Para iniciar o drone com docker-compose
* `docker-compose up -d`

Para iniciar o drone com docker-compose
* `docker-compose up -d`

## Parando a execução
Finalize a execução dos container
* `docker-compose stop `

## Pare e remova os containers
Pare e remova os container, networks, volumes, and images criados por ele
* `docker-compose down `

## Acessar os logs dos containers
Para acompanhar os logs na tela do terminal
* `docker-compose logs -f`