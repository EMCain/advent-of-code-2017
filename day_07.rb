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

test_parents = get_parents('test_inputs/day_07.txt')
# puts get_bottom('test_inputs/day_07.txt')
#
parents = get_parents('inputs/day_07.txt')
# puts get_bottom('inputs/day_07.txt')

# part 2

def get_total_weight(filename, parents)
  hash = {}
  IO.foreach(filename) do | line |
    match = /(?<disc>\w+) \((?<weight>\d+)\)/.match(line.strip)
    if !match.nil?
      p match.named_captures
      name = match.named_captures['disc']
      weight = match.named_captures['weight'].to_i
      while parents.keys.include? name
        if !hash.keys.include? name
          hash[name] = 0
        end
        hash[name] += weight
        puts name + " weighs " + hash[name].to_s

        name = parents[name]

      end
    end
  end
  return hash
end

get_total_weight('test_inputs/day_07.txt', test_parents)
