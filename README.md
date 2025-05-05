## Construindo um blog com GO

Construção do zero da estrutura de um blog em Go com suporte a arquivos Markdown, abordando os seguintes pontos:

    - Criação de um servidor HTTP simples com a net/http;
    - Organização de rotas e handlers em arquivos separados;
    - Uso de html/template para renderizar páginas dinâmicas;
    - Criação de uma estrutura de templates reutilizáveis (base.html);
    - Conversão de arquivos .md em HTML usando a biblioteca Goldmark;
    - Carregamento dinâmico de posts escritos em .yaml, com campos de metadados e conteúdo em Markdown;
    - Implementação de handlers para listar posts e exibir posts individuais com base no slug da URL.

Com essa estrutura, você tem um blog funcional e simples, com suporte a Markdown e fácil de manter — especialmente útil para projetos pessoais.
