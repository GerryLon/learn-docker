#include <stdio.h>
#include <unistd.h>
#include <stdlib.h>

int main() {
  // 相应目录必须已经存在, 否则会报错: chroot: No such file or directory
  const char* tmpRoot = "/opt/soft";
  char *sret;

  // 要先调chdir才行, 否则进程没有"当前目录"
  // 参看chroot的man 手册:
  // If the program is not currently running with an altered root directory, it should be noted that
  // chroot() has no effect on the process's current directory.
  if (chdir(tmpRoot) != 0)  {
    perror("chdir");
    exit(1);
  }

  if (chroot(tmpRoot) != 0) {
    perror("chroot");
    exit(1);
  }
  
  // If buf is NULL, space is allocated as necessary to store the pathname and size is ignored.  
  // This space may later be free(3)'d.
  sret = getcwd(NULL, 0);
  if (!sret) {
    perror("getcwd");
    exit(1);
  }
  printf("cwd: %s\n", sret);
  free(sret);

  return 0;
}
