#!usr/local/bin/ruby -w

print "Enter your name:"
name = gets.chomp
print "Enter your age:"
age = gets.to_i
puts "#{name} is #{12*age} months old"
