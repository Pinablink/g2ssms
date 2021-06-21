# G2ssms

## Apresentação
Biblioteca de suporte a aplicações escritas em Golang, que permite envio de SMS. Essa biblioteca trabalha utilizando um servidor de envio de SMS. Pode ser um envio individual ou massivo. Portanto sua utilização é dependente da assinatura dessa plataforma, onde por um valor você adquiri um pacote de mensagens. Com a assinatura você tem acesso a um Painel Admin, onde pode enviar e gerenciar sua mensagem por esse meio. Além desse painel, existe uma API onde qualquer aplicação pode consumir esse recurso.  

## A biblioteca
Pensado para proporcionar capacidade de comunicação do sistema responsável pelo cadastro das visitas do Cliente GGIZ. Foi adotado esse serviço para envio de notificação ao administrador via SMS, com um "ticket" informativo de que a funcionalidade executou. A empresa https://www.kingsms.com.br/, disponibiliza essa infra-estrutura necessária e um microserviço. O G2ssms é uma abstração para tornar o desenvolvimento mais rápido e disponibilizar essa capacidade as aplicações Golang. Seja a aplicação especifica que lhe deu origem, assim como outras que tiver esse requisito.

👀 Você pode dar uma olhada na documentação que a empresa disponibiliza 👉 https://kingsms.docs.apiary.io/#reference/enviar-sms


## Um exemplo a seguir de implementação
Para utilizar essa abstração em seu código Golang, adicione o pacote em seu ambiente.
<br>
$ go get github.com/Pinablink/g2ssms

Vamos ver então um pouco de código. Logo abaixo a implementação necessária para o envio de um SMS por chamada. Os atributos necessários para o consumo desse microserviço, como boa prática, foi disponibilizado como Variaveis de Ambiente.

```
package main

import (
	"os"
	"fmt"
    "g2ssms/send"  
)

func main() {
    
	
	// Criação do Objeto para envio de sms
	testeObSms := &send.SSendSMS{}
	testeObSms.UrlService = os.Getenv("SSMS_URL_SERVICE")
	testeObSms.Acao = send.SendSms
	testeObSms.Login = os.Getenv("SSMS_LOGIN_SERVICE")
	testeObSms.Token = os.Getenv("SSMS_TOKEN_SERVICE")
	testeObSms.Numero = os.Getenv("SSMS_NUM_DEST_SERVICE")
	testeObSms.Msg = send.SMsg{
		Msg: "Teste de Uso da Lib; DEV: Weber(Pinablink); Mensagem de SMS Oh, Oh!;",
	}

	idResponse, err := testeObSms.SendSMS()

	if err != nil {
		panic(err)
	}

	fmt.Println(idResponse)
}

```
Existem algumas outras execuções possiveis. No entanto a implementação dessa solução disponível, é apenas para enviar um SMS quando uma instância de ***send.SSendSMS*** executar a ***func SMsg()*** .
<br>
Caso ache interessante, esteja a vontade em usar essa solução em seus códigos. 
<br>
E lembre-se essa solução esta consumindo um microserviço disponibilizado pela empresa https://www.kingsms.com.br/.
<br>
