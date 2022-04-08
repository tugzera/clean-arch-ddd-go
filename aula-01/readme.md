# Cenário

Vamos implementar um sistema de vendas online com a possibilidade de realizar pedidos com múltiplos itens, cada um deles com uma quantidade variável, calculando o frete, os impostos, aplicando um cupom de desconto e ainda interagindo com o estoque. Além disso teremos ainda fluxos de pagamento e cancelamento do pedido realizado.


# Testes
1 - Não deve fazer um pedido com cpf inválido OK
2 - Deve fazer um pedido com 3 itens (com descrição, preço e quantidade) OK
3 - Deve fazer um pedido com cupom de desconto (percentual sobre o total do pedido)


# Considere

Refatorar o algoritmo de validação de cpf (OK)

# Sugestões

Utilize a sua linguagem e biblioteca de teste de sua preferência
Faça a modelagem da forma que desejar e não se preocupe por enquanto, vamos implementar juntos na aula seguinte com influências de DDD e Clean Architecture

Dicas
Devem existir no mínimo 2 arquivos, um de teste e outro que implementa os cenários propostos
Tente seguir com disciplina, criando primeiro um teste que falha, depois fazendo e teste passar e refatorando