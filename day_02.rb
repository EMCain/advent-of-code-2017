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

puts("sum for test: " + find_checksum('./test_inputs/day_02_01.txt').to_s)
puts("sum for input: " + find_checksum('./inputs/day_02.txt').to_s)
