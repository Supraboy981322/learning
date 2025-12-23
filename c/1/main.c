#include <stdio.h>
#include <string.h>

int main(int argc, char *argv[]) {
  if (argc < 2) {
    printf("not enough args\n");
    return 1;
  }

  for (int i = 1; i < argc; i++) {
    argv[i-1] = argv[i];
  };argc--;

  for (int i = 0; i < argc; i++) {
    printf("%s ", argv[i]);
  };printf("\n");

  return 0;
}
