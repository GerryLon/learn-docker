#include <stdio.h>
#include <unistd.h>

// 返回当前可执行文件的绝对路径
// 如当前执行为 ./readlink.out 则返回为/data/.../readlink.out 这样的绝对路径
int main() {
	char buf[128];
	readlink("/proc/self/exe", buf, sizeof(buf)/sizeof(char));
	printf("%s\n", buf);

	return 0;
}
