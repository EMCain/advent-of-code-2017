def find_groups(text)
  # find groups
  groups = []
  pattern = /((?<!\!)\{)(.*)((?<!\!)\}(?=[^\}]*)$)/
  garbage_pattern = /^(\<.*\>(?=[^\>]*))$/
  match = pattern.match(text)
  puts match
  while !match.nil?
    puts "text " + text
    puts "match " + match.captures.to_s
    if !garbage_pattern.match(text).nil?
      puts "garbage" + text
      text.slice! garbage_pattern.match(text)
    else
      groups.push(match.captures)
      text.slice! match.captures[0]
      text.slice! match.captures[2]
      # text = match.captures[1]
      match = pattern.match(text)
    end

  end
  puts "groups is " + groups.to_s
  puts groups.length
  return groups
end


a = '{}'
b = '{{{}}}'
c = '{{},{}}'
d = '{{{},{},{{}}}}'
e = '{<{},{},{{}}>}'
f = '{<a>,<a>,<a>,<a>}'

find_groups(a)
find_groups(b)
find_groups(c)
find_groups(d)
find_groups(e)
find_groups(f)
