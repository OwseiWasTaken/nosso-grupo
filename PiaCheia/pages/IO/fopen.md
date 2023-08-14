# fopen #;

## C`#include <stdio.h>` ##;
## C`FILE *fopen(const char *caminho_do_arquivo, const char *modo_do_arquivo);` ##;

### A função _fopen()_ abre o arquivo indicato pela string _caminho_do_arquivo_ e retorna um ponteiro para um manipulador de arquivo ###;
### O argumento _modo_do_arquivo_ define como o arquivo pode ser utilizado ###;


## r ## _read_;
Abre o arquivo para leitura. O manipulador se posiciona no início do arquivo.


## r+ ## _read+write_;
Abre o arquivo para leitura e escrita. O manipulador se posiciona no início do
arquivo, possivelmente subustituindo informação ao escrever.


## w ## _write_;
Apaga o conteúdo do arquivo e o abre para escrita. O manipulador se posiciona
no início do arquivo.


## w+ ## _write+read_;
Apaga o conteúdo do arquivo caso ele já exista, caso contrário, cria um arquivo
para escrita e leitura. O manipulador se posiocina no início do arquivo.


## a ## _append_;
Abre ou cria um arquivo para escrita. O manipulador se posiciona no final do arquivo.


## a+ ## _append+read_;
Abre ou cria um arquivo para escrita e leitura. O manipulador sempre escreve no final do arquivo.
Não está definido onde o manipulador deve se posicionar.
Em sistemas baseados em Unix, o manipulador se posiciona no início do arquivo.


O _modo_do_arquivo_ pode incluir a letra 'b' no final da string `"rb+"`, isso define
que o arquivo será utilizado como informação binária, e não como textual.
Sistemas baseados em Unix não mudam como o arquivo funciona, porém
geralmente é utilizado para maior claridade e portabilidade

