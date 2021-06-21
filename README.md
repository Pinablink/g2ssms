# G2ssms

## Apresentação
Biblioteca de suporte a aplicações escritas em Golang, que permite envio de SMS. Essa biblioteca trabalha utilizando um servidor de envio de SMS. Pode ser um envio individual ou massivo. Portanto sua utilização é dependente da assinatura dessa plataforma, onde por um valor você adquiri um pacote de mensagens. Com a assinatura você tem acesso a um Painel Admin, onde pode enviar e gerenciar sua mensagem por esse meio. Além desse painel, existe uma API onde qualquer aplicação pode consumir esse recurso.  

## A biblioteca
Pensado para proporcionar capacidade de comunicação ao sistema responsável pelo cadastro das visitas do Cliente GGIZ. Foi adotado esse serviço para envio de notificação ao administrador via SMS. A empresa https://www.kingsms.com.br/, disponibiliza essa infra-estrutura necessária e um microserviço. O G2ssms é uma abstração para tornar o desenvolvimento mais rápido e disponibilizar essa capacidade as aplicações Golang. Seja a aplicação especifica que lhe deu origem, assim como outras.
