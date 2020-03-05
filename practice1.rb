#!usr/local/bin/ruby -w

print "Enter a number:"
num = gets.to_i

puts "Positive" if num>0

puts "Negative" if num<0

puts "Zero" if num==0

puts "Divisible by 3" if num%3==0
puts "Divisible by 5" if num%5==0
puts "Millions" if num.between?(1000000,19999999)
puts "2 to the 10 power" if num == 2**10
puts num.to_s.length
if num.to_s.length > 1
  puts num.to_s[1]
else
  puts "Number is not 2 digits long"


end
