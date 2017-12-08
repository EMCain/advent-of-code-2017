def get_parents(filename)
  hash = {}
  IO.foreach(filename) do | line |
    match = /(?<val>\w+)(?: \(\d+\) (->) (?<keys>.*))?/.match(line.strip)
    if !match.named_captures['keys'].nil?
      keys = match.named_captures['keys'].split(', ')
      val = match.named_captures['val']
      for key in keys
        hash[key] = val
      end
    end
  end
  return hash
end


def get_bottom(filename)
  hash = get_parents(filename)
  children = hash.keys
  parents = hash.values
  return (parents - children).uniq
end

puts get_parents('test_inputs/day_07.txt')
puts get_bottom('test_inputs/day_07.txt')

puts get_parents('inputs/day_07.txt')
puts get_bottom('inputs/day_07.txt')
