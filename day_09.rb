def find_groups(text)
  # find groups
  groups = []
  pattern = /(\{)(.*)(\}(?=[^\}]*)$)/
  garbage_pattern = /(\<)(.*)(\>(?=[^\>]*)$)/
  match = pattern.match(text)
  puts match
  while !match.nil?
    if !garbage_pattern.match(text)
      groups.push(match.captures)
    end
    text.slice! match.captures[0]
    text.slice! match.captures[2]
    puts "sliced " + text
    match = pattern.match(text)
  end
  puts "groups is " + groups.to_s
  puts groups.length
  # for each group
    # call itself
  # and there should probably be some kind of counter for the score??? idk
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
