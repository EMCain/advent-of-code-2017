puts('hello')


def find_checksum(filename)
  sum = 0
  IO.foreach(filename) do | line |
    nums = line.strip.split(' ')
    nums = nums.map { |e| e.to_i }
    sum += nums.max - nums.min
  end
  return sum
end

def find_divisible(filename)
  sum = 0
  IO.foreach(filename) do | line |
    nums = line.strip.split(' ')
    nums = nums.map { |e| e.to_i }

    quotient = 0
    while quotient == 0 do
      for numerator in nums
        for denominator in nums
          if numerator != denominator
            if numerator % denominator == 0
              quotient = numerator / denominator
              break
            end
          end
        end
      end
    end
    sum += quotient
  end
  return sum
end
puts "sum for test: " + find_checksum('./test_inputs/day_02_01.txt').to_s
puts "sum for input: " + find_checksum('./inputs/day_02.txt').to_s

puts "divisible for test: " + find_divisible('./test_inputs/day_02_02.txt').to_s
puts "divisible for input: " + find_divisible('./inputs/day_02.txt').to_s
