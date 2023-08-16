# fsync #

{#include
C``
## #include <unistd.h> ##;
## void sync(void); ##;
## int syncfs(int fd); ##;
``
}

{#summary
### A função _sync()_ executa todas as alterações de arquivos pendentes no sistema operacional ###;
### A função _sync()_ não tem como falhar ###;
### A função _syncfs()_ executa todas as alterações de arquivos que pendentes no sistema de arquivos em que o arquivo referido por _fd_ está ###;
}

{#return
### Retorno: ### _int_;
A função _syncfs()_ retorna 0 quando os arquivos são alterados corretamente.
Caso contrário -1 é retornado, e _errno_(indicador de erro universal) é definido apropriadamente
}

{#errors
### Erros: ###;
_EBADFD_ _fd_ Não é uma referência válida a um arquivo;
_EIO_ Um erro ocorreu ao tentar utilizar a mídia física que seria modificada;
_ENOSPC_ não há espaço suficiente para guardar todas as alterações;
_ENOSPC_, _EDQUOT_ não há espaço suficiente para guardar o arquivo (este erro apenas existe em servidores NFS);
}

