n = gets.to_i
points = []

for index in 0..n-1 do
  x, y = gets.split.map(&:to_i)
  points.push([x, y])
end

sorted_points = points.sort{|a, b| if a[1] == b[1] then a[0] <=> b[0] else a[1] <=> b[1] end}

sorted_points.each do |point|
  puts "#{point[0]} #{point[1]}"
end