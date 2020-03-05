#!usr/local/bin/ruby -w
4.times do puts "Hooray" end
4.times {puts "Hooray"}
1.upto(10) do |item|
	puts item
end

5.downto(1){|num| puts num}

0.step(100,10){|num| puts num}
