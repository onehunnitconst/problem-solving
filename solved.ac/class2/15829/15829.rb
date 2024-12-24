len = gets.to_i
str = gets.slice(0, len)

r = 31
m = 1234567891

sum = 0

for index in 0..len-1 do
  id = str[index].ord % 96
  coef = r ** index
  sum = sum + (id * coef % m)
end

result = sum % m

puts result