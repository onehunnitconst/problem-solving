import sys

input = sys.stdin.readline
print = lambda x: sys.stdout.write(str(x) + '\n')

len = int(input())
s = input()[0:len]

r = 31
m = 1234567891

sum = 0

for index in range(len):
  id = ord(s[index]) % 96
  coef = pow(r, index)
  sum += id * coef % m

result = sum % m

print(result)