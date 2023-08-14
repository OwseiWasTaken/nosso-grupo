# fclose #

{
## C`#include <stdio.h>` ##;
## C`int fclose(FILE *stream);` ##;
}

{
### A função _fclose()_ fecha o arquivo indicado pelo pointeiro do manipulador de arquivo _stream_. ###;
### A função *não* deve ser chamada com um pointeiro inválido, ou com um mainpulador de um arquivo que já foi fechado. ###;
### *Nenhuma* ação deve ocorrer com um arquivo após ele ser fechado por _fclose()_. ###;
}

{
### Retorno: ### _int_;
fclose() retorna 0 quando o arquivo é fechado corretamente.;
caso contrário _EOF_(fim do arquivo) é retornado, e _errno_ é definido
}

{
### Erros: ###;
_EBADF_ o manipulador da _stream_ é inválido;
_EIO_ um erro no hardware ocorreu;
_EINTR_ um sinal do sistema operacional interrompeu a função;
_ENOSPC_, _EDQUOT_ não há espaço suficiente para guardar o arquivo (este erro apenas existe em servidores NFS)
}

{
### Notas: ###;
É recomendado rodar [IO.sync](../IO/sync.html) após uma call de fclose() que falhou, para garantir que o mínimo de informação seja perdido
}

